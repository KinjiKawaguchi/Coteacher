'use client';

import React, { useEffect, useState } from 'react';
import { isSignInWithEmailLink, signInWithEmailLink } from 'firebase/auth';
import { useRouter } from 'next/navigation';
import { auth } from '@/libs/utils/auth/FirebaseConfig';
import { userRepo } from '@/repository/user';

function Callback() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLoading(true); // 認証プロセス開始時にisLoadingをtrueに設定
    if (isSignInWithEmailLink(auth, window.location.href)) {
      const email = window.localStorage.getItem('login-email');
      if (email) {
        signInWithEmailLink(auth, email, window.location.href)
          .then(async () => {
            window.localStorage.setItem('email', email);
            window.localStorage.removeItem('login-email');
            const isUserExists = await userRepo.checkUserExistsByEmail(email);
            registerLocalStorage();
            if (isUserExists) {
              router.push('/DashBoard');
            } else {
              router.push('/UserRegister');
            }
          })
          .catch(error => {
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
  }, [router]); // Include the 'router' dependency in the dependency array

  if (isLoading) {
    return <div>読込中...</div>; // 直接「読込中」と表示
  }
  function registerLocalStorage() {
    userRepo.getUserByEmail(localStorage.getItem('email') || '').then(user => {
      if (user) {
        localStorage.setItem('UserID', user.id);
        localStorage.setItem('UserEmail', user.email);
        localStorage.setItem('UserName', user.name);
        localStorage.setItem('UserType', user.user_type.toString());
      }
    });
  }
}

export default Callback;
