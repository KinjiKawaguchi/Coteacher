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
import UserHeader from '@/components/layout/header/UserHeader';
import CreateClass from '@/components/layout/modal/CreateClass';
import { classRepo } from '@/repository/class';
import { Class } from '@/interfaces';

const TeacherDashBoard = () => {
  const [classes, setClasses] = useState<Class[]>([]);

  const fetchClasses = async () => {
    const classes = await classRepo.getClassListByTeacherId(null);
    setClasses(classes || []);
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
            <ClassBox key={cls.id} name={cls.name} />
          ))}
          <ClassBox key="new-class" name="" onClick={onOpen} />{' '}
          <Modal isOpen={isOpen} onClose={onClose} isCentered>
            <CreateClass onClose={onClose} fetchClasses={fetchClasses} />
          </Modal>
        </HStack>
      </Container>
    </ChakraProvider>
  );
};

export default withAuthAndAccountCheck(TeacherDashBoard);
