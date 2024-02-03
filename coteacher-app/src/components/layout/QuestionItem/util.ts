import { Answer, Question } from '@/interfaces';

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
  answerList: Answer[],
  index: number,
  text: string,
  setAnswerList?: (answerList: Answer[]) => void
) {
  const newAnswerList = [...answerList];
  newAnswerList[index].answerText = text;
  if (setAnswerList) {
    setAnswerList(newAnswerList);
  }
}
export function handleOptionTextChange(
  questionList: Question[],
  index: number,
  optionIndex: number,
  text: string,
  setQuestionList: (questionList: Question[]) => void
) {
  const newQuestionList = [...questionList];
  const currentQuestion = newQuestionList[index];
  if (currentQuestion.options) {
    if (currentQuestion.options.length > optionIndex) {
      currentQuestion.options[optionIndex].optionText = text;
    }
  }
  setQuestionList(newQuestionList);
}
export function handleDeleteQuestionOption(
  questionList: Question[],
  index: number,
  optionIndex: number,
  setQuestionList: (questionList: Question[]) => void
) {
  //orderを-1にすることで、削除されたことを示す
  const newQuestionList = [...questionList];
  const currentQuestion = newQuestionList[index];
  if (currentQuestion.options) {
    if (currentQuestion.options.length > optionIndex) {
      currentQuestion.options[optionIndex].order = -1;
    }
    // orderを振り直す
    let order = 0;
    currentQuestion.options.forEach(option => {
      if (option.order !== -1) {
        option.order = order;
        order++;
      }
    });
  }
  setQuestionList(newQuestionList);
}
