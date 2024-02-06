// import MultipleChoiceQuestionComponent from '@/components/layout/QuestionItem/MultipleChoiceQuestion';
import ParagraphTextQuestionComponent from '@/components/layout/QuestionItem/ParagraphTextQuestion';
import RadioQuestionComponent from '@/components/layout/QuestionItem/RadioQuestion';
import TextQuestionComponent from '@/components/layout/QuestionItem/TextQuestion';
import {
  FaRegTrashAlt,
  FaLongArrowAltUp,
  FaLongArrowAltDown,
} from 'react-icons/fa';
import { Label } from '@/components/ui/label';
import { Switch } from '@/components/ui/switch';
import { Question_QuestionType } from '@/gen/proto/coteacher/v1/resources_pb';
import { Question } from '@/interfaces';
import { Box, HStack, Spacer } from '@chakra-ui/react';
import { Button } from '@/components/ui/button';
import { Separator } from '@/components/ui/separator';
import CreateQuestionDropdown from './CreateQuestionDropdown';

interface EditQuestionListProps {
  formId: string;
  questionList: Question[];
  setQuestionList: (questionList: Question[]) => void;
}

export default function EditQuestionList({
  formId,
  questionList,
  setQuestionList,
}: EditQuestionListProps) {
  const handleAddQuestionOptionButton = (index: number) => {
    const updatedList = [...questionList];
    updatedList[index] = {
      ...updatedList[index],
      options: [
        ...(updatedList[index].options || []),
        {
          optionText: '',
          order: updatedList[index].options?.length || 0,
        },
      ],
    };
    setQuestionList(updatedList);
  };
  const handleIsRequiredChange = (index: number) => {
    const updatedList = [...questionList];
    updatedList[index] = {
      ...updatedList[index],
      isRequired: !updatedList[index].isRequired,
    };
    setQuestionList(updatedList);
  };

  const handleForAIProcessingChange = (index: number) => {
    const updatedList = [...questionList];
    updatedList[index] = {
      ...updatedList[index],
      forAiProcessing: !updatedList[index].forAiProcessing,
    };
    setQuestionList(updatedList);
  };

  const handleDeleteButtonClick = (index: number) => {
    const updatedList = [...questionList];
    if (updatedList[index].id) {
      // 既存の質問の場合は order を -1 にする
      updatedList[index] = {
        ...updatedList[index],
        order: -1,
      };
      for (let i = index + 1; i < updatedList.length; i++) {
        updatedList[i] = {
          ...updatedList[i],
          order: updatedList[i].order - 1,
        };
      }
    } else {
      // 新規追加した質問の場合は削除
      updatedList.splice(index, 1);
      for (let i = index; i < updatedList.length; i++) {
        updatedList[i] = {
          ...updatedList[i],
          order: i,
        };
      }
    }
    // 後ろの質問の順序を更新

    setQuestionList(updatedList);
  };

  let visibleQuestionIndex = -1; //!可視の質問のindexを初期化

  return (
    <div>
      {questionList.map((question, index) => {
        if (question.order === -1) return null; //!orderが-1のものは非表示
        visibleQuestionIndex++; //!可視の質問のindexを更新
        return (
          <div key={question.id || `question-${index}`}>
            <HStack>
              <Separator className="my-4" />
              <CreateQuestionDropdown
                formId={formId}
                questionList={questionList}
                setQuestionList={setQuestionList}
                index={index}
                order={visibleQuestionIndex}
              />
            </HStack>
            <Box
              border="1px"
              borderColor="gray.200"
              borderRadius="md"
              p={4}
              mb={4}
            >
              {question.questionType === Question_QuestionType.CHECKBOX && (
                <div>unimplemented</div>
              )}
              {question.questionType === Question_QuestionType.LIST && (
                <div>unimplemented</div>
              )}
              {question.questionType === Question_QuestionType.RADIO && (
                <RadioQuestionComponent
                  questionList={questionList}
                  index={index}
                  editable={true}
                  setQuestionList={setQuestionList}
                />
              )}
              {question.questionType ===
                Question_QuestionType.MULTIPLE_CHOICE && (
                <div>unimplemented</div>
                // <MultipleChoiceQuestionComponent
                //   question={question}
                //   editable={true}
                // />
              )}
              {question.questionType === Question_QuestionType.PARAGRAPH_TEXT &&
                question.textQuestion && (
                  <ParagraphTextQuestionComponent
                    questionList={questionList}
                    index={index}
                    editable={true}
                    setQuestionList={setQuestionList}
                  />
                )}
              {question.questionType === Question_QuestionType.TEXT &&
                question.textQuestion && (
                  <TextQuestionComponent
                    questionList={questionList}
                    editable={true}
                    setQuestionList={setQuestionList}
                    index={index}
                  />
                )}
              {(question.questionType === Question_QuestionType.CHECKBOX ||
                question.questionType === Question_QuestionType.LIST ||
                question.questionType === Question_QuestionType.RADIO ||
                question.questionType ===
                  Question_QuestionType.MULTIPLE_CHOICE) && (
                <Button
                  variant="ghost"
                  onClick={() => handleAddQuestionOptionButton(index)}
                >
                  ＋オプションを追加
                </Button>
              )}
              <HStack>
                <Switch
                  id={`isRequired-${index}`}
                  defaultChecked={question.isRequired}
                  onClick={() => handleIsRequiredChange(index)}
                />
                <Label htmlFor={`isRequired-${index}`}>必須</Label>
                <Switch
                  id={`forAIProcessing-${index}`}
                  defaultChecked={question.forAiProcessing}
                  onClick={() => handleForAIProcessingChange(index)}
                />
                <Label htmlFor={`forAIProcessing-${index}`}>
                  AIへのプロンプトに含める
                </Label>
                <Spacer />
                <Button id={`upButton-${index}`} variant="ghost">
                  <FaLongArrowAltUp />
                </Button>
                <Button id={`downButton-${index}`} variant="ghost">
                  <FaLongArrowAltDown />
                </Button>
                <Button
                  id={`deleteButton-${index}`}
                  onClick={() => handleDeleteButtonClick(index)}
                  variant="ghost"
                >
                  <FaRegTrashAlt />
                </Button>
              </HStack>
            </Box>
          </div>
        );
      })}
      <HStack>
        <Separator className="my-4" />
        <CreateQuestionDropdown
          formId={formId}
          questionList={questionList}
          setQuestionList={setQuestionList}
          index={questionList.length}
          order={visibleQuestionIndex + 1}
        />
      </HStack>
    </div>
  );
}
