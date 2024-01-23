// import ListQuestionComponent from '@/components/layout/QuestionItem/ListQuestion';
// import MultipleChoiceQuestionComponent from '@/components/layout/QuestionItem/MultipleChoiceQuestion';
import ParagraphTextQuestionComponent from '@/components/layout/QuestionItem/ParagraphTextQuestion';
// import RadioQuestionComponent from '@/components/layout/QuestionItem/RadioQuestion';
// import TextQuestionComponent from '@/components/layout/QuestionItem/TextQuestion';
import React from 'react';
import { Form, Question } from '@/interfaces'; // このパスは適宜変更してください。
import { Box, HStack, Skeleton, Spacer, Stack } from '@chakra-ui/react';
import { Question_QuestionType } from '@/gen/proto/coteacher/v1/resources_pb';
import TextQuestionComponent from '@/components/layout/QuestionItem/TextQuestion';
import { Separator } from '@/components/ui/separator';
import { Button } from '@/components/ui/button';
import { callOpenAI } from '@/libs/services/openai';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTrigger,
} from '@/components/ui/dialog';
import { DialogTitle } from '@radix-ui/react-dialog';
interface QuestionListProps {
  questionList: Question[];
  setQuestionList: (questionList: Question[]) => void;
  hasViewPermission: boolean;
  form: Form | null;
}

export default function QuestionList({
  questionList,
  setQuestionList,
  hasViewPermission,
  form,
}: QuestionListProps) {
  const [isLoading, setIsLoading] = React.useState(false);
  const [responseContent, setResponseContent] = React.useState('');
  const [answerList, setAnswerList] = React.useState<string[]>([]);

  const handleAnswerSubmit = async () => {
    try {
      setResponseContent(''); // レスポンス内容をリセット
      setIsLoading(true); // ローディング開始
      console.log(answerList);
      if (form) {
        const response = await callOpenAI(form, questionList, answerList);
        setResponseContent(response.choices[0].message.content || ''); // レスポンス内容を設定
        console.log(response);
      }
    } catch (error) {
      console.error('OpenAI API call failed:', error);
    } finally {
      setIsLoading(false); // ローディング終了
    }
  };
  return (
    <div>
      {questionList.map((question, index) => {
        if (question.order === -1) return null;
        return (
          <div key={question.id || index}>
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
                question.textQuestion && (
                  <div>unimplemented</div>
                  // <ListQuestionComponent question={question} editable={false} />
                )}
              {question.questionType === Question_QuestionType.RADIO &&
                question.textQuestion && (
                  <div>unimplemented</div>
                  // <RadioQuestionComponent question={question} editable={false} />
                )}
              {question.questionType ===
                Question_QuestionType.MULTIPLE_CHOICE &&
                question.textQuestion && (
                  <div>unimplemented</div>
                  // <MultipleChoiceQuestionComponent
                  //   question={question}
                  //   editable={false}
                  // />
                )}
              {question.questionType === Question_QuestionType.PARAGRAPH_TEXT &&
                question.textQuestion && (
                  <ParagraphTextQuestionComponent
                    questionList={questionList}
                    index={index}
                    editable={false}
                    setQuestionList={setQuestionList}
                    answerList={answerList}
                    setAnswerList={setAnswerList}
                  />
                )}
              {question.questionType === Question_QuestionType.TEXT &&
                question.textQuestion && (
                  <TextQuestionComponent
                    questionList={questionList}
                    index={index}
                    editable={false}
                    setQuestionList={setQuestionList}
                    answerList={answerList}
                    setAnswerList={setAnswerList}
                  />
                )}
            </Box>
          </div>
        );
      })}
      {hasViewPermission && questionList.length > 0 && (
        <>
          <Separator className="my-4" />
          <HStack>
            <Spacer />
            <Dialog>
              <DialogTrigger asChild>
                <Button onClick={handleAnswerSubmit} disabled={isLoading}>
                  送信
                </Button>
              </DialogTrigger>
              <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                  <DialogTitle>回答</DialogTitle>
                </DialogHeader>
                {isLoading && (
                  <Stack>
                    <Skeleton height="20px" />
                    <Skeleton height="20px" />
                    <Skeleton height="20px" />
                  </Stack>
                )}
                {responseContent && <Box p={4}>{responseContent}</Box>}
              </DialogContent>
            </Dialog>
          </HStack>
        </>
      )}
    </div>
  );
}
