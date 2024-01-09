import React from 'react';
import {
  Popover,
  PopoverTrigger,
  PopoverBody,
  PopoverContent,
  Portal,
  IconButton,
  Button,
} from '@chakra-ui/react';
import { FaEllipsisH, FaPassport, FaTrashAlt } from 'react-icons/fa';
import { classInvitationCodeRepo } from '@/repository/classInvitationCode';
import { studentRepo } from '@/repository/student';
import { classRepo } from '@/repository/class';

interface ClassBoxPopUpProps {
  classId: string;
}

const issueInvitationCode = async (classId: string) => {
  const cic = await classInvitationCodeRepo.issueClassInvitationCode(classId);
  alert(cic.invitationCode);
};

const quitClass = async (classId: string) => {
  await studentRepo.quitClass(classId);
  alert('授業から退出しました。');
  // リロード
  window.location.reload(); // TODO: リロードせずに内容を変更
};

const deleteClass = async (classId: string) => {
  await classRepo.deleteClass(classId);
  alert('授業を削除しました。');
  // リロード
  window.location.reload(); // TODO: リロードせずに内容を変更
};

const ClassBoxPopUp = ({ classId }: ClassBoxPopUpProps) => {
  const userType = localStorage.getItem('UserType');

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
    <Popover>
      <PopoverTrigger>
        <IconButton aria-label="Search database" icon={<FaEllipsisH />} />
      </PopoverTrigger>
      <Portal>
        <PopoverContent>
          <PopoverBody>
            {userType === '1' && (
              <Button
                leftIcon={<FaTrashAlt />}
                colorScheme="teal"
                variant="solid"
                onClick={handleQuitClass}
              >
                授業から退出
              </Button>
            )}
            {userType === '2' && (
              <>
                <Button
                  leftIcon={<FaPassport />}
                  colorScheme="teal"
                  variant="solid"
                  onClick={handleIssueInvitationCode}
                >
                  授業招待コードを発行
                </Button>
                <Button
                  leftIcon={<FaTrashAlt />}
                  colorScheme="teal"
                  variant="solid"
                  onClick={handleDeleteClass}
                >
                  授業を削除
                </Button>
              </>
            )}
          </PopoverBody>
        </PopoverContent>
      </Portal>
    </Popover>
  );
};

export default ClassBoxPopUp;
