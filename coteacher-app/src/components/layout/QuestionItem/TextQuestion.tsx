import React from 'react';
import { FormControl, FormLabel, HStack, Text } from '@chakra-ui/react';
import { Input } from '@/components/ui/input';
import { QuestionComponentProps } from '@/interfaces'; // 必要に応じてパスを変更してください。
import { Switch } from '@/components/ui/switch';
import { handleQuestionTextChange, handleMaxLengthChange } from './util';

// TextQuestionComponent: テキストベースの質問フォームをレンダリングするコンポーネント
const TextQuestionComponent: React.FC<QuestionComponentProps> = ({
  questionList,
  index,
  editable,
  setQuestionList,
}) => {
  const question = questionList[index];
  const { questionText, textQuestion, isRequired } = question;

  return (
    <>
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
            placeholder={question.questionText}
          />
        ) : (
          <FormLabel>{questionText}</FormLabel>
        )}
        <Input
          disabled={editable}
          type="text"
          maxLength={textQuestion?.maxLength}
          placeholder="Your answer"
        />
        {editable && (
          <HStack>
            <Switch defaultChecked={textQuestion !== undefined} />
            <Text fontSize="sm" color="gray.500">
              Maximum length:
            </Text>
            <Input
              type="number"
              value={textQuestion?.maxLength || ''}
              disabled={!textQuestion}
              width="50px"
              onChange={e =>
                handleMaxLengthChange(
                  questionList,
                  index,
                  parseInt(e.target.value),
                  setQuestionList
                )
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
