package main

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_readLines(t *testing.T) {
	tests := []struct {
		name      string
		reader    io.Reader
		wantLines []string
		wantErr   bool
	}{
		{
			name:      "Normal case with multiple lines",
			reader:    strings.NewReader("line 1\nline 2\nline 3"),
			wantLines: []string{"line 1", "line 2", "line 3"},
			wantErr:   false,
		},
		{
			name:      "Empty input",
			reader:    strings.NewReader(""),
			wantLines: nil, // bufio.Scanner с пустым вводом вернет nil срез, не пустой
			wantErr:   false,
		},
		{
			name:      "Single line without newline",
			reader:    strings.NewReader("single line"),
			wantLines: []string{"single line"},
			wantErr:   false,
		},
		{
			name:      "Input with blank lines",
			reader:    strings.NewReader("alpha\n\nbeta"),
			wantLines: []string{"alpha", "", "beta"},
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLines, err := readLines(tt.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("readLines() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}

func Test_writeLines(t *testing.T) {
	tests := []struct {
		name    string
		lines   []string
		want    string
		wantErr bool
	}{
		{
			name:    "Normal case with multiple lines",
			lines:   []string{"line 1", "line 2", "line 3"},
			want:    "line 1\nline 2\nline 3\n",
			wantErr: false,
		},
		{
			name:    "Empty slice",
			lines:   []string{},
			want:    "",
			wantErr: false,
		},
		{
			name:    "Slice with empty string",
			lines:   []string{"hello", "", "world"},
			want:    "hello\n\nworld\n",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Используем bytes.Buffer для имитации io.Writer в памяти
			var buf bytes.Buffer
			err := writeLines(&buf, tt.lines)

			if (err != nil) != tt.wantErr {
				t.Errorf("writeLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got := buf.String(); got != tt.want {
				t.Errorf("writeLines() output = %q, want %q", got, tt.want)
			}
		})
	}
}

func Test_getUnique(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []string
	}{
		{
			name:  "Normal case with duplicates",
			lines: []string{"a", "a", "b", "c", "c", "c", "d"},
			want:  []string{"a", "b", "c", "d"},
		},
		{
			name:  "No duplicates",
			lines: []string{"a", "b", "c", "d"},
			want:  []string{"a", "b", "c", "d"},
		},
		{
			name:  "All elements are the same",
			lines: []string{"x", "x", "x", "x"},
			want:  []string{"x"},
		},
		{
			name:  "Empty slice",
			lines: []string{},
			want:  []string{},
		},
		{
			name:  "Single element slice",
			lines: []string{"lonely"},
			want:  []string{"lonely"},
		},
		{
			name:  "Duplicates at the beginning and end",
			lines: []string{"start", "start", "middle", "end", "end"},
			want:  []string{"start", "middle", "end"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]string, len(tt.lines))
			copy(inputCopy, tt.lines)

			got := getUnique(inputCopy)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
