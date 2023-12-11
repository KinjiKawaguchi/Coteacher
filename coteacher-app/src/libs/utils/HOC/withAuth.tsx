import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { auth } from '@/libs/utils/auth/FirebaseConfig'; // Firebase設定のインポート
import { checkStudentExist } from '@/libs/services/api'; // アカウント確認APIのインポート

const withAuth: (Component: React.FC<any>) => React.FC<any> = (Component) => {
  const WithAuthComponent: React.FC<any> = (props) => {
    const [isLoading, setIsLoading] = useState(true);
    const router = useRouter();

    useEffect(() => {
      const unsubscribe = auth.onAuthStateChanged(async (user) => {
        if (user?.email) {
          if (await checkStudentExist(user.email)) {
            router.push('/ClassSelect');
          }
          setIsLoading(false);
        } else {
          router.push('/');
          setIsLoading(false);
        }
      });

      // コンポーネントのアンマウント時にリスナーを解除
      return () => unsubscribe();
    }, [router]);

    if (isLoading) {
      return <div>Loading...</div>;
    }

    return <Component {...props} />;
  };

  WithAuthComponent.displayName = `withAuth(${
    Component.displayName || Component.name
  })`;

  return WithAuthComponent;
};

export default withAuth;
