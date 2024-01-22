import { Question } from '@/interfaces';

export default function handleQuestionTextChange(
  questionList: Question[],
  index: number,
  text: string,
  setQuestionList: (questionList: Question[]) => void
) {
  const newQuestionList = [...questionList];
  newQuestionList[index].questionText = text;
  setQuestionList(newQuestionList);
}
