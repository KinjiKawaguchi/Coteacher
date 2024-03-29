import { Question_QuestionType } from '@/gen/proto/coteacher/v1/resources_pb';
import { UserType } from '@/gen/proto/coteacher/v1/user_pb';

export type Class = {
  id: string;
  name: string;
  teacherId: string;
};

export type User = {
  id: string;
  name: string;
  email: string;
  user_type: UserType;
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
  systemPrompt: string;
  usage?: number; // for frontend
};

export type Question = {
  id?: string;
  formId: string;
  questionType: Question_QuestionType;
  questionText: string;
  textQuestion?: TextQuestion;
  options?: Option[];
  isRequired: boolean;
  forAiProcessing: boolean;
  order: number; // -1は削除済み
};

export type TextQuestion = {
  id?: string;
  questionId?: string;
  maxLength: number;
};

export type Option = {
  id?: string;
  questionId?: string;
  optionText: string;
  order: number;
};

export interface QuestionComponentProps {
  questionList: Question[];
  index: number;
  editable: boolean;
  setQuestionList: (questionList: Question[]) => void;
  answerList?: Answer[];
  setAnswerList?: (answerList: Answer[]) => void;
  answerText?: string;
}

export type Answer = {
  id?: string;
  responseId?: string;
  question: Question;
  order: number;
  answerText?: string;
  selectedOptionList?: Option[];
};

export type Response = {
  id?: string;
  formId: string;
  student: User;
  answerList: Answer[];
  aiResponse: string;
};
