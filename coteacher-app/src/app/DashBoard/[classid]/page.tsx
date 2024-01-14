export default function ClassPage({ params }: { params: { classid: string } }) {
  return <div>Class {params.classid}</div>;
}