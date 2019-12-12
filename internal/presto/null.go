// Copyright 2019 Grabtaxi Holdings PTE LTE (GRAB), All rights reserved.
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file

package presto

import (
	"fmt"

	"github.com/grab/talaria/internal/encoding/typeof"
)

const maxNullCount = 25000

var zNulls = make([]bool, maxNullCount)
var zInt32 = make([]int32, maxNullCount)
var zInt64 = make([]int64, maxNullCount)
var zFloat64 = make([]float64, maxNullCount)
var zBool = make([]bool, maxNullCount)

func init() {
	for i := 0; i < maxNullCount; i++ {
		zNulls[i] = true
	}
}

// NullColumn creates a new null block
func NullColumn(t typeof.Type, count int) Column {
	switch t {
	case typeof.String:
		return &PrestoThriftVarchar{
			Nulls: zNulls[:count],
			Sizes: zInt32[:count],
			Bytes: []byte{},
		}
	case typeof.Int32:
		return &PrestoThriftInteger{
			Nulls: zNulls[:count],
			Ints:  zInt32[:count],
		}
	case typeof.Int64:
		return &PrestoThriftBigint{
			Nulls: zNulls[:count],
			Longs: zInt64[:count],
		}
	case typeof.Float64:
		return &PrestoThriftDouble{
			Nulls:   zNulls[:count],
			Doubles: zFloat64[:count],
		}
	case typeof.Bool:
		return &PrestoThriftBoolean{
			Nulls:    zNulls[:count],
			Booleans: zBool[:count],
		}
	case typeof.Timestamp:
		return &PrestoThriftTimestamp{
			Nulls:      zNulls[:count],
			Timestamps: zInt64[:count],
		}
	case typeof.JSON:
		return &PrestoThriftJson{
			Nulls: zNulls[:count],
			Sizes: zInt32[:count],
			Bytes: []byte{},
		}
	}

	panic(fmt.Errorf("presto: unknown type %v", t))
}
