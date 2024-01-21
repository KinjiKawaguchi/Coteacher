import React, { ReactElement, useEffect, useState } from 'react';
import {
  ChakraProvider,
  Container,
  HStack,
  useDisclosure,
  Modal,
} from '@chakra-ui/react';
import theme from '@/theme';
import ClassBox from './ClassBox/';
import withAuthAndAccountCheck from '@/libs/utils/HOC/withAccountCheck';
import UserHeader from '@/components/layout/header/UserHeader';
import { useRouter } from 'next/navigation';
import { Class } from '@/interfaces';

interface CommonDashBoardProps {
  fetchClasses: () => Promise<Class[]>;
  ModalContent: ({
    onClose,
    fetchClasses,
    setClasses,
  }: {
    onClose: () => void;
    fetchClasses: () => Promise<Class[]>;
    setClasses: React.Dispatch<React.SetStateAction<Class[]>>;
  }) => ReactElement;
}

const CommonDashBoard: React.FC<CommonDashBoardProps> = ({
  fetchClasses,
  ModalContent,
}) => {
  const router = useRouter();
  const [classes, setClasses] = useState<Class[]>([]);
  const { isOpen, onOpen, onClose } = useDisclosure();

  useEffect(() => {
    fetchClasses().then(setClasses);
  }, [fetchClasses]);

  const handleClassClick = (classId: string) => {
    if (classId) {
      router.push(`/DashBoard/${classId}`);
    } else {
      onOpen();
    }
  };

  return (
    <ChakraProvider theme={theme}>
      <Container maxWidth="container.sm">
        <UserHeader />
        <HStack spacing={4} wrap="wrap" justify="center">
          {classes.map(cls => (
            <ClassBox
              key={cls.id}
              classData={cls}
              onClick={() => handleClassClick(cls.id)}
            />
          ))}
          <ClassBox
            key="new-class"
            classData={{ id: '', name: '', teacherId: '' }}
            onClick={() => handleClassClick('')}
          />
          <Modal isOpen={isOpen} onClose={onClose} isCentered>
            <ModalContent
              onClose={onClose}
              fetchClasses={fetchClasses}
              setClasses={setClasses}
            />
          </Modal>
        </HStack>
      </Container>
    </ChakraProvider>
  );
};

export default withAuthAndAccountCheck(CommonDashBoard);
