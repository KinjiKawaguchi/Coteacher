export default function ResponseList({
  params,
}: {
  params: { formid: string };
}) {
  return (
    <div>
      <h1>ResponseList{params.formid}</h1>
    </div>
  );
}
