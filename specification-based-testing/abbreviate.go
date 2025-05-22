package main

import (
	"fmt"
)

// Adapted from Apache's Commons Lang. Check the original Java implementation here: https://github.com/apache/commons-lang/blob/e0b474c0d015f89a52c4cf8866fa157dd89e7d1c/src/main/java/org/apache/commons/lang3/StringUtils.java#L332C26-L332C36
//
// Abbreviates a String using a given replacement marker. This will turn
// "Now is the time for all good men" into "...is the time for..." if "..." was defined
// as the replacement marker.
//
// It allows you to specify a "left edge" offset.  Note that this left edge is not necessarily going to
// be the leftmost character in the result, or the first character following the
// replacement marker, but it will appear somewhere in the result.
//
// In no case will it return a String of length greater than [maxWidth]
func Abbreviate(str, abbrevMarker string, offset, maxWidth int) (string, error) {
	if str == "" && abbrevMarker == "" {
		return str, nil
	} else if str != "" && abbrevMarker == "" && maxWidth > 0 {
		if maxWidth > len(str) {
			return str, nil
		}
		return str[:maxWidth], nil
	} else if str == "" || abbrevMarker == "" {
		return str, nil
	}

	abbrevMarkerLength := len(abbrevMarker)
	minAbbrevWidth := abbrevMarkerLength + 1
	minAbbrevWidthOffset := abbrevMarkerLength*2 + 1

	lenStr := len(str)

	if maxWidth < minAbbrevWidth {
		return "", fmt.Errorf("minimum abbreviation width is %d", minAbbrevWidth)
	}

	if lenStr <= maxWidth {
		return str, nil
	}

	if offset > lenStr {
		offset = lenStr
	}

	// wtf is this? is it really necessary?
	if lenStr-offset < maxWidth-abbrevMarkerLength {
		offset = lenStr - (maxWidth - abbrevMarkerLength)
	}

	if offset <= abbrevMarkerLength+1 {
		return str[:maxWidth-abbrevMarkerLength] + abbrevMarker, nil
	}

	if maxWidth < minAbbrevWidthOffset {
		return "", fmt.Errorf("minimum abbreviation width with offset is %d", minAbbrevWidthOffset)
	}

	if offset+maxWidth-abbrevMarkerLength < lenStr {
		abbrevPart, err := Abbreviate(str[offset:], abbrevMarker, 0, maxWidth-abbrevMarkerLength)
		if err != nil {
			return "", err
		}
		return abbrevMarker + abbrevPart, nil
	}

	return abbrevMarker + str[lenStr-(maxWidth-abbrevMarkerLength):], nil
}
