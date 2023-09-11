// @generated Code generated by gen-atomicwrapper.

// Copyright (c) 2020-2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package atomic

import (
	"time"
)

// Time is an atomic type-safe wrapper for time.Time values.
type Time struct {
	_ nocmp // disallow non-atomic comparison

	v Value
}

var _zeroTime time.Time

// NewTime creates a new Time.
func NewTime(val time.Time) *Time {
	x := &Time{}
	if val != _zeroTime {
		x.Store(val)
	}
	return x
}

// Load atomically loads the wrapped time.Time.
func (x *Time) Load() time.Time {
	return unpackTime(x.v.Load())
}

// Store atomically stores the passed time.Time.
func (x *Time) Store(val time.Time) {
	x.v.Store(packTime(val))
}
