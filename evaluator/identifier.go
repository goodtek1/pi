package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := object.Builtins[node.Value]; ok {
		return builtin
	}
	return object.NewError("identifier not found: " + node.Value)
}
