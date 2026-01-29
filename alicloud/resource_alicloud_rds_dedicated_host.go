// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudRdsDedicatedHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudRdsDedicatedHostCreate,
		Read:   resourceAliCloudRdsDedicatedHostRead,
		Delete: resourceAliCloudRdsDedicatedHostDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"dedicated_host_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dedicated_host_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudRdsDedicatedHostCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceAliCloudRdsDedicatedHostRead(d, meta)
}

func resourceAliCloudRdsDedicatedHostRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	rdsServiceV2 := RdsServiceV2{client}

	objectRaw, err := rdsServiceV2.DescribeRdsDedicatedHost(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_rds_dedicated_host DescribeRdsDedicatedHost Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("dedicated_host_group_id", objectRaw["DedicatedHostGroupId"])

	dedicatedHostsRawArrayObj, _ := jsonpath.Get("$.DedicatedHosts.DedicatedHosts[*]", objectRaw)
	dedicatedHostsRawArray := make([]interface{}, 0)
	if dedicatedHostsRawArrayObj != nil {
		dedicatedHostsRawArray = convertToInterfaceArray(dedicatedHostsRawArrayObj)
	}
	dedicatedHostsRaw := make(map[string]interface{})
	if len(dedicatedHostsRawArray) > 0 {
		dedicatedHostsRaw = dedicatedHostsRawArray[0].(map[string]interface{})
	}

	d.Set("dedicated_host_id", dedicatedHostsRaw["DedicatedHostId"])

	return nil
}

func resourceAliCloudRdsDedicatedHostDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Dedicated Host. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
