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
	mod "github.com/craterdog/go-missing-utilities/v2"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const template = `
	<mixedName>
	<MixedName>
	<mixedName_>
	<mixed-name>
	<MixedName_>
	<MIXED_NAME>
`

const expected = `
	mixedValue
	mixedValue
	mixedValue
	mixed-value
	MixedValue
	MIXED_VALUE
`

const reserved = `
	string
	string
	string_
	string
	String
	STRING
`

func TestStringManipulation(t *tes.T) {
	var mixedCase = "helloWorld"
	var lowerCase = mod.MakeLowerCase(mixedCase)
	ass.Equal(t, "helloWorld", lowerCase)

	var snakeCase = mod.MakeSnakeCase(mixedCase)
	ass.Equal(t, "hello-world", snakeCase)

	var upperCase = mod.MakeUpperCase(mixedCase)
	ass.Equal(t, "HelloWorld", upperCase)

	var allCaps = mod.MakeAllCaps(mixedCase)
	ass.Equal(t, "HELLO_WORLD", allCaps)

	var plural = mod.MakePlural(mixedCase)
	ass.Equal(t, "helloWorlds", plural)

	mixedCase = "HelloWorld"
	lowerCase = mod.MakeLowerCase(mixedCase)
	ass.Equal(t, "helloWorld", lowerCase)

	snakeCase = mod.MakeSnakeCase(mixedCase)
	ass.Equal(t, "hello-world", snakeCase)

	upperCase = mod.MakeUpperCase(mixedCase)
	ass.Equal(t, "HelloWorld", upperCase)

	allCaps = mod.MakeAllCaps(mixedCase)
	ass.Equal(t, "HELLO_WORLD", allCaps)

	plural = mod.MakePlural(mixedCase)
	ass.Equal(t, "HelloWorlds", plural)

	var singular = "mess"
	plural = mod.MakePlural(singular)
	ass.Equal(t, "messes", plural)

	plural = "pixies"
	var unchanged = mod.MakePlural(plural)
	ass.Equal(t, plural, unchanged)

	var actual = mod.ReplaceAll(template, "mixedName", "mixedValue")
	ass.Equal(t, expected, actual)

	actual = mod.ReplaceAll(template, "MixedName", "string")
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
	ass.True(t, mod.IsUndefined(emptyString))
	ass.False(t, mod.IsDefined(emptyString))

	var hello = "Hello World"
	ass.False(t, mod.IsUndefined(hello))
	ass.True(t, mod.IsDefined(hello))

	var array []int
	ass.True(t, mod.IsUndefined(array))
	ass.False(t, mod.IsDefined(array))

	array = []int{1, 2, 3}
	ass.False(t, mod.IsUndefined(array))
	ass.True(t, mod.IsDefined(array))

	array = array[1:]
	ass.False(t, mod.IsUndefined(array))
	ass.True(t, mod.IsDefined(array))

	var m map[string]int
	ass.True(t, mod.IsUndefined(m))
	ass.False(t, mod.IsDefined(m))

	m = map[string]int{
		"hello": 1,
		"world": 2,
	}
	ass.False(t, mod.IsUndefined(m))
	ass.True(t, mod.IsDefined(m))

	var anything any
	ass.True(t, mod.IsUndefined(anything))
	ass.False(t, mod.IsDefined(anything))

	anything = 5
	ass.False(t, mod.IsUndefined(anything))
	ass.True(t, mod.IsDefined(anything))

	var nilPointer Interface
	ass.True(t, mod.IsUndefined(nilPointer))
	ass.False(t, mod.IsDefined(nilPointer))

	nilPointer = nil
	ass.True(t, mod.IsUndefined(nilPointer))
	ass.False(t, mod.IsDefined(nilPointer))

	var pointer Interface = &Class{}
	ass.False(t, mod.IsUndefined(pointer))
	ass.True(t, mod.IsDefined(pointer))
	ass.True(t, mod.ImplementsType(pointer, (*Interface)(nil)))
	ass.False(t, mod.ImplementsType(anything, (*Interface)(nil)))
}
