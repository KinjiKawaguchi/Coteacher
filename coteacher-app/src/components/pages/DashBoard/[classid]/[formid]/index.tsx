'use client';

import { isEqual } from 'lodash'; // lodashをインポート
import UserHeader from '@/components/layout/header/UserHeader';
import { Button } from '@/components/ui/button';
import { formRepo } from '@/repository/form';
import { questionRepo } from '@/repository/question';
import theme from '@/theme';
import {
  ChakraProvider,
  Spinner,
  Container,
  HStack,
  Spacer,
} from '@chakra-ui/react';
import { ArrowLeft } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Question, Form } from '@/interfaces';
import QuestionList from './QuestionList';
import EditQuestionList from './EditQuestionList';
import toast from '@/libs/utils/toast';
import FormSetting from './FormSetting';
import ResponseList from './ResponseList';

export default function FormView({ params }: { params: { formid: string } }) {
  const [hasEditPermission, setHasEditPermission] = useState<boolean | null>(
    null
  );
  const [hasViewPermission, setHasViewPermission] = useState<boolean | null>(
    null
  );
  const [isLoading, setIsLoading] = useState(true);
  const [form, setForm] = useState<Form | null>(null);
  const [questionList, setQuestionList] = useState<Question[]>([]);
  const [remoteQuestionList, setRemoteQuestionList] = useState<Question[]>([]);
  const [showSaveButton, setShowSaveButton] = useState(false); // 保存ボタン表示用の状態
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      setHasEditPermission(
        await formRepo.checkFormEditPermission(params.formid)
      );
      setHasViewPermission(
        await formRepo.checkFormViewPermission(params.formid)
      );
      setForm(await formRepo.getFormById(params.formid));

      const fetchedQuestions = await questionRepo.getQuestionListByFormId(
        params.formid
      );
      setRemoteQuestionList(fetchedQuestions);
      const sortedQuestions = fetchedQuestions.sort(
        (a, b) => a.order - b.order
      ); // orderに基づいてソート
      setQuestionList(sortedQuestions);
      setIsLoading(false);
    };

    fetchData();
  }, []);

  // remoteQuestionList と questionList を比較する useEffect
  useEffect(() => {
    setShowSaveButton(!isEqual(remoteQuestionList, questionList));
  }, [remoteQuestionList, questionList]);

  if (hasEditPermission === null || hasViewPermission === null || isLoading) {
    return <Spinner />;
  }

  if (!hasEditPermission && !hasViewPermission) {
    return <div>アクセス権限がありません</div>;
  }

  const router = useRouter();
  const navBack = () => router.back();

  // 保存ボタンのクリックハンドラー（ここに保存ロジックを実装）
  const handleSave = async () => {
    setSaving(true);
    const savedQuestionList = await questionRepo.saveQuestionList(questionList);
    setRemoteQuestionList(savedQuestionList);
    setQuestionList(savedQuestionList);
    toast({
      title: '保存しました',
      status: 'success',
    });
    setSaving(false);
  };

  return (
    <ChakraProvider theme={theme}>
      <HStack>
        <Button variant="ghost" onClick={navBack}>
          <ArrowLeft />
        </Button>
        <Spacer />
        {showSaveButton &&
          (saving ? (
            <Spinner />
          ) : (
            <Button onClick={handleSave}>保存する</Button>
          ))}
        <Spacer />
      </HStack>
      <Container maxWidth="container.sm">
        <UserHeader />
        <Tabs defaultValue="question">
          <TabsList>
            <TabsTrigger value="question">{form?.name}</TabsTrigger>
            {hasEditPermission && (
              <>
                <TabsTrigger value="edit">編集</TabsTrigger>
                <TabsTrigger value="response">回答一覧</TabsTrigger>
                <TabsTrigger value="setting">設定</TabsTrigger>
              </>
            )}
          </TabsList>
          <TabsContent value="question">
            <QuestionList
              questionList={questionList}
              setQuestionList={setQuestionList}
              hasViewPermission={hasViewPermission}
              form={form}
            />
          </TabsContent>
          {hasEditPermission && (
            <>
              <TabsContent value="edit">
                <EditQuestionList
                  formId={params.formid}
                  questionList={questionList}
                  setQuestionList={setQuestionList}
                />
              </TabsContent>
              <TabsContent value="response">
                <ResponseList params={params} />
              </TabsContent>
              <TabsContent value="setting">
                {form && <FormSetting form={form} setForm={setForm} />}
              </TabsContent>
            </>
          )}
        </Tabs>
      </Container>
    </ChakraProvider>
  );
}
