package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	shell := buildShellFromLogs("./input.txt")

	fmt.Println("Puzzle1 \t", puzzle1(shell))
	fmt.Println("Puzzle2 \t", puzzle2(shell))
}

func puzzle1(shell *shell) int64 {
	var result int64

	rootDir := shell.rootDir
	walkDir(rootDir, func(node *node) {
		if node.nodeType == dirNode && node.size <= 100_000 {
			result += node.size
		}
	})

	return result
}

func puzzle2(shell *shell) int64 {
	rootDir := shell.rootDir

	var totalDiskSize int64 = 70_000_000
	var minFreeRequired int64 = 30_000_000

	totalUsed := rootDir.size
	unused := totalDiskSize - totalUsed

	if unused >= minFreeRequired {
		return 0
	}

	needToFree := minFreeRequired - unused

	result := rootDir.size
	walkDir(rootDir, func(node *node) {
		if node.nodeType == dirNode && node.size >= needToFree {
			if node.size < result {
				result = node.size
			}
		}
	})

	return result
}

func buildShellFromLogs(logfile string) *shell {
	shell := newShell(newFS())

	f, err := os.Open(logfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if len(line) < 1 {
			continue
		}

		switch true {
		case strings.HasPrefix(line, prompt):
			args := strings.Split(line, " ")
			if args[1] == cdCmd {
				shell.cd(args[2])
			}
		case strings.HasPrefix(line, dirCmd):
			dirInfo := strings.Split(line, " ")
			shell.mkdir(dirInfo[1])
		default:
			fileInfo := strings.Split(line, " ")
			size, _ := strconv.ParseInt(fileInfo[0], 10, 0)
			shell.fallocate(fileInfo[1], size)
		}
	}

	return shell
}
