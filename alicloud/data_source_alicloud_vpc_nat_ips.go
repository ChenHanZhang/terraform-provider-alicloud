package alicloud

import (
	"fmt"
	"regexp"
	"time"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceAlicloudVpcNatIps() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudVpcNatIpsRead,
		Schema: map[string]*schema.Schema{
			"nat_gateway_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"nat_ip_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ids": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.ValidateRegexp,
				ForceNew:     true,
			},
			"names": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"nat_ip_name": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"nat_ip_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Available", "Creating", "Deleting"}, false),
			},
			"output_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_default": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"nat_gateway_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nat_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nat_ip_cidr": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nat_ip_description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nat_ip_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"nat_ip_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAlicloudVpcNatIpsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	action := "ListNatIps"
	request := make(map[string]interface{})
	request["NatGatewayId"] = d.Get("nat_gateway_id")
	if v, ok := d.GetOk("nat_ip_cidr"); ok {
		request["NatIpCidr"] = v
	}
	if v, ok := d.GetOk("nat_ip_name"); ok {
		request["NatIpName"] = v
	}
	if v, ok := d.GetOk("nat_ip_ids"); ok {
		request["NatIpIds"] = v
	}
	if v, ok := d.GetOk("status"); ok {
		request["NatIpStatus"] = v
	}
	request["RegionId"] = client.RegionId
	if v, ok := d.GetOk("status"); ok {
		request["NatIpStatus"] = v
	}
	request["MaxResults"] = PageSizeLarge
	var objects []map[string]interface{}
	var natIpNameRegex *regexp.Regexp
	if v, ok := d.GetOk("name_regex"); ok {
		r, err := regexp.Compile(v.(string))
		if err != nil {
			return WrapError(err)
		}
		natIpNameRegex = r
	}

	idsMap := make(map[string]string)
	if v, ok := d.GetOk("ids"); ok {
		for _, vv := range v.([]interface{}) {
			if vv == nil {
				continue
			}
			idsMap[vv.(string)] = vv.(string)
		}
	}
	var response map[string]interface{}
	var err error
	for {
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			request["ClientToken"] = buildClientToken("ListNatIps")
			response, err = client.RpcPost("Vpc", "2016-04-28", action, nil, request, true)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			return nil
		})
		addDebug(action, response, request)
		if err != nil {
			return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_vpc_nat_ips", action, AlibabaCloudSdkGoERROR)
		}
		resp, err := jsonpath.Get("$.NatIps", response)
		if err != nil {
			return WrapErrorf(err, FailedGetAttributeMsg, action, "$.NatIps", response)
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})
			if natIpNameRegex != nil && !natIpNameRegex.MatchString(fmt.Sprint(item["NatIpName"])) {
				continue
			}
			if len(idsMap) > 0 {
				if _, ok := idsMap[fmt.Sprintf("%s:%s", item["NatGatewayId"], item["NatIpId"])]; !ok {
					continue
				}
			}
			objects = append(objects, item)
		}
		if nextToken, ok := response["NextToken"].(string); ok && nextToken != "" {
			request["NextToken"] = nextToken
		} else {
			break
		}
	}
	ids := make([]string, 0)
	names := make([]interface{}, 0)
	s := make([]map[string]interface{}, 0)
	for _, object := range objects {
		mapping := map[string]interface{}{
			"is_default":         object["IsDefault"],
			"nat_gateway_id":     object["NatGatewayId"],
			"nat_ip":             object["NatIp"],
			"nat_ip_cidr":        object["NatIpCidr"],
			"nat_ip_description": object["NatIpDescription"],
			"id":                 fmt.Sprintf("%s:%s", object["NatGatewayId"], object["NatIpId"]),
			"nat_ip_id":          fmt.Sprint(object["NatIpId"]),
			"nat_ip_name":        object["NatIpName"],
			"status":             object["NatIpStatus"],
		}
		ids = append(ids, fmt.Sprint(mapping["id"]))
		names = append(names, object["NatIpName"])
		s = append(s, mapping)
	}

	d.SetId(dataResourceIdHash(ids))
	if err := d.Set("ids", ids); err != nil {
		return WrapError(err)
	}

	if err := d.Set("names", names); err != nil {
		return WrapError(err)
	}

	if err := d.Set("ips", s); err != nil {
		return WrapError(err)
	}
	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), s)
	}

	return nil
}
