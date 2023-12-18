import React from 'react';
import {
  ChakraProvider,
  Button,
  Flex,
  Text,
  VStack,
  Image,
  Modal,
} from '@chakra-ui/react';
import StudentAuth from '@/components/layout/modal/StudentAuth';
import app_icon from '@/images/app-icon.svg';
import theme from '@/theme';
import { useRouter } from 'next/navigation';

type AuthSectionProps = {
  isAuthenticated: boolean;
  onOpen: () => void;
  onClose: () => void;
  isOpen: boolean;
};

// ボタンのプロパティの型定義
type LoginButtonProps = {
  children: React.ReactNode;
  onClick?: () => void; // ボタンクリックのためのオプショナルなonClickハンドラーを追加
};

// ログインボタンコンポーネント
const LoginButton: React.FC<LoginButtonProps> = ({ children, onClick }) => (
  <Button
    onClick={onClick}
    colorScheme="teal"
    size="lg"
    borderRadius="30px"
    variant="outline"
  >
    {children}
  </Button>
);

const AuthSection: React.FC<AuthSectionProps> = ({
  isAuthenticated,
  onOpen,
  onClose,
  isOpen,
}) => {
  const router = useRouter();

  const navClassSelect = () => {
    router.push('/ClassSelect');
  };

  return (
    <ChakraProvider theme={theme}>
      <Flex>
        <VStack spacing={8} textAlign="center">
          <Image src={app_icon.src} alt="logo" boxSize="200px" />
          <Text fontSize="4xl">あなたの先生</Text>
          {isAuthenticated ? (
            <LoginButton onClick={navClassSelect}>ダッシュボード</LoginButton>
          ) : (
            <>
              <LoginButton>先生ログイン</LoginButton>
              <LoginButton onClick={onOpen}>生徒ログイン</LoginButton>
              <Modal isOpen={isOpen} onClose={onClose} isCentered>
                <StudentAuth onClose={onClose} />
              </Modal>
            </>
          )}
        </VStack>
      </Flex>
    </ChakraProvider>
  );
};

export default AuthSection;
