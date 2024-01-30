import React from 'react';
import { FormControl, FormLabel, HStack, Text } from '@chakra-ui/react';
import { Textarea } from '@/components/ui/textarea';
import { Input } from '@/components/ui/input';
import { QuestionComponentProps } from '@/interfaces'; // 必要に応じてパスを変更してください。
import {
  handleQuestionTextChange,
  handleMaxLengthChange,
  handleAnswerChange,
} from './util';
import { Switch } from '@/components/ui/switch';

// ParagraphTextQuestionComponent: パラグラフ形式の質問フォームをレンダリングするコンポーネント
const ParagraphTextQuestionComponent: React.FC<QuestionComponentProps> = ({
  questionList,
  index,
  editable,
  setQuestionList,
  answerList,
  setAnswerList,
  answerText,
}) => {
  const question = questionList[index];
  const { questionText, isRequired, textQuestion } = question;

  return (
    <FormControl isRequired={isRequired}>
      {editable ? (
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
          placeholder="Edit the question"
        />
      ) : (
        <FormLabel>{questionText}</FormLabel>
      )}
      {answerText && (
        <Textarea
          disabled={editable}
          maxLength={textQuestion?.maxLength}
          placeholder="Your answer"
          value={answerText}
        />
      )}
      {!answerText && (
        <Textarea
          disabled={editable}
          placeholder="Your answer"
          maxLength={textQuestion?.maxLength}
          onChange={e =>
            handleAnswerChange(
              answerList || [''],
              index,
              e.target.value,
              setAnswerList
            )
          }
        />
      )}
      {editable && (
        <HStack>
          <Switch defaultChecked={textQuestion !== undefined} />
          <Text fontSize="sm" color="gray.500">
            Maximum length:
          </Text>
          <Input
            type="number"
            placeholder="Maximum length"
            defaultValue={textQuestion?.maxLength}
            onChange={e =>
              handleMaxLengthChange(
                questionList,
                index,
                parseInt(e.target.value),
                setQuestionList
              )
            }
          />
        </HStack>
      )}
      {!editable && textQuestion?.maxLength && (
        <HStack>
          <Text fontSize="sm" color="gray.500">
            Maximum length: {textQuestion.maxLength} characters
          </Text>
        </HStack>
      )}
    </FormControl>
  );
};

export default ParagraphTextQuestionComponent;
