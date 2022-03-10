package stmt

import (
	"fmt"
	"github.com/viant/velty/est"
	"github.com/viant/velty/est/op"
	"reflect"
	"unsafe"
)

type directAppender struct {
	x *op.Operand
}

func (a *directAppender) appendString(state *est.State) unsafe.Pointer {
	ptr := state.Pointer(*a.x.Offset)
	val := *(*string)(ptr)
	state.Buffer.AppendString(val)
	return ptr
}

func (a *directAppender) appendInt(state *est.State) unsafe.Pointer {
	ptr := state.Pointer(*a.x.Offset)
	val := *(*int)(ptr)
	state.Buffer.AppendInt(val)
	return ptr

}

func (a *directAppender) appendBool(state *est.State) unsafe.Pointer {
	ptr := state.Pointer(*a.x.Offset)
	val := *(*bool)(ptr)
	state.Buffer.AppendBool(val)
	return ptr
}

func Selector(expr *op.Expression) est.New {
	return func(control est.Control) (est.Compute, error) {
		x, err := expr.Operand(control)
		if err != nil {
			return nil, err
		}
		//TODO check if you can use direct appnder
		result := &directAppender{x: x}
		switch expr.Type.Kind() {
		case reflect.Int:
			return result.appendInt, nil
		case reflect.String:
			return result.appendInt, nil
		case reflect.Bool:
			return result.appendBool, nil
		default:
			return nil, fmt.Errorf("unsupported append selector: %s", expr.Type.String())
		}
	}
}
