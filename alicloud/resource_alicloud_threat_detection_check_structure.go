// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudThreatDetectionCheckStructure() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudThreatDetectionCheckStructureCreate,
		Read:   resourceAliCloudThreatDetectionCheckStructureRead,
		Delete: resourceAliCloudThreatDetectionCheckStructureDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{},
	}
}

func resourceAliCloudThreatDetectionCheckStructureCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceAliCloudThreatDetectionCheckStructureRead(d, meta)
}

func resourceAliCloudThreatDetectionCheckStructureRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	threatDetectionServiceV2 := ThreatDetectionServiceV2{client}

	objectRaw, err := threatDetectionServiceV2.DescribeThreatDetectionCheckStructure(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_threat_detection_check_structure DescribeThreatDetectionCheckStructure Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	return nil
}

func resourceAliCloudThreatDetectionCheckStructureDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Check Structure. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
