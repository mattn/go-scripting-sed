package sed

import (
	"testing"

	"github.com/lestrrat/go-scripting/cmd"
)

func TestSed(t *testing.T) {
	r, err := cmd.New("cat", "testdata/example.txt").
		BailOnError(true).
		CaptureStderr(true).
		CaptureStdout(true).
		Filter(Sed("hello", "こんにちわ")).
		Do(nil)
	if err != nil {
		t.Fatal(err)
	}

	want := "こんにちわ foo\nbar こんにちわ\nherro baz\n"
	got := r.OutputString()
	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}
