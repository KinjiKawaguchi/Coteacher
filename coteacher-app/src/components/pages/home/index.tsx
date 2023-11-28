'use client';
import React from 'react';
import {
  ChakraProvider,
  Button,
  Text,
  VStack,
  Image,
  Center,
  Container,
  useDisclosure,
  Modal,
} from '@chakra-ui/react';
import theme from '@/theme'; // Import custom theme
import LargeWithAppLinksAndSocial from '@/components/footer';
import app_icon from '@/images/app-icon.svg';
import StudentAuth from '@/components/modal/StudentAuth';

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

// IntroductionText コンポーネントは再利用可能で、テキストの内容を変更できます
const IntroductionText = () => (
  <Text mt={4} align="left" letterSpacing={3}>
    あなたの先生は、先生一人では教えきれない授業で、先生が用意したフォーマットに従って、生徒が疑問を解決できるサービスです。
    <br />
    <br />
    生徒は、自分の学習スタイルやペースに合わせて、先生の用意した人工知能と対話することができます。
    <br />
    <br />
    私たちの目標は、教育の質を高め、学習の可能性を広げることです。あなたの先生を通じて、学習はもっと柔軟で、アクセスしやすく、そして楽しくなります。
  </Text>
);

const HomeView = () => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <ChakraProvider theme={theme}>
      <Center py={5}>
        <Container maxWidth="container.sm">
          <VStack spacing={8} textAlign="center">
            <Image src={app_icon.src} alt="logo" boxSize="200px" />
            <Text fontSize="4xl">あなたの先生</Text>
            <LoginButton>先生ログイン</LoginButton>
            <LoginButton onClick={onOpen}>生徒ログイン</LoginButton>
            <Modal isOpen={isOpen} onClose={onClose} isCentered>
              <StudentAuth onClose={onClose} />
            </Modal>
            <IntroductionText />
          </VStack>
        </Container>
      </Center>
      <LargeWithAppLinksAndSocial />
    </ChakraProvider>
  );
};

export default HomeView;
