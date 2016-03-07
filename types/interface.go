package types

type ValueAppender interface {
	AppendValue(b []byte, quote bool) ([]byte, error)
}

//------------------------------------------------------------------------------

// Q is a ValueAppender that represents safe SQL query.
type Q string

var _ ValueAppender = Q("")

func (q Q) AppendValue(dst []byte, quote bool) ([]byte, error) {
	return append(dst, string(q)...), nil
}

//------------------------------------------------------------------------------

// F is a ValueAppender that represents SQL field, e.g. table or column name.
type F string

var _ ValueAppender = F("")

func (f F) AppendValue(dst []byte, quote bool) ([]byte, error) {
	return AppendField(dst, string(f), quote), nil
}