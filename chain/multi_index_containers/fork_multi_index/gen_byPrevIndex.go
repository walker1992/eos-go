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
package fork_multi_index

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/eosspark/container/templates"
	rbt "github.com/eosspark/container/templates/tree"
	"github.com/eosspark/eos-go/common"
)

// template type Map(K,V,Compare)

func assertByPrevIndexImplementation() {
	var _ templates.Map = (*byPrevIndex)(nil)
}

// Map holds the elements in a red-black Tree
type byPrevIndex struct {
	*rbt.Tree
}

// NewWith instantiates a Tree map with the custom comparator.
func newByPrevIndex() *byPrevIndex {
	return &byPrevIndex{Tree: rbt.NewWith(byPrevCompare, false)}
}

func copyFromByPrevIndex(tm *byPrevIndex) *byPrevIndex {
	return &byPrevIndex{Tree: rbt.CopyFrom(tm.Tree)}
}

type multiByPrevIndex = byPrevIndex

func newMultiByPrevIndex() *multiByPrevIndex {
	return &byPrevIndex{Tree: rbt.NewWith(byPrevCompare, true)}
}

func copyMultiFromByPrevIndex(tm *byPrevIndex) *byPrevIndex {
	return &byPrevIndex{Tree: rbt.CopyFrom(tm.Tree)}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *byPrevIndex) Put(key common.BlockIdType, value IndexKey) {
	m.Tree.Put(key, value)
}

func (m *byPrevIndex) Insert(key common.BlockIdType, value IndexKey) iteratorByPrevIndex {
	return iteratorByPrevIndex{m.Tree.Insert(key, value)}
}

// Get searches the element in the map by key and returns its value or nil if key is not found in Tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *byPrevIndex) Get(key common.BlockIdType) iteratorByPrevIndex {
	return iteratorByPrevIndex{m.Tree.Get(key)}
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *byPrevIndex) Remove(key common.BlockIdType) {
	m.Tree.Remove(key)
}

// Keys returns all keys in-order
func (m *byPrevIndex) Keys() []common.BlockIdType {
	keys := make([]common.BlockIdType, m.Tree.Size())
	it := m.Tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key().(common.BlockIdType)
	}
	return keys
}

// Values returns all values in-order based on the key.
func (m *byPrevIndex) Values() []IndexKey {
	values := make([]IndexKey, m.Tree.Size())
	it := m.Tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value().(IndexKey)
	}
	return values
}

// Each calls the given function once for each element, passing that element's key and value.
func (m *byPrevIndex) Each(f func(key common.BlockIdType, value IndexKey)) {
	Iterator := m.Iterator()
	for Iterator.Next() {
		f(Iterator.Key(), Iterator.Value())
	}
}

// Find passes each element of the container to the given function and returns
// the first (key,value) for which the function is true or nil,nil otherwise if no element
// matches the criteria.
func (m *byPrevIndex) Find(f func(key common.BlockIdType, value IndexKey) bool) (k common.BlockIdType, v IndexKey) {
	Iterator := m.Iterator()
	for Iterator.Next() {
		if f(Iterator.Key(), Iterator.Value()) {
			return Iterator.Key(), Iterator.Value()
		}
	}
	return
}

// String returns a string representation of container
func (m byPrevIndex) String() string {
	str := "TreeMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}

// Iterator holding the Iterator's state
type iteratorByPrevIndex struct {
	rbt.Iterator
}

// Iterator returns a stateful Iterator whose elements are key/value pairs.
func (m *byPrevIndex) Iterator() iteratorByPrevIndex {
	return iteratorByPrevIndex{Iterator: m.Tree.Iterator()}
}

// Begin returns First Iterator whose position points to the first element
// Return End Iterator when the map is empty
func (m *byPrevIndex) Begin() iteratorByPrevIndex {
	return iteratorByPrevIndex{m.Tree.Begin()}
}

// End returns End Iterator
func (m *byPrevIndex) End() iteratorByPrevIndex {
	return iteratorByPrevIndex{m.Tree.End()}
}

// Value returns the current element's value.
// Does not modify the state of the Iterator.
func (Iterator *iteratorByPrevIndex) Value() IndexKey {
	return Iterator.Iterator.Value().(IndexKey)
}

// Key returns the current element's key.
// Does not modify the state of the Iterator.
func (Iterator *iteratorByPrevIndex) Key() common.BlockIdType {
	return Iterator.Iterator.Key().(common.BlockIdType)
}

func (m *byPrevIndex) LowerBound(key common.BlockIdType) iteratorByPrevIndex {
	return iteratorByPrevIndex{m.Tree.LowerBound(key)}
}

func (m *byPrevIndex) UpperBound(key common.BlockIdType) iteratorByPrevIndex {
	return iteratorByPrevIndex{m.Tree.UpperBound(key)}

}

// ToJSON outputs the JSON representation of the map.
type pairByPrevIndex struct {
	Key common.BlockIdType `json:"key"`
	Val IndexKey           `json:"val"`
}

func (m byPrevIndex) MarshalJSON() ([]byte, error) {
	elements := make([]pairByPrevIndex, 0, m.Size())
	it := m.Iterator()
	for it.Next() {
		elements = append(elements, pairByPrevIndex{it.Key(), it.Value()})
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *byPrevIndex) UnmarshalJSON(data []byte) error {
	elements := make([]pairByPrevIndex, 0)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for _, pair := range elements {
			m.Put(pair.Key, pair.Val)
		}
	}
	return err
}