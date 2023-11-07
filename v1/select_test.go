package orm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelector_Build(t *testing.T) {
	test_cases := []struct {
		name      string
		q         QueryBuilder
		wantQuery *Query
		wantError error
	}{
		{
			name: "default from",
			q:    NewSelector[TestModel](),
			wantQuery: &Query{
				SQL: "SELECT * FROM `TestModel`",
			},
			wantError: nil,
		},
		{
			name: "with from",
			q:    NewSelector[TestModel]().From("`TestTable`"),
			wantQuery: &Query{
				SQL: "SELECT * FROM `TestTable`",
			},
			wantError: nil,
		},
		{
			name: "with db",
			q:    NewSelector[TestModel]().From("`TestDb`.`TestTable`"),
			wantQuery: &Query{
				SQL: "SELECT * FROM `TestDb`.`TestTable`",
			},
			wantError: nil,
		},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			query, err := tc.q.Build()
			assert.Equal(t, tc.wantError, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantQuery, query)
		})
	}
}
