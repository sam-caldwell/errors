package errors

const (
	// ArraySizeError - indicates that the given array size does not match the expected size.
	ArraySizeError = "array size error"

	// BoundsCheckError - indicates that a value fails to meet the expected bounds
	BoundsCheckError = "bounds check error"

	// CannotOpenFile - indicates that the file cannot be opened
	CannotOpenFile = "cannot open file"

	// CannotReadFile - indicates that the file cannot be read
	CannotReadFile = "cannot read file"

	// Details - Modifier (any) to add Details to an error
	Details = "(%v)"

	// DuplicateEntry - indicates that a duplicate is found in some set/list/object
	DuplicateEntry = "duplicate entry"

	// EmptySet - indicates that a set/list is empty
	EmptySet = "empty set"

	// FailedToOpenLogFile - indicates that the log file could not be opened
	FailedToOpenLogFile = "failed to open log file"

	// IndexOutOfRange - indicates that the index is out of range (e.g. for a list)
	IndexOutOfRange = "index out of range"

	// InternalError - indicates some unexpected/undefined internal error has occurred
	InternalError = "internal error"

	// InvalidApplicationName - indicates that the application name is invalid
	InvalidApplicationName = "invalid application name"

	// InvalidCommand - indicates that the command input is invalid
	InvalidCommand = "unknown command"

	// InvalidContextId - indicate that the contextId is invalid
	InvalidContextId = "invalid contextId"

	// InvalidInput - indicates that a given input is invalid
	InvalidInput = "invalid input"

	// InvalidLogLevel - indicates that the log level is invalid
	InvalidLogLevel = "invalid log level"

	// InvalidFilePath - invalid file/path
	InvalidFilePath = "invalid file/path"

	// LockCheckFailed - indicates that the lock-check operation failed
	LockCheckFailed = "lock check failed"

	// MalformedGuid - indicates that a given input is a malformed GUID
	MalformedGuid = "malformed GUID"

	// MalformedInput - indicates that a given input is malformed
	MalformedInput = "malformed input"

	// MissingArguments - indicates that expected arguments are missing
	MissingArguments = "missing argument"

	// MissingColor - indicates that a color value is missing
	MissingColor = "Missing color"

	// MissingContextId - indicates that context id is missing
	MissingContextId = "missing contextId (uuid)"

	// MissingExtension - indicates a file extension is missing
	MissingExtension = "missing extension"

	// MissingField - indicates that a field is missing
	MissingField = "missing field"

	// NilOrEmptyString - indicates that a given string is unexpectedly nil or empty
	NilOrEmptyString = "empty or nil string not allowed"

	// NotEnoughSymbols - indicates that too few symbols are given
	NotEnoughSymbols = "too few symbols"

	// NotFound - indicates that some entity is not found
	NotFound = "not found"

	// NotImplemented - indicates that some functionality is not yet implemented
	NotImplemented = "not implemented"

	// NotInitialized - indicates that some entity is not initialized
	NotInitialized = "not initialized"

	// OverflowError - indicates that an overflow has occurred
	OverflowError = "overflow error"

	// ReadOnly - indicates that an entity is in an unexpected read-only state
	ReadOnly = "read only"

	// TimeConversionError - indicates an error converting time
	TimeConversionError = "error converting time"

	// TypeMismatch - indicates that the given entity is of the wrong type
	TypeMismatch = "type mismatch"

	// UnableToDetectFamily - indicates that operating system family cannot be determined
	UnableToDetectFamily = "unable to detect operating system family"

	// UninitializedValue - indicates that the value has not been initialized as expected
	UninitializedValue = "uninitialized value"

	// UnknownCommand - indicates that the command is unknown
	UnknownCommand = "Unknown command"

	//UnrecoverableFilenameCollision - A filename collision has occurred for which there is no recovery option
	UnrecoverableFilenameCollision = "unrecoverable filename collision"

	// UnsupportedType - indicates an unsupported type is detected
	UnsupportedType = "Unsupported type"

	// UnsupportedOpsys - indicates that the operating system is unsupported
	UnsupportedOpsys = "unsupported operating system"

	// UnsupportedCpuArchitecture - indicates that the CPU Architecture is unsupported
	UnsupportedCpuArchitecture = "unsupported cpu architecture"

	// UnsupportedLanguage - indicates that the given language is unsupported
	UnsupportedLanguage = "unsupported language"

	// UnsupportedVersion - indicates that the version is not supported
	UnsupportedVersion = "unsupported version"

	// ValueTooLarge - indicates that the given value is too large
	ValueTooLarge = "value too large"

	// ValueTooSmall - indicates that the given value is too small
	ValueTooSmall = "value too small"
)
