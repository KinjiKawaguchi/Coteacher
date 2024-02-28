from coteacher.v1 import resources_pb2 as _resources_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateStudentClassRequest(_message.Message):
    __slots__ = ("student_id", "class_id")
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    student_id: str
    class_id: str
    def __init__(self, student_id: _Optional[str] = ..., class_id: _Optional[str] = ...) -> None: ...

class CreateStudentClassResponse(_message.Message):
    __slots__ = ("student_class",)
    STUDENT_CLASS_FIELD_NUMBER: _ClassVar[int]
    student_class: _resources_pb2.StudentClass
    def __init__(self, student_class: _Optional[_Union[_resources_pb2.StudentClass, _Mapping]] = ...) -> None: ...

class GetStudentListByClassIDRequest(_message.Message):
    __slots__ = ("class_id",)
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    class_id: str
    def __init__(self, class_id: _Optional[str] = ...) -> None: ...

class GetStudentListByClassIDResponse(_message.Message):
    __slots__ = ("students",)
    STUDENTS_FIELD_NUMBER: _ClassVar[int]
    students: _containers.RepeatedCompositeFieldContainer[_resources_pb2.User]
    def __init__(self, students: _Optional[_Iterable[_Union[_resources_pb2.User, _Mapping]]] = ...) -> None: ...

class GetClassListByStudentIDRequest(_message.Message):
    __slots__ = ("student_id",)
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    student_id: str
    def __init__(self, student_id: _Optional[str] = ...) -> None: ...

class GetClassListByStudentIDResponse(_message.Message):
    __slots__ = ("classes",)
    CLASSES_FIELD_NUMBER: _ClassVar[int]
    classes: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Class]
    def __init__(self, classes: _Optional[_Iterable[_Union[_resources_pb2.Class, _Mapping]]] = ...) -> None: ...

class DeleteStudentClassRequest(_message.Message):
    __slots__ = ("student_id", "class_id")
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    student_id: str
    class_id: str
    def __init__(self, student_id: _Optional[str] = ..., class_id: _Optional[str] = ...) -> None: ...

class DeleteStudentClassResponse(_message.Message):
    __slots__ = ("student_class",)
    STUDENT_CLASS_FIELD_NUMBER: _ClassVar[int]
    student_class: _resources_pb2.StudentClass
    def __init__(self, student_class: _Optional[_Union[_resources_pb2.StudentClass, _Mapping]] = ...) -> None: ...
