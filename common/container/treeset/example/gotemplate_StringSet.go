// Code generated by gotemplate. DO NOT EDIT.

// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treeset implements a Tree backed by a red-black Tree.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package example

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/eosspark/eos-go/common"
	"github.com/eosspark/eos-go/common/container"
	rbt "github.com/eosspark/eos-go/common/container/redblacktree"
	"github.com/eosspark/eos-go/crypto/rlp"
)

// template type Set(V,Compare,Multi)

func assertStringSetImplementation() {
	var _ container.Set = (*StringSet)(nil)
}

// Set holds elements in a red-black Tree
type StringSet struct {
	*rbt.Tree
}

var itemExistsStringSet = struct{}{}

// NewWith instantiates a new empty set with the custom comparator.

func NewStringSet(Value ...string) *StringSet {
	set := &StringSet{Tree: rbt.NewWith(StringComparator, false)}
	set.Add(Value...)
	return set
}

func CopyFromStringSet(ts *StringSet) *StringSet {
	return &StringSet{Tree: rbt.CopyFrom(ts.Tree)}
}

func StringSetIntersection(a *StringSet, b *StringSet, callback func(elem string)) {
	aIterator := a.Iterator()
	bIterator := b.Iterator()

	if !aIterator.First() || !bIterator.First() {
		return
	}

	for aHasNext, bHasNext := true, true; aHasNext && bHasNext; {
		comp := StringComparator(aIterator.Value(), bIterator.Value())
		switch {
		case comp > 0:
			bHasNext = bIterator.Next()
		case comp < 0:
			aHasNext = aIterator.Next()
		default:
			callback(aIterator.Value())
			aHasNext = aIterator.Next()
			bHasNext = bIterator.Next()
		}
	}
}

// Add adds the item one to the set.Returns false and the interface if it already exists
func (set *StringSet) AddItem(item string) (bool, string) {
	itr := set.Tree.Insert(item, itemExistsStringSet)
	if itr.IsEnd() {
		return false, item
	}
	return true, itr.Key().(string)
}

// Add adds the items (one or more) to the set.
func (set *StringSet) Add(items ...string) {
	for _, item := range items {
		set.Tree.Put(item, itemExistsStringSet)
	}
}

// Remove removes the items (one or more) from the set.
func (set *StringSet) Remove(items ...string) {
	for _, item := range items {
		set.Tree.Remove(item)
	}

}

// Values returns all items in the set.
func (set *StringSet) Values() []string {
	keys := make([]string, set.Size())
	it := set.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Value()
	}
	return keys
}

// Contains checks weather items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *StringSet) Contains(items ...string) bool {
	for _, item := range items {
		if iter := set.Get(item); iter.IsEnd() {
			return false
		}
	}
	return true
}

// String returns a string representation of container
func (set *StringSet) String() string {
	str := "TreeSet\n"
	items := make([]string, 0)
	for _, v := range set.Tree.Keys() {
		items = append(items, fmt.Sprintf("%v", v))
	}
	str += strings.Join(items, ", ")
	return str
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
type IteratorStringSet struct {
	rbt.Iterator
}

// Iterator holding the iterator's state
func (set *StringSet) Iterator() IteratorStringSet {
	return IteratorStringSet{Iterator: set.Tree.Iterator()}
}

// Begin returns First Iterator whose position points to the first element
// Return End Iterator when the map is empty
func (set *StringSet) Begin() IteratorStringSet {
	return IteratorStringSet{set.Tree.Begin()}
}

// End returns End Iterator
func (set *StringSet) End() IteratorStringSet {
	return IteratorStringSet{set.Tree.End()}
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator IteratorStringSet) Value() string {
	return iterator.Iterator.Key().(string)
}

// Each calls the given function once for each element, passing that element's index and value.
func (set *StringSet) Each(f func(value string)) {
	iterator := set.Iterator()
	for iterator.Next() {
		f(iterator.Value())
	}
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (set *StringSet) Find(f func(value string) bool) (v string) {
	iterator := set.Iterator()
	for iterator.Next() {
		if f(iterator.Value()) {
			return iterator.Value()
		}
	}
	return
}

func (set *StringSet) LowerBound(item string) IteratorStringSet {
	return IteratorStringSet{set.Tree.LowerBound(item)}
}

func (set *StringSet) UpperBound(item string) IteratorStringSet {
	return IteratorStringSet{set.Tree.UpperBound(item)}
}

// ToJSON outputs the JSON representation of the set.
func (set StringSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *StringSet) UnmarshalJSON(data []byte) error {
	elements := make([]string, 0)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Tree = rbt.NewWith(StringComparator, false)
		set.Add(elements...)
	}
	return err
}

func (set StringSet) Pack() (re []byte, err error) {
	re = append(re, common.WriteUVarInt(set.Size())...)
	set.Each(func(value string) {
		reVal, _ := rlp.EncodeToBytes(value)
		re = append(re, reVal...)
	})
	return re, nil
}

func (set *StringSet) Unpack(in []byte) (int, error) {
	set.Tree = rbt.NewWith(StringComparator, false)

	decoder := rlp.NewDecoder(in)
	l, err := decoder.ReadUvarint64()
	if err != nil {
		return 0, err
	}

	for i := 0; i < int(l); i++ {
		v := new(string)
		decoder.Decode(v)
		set.Add(*v)
	}
	return decoder.GetPos(), nil
}
