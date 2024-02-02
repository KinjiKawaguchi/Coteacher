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
};

const CreateQuestionDropdown: React.FC<CreateQuestionDropdownProps> = ({
  formId,
  questionList,
  setQuestionList,
  index,
}) => {
  const handleAddQuestionButtonClick = (
    index: number,
    questionType: Question_QuestionType
  ) => {
    const updatedList = [...questionList];

    let newQuestion: Question = {
      order: index + 1,
      questionType: questionType, // 質問を選べるようにする TODO: 質問を選べるように
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

    // 新しい質問をリストに追加
    updatedList.splice(index + 1, 0, newQuestion);

    // 追加した質問の後ろの質問の順序を更新
    for (let i = index + 2; i < updatedList.length; i++) {
      updatedList[i].order += 1;
    }

    setQuestionList(updatedList);
    console.log(updatedList);
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
