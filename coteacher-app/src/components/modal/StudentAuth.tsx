import React, { useState } from 'react';
import { getAuth, sendSignInLinkToEmail } from 'firebase/auth';
import '@/utils/firebase/FirebaseConfig';
import {
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  ModalCloseButton,
  Text,
  Button,
  VStack,
  Input,
  Flex,
  useToast,
} from '@chakra-ui/react';

type StudentAuthProps = {
  onClose: () => void;
};

export default function StudentAuth({ onClose }: StudentAuthProps) {
  const [email, setEmail] = useState('');
  const toast = useToast();
  const auth = getAuth();

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleSubmit();
    }
  };

  const sendEmailLink = async (email: string) => {
    const actionCodeSettings = {
      // TODO: 本番環境URLに変更
      url: 'http://localhost:3000/finishSignUp?cartId=1234',
      handleCodeInApp: true,
    };

    try {
      console.log(email);
      await sendSignInLinkToEmail(auth, email, actionCodeSettings);
      showToast(
        'ログインリンクを送信しました',
        'メールを確認してログインしてください。',
        'success'
      );
      window.localStorage.setItem('emailForSignIn', email);
    } catch (error) {
      console.error(error); // エラー情報をコンソールに出力
      showToast('エラーが発生しました', null, 'error');
    }
  };

  const handleSubmit = async () => {
    if (!email.match(/^[^\s@]+@[^\s@]+\.[^\s@]+$/)) {
      showToast(
        '無効なメールアドレス',
        '正しいメールアドレスを入力してください。',
        'warning'
      );
      return;
    }
    await sendEmailLink(email);
  };

  const showToast = (
    title: string,
    description: React.ReactNode,
    status: 'info' | 'warning' | 'success' | 'error' | 'loading' | undefined
  ) => {
    toast({
      title: title,
      description: description,
      status: status,
      duration: 3000,
      isClosable: true,
    });
  };

  return (
    <>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>学生認証</ModalHeader>
        <ModalCloseButton onClick={onClose} />
        <ModalBody>
          <VStack spacing={8} textAlign="center">
            <Text>
              入力したメールアドレスにログイン用リンクが送信されます。
            </Text>
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
              onClick={handleSubmit}
            >
              ログインリンクを送信
            </Button>
          </Flex>
        </ModalFooter>
      </ModalContent>
    </>
  );
}
