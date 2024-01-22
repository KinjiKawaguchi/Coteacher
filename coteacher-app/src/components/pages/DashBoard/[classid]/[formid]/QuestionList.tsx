// import ListQuestionComponent from '@/components/layout/QuestionItem/ListQuestion';
// import MultipleChoiceQuestionComponent from '@/components/layout/QuestionItem/MultipleChoiceQuestion';
// import ParagraphTextQuestionComponent from '@/components/layout/QuestionItem/ParagraphTextQuestion';
// import RadioQuestionComponent from '@/components/layout/QuestionItem/RadioQuestion';
// import TextQuestionComponent from '@/components/layout/QuestionItem/TextQuestion';
import React from 'react';
import { Question } from '@/interfaces'; // このパスは適宜変更してください。
// import { Box } from '@chakra-ui/react';
// import { Question_QuestionType } from '@/gen/proto/coteacher/v1/resources_pb';
interface QuestionListProps {
  questionList: Question[];
}

export default function QuestionList({ questionList }: QuestionListProps) {
  console.log(questionList);
  return (
    <div>
      {/* {questionList.map((question, index) => (
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
                <ListQuestionComponent question={question} editable={false} />
              )}
            {question.questionType === Question_QuestionType.RADIO &&
              question.textQuestion && (
                <RadioQuestionComponent question={question} editable={false} />
              )}
            {question.questionType === Question_QuestionType.MULTIPLE_CHOICE &&
              question.textQuestion && (
                <MultipleChoiceQuestionComponent
                  question={question}
                  editable={false}
                />
              )}
            {question.questionType === Question_QuestionType.PARAGRAPH_TEXT &&
              question.textQuestion && (
                <ParagraphTextQuestionComponent
                  question={question}
                  editable={false}
                />
              )}
            {question.questionType === Question_QuestionType.TEXT &&
              question.textQuestion && (
                <TextQuestionComponent question={question} editable={false} />
              )}
          </Box>
        </div>
      ))}{' '} */}
    </div>
  );
}
