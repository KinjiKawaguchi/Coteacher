from coteacher.v1 import resources_pb2 as _resources_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateClassInvitationCodeRequest(_message.Message):
    __slots__ = ("class_id", "expiration_date")
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    EXPIRATION_DATE_FIELD_NUMBER: _ClassVar[int]
    class_id: str
    expiration_date: _timestamp_pb2.Timestamp
    def __init__(self, class_id: _Optional[str] = ..., expiration_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class CreateClassInvitationCodeResponse(_message.Message):
    __slots__ = ("class_invitation_code",)
    CLASS_INVITATION_CODE_FIELD_NUMBER: _ClassVar[int]
    class_invitation_code: _resources_pb2.ClassInvitationCode
    def __init__(self, class_invitation_code: _Optional[_Union[_resources_pb2.ClassInvitationCode, _Mapping]] = ...) -> None: ...

class GetClassInvitationCodeByIDRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetClassInvitationCodeByIDResponse(_message.Message):
    __slots__ = ("class_invitation_code",)
    CLASS_INVITATION_CODE_FIELD_NUMBER: _ClassVar[int]
    class_invitation_code: _resources_pb2.ClassInvitationCode
    def __init__(self, class_invitation_code: _Optional[_Union[_resources_pb2.ClassInvitationCode, _Mapping]] = ...) -> None: ...

class GetClassInvitationCodeListByClassIDRequest(_message.Message):
    __slots__ = ("class_id",)
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    class_id: str
    def __init__(self, class_id: _Optional[str] = ...) -> None: ...

class GetClassInvitationCodeListByClassIDResponse(_message.Message):
    __slots__ = ("class_invitation_code_list",)
    CLASS_INVITATION_CODE_LIST_FIELD_NUMBER: _ClassVar[int]
    class_invitation_code_list: _containers.RepeatedCompositeFieldContainer[_resources_pb2.ClassInvitationCode]
    def __init__(self, class_invitation_code_list: _Optional[_Iterable[_Union[_resources_pb2.ClassInvitationCode, _Mapping]]] = ...) -> None: ...

class UpdateClassInvitationCodeRequest(_message.Message):
    __slots__ = ("id", "expiration_date", "is_active")
    ID_FIELD_NUMBER: _ClassVar[int]
    EXPIRATION_DATE_FIELD_NUMBER: _ClassVar[int]
    IS_ACTIVE_FIELD_NUMBER: _ClassVar[int]
    id: str
    expiration_date: _timestamp_pb2.Timestamp
    is_active: bool
    def __init__(self, id: _Optional[str] = ..., expiration_date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., is_active: bool = ...) -> None: ...

class UpdateClassInvitationCodeResponse(_message.Message):
    __slots__ = ("class_invitation_code",)
    CLASS_INVITATION_CODE_FIELD_NUMBER: _ClassVar[int]
    class_invitation_code: _resources_pb2.ClassInvitationCode
    def __init__(self, class_invitation_code: _Optional[_Union[_resources_pb2.ClassInvitationCode, _Mapping]] = ...) -> None: ...
