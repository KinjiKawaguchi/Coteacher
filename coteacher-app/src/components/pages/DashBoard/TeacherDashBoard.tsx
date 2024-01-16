'use client';

import React, { useEffect, useState } from 'react';
import { Container, HStack, useDisclosure, Modal } from '@chakra-ui/react';
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
    <>
      <Container maxWidth="container.sm">
        <UserHeader />
        <HStack spacing={4} wrap="wrap" justify="center">
          {memoizedClasses.map(cls => (
            <ClassBox key={cls.id} id={cls.id} name={cls.name} />
          ))}
          <ClassBox key="new-class" id="" name="" onClick={onOpen} />{' '}
          <Modal isOpen={isOpen} onClose={onClose} isCentered>
            <CreateClass onClose={onClose} fetchClasses={fetchClasses} />
          </Modal>
        </HStack>
      </Container>
    </>
  );
};

export default withAuthAndAccountCheck(TeacherDashBoard);
