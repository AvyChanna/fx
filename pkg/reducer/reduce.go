package reducer

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	fxJson "github.com/antonmedv/fx/pkg/json"
	fxTheme "github.com/antonmedv/fx/pkg/theme"
)

func Reduce(input interface{}, args []string, theme fxTheme.Theme, fxrc string) int {
	path, ok := SplitSimplePath(args)
	if ok {
		output := GetBySimplePath(input, path)
		Echo(output, theme)
		return 0
	}
	var cmd *exec.Cmd = CreateNodejs(args, fxrc)

	// TODO: Reimplement stringify with io.Reader.
	cmd.Stdin = strings.NewReader(fxJson.Stringify(input))
	output, err := cmd.CombinedOutput()
	if err != nil {
		exitCode := 1
		status, ok := err.(*exec.ExitError)
		if ok {
			exitCode = status.ExitCode()
		} else {
			fmt.Println(err.Error())
		}
		fmt.Print(string(output))
		return exitCode
	}

	dec := json.NewDecoder(bytes.NewReader(output))
	dec.UseNumber()
	object, err := fxJson.Parse(dec)
	if err != nil {
		fmt.Print(string(output))
		return 0
	}
	Echo(object, theme)
	if dec.InputOffset() < int64(len(output)) {
		fmt.Print(string(output[dec.InputOffset():]))
	}
	return 0
}
