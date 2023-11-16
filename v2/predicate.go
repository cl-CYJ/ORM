package orm

type Expression interface {
	expr()
}

type Predicate struct {
	left  Expression
	op    op
	right Expression //
}

func (p Predicate) expr() {

}

func (p Predicate) AND(r Predicate) Predicate {
	return Predicate{
		left:  p,
		op:    opAND,
		right: r,
	}
}

func (p Predicate) OR(r Predicate) Predicate {
	return Predicate{
		left:  p,
		op:    opOR,
		right: r,
	}
}

func NOT(p Predicate) Predicate {
	return Predicate{
		op:    opNOT,
		right: p,
	}
}
