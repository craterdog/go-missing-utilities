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
	osx "os"
	sts "strings"
)

func pathExists(
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

func removePath(
	path string,
) {
	var err = osx.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

func renamePath(
	oldPath string,
	newPath string,
) {
	var err = osx.Rename(oldPath, newPath)
	if err != nil {
		panic(err)
	}
}

func homeDirectory() string {
	directory, err := osx.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return directory + "/"
}

func makeDirectory(
	directory string,
) {
	var err = osx.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}

func remakeDirectory(
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

func readDirectory(
	directory string,
) []string {
	if !sts.HasSuffix(directory, "/") {
		directory += "/"
	}
	var files, err = osx.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	var filenames = make([]string, 0, len(files))
	for _, file := range files {
		var filename = file.Name()
		if sts.HasPrefix(filename, ".") {
			// Skip hiddend files.
			continue
		}
		filenames = append(filenames, filename)
	}
	return filenames
}

func readFile(
	filename string,
) string {
	var bytes, err = osx.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var source = string(bytes)
	return source
}

func writeFile(
	filename string,
	source string,
) {
	var bytes = []byte(source)
	var err = osx.WriteFile(filename, bytes, 0644)
	if err != nil {
		panic(err)
	}
}
