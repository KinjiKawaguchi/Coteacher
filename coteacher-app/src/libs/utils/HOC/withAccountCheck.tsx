import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router'; // Corrected import
import { auth } from '@/libs/utils/auth/FirebaseConfig';
import { checkStudentExist } from '@/libs/services/api';
import { Spinner } from '@chakra-ui/react'; // Chakra UI Spinner import

const withAuthAndAccountCheck = <P extends object>(
  Component: React.ComponentType<P>
): React.FC<P> => {
  const WithAuthAndAccountCheck: React.FC<P> = props => {
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState<string | null>(null); // Error state
    const router = useRouter();

    useEffect(() => {
      const unsubscribe = auth.onAuthStateChanged(async user => {
        if (user) {
          if (user.email) {
            try {
              const isStudentExists = await checkStudentExist(user.email);
              if (!isStudentExists) {
                router.push('/RegisterStudent');
              } else {
                setIsLoading(false);
              }
            } catch (err) {
              console.error(err);
              setError(
                err instanceof Error ? err.message : 'An error occurred'
              );
              setIsLoading(false);
            }
          }
        } else {
          router.push('/');
          setIsLoading(false);
        }
      });

      return () => unsubscribe();
    }, [router]);

    if (isLoading) {
      return (
        <div>
          <Spinner /> {/* Chakra UI Spinner */}
          <p>Loading...</p>
        </div>
      );
    }

    if (error) {
      return <div>Error: {error}</div>; // Error display
    }

    return <Component {...(props as P)} />;
  };

  WithAuthAndAccountCheck.displayName = `withAuthAndAccountCheck(${
    Component.displayName || Component.name || 'Component'
  })`;

  return WithAuthAndAccountCheck;
};

export default withAuthAndAccountCheck;
