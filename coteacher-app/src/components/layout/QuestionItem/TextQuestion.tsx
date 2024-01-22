import React from 'react';
import { FormControl, FormLabel, HStack, Text } from '@chakra-ui/react';
import { Input } from '@/components/ui/input';
import { QuestionComponentProps } from '@/interfaces'; // このパスは適宜変更してください。
import { Switch } from '@/components/ui/switch';
import handleQuestionTextChange from './util';
//TODO: MaxLength周りの実装をする

const TextQuestionComponent: React.FC<QuestionComponentProps> = ({
  questionList,
  index,
  editable,
  setQuestionList,
}: QuestionComponentProps) => {
  const question = questionList[index];
  const { questionText, textQuestion, isRequired } = question;

  const handleMaxLengthChange = (index: number, maxLength: number) => {
    const newQuestionList = [...questionList];
    const currentQuestion = newQuestionList[index];
    if (currentQuestion.textQuestion) {
      currentQuestion.textQuestion.maxLength = maxLength;
    }
    setQuestionList(newQuestionList);
  };

  // const handleMaxLengthUsageChange = (index: number) => {
  //   const newQuestionList = [...questionList];
  //   const currentQuestion = newQuestionList[index];

  //   if (currentQuestion.textQuestion) {
  //     // textQuestion を未定義にする
  //     currentQuestion.textQuestion = undefined;
  //   } else {
  //     // textQuestion が未定義の場合、新しいオブジェクトを作成
  //     currentQuestion.textQuestion = {
  //       ...(currentQuestion.textQuestion ?? {}),
  //       questionId: currentQuestion.id,
  //       max_length: 100, // ここに適切な初期値を設定
  //     };
  //   }
  //   setQuestionList(newQuestionList);
  // };

  return (
    <>
      <FormControl isRequired={isRequired}>
        {editable && (
          <Input
            type="text"
            value={question.questionText}
            onChange={e =>
              handleQuestionTextChange(
                questionList,
                index,
                e.target.value,
                setQuestionList
              )
            }
            placeholder={question.questionText}
          />
        )}
        {!editable && <FormLabel>{questionText}</FormLabel>}
        <Input
          disabled={editable}
          type="text"
          maxLength={100}
          placeholder="Your answer"
        />
        {editable && (
          <HStack>
            <Switch
              defaultChecked={question.textQuestion !== undefined}
              //onClick={() => handleMaxLengthUsageChange(index)}
            />
            <Text fontSize="sm" color="gray.500">
              Maximum length:
            </Text>
            <Input
              type="number"
              value={textQuestion?.maxLength}
              disabled={!question.textQuestion}
              width="50px"
              onChange={e =>
                handleMaxLengthChange(index, parseInt(e.target.value))
              }
            />
            <Text fontSize="sm" color="gray.500">
              characters
            </Text>
          </HStack>
        )}
        {!editable && textQuestion?.maxLength && (
          <Text fontSize="sm" color="gray.500">
            Maximum length: {textQuestion.maxLength} characters
          </Text>
        )}
      </FormControl>
    </>
  );
};

export default TextQuestionComponent;
