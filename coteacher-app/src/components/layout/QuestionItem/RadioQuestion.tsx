import React from 'react';
import { QuestionComponentProps } from '@/interfaces';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';
import { Label } from '@/components/ui/label';
import { HStack } from '@chakra-ui/react';
import { Button } from '@/components/ui/button';
import { MdCancel } from 'react-icons/md';
import { Input } from '@/components/ui/input';
import {
  handleDeleteQuestionOption,
  handleOptionTextChange,
  handleQuestionTextChange,
} from './util';

const RadioQuestionComponent: React.FC<QuestionComponentProps> = ({
  questionList,
  index,
  editable,
  setQuestionList,
}) => {
  const question = questionList[index];
  const { questionText, options, isRequired } = question;

  // Find the first option with order=0 or use the first option as default
  const defaultOption =
    options?.find(option => option.order === 0) || options?.[0];

  // Filter out options with order=-1
  const visibleOptions = options?.filter(option => option.order !== -1) || [];

  return (
    <>
      {editable && (
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
      )}
      {!editable && (
        <div className="mb-2">
          {questionText} {isRequired && <span className="text-red-500">*</span>}
        </div>
      )}
      <RadioGroup defaultValue={defaultOption?.optionText}>
        {visibleOptions.map((option, optionIndex) => (
          <div
            key={option.id || optionIndex}
            className="flex items-center space-x-2"
          >
            <HStack>
              <RadioGroupItem
                value={option.optionText}
                id={`radio-${option.id || optionIndex}`}
              />
              {editable && (
                <Input
                  onChange={e =>
                    handleOptionTextChange(
                      questionList,
                      index,
                      optionIndex,
                      e.target.value,
                      setQuestionList
                    )
                  }
                  value={option.optionText}
                />
              )}
              {!editable && (
                <Label htmlFor={`radio-${option.id || optionIndex}`}>
                  {option.optionText}
                </Label>
              )}
              {editable && (
                <Button
                  onClick={() =>
                    handleDeleteQuestionOption(
                      questionList,
                      index,
                      optionIndex,
                      setQuestionList
                    )
                  }
                  variant="ghost"
                >
                  <MdCancel />
                </Button>
              )}
            </HStack>
          </div>
        ))}
      </RadioGroup>
    </>
  );
};

export default RadioQuestionComponent;
