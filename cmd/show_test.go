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
	}{
		{command: "sample-go-cli show", want: "show called: optint: 0, optstr: default"},
		{command: "sample-go-cli show --int 10", want: "show called: optint: 10, optstr: default"},
		{command: "sample-go-cli show --str test", want: "show called: optint: 0, optstr: test"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := NewCmdShow()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
