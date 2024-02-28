from coteacher.v1 import resources_pb2 as _resources_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetQuestionListByFormIdRequest(_message.Message):
    __slots__ = ("form_id",)
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    form_id: str
    def __init__(self, form_id: _Optional[str] = ...) -> None: ...

class GetQuestionListByFormIdResponse(_message.Message):
    __slots__ = ("questions",)
    QUESTIONS_FIELD_NUMBER: _ClassVar[int]
    questions: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Question]
    def __init__(self, questions: _Optional[_Iterable[_Union[_resources_pb2.Question, _Mapping]]] = ...) -> None: ...

class SaveQuestionListRequest(_message.Message):
    __slots__ = ("questions",)
    QUESTIONS_FIELD_NUMBER: _ClassVar[int]
    questions: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Question]
    def __init__(self, questions: _Optional[_Iterable[_Union[_resources_pb2.Question, _Mapping]]] = ...) -> None: ...

class SaveQuestionListResponse(_message.Message):
    __slots__ = ("questions",)
    QUESTIONS_FIELD_NUMBER: _ClassVar[int]
    questions: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Question]
    def __init__(self, questions: _Optional[_Iterable[_Union[_resources_pb2.Question, _Mapping]]] = ...) -> None: ...
