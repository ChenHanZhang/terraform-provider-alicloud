// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudPolardbDbClusterPerformance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudPolardbDbClusterPerformanceCreate,
		Read:   resourceAliCloudPolardbDbClusterPerformanceRead,
		Delete: resourceAliCloudPolardbDbClusterPerformanceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"db_node_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudPolardbDbClusterPerformanceCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceAliCloudPolardbDbClusterPerformanceRead(d, meta)
}

func resourceAliCloudPolardbDbClusterPerformanceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	polardbServiceV2 := PolardbServiceV2{client}

	objectRaw, err := polardbServiceV2.DescribePolardbDbClusterPerformance(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_polardb_db_cluster_performance DescribePolardbDbClusterPerformance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("db_node_id", objectRaw["DBNodeId"])
	d.Set("end_time", objectRaw["EndTime"])

	objectRaw, err = polardbServiceV2.DescribeDbClusterPerformanceDescribeDBClusterPerformance(d.Id())
	if err != nil && !NotFoundError(err) {
		return WrapError(err)
	}

	d.Set("end_time", objectRaw["EndTime"])
	d.Set("start_time", objectRaw["StartTime"])
	d.Set("db_cluster_id", objectRaw["DBClusterId"])

	return nil
}

func resourceAliCloudPolardbDbClusterPerformanceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Db Cluster Performance. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
