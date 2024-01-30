import ParagraphTextQuestionComponent from '@/components/layout/QuestionItem/ParagraphTextQuestion';
import TextQuestionComponent from '@/components/layout/QuestionItem/TextQuestion';
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion';
import { Question_QuestionType } from '@/gen/proto/coteacher/v1/resources_pb';
import { Response } from '@/interfaces';
import { Box } from '@chakra-ui/react';

export default function ResponseList() {
  const responseList: Response[] = [];
  return (
    <Accordion type="single" collapsible className="w-full">
      {responseList.map((response, index) => (
        <AccordionItem key={response.id} value={response.id}>
          <AccordionTrigger>#{index}</AccordionTrigger>
          {'// TODO: Usernameにする'}
          <AccordionContent>
            {response.answerList.map(answer => (
              <div key={answer.id}>
                <Box
                  border="1px"
                  borderColor="gray.200"
                  borderRadius="md"
                  p={4}
                  mb={4}
                >
                  {answer.question.questionType ===
                    Question_QuestionType.PARAGRAPH_TEXT && (
                    <ParagraphTextQuestionComponent
                      questionList={[answer.question]}
                      index={0}
                      editable={false}
                      setQuestionList={() => {}}
                      answerText={answer.answerText}
                    />
                  )}
                  {answer.question.questionType ===
                    Question_QuestionType.TEXT && (
                    <TextQuestionComponent
                      questionList={[answer.question]}
                      index={0}
                      editable={false}
                      setQuestionList={() => {}}
                      answerText={answer.answerText}
                    />
                  )}
                  {answer.question.questionType ===
                    Question_QuestionType.CHECKBOX && <div>unimplemented</div>}
                  {answer.question.questionType ===
                    Question_QuestionType.RADIO && <div>unimplemented</div>}
                  {answer.question.questionType ===
                    Question_QuestionType.LIST && <div>unimplemented</div>}
                  {answer.question.questionType ===
                    Question_QuestionType.MULTIPLE_CHOICE && (
                    <div>unimplemented</div>
                  )}
                </Box>
              </div>
            ))}
            <Box
              border="1px"
              borderColor="gray.200"
              borderRadius="md"
              p={4}
              mb={4}
            >
              <div>AI回答</div>
              <br />
              <div>{response.aiResponse}</div>
            </Box>
          </AccordionContent>
        </AccordionItem>
      ))}
    </Accordion>
  );
}
