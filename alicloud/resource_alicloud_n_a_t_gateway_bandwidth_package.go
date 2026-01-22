// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
package alicloud

import (
	"log"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudNATGatewayBandwidthPackage() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudNATGatewayBandwidthPackageCreate,
		Read:   resourceAliCloudNATGatewayBandwidthPackageRead,
		Delete: resourceAliCloudNATGatewayBandwidthPackageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bandwidth_package_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"payment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"region_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAliCloudNATGatewayBandwidthPackageCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceAliCloudNATGatewayBandwidthPackageRead(d, meta)
}

func resourceAliCloudNATGatewayBandwidthPackageRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	nATGatewayServiceV2 := NATGatewayServiceV2{client}

	objectRaw, err := nATGatewayServiceV2.DescribeNATGatewayBandwidthPackage(d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_n_a_t_gateway_bandwidth_package DescribeNATGatewayBandwidthPackage Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("region_id", objectRaw["RegionId"])
	d.Set("status", objectRaw["Status"])
	d.Set("bandwidth_package_id", objectRaw["BandwidthPackageId"])

	return nil
}

func resourceAliCloudNATGatewayBandwidthPackageDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] Cannot destroy resource AliCloud Resource Bandwidth Package. Terraform will remove this resource from the state file, however resources may remain.")
	return nil
}
