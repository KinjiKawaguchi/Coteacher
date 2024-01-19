import React from 'react';
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
import { useRouter } from 'next/navigation';

export type FormTableProps = {
  forms: Form[]; // 新しいプロパティを追加
};

export const FormTable: React.FC<FormTableProps> = ({ forms }) => {
  const userType = localStorage.getItem('UserType');

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

  const router = useRouter();

  const handleRowClick = (row: any) => {
    const formId = row.original.id; // formIdはrow.originalから取得
    router.push(`/dashboard/${formId}`);
  };

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
            <TableRow key={row.id} onClick={() => handleRowClick(row)}>
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
