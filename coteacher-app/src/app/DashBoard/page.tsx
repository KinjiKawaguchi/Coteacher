'use client';
import React, { useState, useEffect } from 'react';
import StudentDashBoard from '@/components/pages/DashBoard/StudentDashBoard'; // Corrected import if needed
import TeacherDashBoard from '@/components/pages/DashBoard/TeacherDashBoard';

export default function DashBoard() {
  const [userType, setUserType] = useState('');

  useEffect(() => {;
    // Check if running in a browser environment
    if (typeof window !== 'undefined') {
      setUserType(localStorage.getItem('UserType') || '');
    }
  }, []);

  if (userType === 'Student') {
    return <StudentDashBoard />;
  } else if (userType === 'Teacher') {
    // Assuming 'Teacher' is the other user type
    return <TeacherDashBoard />;
  }

  return null; // Or some loading indicator or fallback UI
}
