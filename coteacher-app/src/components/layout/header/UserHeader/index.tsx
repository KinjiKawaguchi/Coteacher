import { ChakraProvider, Flex, HStack } from '@chakra-ui/react';
import React from 'react';
import UserCard from '@/components/elements/UserCard';
import theme from '@/theme'; // Import custom theme

const UserHeader = () => {
  return (
    <ChakraProvider theme={theme}>
      <Flex direction="column" justify="left" align="left">
        <HStack marginTop="4">
          <UserCard />
        </HStack>
      </Flex>
    </ChakraProvider>
  );
};

export default UserHeader;
