'use client';

import React, { useState, useEffect } from 'react';
import { Text, HStack } from '@chakra-ui/react';
import { Avatar, AvatarImage } from '@/components/ui/avatar';

import { Github, LifeBuoy, LogOut, Settings, User } from 'lucide-react';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { auth } from '@/libs/utils/auth/FirebaseConfig'; // Firebase認証のインポート
import { useRouter } from 'next/navigation'; // ルーターのインポート

const UserCard = () => {
  const [userType, setUserType] = useState('');
  const [userName, setUserName] = useState('');
  const [isPopoverOpen, setIsPopoverOpen] = useState(false);
  const [isUserPresent, setIsUserPresent] = useState(false); // New state variable

  const router = useRouter(); // ルーターのインスタンスを生成

  const handleLogout = () => {
    auth.signOut(); // Firebaseからログアウト
    setIsUserPresent(false); // ユーザー状態を更新
    localStorage.clear(); // localStorageをクリア
    router.push('/'); // ホームページにリダイレクト
  };

  useEffect(() => {
    if (typeof window !== 'undefined') {
      const userID = localStorage.getItem('UserID');
      if (userID === null) {
        setIsUserPresent(false);
        return;
      }

      setIsUserPresent(true);
      setUserType(localStorage.getItem('UserType') || 'Guest');
      setUserName(localStorage.getItem('UserName') || 'User');
    }
  }, []);

  const displayName = `[${userType}] ${userName}`;

  const handleTogglePopover = () => setIsPopoverOpen(!isPopoverOpen);

  if (!isUserPresent) {
    return null; // Or return some fallback UI if needed
  }
  return (
    //TODO: すべてのボタンを実装する
    <DropdownMenu>
      <DropdownMenuTrigger>
        <HStack spacing={4} align="center" onClick={handleTogglePopover}>
          <Avatar>
            <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
          </Avatar>
          <Text fontWeight="bold">{displayName}</Text>
        </HStack>{' '}
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56">
        <DropdownMenuLabel>My Account</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem>
            <User className="mr-2 h-4 w-4" />
            <span>Profile</span>
            <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
          </DropdownMenuItem>
          <DropdownMenuItem>
            <Settings className="mr-2 h-4 w-4" />
            <span>Settings</span>
            <DropdownMenuShortcut>⌘S</DropdownMenuShortcut>
          </DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem>
          <Github className="mr-2 h-4 w-4" />
          <span>GitHub</span>
        </DropdownMenuItem>
        <DropdownMenuItem>
          <LifeBuoy className="mr-2 h-4 w-4" />
          <span>Support</span>
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={handleLogout}>
          <LogOut className="mr-2 h-4 w-4" />
          <span>Log out</span>
          <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default UserCard;
