package main

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/macheal/tengo/v2"
)

const V_RETURN = "V_RETURN"

func main() {
	src := `
	V_INTERNAL_F := func() {
		list:={}
		list.b = "b"
		list.c = "c"
		list.a = "a"

		return list
	}
	V_RETURN := V_INTERNAL_F()
`

	res, err := run(src)
	if err != nil {
		panic(err)
	}
	j_str, err := json.Marshal(res)
	println(string(j_str), err)

}

func run(src string) (res interface{}, err error) {
	// 5 seconds context timeout is enough for an example.
	ctx := context.Background()

	script := tengo.NewScript([]byte(src))

	compiled, err := script.Compile()
	if err != nil {
		return nil, err
	}

	if err := compiled.RunContext(ctx); err != nil {
		return nil, err

	}

	if err := compiled.RunContext(context.Background()); err != nil {

		return nil, err
	}

	//logger.LogDebugf( " 4. RunContext ok ")

	v := compiled.Get(V_RETURN)
	if err := v.Error(); err != nil {
		return nil, err
	}
	if !v.IsUndefined() {
		return v.Value(), nil

	}
	return nil, errors.New("err.")
}
