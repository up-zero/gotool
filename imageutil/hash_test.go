package imageutil

import (
	"testing"
)

func TestAHash(t *testing.T) {
	img, _ := Open("test.png")
	t.Log(AHash(img)) // 17277923389503686655
}

func TestDHash(t *testing.T) {
	img, _ := Open("test.png")
	t.Log(DHash(img)) // 7430675043070621677
}
