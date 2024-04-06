package ref

import "reflect"

func CallMethod(method reflect.Method, reciever any, args ...any) []any {
	return CallFunc(method.Func, append([]any{reciever}, args...))
}

func CallFunc(f reflect.Value, args []any) []any {
	inputs := make([]reflect.Value, len(args))
	for i, arg := range args {
		inputs[i] = reflect.ValueOf(arg)
	}
	outputs := f.Call(inputs)

	results := make([]any, len(outputs))
	for i, output := range outputs {
		results[i] = output.Interface()
	}

	return results
}
