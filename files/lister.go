package files

import (
	"os"
	"path/filepath"
)

type Lister struct {
	wildCard string
}

func NewLister(wildCard string) *Lister {
	return &Lister{wildCard: wildCard}
}

func (l Lister) ListFiles(path string) ([]string, error) {
	result := []string{}
	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, currentError error) error {
		if currentError != nil {
			return currentError
		}

		if info.IsDir() {
			return nil
		}

		matched, err := filepath.Match(l.wildCard, info.Name())
		if err != nil {
			return err
		}

		if matched {
			result = append(result, filepath.Join(currentPath))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil

}
