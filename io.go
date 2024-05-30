package main

import (
	"bytes"
	"io"
	"strings"
)

// Wrapping a writer
type MyWriter struct {
	w io.Writer
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	return w.w.Write(p)
}

// Wrapping a reader
type MyReader struct {
	r io.Reader
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	return r.r.Read(p)
}

// Never ending reader
type NeverEnding byte

func (n NeverEnding) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = byte(n)
	}
	return len(p), nil
}

func main() {
	buf := make([]byte, 10)

	var r io.Reader
	r = NeverEnding('A')
	r = io.LimitReader(r, 10)
	r = strings.NewReader("Hello, World!")
	pr, pw := io.Pipe()

	_, err = io.ReadAll(r)

	if err != nil {
		panic(err)
	}

	r = io.MultiReader(r, r)
	r = io.TeeReader(r, &w)

	var w bytes.Buffer
	w = io.MultiWriter(&w, &w)

	_, err = io.writeString(w, "Hello, World!")
	_, err = io.Copy(w, r)
}
