"use client";

import React from 'react';
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
} from '@chakra-ui/react';

function RegisterStudentView() {
  const [name, setName] = React.useState('');
  const navigator = useRouter();

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleRegister();
    }
  };

  const handleRegister = async () => {
    if (name.length >= 4) {
      try {
        console.log(name);
        const res = await createStudent(name);
        if (res.status === 200) {
          console.log(res);
          navigator.push('/ClassSelect');
        }
      } catch (error) {
        console.log(error);
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
          />
        </Flex>
        <Button onClick={handleRegister}>登録</Button>
      </VStack>
    </ChakraProvider>
  );
}

export default withAuth(RegisterStudentView);
