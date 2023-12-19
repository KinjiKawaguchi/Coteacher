async function getUser(email: string) {
  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/User/Get?Email=${email}`,
    {
      method: 'GET',
    }
  );

  if (!response.ok) {
    console.error('Response error:', response.status);
    return null;
  }

  try {
    const data = await response.json();
    console.log(data);
    return data;
  } catch (error) {
    console.error('Error parsing JSON:', error);
    return null;
  }
}

async function checkUserExist(email: string) {
  const data = await getUser(email!);
  if (data.user.UserType === 'Teacher' || data.user.UserType === 'Student') {
    localStorage.setItem('UserID', data.user.ID);
    localStorage.setItem('UserEmail', data.user.Email);
    localStorage.setItem('UserName', data.user.Name);
    localStorage.setItem('UserType', data.user.UserType);
    console.log(localStorage.getItem('UserID'));
    return true;
  } else {
    return false;
  }
}

async function checkStudentExist(email: string) {
  const data = await getUser(email!);
  if(!data) {
    return false;
  }
  if (data.user.UserType === 'Student') {
    localStorage.setItem('UserID', data.user.ID);
    localStorage.setItem('UserEmail', data.user.Email);
    localStorage.setItem('UserName', data.user.Name);
    localStorage.setItem('UserType', data.user.UserType);
    console.log(localStorage.getItem('UserID'));
    return true;
  } else {
    return false;
  }
}

async function checkTeacherExist(email: string) {
  const data = await getUser(email!);
  if(!data) {
    return false;
  }
  if (data.user.UserType === 'Teacher') {
    localStorage.setItem('UserID', data.user.ID);
    localStorage.setItem('UserEmail', data.user.Email);
    localStorage.setItem('UserName', data.user.Name);
    localStorage.setItem('UserType', data.user.UserType);
    console.log(localStorage.getItem('UserID'));
    return true;
  } else {
    return false;
  }
}

async function createStudent(name: string) {
  const email = window.localStorage.getItem('email');
  // POSTリクエストに変更
  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/User/Create?Email=${email}&Name=${name}&UserType=Student`,
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
    localStorage.setItem('UserID', data.user.ID);
    localStorage.setItem('UserEmail', data.user.Email);
    localStorage.setItem('UserName', data.user.Name);
    localStorage.setItem('UserType', data.user.UserType);
    console.log(localStorage.getItem('UserID'));
    return response;
  } catch (error) {
    console.error('Error parsing JSON:', error);
  }
}

async function getParticipatingClass() {
  const StudentID = localStorage.getItem('UserID');

  const response = await fetch(
    `https://api-image-pgfe7sqiia-an.a.run.app/StudentClass/GetList?StudentID=${StudentID}`,
    {
      method: 'GET',
    }
  );

  if (!response.ok) {
    console.error('Response error:', response.status);
    return null;
  }

  try {
    const data = await response.json();
    console.log(data);
    return data;
  } catch (error) {
    return null;
  }
}

async function participateClass(invitationCode: string) {
  const StudentID = localStorage.getItem('UserID');

  try {
    const response = await fetch(
      `https://api-image-pgfe7sqiia-an.a.run.app/StudentClass/Create?StudentID=${StudentID}&InvitationCode=${invitationCode}`,
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
  } catch (error: unknown) {
    const errorMessage =
      error instanceof Error ? error.message : 'Network error';
    return { status: 'network-error', error: errorMessage };
  }
}

export {
  getUser,
  checkUserExist,
  checkTeacherExist,
  checkStudentExist,
  createStudent,
  getParticipatingClass,
  participateClass,
};
