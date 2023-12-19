import React, { useState } from 'react';
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
} from '@chakra-ui/react';
import { EMAIL_REGEX } from '@/constants';
import { sendEmailLink } from '@/libs/utils/auth/auth';
import toast from '@/libs/utils/toast';

type StudentAuthProps = {
  onClose: () => void;
};

export default function StudentAuth({ onClose }: StudentAuthProps) {
  const [email, setEmail] = useState('');

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleSubmit();
    }
  };

  const handleSubmit = async () => {
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
