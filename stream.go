package main

import (
	"encoding/json"
	"fmt"
	"io"

	fxJson "github.com/antonmedv/fx/pkg/json"
	fxReducer "github.com/antonmedv/fx/pkg/reducer"
	fxTheme "github.com/antonmedv/fx/pkg/theme"
)

func stream(dec *json.Decoder, object interface{}, args []string, theme fxTheme.Theme, fxrc string) int {
	var err error
	for {
		if object != nil {
			fxReducer.Reduce(object, args, theme, fxrc)
		}
		object, err = fxJson.Parse(dec)
		if err == io.EOF {
			return 0
		}
		if err != nil {
			fmt.Println("JSON Parse Error:", err.Error())
			return 1
		}
	}
}
