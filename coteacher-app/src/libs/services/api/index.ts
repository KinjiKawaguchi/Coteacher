async function checkStudentExist(email: string) {
  // const API_URL = process.env.API_URL!;
  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/Student/CheckAcountExist/${email}`
  );
  const data = await response.json();
  console.log(data.exist);
  return data.exist;
}

async function createStudent(name: string) {
  const email = window.localStorage.getItem('email');
  // POSTリクエストに変更
  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/Student/Create?Email=${email}&Name=${name}`, {
      method: 'POST', // POSTメソッドを指定
    }
  );

  // レスポンスのステータスをチェック
  if (!response.ok) {
    console.error('Response error:', response.status);
    return;
  }

  try {
    const data = await response.json();
    console.log(data);
    return data;
  } catch (error) {
    console.error('Error parsing JSON:', error);
  }
}


export { checkStudentExist, createStudent };
