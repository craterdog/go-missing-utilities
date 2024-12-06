/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies.  All Rights Reserved.    .
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
	uti "github.com/craterdog/go-missing-utilities/v2"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

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

func TestStringManipulation(t *tes.T) {
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

type Class struct {
}

func (v *Class) DoNothing() {
}

func TestObjectReflection(t *tes.T) {
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
