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
import { studentRepo } from '@/repository/student';
import { Class } from '@/interfaces';

type ParticipateClassProps = {
  onClose: () => void;
  fetchClasses: () => Promise<Class[]>;
  setClasses: React.Dispatch<React.SetStateAction<Class[]>>;
};

export default function ParticipateClass({
  onClose,
  fetchClasses,
  setClasses,
}: ParticipateClassProps) {
  const [invitationCode, setInvitationCode] = useState('');
  const [isParticipating, setIsParticipating] = useState<boolean>(false);

  const handleInvitationCodeChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setInvitationCode(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter' && invitationCode === event.currentTarget.value) {
      handleSubmit();
    }
  };

  const handleSubmit = async () => {
    try {
      setIsParticipating(true);
      const res = await studentRepo.participateClass(invitationCode);

      if (res) {
        // Success toast
        toast({
          status: 'success',
          title: 'Success',
          description: res.name + 'に参加しました。',
        });

        const classes = await fetchClasses();
        setClasses(classes);
        onClose();
      }
      // TODO: ここを動作させる
      // } else {
      //   // Error toasts
      //   let toastStatus = 'error';
      //   let toastTitle = 'Error';
      //   const toastDescription = res.error;

      //   if (res.status === 404) {
      //     toastTitle = '無効なコードです。';
      //   } else if (res.status === 409) {
      //     toastStatus = 'warning';
      //     toastTitle = '既に参加している授業です。';
      //   } else if (res.status === 500) {
      //     toastTitle = 'Server Error';
      //   } else if (res.status === 'network-error') {
      //     toastTitle = 'Network Error';
      //   }

      //   toast({
      //     status: toastStatus as 'success' | 'error' | 'warning' | 'info',
      //     title: toastTitle,
      //     description: toastDescription,
      //   });
    } catch (error) {
      console.error(error);
    } finally {
      setIsParticipating(false);
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
            {isParticipating ? (
              <Spinner />
            ) : (
              <Button
                colorScheme="teal"
                size="lg"
                borderRadius="30px"
                variant="outline"
                onClick={handleSubmit}
              >
                授業に参加
              </Button>
            )}
          </Flex>
        </ModalFooter>
      </ModalContent>
    </>
  );
}
