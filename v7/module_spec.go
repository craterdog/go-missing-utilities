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

/*
Package "module" defines the global functions provided by this module.  The
functions fill is some gaps in the Go language and native libraries.  They make
it easy to perform the things that should be simple in Go but aren't for various
reasons.  The functions cover the following areas:
  - File System
  - Arrays
  - Maps
  - Strings
  - Reflection
*/
package module

import (
	fmt "fmt"
	cmp "math/cmplx"
	osx "os"
	ref "reflect"
	run "runtime"
	sor "sort"
	stc "strconv"
	sts "strings"
	uni "unicode"
)

// GLOBAL FUNCTIONS

// File System

/*
PathExists checks whether or not the specified file system path is defined.  An
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
file system path.
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
MakeDirectory creates all directories in the specified file system directory
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
specified file system directory path.
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
ReadFile returns the contents of the specified file from the file system as a
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
file in the file system.
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
elements as the specified array.  The result is not a deep copy.
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
ArraysAreEqual[V comparable] determines whether or not the specified arrays have
the same elements.
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

// Maps

/*
CopyMap[K comparable, V any] returns a copy of the specified map with the same
size and key-value pairs as the specified map.  The result is not a deep copy.
*/
func CopyMap[K comparable, V any](
	map_ map[K]V,
) map[K]V {
	var duplicate = make(map[K]V)
	for key, value := range map_ {
		duplicate[key] = value
	}
	return duplicate
}

