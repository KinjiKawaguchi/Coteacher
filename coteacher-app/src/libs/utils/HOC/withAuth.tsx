import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation'; // Corrected import
import { auth } from '@/libs/utils/auth/FirebaseConfig'; // Firebase configuration import
import { userRepo } from '@/repository/user';
import { CheckUserExistsByEmailRequest } from '@/gen/proto/coteacher/v1/user_pb';

// Define a type for the component props if known, or use generics for flexibility
const withAuth = <P extends object>(
  Component: React.ComponentType<P>
): React.FC<P> => {
  const WithAuthComponent: React.FC<P> = props => {
    const [isLoading, setIsLoading] = useState(true);
    const router = useRouter();

    useEffect(() => {
      const unsubscribe = auth.onAuthStateChanged(async user => {
        if (user?.email) {
          const request = new CheckUserExistsByEmailRequest();
          request.email = user.email;
          if (await userRepo.checkUserExistsByEmail(request)) {
            router.push('/DashBoard');
          }
          setIsLoading(false);
        } else {
          router.push('/');
          setIsLoading(false);
        }
      });
      // Unsubscribe the listener when the component unmounts
      return () => unsubscribe();
    }, [router]);

    if (isLoading) {
      return <div>Loading...</div>;
    }

    return <Component {...(props as P)} />;
  };

  WithAuthComponent.displayName = `withAuth(${
    Component.displayName || Component.name || 'Component'
  })`;

  return WithAuthComponent;
};

export default withAuth;
