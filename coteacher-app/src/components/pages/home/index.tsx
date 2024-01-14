'use client';
import React, { useState, useEffect } from 'react';
import {
  ChakraProvider,
  Text,
  VStack,
  Container,
  Image,
} from '@chakra-ui/react';
import theme from '@/theme'; // Import custom theme
import LargeWithAppLinksAndSocial from '@/components/layout/footer';
import AuthSection from '@/components/pages/home/AuthSection';
import { auth } from '@/libs/utils/auth/FirebaseConfig'; // Import Firebase config
import app_icon from '@/images/app-icon.svg';

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
  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);

  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged(user => {
      setIsAuthenticated(!!user); // userが存在する場合はtrue, そうでなければfalse
    });

    return () => unsubscribe();
  }, []);

  return (
    <ChakraProvider theme={theme}>
      <Container maxWidth="container.sm">
        <VStack spacing={8} textAlign="center">
          <Image src={app_icon.src} alt="logo" boxSize="200px" />
          <Text fontSize="4xl">あなたの先生</Text>
          <AuthSection isAuthenticated={isAuthenticated} />
          <IntroductionText />
        </VStack>
      </Container>
      <LargeWithAppLinksAndSocial />
    </ChakraProvider>
  );
};

export default HomeView;
