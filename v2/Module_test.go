/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package module_test

import (
	fmt "fmt"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ass "github.com/stretchr/testify/assert"
	stc "strconv"
	tes "testing"
)

type Foolish interface {
	GetFoo() int
	GetBar() string
	GetAny() any
}

func FooBar(foo int, bar string, baz any) Foolish {
	return &foobar{foo, bar, baz}
}

type foobar struct {
	foo int
	bar string
	Nil any
}

func (v *foobar) GetFoo() int    { return v.foo }
func (v foobar) GetFoo2() int    { return v.foo }
func (v *foobar) GetBar() string { return v.bar }
func (v foobar) GetBar2() string { return v.bar }
func (v *foobar) GetAny() any    { return nil }
func (v foobar) GetAny2() any    { return nil }
func (v *foobar) String() string { return v.bar + "-" + stc.Itoa(v.foo) }

func TestImplementsType(t *tes.T) {
	var aspect Foolish
	var value any
	ass.False(t, uti.ImplementsType(value, (*Foolish)(nil)))
	value = "foolish"
	ass.False(t, uti.ImplementsType(value, (*Foolish)(nil)))
	value = FooBar(5, "five", aspect)
	ass.True(t, uti.ImplementsType(value, (*Foolish)(nil)))
}

func TestIsDefined(t *tes.T) {
	var integer int
	ass.True(t, uti.IsDefined(integer))
	integer = 5
	ass.True(t, uti.IsDefined(integer))

	var name string
	ass.False(t, uti.IsDefined(name))
	name = ""
	ass.False(t, uti.IsDefined(name))
	name = "foobar"
	ass.True(t, uti.IsDefined(name))

	var slice []int
	ass.False(t, uti.IsDefined(slice))
	slice = []int{}
	ass.True(t, uti.IsDefined(slice))
	slice = []int{1, 2, 3}
	ass.True(t, uti.IsDefined(slice))
}

var booleanFalse = false
var booleanTrue = true
var byte16 = byte(16)
var rune1024 = rune(1024)
var uint85 = uint8(5)
var int13 = 13
var float = 1.23e10
var complex4 = 4 + 0i
var complex5i = 5i
var stringHello = "Hello World!"

func TestPrimitives(t *tes.T) {
	fmt.Println("Primitives")
	fmt.Println(uti.Format(booleanFalse))
	fmt.Println(uti.Format(byte16))
	fmt.Println(uti.Format(rune1024))
	fmt.Println(uti.Format(uint85))
	fmt.Println(uti.Format(int13))
	fmt.Println(uti.Format(float))
	fmt.Println(uti.Format(complex5i))
	fmt.Println(uti.Format(stringHello))
	fmt.Println()
	var mapOfAny = map[any]any{
		booleanFalse: booleanFalse,
		byte16:       byte16,
		rune1024:     rune1024,
		complex5i:    complex5i,
		uint85:       uint85,
		int13:        int13,
		float:        float,
		complex4:     complex4,
		stringHello:  stringHello,
		booleanTrue:  booleanTrue,
	}
	fmt.Println(uti.Format(mapOfAny))
	fmt.Println()
}

func TestArrays(t *tes.T) {
	fmt.Println("Arrays")
	var empty = []int{}
	fmt.Println(uti.Format(empty))

	var array = []any{nil, nil}
	array[0] = array
	array[1] = &Association{
		key:   "aKey",
		value: "aValue",
	}
	fmt.Println(uti.Format(array))

	var first = []int{1, 2, 3}
	var second = uti.CopyArray(first)
	ass.True(t, uti.ArraysAreEqual(first, first))
	ass.True(t, uti.ArraysAreEqual(first, second))
	ass.True(t, uti.ArraysAreEqual(second, first))
	first[1] = 5
	ass.False(t, uti.ArraysAreEqual(first, second))
	ass.False(t, uti.ArraysAreEqual(second, first))
	fmt.Println(uti.Format(first))
	fmt.Println()
}

func TestMaps(t *tes.T) {
	fmt.Println("Maps")
	var empty = map[string]int{}
	fmt.Println(uti.Format(empty))

	var first = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	var second = uti.CopyMap(first)
	ass.True(t, uti.MapsAreEqual(first, first))
	ass.True(t, uti.MapsAreEqual(first, second))
	ass.True(t, uti.MapsAreEqual(second, first))
	first["two"] = 5
	ass.False(t, uti.MapsAreEqual(first, second))
	second = uti.CopyMap(first)
	second["four"] = 4
	ass.False(t, uti.MapsAreEqual(first, second))
	fmt.Println(uti.Format(first))
	fmt.Println()
}

type triangle struct {
	X float64
	Y float64
	r float64
}

func TestStructures(t *tes.T) {
	fmt.Println("Structures")
	var structure = triangle{
		X: 3.0,
		Y: 4.0,
		r: 5.0,
	}
	fmt.Println(uti.Format(structure))
	fmt.Println()
}

type Sequential interface {
	AsArray() []string
}

type Array struct {
	attribute []string
}

func (v *Array) AsArray() []string {
	return v.attribute
}

type Association struct {
	key   any
	value any
}

func (v *Association) GetKey() any {
	return v.key
}

func (v *Association) GetValue() any {
	return v.value
}

type Map struct {
	associations []*Association
}

