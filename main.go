package main

import "fmt"
import "os"
import "flag"

func printDir(path string, dir string, padding []rune, maxDepth int, isDir bool) {
	if maxDepth < 0 {
		return
	}

	fullPath := path + dir

	fmt.Printf("%s%s\n", string(padding), dir)

	if padding[len(padding)-2] == '├' {
        padding[len(padding)-2] = '│'
		padding[len(padding)-1] = ' '
    } else if padding[len(padding)-2] == '└' {
		padding[len(padding)-2] = ' '
        padding[len(padding)-1] = ' '
    }

	if !isDir {
		return
	}

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		fmt.Println("Read dir err:", err)
		os.Exit(1)
	}

	for i, entry := range entries {
		name := entry.Name()

		if i == (len(entries) - 1) {
			padding = append(padding, '└', '─')

		} else {
			padding = append(padding, '├', '─')
		}

		printDir(fullPath + "/", name, padding, maxDepth - 1, entry.IsDir())

		padding = padding[0:len(padding) - 2]
	}
}

func main() {
	folderPath := flag.String("f", "./", "folder path")
	maxDepth := flag.Int("d", 64, "max depth")

	flag.Parse()

	if len(*folderPath) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	padding := make([]rune, 0)
    padding = append(padding, ' ', ' ')

	printDir("", *folderPath, padding, *maxDepth, true)
}
