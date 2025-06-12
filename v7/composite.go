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
