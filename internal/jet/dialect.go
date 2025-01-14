package jet

// Dialect interface
type Dialect interface {
	Name() string
	PackageName() string
	OperatorSerializeOverride(operator string) SerializeOverride
	FunctionSerializeOverride(function string) SerializeOverride
	AliasQuoteChar() byte
	IdentifierQuoteChar() byte
	ArgumentPlaceholder() QueryPlaceholderFunc
}

// SerializeFunc func
type SerializeFunc func(statement StatementType, out *SQLBuilder, options ...SerializeOption)

// SerializeOverride func
type SerializeOverride func(expressions ...Expression) SerializeFunc

// QueryPlaceholderFunc func
type QueryPlaceholderFunc func(ord int) string

// DialectParams struct
type DialectParams struct {
	Name                       string
	PackageName                string
	OperatorSerializeOverrides map[string]SerializeOverride
	FunctionSerializeOverrides map[string]SerializeOverride
	AliasQuoteChar             byte
	IdentifierQuoteChar        byte
	ArgumentPlaceholder        QueryPlaceholderFunc
}

// NewDialect creates new dialect with params
func NewDialect(params DialectParams) Dialect {
	return &dialectImpl{
		name:                       params.Name,
		packageName:                params.PackageName,
		operatorSerializeOverrides: params.OperatorSerializeOverrides,
		functionSerializeOverrides: params.FunctionSerializeOverrides,
		aliasQuoteChar:             params.AliasQuoteChar,
		identifierQuoteChar:        params.IdentifierQuoteChar,
		argumentPlaceholder:        params.ArgumentPlaceholder,
	}
}

type dialectImpl struct {
	name                       string
	packageName                string
	operatorSerializeOverrides map[string]SerializeOverride
	functionSerializeOverrides map[string]SerializeOverride
	aliasQuoteChar             byte
	identifierQuoteChar        byte
	argumentPlaceholder        QueryPlaceholderFunc

	supportsReturning bool
}

func (d *dialectImpl) Name() string {
	return d.name
}

func (d *dialectImpl) PackageName() string {
	return d.packageName
}

func (d *dialectImpl) OperatorSerializeOverride(operator string) SerializeOverride {
	if d.operatorSerializeOverrides == nil {
		return nil
	}
	return d.operatorSerializeOverrides[operator]
}

func (d *dialectImpl) FunctionSerializeOverride(function string) SerializeOverride {
	if d.functionSerializeOverrides == nil {
		return nil
	}
	return d.functionSerializeOverrides[function]
}

func (d *dialectImpl) AliasQuoteChar() byte {
	return d.aliasQuoteChar
}

func (d *dialectImpl) IdentifierQuoteChar() byte {
	return d.identifierQuoteChar
}

func (d *dialectImpl) ArgumentPlaceholder() QueryPlaceholderFunc {
	return d.argumentPlaceholder
}
