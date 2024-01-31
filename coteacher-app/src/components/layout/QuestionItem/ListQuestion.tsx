import React from 'react';
import { Question } from '@/interfaces';
// import {
//   Select,
//   SelectContent,
//   SelectGroup,
//   SelectItem,
//   SelectTrigger,
//   SelectValue,
// } from '@/components/ui/select';

type ListQuestionComponentProps = {
  question: Question;
  editable: boolean;
};

const ListQuestionComponent: React.FC<ListQuestionComponentProps> = (
  {
    // question,
    //editable,
  }
) => {
  // const { questionText, options, isRequired } = question;

  return (
    <>
      {/* <div className="mb-2">
        {questionText} {isRequired && <span className="text-red-500">*</span>}
      </div>
      <Select>
        <SelectTrigger className="w-[180px]">
          <SelectValue placeholder="Select an option" />
        </SelectTrigger>
        <SelectContent>
          <SelectGroup>
            {options?.map((option, index) => (
              <SelectItem key={option.id || index} value={option.optionText}>
                {option.optionText}
              </SelectItem>
            ))}
          </SelectGroup>
        </SelectContent>
      </Select> */}
    </>
  );
};

export default ListQuestionComponent;
