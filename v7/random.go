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
	var random, _ = ran.Int(ran.Reader, big.NewInt(int64(2)))
	return int(random.Int64()) > 0
}

func randomOrdinal(maximum Ordinal) Ordinal {
	var random, _ = ran.Int(ran.Reader, big.NewInt(int64(maximum)))
	return Ordinal(random.Int64() + 1)
}

func randomProbability() Probability {
	var maximum = Ordinal(1 << 53) // 53 bits for the sign and mantissa.
	return Probability(float64(randomOrdinal(maximum)) / float64(maximum))
}

func randomBytes(size Size) []byte {
	var bytes = make([]byte, size)
	_, _ = ran.Read(bytes) // This call can never fail.
	return bytes
}
