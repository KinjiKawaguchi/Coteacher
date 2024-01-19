import React, { useState } from 'react';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { FaPlus } from 'react-icons/fa';
import { CreateFormInput, formRepo } from '@/repository/form';
import toast from '@/libs/utils/toast';
import { DialogClose } from '@radix-ui/react-dialog';

// Propsの型定義
interface CreateFormDialogProps {
  params: { classid: string };
  updateFormList: () => Promise<void>;
}

const CreateFormDialog: React.FC<CreateFormDialogProps> = ({
  params,
  updateFormList,
}) => {
  const [formName, setFormName] = useState<string>('');
  const [description, setDescription] = useState<string>('');
  const [usageLimit, setUsageLimit] = useState<number>(0);

  const handleFormNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setFormName(event.target.value);
  };

  const handleDescriptionChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setDescription(event.target.value);
  };

  const handleUsageLimitChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setUsageLimit(Number(event.target.value));
  };

  const handleCreate = async () => {
    const createFormInput: CreateFormInput = {
      classId: params.classid,
      name: formName,
      description,
      usageLimit,
    };
    await formRepo.createForm(createFormInput);
    await updateFormList();
    toast({
      status: 'success',
      title: 'フォームを作成しました',
    });
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>
          <FaPlus />
          新しいフォームを作成
        </Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>新しいフォームを作成</DialogTitle>
          <DialogDescription>
            入力された情報でフォームを作成します。
          </DialogDescription>
        </DialogHeader>
        <Input
          type="text"
          value={formName}
          placeholder="フォームの名前"
          onChange={handleFormNameChange}
        />
        <Input
          type="text"
          value={description}
          placeholder="フォームの説明"
          onChange={handleDescriptionChange}
        />
        <Input
          type="number"
          value={usageLimit}
          placeholder="回答数の上限"
          onChange={handleUsageLimitChange}
        />
        <DialogClose asChild>
          <Button onClick={handleCreate}>作成</Button>
        </DialogClose>
      </DialogContent>
    </Dialog>
  );
};

export default CreateFormDialog;
