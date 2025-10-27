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
  - Sequences (arrays, slices and maps)
  - Controllers (finite state machines)
  - Strings
  - Codex
  - Random
  - Reflection
*/
package module

// GLOBAL TYPES

/*
IteratorLike[V any] is an interface that declares the complete set of methods
that must be supported by each iterator.
*/
type IteratorLike[V any] interface {
	IsEmpty() bool
	ToStart()
	ToEnd()
	HasPrevious() bool
	GetPrevious() V
	HasNext() bool
	GetNext() V
	GetSize() uint
	GetSlot() uint
	SetSlot(
		slot uint,
	)
}

/*
Event is a constrained type representing an event type in a state machine.
Using a string type for an event makes it easier to print out in a human
readable way.
*/
type Event string

/*
State is a constrained type representing a state in a state machine.  Using a
string type for a state makes it easier to print out in a human readable way.
*/
type State string

/*
Transitions is a constrained type representing a row of states in a state machine.
*/
type Transitions []State

/*
ControllerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete controller-like class.

A controller-like class implements a finite state machine with possible event
types. It enforces the possible states of the state machine and allowed
transitions between states given a finite set of possible event types. It
implements a finite state machine with the following table structure:

	        -----------------------------------
	events: | [event1,  event2,  ... eventM ] |
	        -----------------------------------
	state1: | [invalid, state2,  ... invalid] |
	state2: | [state3,  stateN,  ... invalid] |
	        |                ...              |
	stateN: | [state1,  invalid, ... state3 ] |
	        -----------------------------------

The first row of the state machine defines the possible events that can occur.
Each subsequent row defines a state and the possible transitions from that
state to the next state for each possible event. Transitions marked as "invalid"
cannot occur. The state machine always starts in the first state of the finite
state machine (e.g. state1).
*/
type ControllerLike interface {
	ProcessEvent(
		event Event,
	) State
	GetState() State
	SetState(
		state State,
	)
	GetEvents() []Event
	GetTransitions() map[State]Transitions
}

// GLOBAL FUNCTIONS

// Sequences

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
Iterator[V any] returns an iterator over the specified array.  The
iterator moves forwards or backwards over the array landing in the slots between
the items in the array.  From a given slot the previous and next items are
accessible to the iterator.
*/
func Iterator[V any](
	array []V,
) IteratorLike[V] {
	return createIterator[V](array)
}

/*
ArraySize[V any] returns the current size of the specified array.
*/
func ArraySize[V any](
	array []V,
) uint {
	return arraySize(array)
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
MapSize[K comparable, V any] returns the current size of the specified map.
*/
func MapSize[K comparable, V any](
	map_ map[K]V,
) uint {
	return mapSize(map_)
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

// Controllers

/*
Controller returns a state machine that can be used to control an automaton.
*/
func Controller(
	events []Event,
	transitions map[State]Transitions,
	initialState State,
) ControllerLike {
	return createController(events, transitions, initialState)
}

/*
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
RenamePath renames an old file system path to a new one.
*/
func RenamePath(
	oldPath string,
	newPath string,
) {
	renamePath(oldPath, newPath)
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

// Codex

func Base16Encode(
	bytes []byte,
) string {
	return base16Encode(bytes)
}

func Base16Decode(
	encoded string,
) []byte {
	return base16Decode(encoded)
}

func Base32Encode(
	bytes []byte,
) string {
	return base32Encode(bytes)
}

func Base32Decode(
	encoded string,
) []byte {
	return base32Decode(encoded)
}

func Base64Encode(
	bytes []byte,
) string {
	return base64Encode(bytes)
}

func Base64Decode(
	encoded string,
) []byte {
	return base64Decode(encoded)
}

// Random

func RandomBoolean() bool {
	return randomBoolean()
}

func RandomOrdinal(
	maximum uint,
) uint {
	return randomOrdinal(maximum)
}

func RandomProbability() float64 {
	return randomProbability()
}

func RandomBytes(
	size uint,
) []byte {
	return randomBytes(size)
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
