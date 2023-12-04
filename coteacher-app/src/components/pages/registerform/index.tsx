import React from 'react';
import {
  ChakraProvider,
  Button,
  Text,
  VStack,
  Flex,
  Input,
} from '@chakra-ui/react';
import withAuth from '@/libs/utils/HOC/withAuth';
import theme from '@/theme';

function RegisterStudentView() {
  return (
    <ChakraProvider theme={theme}>
      <VStack spacing={8} textAlign="center">
        <Text>名前を入れて始めましょう。</Text>
        <Flex>
          <Input placeholder="名前" />
        </Flex>
        <Button>登録</Button>
      </VStack>
    </ChakraProvider>
  );
}

export default withAuth(RegisterStudentView);
