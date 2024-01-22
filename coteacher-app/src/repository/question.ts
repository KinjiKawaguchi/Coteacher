import {
  createPromiseClient,
  PromiseClient,
  Transport,
} from '@bufbuild/connect';
import { backendGrpcTransport } from '@/config/connectRpc';
import {
  Question as INTERFACE_Question,
  Option as INTERFACE_Option,
  TextQuestion as INTERFACE_TextQuestion,
} from '@/interfaces';
import {
  GetQuestionListByFormIdRequest,
  SaveQuestionListRequest,
} from '@/gen/proto/coteacher/v1/question_pb';
import { QuestionService } from '@/gen/proto/coteacher/v1/question_connect';
import {
  Question,
  Question_TextQuestion,
  QuestionOption,
} from '@/gen/proto/coteacher/v1/resources_pb';

class QuestionRepository {
  readonly cli: PromiseClient<typeof QuestionService>;

  constructor(t: Transport) {
    this.cli = createPromiseClient(QuestionService, t);
  }

  async getQuestionListByFormId(formId: string): Promise<INTERFACE_Question[]> {
    const req = new GetQuestionListByFormIdRequest();
    req.formId = formId;
    const res = await this.cli.getQuestionListByFormId(req);

    const questions: INTERFACE_Question[] = res.questions.map(q => {
      let textQuestion: INTERFACE_TextQuestion | undefined;
      if (q.textQuestion) {
        textQuestion = {
          id: q.textQuestion.id,
          questionId: q.textQuestion.questionId,
          maxLength: q.textQuestion.maxLength,
        };
      }

      let options: INTERFACE_Option[] | undefined;
      if (q.options && q.options.length) {
        options = q.options.map(o => ({
          id: o.id,
          questionId: o.questionId,
          optionText: o.optionText,
          order: o.order,
        }));
      }

      return {
        id: q.id,
        formId: q.formId,
        questionType: q.questionType,
        questionText: q.questionText,
        textQuestion: textQuestion,
        options: options,
        isRequired: q.isRequired,
        forAiProcessing: q.forAiProcessing,
        order: q.order,
      };
    });
    return questions;
  }

  async saveQuestionList(
    questionList: INTERFACE_Question[]
  ): Promise<INTERFACE_Question[]> {
    const req = new SaveQuestionListRequest();
    req.questions = questionList.map(q => {
      const question = new Question();
      question.id = q.id || '';
      question.formId = q.formId;
      question.questionType = q.questionType;
      question.questionText = q.questionText;
      question.isRequired = q.isRequired;
      question.forAiProcessing = q.forAiProcessing;
      question.order = q.order;

      if (q.textQuestion) {
        const textQuestion = new Question_TextQuestion();
        textQuestion.id = q.textQuestion.id || ''; // デフォルト値を設定
        textQuestion.questionId = q.textQuestion.questionId || ''; // デフォルト値を設定
        textQuestion.maxLength = q.textQuestion.maxLength;
        question.textQuestion = textQuestion;
      }

      if (q.options && q.options.length) {
        question.options = q.options.map(o => {
          const option = new QuestionOption();
          option.id = o.id || '';
          option.questionId = o.questionId || '';
          option.optionText = o.optionText;
          option.order = o.order;
          return option;
        });
      }

      return question;
    });
    const res = await this.cli.saveQuestionList(req);
    const questions: INTERFACE_Question[] = res.questions.map(q => {
      let textQuestion: INTERFACE_TextQuestion | undefined;
      if (q.textQuestion) {
        textQuestion = {
          id: q.textQuestion.id,
          questionId: q.textQuestion.questionId,
          maxLength: q.textQuestion.maxLength,
        };
      }

      let options: INTERFACE_Option[] | undefined;
      if (q.options && q.options.length) {
        options = q.options.map(o => ({
          id: o.id,
          questionId: o.questionId,
          optionText: o.optionText,
          order: o.order,
        }));
      }

      return {
        id: q.id,
        formId: q.formId,
        questionType: q.questionType,
        questionText: q.questionText,
        textQuestion: textQuestion,
        options: options,
        isRequired: q.isRequired,
        forAiProcessing: q.forAiProcessing,
        order: q.order,
      };
    });

    return questions;
  }
}

export const questionRepo = new QuestionRepository(backendGrpcTransport);
