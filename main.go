package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"flag"

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
/*	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Usage: dirsizecompare DIR1 DIR2 [-f]")
		return
	} */

	var orig_sz int64
	var dest_sz int64
	var pct float64

	followPtr := flag.Bool("f", false, "follow")
	dir1Ptr := flag.String("d1", "", "DIR1")
	dir2Ptr := flag.String("d2", "", "DIR2")
	flag.Parse()
//	fmt.Println("f:", *followPtr)
	if len(*dir1Ptr) < 1 &&  len(*dir2Ptr) < 1 {
		fmt.Println("Usage: dirsizecompare -d1=DIR1 -d2=DIR2 [-f]")
		return
	}

	//orig_sz = getDirSize(os.Args[1])
	//dest_sz = getDirSize(os.Args[2])
	orig_sz = getDirSize(*dir1Ptr)
	dest_sz = getDirSize(*dir2Ptr)
	fmt.Println("Original size: ", orig_sz)
	fmt.Println("Destination size: ", dest_sz)
	pct = (float64(dest_sz) / float64(orig_sz)) * 100
	fmt.Printf("%.2f%% \n", pct)

	
	if *followPtr {
		for true {
			time.Sleep(5 * time.Second)
			screen.Clear()
			screen.MoveTopLeft()
			dest_sz = getDirSize(*dir2Ptr)
			fmt.Println("Original size: ", orig_sz)
			fmt.Println("Destination size: ", dest_sz)
			pct = (float64(dest_sz) / float64(orig_sz)) * 100
			fmt.Printf("%.2f%% \n\nPress CTRL+C to end...\n", pct)
		}
	}
}
