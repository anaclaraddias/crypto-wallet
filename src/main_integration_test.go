package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainIntegration_Case1(t *testing.T) {
	inputJSON := `{"type": "DEPOSIT", "asset": "BTC", "amount": 1.5}`
	input := inputJSON + "\n"

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	rIn, wIn, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdin pipe: %v", err)
	}

	os.Stdin = rIn
	go func() {
		_, _ = wIn.Write([]byte(input))
		wIn.Close()
	}()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	os.Stdout = w

	main()

	w.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read captured output: %v", err)
	}
	os.Stdout = oldStdout

	output := buf.String()

	expected := `{"BTC":1.5}`

	if !strings.Contains(output, expected) {
		t.Fatalf("expected output to contain %s, got:\n%s", expected, output)
	}
}

func TestStop(t *testing.T) {
	inputJSON := `stop`
	input := inputJSON + "\n"

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	rIn, wIn, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdin pipe: %v", err)
	}

	os.Stdin = rIn
	go func() {
		_, _ = wIn.Write([]byte(input))
		wIn.Close()
	}()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	os.Stdout = w

	main()

	w.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read captured output: %v", err)
	}
	os.Stdout = oldStdout

	output := buf.String()

	expected := `>`

	if !strings.Contains(output, expected) {
		t.Fatalf("expected output to contain %s, got:\n%s", expected, output)
	}
}

func TestMainIntegration_InvalidJson(t *testing.T) {
	inputJSON := `{"INVALID}`
	input := inputJSON + "\n"

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	rIn, wIn, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create stdin pipe: %v", err)
	}

	os.Stdin = rIn
	go func() {
		_, _ = wIn.Write([]byte(input))
		wIn.Close()
	}()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	os.Stdout = w

	main()

	w.Close()
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read captured output: %v", err)
	}
	os.Stdout = oldStdout

	output := buf.String()

	expected := `{}`

	if !strings.Contains(output, expected) {
		t.Fatalf("expected output to contain %s, got:\n%s", expected, output)
	}
}
