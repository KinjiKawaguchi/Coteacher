from coteacher.v1 import resources_pb2 as _resources_pb2
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class UserType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = ()
    USER_TYPE_UNSPECIFIED: _ClassVar[UserType]
    USER_TYPE_STUDENT: _ClassVar[UserType]
    USER_TYPE_TEACHER: _ClassVar[UserType]
USER_TYPE_UNSPECIFIED: UserType
USER_TYPE_STUDENT: UserType
USER_TYPE_TEACHER: UserType

class CreateUserRequest(_message.Message):
    __slots__ = ("name", "email", "user_type")
    NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    name: str
    email: str
    user_type: UserType
    def __init__(self, name: _Optional[str] = ..., email: _Optional[str] = ..., user_type: _Optional[_Union[UserType, str]] = ...) -> None: ...

class CreateUserResponse(_message.Message):
    __slots__ = ("user", "user_type")
    USER_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    user: _resources_pb2.User
    user_type: UserType
    def __init__(self, user: _Optional[_Union[_resources_pb2.User, _Mapping]] = ..., user_type: _Optional[_Union[UserType, str]] = ...) -> None: ...

class GetUserByIDRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class GetUserByIDResponse(_message.Message):
    __slots__ = ("user", "user_type")
    USER_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    user: _resources_pb2.User
    user_type: UserType
    def __init__(self, user: _Optional[_Union[_resources_pb2.User, _Mapping]] = ..., user_type: _Optional[_Union[UserType, str]] = ...) -> None: ...

class GetUserByEmailRequest(_message.Message):
    __slots__ = ("email",)
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    email: str
    def __init__(self, email: _Optional[str] = ...) -> None: ...

class GetUserByEmailResponse(_message.Message):
    __slots__ = ("user", "user_type")
    USER_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    user: _resources_pb2.User
    user_type: UserType
    def __init__(self, user: _Optional[_Union[_resources_pb2.User, _Mapping]] = ..., user_type: _Optional[_Union[UserType, str]] = ...) -> None: ...

class UpdateUserRequest(_message.Message):
    __slots__ = ("id", "name", "email", "user_type")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    email: str
    user_type: UserType
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., email: _Optional[str] = ..., user_type: _Optional[_Union[UserType, str]] = ...) -> None: ...

class UpdateUserResponse(_message.Message):
    __slots__ = ("user", "user_type")
    USER_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    user: _resources_pb2.User
    user_type: UserType
    def __init__(self, user: _Optional[_Union[_resources_pb2.User, _Mapping]] = ..., user_type: _Optional[_Union[UserType, str]] = ...) -> None: ...

class DeleteUserRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class DeleteUserResponse(_message.Message):
    __slots__ = ("user", "user_type")
    USER_FIELD_NUMBER: _ClassVar[int]
    USER_TYPE_FIELD_NUMBER: _ClassVar[int]
    user: _resources_pb2.User
    user_type: UserType
    def __init__(self, user: _Optional[_Union[_resources_pb2.User, _Mapping]] = ..., user_type: _Optional[_Union[UserType, str]] = ...) -> None: ...

class CheckUserExistsByEmailRequest(_message.Message):
    __slots__ = ("email",)
    EMAIL_FIELD_NUMBER: _ClassVar[int]
    email: str
    def __init__(self, email: _Optional[str] = ...) -> None: ...

class CheckUserExistsByEmailResponse(_message.Message):
    __slots__ = ("exists",)
    EXISTS_FIELD_NUMBER: _ClassVar[int]
    exists: bool
    def __init__(self, exists: bool = ...) -> None: ...
