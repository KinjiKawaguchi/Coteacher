import { Question } from '@/interfaces';

export function handleQuestionTextChange(
  questionList: Question[],
  index: number,
  text: string,
  setQuestionList: (questionList: Question[]) => void
) {
  const newQuestionList = [...questionList];
  newQuestionList[index].questionText = text;
  setQuestionList(newQuestionList);
}

export function handleMaxLengthChange(
  questionList: Question[],
  index: number,
  maxLength: number,
  setQuestionList: (questionList: Question[]) => void
) {
  const newQuestionList = [...questionList];
  const currentQuestion = newQuestionList[index];
  if (currentQuestion.textQuestion) {
    currentQuestion.textQuestion.maxLength = maxLength;
  }
  setQuestionList(newQuestionList);
}

export function handleAnswerChange(
  answerList: string[],
  index: number,
  text: string,
  setAnswerList?: (answerList: string[]) => void
) {
  const newAnswerList = [...answerList];
  newAnswerList[index] = text;
  if (setAnswerList) {
    setAnswerList(newAnswerList);
  }
}
