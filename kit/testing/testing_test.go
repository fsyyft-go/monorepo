// Copyright 2025 fsyyft-go
//
// Licensed under the MIT License. See LICENSE file in the project root for full license information.

package testing

import (
	"bytes"
	"os"
	"testing"
)

func TestPrintln(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe error: %v", err)
	}
	oldStdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	Println("测试信息", 123)
	if err := w.Close(); err != nil {
		t.Fatalf("w.Close error: %v", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("ReadFrom error: %v", err)
	}
	output := buf.String()
	want := "=-=       测试信息 123\n"
	if output != want {
		t.Errorf("Println output = %q, want %q", output, want)
	}
}

func TestPrintf(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe error: %v", err)
	}
	oldStdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	Printf("当前进度：%d%%\n", 50)
	if err := w.Close(); err != nil {
		t.Fatalf("w.Close error: %v", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("ReadFrom error: %v", err)
	}
	output := buf.String()
	want := "=-=       当前进度：50%\n"
	if output != want {
		t.Errorf("Printf output = %q, want %q", output, want)
	}
}
