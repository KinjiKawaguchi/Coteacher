async function checkStudentExist(email: string) {
  // const API_URL = process.env.API_URL!;
  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/Student/CheckAccountExist?Email=${email}`,
    {
      method: 'GET',
    }
  );
  const data = await response.json();
  console.log(data.exist);
  return data.exist;
}

async function createStudent(name: string) {
  const email = window.localStorage.getItem('email');
  // POSTリクエストに変更
  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/Student/Create?Email=${email}&Name=${name}`,
    {
      method: 'POST',
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
    return response;
  } catch (error) {
    console.error('Error parsing JSON:', error);
  }
}

async function getParticipatingClass() {
  // const StudentID = window.localStorage.getItem('StudentID');
  const StudentID = 'e08397af-2d13-4814-b671-9831bb2e395b';

  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/StudentClass/GetParticipatingClass?StudentID=${StudentID}`,
    {
      method: 'GET',
    }
  );

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

async function participateClass(invitationCode: string) {
  const StudentID = 'e08397af-2d13-4814-b671-9831bb2e395b';

  try {
    const response = await fetch(
      `https://api-image-pgfe7sqiia-an.a.run.app/StudentClass/ParticipateClass?StudentID=${StudentID}&InvitationCode=${invitationCode}`,
      {
        method: 'GET',
      }
    );

    const data = await response.json(); // Always parse the JSON response

    if (response.ok) {
      return { status: response.status, message: data.message };
    } else {
      // For error cases, include the status code and error message
      return {
        status: response.status,
        error: data.message || 'An error occurred',
      };
    }
  } catch (error: any) {
    return { status: 'network-error', error: error.message || 'Network error' };
  }
}

export {
  checkStudentExist,
  createStudent,
  getParticipatingClass,
  participateClass,
};
