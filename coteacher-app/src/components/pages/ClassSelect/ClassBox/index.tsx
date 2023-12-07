import React from 'react';
import { Box, Text, useColorModeValue } from '@chakra-ui/react';
import { FaBook } from 'react-icons/fa';

interface ClassBoxProps {
  name: string;
}

const ClassBox: React.FC<ClassBoxProps> = ({ name }) => {
  const bgColor = useColorModeValue('gray.100', 'gray.600');
  const textColor = useColorModeValue('gray.800', 'white');
  name = name === '' ? '+' : name;

  return (
    <Box
      p={4}
      w={['150px', '200px']} // 幅を固定
      bg={bgColor}
      boxShadow="md"
      borderRadius="lg"
      m={2}
      textAlign="center"
      _hover={{ boxShadow: 'lg', transform: 'scale(1.05)' }}
      role="group"
    >
      <FaBook size="24" color={textColor} aria-hidden="true" />
      <Text fontSize={['md', 'lg']} fontWeight="bold" color={textColor}>
        {name}
      </Text>
    </Box>
  );
};

export default ClassBox;
