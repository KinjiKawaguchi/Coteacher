import React from 'react';
import { QuestionComponentProps } from '@/interfaces';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';
import { FormControl, FormLabel, HStack } from '@chakra-ui/react';
import { Button } from '@/components/ui/button';
import { MdCancel } from 'react-icons/md';
import { Input } from '@/components/ui/input';
import {
  handleDeleteQuestionOption,
  handleOptionTextChange,
  handleQuestionTextChange,
} from './util';
import { Label } from '@/components/ui/label';

const RadioQuestionComponent: React.FC<QuestionComponentProps> = ({
  questionList,
  index,
  editable,
  setQuestionList,
}) => {
  const question = questionList[index];
  const { questionText, options } = question;

  // Find the first option with order=0 or use the first option as default
  const defaultOption =
    options?.find(option => option.order === 0) || options?.[0];

  return (
    <FormControl isRequired={question.isRequired}>
      {editable && (
        <>
          <Input
            value={questionText}
            onChange={e =>
              handleQuestionTextChange(
                questionList,
                index,
                e.target.value,
                setQuestionList
              )
            }
          />
          {options?.map(
            (option, optionIndex) =>
              option.order !== -1 && ( // ここで描画するかどうかを判断
                <HStack key={option.id}>
                  <Input
                    value={option.optionText}
                    onChange={e =>
                      handleOptionTextChange(
                        questionList,
                        index,
                        optionIndex,
                        e.target.value,
                        setQuestionList
                      )
                    }
                  />
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() =>
                      handleDeleteQuestionOption(
                        questionList,
                        index,
                        optionIndex,
                        setQuestionList
                      )
                    }
                  >
                    <MdCancel />
                  </Button>
                </HStack>
              )
          )}
        </>
      )}
      {!editable && (
        <>
          <FormLabel>{question.questionText}</FormLabel>
          <RadioGroup defaultValue={String(defaultOption?.order)}>
            {options?.map(
              (option, optionIndex) =>
                option.order !== -1 && (
                  <HStack key={`option-${option.id || optionIndex}`}>
                    <RadioGroupItem
                      value={option.optionText}
                      id={`radio-${option.id || optionIndex}`}
                    />
                    <Label htmlFor={`radio-${option.id || optionIndex}`}>
                      {option.optionText}
                    </Label>
                  </HStack>
                )
            )}
          </RadioGroup>
        </>
      )}
    </FormControl>
  );
};

export default RadioQuestionComponent;
