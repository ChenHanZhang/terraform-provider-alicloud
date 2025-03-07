package alicloud

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAlicloudClickHouseAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudClickHouseAccountCreate,
		Read:   resourceAlicloudClickHouseAccountRead,
		Update: resourceAlicloudClickHouseAccountUpdate,
		Delete: resourceAlicloudClickHouseAccountDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
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
				ValidateFunc: StringMatch(regexp.MustCompile(`^[a-z][a-z0-9_]{1,15}`), "The account_name most consist of lowercase letters, numbers, and underscores, starting with a lowercase letter"),
			},
			"account_password": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringMatch(regexp.MustCompile(`[a-zA-Z!#$%^&*()_+-=]{8,32}`), "account_password must consist of uppercase letters, lowercase letters, numbers, and special characters"),
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
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: StringInSlice([]string{"Normal", "Super"}, false),
			},
			"dml_authority": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: StringInSlice([]string{"all", "readOnly,modify"}, false),
			},
			"ddl_authority": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allow_databases": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_databases": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "Field 'total_databases' has been deprecated from version 1.223.1 and it will be removed in the future version.",
			},
			"allow_dictionaries": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_dictionaries": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: "Field 'total_dictionaries' has been deprecated from version 1.223.1 and it will be removed in the future version.",
			},
		},
	}
}

func resourceAlicloudClickHouseAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "CreateAccount"
	request := make(map[string]interface{})
	var err error
	if v, ok := d.GetOk("account_description"); ok {
		request["AccountDescription"] = v
	}
	request["AccountName"] = d.Get("account_name")
	request["AccountPassword"] = d.Get("account_password")
	request["DBClusterId"] = d.Get("db_cluster_id")
	if d.Get("type") == "Super" {
		action = "CreateSQLAccount"
		request["AccountType"] = d.Get("type")
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = client.RpcPost("clickhouse", "2019-11-11", action, nil, request, false)
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_click_house_account", action, AlibabaCloudSdkGoERROR)
	}

	d.SetId(fmt.Sprint(request["DBClusterId"], ":", request["AccountName"]))

	return resourceAlicloudClickHouseAccountUpdate(d, meta)
}
func resourceAlicloudClickHouseAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	clickhouseService := ClickhouseService{client}
	object, err := clickhouseService.DescribeClickHouseAccount(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_click_house_account clickhouseService.DescribeClickHouseAccount Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}
	d.Set("account_name", parts[1])
	d.Set("db_cluster_id", parts[0])
	d.Set("account_description", object["AccountDescription"])
	d.Set("status", object["AccountStatus"])
	d.Set("type", object["AccountType"])

	authority, err := clickhouseService.DescribeClickHouseAccountAuthority(d.Id())
	d.Set("dml_authority", authority["DmlAuthority"])
	d.Set("ddl_authority", authority["DdlAuthority"])

	d.Set("allow_databases", convertArrayToString(authority["AllowDatabases"], ","))
	d.Set("allow_dictionaries", convertArrayToString(authority["AllowDictionaries"], ","))
	d.Set("total_databases", convertArrayToString(authority["TotalDatabases"], ","))
	d.Set("total_dictionaries", convertArrayToString(authority["TotalDictionaries"], ","))
	if err != nil {
		return WrapError(err)
	}

	return nil
}
func resourceAlicloudClickHouseAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}

	update := false
	d.Partial(true)
	request := map[string]interface{}{
		"AccountName": parts[1],
		"DBClusterId": parts[0],
	}
	if !d.IsNewResource() && d.HasChange("account_description") {
		update = true
	}
	if v, ok := d.GetOk("account_description"); ok {
		request["AccountDescription"] = v
	}
	if update {
		action := "ModifyAccountDescription"
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("clickhouse", "2019-11-11", action, nil, request, false)
			if err != nil {
				if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("account_description")
	}
	update = false
	request = map[string]interface{}{
		"AccountName": parts[1],
		"DBClusterId": parts[0],
	}
	if !d.IsNewResource() && d.HasChange("account_password") {
		update = true
	}
	request["AccountPassword"] = d.Get("account_password")
	if update {
		action := "ResetAccountPassword"
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("clickhouse", "2019-11-11", action, nil, request, false)
			if err != nil {
				if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("account_password")
	}

	update = false
	request = map[string]interface{}{
		"AccountName": parts[1],
		"DBClusterId": parts[0],
	}
	request["RegionId"] = client.RegionId
	if d.HasChange("dml_authority") {
		update = true
	}
	if v, ok := d.GetOk("dml_authority"); ok {
		request["DmlAuthority"] = v
	}

	if d.HasChange("ddl_authority") {
		update = true
	}
	if v, ok := d.GetOkExists("ddl_authority"); ok {
		request["DdlAuthority"] = v
	}

	if d.HasChange("allow_databases") {
		update = true
	}
	if v, ok := d.GetOk("allow_databases"); ok {
		request["AllowDatabases"] = v
	}

	if d.HasChange("total_databases") {
		update = true
	}
	if v, ok := d.GetOk("total_databases"); ok {
		request["TotalDatabases"] = v
	}

	if d.HasChange("allow_dictionaries") {
		update = true
	}
	if v, ok := d.GetOk("allow_dictionaries"); ok {
		request["AllowDictionaries"] = v
	}

	if d.HasChange("total_dictionaries") {
		update = true
	}
	if v, ok := d.GetOk("total_dictionaries"); ok {
		request["TotalDictionaries"] = v
	}

	if update {
		action := "ModifyAccountAuthority"
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = client.RpcPost("clickhouse", "2019-11-11", action, nil, request, false)
			if err != nil {
				if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("account_password")
	}
	d.Partial(false)

	return resourceAlicloudClickHouseAccountRead(d, meta)
}
func resourceAlicloudClickHouseAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	clickhouseService := ClickhouseService{client}
	parts, err := ParseResourceId(d.Id(), 2)
	if err != nil {
		return WrapError(err)
	}
	action := "DeleteAccount"
	var response map[string]interface{}
	request := map[string]interface{}{
		"AccountName": parts[1],
		"DBClusterId": parts[0],
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = client.RpcPost("clickhouse", "2019-11-11", action, nil, request, false)
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(action, response, request)
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}
	stateConf := BuildStateConf([]string{"Deleting"}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, clickhouseService.ClickhouseStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		if NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}
