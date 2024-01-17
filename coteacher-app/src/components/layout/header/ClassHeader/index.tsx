import React, { useEffect, useState } from 'react';
import { HStack, Text, Box, Spinner, Center } from '@chakra-ui/react';
import { classRepo } from '@/repository/class';
import { userRepo } from '@/repository/user';
import { Class, User } from '@/interfaces';
import { FaChalkboardTeacher, FaEnvelope } from 'react-icons/fa';

export type ClassHeaderProps = {
  classId: string;
};

export const ClassHeader: React.FC<ClassHeaderProps> = ({ classId }) => {
  const [classData, setClassData] = useState<Class | null>(null);
  const [teacherData, setTeacherData] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true);
      try {
        const fetchedClassData = await classRepo.getClassById(classId);
        setClassData(fetchedClassData);
        if (fetchedClassData?.teacherId) {
          const fetchedTeacherData = await userRepo.getUserById(
            fetchedClassData.teacherId
          );
          setTeacherData(fetchedTeacherData);
        }
      } catch (error) {
        console.error('Error fetching data', error);
        // ここでエラーハンドリングを行う
      } finally {
        setIsLoading(false);
      }
    };
    fetchData();
  }, [classId]);

  if (isLoading) {
    return <Spinner />;
  }

  return (
    <Center>
      <HStack spacing={4}>
        {classData && (
          <Box>
            <Text fontSize="lg" fontWeight="bold">
              {classData.name}
            </Text>
          </Box>
        )}
        {teacherData && (
          <Box>
            <HStack>
              <FaChalkboardTeacher />
              <Text fontSize="lg"> {teacherData.name}</Text>
            </HStack>
            <HStack>
              <FaEnvelope />
              <Text fontSize="md">{teacherData.email}</Text>
            </HStack>
          </Box>
        )}
      </HStack>
    </Center>
  );
};
