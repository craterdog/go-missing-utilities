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
  - Composites (arrays, slices and maps)
  - Strings
  - Reflection
*/
package module

// GLOBAL TYPES

// GLOBAL FUNCTIONS

// File System

/*
PathExists checks whether or not the specified file system path is defined.  An
empty string or a nil pointer is considered to be undefined.
*/
func PathExists(
	path string,
) bool {
	return pathExists(path)
}

/*
RemovePath recursively removes all directories and files found in the specified
file system path.
*/
func RemovePath(
	path string,
) {
	removePath(path)
}

/*
HomeDirectory returns the home directory path for the current user.
*/
func HomeDirectory() string {
	return homeDirectory()
}

/*
MakeDirectory creates all directories in the specified file system directory
path.
*/
func MakeDirectory(
	directory string,
) {
	makeDirectory(directory)
}

/*
RemakeDirectory recursively removes all files and subdirectories from the
specified file system directory path.
*/
func RemakeDirectory(
	directory string,
) {
	remakeDirectory(directory)
}

/*
ReadDirectory returns an array containing the filenames of the files in the
specified directory.
*/
func ReadDirectory(
	directory string,
) []string {
	return readDirectory(directory)
}

/*
ReadFile returns the contents of the specified file from the file system as a
string.
*/
func ReadFile(
	filename string,
) string {
	return readFile(filename)
}

/*
WriteFile writes the specified source string as the contents of the specified
file in the file system.
*/
func WriteFile(
	filename string,
	source string,
) {
	writeFile(filename, source)
}

// Composites

/*
Relative indexing allows an index to be a relative positive (or negative)
ordinal index of a value in a sequence.  The indices are ordinal rather than
cardinal (zero-based) which never really made sense except for pointer offsets.
What is the "zeroth value" in a sequence anyway?  It's the "first value",
right?  So we start a fresh...

The relative indexing approach allows for positive indices starting at the
beginning of a sequence—and negative indices starting at the end of the
sequence, as follows:

	    1           2           3             N
	[value 1] . [value 2] . [value 3] ... [value N]
	   -N        -(N-1)      -(N-2)          -1

Notice that because the indices are ordinal based, the positive and negative
indices are symmetrical.  A relative index can NEVER be zero.

RelativeToCardinal transforms a relative (ordinal-based) index into the
corresponding zero-based index.  The following transformation is performed:

	[-size..-1] or [1..size] => [0..size)

Notice that the specified relative index cannot be zero since zero is NOT an
ordinal number.
*/
func RelativeToCardinal(
	relative int,
	size uint,
) int {
	return relativeToCardinal(relative, size)
}

/*
CardinalToRelative transforms a cardinal (zero-based) index into the
corresponding relative (ordinal-based) index.  The following transformation
is performed:

	[0..size) => [1..size]

The transformation always chooses the positive ordinal range.
*/

func CardinalToRelative(
	cardinal int,
	size uint,
) int {
	return cardinalToRelative(cardinal, size)
}

/*
CopyArray[V any] returns a copy of the specified array with the same size and
elements as the specified array.  The result is not a deep copy.
*/
func CopyArray[V any](
	array []V,
) []V {
	return copyArray(array)
}

/*
ArraysAreEqual[V comparable] determines whether or not the specified arrays have
the same elements.
*/
func ArraysAreEqual[V comparable](
	first []V,
	second []V,
) bool {
	return arraysAreEqual(first, second)
}

/*
CombineArrays[V any] returns a new array containing the concatenation
of the specified arrays.
*/
func CombineArrays[V any](
	first []V,
	second []V,
) []V {
	return combineArrays(first, second)
}

/*
CopyMap[K comparable, V any] returns a copy of the specified map with the same
size and key-value pairs as the specified map.  The result is not a deep copy.
*/
func CopyMap[K comparable, V any](
	map_ map[K]V,
) map[K]V {
	return copyMap(map_)
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
	return mapsAreEqual(first, second)
}

/*
CombineMaps[K comparable, V any] returns a new map containing the concatenation
of the specified maps.
*/
func CombineMaps[K comparable, V any](
	first map[K]V,
	second map[K]V,
) map[K]V {
	return combineMaps(first, second)
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
	return makeAllCaps(mixedCase)
}

/*
MakeLowerCase modifies the specified mixed case string into a corresponding
string starting with a lowercase letter.  All other letters remain unchanged.
*/
func MakeLowerCase(
	mixedCase string,
) string {
	return makeLowerCase(mixedCase)
}

/*
MakePlural attempts to modify the specified mixed case string to make it plural.
It does not use much intelligence to attempt this but gets most cases correct.
*/
func MakePlural(
	mixedCase string,
) string {
	return makePlural(mixedCase)
}

/*
MakeSnakeCase modifies the specified mixed case string into a corresponding all
lowercase string using "-"s to separate the words found in the mixed case
string.
*/
func MakeSnakeCase(
	mixedCase string,
) string {
	return makeSnakeCase(mixedCase)
}

/*
MakeUpperCase modifies the specified mixed case string into a corresponding
string starting with an uppercase letter.  All other letters remain unchanged.
*/
func MakeUpperCase(
	mixedCase string,
) string {
	return makeUpperCase(mixedCase)
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
	return replaceAll(template, name, value)
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
	return format(value)
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
	return implementsInterface(value, pointer)
}

/*
IsDefined checks whether or not the specified value is defined in a meaningful
way.  Empty strings and nil pointers are considered as being undefined.
*/
func IsDefined(
	value any,
) bool {
	return isDefined(value)
}

/*
IsUndefined checks whether or not the specified value is undefined.  Empty
strings and nil pointers are considered as being undefined.
*/
func IsUndefined(
	value any,
) bool {
	return !isDefined(value)
}
