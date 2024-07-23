// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Written by Cristian Maglie.

package json_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.bug.st/json"
)

// Test the extra features provided by the go.bug.st/json package

func TestRequiredField(t *testing.T) {
	type s struct {
		Required  string `json:"req,required"`
		Optional  string `json:"opt"`
		OmitEmpty string `json:"omit,omitempty"`
	}
	{
		var x s
		err := json.Unmarshal([]byte(`{}`), &x)
		require.EqualError(t, err, "json: undefined required field req into json_test.s")
	}
	{
		var x s
		err := json.Unmarshal([]byte(`{"req":"hello"}`), &x)
		require.NoError(t, err)
		require.Equal(t, s{Required: "hello"}, x)
	}
	{
		var x s
		err := json.Unmarshal([]byte(`{"opt":"hello"}`), &x)
		require.EqualError(t, err, "json: undefined required field req into json_test.s")
	}
	{
		var x s
		err := json.Unmarshal([]byte(`{"omit":"hello"}`), &x)
		require.EqualError(t, err, "json: undefined required field req into json_test.s")
	}
	{
		var x s
		err := json.Unmarshal([]byte(`{"opt":"hello","omit":"hello"}`), &x)
		require.EqualError(t, err, "json: undefined required field req into json_test.s")
	}
	{
		var x s
		err := json.Unmarshal([]byte(`{"req":"HI!","opt":"hello","omit":"hello"}`), &x)
		require.NoError(t, err)
		require.Equal(t, s{Required: "HI!", Optional: "hello", OmitEmpty: "hello"}, x)
	}
}
