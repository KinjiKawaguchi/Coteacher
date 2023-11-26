import React from "react";
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
  Image,
  Flex
} from "@chakra-ui/react";
import app_icon from '@/images/app-icon.svg';

type StudentAuthProps = {
  onClose: () => void;
};

export default function StudentAuth({ onClose }: StudentAuthProps) {
  return (
    <>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>
          <VStack spacing={8} textAlign="center">
            <Image src={app_icon.src} alt="logo" boxSize="100px" />
            <Text fontSize="4xl">生徒ログイン</Text>
          </VStack>
        </ModalHeader>
        <ModalCloseButton onClick={onClose} />
        <ModalBody>
          <VStack spacing={8} textAlign="center">
            <Text>入力したメールアドレスにログイン用リンクが送信されます。</Text>
            <Input placeholder="メールアドレス" />
          </VStack>
        </ModalBody>
        <ModalFooter>
          <Flex w="100%" justifyContent="center">
          <Button colorScheme="teal" size="lg" borderRadius="30px" variant="outline">
            ログインリンクを送信
          </Button>
          </Flex>
        </ModalFooter>
      </ModalContent>
    </>
  );
}
