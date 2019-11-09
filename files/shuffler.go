package files

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type RandShuffler struct {
	shuffledEntriesCount int
	prefixCharCount      int
}

func NewRandShuffler(shuffledEntriesCount int) *RandShuffler {
	rand.Seed(time.Now().UnixNano())
	return &RandShuffler{
		shuffledEntriesCount: shuffledEntriesCount,
		prefixCharCount:      int(math.Log10(float64(shuffledEntriesCount))) + 1,
	}
}

func (s RandShuffler) Shuffle(fileName string) string {
	randomInt := rand.Intn(s.shuffledEntriesCount)
	formatPattern := fmt.Sprintf("%%0%dd-%%s", s.prefixCharCount)
	return fmt.Sprintf(formatPattern, randomInt, fileName)
}
