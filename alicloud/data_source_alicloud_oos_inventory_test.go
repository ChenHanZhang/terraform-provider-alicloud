package alicloud

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAliCloudOosInventorySchemaDataSource_basic(t *testing.T) {
	resourceId := "data.alicloud_oos_inventory_schema.default"
	checkoutSupportedRegions(t, true, connectivity.OOSSupportRegions)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
data "alicloud_oos_inventory_schema" "default" {
  type_name = "ACS:InstanceInformation"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceId, "ids.#", "1"),
					resource.TestCheckResourceAttr(resourceId, "ids.0", "ACS:InstanceInformation"),
					resource.TestCheckResourceAttr(resourceId, "schemas.#", "1"),
					resource.TestCheckResourceAttr(resourceId, "schemas.0.id", "ACS:InstanceInformation"),
					resource.TestCheckResourceAttr(resourceId, "schemas.0.type_name", "ACS:InstanceInformation"),
					resource.TestCheckResourceAttrSet(resourceId, "schemas.0.version"),
					resource.TestCheckResourceAttrSet(resourceId, "schemas.0.attributes_json"),
					resource.TestCheckResourceAttrSet(resourceId, "schemas.0.schema_json"),
					resource.TestCheckResourceAttrSet(resourceId, "schemas_json"),
				),
			},
		},
	})
}

func TestAccAliCloudOosInventorySearchDataSource_basic(t *testing.T) {
	resourceId := "data.alicloud_oos_inventory_search.default"
	checkoutSupportedRegions(t, true, connectivity.OOSSupportRegions)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
data "alicloud_oos_inventory_search" "default" {
  filter {
    name     = "ACS:InstanceInformation.InstanceId"
    operator = "Equal"
    values   = ["i-tf-testacc-not-exist"]
  }
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceId, "ids.#", "0"),
					resource.TestCheckResourceAttr(resourceId, "entities.#", "0"),
					resource.TestCheckResourceAttr(resourceId, "entities_json", "[]"),
				),
			},
		},
	})
}

func TestAccAliCloudOosInventoryEntriesDataSource_basic(t *testing.T) {
	resourceId := "data.alicloud_oos_inventory_entries.default"
	checkoutSupportedRegions(t, true, connectivity.OOSSupportRegions)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithEnvVariable(t, "ALICLOUD_OOS_INVENTORY_INSTANCE_ID")
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
data "alicloud_oos_inventory_entries" "default" {
  instance_id = "%s"
  type_name   = "ACS:InstanceInformation"
}
`, os.Getenv("ALICLOUD_OOS_INVENTORY_INSTANCE_ID")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceId, "instance_id", os.Getenv("ALICLOUD_OOS_INVENTORY_INSTANCE_ID")),
					resource.TestCheckResourceAttr(resourceId, "type_name", "ACS:InstanceInformation"),
					resource.TestCheckResourceAttrSet(resourceId, "entries_json"),
				),
			},
		},
	})
}

func TestExpandOosInventoryFilters(t *testing.T) {
	request := make(map[string]interface{})
	expandOosInventoryFilters(request, []interface{}{
		map[string]interface{}{
			"name":     "ACS:InstanceInformation.InstanceId",
			"operator": "Equal",
			"values":   []interface{}{"i-bp1", "i-bp2"},
		},
	})

	expected := map[string]interface{}{
		"Filter.1.Name":     "ACS:InstanceInformation.InstanceId",
		"Filter.1.Operator": "Equal",
		"Filter.1.Value.1":  "i-bp1",
		"Filter.1.Value.2":  "i-bp2",
	}
	for key, value := range expected {
		if request[key] != value {
			t.Fatalf("expected %s to be %v, got %v", key, value, request[key])
		}
	}
}

func TestExpandOosInventoryAggregators(t *testing.T) {
	request := make(map[string]interface{})
	expandOosInventoryAggregators(request, []interface{}{"ACS:Application.Name", "ACS:Application.Version"})

	expected := map[string]interface{}{
		"Aggregator.1": "ACS:Application.Name",
		"Aggregator.2": "ACS:Application.Version",
	}
	for key, value := range expected {
		if request[key] != value {
			t.Fatalf("expected %s to be %v, got %v", key, value, request[key])
		}
	}
}

func TestFlattenOosInventoryJsonList(t *testing.T) {
	rawObjects := []interface{}{
		map[string]interface{}{
			"Id": "i-bp1",
			"ACS:InstanceInformation": map[string]interface{}{
				"Content": map[string]interface{}{
					"PlatformName": "ubuntu",
				},
			},
		},
	}

	result, err := flattenOosInventoryJsonList(rawObjects)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("expected 1 result, got %d", len(result))
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal([]byte(result[0].(string)), &decoded); err != nil {
		t.Fatalf("expected valid JSON, got %v", err)
	}
	if decoded["Id"] != "i-bp1" {
		t.Fatalf("expected decoded Id to be i-bp1, got %v", decoded["Id"])
	}
}
