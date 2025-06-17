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

package module

import (
	fmt "fmt"
)

func relativeToZeroBased[V any](
	array []V,
	relative Index,
) int {
	var zeroBased int
	var size = Index(len(array))
	switch {
	case size == 0:
		// The Array is empty.
		panic("Cannot index an empty Array.")
	case relative == 0:
		// Zero is not a relative index.
		panic("Indices must be positive or negative ordinals, not zero.")
	case relative < -size || relative > size:
		// The index is outside the bounds of the specified range.
		var message = fmt.Sprintf(
			"The specified index is outside the allowed ranges [-%v..-1] and [1..%v]: %v",
			size,
			size,
			relative)
		panic(message)
	case relative < 0:
		// Convert a negative index.
		zeroBased = int(relative + size)
	case relative > 0:
		// Convert a positive index.
		zeroBased = int(relative - 1)
	}
	return zeroBased
}

func zeroBasedToRelative[V any](
	array []V,
	zeroBased int,
) Index {
	var relative Index
	var size = len(array)
	switch {
	case size == 0:
		// The Array is empty.
		panic("Cannot index an empty Array.")
	case zeroBased < 0 || zeroBased >= size:
		// The index is outside the bounds of the array.
		var message = fmt.Sprintf(
			"The specified index is outside the allowed range [0..%v): %v",
			size,
			zeroBased,
		)
		panic(message)
	default:
		relative = Index(zeroBased + 1)
	}
	return relative
}

func copyArray[V any](
	array []V,
) []V {
	var size = len(array)
	var duplicate = make([]V, size)
	copy(duplicate, array)
	return duplicate
}

func arraysAreEqual[V comparable](
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

func combineArrays[V any](
	first []V,
	second []V,
) []V {
	return append(first, second...)
}

func copyMap[K comparable, V any](
	map_ map[K]V,
) map[K]V {
	var duplicate = make(map[K]V)
	for key, value := range map_ {
		duplicate[key] = value
	}
	return duplicate
}

func mapsAreEqual[K comparable, V comparable](
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

func combineMaps[K comparable, V any](
	first map[K]V,
	second map[K]V,
) map[K]V {
	var result = make(map[K]V)
	for key, value := range first {
		result[key] = value
	}
	for key, value := range second {
		result[key] = value
	}
	return result
}
