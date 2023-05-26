package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"google.golang.org/api/dataproc/v1"
)

// Tests schema version migration by creating a cluster with an old version of the provider (4.65.0)
// and then updating it with the current version the provider.
func TestAccDataprocClusterLabelsMigration_withoutLabels_withoutChanges(t *testing.T) {
	SkipIfVcr(t)
	t.Parallel()

	rnd := RandString(t, 10)
	var cluster dataproc.Cluster
	oldVersion := map[string]resource.ExternalProvider{
		"google": {
			VersionConstraint: "4.65.0", // a version that doesn't separate user defined labels and system labels
			Source:            "registry.terraform.io/hashicorp/google",
		},
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { AccTestPreCheck(t) },
		CheckDestroy: testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config:            testAccDataprocClusterLabelsMigration_withoutLabels(rnd),
				ExternalProviders: oldVersion,
			},
			{
				Config:                   testAccDataprocClusterLabelsMigration_withoutLabels(rnd),
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckNoResourceAttr("google_dataproc_cluster.with_labels", "labels.%"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "4"),
				),
			},
			{
				Config:                   testAccDataprocCluster_withLabels(rnd),
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.%", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.key1", "value1"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "5"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.key1", "value1"),
				),
			},
		},
	})
}

func TestAccDataprocClusterLabelsMigration_withLabels_withoutChanges(t *testing.T) {
	SkipIfVcr(t)
	t.Parallel()

	rnd := RandString(t, 10)
	var cluster dataproc.Cluster
	oldVersion := map[string]resource.ExternalProvider{
		"google": {
			VersionConstraint: "4.65.0", // a version that doesn't separate user defined labels and system labels
			Source:            "registry.terraform.io/hashicorp/google",
		},
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { AccTestPreCheck(t) },
		CheckDestroy: testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config:            testAccDataprocCluster_withLabels(rnd),
				ExternalProviders: oldVersion,
			},
			{
				Config:                   testAccDataprocCluster_withLabels(rnd),
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.%", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.key1", "value1"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "5"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.key1", "value1"),
				),
			},
			{
				Config:                   testAccDataprocCluster_withLabelsUpdate(rnd),
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					// We only provide two, so expect 2.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.%", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.key2", "value2"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "5"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.key2", "value2"),
				),
			},
		},
	})
}

func TestAccDataprocClusterLabelsMigration_withUpdate(t *testing.T) {
	SkipIfVcr(t)
	t.Parallel()

	rnd := RandString(t, 10)
	var cluster dataproc.Cluster
	oldVersion := map[string]resource.ExternalProvider{
		"google": {
			VersionConstraint: "4.65.0", // a version that doesn't separate user defined labels and system labels
			Source:            "registry.terraform.io/hashicorp/google",
		},
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { AccTestPreCheck(t) },
		CheckDestroy: testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config:            testAccDataprocClusterLabelsMigration_withoutLabels(rnd),
				ExternalProviders: oldVersion,
			},
			{
				Config:                   testAccDataprocCluster_withLabels(rnd),
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.%", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.key1", "value1"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "5"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.key1", "value1"),
				),
			},
			{
				Config:                   testAccDataprocClusterLabelsMigration_withoutLabels(rnd),
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckNoResourceAttr("google_dataproc_cluster.with_labels", "labels.%"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "4"),
				),
			},
		},
	})
}

func TestAccDataprocClusterLabelsMigration_withRemoval(t *testing.T) {
	SkipIfVcr(t)
	t.Parallel()

	rnd := RandString(t, 10)
	var cluster dataproc.Cluster
	oldVersion := map[string]resource.ExternalProvider{
		"google": {
			VersionConstraint: "4.65.0", // a version that doesn't separate user defined labels and system labels
			Source:            "registry.terraform.io/hashicorp/google",
		},
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { AccTestPreCheck(t) },
		CheckDestroy: testAccCheckDataprocClusterDestroy(t),
		Steps: []resource.TestStep{
			{
				Config:            testAccDataprocCluster_withLabels(rnd),
				ExternalProviders: oldVersion,
			},
			{
				Config:                   testAccDataprocClusterLabelsMigration_withoutLabels(rnd),
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckNoResourceAttr("google_dataproc_cluster.with_labels", "labels.%"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "4"),
				),
			},
			{
				Config:                   testAccDataprocCluster_withLabels(rnd),
				ProtoV5ProviderFactories: ProtoV5ProviderFactories(t),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataprocClusterExists(t, "google_dataproc_cluster.with_labels", &cluster),

					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.%", "1"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "labels.key1", "value1"),
					// We only provide one, but GCP adds three and goog-dataproc-autozone is added internally, so expect 5.
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.%", "5"),
					resource.TestCheckResourceAttr("google_dataproc_cluster.with_labels", "effective_labels.key1", "value1"),
				),
			},
		},
	})
}

func testAccDataprocClusterLabelsMigration_withoutLabels(rnd string) string {
	return fmt.Sprintf(`
resource "google_dataproc_cluster" "with_labels" {
  name   = "tf-test-dproc-%s"
  region = "us-central1"
}
`, rnd)
}
