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

func printDirRecur(path string, name string, padding string, depthLimit int, entryKind EntryKind, entryPos EntryPos) error {
	if depthLimit < 0 {
		return nil
	}

	if entryPos == Sibling {
		fmt.Printf("%s├─", padding)
	} else if entryPos == Last {
		fmt.Printf("%s└─", padding)
	}

	fmt.Printf("%s\n", name)

	if entryKind == FileEntry {
		return nil
	}

	fullPath := path + name + "/"
	entries, err := os.ReadDir(fullPath)
	if err != nil {
		return err
	}

	if entryPos == Sibling {
		childrenPadding := "│  "
		padding += childrenPadding
	} else if entryPos == Last {
		childrenPadding := "   "
		padding += childrenPadding
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

		if err := printDirRecur(fullPath, name, padding, depthLimit - 1, kind, pos); err != nil {
			return err
		}
	}
	return nil
}

func printDir(path string, maxDepth int) error {
	return printDirRecur("", path, "", maxDepth, DirEntry, Root)
}

func main() {
	folderPath := flag.String("f", "./", "folder path")
	maxDepth := flag.Int("d", 64, "max depth")

	flag.Parse()

	if folderPath == nil {
		flag.Usage()
		os.Exit(1)
	}

	if err := printDir(*folderPath, *maxDepth); err != nil {
		fmt.Println("Print err:", err)
		os.Exit(1)
	}
}
