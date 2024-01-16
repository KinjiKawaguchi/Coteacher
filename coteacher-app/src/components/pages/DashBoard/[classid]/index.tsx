import React, { useEffect, useState } from 'react';
import NotFound from '@/app/not-found';
import { classRepo } from '@/repository/class';
import { Spinner } from '@chakra-ui/react';
import ClassEditView from './ClassEditView';
import ClassView from './ClassView';

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

  if (hasEditPermission) {
    return <ClassEditView params={params} />;
  } else if (hasViewPermission) {
    return <ClassView params={params} />;
  } else {
    return <NotFound />;
  }
}
