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
};

export type Form = {
  id: string;
  classId: string;
  name: string;
  description: string;
  usageLimit: number;
};
