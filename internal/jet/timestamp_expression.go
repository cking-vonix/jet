package jet

// TimestampExpression interface
type TimestampExpression interface {
	Expression

	EQ(rhs TimestampExpression) BoolExpression
	NOT_EQ(rhs TimestampExpression) BoolExpression
	IS_DISTINCT_FROM(rhs TimestampExpression) BoolExpression
	IS_NOT_DISTINCT_FROM(rhs TimestampExpression) BoolExpression

	LT(rhs TimestampExpression) BoolExpression
	LT_EQ(rhs TimestampExpression) BoolExpression
	GT(rhs TimestampExpression) BoolExpression
	GT_EQ(rhs TimestampExpression) BoolExpression
}

type timestampInterfaceImpl struct {
	parent TimestampExpression
}

func (t *timestampInterfaceImpl) EQ(rhs TimestampExpression) BoolExpression {
	return eq(t.parent, rhs)
}

func (t *timestampInterfaceImpl) NOT_EQ(rhs TimestampExpression) BoolExpression {
	return notEq(t.parent, rhs)
}

func (t *timestampInterfaceImpl) IS_DISTINCT_FROM(rhs TimestampExpression) BoolExpression {
	return isDistinctFrom(t.parent, rhs)
}

func (t *timestampInterfaceImpl) IS_NOT_DISTINCT_FROM(rhs TimestampExpression) BoolExpression {
	return isNotDistinctFrom(t.parent, rhs)
}

func (t *timestampInterfaceImpl) LT(rhs TimestampExpression) BoolExpression {
	return lt(t.parent, rhs)
}

func (t *timestampInterfaceImpl) LT_EQ(rhs TimestampExpression) BoolExpression {
	return ltEq(t.parent, rhs)
}

func (t *timestampInterfaceImpl) GT(rhs TimestampExpression) BoolExpression {
	return gt(t.parent, rhs)
}

func (t *timestampInterfaceImpl) GT_EQ(rhs TimestampExpression) BoolExpression {
	return gtEq(t.parent, rhs)
}

//-------------------------------------------------

type timestampExpressionWrapper struct {
	timestampInterfaceImpl
	Expression
}

func newTimestampExpressionWrap(expression Expression) TimestampExpression {
	timestampExpressionWrap := timestampExpressionWrapper{Expression: expression}
	timestampExpressionWrap.timestampInterfaceImpl.parent = &timestampExpressionWrap
	return &timestampExpressionWrap
}

// TimestampExp is timestamp expression wrapper around arbitrary expression.
// Allows go compiler to see any expression as timestamp expression.
// Does not add sql cast to generated sql builder output.
func TimestampExp(expression Expression) TimestampExpression {
	return newTimestampExpressionWrap(expression)
}
