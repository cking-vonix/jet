package jet

import (
	"gotest.tools/assert"
	"testing"
)

func TestClauseSelect_Serialize(t *testing.T) {
	defer func() {
		r := recover()
		assert.Equal(t, r, "jet: SELECT clause has to have at least one projection")
	}()

	selectClause := &ClauseSelect{}
	selectClause.Serialize(SelectStatementType, &SQLBuilder{})
}
