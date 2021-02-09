package util

import "github.com/drand/drand/fs"

// TODO get rid of drand dependency

const defaultRelativePath = "wasm"

func LocateFile(fileName string, relativePath ...string) string {
	relPath := defaultRelativePath
	if len(relativePath) > 0 {
		relPath = relativePath[0]
	}
	// check if this file exists
	exists, err := fs.Exists(fileName)
	if err != nil {
		panic(err)
	}
	if exists {
		return fileName
	}

	// walk up the directory tree to find the Wasm repo folder
	path := relPath
	exists, err = fs.Exists(path)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		path = "../" + path
		exists, err = fs.Exists(path)
		if err != nil {
			panic(err)
		}
		if exists {
			break
		}
	}

	// check if file is in Wasm repo
	path = path + "/" + fileName
	exists, err = fs.Exists(path)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("Missing wasm file: " + fileName)
	}
	return path
}
