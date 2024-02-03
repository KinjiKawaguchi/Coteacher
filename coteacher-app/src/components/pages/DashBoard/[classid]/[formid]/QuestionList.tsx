// import ListQuestionComponent from '@/components/layout/QuestionItem/ListQuestion';
// import MultipleChoiceQuestionComponent from '@/components/layout/QuestionItem/MultipleChoiceQuestion';
import ParagraphTextQuestionComponent from '@/components/layout/QuestionItem/ParagraphTextQuestion';
import RadioQuestionComponent from '@/components/layout/QuestionItem/RadioQuestion';
// import TextQuestionComponent from '@/components/layout/QuestionItem/TextQuestion';
import React, { useEffect } from 'react';
import { Answer, Form, Question } from '@/interfaces'; // このパスは適宜変更してください。
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
import { responseRepo } from '@/repository/response';
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
  const [answerList, setAnswerList] = React.useState<Answer[]>([]);

  // `questionList` が変更されたときにのみ回答リストを再計算
  useEffect(() => {
    const validAnswers = questionList.reduce((acc, question) => {
      if (question.order === -1) return acc; // '-1' で省かれた質問を無視

      const answer: Answer = {
        question: question,
        order: acc.length, // 累積リストの長さを新しい `order` として使用
        answerText: '',
        selectedOptionList: [],
      };

      acc.push(answer); // 有効な回答を累積リストに追加
      return acc;
    }, [] as Answer[]);

    setAnswerList(validAnswers); // 最終的な回答リストを設定
  }, [questionList]);

  const handleAnswerSubmit = async () => {
    try {
      setResponseContent(''); // レスポンス内容をリセット
      setIsLoading(true); // ローディング開始
      const submitResponseReponse = await responseRepo.submitResponse(
        form?.id || '',
        answerList
      ); // 回答を保存
      if (!submitResponseReponse.success)
        throw new Error('Failed to submit response');
      if (form) {
        const AiResponse = (await callOpenAI(form, questionList, answerList))
          .choices[0].message.content;
        if (!AiResponse) throw new Error('AI response is empty');
        setResponseContent(AiResponse);
        responseRepo.submitAiResponse(
          submitResponseReponse.responseId,
          AiResponse
        );
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
              {question.questionType === Question_QuestionType.CHECKBOX && (
                <div>unimplemented</div>
              )}
              {question.questionType === Question_QuestionType.LIST && (
                <div>unimplemented</div>
                // <ListQuestionComponent question={question} editable={false} />
              )}
              {question.questionType === Question_QuestionType.RADIO && (
                <RadioQuestionComponent
                  questionList={questionList}
                  index={index}
                  editable={false}
                  setQuestionList={setQuestionList}
                />
              )}
              {question.questionType ===
                Question_QuestionType.MULTIPLE_CHOICE && (
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
