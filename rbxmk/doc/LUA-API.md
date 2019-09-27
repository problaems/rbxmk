[DRAFT]

# Lua API reference
This document describes the Lua API provided by rbxmk.

## Syntax
This document uses a syntax to describe the Lua API provided by rbxmk. This
section provides an overview of the syntax.

<details><summary>Expand</summary>

*Within this section, the `<>` syntax indicates a stand-in for literal text.*

There are several kinds of declarations:

- `module <module>`: Declares a set of values.
- `type <type>: <declaration>`: Declares a value type.
- `variable <var>: <type>`: Declares a value.
- `function <func><arguments><returns>`: Declares a Lua function.

`<module>` is the name of the module.

**Example:** `module rbxmk`

`<type>` is the name of the type. `<declaration>` describes the type.

- `type Foobar: string`: Declared as another type.
- `type Foobar: This | That`: Can be of type `This` or type `That`.
- `type Foobar: string?`: Nullable string. Shorthand for `string | nil`.
- `type Foobar: any`: Can be a value of any type.
- `type Foobar: []string`: An array of `string` values.
- `type Foobar: map[string]int`: A map of `string` keys to `int` values.
- `type Foobar: struct{field1: int; field2: string}`: A table with fields.
- `type Foobar: struct{~field1: int}`: `field1` is readonly.
- `type Foobar: struct{Embedded}`: The fields of struct `Embedded` are merged into `Foobar`.

**Example:** `type SourceOption: Config | API | Format`

`<var>` is the name of the variable. It may have a qualifier to indicate where
is is declared.

- `variable foobar: string`: Variable "foobar" of type `string`.
- `variable module.foobar: <type>`: Declared to be within `module`.

**Example:** `variable rbxmk.config: Config`

`<func>` Is the name of the function. It may have a qualifier to indicate where
is is declared.

- `function foobar()`: Function "foobar".
- `function module.foobar()`: Declared to be within `module`.
- `function Type:foobar()`: Declared to be a method on `Type`.

`<arguments>` is a list of parameters passed to the function, enclosed in
parentheses. `<returns>` is a list of parameters returned by the function. Each
parameter is a name paired with a type declaration.

- `function foobar(arg1: int, arg2: string)`: Two arguments of type `int` and
  `string`; nothing returned.
- `function foobar(arg1: int, arg2: ...string)`: One `int` argument, variable
  number of `string` arguments.
- `function foobar(arg1, arg2: int)`: Two `int` arguments.
- `function foobar(arg1: int = 42)`: Default value when the argument is
  unspecified. Implies the type is nullable.
- `function foobar() int`: Unnamed `int` return value.
- `function foobar() (int, string)`: Unnamed `int` and `string` return values.
- `function foobar() (ret1, ret2: int)`: Named `int` return values.
- `function foobar() (ret1 ...string)`: Named variable `string` return values.

Functions may be declared multiple times, indicating an alteration in the
arguments that can be passed to it. Arguments may also be a literal value,
indicating that the value is required for that alteration to be used:

- `function foobar("Foo", arg2: int)`
- `function foobar("Bar", arg2: string)`

A type is usually implemented as a userdata with a metatable, which use
metamethods to enforce type behaviors. If a type has a nonstandard behavior,
then its metamethod will be declared as `function Type:__metamethod()`. Standard
and nonstandard behaviors may be merged; e.g. a type declares a number of
methods (implemented using `__index`), but also declares `__index` directly to
describe a nonstandard behavior.

**Example:** `function Instance:FindFirstChild(name: string, recursive: bool = false) Instance?`
</details>

## Globals
The following values are present in the global environment:

Kind     | Name       | Description
---------|------------|------------
module   | `lua`      | Abridged version of the Lua standard library. Merged into _G.
module   | `roblox`   | Compatibility layer for the Roblox environment. Merged into _G.
module   | `types`    | Constructors for Roblox types.
module   | `rbxmk`    | Interface to rbxmk.
module   | `os`       | Additional functions for interacting with the operating system.
function | `type`     | Detects additional types.

