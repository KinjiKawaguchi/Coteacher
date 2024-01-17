import React, { useEffect, useState } from 'react';
import {
  ColumnDef,
  useReactTable,
  getCoreRowModel,
  flexRender,
} from '@tanstack/react-table';
import {
  Table,
  TableBody,
  TableCell,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { Form } from '@/interfaces';
import { formRepo } from '@/repository/form';
import { responseRepo } from '@/repository/response';

export type FormTableProps = {
  classId: string;
};

export const FormTable: React.FC<FormTableProps> = ({ classId }) => {
  const [forms, setForms] = useState<Form[]>([]);
  const userType = localStorage.getItem('UserType');

  useEffect(() => {
    let isMounted = true;

    const fetchForms = async () => {
      const formList = await formRepo.getFormListByClassId(classId);
      if (!isMounted) return;

      const fetchUsageData = async form => {
        if (userType === '1') {
          return await responseRepo.getNumberOfResponseByStudentId(
            null,
            form.id
          );
        } else if (userType === '2') {
          return await responseRepo.getNumberOfResponseByFormId(form.id);
        }
        return 0;
      };

      const formsWithUsage = await Promise.all(
        formList.map(async form => ({
          ...form,
          usage: await fetchUsageData(form),
        }))
      );

      setForms(formsWithUsage);
    };

    fetchForms();

    return () => {
      isMounted = false;
    };
  }, [classId]);

  const columns: ColumnDef<Form>[] = [
    {
      accessorKey: 'name',
      header: 'フォーム名',
      cell: ({ row }) => row.original.name,
    },
    {
      accessorKey: 'description',
      header: '説明',
      cell: ({ row }) => row.original.description,
    },
    {
      accessorKey: 'usageLimit',
      header: '使用制限回数',
      cell: ({ row }) => row.original.usageLimit,
    },
    // userTypeに応じて条件分岐したカラムを追加します
    ...(userType === '1'
      ? [
          {
            accessorKey: 'usage',
            header: '使用回数',
            cell: ({ row }) => row.original.usage,
          },
        ]
      : []),
    ...(userType === '2'
      ? [
          {
            accessorKey: 'usage',
            header: '総使用回数',
            cell: ({ row }) => row.original.usage,
          },
        ]
      : []),
  ];

  const table = useReactTable({
    data: forms,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div className="rounded-md border">
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map(headerGroup => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map(header => (
                <TableCell key={header.id}>
                  {header.isPlaceholder
                    ? null
                    : flexRender(
                        header.column.columnDef.header,
                        header.getContext()
                      )}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows.map(row => (
            <TableRow key={row.id}>
              {row.getVisibleCells().map(cell => (
                <TableCell key={cell.id}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
};
