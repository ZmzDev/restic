package main

import (
	"bytes"
	"context"
	"os"
	"testing"

	rtest "github.com/restic/restic/internal/test"
)

func testRunCheck(t testing.TB, gopts GlobalOptions) {
	t.Helper()
	output, err := testRunCheckOutput(gopts, true)
	if err != nil {
		t.Error(output)
		t.Fatalf("unexpected error: %+v", err)
	}
}

func testRunCheckMustFail(t testing.TB, gopts GlobalOptions) {
	t.Helper()
	_, err := testRunCheckOutput(gopts, false)
	rtest.Assert(t, err != nil, "expected non nil error after check of damaged repository")
}

func testRunCheckOutput(gopts GlobalOptions, checkUnused bool) (string, error) {
	buf := bytes.NewBuffer(nil)

	globalOptions.stdout = buf
	defer func() {
		globalOptions.stdout = os.Stdout
	}()

	opts := CheckOptions{
		ReadData:    true,
		CheckUnused: checkUnused,
	}

	err := runCheck(context.TODO(), opts, gopts, nil)
	return buf.String(), err
}
