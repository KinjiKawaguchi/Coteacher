import React from 'react';
import { Box, Text, useColorModeValue } from '@chakra-ui/react';
import { FaBook } from 'react-icons/fa';
import { useRouter } from 'next/navigation';

interface ClassBoxProps {
  id: string;
  name: string;
  onClick?: () => void; // Add this line
}

const ClassBox = ({ id, name, onClick }: ClassBoxProps) => {
  const router = useRouter();

  const bg = useColorModeValue('gray.100', 'gray.600');
  const text = useColorModeValue('gray.800', 'white');

  const handleClick = () => {
    onClick?.();
    router.push(`/Classes/${id}`); // 正しいパス
  };

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
      onClick={handleClick}
    >
      <FaBook size="24" color={text} aria-hidden="true" />
      <Text fontSize={['md', 'lg']} fontWeight="bold" color={text}>
        {name || '+'}
      </Text>
    </Box>
  );
};

export default ClassBox;
