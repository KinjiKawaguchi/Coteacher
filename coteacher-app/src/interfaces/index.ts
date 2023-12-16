export type Class = {
  ID: string;
  Name: string;
  TeacherID: string;
};

export type User = {
  Type: 'Student' | 'Teacher';
  ID: string;
  Name: string;
  Email: string;
};
