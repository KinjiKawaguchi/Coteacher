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
    const auth = getAuth();
    await sendSignInLinkToEmail(auth, email, actionCodeSettings);
    window.localStorage.setItem('login-email', email);
    toast({
      status: 'success',
      title: 'メールを送信しました',
      description: 'メールを確認してログインしてください。',
    });
  } catch (error) {
    console.error('Error sending sign-in email link:', error);
    toast({ status: 'error', title: 'エラーが発生しました' });
  }
};
