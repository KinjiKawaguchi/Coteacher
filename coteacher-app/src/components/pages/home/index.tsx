import React from 'react';
import { ChakraProvider, Button, Text, VStack, Image, Center, Container } from '@chakra-ui/react';
import theme from '@/theme'; // Import custom theme
import LargeWithAppLinksAndSocial from '@/components/footer';
import app_icon from '@/images/app-icon.svg';

const HomeView = () => {
  return (
    <ChakraProvider theme={theme}>
      <Center py={5}>
        <Container maxWidth="container.sm">
          <VStack spacing={8} textAlign="center">
          <Image src={app_icon.src} alt="logo" width={200} height={200} />
          <Text fontSize="4xl">あなたの先生</Text>
          <Button colorScheme="teal" size="lg" borderRadius="30px" variant="outline">
            先生ログイン
          </Button>
          <Button colorScheme="teal" size="lg" borderRadius="30px" variant="outline">
            生徒ログイン
          </Button>
          <Text mt={4} align="left" letterSpacing={3}>
          あなたの先生は、先生一人では教えきれない授業で、先生が用意したフォーマットに従って、生徒が疑問を解決できるサービスです。
          <br />
          <br />
          生徒は、自分の学習スタイルやペースに合わせて、先生の用意した人工知能と対話することができます。 
          <br />
          <br />
          私たちの目標は、教育の質を高め、学習の可能性を広げることです。あなたの先生を通じて、学習はもっと柔軟で、アクセスしやすく、そして楽しくなります。
            
          </Text> 
          </VStack>
        </Container>
      </Center>
      <LargeWithAppLinksAndSocial />
    </ChakraProvider>
  );
}

export default HomeView;
