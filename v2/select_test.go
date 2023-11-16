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
		{
			name: "single predicate",
			q:    NewSelector[TestModel]().Where(C("name").Eq(18)),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `TestModel` WHERE `name` = ?",
				Args: []any{18},
			},
			wantError: nil,
		},
		{
			name: "multi predicate",
			q:    NewSelector[TestModel]().Where(C("id").Eq(123).AND(C("name").Eq("chuyingjie")), NOT(C("age").GT(35))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `TestModel` WHERE `id` = ? AND `name` = ? AND  NOT `age` > ?",
				Args: []any{123, "chuyingjie", 35},
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
