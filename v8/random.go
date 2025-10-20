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
	ran "crypto/rand"
	big "math/big"
)

func randomBoolean() bool {
	// A random boolean is in the range [0..1].
	var random, _ = ran.Int(ran.Reader, big.NewInt(int64(2)))
	return random.Int64() > 0
}

func randomOrdinal(
	maximum uint,
) uint {
	// A random integer is in the range [0..maximum).
	var random, _ = ran.Int(ran.Reader, big.NewInt(int64(maximum)))
	// Convert [0..maximum) to [1..maximum].
	return uint(random.Uint64() + 1)
}

func randomProbability() float64 {
	// Use 53 bits for the sign and mantissa only.
	var maximum = uint(1 << 53)
	// A random probability is in the range (0.0..1.0] since something with
	// zero probability will never occur so we use [1..maximum]/maximum.
	return float64(randomOrdinal(maximum)) / float64(maximum)
}

func randomBytes(
	size uint,
) []byte {
	var bytes = make([]byte, size)
	_, _ = ran.Read(bytes) // This call should never fail.
	return bytes
}
