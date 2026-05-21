package imageutil

import (
	"image"
	"image/color"
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
	img := image.NewGray(image.Rect(0, 0, 8, 8))

	// 较大的 2x2 Blob，面积应为 4
	img.SetGray(1, 1, color.Gray{Y: 255})
	img.SetGray(2, 1, color.Gray{Y: 255})
	img.SetGray(1, 2, color.Gray{Y: 255})
	img.SetGray(2, 2, color.Gray{Y: 255})

	// 较小的单像素 Blob，面积应为 1
	img.SetGray(6, 6, color.Gray{Y: 255})

	maxBlob := FindBlobs(img).GetLargestBlob()
	if maxBlob == nil {
		t.Fatal("expected max blob, got nil")
	}

	if maxBlob.Area != 4 {
		t.Fatalf("expected max blob area 4, got %d", maxBlob.Area)
	}

	expectedBounds := image.Rect(1, 1, 3, 3)
	if maxBlob.Bounds != expectedBounds {
		t.Fatalf("expected bounds %v, got %v", expectedBounds, maxBlob.Bounds)
	}
}

func TestBlobResultGetMaxBlobEmpty(t *testing.T) {
	if (&BlobResult{}).GetLargestBlob() != nil {
		t.Fatal("expected nil for empty blob result")
	}

	var result *BlobResult
	if result.GetLargestBlob() != nil {
		t.Fatal("expected nil for nil blob result")
	}
}
