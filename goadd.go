// #Package by van9
package goadd

import "golang.org/x/exp/constraints"

// A number is an interface that represent any int/float
type number interface {
	constraints.Float | constraints.Integer
}

// Add adds two numbers together and return same type
//
// https://www.mathsisfun.com/numbers/addition.html
func Add[T number](x, y T) T {
	return x + y
}
