package interceptor

import (
	"context"
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitializeFirebaseApp(logger *slog.Logger) *auth.Client {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./secrets/firebase-sa-key.json") // Firebaseのサービスアカウントキーへのパス
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logger.Error("Firebaseアプリの初期化に失敗しました: %v", err)
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		logger.Error("Firebase Authクライアントの初期化に失敗しました: %v", err)
	}
	return authClient
}

// AuthInterceptor は、connect 用の認証インターセプターです。
func AuthInterceptor(authClient *auth.Client) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// リクエストヘッダーからAuthorizationヘッダーを取得
			token := req.Header().Get("Authorization")

			// Bearer トークンを抽出
			const prefix = "Bearer "
			if len(token) >= len(prefix) && token[:len(prefix)] == prefix {
				token = token[len(prefix):]
			} else {
				// Authorizationヘッダーが不適切な形式の場合
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("不正なAuthorizationヘッダー"))
			}

			// JWTトークンの検証
			_, err := authClient.VerifyIDToken(ctx, token)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("トークンの検証に失敗しました: %v", err))
			}

			// 検証が成功した場合はリクエストを処理
			return next(ctx, req)
		}
	}
}
