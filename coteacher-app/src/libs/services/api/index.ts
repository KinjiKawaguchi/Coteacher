async function checkStudentExist(email: string) {
  // const API_URL = process.env.API_URL!;
  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/Student/CheckAcountExist/${email}`
  );
  const data = await response.json();
  console.log(data.exist);
  return data.exist;
}

export { checkStudentExist };