/*
MapsAreEqual[K comparable, V comparable] determines whether or not the specified
maps have the same key-value pairs.  This function is deterministic even though
Go maps are not.
*/
func MapsAreEqual[K comparable, V comparable](
	first map[K]V,
	second map[K]V,
) bool {
	if len(first) != len(second) {
		return false
	}
	for key, value := range first {
		if second[key] != value {
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
MakePlural attempts to modify the specified mixed case string to make it plural.
It does not use much intelligence to attempt this but gets most cases correct.
*/
func MakePlural(
	mixedCase string,
) string {
	var plural string
	switch {
	case sts.HasSuffix(mixedCase, "ch"):
		plural = mixedCase + "es"
	case sts.HasSuffix(mixedCase, "es"):
		plural = mixedCase
	case sts.HasSuffix(mixedCase, "f"):
		plural = sts.TrimSuffix(mixedCase, "f") + "ves"
	case sts.HasSuffix(mixedCase, "fe"):
		plural = sts.TrimSuffix(mixedCase, "fe") + "ves"
	case sts.HasSuffix(mixedCase, "s"):
		plural = mixedCase + "es"
	case sts.HasSuffix(mixedCase, "sh"):
		plural = mixedCase + "es"
	case sts.HasSuffix(mixedCase, "tz"):
		plural = mixedCase + "es"
	case sts.HasSuffix(mixedCase, "x"):
		plural = mixedCase + "es"
	case sts.HasSuffix(mixedCase, "y") && !sts.HasSuffix(mixedCase, "ey"):
		plural = sts.TrimSuffix(mixedCase, "y") + "ies"
	case sts.HasSuffix(mixedCase, "z") && !sts.HasSuffix(mixedCase, "zz"):
		plural = mixedCase + "zes"
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
  - <lowerCaseName_>   -> lowerCaseValue[_]  {convert value to unique ⃰lower case}
  - <~lowerCaseName>   -> lowerCaseValue     {convert value to lower case}
  - <~snake-case-name> -> snake-case-value   {convert value to snake case}
  - <~UpperCaseName>   -> UpperCaseValue     {convert value to upper case}
  - <~ALL_CAPS_NAME>   -> ALL_CAPS_VALUE     {convert value to all caps with _'s}

⃰A trailing underscore "_" is added if the value collides with a Go keyword.
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

/*
Format returns a canonical string describing any value in Go.  It takes into
account the nesting depth of all compound values (i.e. arrays, maps and structs)
and indents each four spaces per nesting level.  This function does not call the
Go "Stringer" interface on any of the values even if the value supports it since
this the "Stringer" interface does not take into account the nesting depth.

That said, the Go "Stringer" interface can be safely implemented using the
Format function as follows:

	 func (v *MyClass) String() string {
		 return uti.Format(v)
	 }

There should be no risk of infinite recursion from Format() calling String()
calling Format() calling String()...
*/
func Format(
	value any,
) string {
	var reflected = ref.ValueOf(value)
	return formatValue(reflected, 0)
}

// Reflection

/*
ImplementsInterface checks whether or not the specified value implements the
specified interface.  It can be used as follows:

	type MyInterface interface {
		DoSomething()
	}

	type MyStruct struct{}

	func (v *MyStruct) DoSomething() {}

	func main() {
		var myValue any = &MyStruct{}
		var myInterface *MyInterface
		if ImplementsInterface(myValue, myInterface) {
			var actual MyInterface = myValue.(MyInterface)
			fmt.Println("myValue implements MyInterface:", actual)
		}
	}

NOTE: The interface argument that gets passed into the ImplementsInterface() call
must be a pointer to the interface since the argument is of type any.
*/
func ImplementsInterface(
	value any,
	pointer any,
) bool {
	if IsUndefined(value) {
		return false
	}
	var valueType = ref.TypeOf(value)
	var interface_ = ref.TypeOf(pointer).Elem()
	return valueType.Implements(interface_)
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

// Private

const maximumDepth = 8

func formatArray(
	reflected ref.Value,
	depth uint,
) string {
	var result = "["
	var size = reflected.Len()
	if size == 0 {
		// This is an empty array.
		result += " "
	} else {
		// This is a multivalued array.
		if depth < maximumDepth {
			depth++
			for index := 0; index < size; index++ {
				result += formatNewline(depth)
				var value = reflected.Index(index)
				result += formatValue(value, depth)
			}
			depth--
			result += formatNewline(depth)
		} else {
			result += "..."
		}
	}
	var typeName = formatType(reflected.Type())
	result += "](" + typeName + ")"
	return result
}

func formatAssociation(
	key ref.Value,
	value ref.Value,
	depth uint,
) string {
	var result = formatValue(key, depth)
	result += ": "
	result += formatValue(value, depth)
	return result
}

func formatAssociations(
	reflected ref.Value,
	depth uint,
) string {
	var result string
	var size = reflected.Len()
	if size == 0 {
		// This is an empty sequence of associations.
		result += ":"
	} else {
		// This is a multivalued sequence of associations.
		if depth < maximumDepth {
			depth++
			for index := 0; index < size; index++ {
				result += formatNewline(depth)
				var association = reflected.Index(index)
				var key = association.MethodByName("GetKey").Call(
					[]ref.Value{},
				)[0]
				var value = association.MethodByName("GetValue").Call(
					[]ref.Value{},
				)[0]
				result += formatAssociation(key, value, depth)
			}
			depth--
			result += formatNewline(depth)
		} else {
			result += "..."
		}
	}
	return result
}

func formatBoolean(
	reflected ref.Value,
	depth uint,
) string {
	var value = reflected.Bool()
	return stc.FormatBool(value)
}

func formatChannel(
	reflected ref.Value,
	depth uint,
) string {
	var direction string
	var reflectedType = reflected.Type()
	switch reflectedType.ChanDir() {
	case ref.RecvDir:
		direction = "Receive"
	case ref.SendDir:
		direction = "Send"
	case ref.BothDir:
		direction = "Both"
	}
	var result = "["
	depth++
	result += formatNewline(depth)
	result += "Direction: " + direction
	result += formatNewline(depth)
	result += "Capacity: " + stc.Itoa(reflected.Cap())
	result += formatNewline(depth)
	result += "Size: " + stc.Itoa(reflected.Len())
	depth--
	result += formatNewline(depth)
	var typeName = formatType(reflected.Type())
	result += "](" + typeName + ")"
	return result
}

func formatComplex(
	reflected ref.Value,
	depth uint,
) string {
	var value = reflected.Complex()
	return stc.FormatComplex(complex128(value), 'G', -1, 64)
}

func formatFloat(
	reflected ref.Value,
	depth uint,
) string {
	var value = reflected.Float()
	var result = stc.FormatFloat(float64(value), 'G', -1, 64)
	if !sts.Contains(result, ".") && !sts.Contains(result, "E") {
		result += ".0"
	}
	return result
}

func formatFunction(
	reflected ref.Value,
	depth uint,
) string {
	// Format the signature type rather than the function definition.
	var functionName = run.FuncForPC(reflected.Pointer()).Name()
	var functionSignature = formatType(reflected.Type())
	if IsDefined(functionName) {
		var sections = sts.Split(functionName, ".")
		functionName = sections[len(sections)-1]
		functionSignature = sts.TrimPrefix(functionSignature, "func")
		functionSignature = "func " + functionName + functionSignature
	}
	return functionSignature
}

func formatInstance(
	reflected ref.Value,
	depth uint,
) string {
	var result string
	if depth < maximumDepth {
		depth++
		var reflectedType = reflected.Type()
		var count = reflectedType.NumMethod()
		for index := 0; index < count; index++ {
			var methodName = reflectedType.Method(index).Name
			if sts.HasPrefix(methodName, "Get") {
				var method = reflected.MethodByName(methodName)
				var methodType = method.Type()
				if methodType.NumIn() == 0 && methodType.NumOut() == 1 {
					result += formatNewline(depth)
					var attributeName = sts.TrimPrefix(methodName, "Get")
					var attributeValue = method.Call(
						[]ref.Value{},
					)[0]
					result += attributeName
					result += ": "
					if methodName == "GetClass" {
						// Just format the class type to avoid any recursion.
						var classType = methodType.Out(0)
						result += formatType(classType)
					} else {
						result += formatValue(attributeValue, depth)
					}
				}
			}
		}
		depth--
		result += formatNewline(depth)
	} else {
		result += "..."
	}
	return result
}

func formatInteger(
	reflected ref.Value,
	depth uint,
) string {
	var value = reflected.Int()
	return stc.FormatInt(int64(value), 10)
}

func formatInterface(
	reflected ref.Value,
	depth uint,
) string {
	// NOTE:
	// Since a class that implements an iterface must implement all methods
	// defined in that interface we can just format the value behind the
	// interface.
	var value = reflected.Elem()
	var result = formatValue(value, depth)
	return result
}

var typeMap = map[ref.Kind]uint8{
	ref.Bool:          0,
	ref.Uint8:         1,
	ref.Uint16:        2,
	ref.Uint32:        3,
	ref.Uint64:        4,
	ref.Uint:          5,
	ref.Int8:          6,
	ref.Int16:         7,
	ref.Int64:         8,
	ref.Int:           9,
	ref.Float32:       10,
	ref.Float64:       11,
	ref.Complex64:     12,
	ref.Complex128:    13,
	ref.Int32:         14,
	ref.String:        15,
	ref.Func:          16,
	ref.Chan:          17,
	ref.Array:         18,
	ref.Slice:         19,
	ref.Map:           20,
	ref.Struct:        21,
	ref.Pointer:       22,
	ref.Interface:     23,
	ref.Uintptr:       24,
	ref.UnsafePointer: 25,
}

func formatMap(
	reflected ref.Value,
	depth uint,
) string {
	// NOTE:
	// The intrinsic Go map data type is non-deterministic.  The ordering of the
	// keys is determined by a hash function which means that two maps with the
	// same keys will likely return the keys in a different order.  This also
	// means that the same code will likely run differently each time it is
	// executed.  It is important—for testing and debugging purposes—that the
	// formatting functionality be deterministic, even for Go maps.  This
	// private function attempts to ensure determinism.  The keys are sorted
	// before formatting with the following sorting criteria:
	//
	// Key type ordering (see the typeMap data structure above):
	//  * booleans
	//  * unsigned integers
	//  * signed integers
	//  * floats
	//  * complex numbers
	//  * runes
	//  * strings
	//
	// Value ordering:
	//  * false before true
	//  * complex values by their amplitudes
	//  * runes by their unicode numbers
	//  * strings alphabetically by the unicode number of their characters
	//
	var result = "["
	var size = reflected.Len()
	if size == 0 {
		// This is an empty map.
		result += ":"
	} else {
		// This is a multivalued map.
		if depth < maximumDepth {
			depth++
			// First sort the keys since Go maps are deterministic.
			var keys = reflected.MapKeys()
			sor.SliceStable(
				keys,
				func(i, j int) bool {
					// Convert wrapper types into their element types.
					var firstKey = keys[i]
					var secondKey = keys[j]
					if firstKey.Kind() == ref.Interface {
						firstKey = firstKey.Elem()
					}
					if secondKey.Kind() == ref.Interface {
						secondKey = secondKey.Elem()
					}
					// Sort by key type if the keys have different types.
					if firstKey.Kind() != secondKey.Kind() {
						var firstType = typeMap[firstKey.Kind()]
						var secondType = typeMap[secondKey.Kind()]
						return firstType < secondType
					}
					// Sort by key value if they have the same type.
					switch firstKey.Kind() {
					case ref.Bool:
						return !(firstKey.Bool()) && secondKey.Bool()
					case ref.Int, ref.Int8, ref.Int16, ref.Int32, ref.Int64:
						return firstKey.Int() < secondKey.Int()
					case ref.Uint, ref.Uint8, ref.Uint16, ref.Uint32, ref.Uint64:
						return firstKey.Uint() < secondKey.Uint()
					case ref.Float32, ref.Float64:
						return firstKey.Float() < secondKey.Float()
					case ref.Complex64, ref.Complex128:
						var firstAmplitude = cmp.Abs(firstKey.Complex())
						var secondAmplitude = cmp.Abs(secondKey.Complex())
						return firstAmplitude < secondAmplitude
					case ref.String:
						return firstKey.String() < secondKey.String()
					default:
						var message = fmt.Sprintf(
							"Attempted to compare an unknown key type: %v of type %T",
							firstKey.Interface(),
							firstKey.Interface(),
						)
						panic(message)
					}
				},
			)
			// Format the key-value pairs in order.
			for _, key := range keys {
				result += formatNewline(depth)
				var value = reflected.MapIndex(key)
				result += formatAssociation(key, value, depth)
			}
			depth--
			result += formatNewline(depth)
		} else {
			result += "..."
		}
	}
	var typeName = formatType(reflected.Type())
	result += "](" + typeName + ")"
	return result
}

func formatNewline(
	depth uint,
) string {
	var result = "\n"
	var indentation = "    "
	var level uint
	for level < depth {
		result += indentation
		level++
	}
	return result
}

func formatPointer(
	reflected ref.Value,
	depth uint,
) string {
	var result = "&["
	switch {
	case reflected.MethodByName("GetKeys").IsValid():
		// Format the sequence of associations.
		var associations = reflected.MethodByName("AsArray").Call(
			[]ref.Value{},
		)[0]
		result += formatAssociations(associations, depth)
	case reflected.MethodByName("AsArray").IsValid():
		// Format the sequence of values.
		var values = reflected.MethodByName("AsArray").Call(
			[]ref.Value{},
		)[0]
		result += formatSequence(values, depth)
	case reflected.NumMethod() > 0:
		// Format the instance of a class.
		result += formatInstance(reflected, depth)
	default:
		// Dereference the pointer.
		var value = reflected.Elem()
		result += formatValue(value, depth)
	}
	var typeName = formatType(reflected.Type())
	result += "](" + typeName + ")"
	return result
}

func formatRune(
	reflected ref.Value,
	depth uint,
) string {
	var value = rune(reflected.Int())
	return stc.QuoteRune(value)
}

func formatSequence(
	reflected ref.Value,
	depth uint,
) string {
	var result string
	var size = reflected.Len()
	if size == 0 {
		// This is an empty sequence.
		result += " "
	} else {
		// This is a multivalued sequence.
		if depth < maximumDepth {
			depth++
			for index := 0; index < size; index++ {
				result += formatNewline(depth)
				var value = reflected.Index(index)
				result += formatValue(value, depth)
			}
			depth--
			result += formatNewline(depth)
		} else {
			result += "..."
		}
	}
	return result
}

func formatString(
	reflected ref.Value,
	depth uint,
) string {
	var value = reflected.String()
	return stc.Quote(value)
}

func formatStructure(
	reflected ref.Value,
	depth uint,
) string {
	var result = "["
	if depth < maximumDepth {
		depth++
		var fields = ref.VisibleFields(reflected.Type())
		for index, field := range fields {
			result += formatNewline(depth)
			var name = field.Name
			result += name
			result += ": "
			if field.IsExported() {
				var value = reflected.Field(index)
				result += formatValue(value, depth)
			} else {
				result += "<private>"
			}
		}
		depth--
		result += formatNewline(depth)
	} else {
		result += "..."
	}
	var typeName = formatType(reflected.Type())
	result += "](" + typeName + ")"
	return result
}

func formatType(
	reflectedType ref.Type,
) string {
	var result string
	switch reflectedType.Kind() {
	case ref.Interface:
		// Interfaces may be named or the "any" interface type.
		result = reflectedType.Name()
		if len(result) == 0 {
			result = reflectedType.String()
		}
		if result == "interface {}" {
			result = "any"
		}
	case ref.Pointer:
		// Format the pointer type recursively.
		result = "*"
		var referencedType = reflectedType.Elem()
		result += formatType(referencedType)
	case ref.UnsafePointer:
		// Mark it as unsafe.
		result = "<unsafe>"
	case ref.Func:
		// Format the function signature type recursively.
		result = "func("
		var count = reflectedType.NumIn()
		for index := 0; index < count; index++ {
			var argumentType = reflectedType.In(index)
			result += formatType(argumentType)
			if index < count-1 {
				result += ", "
			}
		}
		result += ")"
		count = reflectedType.NumOut()
		if count > 0 {
			result += " ("
			for index := 0; index < count; index++ {
				var argumentType = reflectedType.Out(index)
				result += formatType(argumentType)
				if index < count-1 {
					result += ", "
				}
			}
			result += ")"
		}
	case ref.Struct:
		// Format the structure type.
		result = reflectedType.Name()
		if len(result) == 0 {
			result = "struct"
		}
	case ref.Chan:
		// Format the channel type recursively.
		result = "chan "
		var elementType = reflectedType.Elem()
		result += formatType(elementType)
	case ref.Array, ref.Slice:
		// Format the array type recursively.
		result = "array["
		var elementType = reflectedType.Elem()
		result += formatType(elementType)
		result += "]"
	case ref.Map:
		// Format the map type recursively.
		result = "map["
		var keyType = reflectedType.Key()
		result += formatType(keyType)
		result += ", "
		var valueType = reflectedType.Elem()
		result += formatType(valueType)
		result += "]"
	default:
		// Everything else should be a named type or an intrinsic type.
		result = reflectedType.Name()
		if len(result) == 0 {
			result = reflectedType.String()
		}
	}
	return result
}

func formatUnsafe(
	reflected ref.Value,
	depth uint,
) string {
	return "<unsafe>"
}

func formatUnsigned(
	reflected ref.Value,
	depth uint,
) string {
	var value = reflected.Uint()
	return "0x" + stc.FormatUint(uint64(value), 16)
}

func formatValue(
	reflected ref.Value,
	depth uint,
) string {
	if !reflected.IsValid() {
		return "<nil>"
	}
	switch reflected.Kind() {
	case ref.Bool:
		return formatBoolean(reflected, depth)

	case ref.Uint, ref.Uint8, ref.Uint16, ref.Uint32, ref.Uint64, ref.Uintptr:
		return formatUnsigned(reflected, depth)

	case ref.Int, ref.Int8, ref.Int16, ref.Int64:
		return formatInteger(reflected, depth)

	case ref.Float32, ref.Float64:
		return formatFloat(reflected, depth)

	case ref.Complex64, ref.Complex128:
		return formatComplex(reflected, depth)

	case ref.Int32:
		return formatRune(reflected, depth)

	case ref.String:
		return formatString(reflected, depth)

	case ref.Func:
		return formatFunction(reflected, depth)

	case ref.Chan:
		return formatChannel(reflected, depth)

	case ref.Array, ref.Slice:
		return formatArray(reflected, depth)

	case ref.Map:
		return formatMap(reflected, depth)

	case ref.Struct:
		return formatStructure(reflected, depth)

	case ref.Pointer:
		return formatPointer(reflected, depth)

	case ref.Interface:
		return formatInterface(reflected, depth)

	case ref.UnsafePointer:
		return formatUnsafe(reflected, depth)

	default:
		var message = fmt.Sprintf(
			"Attempted to format an unsupported type: %v",
			reflected.Kind(),
		)
		panic(message)
	}
}
