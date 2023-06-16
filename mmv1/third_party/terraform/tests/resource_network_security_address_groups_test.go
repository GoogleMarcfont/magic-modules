<<<<<<<< HEAD:mmv1/third_party/terraform/tests/resource_network_security_address_groups_test.go
package google
========
<% autogen_exception -%>
package networksecurity_test
<% unless version == 'ga' -%>
>>>>>>>> d20c173b9 (Generate Mmv1 test files to the service packages):mmv1/third_party/terraform/services/networksecurity/resource_network_security_address_group_test.go.erb

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-provider-google/google/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccNetworkSecurityAddressGroups_update(t *testing.T) {
	t.Parallel()

<<<<<<<< HEAD:mmv1/third_party/terraform/tests/resource_network_security_address_groups_test.go
	addressGroupsName := fmt.Sprintf("tf-test-address-group-%s", acctest.RandString(t, 10))
========
<<<<<<< HEAD:mmv1/third_party/terraform/tests/resource_network_security_address_groups_test.go.erb
    addressGroupsName := fmt.Sprintf("tf-test-address-group-%s", acctest.RandString(t, 10))
=======
<<<<<<< HEAD:mmv1/third_party/terraform/tests/resource_network_security_address_groups_test.go.erb
    addressGroupsName := fmt.Sprintf("tf-test-address-group-%s", RandString(t, 10))
>>>>>>> c13a90bef (Generate Mmv1 test files to the service packages):mmv1/third_party/terraform/services/networksecurity/resource_network_security_address_group_test.go.erb
>>>>>>>> d20c173b9 (Generate Mmv1 test files to the service packages):mmv1/third_party/terraform/services/networksecurity/resource_network_security_address_group_test.go.erb
	projectName := GetTestProjectFromEnv()
=======
    addressGroupsName := fmt.Sprintf("tf-test-address-group-%s", acctest.RandString(t, 10))
	projectName := envvar.GetTestProjectFromEnv()
>>>>>>> 12945f953 (Generate Mmv1 test files to the service packages):mmv1/third_party/terraform/services/networksecurity/resource_network_security_address_group_test.go.erb

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckNetworkSecurityAddressGroupDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkSecurityAddressGroups_basic(addressGroupsName, projectName),
			},
			{
				ResourceName:      "google_network_security_address_group.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccNetworkSecurityAddressGroups_update(addressGroupsName, projectName),
			},
			{
				ResourceName:      "google_network_security_address_group.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetworkSecurityAddressGroups_basic(addressGroupsName, projectName string) string {
	return fmt.Sprintf(`
resource "google_network_security_address_group" "foobar" {
    name        = "%s"
    parent 		= "projects/%s"
    location    = "us-central1"
    description = "my address groups"
    type        = "IPV4"
    capacity    = "100"
    labels      = {
		foo = "bar"
    }
    items 		= ["208.80.154.224/32"]
}
`, addressGroupsName, projectName)
}

func testAccNetworkSecurityAddressGroups_update(addressGroupsName, projectName string) string {
	return fmt.Sprintf(`
resource "google_network_security_address_group" "foobar" {
    name        = "%s"
	parent 		= "projects/%s"
    location    = "us-central1"
    description = "my address groups. Update"
    type        = "IPV4"
    capacity    = "100"
    labels      = {
		foo = "foo"
    }
    items 		= ["208.80.155.224/32", "208.80.154.224/32"]
}
`, addressGroupsName, projectName)
}
