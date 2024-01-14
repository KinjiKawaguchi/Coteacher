import React from 'react';
import { ChakraProvider, Flex, Text, VStack, Image } from '@chakra-ui/react';
import { Button } from '@/components/ui/button';
import StudentAuthDialog from './StudentAuthDialog';
import TeacherAuthDialog from './TeacherAuthDialog';
import app_icon from '@/images/app-icon.svg';
import theme from '@/theme';
import { useRouter } from 'next/navigation';

type AuthSectionProps = {
  isAuthenticated: boolean;
};

const AuthSection: React.FC<AuthSectionProps> = ({ isAuthenticated }) => {
  const router = useRouter();

  const navDashBoard = () => {
    router.push('/DashBoard');
  };

  return (
    <ChakraProvider theme={theme}>
      <Flex>
        <VStack spacing={8} textAlign="center">
          <Image src={app_icon.src} alt="logo" boxSize="200px" />
          <Text fontSize="4xl">あなたの先生</Text>
          {isAuthenticated ? ( // TODO: ここのチラツキをなくす
            <Button onClick={navDashBoard}>ダッシュボード</Button>
          ) : (
            <>
              <TeacherAuthDialog />
              <StudentAuthDialog />
            </>
          )}
        </VStack>
      </Flex>
    </ChakraProvider>
  );
};

export default AuthSection;
