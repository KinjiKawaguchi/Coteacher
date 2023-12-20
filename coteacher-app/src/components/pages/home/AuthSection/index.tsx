import React, { useState } from 'react';
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
import TeacherAuth from '@/components/layout/modal/TeacherAuth';
import app_icon from '@/images/app-icon.svg';
import theme from '@/theme';
import { useRouter } from 'next/navigation';

type AuthSectionProps = {
  isAuthenticated: boolean;
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
}) => {
  const router = useRouter();
  const [isTeacherModalOpen, setIsTeacherModalOpen] = useState(false);
  const [isStudentModalOpen, setIsStudentModalOpen] = useState(false);

  const handleTeacherModalOpen = () => setIsTeacherModalOpen(true);
  const handleTeacherModalClose = () => setIsTeacherModalOpen(false);

  const handleStudentModalOpen = () => setIsStudentModalOpen(true);
  const handleStudentModalClose = () => setIsStudentModalOpen(false);

  const navDashBoard = () => {
    router.push('/DashBoard');
  };

  return (
    <ChakraProvider theme={theme}>
      <Flex>
        <VStack spacing={8} textAlign="center">
          <Image src={app_icon.src} alt="logo" boxSize="200px" />
          <Text fontSize="4xl">あなたの先生</Text>
          {isAuthenticated ? (
            <LoginButton onClick={navDashBoard}>ダッシュボード</LoginButton>
          ) : (
            <>
              <LoginButton onClick={handleTeacherModalOpen}>先生ログイン</LoginButton>
                <Modal isOpen={isTeacherModalOpen} onClose={handleTeacherModalClose} isCentered>
                  <TeacherAuth onClose={handleTeacherModalClose} />
                </Modal>
              <LoginButton onClick={handleStudentModalOpen}>生徒ログイン</LoginButton>
              <Modal isOpen={isStudentModalOpen} onClose={handleStudentModalClose} isCentered>
                <StudentAuth onClose={handleStudentModalClose} />
              </Modal>
            </>
          )}
        </VStack>
      </Flex>
    </ChakraProvider>
  );
};

export default AuthSection;
