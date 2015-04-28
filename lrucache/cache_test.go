// Copyright 2015 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lrucache_test

import (
	"testing"

	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
	"github.com/jacobsa/util/lrucache"
)

func TestCache(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////////////////

const capacity = 3

type CacheTest struct {
	cache lrucache.Cache
}

func init() { RegisterTestSuite(&CacheTest{}) }

func (t *CacheTest) SetUp(ti *TestInfo) {
	panic("TODO")
}

////////////////////////////////////////////////////////////////////////
// Test functions
////////////////////////////////////////////////////////////////////////

func (t *CacheTest) LookUpInEmptyCache() {
	ExpectEq(nil, t.cache.LookUp(""))
	ExpectEq(nil, t.cache.LookUp("taco"))
}

func (t *CacheTest) InsertNilValue() {
	ExpectThat(
		func() { t.cache.Insert("taco", nil) },
		Panics(HasSubstr("nil value")),
	)
}

func (t *CacheTest) LookUpUnknownKey() {
	AssertFalse(true, "TODO")
}

func (t *CacheTest) FillUpToCapacity() {
	AssertFalse(true, "TODO")
}

func (t *CacheTest) ExpiresLeastRecentlyUsed() {
	AssertFalse(true, "TODO")
}

func (t *CacheTest) Overwrite() {
	AssertFalse(true, "TODO")
}

func (t *CacheTest) Encode_EmptyCache() {
	AssertFalse(true, "TODO")
}

func (t *CacheTest) Encode_PreservesLruOrderAndCapacity() {
	AssertFalse(true, "TODO")
}
