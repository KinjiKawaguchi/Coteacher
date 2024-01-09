import React from 'react';
import { Box, HStack, Text, useColorModeValue } from '@chakra-ui/react';
import { FaBook } from 'react-icons/fa';
import ClassBoxPopUp from './ClassBoxPopUp';

interface ClassBoxProps {
  id: string;
  name: string;
  onClick?: () => void;
}

const ClassBox = ({ id, name, onClick }: ClassBoxProps) => {
  const bg = useColorModeValue('gray.100', 'gray.600');
  const text = useColorModeValue('gray.800', 'white');

  return (
    <Box
      p={4}
      w={['150px', '200px']} // 幅を固定
      h="auto" // 高さは内容に応じて自動調整
      bg={bg}
      boxShadow="md"
      borderRadius="lg"
      m={2}
      textAlign="center"
      _hover={{ boxShadow: 'lg', transform: 'scale(1.05)' }}
      role="group"
      onClick={onClick}
    >
      <HStack justifyContent="space-between" mb={2}>
        <FaBook size="24" color={text} aria-hidden="true" />
        <ClassBoxPopUp classId={id} />
      </HStack>
      <Text
        fontSize={['md', 'lg']}
        fontWeight="bold"
        color={text}
        noOfLines={1}
        isTruncated
      >
        {name || '+'}
      </Text>
    </Box>
  );
};

export default ClassBox;
