// Package parth provides functions for accessing path segments.
//
// When returning an int of any size, the first whole number within the
// specified segment will be returned.  When returning a float of any size,
// the first decimal number within the specified segment will be returned.
package parth

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// SegmentToString receives an int representing a path segment, and returns both
// the specified segment as a string and a nil error.  If any error is
// encountered, a zero value string and error are returned.
func SegmentToString(path string, i int) (string, error) {
	if i >= 0 {
		return posSegToString(path, i)
	}
	return negSegToString(path, i)
}

// SegmentToInt64 receives an int representing a path segment, and returns both
// the specified segment as an int64 and a nil error.  If any error is
// encountered, a zero value int64 and error are returned.
func SegmentToInt64(path string, i int) (int64, error) {
	var s string
	var err error
	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}
	if s, err = findFirstIntString(s); err != nil {
		return 0, err
	}
	var v int64
	if v, err = strconv.ParseInt(s, 10, 64); err != nil {
		return 0, err
	}
	return v, nil
}

// SegmentToInt32 receives an int representing a path segment, and returns both
// the specified segment as an int32 and a nil error.  If any error is
// encountered, a zero value int32 and error are returned.
func SegmentToInt32(path string, i int) (int32, error) {
	var s string
	var err error
	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}
	if s, err = findFirstIntString(s); err != nil {
		return 0, err
	}
	var v int64
	if v, err = strconv.ParseInt(s, 10, 32); err != nil {
		return 0, err
	}
	return int32(v), nil
}

// SegmentToInt16 receives an int representing a path segment, and returns both
// the specified segment as an int16 and a nil error.  If any error is
// encountered, a zero value int16 and error are returned.
func SegmentToInt16(path string, i int) (int16, error) {
	var s string
	var err error
	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}
	if s, err = findFirstIntString(s); err != nil {
		return 0, err
	}
	var v int64
	if v, err = strconv.ParseInt(s, 10, 16); err != nil {
		return 0, err
	}
	return int16(v), nil
}

// SegmentToInt8 receives an int representing a path segment, and returns both
// the specified segment as an int8 and a nil error.  If any error is
// encountered, a zero value int8 and error are returned.
func SegmentToInt8(path string, i int) (int8, error) {
	var s string
	var err error
	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}
	if s, err = findFirstIntString(s); err != nil {
		return 0, err
	}
	var v int64
	if v, err = strconv.ParseInt(s, 10, 8); err != nil {
		return 0, err
	}
	return int8(v), nil
}

// SegmentToInt receives an int representing a path segment, and returns both
// the specified segment as an int and a nil error.  If any error is
// encountered, a zero value int and error are returned.
func SegmentToInt(path string, i int) (int, error) {
	var s string
	var err error
	if s, err = SegmentToString(path, i); err != nil {
		return 0, err
	}
	if s, err = findFirstIntString(s); err != nil {
		return 0, err
	}
	var v int64
	if v, err = strconv.ParseInt(s, 10, 0); err != nil {
		return 0, err
	}
	return int(v), nil
}

// SegmentToBool receives an int representing a path segment, and returns both
// the specified segment as a bool and a nil error.  If any error is
// encountered, a zero value bool and error are returned.
func SegmentToBool(path string, i int) (bool, error) {
	var s string
	var err error
	if s, err = SegmentToString(path, i); err != nil {
		return false, err
	}
	var v bool
	if v, err = strconv.ParseBool(s); err != nil {
		return false, err
	}
	return v, nil
}

// SegmentToFloat64 receives an int representing a path segment, and returns
// both the specified segment as a float64 and a nil error.  If any error is
// encountered, a zero value float64 and error are returned.
func SegmentToFloat64(path string, i int) (float64, error) {
	var s string
	var err error
	if s, err = SegmentToString(path, i); err != nil {
		return 0.0, err
	}
	if s, err = findFirstFloatString(s); err != nil {
		return 0.0, err
	}
	var v float64
	if v, err = strconv.ParseFloat(s, 64); err != nil {
		return 0.0, err
	}
	return v, nil
}

// SegmentToFloat32 receives an int representing a path segment, and returns
// both the specified segment as a float32 and a nil error.  If any error is
// encountered, a zero value float32 and error are returned.
func SegmentToFloat32(path string, i int) (float32, error) {
	var s string
	var err error
	if s, err = SegmentToString(path, i); err != nil {
		return 0.0, err
	}
	if s, err = findFirstFloatString(s); err != nil {
		return 0.0, err
	}
	var v float64
	if v, err = strconv.ParseFloat(s, 32); err != nil {
		return 0.0, err
	}
	return float32(v), nil
}

