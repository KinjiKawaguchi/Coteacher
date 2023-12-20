import React, { useState } from 'react';
import {
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalFooter,
  ModalCloseButton,
  Button,
  VStack,
  Input,
  Flex,
  Spinner,
} from '@chakra-ui/react';
import toast from '@/libs/utils/toast';
import { CreateClassRequest } from '@/libs/services/api';

type ParticipateClassProps = {
  onClose: () => void;
  fetchClasses: () => Promise<void>;
};

export default function CreateClass({
  onClose,
  fetchClasses,
}: ParticipateClassProps) {
  const [className, setClassName] = useState('');
  const [isCreating, setIsCreating] = useState(false);

  const handleClassNameChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setClassName(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleSubmit();
    }
  };

  const handleSubmit = async () => {
    try {
      setIsCreating(true);
      const res = await CreateClassRequest(className);

      if (res.status === 200) {
        // Success toast
        toast({
          status: 'success',
          title: 'Success',
          description: res.message,
        });
        await fetchClasses();
        onClose();
      } else {
        // Error toasts
        let toastStatus = 'error';
        let toastTitle = 'Error';
        const toastDescription = res.error;

        if (res.status === 404) {
          toastTitle = '無効なコードです。';
        } else if (res.status === 409) {
          toastStatus = 'warning';
          toastTitle = '既に参加している授業です。';
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
    } catch (error) {
      console.error(error);
    } finally {
      setIsCreating(false);
    }
  };

  return (
    <>
      <ModalOverlay />
      <ModalContent>
        <ModalHeader>作成する授業名を入力</ModalHeader>
        <ModalCloseButton onClick={onClose} />
        <ModalBody>
          <VStack spacing={8} textAlign="center">
            <Input
              placeholder="授業名"
              value={className}
              onChange={handleClassNameChange}
              onKeyDown={handleKeyDown}
            />
          </VStack>
        </ModalBody>
        <ModalFooter>
          <Flex w="100%" justifyContent="center">
            {isCreating ? (
              <Spinner />
            ) : (
              <Button
                colorScheme="teal"
                size="lg"
                borderRadius="30px"
                variant="outline"
                onClick={handleSubmit}
              >
                授業を作成
              </Button>
            )}
          </Flex>
        </ModalFooter>
      </ModalContent>
    </>
  );
}
