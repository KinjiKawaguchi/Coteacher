// import React from 'react';
// import { Question } from '@/interfaces';
// import { Checkbox } from '@/components/ui/checkbox';

// type MultipleChoiceQuestionComponentProps = {
//   question: Question;
//   editable: boolean;
// };

// const MultipleChoiceQuestionComponent: React.FC<
//   MultipleChoiceQuestionComponentProps
// > = ({ question }) => {
//   const { questionText, options, isRequired } = question;

//   return (
//     <>
//       <div className="mb-2">
//         {questionText} {isRequired && <span className="text-red-500">*</span>}
//       </div>
//       {options?.map((option, index) => (
//         <div key={option.id || index} className="flex items-center space-x-2">
//           <Checkbox id={`option-${index}`} />
//           <label
//             htmlFor={`option-${index}`}
//             className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
//           >
//             {option.optionText}
//           </label>
//         </div>
//       ))}
//     </>
//   );
// };

// export default MultipleChoiceQuestionComponent;
