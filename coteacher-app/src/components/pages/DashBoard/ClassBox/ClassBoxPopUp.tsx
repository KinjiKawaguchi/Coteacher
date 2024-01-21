import React from 'react';
import { MdDeleteForever, MdOutlinePublish } from 'react-icons/md';
import { IoIosExit } from 'react-icons/io';
import { TiThMenu } from 'react-icons/ti';
import { classInvitationCodeRepo } from '@/repository/classInvitationCode';
import { studentRepo } from '@/repository/student';
import { classRepo } from '@/repository/class';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuShortcut,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import toast from '@/libs/utils/toast';

interface ClassBoxPopUpProps {
  classId: string;
}

const issueInvitationCode = async (classId: string) => {
  const cic = await classInvitationCodeRepo.issueClassInvitationCode(classId);
  toast({
    status: 'success',
    title: '招待コードを発行しました。',
    description: cic.invitationCode,
  });
};

const quitClass = async (classId: string) => {
  await studentRepo.quitClass(classId);
  toast({
    status: 'success',
    title: '授業から退出しました。',
  });
  window.location.reload(); // TODO: リロードせずに内容を変更
};

const deleteClass = async (classId: string) => {
  await classRepo.deleteClass(classId);
  toast({
    status: 'success',
    title: '授業を削除しました。',
  }); // リロード
  window.location.reload(); // TODO: リロードせずに内容を変更
};

const ClassBoxPopUp = ({ classId }: ClassBoxPopUpProps) => {
  const userType = localStorage.getItem('UserType');

  const handleMenuItemClick = async (
    e: React.MouseEvent<HTMLDivElement, MouseEvent>,
    action: () => Promise<void>
  ) => {
    e.stopPropagation(); // 親コンポーネントへのイベント伝播を停止
    await action(); // 非同期アクションを実行
  };

  const handleIssueInvitationCode = async () => {
    await issueInvitationCode(classId);
  };

  const handleQuitClass = async () => {
    await quitClass(classId);
  };

  const handleDeleteClass = async () => {
    await deleteClass(classId);
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger>
        <button aria-label="メニュー" className="icon-button">
          <TiThMenu />
        </button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56">
        <DropdownMenuLabel>授業操作</DropdownMenuLabel>
        <DropdownMenuGroup>
          {userType === '1' && (
            <DropdownMenuItem
              onClick={e => handleMenuItemClick(e, handleQuitClass)}
            >
              <IoIosExit className="mr-2 h-4 w-4" />
              <span>授業を退出</span>
              <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
            </DropdownMenuItem>
          )}
          {userType === '2' && (
            <>
              <DropdownMenuItem
                onClick={e => handleMenuItemClick(e, handleIssueInvitationCode)}
              >
                <MdOutlinePublish className="mr-2 h-4 w-4" />
                <span>招待コードを発行</span>
                <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
              </DropdownMenuItem>
              <DropdownMenuItem
                onClick={e => handleMenuItemClick(e, handleDeleteClass)}
              >
                <MdDeleteForever className="mr-2 h-4 w-4" />
                <span>授業を削除</span>
                <DropdownMenuShortcut>⌘S</DropdownMenuShortcut>
              </DropdownMenuItem>
            </>
          )}
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default ClassBoxPopUp;
