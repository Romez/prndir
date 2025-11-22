package main

import "fmt"
import "os"
import "flag"

type EntryKind int
type EntryPos int

const (
	Root EntryPos = iota
	Last
	Sibling
)

const (
	FileEntry EntryKind = iota
	DirEntry
)

func printDirRecur(path string, name string, padding []rune, depthLimit int, entryKind EntryKind, entryPos EntryPos) {
	if depthLimit < 0 {
		return
	}

	if entryPos == Sibling {
		fmt.Printf("%s├─", string(padding))
	} else if entryPos == Last {
		fmt.Printf("%s└─", string(padding))
	}

	fmt.Printf("%s\n", name)

	if entryKind == FileEntry {
		return
	}

	fullPath := path + name + "/"
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		fmt.Println("Read dir err:", err)
		os.Exit(1)
	}

	if entryPos == Sibling {
		padding = append(padding, '│', ' ', ' ')
	} else if entryPos == Last {
		padding = append(padding, ' ', ' ', ' ')
	}

	entriesLen := len(entries)
	
	for i, entry := range entries {
		name := entry.Name()

		kind := FileEntry
		if entry.IsDir() {
			kind = DirEntry
		}

		pos := Sibling
		if i == (entriesLen - 1) {
			pos = Last
		}

		printDirRecur(fullPath, name, padding, depthLimit - 1, kind, pos)
	}
}

func printDir(path string, maxDepth int) {
	padding := make([]rune, 0)
	printDirRecur("", path, padding, maxDepth, DirEntry, Root)
}

func main() {
	folderPath := flag.String("f", "./", "folder path")
	maxDepth := flag.Int("d", 64, "max depth")

	flag.Parse()

	if folderPath == nil {
		flag.Usage()
		os.Exit(1)
	}

	printDir(*folderPath, *maxDepth)
}
