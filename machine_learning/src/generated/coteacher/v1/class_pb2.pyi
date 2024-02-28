from coteacher.v1 import resources_pb2 as _resources_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateClassRequest(_message.Message):
    __slots__ = ("name", "teacher_id")
    NAME_FIELD_NUMBER: _ClassVar[int]
    TEACHER_ID_FIELD_NUMBER: _ClassVar[int]
    name: str
    teacher_id: str
    def __init__(self, name: _Optional[str] = ..., teacher_id: _Optional[str] = ...) -> None: ...

class CreateClassResponse(_message.Message):
    __slots__ = ()
    CLASS_FIELD_NUMBER: _ClassVar[int]
    def __init__(self, **kwargs) -> None: ...

class GetClassByIDRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetClassByIDResponse(_message.Message):
    __slots__ = ()
    CLASS_FIELD_NUMBER: _ClassVar[int]
    def __init__(self, **kwargs) -> None: ...

class GetClassListByTeacherIDRequest(_message.Message):
    __slots__ = ("teacher_id",)
    TEACHER_ID_FIELD_NUMBER: _ClassVar[int]
    teacher_id: str
    def __init__(self, teacher_id: _Optional[str] = ...) -> None: ...

class GetClassListByTeacherIDResponse(_message.Message):
    __slots__ = ("classes",)
    CLASSES_FIELD_NUMBER: _ClassVar[int]
    classes: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Class]
    def __init__(self, classes: _Optional[_Iterable[_Union[_resources_pb2.Class, _Mapping]]] = ...) -> None: ...

class UpdateClassRequest(_message.Message):
    __slots__ = ("id", "name", "teacher_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    TEACHER_ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    teacher_id: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., teacher_id: _Optional[str] = ...) -> None: ...

class UpdateClassResponse(_message.Message):
    __slots__ = ()
    CLASS_FIELD_NUMBER: _ClassVar[int]
    def __init__(self, **kwargs) -> None: ...

class DeleteClassRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteClassResponse(_message.Message):
    __slots__ = ()
    CLASS_FIELD_NUMBER: _ClassVar[int]
    def __init__(self, **kwargs) -> None: ...

class CheckClassEditPermissionRequest(_message.Message):
    __slots__ = ("class_id", "teacher_id")
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    TEACHER_ID_FIELD_NUMBER: _ClassVar[int]
    class_id: str
    teacher_id: str
    def __init__(self, class_id: _Optional[str] = ..., teacher_id: _Optional[str] = ...) -> None: ...

class CheckClassEditPermissionResponse(_message.Message):
    __slots__ = ("has_permission",)
    HAS_PERMISSION_FIELD_NUMBER: _ClassVar[int]
    has_permission: bool
    def __init__(self, has_permission: bool = ...) -> None: ...

class CheckClassViewPermissionRequest(_message.Message):
    __slots__ = ("class_id", "student_id")
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    class_id: str
    student_id: str
    def __init__(self, class_id: _Optional[str] = ..., student_id: _Optional[str] = ...) -> None: ...

class CheckClassViewPermissionResponse(_message.Message):
    __slots__ = ("has_permission",)
    HAS_PERMISSION_FIELD_NUMBER: _ClassVar[int]
    has_permission: bool
    def __init__(self, has_permission: bool = ...) -> None: ...
