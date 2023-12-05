import { getAuth, sendSignInLinkToEmail } from 'firebase/auth';
import toast from '@/libs/utils/toast/index';
import { LOGIN_LINK_URL } from '@/constants';
import '@/libs/utils/auth/FirebaseConfig';

export const sendEmailLink = async (email: string) => {
  const actionCodeSettings = {
    url: LOGIN_LINK_URL,
    handleCodeInApp: true,
  };

  try {
    console.log(email);
    const auth = getAuth();
    await sendSignInLinkToEmail(auth, email, actionCodeSettings)
      .then(() => {
        window.localStorage.setItem('email', email);
      })
      .catch((error) => {
        toast({ status: 'error', title: 'エラーが発生しました' });
        return;
      });
    toast({
      status: 'success',
      title: 'メールを送信しました',
      description: 'メールを確認してログインしてください。',
    });
    window.localStorage.setItem('emailForSignIn', email);
  } catch (error) {
    console.error(error);
    toast({ status: 'error', title: 'エラーが発生しました' });
  }
};