// SpanToString receives two int values representing path segments, and
// returns the content of and between those segments as a string and a nil
// error.  If any error is encountered, a zero value string and error are
// returned.  The segments can be of negative values, but firstSeg must come
// before the lastSeg.
func SpanToString(path string, firstSeg, lastSeg int) (string, error) {
	i := findPathIndexes(path)
	f := firstSeg
	l := lastSeg

	if f < 0 {
		f = len(i) + f - 1
	}
	if l >= 0 {
		l++
	} else {
		l = len(i) + l
	}

	if f > len(i)-2 || f < 0 {
		return "", fmt.Errorf("first path segment index %d does not exist", firstSeg)
	}
	if l > len(i) || l < 0 {
		return "", fmt.Errorf("last path segment index %d does not exist", lastSeg)
	}
	if f > l { // or equal?
		return "", fmt.Errorf("first segment must come before the last segment")
	}
	if i[f] >= i[l] { // or equal?
		f, l = l-1, f+1
	}

	return path[i[f]:i[l]], nil
}

func posSegToString(path string, i int) (string, error) {
	c, ind0, ind1 := 0, 0, 0
	for n := 0; n < len(path); n++ {
		if path[n] == '/' {
			if c == i {
				if n+1 < len(path) && path[n+1] != '/' {
					ind0 = n + 1
				} else {
					break
				}
			}
			if c > i {
				ind1 = n
				break
			}
			c++
		} else if n == 0 {
			if c == i {
				ind0 = n
			}
			c++
		} else if n == len(path)-1 {
			if c > i {
				ind1 = n + 1
			}
			break
		}
	}
	if i < 0 || ind1 == 0 {
		return "", fmt.Errorf("path segment index %d does not exist", i)
	}
	return path[ind0:ind1], nil
}

func negSegToString(path string, i int) (string, error) {
	i = i * -1
	c, ind0, ind1 := 1, 0, 0
	for n := len(path) - 1; n >= 0; n-- {
		if path[n] == '/' {
			if c == i {
				if n-1 >= 0 && path[n-1] != '/' {
					ind1 = n
				} else {
					break
				}
			}
			if c > i {
				ind0 = n + 1
				break
			}
			c++
		} else if n == len(path)-1 {
			if c == i {
				ind1 = n + 1
			}
			c++
		} else if n == 0 {
			if c > i {
				ind0 = n + 1
			}
			break
		}
	}
	if i < 1 || ind0 == 0 {
		return "", fmt.Errorf("path segment index %d does not exist", i*-1)
	}
	return path[ind0:ind1], nil
}

func findFirstIntString(s string) (string, error) {
	ind, l := 0, 0
	for n := 0; n < len(s); n++ {
		if unicode.IsDigit(rune(s[n])) {
			if l == 0 {
				ind = n
			}
			l++
		} else if s[n] == '-' {
			if l == 0 {
				ind = n
				l++
			} else {
				break
			}
		} else {
			if l == 0 && s[n] == '.' {
				if n+1 < len(s) && unicode.IsDigit(rune(s[n+1])) {
					return "0", nil
				}
				break
			}
			if l > 0 {
				break
			}
		}
	}

	if l == 0 {
		return "", errors.New("path segment does not contain int")
	}
	return s[ind : ind+l], nil
}

func findFirstFloatString(s string) (string, error) {
	c, ind, l := 0, 0, 0
	for n := 0; n < len(s); n++ {
		if unicode.IsDigit(rune(s[n])) {
			if l == 0 {
				ind = n
			}
			l++
		} else if s[n] == '-' {
			if l == 0 {
				ind = n
				l++
			} else {
				break
			}
		} else if s[n] == '.' {
			if l == 0 {
				ind = n
			}
			if c > 0 {
				break
			}
			l++
			c++
		} else if s[n] == 'e' && l > 0 && n+1 < len(s) && s[n+1] == '+' {
			l++
		} else if s[n] == '+' && l > 0 && s[n-1] == 'e' {
			if n+1 < len(s) && unicode.IsDigit(rune(s[n+1])) {
				l++
				continue
			}
			l--
			break
		} else {
			if l > 0 {
				break
			}
		}
	}

	if l == 0 || s[ind:ind+l] == "." {
		return "", errors.New("path segment does not contain float")
	}
	return s[ind : ind+l], nil
}

func findPathIndexes(path string) []int {
	i := make([]int, 1, len(path))
	for n := 0; n < len(path); n++ {
		if (n > 0 && path[n] == '/') || n == len(path)-1 {
			if n == len(path)-1 && path[n] != '/' {
				n++
			}
			i = append(i, n)
		}
	}
	return i
}
