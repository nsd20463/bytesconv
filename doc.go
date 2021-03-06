// Copyright 2016 Nicolas Dade. Derived from the go 1.7.4
// standard library, which itself was:
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bytesconv implements conversions from []byte representations
// of basic data types. These are the routines missing from the standard
// library's bytesconv. When parsing input data, which naturally arrives as
// []byte, these make parsing the input free of garbage creation. (Using
// the bytesconv.ParseXxx functions creates a temporary string which the compiler
// isn't (in 1.7) good enough to eliminate. There's a long discussion about
// the need for these []byte parsing functions in the go, and the masters
// are still, years later, waiting for the compiler to get it right. I can't
// wait for my code to get it right, hence this largely C&Ved package.)
//
// Note that only conversions from []byte are implemented. Converting to []byte
// is already handled by the standard library's strconv package's AppendXxx
// functions.
//
// Numeric Conversions
//
// ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:
//
//	b, err := bytesconv.ParseBool([]byte("true"))
//	f, err := bytesconv.ParseFloat([]byte("3.1415"), 64)
//	i, err := bytesconv.ParseInt([]byte("-42"), 10, 64)
//	u, err := bytesconv.ParseUint([]byte("42"), 10, 64)
//
// The parse functions return the widest type (float64, int64, and uint64),
// but if the size argument specifies a narrower width the result can be
// converted to that narrower type without data loss:
//
//	s := []byte("2147483647") // biggest int32
//	i64, err := bytesconv.ParseInt(s, 10, 32)
//	...
//	i := int32(i64)
//
package bytesconv
