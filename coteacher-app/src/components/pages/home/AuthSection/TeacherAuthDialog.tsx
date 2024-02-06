import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Alert, AlertDescription } from '@/components/ui/alert';

import { EMAIL_REGEX } from '@/constants';
import { sendEmailLink } from '@/libs/utils/auth/auth';
import toast from '@/libs/utils/toast';
import { teacherRepo } from '@/repository/teacher';
import EmailInput from './EmailInput';
import { Button } from '@/components/ui/button';

const TeacherAuthDialog = () => {
  const handleSubmit = async (email: string) => {
    if (!EMAIL_REGEX.test(email)) {
      toast({
        status: 'warning',
        title: '無効なメールアドレスです',
        description: '正しいメールアドレスを入力してください',
      });
      return;
    }
    const isTeacherExists = await teacherRepo.checkTeacherExistsByEmail(email);
    if (isTeacherExists) {
      sendEmailLink(email, 'Teacher');
    } else {
      toast({
        status: 'warning',
        title: 'アカウントが存在しません',
        description: '先生用アカウントの新規登録は現在受け付けておりません',
      });
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>先生向け認証</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>先生用ログイン</DialogTitle>
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
