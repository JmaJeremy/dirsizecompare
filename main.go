package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/inancgumus/screen"
)

func getDirSize(dir string) int64 {
	var totalSize int64 = 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			totalSize += info.Size()
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return -1
	}

	return totalSize
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./dirsizecompare [origin] [destination]")
		return
	}

	var orig_sz int64
	var dest_sz int64
	var pct float64

	orig_sz = getDirSize(os.Args[1])
	dest_sz = getDirSize(os.Args[2])
	fmt.Println("Original size: ", orig_sz)
	fmt.Println("Destination size: ", dest_sz)
	pct = (float64(dest_sz) / float64(orig_sz)) * 100
	fmt.Printf("%.2f%% \n", pct)

	for true {
		time.Sleep(5 * time.Second)
		screen.Clear()
		screen.MoveTopLeft()
		dest_sz = getDirSize(os.Args[2])
		fmt.Println("Original size: ", orig_sz)
		fmt.Println("Destination size: ", dest_sz)
		pct = (float64(dest_sz) / float64(orig_sz)) * 100
		fmt.Printf("%.2f%% \n\nPress CTRL+C to end...\n", pct)
	}
}
