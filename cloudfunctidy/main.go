package cloudfunctidy

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"cloud.google.com/go/firestore"
	vision "cloud.google.com/go/vision/apiv1"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

var projectID string = "demos-xy"
var maxNoOfLabels int = 1

type GCSEvent struct {
	Name        string    `json:"name"`
	Bucket      string    `json:"bucket"`
	ContentType string    `json:"contentType"`
	Updated     time.Time `json:"updated"`
}

type Image struct {
	URI    string
	Labels []Label
}

type Label struct {
	Description string
	Score       float32
}

func newImage(uri string, annotations []*visionpb.EntityAnnotation) Image {
	imageLabels := []Label{}
	for i, annotation := range annotations {
		if i == maxNoOfLabels {
			break
		}
		imageLabels = append(
			imageLabels,
			Label{
				Description: annotation.Description,
				Score:       annotation.Score,
			},
		)
	}
	return Image{
		URI:    uri,
		Labels: imageLabels,
	}
}

func LabelImage(c context.Context, e GCSEvent) error {
	fileURI := fmt.Sprintf("gs://%s/%s", e.Bucket, e.Name)
	labels, err := getLabelsUsingCloudVisionAPI(c, fileURI)
	if err != nil {
		return fmt.Errorf("Unable to get labels: %v", err)
	}
	err = uploadImageLabelsToFirestore(c, fileURI, labels)
	if err != nil {
		return fmt.Errorf("Unable to upload Image Labels to Firestore: %v", err)
	}
	return nil
}

func uploadImageLabelsToFirestore(
	c context.Context,
	fileURI string,
	labels []*visionpb.EntityAnnotation,
) error {
	client, err := firestore.NewClient(c, projectID)
	if err != nil {
		return fmt.Errorf("Failed to create new firestore client: %v", err)
	}
	defer client.Close()
	toUploadImageLabel := newImage(fileURI, labels)
	firestoreID := getFileID(fileURI)
	_, err = client.Collection("imageLabels").Doc(firestoreID).Set(c, toUploadImageLabel)
	if err != nil {
		return fmt.Errorf("Unable to upload labels to Firestore: %v", err)
	}
	return nil
}

func getLabelsUsingCloudVisionAPI(c context.Context, fileURI string) ([]*visionpb.EntityAnnotation, error) {
	visionClient, err := vision.NewImageAnnotatorClient(c)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer visionClient.Close()
	image := vision.NewImageFromURI(fileURI)
	var annotations []*visionpb.EntityAnnotation
	annotations, err = visionClient.DetectLabels(c, image, nil, maxNoOfLabels)
	if err != nil {
		return []*visionpb.EntityAnnotation{}, fmt.Errorf("VisionClient.DetectLabels: %v", err)
	}
	return annotations, nil
}

func getFileID(uri string) string {
	re, _ := regexp.Compile("([A-z]+).jpg$")
	return re.FindStringSubmatch(uri)[1]
}
