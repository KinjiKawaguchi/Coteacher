// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts,import_extension=none"
// @generated from file coteacher/v1/student_class.proto (package coteacher.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateStudentClassRequest, CreateStudentClassResponse, DeleteStudentClassRequest, DeleteStudentClassResponse, GetClassListByStudentIDRequest, GetClassListByStudentIDResponse, GetStudentListByClassIDRequest, GetStudentListByClassIDResponse } from "./student_class_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service coteacher.v1.StudentClassService
 */
export const StudentClassService = {
  typeName: "coteacher.v1.StudentClassService",
  methods: {
    /**
     * @generated from rpc coteacher.v1.StudentClassService.CreateStudentClass
     */
    createStudentClass: {
      name: "CreateStudentClass",
      I: CreateStudentClassRequest,
      O: CreateStudentClassResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc coteacher.v1.StudentClassService.GetStudentListByClassID
     */
    getStudentListByClassID: {
      name: "GetStudentListByClassID",
      I: GetStudentListByClassIDRequest,
      O: GetStudentListByClassIDResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc coteacher.v1.StudentClassService.GetClassListByStudentID
     */
    getClassListByStudentID: {
      name: "GetClassListByStudentID",
      I: GetClassListByStudentIDRequest,
      O: GetClassListByStudentIDResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc coteacher.v1.StudentClassService.DeleteStudentClass
     */
    deleteStudentClass: {
      name: "DeleteStudentClass",
      I: DeleteStudentClassRequest,
      O: DeleteStudentClassResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

