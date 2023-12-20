'use client';

import React, { useEffect, useState } from 'react';
import {
  ChakraProvider,
  Container,
  HStack,
  useDisclosure,
  Modal,
} from '@chakra-ui/react';
import theme from '@/theme';
import ClassBox from './ClassBox';
import withAuthAndAccountCheck from '@/libs/utils/HOC/withAccountCheck';
import { getParticipatingClass } from '@/libs/services/api';
import ParticipateClass from '@/components/layout/modal/ParticipateClass';
import UserHeader from '@/components/layout/header/UserHeader';

type Class = {
  ID: string;
  Name: string;
  TeacherID: string;
};

const StudentDashBoard = () => {
  const [classes, setClasses] = useState<Class[]>([]);

  const fetchClasses = async () => {
    const { classes: resClasses } = await getParticipatingClass();
    setClasses(resClasses || []);
  };

  useEffect(() => {
    fetchClasses();
  }, []);

  const memoizedClasses = React.useMemo(() => classes, [classes]);

  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <ChakraProvider theme={theme}>
      <Container maxWidth="container.sm">
        <UserHeader />
        <HStack spacing={4} wrap="wrap" justify="center">
          {memoizedClasses.map(cls => (
            <ClassBox key={cls.ID} id={cls.ID} name={cls.Name} />
          ))}
          <ClassBox id="new-class" name="" onClick={onOpen} />{' '}
          <Modal isOpen={isOpen} onClose={onClose} isCentered>
            <ParticipateClass onClose={onClose} fetchClasses={fetchClasses} />
          </Modal>
        </HStack>
      </Container>
    </ChakraProvider>
  );
};

export default withAuthAndAccountCheck(StudentDashBoard);
