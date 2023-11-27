import React, { useState } from "react";
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
  useToast
} from "@chakra-ui/react";


type StudentAuthProps = {
  onClose: () => void;
};

export default function StudentAuth({ onClose }: StudentAuthProps) {
  const [email, setEmail] = useState("");
  const toast = useToast(); // トースト通知のためのフック

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(event.target.value);
  };

  const checkAccountExistence = async (email: string) => {
    try {
      console.log(email);
      const response = await fetch(`https://api-image-pgfe7sqiia-an.a.run.app/CheckAcountExist/${email}`);
      // ログ出力      
      console.log(response);

      // APIが404を返した場合、アカウントが存在しないと判断
      if (response.status === 404) {
        return false;
      }

      // その他の応答は成功と見なし、アカウントが存在すると判断
      const data: any = await response.json();
      return data.exist;
    } catch (error) {
      toast({
        title: "エラー",
        description: "アカウントの確認中にエラーが発生しました。",
        status: "error",
        duration: 3000,
        isClosable: true,
      });
      return false; // エラーが発生した場合はfalseを返す
    }
  };

  const handleSubmit = async () => {
    // メールアドレスの形式を確認する処理を追加 (オプショナル)
    if (!email.match(/^[^\s@]+@[^\s@]+\.[^\s@]+$/)) {
      toast({
        title: "無効なメールアドレス",
        description: "正しいメールアドレスを入力してください。",
        status: "warning",
        duration: 3000,
        isClosable: true,
      });
      return;
    }

    // アカウントの存在を確認
    const accountExists = await checkAccountExistence(email);
    if (accountExists) {
      // Firebase Authにログインリンクを送信
      
      toast({
        title: "ログインリンクを送信",
        description: "入力したメールアドレスにログイン用リンクを送信しました。",
        status: "success",
        duration: 3000,
        isClosable: true,
      });
        } else {

    }
  };
  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleSubmit();
    }
  };

  return (
    <>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader> {/* 他のコンポーネントは変更なし */} </ModalHeader>
        <ModalCloseButton onClick={onClose} />
        <ModalBody>
          <VStack spacing={8} textAlign="center">
            <Text>入力したメールアドレスにログイン用リンクが送信されます。</Text>
            <Input 
              placeholder="example@example.test" 
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
