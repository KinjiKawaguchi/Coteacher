import { Answer, Form, Question } from '@/interfaces';
import OpenAI from 'openai';

// グローバルに一度だけOpenAIインスタンスを作成
const openai = new OpenAI({
  apiKey: process.env.NEXT_PUBLIC_OPENAI_API_KEY,
  dangerouslyAllowBrowser: true, // ***注意*** クライアントサイドの実行を許可
});

function createContent(
  form: Form,
  questionList: Question[],
  answerList: Answer[]
): string {
  const questionsContent = questionList
    .filter(question => question.order !== -1 && question.forAiProcessing)
    .map(question => {
      // 回答を order プロパティに基づいて探す
      const answer = answerList.find(answer => answer.order === question.order);
      return `{
      question_type: '${question.questionType}',
      question_text: '${question.questionText}',
      answer: '${answer ? answer.answerText : '回答なし'}'
    }`;
    })
    .join(',');

  return `{
    instructions: '${form.systemPrompt}',
    able_conversation: false,
    questions: [${questionsContent}]
  }`;
}
// OpenAI APIを呼び出す関数
export async function callOpenAI(
  form: Form,
  questionList: Question[],
  answerList: Answer[]
) {
  const content = createContent(form, questionList, answerList);

  const userName = localStorage.getItem('UserName') || '';
  // OpenAIのAPIリクエスト

  console.log(content);
  const response = await openai.chat.completions.create({
    model: 'gpt-4',
    messages: [
      {
        role: 'system',
        content:
          'あなたは、多くの生徒を持つ先生が拾いきれない疑問、質問、意見などに対して代わりに生徒と対話を行います。\n' +
          '先生が用意したAIに提供する情報の生徒からの入力は以下のフォーマットによって与えられます。\n' +
          '{\n' +
          'instructions:"先生からの指示"\n' +
          'able_conversation:一問一答かどうか\n' +
          'questions{\n' +
          '[question_type:AIの入力のために先生が用意した質問のタイプ\n' +
          'question_text:AIの入力のために先生が用意した質問\n' +
          'answer:その質問に対する生徒の回答\n' +
          ']}\n' +
          '}\n' +
          'あなたの最も重要視すべきことは、生徒が理解し、定着し、応用できるような学習を促すことと、生徒の学習意欲や自主性を高めることです。' +
          '生徒の理解度や意欲を常に把握し、それに応じた対応をしましょう。' +
          '生徒が自分で考え、答えを導き出すことができるよう、ヒントやサポートを与えましょう。' +
          '生徒が学ぶ楽しさや喜びを感じられるよう、コミュニケーションを大切にしましょう。' +
          'とにかく元気にポジティブな返答をしましょう。',
      },
      {
        role: 'user',
        content: content,
      },
      {
        role: 'assistant',
        content:
          userName +
          'さん、先生の補助をする、元気でポジティブなGPT先生だよ！！！よろしく！！\nただし解決のための直接的な回答は出来ないから、これから教えるヒントを活かしてがんばって！！',
      },
    ],
    temperature: 2,
    max_tokens: 5000,
    top_p: 0.5,
    frequency_penalty: 0,
    presence_penalty: 0,
  });
  console.log(response);
  return response;
}
