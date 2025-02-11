// Copyright 2022 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package seclang

import (
	"testing"

	"github.com/corazawaf/coraza/v2"
	"github.com/corazawaf/coraza/v2/types/variables"
	"github.com/stretchr/testify/require"
)

/*
func TestRequestBodyAccessOff(t *testing.T) {
	waf := coraza.NewWaf()
	parser, _ := NewParser(waf)
	if err := parser.FromString(`
	SecRequestBodyAccess Off
	`); err != nil {
		t.Fatal(err)
	}
	tx := waf.NewTransaction()
	tx.ProcessURI("/", "POST", "http/1.1")
	tx.RequestBodyBuffer.Write([]byte("test=123"))
	tx.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")
	tx.ProcessRequestHeaders()
	tx.ProcessRequestBody()
	if len(tx.GetCollection(variables.ArgsPost).Data()) != 0 {
		t.Error("Should not have args")
	}
}*/

func TestRequestBodyAccessOn(t *testing.T) {
	waf := coraza.NewWaf()
	parser, _ := NewParser(waf)

	err := parser.FromString(`
	SecRequestBodyAccess On
	`)
	require.NoError(t, err)

	tx := waf.NewTransaction()
	tx.ProcessURI("/", "POST", "http/1.1")
	_, err = tx.RequestBodyBuffer.Write([]byte("test=123"))
	require.NoError(t, err)

	tx.AddRequestHeader("Content-Type", "application/x-www-form-urlencoded")
	tx.ProcessRequestHeaders()
	_, err = tx.ProcessRequestBody()
	require.NoError(t, err)
	require.NotEmpty(t, tx.GetCollection(variables.ArgsPost).Data(), "should have args")
}
