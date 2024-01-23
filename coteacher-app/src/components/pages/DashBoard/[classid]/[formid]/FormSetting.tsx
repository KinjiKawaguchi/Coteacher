import { Form } from '@/interfaces';
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion';
import { isEqual } from 'lodash'; // lodashをインポート
import { HStack, Spacer, Textarea, VStack } from '@chakra-ui/react';
import { Label } from '@/components/ui/label';
import { useEffect, useState } from 'react';
import { Input } from '@/components/ui/input';
import { formRepo } from '@/repository/form';
import toast from '@/libs/utils/toast';
import { Button } from '@/components/ui/button';

interface FormSettingProps {
  form: Form;
  setForm: (form: Form) => void;
}

export default function FormSetting({ form, setForm }: FormSettingProps) {
  const [newForm, setNewForm] = useState<Form>(form);
  const [showSaveButton, setShowSaveButton] = useState(false); // 保存ボタン表示用の状態

  const handleFormNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewForm({ ...newForm, name: e.target.value });
  };

  const handleFormDescriptionChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setNewForm({ ...newForm, description: e.target.value });
  };

  const handleFormUsageLimitChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setNewForm({ ...newForm, usageLimit: parseInt(e.target.value) });
  };

  const handleFormSystemPromptChange = (
    e: React.ChangeEvent<HTMLTextAreaElement>
  ) => {
    setNewForm({ ...newForm, systemPrompt: e.target.value });
  };

  const handleSaveButtonClick = async () => {
    setForm(await formRepo.updateForm(newForm));
    toast({
      title: 'Success!',
      description: 'Form settings updated.',
      status: 'success',
    });
  };

  useEffect(() => {
    setShowSaveButton(!isEqual(form, newForm));
  }, [form, newForm]);

  return (
    <>
      <VStack>
        <Label>Form Name</Label>
        <Input value={newForm.name} onChange={handleFormNameChange} />
        <Label>Form Description</Label>
        <Input
          value={newForm.description}
          onChange={handleFormDescriptionChange}
        />
        <Label>Usage Limit</Label>
        <Input
          type="number"
          value={newForm.usageLimit}
          onChange={handleFormUsageLimitChange}
        />
        <Accordion type="single" collapsible className="w-full">
          <AccordionItem value="item-1">
            <AccordionTrigger>システムプロンプト</AccordionTrigger>
            <AccordionContent>
              <Textarea
                placeholder="システムプロンプト"
                value={newForm.systemPrompt}
                onChange={handleFormSystemPromptChange}
              />
            </AccordionContent>
          </AccordionItem>
        </Accordion>
      </VStack>
      {showSaveButton && (
        <HStack>
          <Spacer />
          <Button onClick={handleSaveButtonClick}>Save</Button>
        </HStack>
      )}
    </>
  );
}
