import React from 'react';
import { usePathname } from 'next/navigation';
import { checkAccesable, checkEditable } from '@/libs/services/api';

export default async function ClassPageView() {
  const pathname = usePathname();

  const classID = pathname.split('/')[2];

  const accessable = await checkAccesable(classID);
  const editable = await checkEditable(classID);

  if (!accessable) {
    return <h1>Access denied</h1>;
  }

  if (!editable) {
    return <h1>Read only</h1>;
  }
  return (
    <div>
      <h1>Class ID: {classID}</h1>
    </div>
  );
}
