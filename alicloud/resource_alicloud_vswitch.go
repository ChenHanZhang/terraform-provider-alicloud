package alicloud

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAlicloudVswitch() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudVswitchCreate,
		Read:   resourceAlicloudVswitchRead,
		Update: resourceAlicloudVswitchUpdate,
		Delete: resourceAlicloudVswitchDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"all": {
				Optional: true,
				Type:     schema.TypeBool,
			},
			"available_ip_address_count": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"cidr_block": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeString,
			},
			"create_time": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"description": {
				Optional: true,
				Type:     schema.TypeString,
			},
			"enable_ipv6": {
				Optional: true,
				Type:     schema.TypeBool,
			},
			"ipv6_cidr_block": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"ipv6_cidr_block_mask": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_default": {
				Computed: true,
				Type:     schema.TypeBool,
			},
			"network_acl_id": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"resource_group_id": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"resource_type": {
				Optional: true,
				Type:     schema.TypeString,
			},
			"route_table_id": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"status": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"tags": tagsSchema(),
			"vswitch_id": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeString,
			},
			"vswitch_name": {
				Optional:      true,
				Computed:      true,
				Type:          schema.TypeString,
				ConflictsWith: []string{"name"},
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"vswitch_name"},
				Deprecated:    "Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.",
			},
			"vpc_id": {
				Required: true,
				ForceNew: true,
				Type:     schema.TypeString,
			},
			"zone_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"availability_zone"},
			},
			"availability_zone": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ForceNew:      true,
				ConflictsWith: []string{"zone_id"},
				Deprecated:    "Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.",
			},
		},
	}
}

func resourceAlicloudVswitchCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcService := VpcService{client}
	request := map[string]interface{}{
		"RegionId": client.RegionId,
	}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}

	if v, ok := d.GetOk("cidr_block"); ok {
		request["CidrBlock"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("ipv6_cidr_block_mask"); ok {
		request["Ipv6CidrBlock"] = v
	}
	if v, ok := d.GetOk("vswitch_name"); ok {
		request["VSwitchName"] = v
	} else if v, ok := d.GetOk("name"); ok {
		request["VSwitchName"] = v
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		request["VpcId"] = v
	}
	if v, ok := d.GetOk("vpc_ipv6_cidr_block"); ok {
		request["VpcIpv6CidrBlock"] = v
	}
	if v, ok := d.GetOk("zone_id"); ok {
		request["ZoneId"] = v
	} else if v, ok := d.GetOk("availability_zone"); ok {
		request["ZoneId"] = v
	} else {
		return WrapError(Error(`[ERROR] Argument "availability_zone" or "zone_id" must be set one!`))
	}

	var response map[string]interface{}
	action := "CreateVSwitch"
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(client.GetRetryTimeout(d.Timeout(schema.TimeoutCreate)), func() *resource.RetryError {
		request["ClientToken"] = buildClientToken("CreateVSwitch")
		resp, err := conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &runtime)
		if err != nil {
			if NeedRetry(err) || IsExpectedErrors(err, []string{"CreateVSwitch.IncorrectStatus.cbnStatus", "IncorrectStatus.%s", "IncorrectStatus.cbnStatus", "IncorrectVSwitchStatus", "InvalidStatus.RouteEntry", "OperationConflict", "OperationFailed.DistibuteLock", "OperationFailed.IdempotentTokenProcessing", "OperationFailed.NotifyCenCreate", "TaskConflict"}) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		response = resp
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_vswitch", action, AlibabaCloudSdkGoERROR)
	}

	if v, err := jsonpath.Get("$.VSwitchId", response); err != nil || v == nil {
		return WrapErrorf(err, IdMsg, "alicloud_vswitch")
	} else {
		d.SetId(fmt.Sprint(v))
	}
	stateConf := BuildStateConf([]string{}, []string{"Available"}, d.Timeout(schema.TimeoutCreate), 5*time.Second, vpcService.VswitchStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return resourceAlicloudVswitchUpdate(d, meta)
}

func resourceAlicloudVswitchRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcService := VpcService{client}

	object, err := vpcService.DescribeVswitch(d.Id())
	if err != nil {
		if NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_vswitch vpcService.DescribeVswitch Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	availableIpAddressCount47 := object["AvailableIpAddressCount"]
	d.Set("available_ip_address_count", availableIpAddressCount47)

	cidrBlock92 := object["CidrBlock"]
	d.Set("cidr_block", cidrBlock92)

	creationTime88 := object["CreationTime"]
	d.Set("create_time", creationTime88)

	description11 := object["Description"]
	d.Set("description", description11)

	ipv6CidrBlock3 := object["Ipv6CidrBlock"]
	d.Set("ipv6_cidr_block", ipv6CidrBlock3)

	if v, ok := object["Ipv6CidrBlock"]; ok && fmt.Sprint(v) != "" {
		_, cidrBlock := GetIPv6SubnetAddr(v.(string))
		d.Set("ipv6_cidr_block_mask", cidrBlock)
	}

	isDefault88 := object["IsDefault"]
	d.Set("is_default", isDefault88)

	networkAclId18 := object["NetworkAclId"]
	d.Set("network_acl_id", networkAclId18)

	resourceGroupId56 := object["ResourceGroupId"]
	d.Set("resource_group_id", resourceGroupId56)
	routeTableRouteTableId19, _ := jsonpath.Get("$.RouteTable.RouteTableId", object)
	d.Set("route_table_id", routeTableRouteTableId19)

	status7 := object["Status"]
	d.Set("status", status7)
	tagsMap := make(map[string]interface{})
	tagsRaw, _ := jsonpath.Get("$.Tag", object)
	if tagsRaw != nil {
		for _, value0 := range tagsRaw.([]interface{}) {
			tags := value0.(map[string]interface{})
			key := tags["Key"].(string)
			value := tags["Value"]
			if !ignoredTags(key, value) {
				tagsMap[key] = value
			}
		}
	}
	if len(tagsMap) > 0 {
		d.Set("tags", tagsMap)
	}

	vSwitchName57 := object["VSwitchName"]
	d.Set("vswitch_name", vSwitchName57)
	d.Set("name", vSwitchName57)

	vpcId52 := object["VpcId"]
	d.Set("vpc_id", vpcId52)

	zoneId75 := object["ZoneId"]
	d.Set("zone_id", zoneId75)
	d.Set("availability_zone", zoneId75)

	return nil
}

func resourceAlicloudVswitchUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}
	vpcService := VpcService{client}
	d.Partial(true)
	update := false
	request := map[string]interface{}{
		"VSwitchId": d.Id(),
		"RegionId":  client.RegionId,
	}

	if !d.IsNewResource() && d.HasChange("description") {
		update = true
		if v, ok := d.GetOk("description"); ok {
			request["Description"] = v
		}
	}
	if v, ok := d.GetOk("enable_ipv6"); ok {
		request["EnableIPv6"] = v
	}
	if !d.IsNewResource() && d.HasChange("ipv6_cidr_block") {
		update = true
		if v, ok := d.GetOk("ipv6_cidr_block"); ok {
			request["Ipv6CidrBlock"] = v
		}
	}

	if !d.IsNewResource() && d.HasChange("ipv6_cidr_block_mask") {
		update = true
		request["Ipv6CidrBlock"] = d.Get("ipv6_cidr_block_mask")
	}

	if !d.IsNewResource() && d.HasChange("vswitch_name") {
		update = true
		if v, ok := d.GetOk("vswitch_name"); ok {
			request["VSwitchName"] = v
		}
	}
	if !d.IsNewResource() && d.HasChange("name") {
		update = true
		request["VSwitchName"] = d.Get("name")
	}

	if v, ok := d.GetOk("vpc_ipv6_cidr_block"); ok {
		request["VpcIpv6CidrBlock"] = v
	}

	if update {
		action := "ModifyVSwitchAttribute"
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(client.GetRetryTimeout(d.Timeout(schema.TimeoutUpdate)), func() *resource.RetryError {
			resp, err := conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
			if err != nil {
				if NeedRetry(err) || IsExpectedErrors(err, []string{"IncorrectStatus.VSwitch", "IncorrectStatus.VpcRouteEntry", "OperationConflict", "OperationFailed.LastTokenProcessing", "ServiceUnavailable"}) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, resp, request)
			return nil
		})
		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
		}
		d.SetPartial("description")
		d.SetPartial("enable_ipv6")
		d.SetPartial("ipv6_cidr_block")
		d.SetPartial("vswitch_name")
		d.SetPartial("vpc_ipv6_cidr_block")
	}

	if d.HasChange("tags") {
		if err := vpcService.SetResourceTags(d, "VSWITCH"); err != nil {
			return WrapError(err)
		}
		d.SetPartial("tags")
	}

	d.Partial(false)
	return resourceAlicloudVswitchRead(d, meta)
}

func resourceAlicloudVswitchDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	vpcService := VpcService{client}
	conn, err := client.NewVpcClient()
	if err != nil {
		return WrapError(err)
	}

	request := map[string]interface{}{

		"VSwitchId": d.Id(),
		"RegionId":  client.RegionId,
	}

	action := "DeleteVSwitch"
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(client.GetRetryTimeout(d.Timeout(schema.TimeoutDelete)), func() *resource.RetryError {
		resp, err := conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2016-04-28"), StringPointer("AK"), nil, request, &util.RuntimeOptions{})
		if err != nil {
			if NeedRetry(err) || IsExpectedErrors(err, []string{"DeleteVSwitch.IncorrectStatus.cbnStatus", "DependencyViolation", "DependencyViolation.EnhancedNatgw", "DependencyViolation.HaVip", "DependencyViolation.MulticastDomain", "DependencyViolation.NetworkAcl", "DependencyViolation.RouteTable", "DependencyViolation.SNAT", "DependencyViolation.SnatEntry", "IncorrectRouteEntryStatus", "IncorrectVSwitchStatus", "InternalError", "LastTokenProcessing", "OperationConflict", "OperationDenied.OtherSubnetProcessing", "SystemBusy", "TaskConflict"}) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, resp, request)
		return nil
	})
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidVSwitchId.NotFound", "InvalidVswitchID.NotFound"}) || NotFoundError(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutDelete), 5*time.Second, vpcService.VswitchStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}
	return nil
}

func GetIPv6SubnetAddr(ipAddr string) (string, int) {
	// Split the IP address and subnet prefix length
	ip, prefix, err := net.ParseCIDR(ipAddr)
	if err != nil {
		return "", 0
	}
	mask, _ := strconv.Atoi(strings.Split(ipAddr, "/")[1])
	// Get the network address by masking the IP address with the subnet prefix
	netAddr := ip.Mask(prefix.Mask)
	// Convert the network address to a string
	netAddrStr := netAddr.String()
	// Get the last8 bits of the address
	last8Bits := netAddr[mask/8-1]
	// Convert the last8 bits to an integer
	return netAddrStr, int(last8Bits)
}
