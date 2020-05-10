package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestShow(t *testing.T) {
	cases := []struct {
		command string
		want    string
		errWant string
	}{
		{command: "sample-go-cli show", want: "", errWant: "Parameter error: Optstr is required"},
		{command: "sample-go-cli show --int 10", want: "", errWant: "Parameter error: Optstr is required"},
		{command: "sample-go-cli show --str test", want: "show called: optint: 0, optstr: test"},
		{command: "sample-go-cli show --str \"test1 test2\"", want: "", errWant: "Parameter error: Optstr is not valid test1 test2"},
		{command: "sample-go-cli show --int 1000 --str 123", want: "", errWant: "Parameter error: Optint cannot be greater than 10"},
		{command: "sample-go-cli show --int -1 --str abc", want: "", errWant: "Parameter error: Optint must be greater than 0"},
		{command: "sample-go-cli show --int 1 --str abc", want: "show called: optint: 1, optstr: abc", errWant: ""},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdShow()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		cmd.SetArgs(cmdArgs[1:])

		if err := cmd.Execute(); err != nil {
			if c.errWant != err.Error() {
				t.Errorf("unexpected error response: errWant:%+v, get:%+v", c.errWant, err.Error())
			}
		}

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
