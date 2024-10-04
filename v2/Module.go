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
	ref "reflect"
	sts "strings"
	uni "unicode"
)

// GLOBAL FUNCTIONS

// String Manipulation

/*
MakeAllCaps modifies the specified mixed case string into a corresponding all
uppercase string using "_"s to separate the words found in the mixed case
string.
*/
func MakeAllCaps(mixedCase string) string {
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
func MakeLowerCase(mixedCase string) string {
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
func MakePlural(mixedCase string) string {
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
func MakeSnakeCase(mixedCase string) string {
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
func MakeUpperCase(mixedCase string) string {
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
  - <anyCaseName>     -> value              {leave value as is}
  - <lowerCaseName_>  -> lowerCaseValue[_]  {convert value to lower case and ensure uniqueness}
  - <snake-case-name> -> snake-case-value   {convert value to snake case}
  - <UpperCaseName_>  -> UpperCaseValue     {convert value to upper case}
  - <ALL_CAPS_NAME>   -> ALL_CAPS_VALUE     {convert value to all caps with underscores}
*/
func ReplaceAll(template string, name string, value string) string {
	// <anyCaseName> -> value
	var anyCaseName = MakeLowerCase(name)
	template = sts.ReplaceAll(template, "<"+anyCaseName+">", value)
	anyCaseName = MakeUpperCase(name)
	template = sts.ReplaceAll(template, "<"+anyCaseName+">", value)

	// <lowerCaseName_> -> lowerCaseValue[_]
	var lowerCaseName = MakeLowerCase(name) + "_"
	var lowerCaseValue = MakeLowerCase(value)
	switch lowerCaseValue {
	// Check to see if the value is a reserved word.
	case "any", "byte", "case", "complex", "copy", "default", "error",
		"false", "import", "interface", "map", "nil", "package", "range",
		"real", "return", "rune", "string", "switch", "true", "type":
		lowerCaseValue += "_"
	}
	template = sts.ReplaceAll(template, "<"+lowerCaseName+">", lowerCaseValue)

	// <snake-case-name> -> snake-case-value
	var snakeCaseName = MakeSnakeCase(name)
	var snakeCaseValue = MakeSnakeCase(value)
	template = sts.ReplaceAll(template, "<"+snakeCaseName+">", snakeCaseValue)

	// <UpperCaseName_> -> UpperCaseValue
	var upperCaseName = MakeUpperCase(name) + "_"
	var upperCaseValue = MakeUpperCase(value)
	template = sts.ReplaceAll(template, "<"+upperCaseName+">", upperCaseValue)

	// <ALL_CAPS_NAME> -> ALL_CAPS_VALUE
	var allCapsName = MakeAllCaps(name)
	var allCapsValue = MakeAllCaps(value)
	template = sts.ReplaceAll(template, "<"+allCapsName+">", allCapsValue)

	return template
}

// Object Reflection

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
	switch actual := value.(type) {
	case string:
		return len(actual) > 0
	default:
		var meta = ref.ValueOf(actual)
		var isPointer = meta.Kind() == ref.Ptr ||
			meta.Kind() == ref.Interface ||
			meta.Kind() == ref.Slice ||
			meta.Kind() == ref.Map ||
			meta.Kind() == ref.Chan ||
			meta.Kind() == ref.Func
		var isNil = isPointer && meta.IsNil()
		return !isNil && meta.IsValid()
	}
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
