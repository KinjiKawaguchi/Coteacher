import { Form, Question } from '@/interfaces';
import OpenAI from 'openai';

// グローバルに一度だけOpenAIインスタンスを作成
const openai = new OpenAI({
  apiKey: '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',
  dangerouslyAllowBrowser: true, // ***注意*** クライアントサイドの実行を許可
});

// ユーザー名の取得
const userName = '河口欣仁';

function createContent(
  form: Form,
  questionList: Question[],
  answerList: string[]
) {
  let content = `{\ninstructions: '${form.systemPrompt}',\nable_conversation: false,\nquestions: [`;

  questionList.forEach((question, index) => {
    if (question.order !== -1) {
      content += `{ \nquestion_type: '${question.questionType}',\nquestion_text: '${question.questionText}',\nanswer: '${answerList[index]}' }`;
      if (index < questionList.length - 1) content += ',';
    }
  });

  content += '\n]\n}';

  return content;
}

// OpenAI APIを呼び出す関数
export async function callOpenAI(
  form: Form,
  questionList: Question[],
  answerList: string[]
) {
  const content = createContent(form, questionList, answerList);

  // OpenAIのAPIリクエスト
  const response = await openai.chat.completions.create({
    model: 'gpt-4',
    messages: [
      {
        role: 'system',
        content:
          'あなたは、多くの生徒を持つ先生が拾いきれない疑問、質問、意見などに対して代わりに生徒と対話を行います。\n先生が用意したAIに提供する情報の生徒からの入力は以下のフォーマットによって与えられます。\n' +
          content,
      },
      {
        role: 'user',
        content: content,
      },
      {
        role: 'assistant',
        content:
          userName +
          'さん、先生の補助をする、GPT先生だよ、よろしく！！\nただし解決のための直接的な回答は出来ないから、これから教えるヒントを活かしてがんばって！！',
      },
    ],
    temperature: 2,
    max_tokens: 5000,
    top_p: 0.5,
    frequency_penalty: 0,
    presence_penalty: 0,
  });

  return response;
}
