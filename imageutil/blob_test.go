package imageutil

import (
	"testing"
)

func TestFindBlobs(t *testing.T) {
	img, _ := Open("mask.jpg")
	result := FindBlobs(img)
	for _, blob := range result.Blobs {
		t.Logf("Blob %d: Area: %d, Bounds: %v, Centroid: %v\n",
			blob.ID, blob.Area, blob.Bounds, blob.Centroid)
	}
}

func TestBlobResultGetMaxBlob(t *testing.T) {
	img, _ := Open("mask.jpg")
	result := FindBlobs(img)
	blob := result.GetLargestBlob()
	if blob == nil {
		t.Fatal("expected largest blob, got nil")
	}
	t.Logf("Blob %d: Area: %d, Bounds: %v, Centroid: %v\n",
		blob.ID, blob.Area, blob.Bounds, blob.Centroid)
}

func TestFindLargestBlob(t *testing.T) {
	img, _ := Open("mask.jpg")
	blob := FindLargestBlob(img)
	if blob == nil {
		t.Fatal("expected largest blob, got nil")
	}
	t.Logf("Blob %d: Area: %d, Bounds: %v, Centroid: %v\n",
		blob.ID, blob.Area, blob.Bounds, blob.Centroid)
}
