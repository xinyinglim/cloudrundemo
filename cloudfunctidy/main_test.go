package cloudfunctidy

import (
	"context"
	"testing"
)

func TestMain(t *testing.T) {
	fileURI := "gs://cloudfunctidy/upload/BeaWilma.jpg"
	c := context.Background()
	labels, err := getLabelsUsingCloudVisionAPI(c, fileURI)
	if err != nil {
		t.Errorf("Unable to get labels: %v", err)
	}
	err = uploadImageLabelsToFirestore(c, fileURI, labels)
	if err != nil {
		t.Errorf("Unable to upload Image Labels to Firestore: %v", err)
	}
	return
}
