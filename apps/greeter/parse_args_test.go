package main

import (
	"errors"
	"testing"
)

func TestParseArgs(t *testing.T) {
	type testConfig struct {
		args []string
		err  error
		config
	}
	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUsage: false, numTimes: 10},
		},
		{
			args:   []string{"abc"},
			err:    errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			config: config{printUsage: false, numTimes: 0},
		},
		{
			args:   []string{"1", "foo"},
			err:    errors.New("Invalid number of arguments"),
			config: config{printUsage: false, numTimes: 0},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.err != nil && err.Error() != tc.err.Error() {
			t.Fatalf("Expected error %v, got %v", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if c.printUsage != tc.config.printUsage {
			t.Fatalf("Expected %v, got %v", tc.config.printUsage, c.printUsage)
		}
		if c.numTimes != tc.config.numTimes {
			t.Fatalf("Expected %v, got %v", tc.config.numTimes, c.numTimes)
		}
	}
}

/*
* this code helps ensure that the parseArgs function works correctly in different situations. It's like a quality control check on the code before it goes out to production.
* By writing these tests, devs can be more confident that their code won't break unexpectedly when they use it.
*
 */
