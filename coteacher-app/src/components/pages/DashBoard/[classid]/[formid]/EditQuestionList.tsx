// import MultipleChoiceQuestionComponent from '@/components/layout/QuestionItem/MultipleChoiceQuestion';
// import ParagraphTextQuestionComponent from '@/components/layout/QuestionItem/ParagraphTextQuestion';
// import RadioQuestionComponent from '@/components/layout/QuestionItem/RadioQuestion';
import TextQuestionComponent from '@/components/layout/QuestionItem/TextQuestion';
import {
  FaRegTrashAlt,
  FaLongArrowAltUp,
  FaLongArrowAltDown,
} from 'react-icons/fa';
import { IoAddCircleOutline } from 'react-icons/io5';
import { Label } from '@/components/ui/label';
import { Switch } from '@/components/ui/switch';
import { Question_QuestionType } from '@/gen/proto/coteacher/v1/resources_pb';
import { Question } from '@/interfaces';
import { Box, HStack, Spacer } from '@chakra-ui/react';
import { Button } from '@/components/ui/button';
import { Separator } from '@/components/ui/separator';

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
    updatedList[index] = {
      ...updatedList[index],
      order: -1,
    };
    setQuestionList(updatedList);
  };

  const handleAddQuestionButtonClick = (index: number) => {
    const updatedList = [...questionList];
    const newQuestion = {
      order: index + 1,
      questionType: Question_QuestionType.TEXT, //TODO: 質問を選べるように
      questionText: '',
      isRequired: false,
      forAiProcessing: false,
      textQuestion: {
        maxLength: 100,
      },
      formId: formId,
    };

    // 新しい質問をリストに追加
    updatedList.splice(index + 1, 0, newQuestion);

    // 追加した質問の後ろの質問の順序を更新
    for (let i = index + 2; i < updatedList.length; i++) {
      updatedList[i].order += 1;
    }

    setQuestionList(updatedList);
  };

  return (
    <div>
      {questionList.map((question, index) => {
        if (question.order === -1) return null;
        // const questionComponentInput: QuestionComponentProps = {
        //   questionList,
        //   index,
        //   editable: true,
        //   setQuestionList,
        // };
        return (
          <div key={question.id || index}>
            <HStack>
              <Separator className="my-4" />
              <Button
                id={`addButton-${index}`}
                variant="ghost"
                onClick={() => {
                  handleAddQuestionButtonClick(index - 1);
                }}
              >
                <IoAddCircleOutline size={24} />
              </Button>
            </HStack>
            <Box
              border="1px"
              borderColor="gray.200"
              borderRadius="md"
              p={4}
              mb={4}
            >
              {question.questionType === Question_QuestionType.CHECKBOX &&
                question.textQuestion && <div>unimplemented</div>}
              {question.questionType === Question_QuestionType.LIST &&
                question.textQuestion && <div>unimplemented</div>}
              {question.questionType === Question_QuestionType.RADIO &&
                question.textQuestion && (
                  <div>unimplemented</div>
                  // <RadioQuestionComponent question={question} editable={true} />
                )}
              {question.questionType ===
                Question_QuestionType.MULTIPLE_CHOICE &&
                question.textQuestion && (
                  <div>unimplemented</div>
                  // <MultipleChoiceQuestionComponent
                  //   question={question}
                  //   editable={true}
                  // />
                )}
              {question.questionType === Question_QuestionType.PARAGRAPH_TEXT &&
                question.textQuestion && (
                  <div>unimplemented</div>
                  // <ParagraphTextQuestionComponent
                  //   question={question}
                  //   editable={true}
                  // />
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
        <Button
          id={`addButton-${questionList.length}`}
          variant="ghost"
          onClick={() => {
            handleAddQuestionButtonClick(questionList.length - 1);
          }}
        >
          <IoAddCircleOutline size={24} />
        </Button>
      </HStack>
    </div>
  );
}
