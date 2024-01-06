import React from 'react';
import {
  Alert,
  AlertIcon,
  Button,
  ModalOverlay,
  ModalContent,
  ModalCloseButton,
  ModalHeader,
  ModalBody,
  ModalFooter,
  Flex,
  VStack,
  Input,
} from '@chakra-ui/react';

import { EMAIL_REGEX } from '@/constants';
import { sendEmailLink } from '@/libs/utils/auth/auth';
import toast from '@/libs/utils/toast';
import { teacherRepo } from '@/repository/teacher';

type TeacherAuthProps = {
  onClose: () => void;
};

export default function TeacherAuth({ onClose }: TeacherAuthProps) {
  const [email, setEmail] = React.useState('');

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key == 'Enter') {
      handleSendEmailLink();
    }
  };

  const handleSendEmailLink = async () => {
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
    <>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>先生ログイン</ModalHeader>
        <ModalCloseButton onClick={onClose} />
        <ModalBody>
          <VStack spacing={8} textAlign="center">
            入力したメールアドレスにログインリンクを送信します
            <Alert status="info">
              <AlertIcon />
              現在、あなたの先生では先生アカウントの新規登録を受け付けておりません。
            </Alert>
            <Input
              placeholder="example@example.com"
              value={email}
              onChange={handleEmailChange}
              onKeyDown={handleKeyDown}
            />
          </VStack>
        </ModalBody>
        <ModalFooter>
          <Flex w="100%" justifyContent="center">
            <Button
              colorScheme="teal"
              size="lg"
              borderRadius="30px"
              variant="outline"
              onClick={handleSendEmailLink}
            >
              ログインリンクを送信
            </Button>
          </Flex>
        </ModalFooter>
      </ModalContent>
    </>
  );
}
