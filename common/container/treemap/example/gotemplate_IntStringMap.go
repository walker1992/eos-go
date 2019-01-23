// Code generated by gotemplate. DO NOT EDIT.

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treemap implements a map backed by red-black Tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package example

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/eosspark/container/utils"
	"github.com/eosspark/eos-go/common/container"
	rbt "github.com/eosspark/eos-go/common/container/tree"
)

// template type Map(K,V,Compare,Multi)

func assertIntStringMapImplementation() {
	var _ container.Map = (*IntStringMap)(nil)
}

// Map holds the elements in a red-black Tree
type IntStringMap struct {
	*rbt.Tree
}

// NewWith instantiates a Tree map with the custom comparator.
func NewIntStringMap() *IntStringMap {
	return &IntStringMap{Tree: rbt.NewWith(utils.IntComparator, false)}
}

func CopyFromIntStringMap(tm *IntStringMap) *IntStringMap {
	return &IntStringMap{Tree: rbt.CopyFrom(tm.Tree)}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *IntStringMap) Put(key int, value string) {
	m.Tree.Put(key, value)
}

func (m *IntStringMap) Insert(key int, value string) IteratorIntStringMap {
	return IteratorIntStringMap{m.Tree.Insert(key, value)}
}

// Get searches the element in the map by key and returns its value or nil if key is not found in Tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *IntStringMap) Get(key int) IteratorIntStringMap {
	return IteratorIntStringMap{m.Tree.Get(key)}
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *IntStringMap) Remove(key int) {
	m.Tree.Remove(key)
}

// Keys returns all keys in-order
func (m *IntStringMap) Keys() []int {
	keys := make([]int, m.Tree.Size())
	it := m.Tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key().(int)
	}
	return keys
}

// Values returns all values in-order based on the key.
func (m *IntStringMap) Values() []string {
	values := make([]string, m.Tree.Size())
	it := m.Tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value().(string)
	}
	return values
}

// Each calls the given function once for each element, passing that element's key and value.
func (m *IntStringMap) Each(f func(key int, value string)) {
	Iterator := m.Iterator()
	for Iterator.Next() {
		f(Iterator.Key(), Iterator.Value())
	}
}

// Find passes each element of the container to the given function and returns
// the first (key,value) for which the function is true or nil,nil otherwise if no element
// matches the criteria.
func (m *IntStringMap) Find(f func(key int, value string) bool) (k int, v string) {
	Iterator := m.Iterator()
	for Iterator.Next() {
		if f(Iterator.Key(), Iterator.Value()) {
			return Iterator.Key(), Iterator.Value()
		}
	}
	return
}

// String returns a string representation of container
func (m IntStringMap) String() string {
	str := "TreeMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}

// Iterator holding the Iterator's state
type IteratorIntStringMap struct {
	rbt.Iterator
}

// Iterator returns a stateful Iterator whose elements are key/value pairs.
func (m *IntStringMap) Iterator() IteratorIntStringMap {
	return IteratorIntStringMap{Iterator: m.Tree.Iterator()}
}

// Begin returns First Iterator whose position points to the first element
// Return End Iterator when the map is empty
func (m *IntStringMap) Begin() IteratorIntStringMap {
	return IteratorIntStringMap{m.Tree.Begin()}
}

// End returns End Iterator
func (m *IntStringMap) End() IteratorIntStringMap {
	return IteratorIntStringMap{m.Tree.End()}
}

// Value returns the current element's value.
// Does not modify the state of the Iterator.
func (iterator IteratorIntStringMap) Value() string {
	return iterator.Iterator.Value().(string)
}

// Key returns the current element's key.
// Does not modify the state of the Iterator.
func (iterator IteratorIntStringMap) Key() int {
	return iterator.Iterator.Key().(int)
}

func (m *IntStringMap) LowerBound(key int) IteratorIntStringMap {
	return IteratorIntStringMap{m.Tree.LowerBound(key)}
}

func (m *IntStringMap) UpperBound(key int) IteratorIntStringMap {
	return IteratorIntStringMap{m.Tree.UpperBound(key)}

}

// ToJSON outputs the JSON representation of the map.
type pairIntStringMap struct {
	Key int    `json:"key"`
	Val string `json:"val"`
}

func (m IntStringMap) MarshalJSON() ([]byte, error) {
	elements := make([]pairIntStringMap, 0, m.Size())
	it := m.Iterator()
	for it.Next() {
		elements = append(elements, pairIntStringMap{it.Key(), it.Value()})
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *IntStringMap) UnmarshalJSON(data []byte) error {
	elements := make([]pairIntStringMap, 0)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for _, pair := range elements {
			m.Put(pair.Key, pair.Val)
		}
	}
	return err
}
