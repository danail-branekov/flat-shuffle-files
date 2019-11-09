package files

import (
	"math/rand"
	"path/filepath"
	"time"
)

type Shuffler interface {
	Shuffle(path string) string
}

type FlatMapShuffler struct {
	targetPath string
	shuffler   Shuffler
}

func NewFlatMapShuffler(targetPath string, shuffler Shuffler) *FlatMapShuffler {
	return &FlatMapShuffler{targetPath: targetPath, shuffler: shuffler}
}

func (m FlatMapShuffler) FlatMapShuffle(paths []string) map[string]string {
	rand.Seed(time.Now().UnixNano())

	result := make(map[string]string)
	for _, p := range paths {
		fileTargetPath := filepath.Join(m.targetPath, m.shuffler.Shuffle(filepath.Base(p)))
		result[p] = fileTargetPath
	}

	return result
}
