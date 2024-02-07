// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts,import_extension=none"
// @generated from file coteacher/v1/response.proto (package coteacher.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetNumberOfResponsesByFormIDRequest, GetNumberOfResponsesByFormIDResponse, GetNumberOfResponsesByStudentIDRequest, GetNumberOfResponsesByStudentIDResponse, GetResponseListByFormIDRequest, GetResponseListByFormIDResponse, SubmitAIResponseRequest, SubmitAIResponseResponse, SubmitResponseRequest, SubmitResponseResponse } from "./response_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service coteacher.v1.ResponseService
 */
export const ResponseService = {
  typeName: "coteacher.v1.ResponseService",
  methods: {
    /**
     * @generated from rpc coteacher.v1.ResponseService.GetNumberOfResponsesByStudentID
     */
    getNumberOfResponsesByStudentID: {
      name: "GetNumberOfResponsesByStudentID",
      I: GetNumberOfResponsesByStudentIDRequest,
      O: GetNumberOfResponsesByStudentIDResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc coteacher.v1.ResponseService.GetNumberOfResponsesByFormID
     */
    getNumberOfResponsesByFormID: {
      name: "GetNumberOfResponsesByFormID",
      I: GetNumberOfResponsesByFormIDRequest,
      O: GetNumberOfResponsesByFormIDResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc coteacher.v1.ResponseService.GetResponseListByFormID
     */
    getResponseListByFormID: {
      name: "GetResponseListByFormID",
      I: GetResponseListByFormIDRequest,
      O: GetResponseListByFormIDResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc coteacher.v1.ResponseService.SubmitResponse
     */
    submitResponse: {
      name: "SubmitResponse",
      I: SubmitResponseRequest,
      O: SubmitResponseResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc coteacher.v1.ResponseService.SubmitAIResponse
     */
    submitAIResponse: {
      name: "SubmitAIResponse",
      I: SubmitAIResponseRequest,
      O: SubmitAIResponseResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

