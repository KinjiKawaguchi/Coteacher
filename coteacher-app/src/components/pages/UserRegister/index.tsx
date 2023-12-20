'use client';

import React, { useState, ChangeEvent, KeyboardEvent, FC } from 'react';
import { useRouter } from 'next/navigation';
import withAuth from '@/libs/utils/HOC/withAuth';
import theme from '@/theme';
import { createStudent } from '@/libs/services/api';
import {
  ChakraProvider,
  Button,
  Text,
  VStack,
  Flex,
  Input,
  Spinner,
} from '@chakra-ui/react';

const RegisterStudentView: FC = () => {
  const [name, setName] = useState<string>('');
  const [isRegistering, setIsRegistering] = useState<boolean>(false);
  const [isRegistered, setIsRegistered] = useState<boolean>(false);
  const navigator = useRouter();

  const handleNameChange = (event: ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
  };

  const handleKeyDown = (event: KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter' && !isRegistered) {
      handleRegister();
    }
  };

  const handleRegister = async () => {
    if (name.length >= 4 && !isRegistered) {
      try {
        setIsRegistering(true);
        const res = await createStudent(name);
        if (res && res.ok) {
          setIsRegistered(true);
          navigator.push('/DashBoard');
        }
      } catch (error) {
        console.error(error);
      } finally {
        setIsRegistering(false);
      }
    }
  };

  return (
    <ChakraProvider theme={theme}>
      <VStack spacing={8} textAlign="center">
        <Text>名前を入れて始めましょう。</Text>
        <Flex>
          <Input
            placeholder="名前"
            value={name}
            onChange={handleNameChange}
            onKeyDown={handleKeyDown}
            disabled={isRegistered}
          />
        </Flex>
        {isRegistering ? (
          <Spinner />
        ) : (
          <Button onClick={handleRegister} disabled={isRegistered}>
            登録
          </Button>
        )}
      </VStack>
    </ChakraProvider>
  );
};

export default withAuth(RegisterStudentView);
