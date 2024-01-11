package prolog

import "github.com/ichiban/prolog/engine"

var (
	// AtomTypeAtom is the term used to represent the atom type.
	AtomTypeAtom = engine.NewAtom("atom")
	// AtomTypeByte is the term used to represent the byte type.
	AtomTypeByte = engine.NewAtom("byte")
	// AtomTypeCharacter is the term used to represent the character type.
	// A character type is a single character (e.g. 'a') identified in the Unicode standard.
	AtomTypeCharacter = engine.NewAtom("character")
	// AtomTypeCharacterCode is the term used to represent the character code type.
	// A character code type is a single character identified by its code point (a number) in the Unicode standard.
	AtomTypeCharacterCode = engine.NewAtom("character_code")
	// AtomTypeCharset is the term used to represent the charset type.
	// A charset type is a set of characters identified by its name in the IANA standard.
	AtomTypeCharset = engine.NewAtom("charset")
	// AtomTypeCryptographicAlgorithm is the term used to represent the cryptographic algorithm type.
	AtomTypeCryptographicAlgorithm = engine.NewAtom("cryptographic_algorithm")
	// AtomTypeDID is the term used to represent the DID type.
	// DID type is a compound with the name "did" and 5 arguments which are the components of the DID, in the form of
	// did(Method, ID, Path, Query, Fragment).
	AtomTypeDID = engine.NewAtom("did")
	// AtomTypeHashAlgorithm is the term used to represent the hash algorithm type.
	AtomTypeHashAlgorithm = engine.NewAtom("hash_algorithm")
	// AtomTypeIOMode is the term used to represent the IO mode type.
	// An IO mode specifies the direction of the IO operation represented as an atom.
	// Possible values are: read, write, append.
	AtomTypeIOMode = engine.NewAtom("io_mode")
	// AtomTypeStream is the term used to represent the stream type.
	AtomTypeStream = engine.NewAtom("stream")
	// AtomTypeText is the term used to represent the text type.
	// A text type is either an atom, a list of characters or a list of character codes.
	AtomTypeText = AtomText
	// AtomTypeList is the term used to represent the list type.
	AtomTypeList = engine.NewAtom("list")
	// AtomTypeNumber is the term used to represent the number type.
	AtomTypeNumber = engine.NewAtom("number")
	// AtomTypeOption is the term used to represent the option type.
	// An option is a compound with the name of the option as functor and one term argument which is
	// the value of the option. For instance: opt(v).
	AtomTypeOption = engine.NewAtom("option")
	// AtomTypePair is the term used to indicate the pair type.
	AtomTypePair = engine.NewAtom("pair")
	// AtomTypeJSON is the term used to indicate the json type.
	AtomTypeJSON = AtomJSON
	// AtomTypeURIComponent is the term used to represent the URI component type.
	AtomTypeURIComponent = engine.NewAtom("uri_component")
)

var (
	// AtomValidEncoding is the atom denoting a valid encoding.
	// The valid encoding atom is a compound with the name of the encoding which is a valid encoding with
	// regard to the predicate where it is used.
	//
	// For instance: valid_encoding(utf8), valid_encoding(hex).
	AtomValidEncoding = engine.NewAtom("encoding")
	// AtomValidEmptyList is the atom denoting a valid empty list.
	AtomValidEmptyList = engine.NewAtom("empty_list")
)

func ValidEncoding(encoding string) engine.Term {
	return AtomValidEncoding.Apply(engine.NewAtom(encoding))
}

func ValidEmptyList() engine.Term {
	return AtomValidEmptyList
}

// ResourceError creates a new resource error exception.
// TODO: to remove once engine.resourceError() is public.
func ResourceError(resource engine.Term, env *engine.Env) engine.Exception {
	return engine.NewException(
		AtomError.Apply(
			engine.NewAtom("resource_error").Apply(resource),
			engine.NewAtom("unknown")), env)
}

var (
	// AtomResourceContext is the atom denoting the "context" resource.
	// The context resource is a contextual data that contains all information needed to
	// process a request and produce a response with the blockchain.
	AtomResourceContext = engine.NewAtom("resource_context")
	// AtomResourceModule is the atom denoting the "module" resource.
	// The module resource is the representation of the module with which the interaction is made.
	// The module resource is denoted as a compound with the name of the module.
	AtomResourceModule = engine.NewAtom("resource_module")
)

func ResourceContext() engine.Term {
	return AtomResourceContext
}

func ResourceModule(module string) engine.Term {
	return AtomResourceModule.Apply(engine.NewAtom(module))
}

// PermissionError creates a new permission error exception.
// TODO: to remove once engine.permissionError() is public.
func PermissionError(operation, permissionType, culprit engine.Term, env *engine.Env) engine.Exception {
	return engine.NewException(
		AtomError.Apply(
			engine.NewAtom("permission_error").Apply(
				operation, permissionType, culprit,
			),
			engine.NewAtom("unknown")), env)
}

var AtomOperationInput = engine.NewAtom("input")

var AtomPermissionTypeStream = engine.NewAtom("stream")

// ExistenceError creates a new existence error exception.
// TODO: to remove once engine.existenceError() is public.
func ExistenceError(objectType, culprit engine.Term, env *engine.Env) engine.Exception {
	return engine.NewException(
		AtomError.Apply(
			engine.NewAtom("existence_error").Apply(
				objectType, culprit,
			),
			engine.NewAtom("unknown")), env)
}

var AtomObjectTypeSourceSink = engine.NewAtom("source_sink")

// UnexpectedError creates a new unexpected error exception.
// TODO: to remove once engine.syntaxError() is public.
func SyntaxError(err error, env *engine.Env) engine.Exception {
	return engine.NewException(
		AtomError.Apply(
			engine.NewAtom("syntax_error").Apply(StringToCharacterListTerm(err.Error())),
			engine.NewAtom("unknown")), env)
}

// WithError adds the error term to the exception term if possible.
// TODO: wait for ichiban/prolog to offer a better way to do this.
func WithError(exception engine.Exception, err error, env *engine.Env) engine.Exception {
	if term, ok := exception.Term().(engine.Compound); ok {
		if term.Functor() == AtomError && term.Arity() == 2 {
			return engine.NewException(term.Functor().Apply(
				term.Arg(0),
				StringToCharacterListTerm(err.Error()),
				term.Arg(1)), env)
		}
	}
	return exception
}