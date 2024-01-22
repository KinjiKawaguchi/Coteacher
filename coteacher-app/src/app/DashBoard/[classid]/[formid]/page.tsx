import FormView from '@/components/pages/DashBoard/[classid]/[formid]';

export default function FormPage({ params }: { params: { formid: string } }) {
  return <FormView params={params} />;
}
