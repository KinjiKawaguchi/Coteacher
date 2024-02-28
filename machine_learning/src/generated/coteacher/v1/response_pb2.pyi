from coteacher.v1 import resources_pb2 as _resources_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetNumberOfResponsesByStudentIDRequest(_message.Message):
    __slots__ = ("student_id", "form_id")
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    student_id: str
    form_id: str
    def __init__(self, student_id: _Optional[str] = ..., form_id: _Optional[str] = ...) -> None: ...

class GetNumberOfResponsesByStudentIDResponse(_message.Message):
    __slots__ = ("number_of_responses",)
    NUMBER_OF_RESPONSES_FIELD_NUMBER: _ClassVar[int]
    number_of_responses: int
    def __init__(self, number_of_responses: _Optional[int] = ...) -> None: ...

class GetNumberOfResponsesByFormIDRequest(_message.Message):
    __slots__ = ("form_id",)
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    form_id: str
    def __init__(self, form_id: _Optional[str] = ...) -> None: ...

class GetNumberOfResponsesByFormIDResponse(_message.Message):
    __slots__ = ("number_of_responses",)
    NUMBER_OF_RESPONSES_FIELD_NUMBER: _ClassVar[int]
    number_of_responses: int
    def __init__(self, number_of_responses: _Optional[int] = ...) -> None: ...

class GetResponseListByFormIDRequest(_message.Message):
    __slots__ = ("form_id",)
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    form_id: str
    def __init__(self, form_id: _Optional[str] = ...) -> None: ...

class GetResponseListByFormIDResponse(_message.Message):
    __slots__ = ("responses", "questions", "students")
    RESPONSES_FIELD_NUMBER: _ClassVar[int]
    QUESTIONS_FIELD_NUMBER: _ClassVar[int]
    STUDENTS_FIELD_NUMBER: _ClassVar[int]
    responses: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Response]
    questions: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Question]
    students: _containers.RepeatedCompositeFieldContainer[_resources_pb2.User]
    def __init__(self, responses: _Optional[_Iterable[_Union[_resources_pb2.Response, _Mapping]]] = ..., questions: _Optional[_Iterable[_Union[_resources_pb2.Question, _Mapping]]] = ..., students: _Optional[_Iterable[_Union[_resources_pb2.User, _Mapping]]] = ...) -> None: ...

class SubmitResponseRequest(_message.Message):
    __slots__ = ("form_id", "student_id", "answers", "ai_response")
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    ANSWERS_FIELD_NUMBER: _ClassVar[int]
    AI_RESPONSE_FIELD_NUMBER: _ClassVar[int]
    form_id: str
    student_id: str
    answers: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Response.Answer]
    ai_response: str
    def __init__(self, form_id: _Optional[str] = ..., student_id: _Optional[str] = ..., answers: _Optional[_Iterable[_Union[_resources_pb2.Response.Answer, _Mapping]]] = ..., ai_response: _Optional[str] = ...) -> None: ...

class SubmitResponseResponse(_message.Message):
    __slots__ = ("success", "response_id")
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    RESPONSE_ID_FIELD_NUMBER: _ClassVar[int]
    success: bool
    response_id: str
    def __init__(self, success: bool = ..., response_id: _Optional[str] = ...) -> None: ...

class SubmitAIResponseRequest(_message.Message):
    __slots__ = ("response_id", "ai_response")
    RESPONSE_ID_FIELD_NUMBER: _ClassVar[int]
    AI_RESPONSE_FIELD_NUMBER: _ClassVar[int]
    response_id: str
    ai_response: str
    def __init__(self, response_id: _Optional[str] = ..., ai_response: _Optional[str] = ...) -> None: ...

class SubmitAIResponseResponse(_message.Message):
    __slots__ = ("success",)
    SUCCESS_FIELD_NUMBER: _ClassVar[int]
    success: bool
    def __init__(self, success: bool = ...) -> None: ...

class CreateDatasetRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class CreateDatasetResponse(_message.Message):
    __slots__ = ("ai_input", "ai_output")
    AI_INPUT_FIELD_NUMBER: _ClassVar[int]
    AI_OUTPUT_FIELD_NUMBER: _ClassVar[int]
    ai_input: _containers.RepeatedScalarFieldContainer[str]
    ai_output: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, ai_input: _Optional[_Iterable[str]] = ..., ai_output: _Optional[_Iterable[str]] = ...) -> None: ...
