from coteacher.v1 import resources_pb2 as _resources_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class CheckStudentExistsByIDRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class CheckStudentExistsByIDResponse(_message.Message):
    __slots__ = ("exists",)
    EXISTS_FIELD_NUMBER: _ClassVar[int]
    exists: bool
    def __init__(self, exists: bool = ...) -> None: ...

class CheckStudentExistsByEmailRequest(_message.Message):
    __slots__ = ("email",)
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    email: str
    def __init__(self, email: _Optional[str] = ...) -> None: ...

class CheckStudentExistsByEmailResponse(_message.Message):
    __slots__ = ("exists",)
    EXISTS_FIELD_NUMBER: _ClassVar[int]
    exists: bool
    def __init__(self, exists: bool = ...) -> None: ...

class ParticipateClassRequest(_message.Message):
    __slots__ = ("user_id", "invitaion_code")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    INVITAION_CODE_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    invitaion_code: str
    def __init__(self, user_id: _Optional[str] = ..., invitaion_code: _Optional[str] = ...) -> None: ...

class ParticipateClassResponse(_message.Message):
    __slots__ = ()
    CLASS_FIELD_NUMBER: _ClassVar[int]
    def __init__(self, **kwargs) -> None: ...

class QuitClassRequest(_message.Message):
    __slots__ = ("user_id", "class_id")
    USER_ID_FIELD_NUMBER: _ClassVar[int]
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    user_id: str
    class_id: str
    def __init__(self, user_id: _Optional[str] = ..., class_id: _Optional[str] = ...) -> None: ...

class QuitClassResponse(_message.Message):
    __slots__ = ()
    CLASS_FIELD_NUMBER: _ClassVar[int]
    def __init__(self, **kwargs) -> None: ...
