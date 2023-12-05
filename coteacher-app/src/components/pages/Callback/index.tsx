'use client';

import React, { useEffect, useState } from 'react';
import { isSignInWithEmailLink, signInWithEmailLink } from 'firebase/auth';
import { useRouter } from 'next/navigation';
import { auth } from '@/libs/utils/auth/FirebaseConfig';

function Callback() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLoading(true); // 認証プロセス開始時にisLoadingをtrueに設定
    if (isSignInWithEmailLink(auth, window.location.href)) {
      const email = window.localStorage.getItem('email');
      if (email) {
        signInWithEmailLink(auth, email, window.location.href)
          .then(() => {
            window.localStorage.removeItem('email');
            router.push('/RegisterStudent');
          })
          .catch((error) => {
            console.error('Error signing in with email link', error);
          })
          .finally(() => {
            setIsLoading(false); // 認証プロセス終了時にisLoadingをfalseに設定
          });
      } else {
        console.log('No email saved in localStorage');
        setIsLoading(false);
      }
    } else {
      setIsLoading(false);
    }
  }, []);

  if (isLoading) {
    return <div>読込中...</div>; // 直接「読込中」と表示
  }
}

export default Callback;
