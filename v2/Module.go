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

/*
Package "module" defines the global functions provided by this module.
*/
package module

import (
	osx "os"
	ref "reflect"
	sts "strings"
	uni "unicode"
)

// GLOBAL FUNCTIONS

// Filesystem

/*
PathExists checks whether or not the specified filesystem path is defined.  An
empty string or a nil pointer is considered to be undefined.
*/
func PathExists(
	path string,
) bool {
	var _, err = osx.Stat(path)
	if err == nil {
		return true
	}
	if osx.IsNotExist(err) {
		return false
	}
	panic(err)
}

/*
RemovePath recursively removes all directories and files found in the specified
filesystem path.
*/
func RemovePath(
	path string,
) {
	var err = osx.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

/*
MakeDirectory creates all directories in the specified filesystem directory
path.
*/
func MakeDirectory(
	directory string,
) {
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

/*
RemakeDirectory recursively removes all files and subdirectories from the
specified filesystem directory path.
*/
func RemakeDirectory(
	directory string,
) {
	var err = osx.RemoveAll(directory)
	if err != nil {
		panic(err)
	}
	err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

/*
ReadFile returns the contents of the specified file from the filesystem as a
string.
*/
func ReadFile(
	filename string,
) string {
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	return source
}

/*
WriteFile writes the specified source string as the contents of the specified
file in the filesystem.
*/
func WriteFile(
	filename string,
	source string,
) {
	var bytes = []byte(source)
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}

// Arrays

/*
CopyArray[V any] returns a copy of the specified array with the same size and
elements as the specified array.
*/
func CopyArray[V any](
	array []V,
) []V {
	var size = len(array)
	var duplicate = make([]V, size)
	copy(duplicate, array)
	return duplicate
}

/*
ArraysAreEqual[V comparable] determines whether or not the specified arrays have the
same elements.
*/
func ArraysAreEqual[V comparable](
	first []V,
	second []V,
) bool {
	if len(first) != len(second) {
		return false
	}
	for index, value := range first {
		if value != second[index] {
			return false
		}
	}
	return true
}

// Strings

/*
MakeAllCaps modifies the specified mixed case string into a corresponding all
uppercase string using "_"s to separate the words found in the mixed case
string.
*/
func MakeAllCaps(
	mixedCase string,
) string {
	var allCaps sts.Builder
	var foundLower bool
	for _, r := range mixedCase {
		switch {
		case uni.IsLower(r):
			foundLower = true
			allCaps.WriteRune(uni.ToUpper(r))
		case uni.IsUpper(r):
			if foundLower {
				allCaps.WriteString("_")
				foundLower = false
			}
			allCaps.WriteRune(r)
		default:
			allCaps.WriteRune(r)
		}
	}
	return allCaps.String()
}

/*
MakeLowerCase modifies the specified mixed case string into a corresponding
string starting with a lowercase letter.  All other letters remain unchanged.
*/
func MakeLowerCase(
	mixedCase string,
) string {
	var lowerCase string
	if len(mixedCase) > 0 {
		runes := []rune(mixedCase)
		runes[0] = uni.ToLower(runes[0])
		lowerCase = string(runes)
	}
	return lowerCase
}

/*
MakePlural attempts to modify the specified mixed case string to make it
plural.  It does not use much intelligence to attempt this.
*/
func MakePlural(
	mixedCase string,
) string {
	var plural string
	switch {
	case sts.HasSuffix(mixedCase, "es"):
		plural = mixedCase
	case sts.HasSuffix(mixedCase, "s"):
		plural = mixedCase + "es"
	default:
		plural = mixedCase + "s"
	}
	return plural
}

/*
MakeSnakeCase modifies the specified mixed case string into a corresponding all
lowercase string using "-"s to separate the words found in the mixed case
string.
*/
func MakeSnakeCase(
	mixedCase string,
) string {
	mixedCase = MakeLowerCase(mixedCase)
	var snakeCase sts.Builder
	for _, r := range mixedCase {
		switch {
		case uni.IsLower(r):
			snakeCase.WriteRune(r)
		case uni.IsUpper(r):
			snakeCase.WriteString("-")
			snakeCase.WriteRune(uni.ToLower(r))
		default:
			snakeCase.WriteRune(r)
		}
	}
	return snakeCase.String()
}

/*
MakeUpperCase modifies the specified mixed case string into a corresponding
string starting with an uppercase letter.  All other letters remain unchanged.
*/
func MakeUpperCase(
	mixedCase string,
) string {
	var upperCase string
	if len(mixedCase) > 0 {
		runes := []rune(mixedCase)
		runes[0] = uni.ToUpper(runes[0])
		upperCase = string(runes)
	}
	return upperCase
}

/*
ReplaceAll replaces each instance of the specified name embedded in angle
brackets (i.e. "<" and ">") with the specified value throughout the specified
template string.  The way the name is shown in the brackets determines what
transformations are done on the value prior to the substitution as follows:
  - <anyCaseName>      -> value              {leave value as is}
  - <lowerCaseName_>   -> lowerCaseValue[_]  {convert value to [unambiguous] lower case}
  - <~lowerCaseName>   -> lowerCaseValue     {convert value to lower case}
  - <~snake-case-name> -> snake-case-value   {convert value to snake case}
  - <~UpperCaseName>   -> UpperCaseValue     {convert value to upper case}
  - <~ALL_CAPS_NAME>   -> ALL_CAPS_VALUE     {convert value to all caps with underscores}
*/
func ReplaceAll(
	template string,
	name string,
	value string,
) string {
	// <anyCaseName> -> value
	var anyCaseName = MakeLowerCase(name)
	template = sts.ReplaceAll(template, "<"+anyCaseName+">", value)
	anyCaseName = MakeUpperCase(name)
	template = sts.ReplaceAll(template, "<"+anyCaseName+">", value)

	// <lowerCaseName_> -> lowerCaseValue[_]
	var lowerCaseName = MakeLowerCase(name) + "_"
	var lowerCaseValue = MakeLowerCase(value)
	switch lowerCaseValue {
	// Check to see if the value is a Go reserved word.
	case "any", "append", "bool", "break", "byte", "cap", "case",
		"chan", "clear", "close", "comparable", "complex", "const",
		"continue", "copy", "default", "defer", "delete", "else",
		"error", "fallthrough", "false", "for", "func", "go", "goto",
		"if", "imag", "import", "int", "interface", "iota", "len",
		"make", "map", "max", "min", "new", "nil", "package", "panic",
		"print", "println", "range", "real", "recover", "return",
		"rune", "select", "string", "struct", "switch", "true", "type",
		"uint", "uintptr", "var":
		lowerCaseValue += "_"
	}
	template = sts.ReplaceAll(template, "<"+lowerCaseName+">", lowerCaseValue)

	// <~lowerCaseName> -> lowerCaseValue
	lowerCaseName = "~" + MakeLowerCase(name)
	lowerCaseValue = MakeLowerCase(value)
	template = sts.ReplaceAll(template, "<"+lowerCaseName+">", lowerCaseValue)

	// <~snake-case-name> -> snake-case-value
	var snakeCaseName = "~" + MakeSnakeCase(name)
	var snakeCaseValue = MakeSnakeCase(value)
	template = sts.ReplaceAll(template, "<"+snakeCaseName+">", snakeCaseValue)

	// <~UpperCaseName> -> UpperCaseValue
	var upperCaseName = "~" + MakeUpperCase(name)
	var upperCaseValue = MakeUpperCase(value)
	template = sts.ReplaceAll(template, "<"+upperCaseName+">", upperCaseValue)

	// <~ALL_CAPS_NAME> -> ALL_CAPS_VALUE
	var allCapsName = "~" + MakeAllCaps(name)
	var allCapsValue = MakeAllCaps(value)
	template = sts.ReplaceAll(template, "<"+allCapsName+">", allCapsValue)

	return template
}

// Reflection

/*
ImplementsType checks whether or not the specified value implements the expected
type of the second specified value.
*/
func ImplementsType(
	value any,
	expectedType any,
) bool {
	if IsDefined(value) {
		var actualType = ref.TypeOf(value)
		var reflectedType = ref.TypeOf(expectedType).Elem()
		return actualType.Implements(reflectedType)
	}
	return false
}

/*
IsDefined checks whether or not the specified value is defined in a meaningful
way.  Empty strings and nil pointers are considered as being undefined.
*/
func IsDefined(
	value any,
) bool {
	// This method addresses the inconsistencies in the Go language with respect
	// to whether or not a value is defined or not.  Go handles interfaces,
	// pointers and various intrinsic types differently.  This makes consistent
	// checking across different types problematic.  We handle it here in one
	// place (hopefully correctly).
	var isDefined bool
	if value != nil {
		switch actual := value.(type) {
		case string:
			return len(actual) > 0
		default:
			var isPointer bool
			var meta = ref.ValueOf(actual)
			switch meta.Kind() {
			case ref.Ptr, ref.Interface, ref.Slice, ref.Map, ref.Chan, ref.Func:
				isPointer = true
			}
			var isNil = isPointer && meta.IsNil()
			isDefined = !isNil && meta.IsValid()
		}
	}
	return isDefined
}

/*
IsUndefined checks whether or not the specified value is undefined.  Empty
strings and nil pointers are considered as being undefined.
*/
func IsUndefined(
	value any,
) bool {
	return !IsDefined(value)
}
