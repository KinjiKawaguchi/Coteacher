import React, { useEffect, useState } from 'react';
import NotFound from '@/app/not-found';
import { classRepo } from '@/repository/class';
import { ChakraProvider, Spinner, Container } from '@chakra-ui/react';
import ClassEditView from './ClassEditView';
import theme from '@/theme';
import ClassView from './ClassView';
import UserHeader from '@/components/layout/header/UserHeader';
import { ClassHeader } from '@/components/layout/header/ClassHeader';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';
import { FormTable } from './FormTable';

export default function Class({ params }: { params: { classid: string } }) {
  const [hasEditPermission, setHasEditPermission] = useState<boolean | null>(
    null
  );
  const [hasViewPermission, setHasViewPermission] = useState<boolean | null>(
    null
  );
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      const editPermission = await classRepo.checkClassEditPermission(
        params.classid
      );
      const viewPermission = await classRepo.checkClassViewPermission(
        params.classid
      );
      setHasEditPermission(editPermission);
      setHasViewPermission(viewPermission);
      setIsLoading(false);
    };

    fetchData();
  }, [params.classid]);

  if (isLoading) {
    return <Spinner />;
  }

  const router = useRouter();

  const navDashBoard = () => {
    router.back();
  };

  return (
    <ChakraProvider theme={theme}>
      <Button onClick={navDashBoard}>←</Button>
      <Container maxWidth="container.sm">
        <UserHeader />
        <ClassHeader classId={params.classid} />
        <FormTable classId={params.classid} />
      </Container>
    </ChakraProvider>
  );

  if (hasEditPermission) {
    return <ClassEditView params={params} />;
  } else if (hasViewPermission) {
    return <ClassView params={params} />;
  } else {
    return <NotFound />;
  }
}
