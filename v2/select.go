package orm

import (
	"context"
	"reflect"
	"strings"
)

type Selector[T any] struct {
	table string
	where []Predicate
	sb    strings.Builder
	args  []any
}

func (s *Selector[T]) Build() (*Query, error) {
	s.sb.WriteString("SELECT * FROM ")
	if s.table == "" {
		var t T
		typ := reflect.TypeOf(t).Name()
		s.sb.WriteByte('`')
		s.sb.WriteString(typ)
		s.sb.WriteByte('`')
	} else {
		s.sb.WriteString(s.table)
	}
	if len(s.where) > 0 {
		s.sb.WriteString(" WHERE ")
		p := s.where[0]
		for i := 1; i < len(s.where); i++ {
			p = p.AND(s.where[i])
		}
		s.buildExpression(p)
	}
	return &Query{
		SQL:  s.sb.String(),
		Args: s.args,
	}, nil
}

func (s *Selector[T]) buildExpression(expr Expression) {
	switch exp := expr.(type) {
	case Column:
		s.sb.WriteByte('`')
		s.sb.WriteString(exp.name)
		s.sb.WriteByte('`')
	case Value:
		s.sb.WriteByte('?')
		s.args = append(s.args, exp.value)
	case Predicate:
		_, lp := exp.left.(Predicate)
		if lp {
			//s.sb.WriteByte('(')
		}
		s.buildExpression(exp.left)
		if lp {
			//s.sb.WriteByte(')')
		}

		s.sb.WriteByte(' ')
		s.sb.WriteString(exp.op.String())
		s.sb.WriteByte(' ')

		_, lp = exp.right.(Predicate)
		if lp {
			//s.sb.WriteByte('(')
		}
		s.buildExpression(exp.right)
		if lp {
			//s.sb.WriteByte(')')
		}
	}
}

func (s *Selector[T]) From(table string) *Selector[T] {
	s.table = table
	return s
}

func (s *Selector[T]) Where(where ...Predicate) *Selector[T] {
	s.where = where
	return s
}

func (s *Selector[T]) Get(ctx context.Context) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Selector[T]) GetMulti(ctx context.Context) (*T, error) {
	//TODO implement me
	panic("implement me")
}

func NewSelector[T any]() *Selector[T] {
	return &Selector[T]{}
}
