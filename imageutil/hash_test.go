package imageutil

import (
	"testing"
)

func TestAHash(t *testing.T) {
	img, _ := Open("test.png")
	t.Log(AHash(img)) // 17277923389503686655
}
