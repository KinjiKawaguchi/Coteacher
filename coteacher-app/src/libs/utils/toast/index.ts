import { createStandaloneToast } from '@chakra-ui/react';

const { toast: toastChakra } = createStandaloneToast();

const toast = ({
  status,
  title = '',
  description = '',
}: {
  status: 'success' | 'error' | 'warning' | 'info';
  title?: string;
  description?: string;
}) => {
  toastChakra({
    title,
    description,
    status,
    duration: 3000,
    isClosable: true,
    position: 'top',
  });
};

export default toast;
