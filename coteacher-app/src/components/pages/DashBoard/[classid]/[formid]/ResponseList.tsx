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
  const responseList: Response[] = [
    {
      id: '1',
      formId: '1',
      studentId: '1',
      answerList: [
        {
          id: '1',
          responseId: '1',
          answerText: 'Python',
          question: {
            id: '1',
            order: 0,
            questionType: 6,
            questionText: '使用している言語はなんですか?',
            textQuestion: {
              id: '1',
              maxLength: 100,
              questionId: '1',
            },
            formId: '1',
            isRequired: true,
            forAiProcessing: true,
          },
        },
        {
          id: '1',
          responseId: '1',
          answerText: 'プログラムが動作しません。',
          question: {
            id: '1',
            order: 0,
            questionType: 6,
            questionText: 'どんな問題に直面していますか?',
            textQuestion: {
              id: '1',
              maxLength: 100,
              questionId: '1',
            },
            formId: '1',
            isRequired: true,
            forAiProcessing: true,
          },
        },
        {
          id: '1',
          responseId: '1',
          answerText:
            'エラーメッセージを翻訳してみると、型の不一致が原因なようです。\nPythonは型を自動で解決してくれると授業で習ったのですが、なぜエラーにナルのでしょうか?',
          question: {
            id: '1',
            order: 0,
            questionType: 5,
            questionText: '問題の原因を可能な限り考察してください。',
            textQuestion: {
              id: '1',
              maxLength: 100,
              questionId: '1',
            },
            formId: '1',
            isRequired: true,
            forAiProcessing: true,
          },
        },
        {
          id: '1',
          responseId: '1',
          answerText: 'Age = 5\nprint("My Age: " + Age)',
          question: {
            id: '1',
            order: 0,
            questionType: 6,
            questionText: '問題のコードををペーストしてください。',
            textQuestion: {
              id: '1',
              maxLength: 100,
              questionId: '1',
            },
            formId: '1',
            isRequired: true,
            forAiProcessing: true,
          },
        },
      ],
      aiResponse: "まず、Pythonの型について理解が必要だね。Pythonは動的型付け言語で、型の自動解決をしてくれるというのは正しいけど、文字列と数値を直接結合することはできないんだ。だから、'MyAge: ' + Ageという式でエラーが出ているんだよ。これを解決するためには、数値を文字列に変換することが必要だよ。Pythonでは、str()という関数を使って数値を文字列に変換することができるんだ。例えば、str(Age)とすると、Ageの値が文字列の値に変換されるよ。これをヒントに、もう一度コードを見直してみてね！頑張って！"
    },
  ];
  return (
    <Accordion type="multiple" className="w-full">
      {responseList.map(response => (
        <AccordionItem key={response.id} value="河口欣仁#1">
          <AccordionTrigger>河口欣仁#1</AccordionTrigger>
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
