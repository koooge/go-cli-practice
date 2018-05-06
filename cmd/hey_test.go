package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestHey(t *testing.T) {
	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: "Hey " + os.Getenv("USER") + "!\n"},
		{args: []string{"hoge"}, want: "Hey hoge!\n"},
		{args: []string{"-n", "fuga"}, want: "Hey fuga!\n"},
		{args: []string{"--name", "piyo"}, want: "Hey piyo!\n"},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdHey()
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != get {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
