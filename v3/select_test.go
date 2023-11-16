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
				SQL: "SELECT * FROM `test_model`",
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
			q:    NewSelector[TestModel]().From("`TestDB`.`TestTable`"),
			wantQuery: &Query{
				SQL: "SELECT * FROM `TestDB`.`TestTable`",
			},
			wantError: nil,
		},
		{
			name: "single predicate",
			q:    NewSelector[TestModel]().Where(C("Name").Eq(18)),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE `name` = ?",
				Args: []any{18},
			},
			wantError: nil,
		},
		{
			name: "multi predicate",
			q:    NewSelector[TestModel]().Where(C("Id").Eq(123).AND(C("Name").Eq("chuyingjie")), NOT(C("Age").GT(35))),
			wantQuery: &Query{
				SQL:  "SELECT * FROM `test_model` WHERE `id` = ? AND `name` = ? AND  NOT `age` > ?",
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
