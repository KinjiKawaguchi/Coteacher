import { FormTable } from './FormTable';

export default function ClassView({ params }: { params: { classid: string } }) {
  return (
    <>
      <div>ClassView</div>
      <FormTable classId={params.classid} />
    </>
  );
}
