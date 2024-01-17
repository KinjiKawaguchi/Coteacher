import { FormTable } from './FormTable';

export default function ClassEditView({
  params,
}: {
  params: { classid: string };
}) {
  return (
    <>
      <div>ClassEditView</div>
      <FormTable classId={params.classid} />
    </>
  );
}