## type function
`function type(value: any) string`

`type` returns a string indicating the type of the given value. In addition to
the standard Lua types, `type` also detects types defined by rbxmk.

## lua module
`module lua`

Abridged version of the Lua standard library. The contents of the module are
merged into the global environment.

Kind     | Name       | Description
---------|------------|------------
variable | `_G`       |
variable | `_VERSION` |
function | `assert`   |
function | `error`    |
function | `ipairs`   |
function | `next`     |
function | `pairs`    |
function | `pcall`    |
function | `print`    |
function | `select`   |
function | `tonumber` |
function | `tostring` |
function | `unpack`   |
function | `xpcall`   |
module   | `math`     |
module   | `string`   | Excludes `dump`.
module   | `table`    | Includes `pack` and `unpack` from 5.2.
module   | `os`       | Only includes `clock`, `date`, `difftime`, and `time`.


## roblox module
`module roblox`

Provides a compatibility layer for the Roblox Lua environment. The contents of
the module are merged into the global environment.

Kind     | Name                      | Description
---------|---------------------------|------------
function | `typeof`                  | Returns the type of a Roblox value.
variable | `Enum`                    | The set of Roblox enums.
...      | type constructors         | Constructors for each Roblox type.

### typeof function
`function typeof(value: any) string`

`typeof` detects types defined by the `roblox` module. Returns "userdata" for
unknown types, which includes types defined by rbxmk.

### Enum variable
`variable Enum: Enums`

`Enum` contains the enums defined by Roblox. The content depends on the globally
configured API value (`rbxmk.config.API`).

### Type constructors
The `roblox` module also contains the constructors for creating Roblox types,
such as `Vector3`, `CFrame`, and so on. For the convenience of writing this
document, they are not repeated here.

## rbxmk module
`module rbxmk`

Kind     | Name     | Description
---------|----------|------------
function | `read`   | Read from a source.
function | `write`  | Write to a source.
function | `load`   | Execute a script.
function | `cred`   | Generate credentials.
variable | `config` | Global configuration.

### read function
`function rbxmk.read(url: string, options: ...SourceOption) Data`

`read` retrieves data from a source. *options* may consist of any combination of
SourceOptions.

### write function
`function rbxmk.write(url: string, data: Data, options: ...SourceOption)`

`write` sends data to a source. *options* may consist of any combination of
SourceOptions.

### load function
`function rbxmk.load(path: string, args: ...any) ...any`

`load` runs the file pointed to by *path* as a Lua script.

### cred function
`function rbxmk.cred(type: string, args: ...any) Cred`

`cred` generates credentials to be used for interactions with sources that
require authentication. *type* indicates how the credentials will be generated,
and determines the remaining arguments, which usually involve an identifier and
a password. *type* is case-insensitive.

It is bad practice to pass a password directly as an argument. Instead, the
password should be received from a script argument, retrieved from an
environment variable using `os.getenv`, or omitted entirely, in which case the
user will be prompted.

#### None
`function rbxmk.cred() Cred`

Prompts for the credential type. Only works for `UserID`, `Username`, `Email`,
or `PhoneNumber`. Defaults to `Username` if an empty string is received.
Remaining arguments, dependent on the type, are also prompted.

#### UserID
`function rbxmk.cred("UserID", id: int?, password: string?) Cred`

Logs in with the user's numeric ID as the identifier. Omitted arguments are
prompted. Throws an error under the following circumstances:

- Prompted *id* cannot be parsed as an integer.
- *id* is less than 1.
- *password* is an empty string.

#### Username
`function rbxmk.cred("Username", name: string?, password: string?) Cred`

Logs in with the user's account name as the identifier. Omitted arguments are
prompted. Throws an error under the following circumstances:

- *name* is an empty string.
- *password* is an empty string.

#### Email
`function rbxmk.cred("Email", address: string?, password: string?) Cred`

Logs in with the user's email as the identifier. Omitted arguments are prompted.
Throws an error under the following circumstances:

- *address* is an empty string.
- *password* is an empty string.

