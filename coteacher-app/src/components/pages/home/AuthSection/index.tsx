import React from 'react';
import { Spinner, Flex, VStack } from '@chakra-ui/react';
import { Button } from '@/components/ui/button';
import StudentAuthDialog from './StudentAuthDialog';
import TeacherAuthDialog from './TeacherAuthDialog';
import { useRouter } from 'next/navigation';

type AuthSectionProps = {
  isAuthenticated: boolean | null;
};

const AuthSection: React.FC<AuthSectionProps> = ({ isAuthenticated }) => {
  const router = useRouter();

  const navDashBoard = () => {
    router.push('/DashBoard');
  };

  if (isAuthenticated === null) {
    return <Spinner />;
  }

  return (
    <Flex>
      <VStack spacing={8} textAlign="center">
        {isAuthenticated ? (
          <Button onClick={navDashBoard}>ダッシュボード</Button>
        ) : (
          <>
            <TeacherAuthDialog />
            <StudentAuthDialog />
          </>
        )}
      </VStack>
    </Flex>
  );
};

export default AuthSection;
