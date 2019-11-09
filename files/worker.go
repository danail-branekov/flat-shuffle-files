package files

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"io/ioutil"
)

type ProcessFunc func(sourcePath, targetPath string) error

type Worker struct {
	processFunc ProcessFunc
}

func NewPrintingWorker() *Worker {
	return &Worker{processFunc: func(sourcePath, targetPath string) error {
		fmt.Printf("%s -> %s\n", sourcePath, targetPath)
		return nil
	}}
}

func NewCopyingWorker() *Worker {
	return &Worker{processFunc: func(sourcePath, targetPath string) error {
		fmt.Printf("Copying: %q -> %q\n", sourcePath, targetPath)
		data, err := ioutil.ReadFile(sourcePath)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(targetPath, data, 0644); err != nil {
			return err
		}
		return nil
	}}
}

func (w Worker) Process(mappedFiles map[string]string) error {
	fmt.Printf("Processing %d files...\n", len(mappedFiles))
	var multiErr *multierror.Error
	for k, v := range mappedFiles {
		if err := w.processFunc(k, v); err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}

	return multiErr.ErrorOrNil()
}
