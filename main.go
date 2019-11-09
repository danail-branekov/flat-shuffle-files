package main

import (
	"flag"
	"flat-shuffle-files/files"
	"fmt"
	"os"
)

type Worker interface {
	Process(files map[string]string) error
}

func main() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] <source-directory>\n\n\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
	}
	fileMaskFlag := flag.String("mask", "*", "mask of files to shuffle, for example '*.mp3'")
	destinationDirFlag := flag.String("destination", ".", "where to move shuffled files to, defaults to current directory. Must be an existing directory")
	pretendFlag := flag.Bool("pretend", false, "print report on which files will be copied")
	flag.Parse()

	fileMask := *fileMaskFlag
	destinationDir := *destinationDirFlag
	pretend := *pretendFlag

	if !isExistingDirectory(destinationDir) {
		exitOnError(fmt.Errorf("Directory %q does not exist", destinationDir))
	}

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	sourceDir := flag.Arg(0)

	lister := files.NewLister(fileMask)
	sourceFiles, err := lister.ListFiles(sourceDir)
	if err != nil {
		exitOnError(err)
	}

	shuffler := files.NewFlatMapShuffler(destinationDir, files.NewRandShuffler(len(sourceFiles)))
	var worker Worker
	if pretend {
		worker = files.NewPrintingWorker()
	} else {
		worker = files.NewCopyingWorker()
	}

	shuffledFiles := shuffler.FlatMapShuffle(sourceFiles)
	if err := worker.Process(shuffledFiles); err != nil {
		exitOnError(err)
	}

	fmt.Printf("\n\n------\n\n%d files processed. Bye!", len(sourceFiles))
}

func isExistingDirectory(dir string) bool {
	stat, err := os.Stat(dir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Cannot stat %q: %v\n", dir, err)
		return false
	}
	return stat.IsDir()
}

func exitOnError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "Unexpected error ocurred: %v", err)
	os.Exit(1)
}
