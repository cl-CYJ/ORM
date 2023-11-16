package orm

type op string

const (
	opEQ  = "="
	opLT  = "<"
	opGT  = ">"
	opAND = "AND"
	opOR  = "OR"
	opNOT = "NOT"
)

func (o op) String() string {
	return string(o)
}

type Column struct {
	name string
}

func (c Column) expr() {

}

type Value struct {
	value any
}

func (c Value) expr() {

}

func exprOf(arg any) Expression {
	switch exp := arg.(type) {
	case Expression:
		return exp
	default:
		return valueOf(arg)
	}
}

func valueOf(arg any) Value {
	return Value{
		value: arg,
	}
}

func C(name string) Column {
	return Column{name}
}

func (c Column) Eq(args any) Predicate {
	return Predicate{
		left:  c,
		op:    opEQ,
		right: exprOf(args),
	}
}

func (c Column) LT(args any) Predicate {
	return Predicate{
		left:  c,
		op:    opEQ,
		right: exprOf(args),
	}
}

func (c Column) GT(args any) Predicate {
	return Predicate{
		left:  c,
		op:    opGT,
		right: exprOf(args),
	}
}
