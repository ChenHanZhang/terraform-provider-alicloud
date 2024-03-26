package alicloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
)

func TestAccAlicloudMarketProductsDataSource(t *testing.T) {
	rand := acctest.RandIntRange(1000000, 9999999)
	resourceId := "data.alicloud_market_products.default"

	testAccConfig := dataSourceTestAccConfigFunc(resourceId,
		"",
		dataSourceMarketProductsConfigDependence)

	nameRegexConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"name_regex":   "BatchCompute",
			"product_type": "MIRROR",
		}),
		fakeConfig: testAccConfig(map[string]interface{}{
			"name_regex":   "BatchCompute_fake",
			"product_type": "MIRROR",
		}),
	}
	//idsConf := dataSourceTestAccConfig{
	//	existConfig: testAccConfig(map[string]interface{}{
	//		"ids":          []string{"cmjj022644"},
	//		"product_type": "MIRROR",
	//	}),
	//	fakeConfig: testAccConfig(map[string]interface{}{
	//		"ids":          []string{"cmjj022644_fake"},
	//		"product_type": "MIRROR",
	//	}),
	//}
	sortConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"sort":         "user_count-desc",
			"product_type": "MIRROR",
		}),
	}
	categoryIdConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"category_id": "53366009",
		}),
	}
	productTypeConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"product_type": "MIRROR",
		}),
	}
	searchTermConf := dataSourceTestAccConfig{
		existConfig: testAccConfig(map[string]interface{}{
			"search_term": "Image",
		}),
	}

	//allConf := dataSourceTestAccConfig{
	//	existConfig: testAccConfig(map[string]interface{}{
	//		"sort":         "created_on-desc",
	//		"category_id":  "56024006",
	//		"product_type": "MIRROR",
	//		"name_regex":   "BatchCompute",
	//		"search_term":  "Image",
	//	}),
	//	fakeConfig: testAccConfig(map[string]interface{}{
	//		"sort":         "created_on-desc",
	//		"category_id":  "56024006",
	//		"product_type": "MIRROR",
	//		"name_regex":   "BatchCompute_fake",
	//		"search_term":  "Image",
	//	}),
	//}

	var existMarketProductsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":                        CHECKSET,
			"ids.0":                        CHECKSET,
			"products.#":                   CHECKSET,
			"products.0.category_id":       CHECKSET,
			"products.0.code":              CHECKSET,
			"products.0.delivery_way":      CHECKSET,
			"products.0.image_url":         CHECKSET,
			"products.0.name":              CHECKSET,
			"products.0.score":             CHECKSET,
			"products.0.short_description": CHECKSET,
			"products.0.supplier_id":       CHECKSET,
			"products.0.supplier_name":     CHECKSET,
			"products.0.target_url":        CHECKSET,
		}
	}

	var fakeMarketProductsMapFunc = func(rand int) map[string]string {
		return map[string]string{
			"ids.#":      "0",
			"products.#": "0",
		}
	}

	var pvtzZoneRecordsCheckInfo = dataSourceAttr{
		resourceId:   resourceId,
		existMapFunc: existMarketProductsMapFunc,
		fakeMapFunc:  fakeMarketProductsMapFunc,
	}

	pvtzZoneRecordsCheckInfo.dataSourceTestCheck(t, rand, nameRegexConf, sortConf, categoryIdConf, productTypeConf, searchTermConf)
}

func dataSourceMarketProductsConfigDependence(name string) string {
	return ""
}
