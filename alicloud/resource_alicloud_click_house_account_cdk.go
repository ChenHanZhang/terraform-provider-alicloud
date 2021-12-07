package alicloud

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/cdk/clickhouse"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAlicloudClickHouseAccountCDK() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudClickHouseAccountCdkCreate,
		Read:   resourceAlicloudClickHouseAccountCdkRead,
		Update: resourceAlicloudClickHouseAccountCdkUpdate,
		Delete: resourceAlicloudClickHouseAccountCdkDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"account_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"account_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-z][a-z0-9_]{1,15}`), "The account_name most consist of lowercase letters, numbers, and underscores, starting with a lowercase letter"),
			},
			"account_password": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[a-zA-Z!#$%^&*()_+-=]{8,32}`), "account_password must consist of uppercase letters, lowercase letters, numbers, and special characters"),
			},
			"db_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAlicloudClickHouseAccountCdkCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(clickhouse.AccountClient)

	err, r := client.Create(d, meta)
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_click_house_account", AlibabaCloudSdkGoERROR)
	}
	d.SetId(r.GetId())
	return resourceAlicloudClickHouseAccountCdkRead(d, meta)
}

func resourceAlicloudClickHouseAccountCdkRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(clickhouse.AccountClient)

	err, r := client.Read(d, meta)
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_click_house_account", AlibabaCloudSdkGoERROR)
	}
	if err = clickhouse.SetAccountState(d, r); err != nil {
		return WrapError(err)
	}
	return nil
}

func resourceAlicloudClickHouseAccountCdkUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(clickhouse.AccountClient)

	err, r := client.Update(d, meta)
	_ = clickhouse.SetAccountState(d, r)

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_click_house_account", AlibabaCloudSdkGoERROR)
	}
	return resourceAlicloudClickHouseAccountCdkRead(d, meta)
}

func resourceAlicloudClickHouseAccountCdkDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(clickhouse.AccountClient)
	err, _ := client.Delete(d, meta)
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_click_house_account", AlibabaCloudSdkGoERROR)
	}
	return nil
}
