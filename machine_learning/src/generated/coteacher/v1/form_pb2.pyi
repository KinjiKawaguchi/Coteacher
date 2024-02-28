from coteacher.v1 import resources_pb2 as _resources_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class CreateFormRequest(_message.Message):
    __slots__ = ("class_id", "name", "description", "usage_limit")
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    USAGE_LIMIT_FIELD_NUMBER: _ClassVar[int]
    class_id: str
    name: str
    description: str
    usage_limit: int
    def __init__(self, class_id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., usage_limit: _Optional[int] = ...) -> None: ...

class CreateFormResponse(_message.Message):
    __slots__ = ("form",)
    FORM_FIELD_NUMBER: _ClassVar[int]
    form: _resources_pb2.Form
    def __init__(self, form: _Optional[_Union[_resources_pb2.Form, _Mapping]] = ...) -> None: ...

class GetFormByIDRequest(_message.Message):
    __slots__ = ("form_id",)
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    form_id: str
    def __init__(self, form_id: _Optional[str] = ...) -> None: ...

class GetFormByIDResponse(_message.Message):
    __slots__ = ("form",)
    FORM_FIELD_NUMBER: _ClassVar[int]
    form: _resources_pb2.Form
    def __init__(self, form: _Optional[_Union[_resources_pb2.Form, _Mapping]] = ...) -> None: ...

class GetFormListByClassIDRequest(_message.Message):
    __slots__ = ("class_id",)
    CLASS_ID_FIELD_NUMBER: _ClassVar[int]
    class_id: str
    def __init__(self, class_id: _Optional[str] = ...) -> None: ...

class GetFormListByClassIDResponse(_message.Message):
    __slots__ = ("forms",)
    FORMS_FIELD_NUMBER: _ClassVar[int]
    forms: _containers.RepeatedCompositeFieldContainer[_resources_pb2.Form]
    def __init__(self, forms: _Optional[_Iterable[_Union[_resources_pb2.Form, _Mapping]]] = ...) -> None: ...

class UpdateFormRequest(_message.Message):
    __slots__ = ("id", "name", "description", "usage_limit", "system_prompt")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    USAGE_LIMIT_FIELD_NUMBER: _ClassVar[int]
    SYSTEM_PROMPT_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    usage_limit: int
    system_prompt: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., usage_limit: _Optional[int] = ..., system_prompt: _Optional[str] = ...) -> None: ...

class UpdateFormResponse(_message.Message):
    __slots__ = ("form",)
    FORM_FIELD_NUMBER: _ClassVar[int]
    form: _resources_pb2.Form
    def __init__(self, form: _Optional[_Union[_resources_pb2.Form, _Mapping]] = ...) -> None: ...

class DeleteFormRequest(_message.Message):
    __slots__ = ("form_id",)
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    form_id: str
    def __init__(self, form_id: _Optional[str] = ...) -> None: ...

class DeleteFormResponse(_message.Message):
    __slots__ = ("form",)
    FORM_FIELD_NUMBER: _ClassVar[int]
    form: _resources_pb2.Form
    def __init__(self, form: _Optional[_Union[_resources_pb2.Form, _Mapping]] = ...) -> None: ...

class CheckFormEditPermissionRequest(_message.Message):
    __slots__ = ("form_id", "teacher_id")
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    TEACHER_ID_FIELD_NUMBER: _ClassVar[int]
    form_id: str
    teacher_id: str
    def __init__(self, form_id: _Optional[str] = ..., teacher_id: _Optional[str] = ...) -> None: ...

class CheckFormEditPermissionResponse(_message.Message):
    __slots__ = ("has_permission",)
    HAS_PERMISSION_FIELD_NUMBER: _ClassVar[int]
    has_permission: bool
    def __init__(self, has_permission: bool = ...) -> None: ...

class CheckFormViewPermissionRequest(_message.Message):
    __slots__ = ("form_id", "student_id")
    FORM_ID_FIELD_NUMBER: _ClassVar[int]
    STUDENT_ID_FIELD_NUMBER: _ClassVar[int]
    form_id: str
    student_id: str
    def __init__(self, form_id: _Optional[str] = ..., student_id: _Optional[str] = ...) -> None: ...

class CheckFormViewPermissionResponse(_message.Message):
    __slots__ = ("has_permission",)
    HAS_PERMISSION_FIELD_NUMBER: _ClassVar[int]
    has_permission: bool
    def __init__(self, has_permission: bool = ...) -> None: ...
