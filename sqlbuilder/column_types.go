package sqlbuilder

//------------------------------------------------------//
type BoolColumn struct {
	boolInterfaceImpl

	baseColumn
}

func NewBoolColumn(name string, nullable NullableColumn) *BoolColumn {
	if !validIdentifierName(name) {
		panic("Invalid column name in bool column")
	}
	boolColumn := &BoolColumn{}
	boolColumn.baseColumn = newBaseColumn(name, nullable, "", boolColumn)

	boolColumn.boolInterfaceImpl.parent = boolColumn

	return boolColumn
}

//------------------------------------------------------//
type NumericColumn struct {
	numericInterfaceImpl
	baseColumn
}

func NewNumericColumn(name string, nullable NullableColumn) *NumericColumn {
	if !validIdentifierName(name) {
		panic("Invalid column name")
	}

	numericColumn := &NumericColumn{}

	numericColumn.numericInterfaceImpl.parent = numericColumn

	numericColumn.baseColumn = newBaseColumn(name, nullable, "", numericColumn)

	return numericColumn
}

//------------------------------------------------------//
type IntegerColumn struct {
	numericInterfaceImpl
	integerInterfaceImpl

	baseColumn
}

// Representation of any integer column
// This function will panic if name is not valid
func NewIntegerColumn(name string, nullable NullableColumn) *IntegerColumn {
	integerColumn := &IntegerColumn{}

	integerColumn.numericInterfaceImpl.parent = integerColumn
	integerColumn.integerInterfaceImpl.parent = integerColumn

	integerColumn.baseColumn = newBaseColumn(name, nullable, "", integerColumn)

	return integerColumn
}

//------------------------------------------------------//
type StringColumn struct {
	stringInterfaceImpl

	baseColumn
}

// Representation of any integer column
// This function will panic if name is not valid
func NewStringColumn(name string, nullable NullableColumn) *StringColumn {
	if !validIdentifierName(name) {
		panic("Invalid column name")
	}

	stringColumn := &StringColumn{}

	stringColumn.stringInterfaceImpl.parent = stringColumn
	stringColumn.stringInterfaceImpl.parent = stringColumn

	stringColumn.baseColumn = newBaseColumn(name, nullable, "", stringColumn)

	return stringColumn
}