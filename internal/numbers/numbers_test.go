// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be found
// in the LICENSE file.

package numbers_test

import (
	"testing"

	"github.com/jelmersnoeck/cacher/internal/numbers"
)

func TestBytesString(t *testing.T) {
	str := "Hello world"
	btsStr := []byte(str)

	if numbers.BytesString(btsStr) != str {
		t.Errorf("Bytes should be the same as provided string.")
		t.FailNow()
	}
}

func TestInt64Bytes(t *testing.T) {
	val := int64(123235)

	conVal, ok := numbers.BytesInt64(numbers.Int64Bytes(val))
	if !ok || conVal != val {
		t.Errorf("Converted bytes for int64 should be the same as value.")
		t.FailNow()
	}
}
