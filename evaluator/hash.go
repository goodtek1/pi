package evaluator

import (
	"pilang/ast"
	"pilang/object"
)

func evalHash(node *ast.Hash, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)
	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return object.NewError("unusable as hash key: %s", key.Type())
		}
		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}
	return &object.Hash{Pairs: pairs}
}
