import React from 'react';
import { Box, Text, useColorModeValue } from '@chakra-ui/react';
import { FaBook } from 'react-icons/fa';

interface ClassBoxProps {
  name: string;
  onClick?: () => void; // Add this line
}

const ClassBox = ({ name, onClick }: ClassBoxProps) => {
  const bg = useColorModeValue('gray.100', 'gray.600');
  const text = useColorModeValue('gray.800', 'white');

  return (
    <Box
      p={4}
      w={['150px', '200px']}
      bg={bg}
      boxShadow="md"
      borderRadius="lg"
      m={2}
      textAlign="center"
      _hover={{ boxShadow: 'lg', transform: 'scale(1.05)' }}
      role="group"
      onClick={onClick}
    >
      <FaBook size="24" color={text} aria-hidden="true" />
      <Text fontSize={['md', 'lg']} fontWeight="bold" color={text}>
        {name || '+'}
      </Text>
    </Box>
  );
};

export default ClassBox;
