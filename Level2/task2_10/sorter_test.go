package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name            string
		config          *Config
		input           string
		want            string
		wantErr         bool
		wantErrorString string
	}{
		{
			name:   "Simple lexicographical sort",
			config: &Config{},
			input:  "c\na\nb",
			want:   "a\nb\nc\n",
		},
		{
			name:   "Reverse sort (-r)",
			config: &Config{Reverse: true},
			input:  "c\na\nb",
			want:   "c\nb\na\n",
		},
		{
			name:   "Unique lines (-u)",
			config: &Config{Unique: true},
			input:  "c\na\nc\nb\na",
			want:   "a\nb\nc\n",
		},
		{
			name:   "Sort by column 2 (-k 2)",
			config: &Config{Column: 2},
			input:  "gamma\t10\nalpha\t30\nbeta\t20",
			want:   "gamma\t10\nbeta\t20\nalpha\t30\n",
		},
		{
			name:   "Sort by column 2 numeric (-k 2 -n)",
			config: &Config{Column: 2, Numeric: true},
			input:  "line\t100\nanother\t20\nlast\t1",
			want:   "last\t1\nanother\t20\nline\t100\n",
		},
		{
			name:   "Sort by column 2 numeric reverse (-k 2 -n -r)",
			config: &Config{Column: 2, Numeric: true, Reverse: true},
			input:  "line\t100\nanother\t20\nlast\t1",
			want:   "line\t100\nanother\t20\nlast\t1\n",
		},
		{
			name:   "Column out of bounds",
			config: &Config{Column: 3},
			input:  "a\t1\nb\t2",
			want:   "a\t1\nb\t2\n",
		},
		{
			name:   "Numeric sort (-n)",
			config: &Config{Numeric: true},
			input:  "100\n20\n1\n-5",
			want:   "-5\n1\n20\n100\n",
		},
		{
			name:   "Human-readable sort (-h)",
			config: &Config{HumanNumeric: true},
			input:  "1G\n5M\n10K",
			want:   "10K\n5M\n1G\n",
		},
		{
			name:   "Human-readable reverse sort (-h -r)",
			config: &Config{HumanNumeric: true, Reverse: true},
			input:  "1G\n5M\n10K",
			want:   "1G\n5M\n10K\n",
		},
		{
			name:   "Month sort (-M)",
			config: &Config{Month: true},
			input:  "Mar\nJan\nFeb",
			want:   "Jan\nFeb\nMar\n",
		},
		{
			name:   "Month sort case-insensitive full names (-M)",
			config: &Config{Month: true},
			input:  "March\njanuary\nFEBRUARY",
			want:   "january\nFEBRUARY\nMarch\n",
		},
		{
			name:   "Ignore trailing blanks (-b)",
			config: &Config{IgnoreBlanks: true},
			input:  "b\na  ",
			want:   "a  \nb\n",
		},
		{
			name:    "Check sorted (correct)",
			config:  &Config{CheckSorted: true},
			input:   "a\nb\nc",
			want:    "",
			wantErr: false,
		},
		{
			name:            "Check sorted (disorder)",
			config:          &Config{CheckSorted: true},
			input:           "c\na\nb",
			want:            "",
			wantErr:         true,
			wantErrorString: "disorder: a",
		},
		{
			name:            "Check sorted numeric (disorder)",
			config:          &Config{CheckSorted: true, Numeric: true},
			input:           "1\n10\n2",
			want:            "",
			wantErr:         true,
			wantErrorString: "disorder: 2",
		},
		{
			name:   "Unique flag with non-identical lines having identical keys",
			config: &Config{Unique: true, Column: 2},
			input:  "b\t10\na\t20\nc\t10",
			want:   "b\t10\nc\t10\na\t20\n",
		},
		{
			name:   "Unique flag with truly identical lines",
			config: &Config{Unique: true, Column: 2},
			input:  "b\t10\na\t20\nb\t10",
			want:   "b\t10\na\t20\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputReader := strings.NewReader(tt.input)
			var outputBuf bytes.Buffer

			err := run(inputReader, &outputBuf, "stdin", tt.config)

			if tt.wantErr {
				if err == nil {
					t.Errorf("run() expected an error, but got nil")
					return
				}
				if tt.wantErrorString != "" && !strings.Contains(err.Error(), tt.wantErrorString) {
					t.Errorf("run() error = %q, want error to contain %q", err.Error(), tt.wantErrorString)
				}
			} else {
				if err != nil {
					t.Errorf("run() returned unexpected error: %v", err)
				}
			}

			got := outputBuf.String()
			if got != tt.want {
				t.Errorf("run() output = %q, want %q", got, tt.want)
			}
		})
	}
}
