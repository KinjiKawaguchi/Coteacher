import { sendSignInLinkToEmail } from 'firebase/auth';
import toast from '@/libs/utils/toast/index';
import { LOGIN_LINK_URL } from '@/constants';
import { auth } from '@/libs/utils/auth/FirebaseConfig';

export const sendEmailLink = async (email: string,userType: string) => {
  const actionCodeSettings = {
    url: LOGIN_LINK_URL,
    handleCodeInApp: true,
  };

  try {
    await sendSignInLinkToEmail(auth, email, actionCodeSettings);
    window.localStorage.setItem('login-email', email);
    window.localStorage.setItem('login-userType', userType);
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
