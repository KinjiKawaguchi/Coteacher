'use client';

import React, { useEffect, useState } from 'react';
import {
  ChakraProvider,
  Center,
  Container,
  HStack,
  Text,
} from '@chakra-ui/react';
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

  const fetchClasses = async () => {
    const { classes: resClasses } = await getParticipatingClass();
    setClasses(resClasses || []);
  };

  useEffect(() => {
    fetchClasses();
  }, []);

  const memoizedClasses = React.useMemo(() => classes, [classes]);

  return (
    <ChakraProvider theme={theme}>
      <Center py={5}>
        <Container>
          <HStack spacing={4} wrap="wrap" justify="center">
            {memoizedClasses.map(cls => (
              <ClassBox key={cls.ID} name={cls.Name} />
            ))}
            <ClassBox key="new-class" name="" />
          </HStack>
          {!memoizedClasses.length && <Text>Loading...</Text>}
        </Container>
      </Center>
    </ChakraProvider>
  );
};

export default withAuthAndAccountCheck(ClassSelectView);
