'use client';
import React, { useState, useEffect } from 'react';
import default_user_icon from '@/images/default-avatar-profile-icon.jpg';
import {
  Image,
  Text,
  HStack,
  Popover,
  PopoverTrigger,
  Button,
  PopoverContent,
  Portal,
} from '@chakra-ui/react';
import { auth } from '@/libs/utils/auth/FirebaseConfig';
//router
import { useRouter } from 'next/navigation';

const UserCard = () => {
  const [userType, setUserType] = useState('');
  const [userName, setUserName] = useState('');
  const [isPopoverOpen, setIsPopoverOpen] = useState(false);
  const router = useRouter();

  useEffect(() => {
    setUserType(localStorage.getItem('UserType') || 'Guest');
    setUserName(localStorage.getItem('UserName') || 'User');
  }, []);

  const displayName = `[${userType}] ${userName}`;

  const handleTogglePopover = () => setIsPopoverOpen(!isPopoverOpen);

  return (
    <Popover isOpen={isPopoverOpen} onClose={() => setIsPopoverOpen(false)}>
      <PopoverTrigger>
        <HStack spacing={4} align="center" onClick={handleTogglePopover}>
          <Image
            borderRadius="full"
            boxSize="50px"
            src={default_user_icon.src}
            alt="User Avatar"
          />
          <Text fontWeight="bold">{displayName}</Text>
        </HStack>
      </PopoverTrigger>
      <Portal>
        <PopoverContent>
          <Button
            onClick={() => {
              localStorage.clear();
              auth.signOut();
              router.push('/');
            }}
          >
            Logout
          </Button>
        </PopoverContent>
      </Portal>
    </Popover>
  );
};

export default UserCard;