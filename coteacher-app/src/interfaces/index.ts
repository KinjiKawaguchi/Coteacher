export type Class = {
  id: string;
  name: string;
  teacherId: string;
};

export type User = {
  id: string;
  name: string;
  email: string;
  user_type: number;
};

export type ClassInvitationCode = {
  id: string;
  classId: string;
  invitationCode: string;
  expirationDate: Date;
  isActive: boolean;
}
