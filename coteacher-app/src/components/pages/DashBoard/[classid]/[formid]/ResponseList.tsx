import React, { useState, useEffect } from 'react';
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
import { responseRepo } from '@/repository/response';

export default function ResponseList({ formId }: { formId: string }) {
  const [responseList, setResponseList] = useState<Response[]>([]);

  console.log('formId', formId);
  useEffect(() => {
    async function fetchResponses() {
      const responses = await responseRepo.getResponseListByFormId(formId);
      setResponseList(responses);
    }

    fetchResponses();
  }, [formId]); // Include formId in the dependency array to refetch responses if formId changes

  return (
    <Accordion type="single" collapsible className="w-full">
      {responseList.map((response, index) => (
        <AccordionItem
          key={response.id ?? index.toString()}
          value={response.id ?? index.toString()}
        >
          <AccordionTrigger>
            Response #{index + 1} By {response.student.name}
          </AccordionTrigger>
          <AccordionContent>
            {response.answerList
              .sort((a, b) => a.order - b.order) // Sort answers by their 'order' property
              .map(answer => (
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
                    {/* Implement other question types as needed */}
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