func (v *Map) GetKeys() []any {
	var size = len(v.associations)
	var keys = make([]any, size)
	for index, association := range v.associations {
		keys[index] = association.GetKey()
	}
	return keys
}

func (v *Map) AsArray() []*Association {
	return v.associations
}

func TestPointers(t *tes.T) {
	fmt.Println("Pointers")
	var integer = 5
	fmt.Println(uti.Format(integer))

	var pointer = &integer
	fmt.Println(uti.Format(pointer))

	var double = &pointer
	fmt.Println(uti.Format(double))

	var class = FooBar(2, "bar", nil)
	fmt.Println(uti.Format(class))

	var array Sequential = &Array{
		attribute: []string{
			"foo",
			"bar",
		},
	}
	fmt.Println(uti.Format(array))

	var map_ = &Map{
		associations: []*Association{
			&Association{"alpha", 1},
			&Association{"beta", 2},
			&Association{"gamma", 3},
		},
	}
	fmt.Println(uti.Format(map_))
	fmt.Println()
}

const template = `
	<mixedName>
	<mixedName_>
	<~mixedName>
	<~mixed-name>
	<MixedName>
	<~MixedName>
	<~MIXED_NAME>
`

const expected = `
	mixedValue
	mixedValue
	mixedValue
	mixed-value
	mixedValue
	MixedValue
	MIXED_VALUE
`

const reserved = `
	string
	string_
	string
	string
	string
	String
	STRING
`

func TestStrings(t *tes.T) {
	var mixedCase = "helloWorld"
	var lowerCase = uti.MakeLowerCase(mixedCase)
	ass.Equal(t, "helloWorld", lowerCase)

	var snakeCase = uti.MakeSnakeCase(mixedCase)
	ass.Equal(t, "hello-world", snakeCase)

	var upperCase = uti.MakeUpperCase(mixedCase)
	ass.Equal(t, "HelloWorld", upperCase)

	var allCaps = uti.MakeAllCaps(mixedCase)
	ass.Equal(t, "HELLO_WORLD", allCaps)

	var plural = uti.MakePlural(mixedCase)
	ass.Equal(t, "helloWorlds", plural)

	mixedCase = "HelloWorld"
	lowerCase = uti.MakeLowerCase(mixedCase)
	ass.Equal(t, "helloWorld", lowerCase)

	snakeCase = uti.MakeSnakeCase(mixedCase)
	ass.Equal(t, "hello-world", snakeCase)

	upperCase = uti.MakeUpperCase(mixedCase)
	ass.Equal(t, "HelloWorld", upperCase)

	allCaps = uti.MakeAllCaps(mixedCase)
	ass.Equal(t, "HELLO_WORLD", allCaps)

	plural = uti.MakePlural(mixedCase)
	ass.Equal(t, "HelloWorlds", plural)

	var singular = "mess"
	plural = uti.MakePlural(singular)
	ass.Equal(t, "messes", plural)

	plural = "pixies"
	var unchanged = uti.MakePlural(plural)
	ass.Equal(t, plural, unchanged)

	var actual = uti.ReplaceAll(template, "mixedName", "mixedValue")
	ass.Equal(t, expected, actual)

	actual = uti.ReplaceAll(template, "MixedName", "string")
	ass.Equal(t, reserved, actual)
}

type Interface interface {
	DoNothing()
}

type Class struct{}

func (v *Class) DoNothing() {
}

func TestReflection(t *tes.T) {
	var emptyString string
	ass.True(t, uti.IsUndefined(emptyString))
	ass.False(t, uti.IsDefined(emptyString))

	var hello = "Hello World"
	ass.False(t, uti.IsUndefined(hello))
	ass.True(t, uti.IsDefined(hello))

	var array []int
	ass.True(t, uti.IsUndefined(array))
	ass.False(t, uti.IsDefined(array))

	array = []int{1, 2, 3}
	ass.False(t, uti.IsUndefined(array))
	ass.True(t, uti.IsDefined(array))

	array = array[1:]
	ass.False(t, uti.IsUndefined(array))
	ass.True(t, uti.IsDefined(array))

	var m map[string]int
	ass.True(t, uti.IsUndefined(m))
	ass.False(t, uti.IsDefined(m))

	m = map[string]int{
		"hello": 1,
		"world": 2,
	}
	ass.False(t, uti.IsUndefined(m))
	ass.True(t, uti.IsDefined(m))

	var anything any
	ass.True(t, uti.IsUndefined(anything))
	ass.False(t, uti.IsDefined(anything))

	anything = 5
	ass.False(t, uti.IsUndefined(anything))
	ass.True(t, uti.IsDefined(anything))

	var nilPointer Interface
	ass.True(t, uti.IsUndefined(nilPointer))
	ass.False(t, uti.IsDefined(nilPointer))

	nilPointer = nil
	ass.True(t, uti.IsUndefined(nilPointer))
	ass.False(t, uti.IsDefined(nilPointer))

	ass.True(t, uti.IsUndefined(nil))
	ass.False(t, uti.IsDefined(nil))

	var pointer Interface = &Class{}
	ass.False(t, uti.IsUndefined(pointer))
	ass.True(t, uti.IsDefined(pointer))
	ass.True(t, uti.ImplementsType(pointer, (*Interface)(nil)))
	ass.False(t, uti.ImplementsType(anything, (*Interface)(nil)))
}
