'use client';

import React, { useEffect, useState } from 'react';
import { ChakraProvider, Center, Container, HStack } from '@chakra-ui/react';
import theme from '@/theme';
import ClassBox from './ClassBox';
import withAuthAndAccountCheck from '@/libs/utils/HOC/withAccountCheck';
import { getParticipatingClass } from '@/libs/services/api';

type Class = {
  ID: string;
  Name: string;
  TeacherID: string;
};

const ClassSelectView = () => {
  const [classes, setClasses] = useState<Class[]>([]);

  useEffect(() => {
    const fetchClasses = async () => {
      const res = await getParticipatingClass();
      if (res && res.classes) {
        setClasses(res.classes);
      }
    };
    fetchClasses();
  }, []);

  return (
    <ChakraProvider theme={theme}>
      <Center py={5}>
        <Container>
          <HStack spacing={4} wrap="wrap" justify="center">
            {classes.map(cls => (
              <ClassBox key={cls.ID} name={cls.Name} />
            ))}
            <ClassBox key="new-class" name="" />
          </HStack>
        </Container>
      </Center>
    </ChakraProvider>
  );
};

export default withAuthAndAccountCheck(ClassSelectView);
