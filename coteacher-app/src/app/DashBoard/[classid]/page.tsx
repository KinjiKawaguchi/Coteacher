'use client';
import Class from '@/components/pages/DashBoard/[classid]';

export default function ClassPage({ params }: { params: { classid: string } }) {
  return <Class params={params} />;
}
