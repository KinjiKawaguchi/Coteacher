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
