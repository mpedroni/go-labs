package main

import "testing"

/* Testing the "HasPrefix" function using the spec-based testing approach

`func HasPrefix(s, prefix string) bool`

The goal of this function is to inform if a string is or isn't prefixed by a given substring.
It receive two parameters: "s" is the string the func will search by a "prefix", which is the second parameter.
The program returns true if the string begins with the prefix, otherwise false.

when "s" starts with prefix, returns true. Otherwise, false.
when "s" is empty, only return true if "prefix" is empty.
when "s" and "prefix" are equal, return true.

## specification-based testing

partitions with a "*" won`t be tested multiple times

- inputs
	- "s" param
		- empty *
		- one character *
		- many characters
	- "prefix" param
		- empty
		- one character *
		- many characters
- inputs relation
	- s has prefix
	- s hasn't prefix
	- prefix in the middle of s *
	- s == prefix

- boundaries testing
	- (on point) len(s) == len(prefix)
	- (off point) len(s) < len(prefix)

- extra scenarios
	- s has prefix but shuffled (false)

## all test cases (with * won't be tested as it is redundant; with ? i'm not sure so i will keep it)


(s="", prefix="" -> true)
(s="", prefix="a" -> false) *
(s="", prefix="abc" -> false) *

(s="a", prefix="" -> true) *
(s="a", prefix="a" -> true) *
(s="a", prefix="b" -> false) *
(s="a", prefix="abc" -> false) (off point boundary)

(s="abc", prefix="" -> true)
(s="abc", prefix="a" -> true)
(s="abc", prefix="b" -> false)
(s="abc", prefix="abc" -> true) (on point boundary)
(s="abc", prefix="bcd" -> false) ?

(s="abc", prefix="ba" -> false) ?
(s="abcd", prefix="bc", false) ?

!! as there is no reason to assume that the program handles s and prefix with 1 or many chars, it seems to be enough testing only the "empty" and "many" partitions

*/

// These are intended to better understand the function behavior and business rules
func TestHasPrefix_ExploratoryTests(t *testing.T) {
	t.Run("s and prefix empty, return true", func(t *testing.T) {
		if !HasPrefix("", "") {
			t.Errorf("expected true, got false")
		}
	})

	t.Run("s and prefix equals, return true", func(t *testing.T) {
		if !HasPrefix("abc", "abc") {
			t.Errorf("expected true, got false")
		}
	})
}

func TestHasPrefix(t *testing.T) {
	tt := []struct {
		name     string
		s        string
		prefix   string
		expected bool
	}{
		{"s empty and prefix empty", "", "", true},
		{"prefix longer than s", "a", "abc", false},
		{"prefix empty", "abc", "", true},
		{"s has prefix", "abc", "a", true},
		{"s hasn't prefix", "abc", "b", false},
		{"s equals prefix", "abc", "abc", true},
		{"s and prefix not equal but same length", "abc", "bcd", false},
		{"s has prefix but shuffled", "abc", "ba", false},
		{"s contains prefix but in the middle", "abcd", "bc", false},
	}

	for _, c := range tt {
		t.Run(c.name, func(t *testing.T) {
			if c.expected != HasPrefix(c.s, c.prefix) {
				t.Errorf("expected %v, got %v", c.expected, !c.expected)
			}
		})
	}
}
