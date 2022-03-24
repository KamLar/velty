package expr

import (
	"github.com/viant/velty/ast"
	"reflect"
)

//Unary represents unary expression
type Unary struct {
	Token ast.Token
	X     ast.Expression
}

func (u *Unary) Type() reflect.Type {
	return u.X.Type()
}
