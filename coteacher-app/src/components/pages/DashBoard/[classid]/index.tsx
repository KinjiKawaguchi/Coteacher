import React, { useEffect, useState } from 'react';
import { classRepo } from '@/repository/class';
import { ChakraProvider, Spinner, Container, Flex } from '@chakra-ui/react';
import theme from '@/theme';
import UserHeader from '@/components/layout/header/UserHeader';
import { ClassHeader } from '@/components/layout/header/ClassHeader';
import { Button } from '@/components/ui/button';
import { useRouter } from 'next/navigation';
import { FormTable } from './FormTable';
import { ArrowLeft } from 'lucide-react';
import CreateFormDialog from './CreateFormDialog';
import { formRepo } from '@/repository/form';
import { Form } from '@/interfaces';
import { responseRepo } from '@/repository/response';

export default function Class({ params }: { params: { classid: string } }) {
  const [hasEditPermission, setHasEditPermission] = useState<boolean | null>(
    null
  );
  const [hasViewPermission, setHasViewPermission] = useState<boolean | null>(
    null
  );
  const [forms, setForms] = useState<Form[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  const updateFormList = async () => {
    const updatedForms = await formRepo.getFormListByClassId(params.classid);
    setForms(updatedForms);
  };

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

      const formList = await formRepo.getFormListByClassId(params.classid);
      const userType = localStorage.getItem('UserType');

      const formsWithUsage = await Promise.all(
        formList.map(async form => {
          const usage =
            userType === '1'
              ? await responseRepo.getNumberOfResponseByStudentId(null, form.id)
              : await responseRepo.getNumberOfResponseByFormId(form.id);
          return { ...form, usage };
        })
      );

      setForms(formsWithUsage);
      setIsLoading(false);
    };

    fetchData();
  }, [params.classid]);

  if (hasEditPermission === null || hasViewPermission === null || isLoading) {
    return <Spinner />;
  }

  if (!hasEditPermission && !hasViewPermission) {
    return <div>権限がありません</div>;
  }

  const router = useRouter();
  const navDashBoard = () => router.back();

  return (
    <ChakraProvider theme={theme}>
      <Button variant="ghost" onClick={navDashBoard}>
        <ArrowLeft />
      </Button>
      <Container maxWidth="container.sm">
        <UserHeader />
        <ClassHeader classId={params.classid} />
        <Flex justifyContent="flex-end">
          <CreateFormDialog params={params} updateFormList={updateFormList} />
        </Flex>
        <FormTable forms={forms} />
      </Container>
    </ChakraProvider>
  );
}
