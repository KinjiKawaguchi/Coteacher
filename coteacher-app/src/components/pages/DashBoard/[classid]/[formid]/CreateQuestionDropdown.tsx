import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { Question_QuestionType } from '@/gen/proto/coteacher/v1/resources_pb';
import { Question } from '@/interfaces';
import { IoAddCircleOutline } from 'react-icons/io5';

type CreateQuestionDropdownProps = {
  formId: string;
  questionList: Question[];
  setQuestionList: (questionList: Question[]) => void;
  index: number;
  order: number;
};

const CreateQuestionDropdown: React.FC<CreateQuestionDropdownProps> = ({
  formId,
  questionList,
  setQuestionList,
  index,
  order,
}) => {
  const handleAddQuestionButtonClick = (
    index: number,
    questionType: Question_QuestionType
  ) => {
    const updatedList = [...questionList];

    let newQuestion: Question = {
      order: order, // 一時的なorder値; 後で更新
      questionType: questionType,
      questionText: '',
      isRequired: false,
      forAiProcessing: false,
      formId: formId,
    };

    // TEXT または PARAGRAPH_TEXT タイプの場合に textQuestion プロパティを追加
    if (
      questionType === Question_QuestionType.TEXT ||
      questionType === Question_QuestionType.PARAGRAPH_TEXT
    ) {
      newQuestion = {
        ...newQuestion,
        textQuestion: { maxLength: 150 },
      };
    } else {
      newQuestion = {
        ...newQuestion,
        options: [{ optionText: '', order: 0 }],
      };
    }

    // 先に既存のorderを更新
    for (let i = index; i < updatedList.length; i++) {
      updatedList[i].order = i + 1;
    }

    // 新しい質問をリストに挿入
    updatedList.splice(index, 0, newQuestion);

    // 質問リストを更新
    setQuestionList(updatedList);
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost">
          <IoAddCircleOutline size={24} />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuItem
          onClick={() =>
            handleAddQuestionButtonClick(index, Question_QuestionType.TEXT)
          }
        >
          テキスト
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() =>
            handleAddQuestionButtonClick(
              index,
              Question_QuestionType.PARAGRAPH_TEXT
            )
          }
        >
          パラグラフ
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() =>
            handleAddQuestionButtonClick(index, Question_QuestionType.RADIO)
          }
        >
          ラジオボタン
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() =>
            handleAddQuestionButtonClick(index, Question_QuestionType.CHECKBOX)
          }
        >
          チェックボックス
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() =>
            handleAddQuestionButtonClick(
              index,
              Question_QuestionType.MULTIPLE_CHOICE
            )
          }
        >
          複数選択
        </DropdownMenuItem>
        <DropdownMenuItem
          onClick={() =>
            handleAddQuestionButtonClick(index, Question_QuestionType.LIST)
          }
        >
          リスト
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default CreateQuestionDropdown;
