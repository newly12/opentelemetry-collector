// Copyright The OpenTelemetry Authors
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

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package internal

// ImmutableByteSlice represents a []byte slice that cannot be mutated.
// The instance of ImmutableByteSlice can be assigned to multiple objects since it's immutable.
type ImmutableByteSlice struct {
	value []byte
}

// NewImmutableByteSlice creates a new ImmutableByteSlice by copying the provided []byte slice.
func NewImmutableByteSlice(val []byte) ImmutableByteSlice {
	is := ImmutableByteSlice{}
	if len(val) != 0 {
		is.value = make([]byte, len(val))
		copy(is.value, val)
	}
	return is
}

// AsRaw returns a copy of the []byte slice.
func (is ImmutableByteSlice) AsRaw() []byte {
	if len(is.value) == 0 {
		return nil
	}
	val := make([]byte, len(is.value))
	copy(val, is.value)
	return val
}

// Len returns length of the []byte slice value.
func (is ImmutableByteSlice) Len() int {
	return len(is.value)
}

// At returns an item from particular index.
func (is ImmutableByteSlice) At(i int) byte {
	return is.value[i]
}

// ImmutableFloat64Slice represents a []float64 slice that cannot be mutated.
// The instance of ImmutableFloat64Slice can be assigned to multiple objects since it's immutable.
type ImmutableFloat64Slice struct {
	value []float64
}

// NewImmutableFloat64Slice creates a new ImmutableFloat64Slice by copying the provided []float64 slice.
func NewImmutableFloat64Slice(val []float64) ImmutableFloat64Slice {
	is := ImmutableFloat64Slice{}
	if len(val) != 0 {
		is.value = make([]float64, len(val))
		copy(is.value, val)
	}
	return is
}

// AsRaw returns a copy of the []float64 slice.
func (is ImmutableFloat64Slice) AsRaw() []float64 {
	if len(is.value) == 0 {
		return nil
	}
	val := make([]float64, len(is.value))
	copy(val, is.value)
	return val
}

// Len returns length of the []float64 slice value.
func (is ImmutableFloat64Slice) Len() int {
	return len(is.value)
}

// At returns an item from particular index.
func (is ImmutableFloat64Slice) At(i int) float64 {
	return is.value[i]
}

// ImmutableUInt64Slice represents a []uint64 slice that cannot be mutated.
// The instance of ImmutableUInt64Slice can be assigned to multiple objects since it's immutable.
type ImmutableUInt64Slice struct {
	value []uint64
}

// NewImmutableUInt64Slice creates a new ImmutableUInt64Slice by copying the provided []uint64 slice.
func NewImmutableUInt64Slice(val []uint64) ImmutableUInt64Slice {
	is := ImmutableUInt64Slice{}
	if len(val) != 0 {
		is.value = make([]uint64, len(val))
		copy(is.value, val)
	}
	return is
}

// AsRaw returns a copy of the []uint64 slice.
func (is ImmutableUInt64Slice) AsRaw() []uint64 {
	if len(is.value) == 0 {
		return nil
	}
	val := make([]uint64, len(is.value))
	copy(val, is.value)
	return val
}

// Len returns length of the []uint64 slice value.
func (is ImmutableUInt64Slice) Len() int {
	return len(is.value)
}

// At returns an item from particular index.
func (is ImmutableUInt64Slice) At(i int) uint64 {
	return is.value[i]
}