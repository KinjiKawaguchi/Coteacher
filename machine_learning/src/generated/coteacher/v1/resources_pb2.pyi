from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class User(_message.Message):
    __slots__ = ("id", "name", "email", "teacher", "student", "created_at", "updated_at")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    TEACHER_FIELD_NUMBER: _ClassVar[int]
    STUDENT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    email: str
    teacher: Teacher
    student: Student
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., email: _Optional[str] = ..., teacher: _Optional[_Union[Teacher, _Mapping]] = ..., student: _Optional[_Union[Student, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class Teacher(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class Student(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ClassInvitationCode(_message.Message):
    __slots__ = ("id", "class_id", "invitation_code", "expiration_date", "is_active", "created_at", "updated_at")
    ID_FIELD_NUMBER: _ClassVar[int]
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    INVITATION_CODE_FIELD_NUMBER: _ClassVar[int]
    EXPIRATION_DATE_FIELD_NUMBER: _ClassVar[int]
    IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    class_id: str
    invitation_code: str
    expiration_date: _timestamp_pb2.Timestamp
    is_active: bool
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., class_id: _Optional[str] = ..., invitation_code: _Optional[str] = ..., expiration_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., is_active: bool = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class Class(_message.Message):
    __slots__ = ("id", "teacher_id", "name", "created_at", "updated_at")
    ID_FIELD_NUMBER: _ClassVar[int]
    TEACHER_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    teacher_id: str
    name: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., teacher_id: _Optional[str] = ..., name: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class StudentClass(_message.Message):
    __slots__ = ("student_id", "class_id", "created_at", "updated_at")
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    student_id: str
    class_id: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, student_id: _Optional[str] = ..., class_id: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class Form(_message.Message):
    __slots__ = ("id", "class_id", "name", "description", "usage_limit", "system_prompt", "created_at", "updated_at")
    ID_FIELD_NUMBER: _ClassVar[int]
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    USAGE_LIMIT_FIELD_NUMBER: _ClassVar[int]
    SYSTEM_PROMPT_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    class_id: str
    name: str
    description: str
    usage_limit: int
    system_prompt: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., class_id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., usage_limit: _Optional[int] = ..., system_prompt: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class Question(_message.Message):
    __slots__ = ("id", "form_id", "question_type", "question_text", "text_question", "options", "is_required", "for_ai_processing", "order", "created_at", "updated_at")
    class QuestionType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        QUESTION_TYPE_UNSPECIFIED: _ClassVar[Question.QuestionType]
        QUESTION_TYPE_CHECKBOX: _ClassVar[Question.QuestionType]
        QUESTION_TYPE_LIST: _ClassVar[Question.QuestionType]
        QUESTION_TYPE_RADIO: _ClassVar[Question.QuestionType]
        QUESTION_TYPE_MULTIPLE_CHOICE: _ClassVar[Question.QuestionType]
        QUESTION_TYPE_PARAGRAPH_TEXT: _ClassVar[Question.QuestionType]
        QUESTION_TYPE_TEXT: _ClassVar[Question.QuestionType]
    QUESTION_TYPE_UNSPECIFIED: Question.QuestionType
    QUESTION_TYPE_CHECKBOX: Question.QuestionType
    QUESTION_TYPE_LIST: Question.QuestionType
    QUESTION_TYPE_RADIO: Question.QuestionType
    QUESTION_TYPE_MULTIPLE_CHOICE: Question.QuestionType
    QUESTION_TYPE_PARAGRAPH_TEXT: Question.QuestionType
    QUESTION_TYPE_TEXT: Question.QuestionType
    class TextQuestion(_message.Message):
        __slots__ = ("id", "question_id", "max_length")
        ID_FIELD_NUMBER: _ClassVar[int]
        QUESTION_ID_FIELD_NUMBER: _ClassVar[int]
        MAX_LENGTH_FIELD_NUMBER: _ClassVar[int]
        id: str
        question_id: str
        max_length: int
        def __init__(self, id: _Optional[str] = ..., question_id: _Optional[str] = ..., max_length: _Optional[int] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    QUESTION_TYPE_FIELD_NUMBER: _ClassVar[int]
    QUESTION_TEXT_FIELD_NUMBER: _ClassVar[int]
    TEXT_QUESTION_FIELD_NUMBER: _ClassVar[int]
    OPTIONS_FIELD_NUMBER: _ClassVar[int]
    IS_REQUIRED_FIELD_NUMBER: _ClassVar[int]
    FOR_AI_PROCESSING_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    form_id: str
    question_type: Question.QuestionType
    question_text: str
    text_question: Question.TextQuestion
    options: _containers.RepeatedCompositeFieldContainer[QuestionOption]
    is_required: bool
    for_ai_processing: bool
    order: int
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., form_id: _Optional[str] = ..., question_type: _Optional[_Union[Question.QuestionType, str]] = ..., question_text: _Optional[str] = ..., text_question: _Optional[_Union[Question.TextQuestion, _Mapping]] = ..., options: _Optional[_Iterable[_Union[QuestionOption, _Mapping]]] = ..., is_required: bool = ..., for_ai_processing: bool = ..., order: _Optional[int] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class QuestionOption(_message.Message):
    __slots__ = ("id", "question_id", "option_text", "order")
    ID_FIELD_NUMBER: _ClassVar[int]
    QUESTION_ID_FIELD_NUMBER: _ClassVar[int]
    OPTION_TEXT_FIELD_NUMBER: _ClassVar[int]
    ORDER_FIELD_NUMBER: _ClassVar[int]
    id: str
    question_id: str
    option_text: str
    order: int
    def __init__(self, id: _Optional[str] = ..., question_id: _Optional[str] = ..., option_text: _Optional[str] = ..., order: _Optional[int] = ...) -> None: ...

class Response(_message.Message):
    __slots__ = ("id", "form_id", "student_id", "answers", "ai_response", "created_at", "updated_at")
    class Answer(_message.Message):
        __slots__ = ("id", "response_id", "question_id", "order", "answer_text", "answer_option_ids")
        ID_FIELD_NUMBER: _ClassVar[int]
        RESPONSE_ID_FIELD_NUMBER: _ClassVar[int]
        QUESTION_ID_FIELD_NUMBER: _ClassVar[int]
        ORDER_FIELD_NUMBER: _ClassVar[int]
        ANSWER_TEXT_FIELD_NUMBER: _ClassVar[int]
        ANSWER_OPTION_IDS_FIELD_NUMBER: _ClassVar[int]
        id: str
        response_id: str
        question_id: str
        order: int
        answer_text: str
        answer_option_ids: _containers.RepeatedScalarFieldContainer[str]
        def __init__(self, id: _Optional[str] = ..., response_id: _Optional[str] = ..., question_id: _Optional[str] = ..., order: _Optional[int] = ..., answer_text: _Optional[str] = ..., answer_option_ids: _Optional[_Iterable[str]] = ...) -> None: ...
    ID_FIELD_NUMBER: _ClassVar[int]
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    ANSWERS_FIELD_NUMBER: _ClassVar[int]
    AI_RESPONSE_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    id: str
    form_id: str
    student_id: str
    answers: _containers.RepeatedCompositeFieldContainer[Response.Answer]
    ai_response: str
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    def __init__(self, id: _Optional[str] = ..., form_id: _Optional[str] = ..., student_id: _Optional[str] = ..., answers: _Optional[_Iterable[_Union[Response.Answer, _Mapping]]] = ..., ai_response: _Optional[str] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...
