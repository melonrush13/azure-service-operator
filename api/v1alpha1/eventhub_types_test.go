// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("Eventhub", func() {
	var (
		key              types.NamespacedName
		created, fetched *Eventhub
	)

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	createEventHub := func(captureDescription CaptureDescription) *Eventhub {
		created := &Eventhub{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "foo",
				Namespace: "default",
			},
			Spec: EventhubSpec{
				Location:      resourcegroupLocation,
				Namespace:     "foo-eventhub-ns-name",
				ResourceGroup: "foo-resource-group",
				Properties: EventhubProperties{
					MessageRetentionInDays: 7,
					PartitionCount:         2,
					CaptureDescription:     captureDescription,
				},
			},
		}
		return created
	}

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.
	Context("Create API", func() {

		It("should create an object successfully (without Capture)", func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}

			created = createEventHub(CaptureDescription{})

			By("creating an API obj")
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())

			fetched = &Eventhub{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(fetched).To(Equal(created))

			By("deleting the created object")
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
			Expect(k8sClient.Get(context.TODO(), key, created)).ToNot(Succeed())
		})

		It("should create an object successfully (with Capture)", func() {
			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}

			capture := CaptureDescription{
				Destination: Destination{
					ArchiveNameFormat: "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}",
					BlobContainer:     "foo-blob-container",
					Name:              "EventHubArchive.AzureBlockBlob",
					StorageAccount: EventHubStorageAccount{
						ResourceGroup: "foo-resource-group",
						AccountName:   "fooaccountname",
					},
				},
				Enabled:           true,
				SizeLimitInBytes:  524288000,
				IntervalInSeconds: 90,
			}

			created = createEventHub(capture)

			By("creating an API obj")
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())

			fetched = &Eventhub{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(fetched).To(Equal(created))

			By("deleting the created object")
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
			Expect(k8sClient.Get(context.TODO(), key, created)).ToNot(Succeed())
		})

	})

})
