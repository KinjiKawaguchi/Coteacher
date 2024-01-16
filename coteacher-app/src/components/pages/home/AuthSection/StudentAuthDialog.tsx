import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';

import { EMAIL_REGEX } from '@/constants';
import { sendEmailLink } from '@/libs/utils/auth/auth';
import toast from '@/libs/utils/toast';
import EmailInput from './EmailInput';

const TeacherAuthDialog = () => {
  const handleSubmit = async (email: string) => {
    if (!EMAIL_REGEX.test(email)) {
      toast({
        status: 'warning',
        title: '無効なメールアドレス',
        description: '正しいメールアドレスを入力してください。',
      });
      return;
    }
    await sendEmailLink(email, 'Student');
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>生徒向け認証</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>生徒用ログイン</DialogTitle>
          <DialogDescription>
            入力されたメールに認証リンクを送信します。
          </DialogDescription>
        </DialogHeader>
        <EmailInput handleSubmit={handleSubmit} />
      </DialogContent>
    </Dialog>
  );
};

export default TeacherAuthDialog;
