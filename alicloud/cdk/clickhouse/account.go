package clickhouse

import (
	"fmt"
	"log"
	"time"

	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type AccountClient struct {
	AccountDescription string
	AccountName        string
	AccountPassword    string
	DbClusterId        string
	Status             string
	Type               string
}

type ClickhouseService struct {
	client *connectivity.AliyunClient
}

func (s *ClickhouseService) DescribeClickHouseAccount(id string) (object map[string]interface{}, err error) {
	return object, nil
}

func (client AccountClient) GetId() string {
	return fmt.Sprint(client.DbClusterId, ":", client.AccountName)
}

func (client AccountClient) Create(d *schema.ResourceData, meta interface{}) (err error, r AccountClient) {
	sdkClient := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}
	action := "CreateAccount"
	request := make(map[string]interface{})
	conn, err := sdkClient.NewClickhouseClient()
	if err != nil {
		return err, r
	}
	if v, ok := d.GetOk("account_description"); ok {
		request["AccountDescription"] = v
	}
	request["AccountName"] = d.Get("account_name")
	request["AccountPassword"] = d.Get("account_password")
	request["DBClusterId"] = d.Get("db_cluster_id")
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-11-11"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})

	return nil, r
}

func (client AccountClient) Read(d *schema.ResourceData, meta interface{}) (err error, r AccountClient) {
	sdkClient := meta.(*connectivity.AliyunClient)
	clickhouseService := ClickhouseService{sdkClient}
	_, err = clickhouseService.DescribeClickHouseAccount(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_click_house_account clickhouseService.DescribeClickHouseAccount Failed!!! %s", err)
			d.SetId("")
			return nil, r
		}
		return err, r
	}
	return nil, r
}

func (client AccountClient) Update(d *schema.ResourceData, meta interface{}) (err error, r AccountClient) {
	sdkClient := meta.(*connectivity.AliyunClient)
	var response map[string]interface{}

	update := false
	request := make(map[string]interface{})
	if d.HasChange("account_description") {
		update = true
		if v, ok := d.GetOk("account_description"); ok {
			request["AccountDescription"] = v
		}
	}
	if update {
		action := "ModifyAccountDescription"
		conn, err := sdkClient.NewClickhouseClient()
		if err != nil {
			return nil, r
		}
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-11-11"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		if err != nil {
			return err, r
		}
	}
	update = false
	if d.HasChange("account_password") {
		update = true
		request = map[string]interface{}{}
		request["AccountPassword"] = d.Get("account_password")
	}
	if update {
		action := "ResetAccountPassword"
		conn, err := sdkClient.NewClickhouseClient()
		if err != nil {
			return err, r
		}
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
			response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-11-11"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		if err != nil {
			return err, r
		}
	}
	d.Partial(false)
	return nil, r
}

func (client AccountClient) Delete(d *schema.ResourceData, meta interface{}) (err error, r AccountClient) {
	sdkClient := meta.(*connectivity.AliyunClient)

	action := "DeleteAccount"
	var request map[string]interface{}
	var response map[string]interface{}
	conn, err := sdkClient.NewClickhouseClient()
	if err != nil {
		return err, r
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-11-11"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if IsExpectedErrors(err, []string{"IncorrectAccountStatus", "IncorrectDBInstanceState"}) || NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	return nil, r
}

func SetAccountState(d *schema.ResourceData, r AccountClient) error {
	if err := d.Set("account_name", r.AccountName); err != nil {
		return err
	}
	if err := d.Set("db_cluster_id", r.DbClusterId); err != nil {
		return err
	}
	if err := d.Set("account_description", r.AccountDescription); err != nil {
		return err
	}
	if err := d.Set("status", r.Status); err != nil {
		return err
	}
	if err := d.Set("type", r.Type); err != nil {
		return err
	}
	return nil
}

func IsExpectedErrors(err error, strings []string) bool {
	return false
}

func NeedRetry(err error) bool {
	return false
}

func StringPointer(action string) *string {
	return nil
}

func NotFoundError(err error) bool {
	return false
}

func incrementalWait(firstDuration time.Duration, increaseDuration time.Duration) func() {
	retryCount := 1
	return func() {
		var waitTime time.Duration
		if retryCount == 1 {
			waitTime = firstDuration
		} else if retryCount > 1 {
			waitTime += increaseDuration
		}
		time.Sleep(waitTime)
		retryCount++
	}
}
