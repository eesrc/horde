package model
//
//Copyright 2019 Telenor Digital AS
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
import (
	"strconv"
)

// storageKey is the type used for keys
type storageKey uint64

const keyBase = 24

func (k storageKey) String() string {
	return strconv.FormatUint(uint64(k), keyBase)
}

// NewKeyFromString creates a new key from a string
func newKeyFromString(id string) (storageKey, error) {
	v, err := strconv.ParseUint(id, keyBase, 64)
	return storageKey(v), err
}
