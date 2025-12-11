package imageutil

import (
	"testing"
)

func TestFindBlobs(t *testing.T) {
	img, _ := Open("mask.png")
	result := FindBlobs(img)
	for _, blob := range result.Blobs {
		t.Logf("Blob %d: Area: %d, Bounds: %v, Centroid: %v\n",
			blob.ID, blob.Area, blob.Bounds, blob.Centroid)
	}
}
