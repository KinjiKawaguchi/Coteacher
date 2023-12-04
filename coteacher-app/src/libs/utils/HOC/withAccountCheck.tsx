import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { auth } from '@/libs/utils/auth/FirebaseConfig'; // Firebase設定のインポート
import { checkStudentExist } from '@/libs/services/api'; // アカウント確認APIのインポート

const withAuthAndAccountCheck: (Component: React.FC<any>) => React.FC<any> = (
  Component
) => {
  return (props: any) => {
    const [isLoading, setIsLoading] = useState(true);
    const router = useRouter();

    useEffect(() => {
      const unsubscribe = auth.onAuthStateChanged(async (user) => {
        if (user) {
          // ユーザーがログインしている場合の処理
          if (user.email) {
            const isStudentExists = await checkStudentExist(user.email);
            if (!isStudentExists) {
              router.push('/RegisterStudent');
            } else {
              setIsLoading(false);
            }
          } else {
            // ユーザーがログインしていない場合の処理
            router.push('/');
          }
        }
      });

      // コンポーネントのアンマウント時にリスナーを解除
      return () => unsubscribe();
    }, []);

    if (isLoading) {
      return <div>Loading...</div>;
    }

    return <Component {...props} />;
  };
};

export default withAuthAndAccountCheck;
