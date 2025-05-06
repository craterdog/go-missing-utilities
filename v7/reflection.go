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
	ref "reflect"
)

func implementsInterface(
	value any,
	pointer any,
) bool {
	if !isDefined(value) {
		return false
	}
	var valueType = ref.TypeOf(value)
	var interface_ = ref.TypeOf(pointer).Elem()
	return valueType.Implements(interface_)
}

func isDefined(
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
