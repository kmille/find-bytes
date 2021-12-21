package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	cont   *bool
	search []byte
)

func findOffsetInFile(filePath string) {
	//fmt.Printf("DEBUG: looking at %q\n", filePath)
	FileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Could not read file %q: %s\n", filePath, err)
		return
	}
	skipped := 0
	for {
		//fmt.Printf("file: %q\n", string(FileContent))
		//fmt.Printf("file: %q\n", hex.EncodeToString(FileContent))
		offset := strings.Index(string(FileContent), string(search))
		if offset == -1 {
			break
		}
		fmt.Printf("Match at offset %d (0x%x) until %d (0x%x) in file %q\n",
			(skipped + offset), (skipped + offset),
			(skipped + offset + len(search)), (skipped + offset + len(search)),
			filePath)
		if !*cont || (offset+len(search)) > len(FileContent) {
			break
		} else {
			skipped = offset + len(search)
			FileContent = FileContent[offset+len(search):]
		}
	}
}

func handleDirectory(path string) {
	//fmt.Printf("DEBUG: looking in %q\n", path)
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Could not read directory %q: %s\n", path, err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			handleDirectory(filepath.Join(path, entry.Name()))
		} else {
			// TODO: call this as a go-routine - but then we have to wait until all go-routines are done
			findOffsetInFile(filepath.Join(path, entry.Name()))
		}
	}
}

func main() {
	var (
		recursive  *bool
		bytes_raw  *string
		bytes_file *string
		fileOrPath string
		err        error
	)
	cont = flag.Bool("c", false, "continue after first match")
	recursive = flag.Bool("r", false, "find recursive")
	bytes_raw = flag.String("b", "", "bytes to look for in hex format")
	bytes_file = flag.String("f", "", "use bytes to look for from file")

	flag.Parse()
	if flag.NArg() == 0 || (len(*bytes_raw) == 0 && len(*bytes_file) == 0) {
		flag.Usage()
		os.Exit(1)
	}
	fileOrPath = flag.Arg(0)

	if len(*bytes_raw) != 0 {
		search, err = hex.DecodeString(*bytes_raw)
		if err != nil {
			fmt.Printf("Could not decode input as hex: %s\n", err)
			os.Exit(1)
		}
	} else {
		search, err = os.ReadFile(*bytes_file)
		if err != nil {
			fmt.Printf("Could not read file %q: %s\n", *bytes_file, err)
			os.Exit(1)
		}
	}

	if *recursive {
		handleDirectory(fileOrPath)
	} else {
		findOffsetInFile(fileOrPath)
	}

}