#### PhoneNumber
`function rbxmk.cred("PhoneNumber", number: string?, password: string?) Cred`

Logs in with the user's phone number as the identifier. Omitted arguments are
prompted. Throws an error under the following circumstances:

- *number* is an empty string.
- *password* is an empty string.

#### Cookie
`function rbxmk.cred("Cookie", format: string, content: string) Cred`

Receives session cookies directly from a string. *format* indicates how
*content* should be decoded. The following formats are supported:

Format   | Description
---------|------------
Header   | Standard HTTP `Set-Cookie` headers.
Registry | Windows registry `.reg` format.
JSON     | JSON format.

#### Studio
`function rbxmk.cred("Studio") Cred`

Attempts to retrieve session cookies from Roblox Studio.

### config variable
`variable rbxmk.config: Config`

`config` is the global Config used by the environment.

## os module
`module os`

The `os` module has several additional functions for working with the operating
system.

### split function
`function os.split(path: string, parts: ...string) ...string`

`split` splits a file path into its component parts. *path* is the path to
split. *parts* is a number of strings indicating which parts of the path return.
Each returned value corresponds to the given part. Throws an error when an
invalid part value is received.

The following parts are available:

Part    | `project/scripts/main.script.lua` | Description
--------|-----------------------------------|------------
`dir`   | `project/scripts`                 | The directory; all but the last element of the path.
`base`  | `main.script.lua`                 | The file name; the last element of the path.
`ext`   | `.lua`                            | The extension; the suffix starting at the last `.` of the last element of the path.
`stem`  | `main.script`                     | The base without the extension.
`fext`  | `.script.lua`                     | The longest format extension that is known by rbxmk.
`fstem` | `main`                            | The base without the format extension.

### join function
`function os.join(elements: ...string) string`

`join` receives a number of path elements, and joins them into a single file
path, adding and normalizing separators as necessary.

An element may contain variables, which begin with a `$` followed by a sequence
of letters, digits, and underscores. Variables will be expanded into their final
values before the string is joined.

The following variables are available:

Variable                                 | Description
-----------------------------------------|------------
`script_directory`, `script_dir`, `sd`   | Expands to the directory of the script currently running.
`script_name`, `sn`                      | Expands to the base name of the script currently running.
`working_directory`, `working_dir`, `wd` | Expands to the current working directory.
`temp_directory`, `temp_dir`, `tmp`      | Expands to the directory for temporary files.

Any other variable returns an empty string. An empty string will also be
returned if a path could not be located.

### getenv function
`function os.getenv(name: string) string?`

`getenv` returns the value of an environment variable of the given name. Returns
nil if the variable is not defined.

### dir function
`function os.dir(path: string) []Directory`

`dir` receives a directory path and returns a list of files within the
directory. Throws an error if the path is not accessible as a directory.

## Types

### API
```
type API struct {}
```

`API` represents descriptions of the classes and enums of a Roblox API.

### Config
```
type Config: struct {
	API: API?
	Defines: Defines?
}
```

`Config` contains options for configuring the behavior of rbxmk functions.

The `API` field specifies an API, which may be used to enhance the encoding and
decoding of Roblox model and place files, or enforce the behaviors of Instance
values.

The `Defines` field contains definitions used by the Lua preprocessor.

#### Copy method
`function Config:Copy() Config`

`Copy` returns a deep copy of the Config.

#### Merge method
`function Config:Merge(other: Config)`

`Merge` merges *other* into the current Config.

### Format
`type Format: string`

Format determines how a source is encoded and decoded. The name of a format name
may also be a file extension. Consequentially, a leading `.` character on the
format is ignored (`.json` is equivalent to `json`).

Format             | Decodes to                | Encodes to
-------------------|---------------------------|-----------
`properties.json`  | `Properties`              | JSON Object.
`properties.xml`   | `Properties`              | XML.
`lua`              | `string`                  | Raw.
`script.lua`       | `Instance` (Script)       | Raw.
`localscript.lua`  | `Instance` (LocalScript)  | Raw.
`modulescript.lua` | `Instance` (ModuleScript) | Raw.
`rbxl`             | `Instance` (DataModel)    | Roblox binary place.
`rbxlx`            | `Instance` (DataModel)    | Roblox XML place.
`rbxm`             | `Instance`                | Roblox binary model.
`rbxmx`            | `Instance`                | Roblox XML model.

