package main

import "testing"

func Test_replace(t *testing.T) {
	type args struct {
		re       string
		repl     string
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replace(tt.args.re, tt.args.repl, tt.args.filename)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
