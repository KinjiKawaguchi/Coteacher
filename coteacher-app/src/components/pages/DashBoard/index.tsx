'use client';

import React, { useState, useEffect } from 'react';
import CommonDashBoard from './CommonDashBoard';
import ParticipateClass from '@/components/layout/modal/ParticipateClass';
import CreateClass from '@/components/layout/modal/CreateClass';
import { studentClassRepo } from '@/repository/studentClass';
import { classRepo } from '@/repository/class';
import { Class } from '@/interfaces';

export default function DashBoard() {
  const [userType, setUserType] = useState<string>('');

  useEffect(() => {
    if (typeof window !== 'undefined') {
      setUserType(localStorage.getItem('UserType') || '');
    }
  }, []);

  const fetchStudentClasses = async (): Promise<Class[]> => {
    return await studentClassRepo.getClassListByStudentId();
  };

  const fetchTeacherClasses = async (): Promise<Class[]> => {
    return await classRepo.getClassListByTeacherId();
  };

  if (userType === '1') {
    return (
      <CommonDashBoard
        fetchClasses={fetchStudentClasses}
        ModalContent={ParticipateClass}
      />
    );
  } else if (userType === '2') {
    return (
      <CommonDashBoard
        fetchClasses={fetchTeacherClasses}
        ModalContent={CreateClass}
      />
    );
  }

  return null;
}