### SourceOption
`type SourceOption: Format | API | Config`

`SourceOption` is used by `rbxmk.read` and `rbxmk.write` to configure their
behavior.

- `Format`: Determines how data is encoded and decoded.
- `API`: Enhances how Instances are encoded and decoded.
- `Config`: Uses the `API` field.

### Instance
```type Instance struct {
	~ClassName: string
	Parent: Instance?
}
```

`Instance` provides a subset of the Roblox Instance API. Enough of the API
implemented to enable manipulation of instance trees.

If `rbxmk.config.API` is defined, then getting and setting properties on the
instance is enforced according to the API description of the instance's class.

#### Instance:\__index
`function Instance:__index(index: string) Value|Instance`

Gets the value of a property of the instance, or a named child of the instance.

When *index* matches a property present in the instance, then the Value of that
property is returned. If an API is defined, then properties are matched
according to the class description instead.

If *index* fails to match a property, then the FindFirstChild method is used to
find a child of the instance.

If *index* also fails to find a child, then an error is thrown.

#### Instance:\__newindex
`function Instance:__newindex(property: string, value: Value)`

Sets the value of a property of the Instance. If an API is defined, then an
error will be thrown if the given property is not a member of the Instance's
class, or if the value type is incorrect for that property.

#### Instance:\__tostring
`function Instance:__tostring() string`

Returns a string representation of the Instance.

#### Instance:\__eq
`function Instance:__index(Instance) bool`

Returns whether the Instance is equal to another. An Instance is only equal to
itself.

#### Instance:ClearAllChildren
`function Instance:ClearAllChildren()`

#### Instance:Clone
`function Instance:Clone() Instance`

#### Instance:Destroy
`function Instance:Destroy()`

#### Instance:FindFirstAncestor
`function Instance:FindFirstAncestor(name: string) Instance?`

#### Instance:FindFirstAncestorOfClass
`function Instance:FindFirstAncestorOfClass(className: string) Instance?`

#### Instance:FindFirstAncestorWhichIsA
`function Instance:FindFirstAncestorWhichIsA(className: string) Instance?`

#### Instance:FindFirstChild
`function Instance:FindFirstChild(name: string, recursive: bool? = false) Instance?`

#### Instance:FindFirstChildOfClass
`function Instance:FindFirstChildOfClass(className: string) Instance?`

#### Instance:FindFirstChildWhichIsA
`function Instance:FindFirstChildWhichIsA(className: string, recursive: bool? = false) Instance?`

#### Instance:GetChildren
`function Instance:GetChildren() []Instance`

#### Instance:GetDescendants
`function Instance:GetDescendants() []Instance`

#### Instance:GetFullName
`function Instance:GetFullName() string`

#### Instance:IsA
`function Instance:IsA(className: string) bool`

#### Instance:IsAncestorOf
`function Instance:IsAncestorOf(descendant: Instance) bool`

#### Instance:IsDescendantOf
`function Instance:IsDescendantOf(ancestor: Instance) bool`

### Properties
`type Properties: map[string]Value`

Contains a set of Instance properties, independent of a class.

#### Properties:SetTo
`function Properties:SetTo(instance: Instance)`

Assigns the value of each property in the set to the property in the instance.

If `rbxmk.config.API` is set, then an error will be thrown if a property in the
set is not a valid member of the instance's class, or if a value type is
incorrect for the property.

### Value
```type Value: any```

Value represents any Roblox type.

### Image
`type Image struct{}`

`Image` represents 2D image data.

### Audio
`type Audio struct{}`

`Audio` represents sound data.

### Mesh
`type Mesh struct{}`

`Mesh` represents 3D mesh data.

### Voxel
`type Voxel struct{}`

`Voxel` represents 3D voxel terrain data.
