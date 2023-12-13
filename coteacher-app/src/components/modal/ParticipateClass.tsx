import React, { useState } from 'react';
import {
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  ModalCloseButton,
  Text,
  Button,
  VStack,
  Input,
  Flex,
} from '@chakra-ui/react';
import toast from '@/libs/utils/toast';
import { participateClass } from '@/libs/services/api';

type ParticipateClassProps = {
  onClose: () => void;
};

export default function ParticipateClass({ onClose }: ParticipateClassProps) {
  const [invitationCode, setInvitationCode] = useState('');

  const handleInvitationCodeChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setInvitationCode(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleSubmit();
    }
  };

  const handleSubmit = async () => {
    const res = await participateClass(invitationCode);

    if (res.status === 200) {
      // Success toast
      toast({
        status: 'success',
        title: 'Success',
        description: res.message,
      });
    } else {
      // Error toasts
      let toastStatus = 'error';
      let toastTitle = 'Error';
      let toastDescription = res.error;

      if (res.status === 404) {
        toastTitle = 'Invalid Code';
      } else if (res.status === 409) {
        toastStatus = 'warning';
        toastTitle = 'Already Participated';
      } else if (res.status === 500) {
        toastTitle = 'Server Error';
      } else if (res.status === 'network-error') {
        toastTitle = 'Network Error';
      }

      toast({
        status: toastStatus as 'success' | 'error' | 'warning' | 'info',
        title: toastTitle,
        description: toastDescription,
      });
    }
  };

  return (
    <>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>授業参加コードを入力</ModalHeader>
        <ModalCloseButton onClick={onClose} />
        <ModalBody>
          <VStack spacing={8} textAlign="center">
            <Input
              placeholder="授業参加コード"
              value={invitationCode}
              onChange={handleInvitationCodeChange}
              onKeyDown={handleKeyDown}
            />
          </VStack>
        </ModalBody>
        <ModalFooter>
          <Flex w="100%" justifyContent="center">
            <Button
              colorScheme="teal"
              size="lg"
              borderRadius="30px"
              variant="outline"
              onClick={handleSubmit}
            >
              授業に参加
            </Button>
          </Flex>
        </ModalFooter>
      </ModalContent>
    </>
  );
}
