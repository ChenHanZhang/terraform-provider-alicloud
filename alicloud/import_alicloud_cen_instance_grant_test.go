package alicloud

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This testcase can not work with multi account.
func SkipTestAccAlicloudCenInstanceGrant_importBasic(t *testing.T) {
	resourceName := "alicloud_cen_instance_grant.default"
	rand := acctest.RandIntRange(1000000, 9999999)
	// multi provideris
	var providers []*schema.Provider
	providerFactories := map[string]func() (*schema.Provider, error){
		"alicloud": func() (*schema.Provider, error) {
			p := Provider()
			providers = append(providers, p.(*schema.Provider))
			return p.(*schema.Provider), nil
		},
	}

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithMultipleAccount(t)
		},

		ProviderFactories: providerFactories,
		CheckDestroy:      testAccCheckCenInstanceGrantDestroyWithProviders(&providers),
		Steps: []resource.TestStep{
			{
				Config: testAccCenInstanceGrantBasic(os.Getenv("ALICLOUD_ACCESS_KEY_2"), os.Getenv("ALICLOUD_SECRET_KEY_2"), os.Getenv("ALICLOUD_ACCOUNT_ID_1"), os.Getenv("ALICLOUD_ACCOUNT_ID_2"), rand),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
