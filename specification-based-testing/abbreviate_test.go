package main

import (
	"testing"
)

func TestAbbreviate(t *testing.T) {
	tests := map[string]struct {
		str, abbrevMarker string
		offset, maxWidth  int
		want              string
		wantErr           string
	}{
		// 1st iteration over partitions
		"when both str and marker are empty, return empty": {
			"", "",
			0, 0,
			"", "",
		},
		"when str is empty, return empty": {
			"", "...",
			0, 0,
			"", "",
		},
		"when offset > len(str), consider offset as str's last char": {
			"abcdefghi", ".",
			10, 3,
			".hi", "",
		},
		"when not possible abbrev str left side due a low offset, abbrev right only": {
			"abcdefghij", "...",
			4, 7,
			"abcd...", "",
		},
		"when offset until str end fits the defined max width": {
			"abcdefghijklmnop", "...",
			9, 10,
			"...jklmnop", "",
		},
		"when abbreviating left side isn't enough to keep offset visible and reach str end, abbreviate both sides": {
			"abcdefghijklmnop", "...",
			8, 10,
			"...ijkl...", "",
		},

		// 2nd iteration over partitions + coverage analysis
		"when marker is empty and len(str) equals max width, return full str": {
			"abcde", "",
			0, 5,
			"abcde", "",
		},
		"when marker is empty and len(str) less than max width, return full str": {
			"abcde", "",
			0, 6,
			"abcde", "",
		},
		"when marker is empty and str doesn't fits max width, abbreviate right side with no marker": {
			"abcde", "",
			0, 4,
			"abcd", "",
		},
		"when max width can't fit the marker with at least one character, returns error": {
			"abcde", "...",
			0, 3,
			"", "minimum abbreviation width is 4",
		},
		"when str is less than max width, return full str": {
			"abcde", "...",
			0, 6,
			"abcde", "",
		},
		"when len(str) is same as max width, return full str": {
			"abcde", "...",
			0, 5,
			"abcde", "",
		},
		"when str must be abbreviated both sides but max length can't fit a both side abbreviation with at least one character, returns error": {
			"abcdefghi", "..",
			4, 4,
			"", "minimum abbreviation width with offset is 5",
		},
		// i think it should return "..defg". Removing the weird if condition on impl line 56 leads to this return
		"when offset is near the end, adjusts to be possible to show at least str end with marker": {
			"abcdefg", "..",
			4, 6,
			"abcd..", "",
		},

		// 3rd iteration with AI analysis
		"when marker is empty and offset > 0, abbreviate right side ignoring offset": {
			"abcdefghij", "",
			5, 4,
			"abcd", "",
		},
		// good one, but same as the line 40's and the suggestion was failing (i needed to fix it)
		"when offset near the end but not enough to fit until str end, abbreviates both sides": {
			"abcdefghijklmnop", "...",
			8, 10,
			"...ijkl...", "",
		},
		"when maxWidth equals marker length plus one, abbreviates with minimum possible size": {
			"abcdefghij", "+++++",
			0, 6,
			"a+++++", "",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Abbreviate(tt.str, tt.abbrevMarker, tt.offset, tt.maxWidth)
			if tt.wantErr == "" && err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if tt.wantErr != "" {
				if err == nil {
					t.Errorf("expected error %q but got nil", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr {
					t.Errorf("error = %q, wantErr = %q", err.Error(), tt.wantErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("got = %q, want = %q", got, tt.want)
			}
		})
	}
}

func TestAbbreviate_Exploratory(t *testing.T) {
	t.Skip()
	tests := []struct {
		name,
		str, abbrevMarker string
		offset, maxWidth int
		want             string
		wantErr          string
	}{
		{
			"abbreviate end",
			"abcdefghijklmno", "...", 0, 10,
			"abcdefg...", "",
		},
		{
			"abbreviate end",
			"abcdefghijklmno", "...", 4, 10,
			"abcdefg...", "",
		},
		{
			"abbreviate middle",
			"abcdefghijklmno", "...", 5, 10,
			"...fghi...", "",
		},
		{
			"still abbreviate middle",
			"abcdefghijklm", "...", 5, 10,
			"...fghi...", "",
		},
		{
			"abbreviate start",
			"abcdefghijkl", "...", 5, 10,
			"...fghijkl", "",
		},
		{
			"abbreviate start",
			"abcdefghijklmno", "...", 10, 10,
			"...ijklmno", "",
		},
		{
			"abbreviate start",
			"abcde", ".", 3, 3,
			".de", "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Abbreviate(tt.str, tt.abbrevMarker, tt.offset, tt.maxWidth)
			if tt.wantErr == "" && err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if tt.wantErr != "" {
				if err == nil {
					t.Errorf("expected error %q but got nil", tt.wantErr)
					return
				}
				if err.Error() != tt.wantErr {
					t.Errorf("error = %q, wantErr = %q", err.Error(), tt.wantErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf("got = %q, want = %q", got, tt.want)
			}
		})
	}
}
