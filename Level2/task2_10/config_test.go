package main

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func Test_parseFlags(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantConfig *Config
	}{
		{
			name: "Default values",
			args: []string{},
			wantConfig: &Config{
				Column: 0, Numeric: false, Reverse: false, Unique: false,
				Month: false, IgnoreBlanks: false, CheckSorted: false, HumanNumeric: false,
			},
		},
		{
			name: "Single flags: -k 3, -n, -r",
			args: []string{"-k", "3", "-n", "-r"},
			wantConfig: &Config{
				Column: 3, Numeric: true, Reverse: true, Unique: false,
				Month: false, IgnoreBlanks: false, CheckSorted: false, HumanNumeric: false,
			},
		},
		{
			name: "All boolean flags",
			args: []string{"-n", "-r", "-u", "-M", "-b", "-c", "-h"},
			wantConfig: &Config{
				Column: 0, Numeric: true, Reverse: true, Unique: true,
				Month: true, IgnoreBlanks: true, CheckSorted: true, HumanNumeric: true,
			},
		},
		{
			name: "Column flag only",
			args: []string{"-k", "5"},
			wantConfig: &Config{
				Column: 5, Numeric: false, Reverse: false, Unique: false,
				Month: false, IgnoreBlanks: false, CheckSorted: false, HumanNumeric: false,
			},
		},
		{
			name: "Human numeric sort",
			args: []string{"-h"},
			wantConfig: &Config{
				Column: 0, Numeric: false, Reverse: false, Unique: false,
				Month: false, IgnoreBlanks: false, CheckSorted: false, HumanNumeric: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }()

			flag.CommandLine = flag.NewFlagSet(tt.name, flag.ExitOnError)
			os.Args = append([]string{"my_sort"}, tt.args...)

			gotConfig := parseFlags()

			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("parseFlags() = %+v, want %+v", gotConfig, tt.wantConfig)
			}
		})
	}
}
