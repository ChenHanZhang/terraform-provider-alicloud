package alicloud

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"time"

	"strconv"

	"github.com/PaesslerAG/jsonpath"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type EcsService struct {
	client *connectivity.AliyunClient
}

func (s *EcsService) JudgeRegionValidation(key, region string) error {
	request := ecs.CreateDescribeRegionsRequest()
	request.RegionId = s.client.RegionId
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeRegions(request)
	})
	if err != nil {
		return fmt.Errorf("DescribeRegions got an error: %#v", err)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	resp, _ := raw.(*ecs.DescribeRegionsResponse)
	if resp == nil || len(resp.Regions.Region) < 1 {
		return GetNotFoundErrorFromString("There is no any available region.")
	}

	var rs []string
	for _, v := range resp.Regions.Region {
		if v.RegionId == region {
			return nil
		}
		rs = append(rs, v.RegionId)
	}
	return fmt.Errorf("'%s' is invalid. Expected on %v.", key, strings.Join(rs, ", "))
}

// DescribeZone validate zoneId is valid in region
func (s *EcsService) DescribeZone(id string) (zone map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeZones"
	request := map[string]interface{}{
		"RegionId": s.client.RegionId,
	}
	response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		return zone, err
	}
	addDebug(action, response, request)
	v, err := jsonpath.Get("$.Zones.Zone", response)
	if err != nil {
		return zone, WrapErrorf(err, FailedGetAttributeMsg, id, "$", response)
	}

	if len(v.([]interface{})) < 1 {
		return zone, WrapError(Error("There is no any availability zone in region %s.", s.client.RegionId))
	}

	zoneIds := []string{}
	for _, z := range v.([]interface{}) {
		tmp := z.(map[string]interface{})
		if tmp["ZoneId"].(string) == id {
			return tmp, nil
		}
		zoneIds = append(zoneIds, tmp["ZoneId"].(string))
	}
	return zone, WrapError(Error("availability_zone %s not exists in region %s, all zones are %s", id, s.client.RegionId, zoneIds))
}

func (s *EcsService) DescribeZones(d *schema.ResourceData) (zones []ecs.Zone, err error) {
	request := ecs.CreateDescribeZonesRequest()
	request.RegionId = s.client.RegionId
	request.InstanceChargeType = d.Get("instance_charge_type").(string)
	request.SpotStrategy = d.Get("spot_strategy").(string)
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeZones(request)
	})
	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, "alicloud_instance_type_families", request.GetActionName(), AlibabaCloudSdkGoERROR)
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response, _ := raw.(*ecs.DescribeZonesResponse)
	if len(response.Zones.Zone) < 1 {
		return zones, WrapError(Error("There is no any availability zone in region %s.", s.client.RegionId))
	}
	if v, ok := d.GetOk("zone_id"); ok {
		zoneIds := []string{}
		for _, z := range response.Zones.Zone {
			if z.ZoneId == v.(string) {
				return []ecs.Zone{z}, nil
			}
			zoneIds = append(zoneIds, z.ZoneId)
		}
		return zones, WrapError(Error("availability_zone %s not exists in region %s, all zones are %s", v.(string), s.client.RegionId, zoneIds))
	} else {
		return response.Zones.Zone, nil
	}
}

func (s *EcsService) DescribeInstance(id string) (instance ecs.Instance, err error) {
	request := ecs.CreateDescribeInstancesRequest()
	request.RegionId = s.client.RegionId
	request.InstanceIds = convertListToJsonString([]interface{}{id})
	request.AdditionalAttributes = &[]string{"META_OPTIONS", "NETWORK_PRIMARY_ENI_IP", "LOGIN_AS_NON_ROOT"}

	var response *ecs.DescribeInstancesResponse
	wait := incrementalWait(1*time.Second, 1*time.Second)
	err = resource.Retry(10*time.Minute, func() *resource.RetryError {
		raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.DescribeInstances(request)
		})
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)

			}
			return resource.NonRetryableError(err)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		response, _ = raw.(*ecs.DescribeInstancesResponse)
		return nil
	})

	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
		return
	}
	if len(response.Instances.Instance) < 1 {
		return instance, WrapErrorf(NotFoundErr("Instance", id), NotFoundMsg, ProviderERROR, response.RequestId)
	}

	return response.Instances.Instance[0], nil
}

func (s *EcsService) DescribeInstanceAttribute(id string) (instance ecs.DescribeInstanceAttributeResponse, err error) {
	request := ecs.CreateDescribeInstanceAttributeRequest()
	request.InstanceId = id
	request.RegionId = s.client.RegionId
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeInstanceAttribute(request)
	})
	if err != nil {
		return instance, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response, _ := raw.(*ecs.DescribeInstanceAttributeResponse)
	if response.InstanceId != id {
		return instance, WrapErrorf(NotFoundErr("Instance", id), NotFoundMsg, ProviderERROR, response.RequestId)
	}

	return *response, nil
}

func (s *EcsService) DescribeInstanceSystemDisk(instanceId, rg, diskId string) (disk ecs.Disk, err error) {
	request := ecs.CreateDescribeDisksRequest()
	if diskId != "" {
		request.DiskIds = convertListToJsonString([]interface{}{diskId})
	} else {
		request.InstanceId = instanceId
		request.DiskType = string(DiskTypeSystem)
		// resource_group_id may cause failure to query the system disk of the instance, because the newly created instance may fail to query through the resource_group_id parameter, so temporarily remove this parameter.
		if rg != "" {
			request.ResourceGroupId = rg
		}
	}
	request.RegionId = s.client.RegionId
	var response *ecs.DescribeDisksResponse
	wait := incrementalWait(1*time.Second, 1*time.Second)
	err = resource.Retry(10*time.Minute, func() *resource.RetryError {
		raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.DescribeDisks(request)
		})
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)

			}
			return resource.NonRetryableError(err)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		response, _ = raw.(*ecs.DescribeDisksResponse)
		return nil
	})
	if err != nil {
		return disk, WrapErrorf(err, DefaultErrorMsg, instanceId, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	if len(response.Disks.Disk) < 1 {
		return disk, WrapErrorf(NotFoundErr("Instance", instanceId+" system disk"), NotFoundWithResponse, response)
	}
	return response.Disks.Disk[0], nil
}

// ResourceAvailable check resource available for zone
func (s *EcsService) ResourceAvailable(zone ecs.Zone, resourceType ResourceType) error {
	for _, res := range zone.AvailableResourceCreation.ResourceTypes {
		if res == string(resourceType) {
			return nil
		}
	}
	return WrapError(Error("%s is not available in %s zone of %s region", resourceType, zone.ZoneId, s.client.Region))
}

func (s *EcsService) DiskAvailable(zone map[string]interface{}, diskCategory DiskCategory) error {
	for _, disk := range zone["AvailableDiskCategories"].(map[string]interface{})["DiskCategories"].([]interface{}) {
		if disk.(string) == string(diskCategory) {
			return nil
		}
	}
	return WrapError(Error("%s is not available in %s zone of %s region", diskCategory, zone["ZoneId"], s.client.Region))
}

func (s *EcsService) JoinSecurityGroups(instanceId string, securityGroupIds []string) error {
	request := ecs.CreateJoinSecurityGroupRequest()
	request.InstanceId = instanceId
	request.RegionId = s.client.RegionId
	for _, sid := range securityGroupIds {
		request.SecurityGroupId = sid
		raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.JoinSecurityGroup(request)
		})
		if err != nil && IsExpectedErrors(err, []string{"InvalidInstanceId.AlreadyExists"}) {
			return WrapErrorf(err, DefaultErrorMsg, instanceId, request.GetActionName(), AlibabaCloudSdkGoERROR)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	}

	return nil
}

func (s *EcsService) LeaveSecurityGroups(instanceId string, securityGroupIds []string) error {
	request := ecs.CreateLeaveSecurityGroupRequest()
	request.InstanceId = instanceId
	request.RegionId = s.client.RegionId
	for _, sid := range securityGroupIds {
		request.SecurityGroupId = sid
		raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.LeaveSecurityGroup(request)
		})
		if err != nil && IsExpectedErrors(err, []string{"InvalidSecurityGroupId.NotFound"}) {
			return WrapErrorf(err, DefaultErrorMsg, instanceId, request.GetActionName(), AlibabaCloudSdkGoERROR)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	}

	return nil
}

func (s *EcsService) DescribeSecurityGroup(id string) (group ecs.DescribeSecurityGroupAttributeResponse, err error) {
	request := ecs.CreateDescribeSecurityGroupAttributeRequest()
	request.SecurityGroupId = id
	request.RegionId = s.client.RegionId

	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeSecurityGroupAttribute(request)
	})
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)

	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidSecurityGroupId.NotFound"}) {
			err = WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return
	}

	response, _ := raw.(*ecs.DescribeSecurityGroupAttributeResponse)
	if response.SecurityGroupId != id {
		err = WrapErrorf(NotFoundErr("Security Group", id), NotFoundMsg, ProviderERROR, response.RequestId)
		return
	}

	return *response, nil
}

func (s *EcsService) DescribeSecurityGroupAttribute(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeSecurityGroupAttribute"
	request := map[string]interface{}{
		"RegionId":        s.client.RegionId,
		"SecurityGroupId": id,
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidSecurityGroupId.NotFound"}) {
			return object, WrapErrorf(NotFoundErr("ECS:SecurityGroup", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$", response)
	}
	if v.(map[string]interface{})["SecurityGroupId"] != id {
		err = WrapErrorf(NotFoundErr("Security Group", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		return
	}
	object = v.(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeSecurityGroupRule(id string) (rule ecs.Permission, err error) {
	parts, err := ParseResourceId(id, 8)
	if err != nil {
		return rule, WrapError(err)
	}

	groupId, direction, ipProtocol, portRange, nicType, cidrIp, policy := parts[0], parts[1], parts[2], parts[3], parts[4], parts[5], parts[6]
	cidrIp, err = compressIPv6OrCIDR(strings.Replace(cidrIp, "_", ":", -1))
	if err != nil {
		return rule, WrapError(err)
	}

	priority, err := strconv.Atoi(parts[7])
	if err != nil {
		return rule, WrapError(err)
	}

	request := ecs.CreateDescribeSecurityGroupAttributeRequest()
	request.RegionId = s.client.RegionId
	request.SecurityGroupId = groupId
	request.Direction = direction
	request.NicType = nicType

	var raw interface{}
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err = s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.DescribeSecurityGroupAttribute(request)
		})
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)

	response, _ := raw.(*ecs.DescribeSecurityGroupAttributeResponse)

	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidSecurityGroupId.NotFound"}) {
			return rule, WrapErrorf(NotFoundErr("SecurityGroup:Rule", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response.RequestId))
		}
		return rule, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}

	if response == nil {
		return rule, GetNotFoundErrorFromString(GetNotFoundMessage("Security Group", groupId))
	}

	for _, ru := range response.Permissions.Permission {
		if strings.ToLower(string(ru.IpProtocol)) == ipProtocol && ru.PortRange == portRange {

			var cidr string
			var prefixListId string
			if direction == string(DirectionIngress) {
				if cidr = ru.SourceCidrIp; cidr == "" {
					cidr = ru.SourceGroupId
				}
				if cidr == "" {
					cidr = ru.Ipv6SourceCidrIp
				}
				prefixListId = ru.SourcePrefixListId
			}

			if direction == string(DirectionEgress) {
				if cidr = ru.DestCidrIp; cidr == "" {
					cidr = ru.DestGroupId
				}
				if cidr == "" {
					cidr = ru.Ipv6DestCidrIp
				}
				prefixListId = ru.DestPrefixListId
			}

			if (cidr == cidrIp || prefixListId == cidrIp) && strings.ToLower(string(ru.Policy)) == policy && ru.Priority == strconv.Itoa(priority) {
				return ru, nil
			}
		}
	}

	return rule, WrapErrorf(NotFoundErr("Security Group Rule", id), NotFoundMsg, ProviderERROR, response.RequestId)

}

func (s *EcsService) DescribeAvailableResources(d *schema.ResourceData, meta interface{}, destination DestinationResource) (zoneId string, validZones []ecs.AvailableZone, requestId string, err error) {
	client := meta.(*connectivity.AliyunClient)
	// Before creating resources, check input parameters validity according available zone.
	// If availability zone is nil, it will return all of supported resources in the current.
	request := ecs.CreateDescribeAvailableResourceRequest()
	request.RegionId = s.client.RegionId
	request.DestinationResource = string(destination)
	if destination == InstanceTypeResource {
		if v, ok := d.GetOk("cpu_core_count"); ok && v.(int) > 0 {
			request.Cores = requests.NewInteger(v.(int))
		}
		if v, ok := d.GetOk("memory_size"); ok && v.(float64) > 0 {
			request.Memory = requests.NewFloat(v.(float64))
		}
	}
	request.IoOptimized = string(IOOptimized)

	if v, ok := d.GetOk("availability_zone"); ok && strings.TrimSpace(v.(string)) != "" {
		zoneId = strings.TrimSpace(v.(string))
	} else if v, ok := d.GetOk("vswitch_id"); ok && strings.TrimSpace(v.(string)) != "" {
		vpcService := VpcService{s.client}
		if vsw, err := vpcService.DescribeVSwitch(strings.TrimSpace(v.(string))); err == nil {
			zoneId = vsw.ZoneId
		}
	}

	if v, ok := d.GetOk("instance_charge_type"); ok && strings.TrimSpace(v.(string)) != "" {
		request.InstanceChargeType = strings.TrimSpace(v.(string))
	}

	if v, ok := d.GetOk("spot_strategy"); ok && strings.TrimSpace(v.(string)) != "" {
		request.SpotStrategy = strings.TrimSpace(v.(string))
	}

	if v, ok := d.GetOk("network_type"); ok && strings.TrimSpace(v.(string)) != "" {
		request.NetworkCategory = strings.TrimSpace(v.(string))
	}

	if v, ok := d.GetOk("is_outdated"); ok && v.(bool) == true {
		request.IoOptimized = string(NoneOptimized)
	}
	if v, ok := d.GetOk("system_disk_category"); ok && strings.TrimSpace(v.(string)) != "" {
		request.SystemDiskCategory = strings.TrimSpace(v.(string))
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	var response *ecs.DescribeAvailableResourceResponse
	if err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.DescribeAvailableResource(request)
		})
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		response, _ = raw.(*ecs.DescribeAvailableResourceResponse)
		return nil
	}); err != nil {
		return "", nil, "", WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	requestId = response.RequestId

	if len(response.AvailableZones.AvailableZone) < 1 {
		err = WrapError(Error("There are no available zones in the API DescribeAvailableResource response: %#v.", response))
		return
	}

	valid := false
	soldout := false
	var expectedZones []string
	for _, zone := range response.AvailableZones.AvailableZone {
		if zone.Status == string(SoldOut) {
			if zone.ZoneId == zoneId {
				soldout = true
			}
			continue
		}
		if zoneId != "" && zone.ZoneId == zoneId {
			valid = true
			validZones = append(make([]ecs.AvailableZone, 1), zone)
			break
		}
		expectedZones = append(expectedZones, zone.ZoneId)
		validZones = append(validZones, zone)
	}
	if zoneId != "" {
		if !valid {
			err = WrapError(Error("Availability zone %s status is not available in the region %s. Expected availability zones: %s. \nDescribeAvailableResource response: %#v.",
				zoneId, client.RegionId, strings.Join(expectedZones, ", "), response))
			return
		}
		if soldout {
			err = WrapError(Error("Availability zone %s status is sold out in the region %s. Expected availability zones: %s. \nDescribeAvailableResource response: %#v.",
				zoneId, client.RegionId, strings.Join(expectedZones, ", "), response))
			return
		}
	}

	if len(validZones) <= 0 {
		err = WrapError(Error("There is no availability resources in the region %s. Please choose another region. \nDescribeAvailableResource response: %#v.", client.RegionId, response))
		return
	}

	return
}

func (s *EcsService) InstanceTypeValidation(targetType, zoneId string, validZones []ecs.AvailableZone) error {

	mapInstanceTypeToZones := make(map[string]string)
	var expectedInstanceTypes []string
	for _, zone := range validZones {
		if zoneId != "" && zoneId != zone.ZoneId {
			continue
		}
		for _, r := range zone.AvailableResources.AvailableResource {
			if r.Type == string(InstanceTypeResource) {
				for _, t := range r.SupportedResources.SupportedResource {
					if t.Status == string(SoldOut) {
						continue
					}
					if targetType == t.Value {
						return nil
					}

					if _, ok := mapInstanceTypeToZones[t.Value]; !ok {
						expectedInstanceTypes = append(expectedInstanceTypes, t.Value)
						mapInstanceTypeToZones[t.Value] = t.Value
					}
				}
			}
		}
	}
	if zoneId != "" {
		return WrapError(Error("The instance type %s is solded out or is not supported in the zone %s. Expected instance types: %s", targetType, zoneId, strings.Join(expectedInstanceTypes, ", ")))
	}
	return WrapError(Error("The instance type %s is solded out or is not supported in the region %s. Expected instance types: %s", targetType, s.client.RegionId, strings.Join(expectedInstanceTypes, ", ")))
}

func (s *EcsService) QueryInstancesWithKeyPair(instanceIdsStr, keyPair string) (instanceIds []string, err error) {
	action := "DescribeInstances"
	var response map[string]interface{}
	client := s.client
	request := make(map[string]interface{})
	request["RegionId"] = s.client.RegionId
	request["PageSize"] = PageSizeLarge
	request["PageNumber"] = 1
	request["InstanceIds"] = instanceIdsStr
	request["KeyPairName"] = keyPair
	for {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			err = WrapErrorf(err, DefaultErrorMsg, keyPair, action, AlibabaCloudSdkGoERROR)
			return
		}
		addDebug(action, response, request)

		v, err := jsonpath.Get("$.Instances.Instance", response)
		if err != nil {
			return instanceIds, WrapErrorf(err, FailedGetAttributeMsg, keyPair, "$.Instances.Instance", response)
		}
		if len(v.([]interface{})) > 0 {
			for _, inst := range v.([]interface{}) {
				instanceIds = append(instanceIds, inst.(map[string]interface{})["InstanceId"].(string))
			}
		}
		if len(v.([]interface{})) < request["PageNumber"].(int) {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}
	return
}

func (s *EcsService) DescribeKeyPair(id string) (keyPair ecs.KeyPair, err error) {
	request := ecs.CreateDescribeKeyPairsRequest()
	request.RegionId = s.client.RegionId
	request.KeyPairName = id
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeKeyPairs(request)
	})
	if err != nil {
		return keyPair, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	object, _ := raw.(*ecs.DescribeKeyPairsResponse)
	if len(object.KeyPairs.KeyPair) < 1 || object.KeyPairs.KeyPair[0].KeyPairName != id {
		return keyPair, WrapErrorf(NotFoundErr("KeyPair", id), NotFoundMsg, ProviderERROR, object.RequestId)
	}
	return object.KeyPairs.KeyPair[0], nil

}

func (s *EcsService) DescribeEcsKeyPairAttachment(id string) (keyPair ecs.KeyPair, err error) {
	index := strings.LastIndexByte(id, ':')
	keyPairName := id[:index]
	keyPair, err = s.DescribeKeyPair(keyPairName)
	if err != nil {
		return keyPair, WrapError(err)
	}
	if keyPair.KeyPairName != keyPairName {
		err = WrapErrorf(NotFoundErr("KeyPairAttachment", id), NotFoundMsg, ProviderERROR)
	}
	return keyPair, nil
}

func (s *EcsService) DescribeDisk(id string) (disk ecs.Disk, err error) {
	request := ecs.CreateDescribeDisksRequest()
	request.DiskIds = convertListToJsonString([]interface{}{id})
	request.RegionId = s.client.RegionId
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeDisks(request)
	})
	if err != nil {
		return disk, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	response, _ := raw.(*ecs.DescribeDisksResponse)
	if len(response.Disks.Disk) < 1 || response.Disks.Disk[0].DiskId != id {
		err = WrapErrorf(NotFoundErr("Disk", id), NotFoundMsg, ProviderERROR, response.RequestId)
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	return response.Disks.Disk[0], nil
}

func (s *EcsService) DescribeEcsDiskAttachment(id string) (disk map[string]interface{}, err error) {
	parts, err := ParseResourceId(id, 2)
	if err != nil {
		return disk, WrapError(err)
	}
	disk, err = s.DescribeEcsDisk(parts[0])
	if err != nil {
		return disk, WrapError(err)
	}

	if disk["InstanceId"] != parts[1] && disk["Status"] != string(InUse) {
		err = WrapErrorf(NotFoundErr("EcsDiskAttachment", id), NotFoundMsg, ProviderERROR)
	}
	return
}

func (s *EcsService) DescribeDisksByType(instanceId string, diskType DiskType) (disk []ecs.Disk, err error) {
	request := ecs.CreateDescribeDisksRequest()
	request.RegionId = s.client.RegionId
	if instanceId != "" {
		request.InstanceId = instanceId
	}
	request.DiskType = string(diskType)

	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeDisks(request)
	})
	if err != nil {
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	resp, _ := raw.(*ecs.DescribeDisksResponse)
	if resp == nil {
		return
	}
	return resp.Disks.Disk, nil
}

func (s *EcsService) DescribeTags(resourceId string, resourceType TagResourceType) (tags []ecs.Tag, err error) {
	request := ecs.CreateDescribeTagsRequest()
	request.RegionId = s.client.RegionId
	request.ResourceType = string(resourceType)
	request.ResourceId = resourceId
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeTags(request)
	})
	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, resourceId, request.GetActionName(), AlibabaCloudSdkGoERROR)
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response, _ := raw.(*ecs.DescribeTagsResponse)
	if len(response.Tags.Tag) < 1 {
		err = WrapErrorf(NotFoundErr("Tags", resourceId), NotFoundMsg, ProviderERROR)
		return
	}

	return response.Tags.Tag, nil
}

func (s *EcsService) DescribeImageById(id string) (image ecs.Image, err error) {
	request := ecs.CreateDescribeImagesRequest()
	request.RegionId = s.client.RegionId
	request.ImageId = id
	request.Status = fmt.Sprintf("%s,%s,%s,%s,%s", "Creating", "Waiting", "Available", "UnAvailable", "CreateFailed")
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeImages(request)
	})
	if err != nil {
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	resp, _ := raw.(*ecs.DescribeImagesResponse)
	if resp == nil || len(resp.Images.Image) < 1 {
		return image, GetNotFoundErrorFromString(GetNotFoundMessage("Image", id))
	}
	return resp.Images.Image[0], nil
}

func (s *EcsService) deleteImage(d *schema.ResourceData) error {

	object, err := s.DescribeImageById(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}
	request := ecs.CreateDeleteImageRequest()

	if force, ok := d.GetOk("force"); ok {
		request.Force = requests.NewBoolean(force.(bool))
	}
	request.RegionId = s.client.RegionId
	request.ImageId = object.ImageId

	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DeleteImage(request)
	})

	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	stateConf := BuildStateConf([]string{}, []string{}, d.Timeout(schema.TimeoutCreate), 3*time.Second, s.ImageStateRefreshFunc(d.Id(), []string{}))
	if _, err := stateConf.WaitForState(); err != nil {
		return WrapErrorf(err, IdMsg, d.Id())
	}

	if v, ok := d.GetOk("delete_auto_snapshot"); ok && v.(bool) {
		errs := map[string]error{}

		for _, item := range object.DiskDeviceMappings.DiskDeviceMapping {
			if item.SnapshotId == "" {
				continue
			}
			request := ecs.CreateDeleteSnapshotRequest()
			request.RegionId = s.client.RegionId
			request.SnapshotId = item.SnapshotId
			request.Force = requests.NewBoolean(true)
			raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
				return ecsClient.DeleteSnapshot(request)
			})
			if err != nil {
				errs[item.SnapshotId] = err
			}
			addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		}
		if len(errs) > 0 {
			errParts := []string{"Errors while deleting associated snapshots:"}
			for snapshotId, err := range errs {
				errParts = append(errParts, fmt.Sprintf("%s: %s", snapshotId, err))
			}
			errParts = append(errParts, "These are no longer managed by Terraform and must be deleted manually.")
			return WrapError(fmt.Errorf("%v", strings.Join(errParts, "\n")))
		}
	}

	return nil
}

func (s *EcsService) updateImage(d *schema.ResourceData) error {

	d.Partial(true)

	err := setTags(s.client, TagResourceImage, d)
	if err != nil {
		return WrapError(err)
	} else {
		d.SetPartial("tags")
	}

	request := ecs.CreateModifyImageAttributeRequest()
	request.RegionId = s.client.RegionId
	request.ImageId = d.Id()

	if d.HasChange("description") || d.HasChange("name") || d.HasChange("image_name") {
		if description, ok := d.GetOk("description"); ok {
			request.Description = description.(string)
		}
		if imageName, ok := d.GetOk("image_name"); ok {
			request.ImageName = imageName.(string)
		} else {
			if imageName, ok := d.GetOk("name"); ok {
				request.ImageName = imageName.(string)
			}
		}
		raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.ModifyImageAttribute(request)
		})
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)

		if err != nil {
			return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
		}

		d.SetPartial("name")
		d.SetPartial("image_name")
		d.SetPartial("description")
	}

	d.Partial(false)
	return nil
}

func (s *EcsService) DescribeNetworkInterface(id string) (networkInterface ecs.NetworkInterfaceSet, err error) {
	request := ecs.CreateDescribeNetworkInterfacesRequest()
	request.RegionId = s.client.RegionId
	eniIds := []string{id}
	request.NetworkInterfaceId = &eniIds
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeNetworkInterfaces(request)
	})
	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response := raw.(*ecs.DescribeNetworkInterfacesResponse)
	if len(response.NetworkInterfaceSets.NetworkInterfaceSet) < 1 ||
		response.NetworkInterfaceSets.NetworkInterfaceSet[0].NetworkInterfaceId != id {
		err = WrapErrorf(NotFoundErr("NetworkInterface", id), NotFoundMsg, ProviderERROR, response.RequestId)
		return
	}

	return response.NetworkInterfaceSets.NetworkInterfaceSet[0], nil
}

func (s *EcsService) DescribeEcsNetworkInterfaceAttachment(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	action := "DescribeNetworkInterfaces"

	client := s.client

	parts, err := ParseResourceId(id, 2)
	if err != nil {
		return nil, WrapError(err)
	}

	request := map[string]interface{}{
		"RegionId":           s.client.RegionId,
		"NetworkInterfaceId": []string{parts[0]},
		"MaxResults":         PageSizeXLarge,
	}

	idExist := false
	for {
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
			if IsExpectedErrors(err, []string{"InvalidEcsId.NotFound", "InvalidEniId.NotFound", "InvalidSecurityGroupId.NotFound", "InvalidVSwitchId.NotFound"}) {
				return object, WrapErrorf(NotFoundErr("Ecs:NetworkInterfaceAttachment", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
			}
			return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		}

		resp, err := jsonpath.Get("$.NetworkInterfaceSets.NetworkInterfaceSet", response)
		if err != nil {
			return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.NetworkInterfaceSets.NetworkInterfaceSet", response)
		}

		if v, ok := resp.([]interface{}); !ok || len(v) < 1 {
			return object, WrapErrorf(NotFoundErr("Ecs:NetworkInterfaceAttachment", id), NotFoundWithResponse, response)
		}

		for _, v := range resp.([]interface{}) {
			if fmt.Sprint(v.(map[string]interface{})["Type"]) == "Member" || fmt.Sprint(v.(map[string]interface{})["Type"]) == "slave" {
				if attachment, ok := v.(map[string]interface{})["Attachment"]; ok {
					attachmentArg := attachment.(map[string]interface{})

					if fmt.Sprint(v.(map[string]interface{})["NetworkInterfaceId"]) == parts[0] && fmt.Sprint(attachmentArg["InstanceId"]) == parts[1] {
						idExist = true
						return v.(map[string]interface{}), nil
					}
				}
			}

			if fmt.Sprint(v.(map[string]interface{})["NetworkInterfaceId"]) == parts[0] && fmt.Sprint(v.(map[string]interface{})["InstanceId"]) == parts[1] {
				idExist = true
				return v.(map[string]interface{}), nil
			}
		}

		if nextToken, ok := response["NextToken"].(string); ok && nextToken != "" {
			request["NextToken"] = nextToken
		} else {
			break
		}
	}

	if !idExist {
		return object, WrapErrorf(NotFoundErr("Ecs:NetworkInterfaceAttachment", id), NotFoundWithResponse, response)
	}

	return object, nil
}

// WaitForInstance waits for instance to given status
func (s *EcsService) WaitForEcsInstance(instanceId string, status Status, timeout int) error {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}
	for {
		instance, err := s.DescribeInstance(instanceId)
		if err != nil && !NotFoundError(err) {
			return err
		}
		if instance.Status == string(status) {
			//Sleep one more time for timing issues
			time.Sleep(DefaultIntervalMedium * time.Second)
			break
		}
		timeout = timeout - DefaultIntervalShort
		if timeout <= 0 {
			return GetTimeErrorFromString(GetTimeoutMessage("ECS Instance", string(status)))
		}
		time.Sleep(DefaultIntervalShort * time.Second)

	}
	return nil
}

// WaitForInstance waits for instance to given status
func (s *EcsService) InstanceStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeInstance(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if object.Status == failState {
				return object, object.Status, WrapError(Error(FailedToReachTargetStatus, object.Status))
			}
		}

		return object, object.Status, nil
	}
}

func (s *EcsService) WaitForDisk(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)

	for {
		object, err := s.DescribeDisk(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}
		// Disk need 3-5 seconds to get ExpiredTime after the status is available
		if object.Status == string(status) && object.ExpiredTime != "" {
			return nil
		}
		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, object.Status, string(status), ProviderERROR)
		}

	}
}

func (s *EcsService) WaitForSecurityGroup(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)

	for {
		_, err := s.DescribeSecurityGroup(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}
		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, Null, string(status), ProviderERROR)
		}

	}
}

func (s *EcsService) WaitForKeyPair(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)

	for {
		_, err := s.DescribeKeyPair(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}
		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, Null, string(status), ProviderERROR)
		}

	}
}

func (s *EcsService) WaitForNetworkInterface(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)

	for {
		object, err := s.DescribeNetworkInterface(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}

		if object.Status == string(status) {
			return nil
		}
		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, object.Status, string(status), ProviderERROR)
		}
	}
}

func (s *EcsService) QueryPrivateIps(eniId string) ([]string, error) {
	if eni, err := s.DescribeNetworkInterface(eniId); err != nil {
		return nil, fmt.Errorf("Describe NetworkInterface(%s) failed, %s", eniId, err)
	} else {
		filterIps := make([]string, 0, len(eni.PrivateIpSets.PrivateIpSet))
		for i := range eni.PrivateIpSets.PrivateIpSet {
			if eni.PrivateIpSets.PrivateIpSet[i].Primary {
				continue
			}
			filterIps = append(filterIps, eni.PrivateIpSets.PrivateIpSet[i].PrivateIpAddress)
		}
		return filterIps, nil
	}
}

func (s *EcsService) WaitForVpcAttributesChanged(instanceId, vswitchId, privateIp string) error {
	deadline := time.Now().Add(5 * time.Minute)
	for {
		if time.Now().After(deadline) {
			return WrapError(Error("Wait for VPC attributes changed timeout"))
		}
		time.Sleep(DefaultIntervalShort * time.Second)

		instance, err := s.DescribeInstance(instanceId)
		if err != nil {
			return WrapError(err)
		}

		if len(privateIp) > 0 && len(instance.VpcAttributes.PrivateIpAddress.IpAddress) > 0 {
			if instance.VpcAttributes.PrivateIpAddress.IpAddress[0] != privateIp {
				continue
			}
		}

		if len(vswitchId) > 0 && instance.VpcAttributes.VSwitchId != vswitchId {
			continue
		}

		return nil
	}
}

func (s *EcsService) WaitForPrivateIpsCountChanged(eniId string, count int) error {
	deadline := time.Now().Add(DefaultTimeout * time.Second)
	for {
		if time.Now().After(deadline) {
			return fmt.Errorf("Wait for private IP addrsses count changed timeout")
		}
		time.Sleep(DefaultIntervalShort * time.Second)

		ips, err := s.QueryPrivateIps(eniId)
		if err != nil {
			return fmt.Errorf("Query private IP failed, %s", err)
		}
		if len(ips) == count {
			return nil
		}
	}
}

func (s *EcsService) WaitForPrivateIpsListChanged(eniId string, ipList []string) error {
	deadline := time.Now().Add(DefaultTimeout * time.Second)
	for {
		if time.Now().After(deadline) {
			return fmt.Errorf("Wait for private IP addrsses list changed timeout")
		}
		time.Sleep(DefaultIntervalShort * time.Second)

		ips, err := s.QueryPrivateIps(eniId)
		if err != nil {
			return fmt.Errorf("Query private IP failed, %s", err)
		}

		if len(ips) != len(ipList) {
			continue
		}

		diff := false
		for i := range ips {
			exist := false
			for j := range ipList {
				if ips[i] == ipList[j] {
					exist = true
					break
				}
			}
			if !exist {
				diff = true
				break
			}
		}

		if !diff {
			return nil
		}
	}
}

func (s *EcsService) WaitForModifySecurityGroupPolicy(id, target string, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		object, err := s.DescribeSecurityGroup(id)
		if err != nil {
			return WrapError(err)
		}
		if object.InnerAccessPolicy == target {
			return nil
		}
		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, object.InnerAccessPolicy, target, ProviderERROR)
		}
	}
}

func (s *EcsService) AttachKeyPair(keyName string, instanceIds []interface{}) error {
	request := ecs.CreateAttachKeyPairRequest()
	request.RegionId = s.client.RegionId
	request.KeyPairName = keyName
	request.InstanceIds = convertListToJsonString(instanceIds)
	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
			return ecsClient.AttachKeyPair(request)
		})
		if err != nil {
			if IsExpectedErrors(err, []string{"ServiceUnavailable"}) {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(request.GetActionName(), raw, request.RpcRequest, request)
		return nil
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, keyName, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	return nil
}

func (s *EcsService) SnapshotStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeSnapshot(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if object.Status == failState {
				return object, object.Status, WrapError(Error(FailedToReachTargetStatus, object.Status))
			}
		}
		return object, object.Status, nil
	}
}

func (s *EcsService) ImageStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeImageById(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if object.Status == failState {
				return object, object.Status, WrapError(Error(FailedToReachTargetStatus, object.Status))
			}
		}
		return object, object.Status, nil
	}
}

func (s *EcsService) TaskStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeTaskById(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}
		for _, failState := range failStates {
			if object.TaskStatus == failState {
				return object, object.TaskStatus, WrapError(Error(FailedToReachTargetStatus, object.TaskStatus))
			}
		}
		return object, object.TaskStatus, nil
	}
}

func (s *EcsService) DescribeTaskById(id string) (task *ecs.DescribeTaskAttributeResponse, err error) {
	request := ecs.CreateDescribeTaskAttributeRequest()
	request.TaskId = id

	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeTaskAttribute(request)
	})
	if err != nil {
		return task, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	task, _ = raw.(*ecs.DescribeTaskAttributeResponse)

	if task.TaskId == "" {
		return task, GetNotFoundErrorFromString(GetNotFoundMessage("task", id))
	}
	return task, nil
}

func (s *EcsService) DescribeSnapshot(id string) (*ecs.Snapshot, error) {
	snapshot := &ecs.Snapshot{}
	request := ecs.CreateDescribeSnapshotsRequest()
	request.RegionId = s.client.RegionId
	request.SnapshotIds = fmt.Sprintf("[\"%s\"]", id)
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeSnapshots(request)
	})
	if err != nil {
		return snapshot, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response := raw.(*ecs.DescribeSnapshotsResponse)
	if len(response.Snapshots.Snapshot) != 1 || response.Snapshots.Snapshot[0].SnapshotId != id {
		return snapshot, WrapErrorf(NotFoundErr("Snapshot", id), NotFoundMsg, ProviderERROR, response.RequestId)
	}
	return &response.Snapshots.Snapshot[0], nil
}

func (s *EcsService) DescribeSnapshotPolicy(id string) (*ecs.AutoSnapshotPolicy, error) {
	policy := &ecs.AutoSnapshotPolicy{}
	request := ecs.CreateDescribeAutoSnapshotPolicyExRequest()
	request.AutoSnapshotPolicyId = id
	request.RegionId = s.client.RegionId
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeAutoSnapshotPolicyEx(request)
	})
	if err != nil {
		return policy, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)

	response := raw.(*ecs.DescribeAutoSnapshotPolicyExResponse)
	if len(response.AutoSnapshotPolicies.AutoSnapshotPolicy) != 1 ||
		response.AutoSnapshotPolicies.AutoSnapshotPolicy[0].AutoSnapshotPolicyId != id {
		return policy, WrapErrorf(NotFoundErr("SnapshotPolicy", id), NotFoundMsg, ProviderERROR, response.RequestId)
	}

	return &response.AutoSnapshotPolicies.AutoSnapshotPolicy[0], nil
}

func (s *EcsService) DescribeReservedInstance(id string) (reservedInstance ecs.ReservedInstance, err error) {
	request := ecs.CreateDescribeReservedInstancesRequest()
	var balance = &[]string{id}
	request.ReservedInstanceId = balance
	request.RegionId = s.client.RegionId
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeReservedInstances(request)
	})
	if err != nil {
		return reservedInstance, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response, _ := raw.(*ecs.DescribeReservedInstancesResponse)

	if len(response.ReservedInstances.ReservedInstance) != 1 ||
		response.ReservedInstances.ReservedInstance[0].ReservedInstanceId != id {
		return reservedInstance, GetNotFoundErrorFromString(GetNotFoundMessage("PurchaseReservedInstance", id))
	}
	return response.ReservedInstances.ReservedInstance[0], nil
}

func (s *EcsService) DescribeLaunchTemplate(id string) (set ecs.LaunchTemplateSet, err error) {

	request := ecs.CreateDescribeLaunchTemplatesRequest()
	request.RegionId = s.client.RegionId
	request.LaunchTemplateId = &[]string{id}

	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeLaunchTemplates(request)
	})
	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response := raw.(*ecs.DescribeLaunchTemplatesResponse)
	if len(response.LaunchTemplateSets.LaunchTemplateSet) != 1 ||
		response.LaunchTemplateSets.LaunchTemplateSet[0].LaunchTemplateId != id {
		err = WrapErrorf(NotFoundErr("LaunchTemplate", id), NotFoundMsg, ProviderERROR, response.RequestId)
		return
	}

	return response.LaunchTemplateSets.LaunchTemplateSet[0], nil

}

func (s *EcsService) DescribeLaunchTemplateVersion(id string, version int) (set ecs.LaunchTemplateVersionSet, err error) {

	request := ecs.CreateDescribeLaunchTemplateVersionsRequest()
	request.RegionId = s.client.RegionId
	request.LaunchTemplateId = id
	request.LaunchTemplateVersion = &[]string{strconv.FormatInt(int64(version), 10)}
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeLaunchTemplateVersions(request)
	})
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidLaunchTemplate.NotFound"}) {
			err = WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
			return
		}
		err = WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response := raw.(*ecs.DescribeLaunchTemplateVersionsResponse)
	if len(response.LaunchTemplateVersionSets.LaunchTemplateVersionSet) != 1 ||
		response.LaunchTemplateVersionSets.LaunchTemplateVersionSet[0].LaunchTemplateId != id {
		err = WrapErrorf(NotFoundErr("LaunchTemplateVersion", id), NotFoundMsg, ProviderERROR, response.RequestId)
		return
	}

	return response.LaunchTemplateVersionSets.LaunchTemplateVersionSet[0], nil

}

func (s *EcsService) WaitForLaunchTemplate(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		object, err := s.DescribeLaunchTemplate(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}
		if object.LaunchTemplateId == id && string(status) != string(Deleted) {
			return nil
		}

		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, Null, string(status), ProviderERROR)
		}
		time.Sleep(DefaultIntervalShort * time.Second)
	}
}

func (s *EcsService) DescribeImageShareByImageId(id string) (imageShare *ecs.DescribeImageSharePermissionResponse, err error) {
	request := ecs.CreateDescribeImageSharePermissionRequest()
	request.RegionId = s.client.RegionId
	parts, err := ParseResourceId(id, 2)
	if err != nil {
		return imageShare, WrapError(err)
	}
	request.ImageId = parts[0]
	raw, err := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeImageSharePermission(request)
	})
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidImageId.NotFound"}) {
			return imageShare, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return imageShare, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	resp, _ := raw.(*ecs.DescribeImageSharePermissionResponse)
	if len(resp.Accounts.Account) == 0 {
		return imageShare, WrapErrorf(NotFoundErr("ModifyImageSharePermission", id), NotFoundMsg, ProviderERROR, resp.RequestId)
	}
	return resp, nil
}

func (s *EcsService) DescribeAutoProvisioningGroup(id string) (group ecs.AutoProvisioningGroup, err error) {
	request := ecs.CreateDescribeAutoProvisioningGroupsRequest()
	ids := []string{id}
	request.AutoProvisioningGroupId = &ids
	request.RegionId = s.client.RegionId
	raw, e := s.client.WithEcsClient(func(ecsClient *ecs.Client) (interface{}, error) {
		return ecsClient.DescribeAutoProvisioningGroups(request)
	})
	if e != nil {
		err = WrapErrorf(e, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
		return
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response, _ := raw.(*ecs.DescribeAutoProvisioningGroupsResponse)
	for _, v := range response.AutoProvisioningGroups.AutoProvisioningGroup {
		if v.AutoProvisioningGroupId == id {
			return v, nil
		}
	}
	err = WrapErrorf(NotFoundErr("AutoProvisioningGroup", id), NotFoundMsg, ProviderERROR, response.RequestId)
	return
}

func (s *EcsService) ListTagResources(id string, resourceType string) (object interface{}, err error) {
	client := s.client
	action := "ListTagResources"
	request := map[string]interface{}{
		"RegionId":     s.client.RegionId,
		"ResourceType": resourceType,
		"ResourceId.1": id,
	}
	tags := make([]interface{}, 0)
	var response map[string]interface{}

	for {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			response, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
			if err != nil {
				if NeedRetry(err) {
					wait()
					return resource.RetryableError(err)
				}
				return resource.NonRetryableError(err)
			}
			addDebug(action, response, request)
			v, err := jsonpath.Get("$.TagResources.TagResource", response)
			if err != nil {
				return resource.NonRetryableError(WrapErrorf(err, FailedGetAttributeMsg, id, "$.TagResources.TagResource", response))
			}
			if v != nil {
				tags = append(tags, v.([]interface{})...)
			}
			return nil
		})
		if err != nil {
			err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
			return
		}
		if response["NextToken"] == nil {
			break
		}
		request["NextToken"] = response["NextToken"]
	}

	return tags, nil
}

func (s *EcsService) tagsToMap(tags []ecs.Tag) map[string]string {
	result := make(map[string]string)
	for _, t := range tags {
		if !s.ecsTagIgnored(t) {
			result[t.TagKey] = t.TagValue
		}
	}

	return result
}

func (s *EcsService) ecsTagIgnored(t ecs.Tag) bool {
	filter := []string{"^aliyun", "^acs:", "^http://", "^https://"}
	for _, v := range filter {
		log.Printf("[DEBUG] Matching prefix %v with %v\n", v, t.TagKey)
		ok, _ := regexp.MatchString(v, t.TagKey)
		if ok {
			log.Printf("[DEBUG] Found Alibaba Cloud specific tag %s (val: %s), ignoring.\n", t.TagKey, t.TagValue)
			return true
		}
	}
	return false
}

func (s *EcsService) SetResourceTags(d *schema.ResourceData, resourceType string) error {

	if d.HasChange("tags") {
		added, removed := parsingTags(d)
		client := s.client

		removedTagKeys := make([]string, 0)
		for _, v := range removed {
			if !ignoredTags(v, "") {
				removedTagKeys = append(removedTagKeys, v)
			}
		}
		if len(removedTagKeys) > 0 {
			action := "UntagResources"
			request := map[string]interface{}{
				"RegionId":     s.client.RegionId,
				"ResourceType": resourceType,
				"ResourceId.1": d.Id(),
			}
			for i, key := range removedTagKeys {
				request[fmt.Sprintf("TagKey.%d", i+1)] = key
			}
			wait := incrementalWait(2*time.Second, 1*time.Second)
			err := resource.Retry(10*time.Minute, func() *resource.RetryError {
				response, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)

					}
					return resource.NonRetryableError(err)
				}
				addDebug(action, response, request)
				return nil
			})
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}
		}
		if len(added) > 0 {
			action := "TagResources"
			request := map[string]interface{}{
				"RegionId":     s.client.RegionId,
				"ResourceType": resourceType,
				"ResourceId.1": d.Id(),
			}
			count := 1
			for key, value := range added {
				request[fmt.Sprintf("Tag.%d.Key", count)] = key
				request[fmt.Sprintf("Tag.%d.Value", count)] = value
				count++
			}

			wait := incrementalWait(2*time.Second, 1*time.Second)
			err := resource.Retry(10*time.Minute, func() *resource.RetryError {
				response, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)

					}
					return resource.NonRetryableError(err)
				}
				addDebug(action, response, request)
				return nil
			})
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}
		}
		d.SetPartial("tags")
	}
	return nil
}

func (s *EcsService) DescribeEcsDedicatedHost(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeDedicatedHosts"
	request := map[string]interface{}{
		"RegionId":         s.client.RegionId,
		"DedicatedHostIds": convertListToJsonString([]interface{}{id}),
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidDedicatedHostId.NotFound", "InvalidLockReason.NotFound"}) {
			return object, WrapErrorf(NotFoundErr("ECS:DedicatedHost", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.DedicatedHosts.DedicatedHost", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.DedicatedHosts.DedicatedHost", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["DedicatedHostId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) EcsDedicatedHostStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsDedicatedHost(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if fmt.Sprint(object["Status"]) == failState {
				return object, fmt.Sprint(object["Status"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["Status"])))
			}
		}
		return object, fmt.Sprint(object["Status"]), nil
	}
}

func (s *EcsService) WaitForAutoProvisioningGroup(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		object, err := s.DescribeAutoProvisioningGroup(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}
		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, object.AutoProvisioningGroupId, id, ProviderERROR)
		}
		time.Sleep(DefaultIntervalShort * time.Second)
	}
}

func (s *EcsService) DescribeEcsCommand(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeCommands"
	request := map[string]interface{}{
		"RegionId":  s.client.RegionId,
		"CommandId": id,
	}
	response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidRegionId.NotFound", "Operation.Forbidden"}) {
			err = WrapErrorf(NotFoundErr("EcsCommand", id), NotFoundMsg, ProviderERROR)
			return object, err
		}
		err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		return object, err
	}
	addDebug(action, response, request)
	v, err := jsonpath.Get("$.Commands.Command", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Commands.Command", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["CommandId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}
func (s *EcsService) DescribeEcsHpcCluster(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeHpcClusters"
	ids, err := json.Marshal([]string{id})
	if err != nil {
		return object, err
	}
	request := map[string]interface{}{
		"RegionId":      s.client.RegionId,
		"HpcClusterIds": string(ids),
	}
	request["ClientToken"] = buildClientToken("DescribeHpcClusters")
	response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
	if err != nil {
		if IsExpectedErrors(err, []string{"NotExists.HpcCluster"}) {
			err = WrapErrorf(NotFoundErr("EcsHpcCluster", id), NotFoundMsg, ProviderERROR)
			return object, err
		}
		err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		return object, err
	}
	addDebug(action, response, request)
	v, err := jsonpath.Get("$.HpcClusters.HpcCluster", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.HpcClusters.HpcCluster", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["HpcClusterId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeEcsAutoSnapshotPolicy(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeAutoSnapshotPolicyEx"
	request := map[string]interface{}{
		"RegionId":             s.client.RegionId,
		"AutoSnapshotPolicyId": id,
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.AutoSnapshotPolicies.AutoSnapshotPolicy", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.AutoSnapshotPolicies.AutoSnapshotPolicy", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["AutoSnapshotPolicyId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) EcsAutoSnapshotPolicyStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsAutoSnapshotPolicy(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if object["Status"].(string) == failState {
				return object, object["Status"].(string), WrapError(Error(FailedToReachTargetStatus, object["Status"].(string)))
			}
		}
		return object, object["Status"].(string), nil
	}
}

func (s *EcsService) DescribeLaunchTemplateVersions(id string, version interface{}) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeLaunchTemplateVersions"
	request := map[string]interface{}{
		"RegionId":              s.client.RegionId,
		"LaunchTemplateId":      id,
		"LaunchTemplateVersion": []interface{}{version},
	}
	response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidLaunchTemplate.NotFound"}) {
			err = WrapErrorf(NotFoundErr("EcsLaunchTemplate", id), NotFoundMsg, ProviderERROR)
			return object, err
		}
		err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		return object, err
	}
	addDebug(action, response, request)
	v, err := jsonpath.Get("$.LaunchTemplateVersionSets.LaunchTemplateVersionSet", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.LaunchTemplateVersionSets.LaunchTemplateVersionSet", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["LaunchTemplateId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) SetResourceTemplateTags(d *schema.ResourceData, resourceType string) error {

	if d.HasChange("template_tags") {
		added, removed := parsingTags(d)
		client := s.client

		removedTagKeys := make([]string, 0)
		for _, v := range removed {
			if !ignoredTags(v, "") {
				removedTagKeys = append(removedTagKeys, v)
			}
		}
		if len(removedTagKeys) > 0 {
			action := "UntagResources"
			request := map[string]interface{}{
				"RegionId":     s.client.RegionId,
				"ResourceType": resourceType,
				"ResourceId.1": d.Id(),
			}
			for i, key := range removedTagKeys {
				request[fmt.Sprintf("TagKey.%d", i+1)] = key
			}
			wait := incrementalWait(2*time.Second, 1*time.Second)
			err := resource.Retry(10*time.Minute, func() *resource.RetryError {
				response, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)

					}
					return resource.NonRetryableError(err)
				}
				addDebug(action, response, request)
				return nil
			})
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}
		}
		if len(added) > 0 {
			action := "TagResources"
			request := map[string]interface{}{
				"RegionId":     s.client.RegionId,
				"ResourceType": resourceType,
				"ResourceId.1": d.Id(),
			}
			count := 1
			for key, value := range added {
				request[fmt.Sprintf("Tag.%d.Key", count)] = key
				request[fmt.Sprintf("Tag.%d.Value", count)] = value
				count++
			}

			wait := incrementalWait(2*time.Second, 1*time.Second)
			err := resource.Retry(10*time.Minute, func() *resource.RetryError {
				response, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
				if err != nil {
					if NeedRetry(err) {
						wait()
						return resource.RetryableError(err)

					}
					return resource.NonRetryableError(err)
				}
				addDebug(action, response, request)
				return nil
			})
			if err != nil {
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}
		}
		d.SetPartial("template_tags")
	}
	return nil
}

func (s *EcsService) DescribeEcsLaunchTemplate(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeLaunchTemplates"
	request := map[string]interface{}{
		"RegionId":         s.client.RegionId,
		"LaunchTemplateId": []string{id},
	}
	response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		return
	}
	addDebug(action, response, request)
	v, err := jsonpath.Get("$.LaunchTemplateSets.LaunchTemplateSet", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.LaunchTemplateSets.LaunchTemplateSet", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["LaunchTemplateId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeEcsSnapshot(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	action := "DescribeSnapshots"

	client := s.client

	ids, err := json.Marshal([]string{id})
	if err != nil {
		return object, err
	}

	request := map[string]interface{}{
		"RegionId":    s.client.RegionId,
		"SnapshotIds": string(ids),
		"PageSize":    PageSizeLarge,
		"PageNumber":  1,
	}

	idExist := false
	for {
		wait := incrementalWait(3*time.Second, 5*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
			if IsExpectedErrors(err, []string{"InvalidFilterKey.NotFound", "InvalidFilterValue", "InvalidSnapshotIds.Malformed,InvalidStatus.NotFound,InvalidSnapshotType.NotFound", "InvalidSnapshotLinkId.NotFound", "InvalidSnapshotType.NotFound", "InvalidSourceDiskType", "InvalidStatus.NotFound", "InvalidUsage"}) {
				return object, WrapErrorf(NotFoundErr("Ecs:Snapshot", id), NotFoundWithResponse, response)
			}
			return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		}

		resp, err := jsonpath.Get("$.Snapshots.Snapshot", response)
		if err != nil {
			return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Snapshots.Snapshot", response)
		}

		if v, ok := resp.([]interface{}); !ok || len(v) < 1 {
			return object, WrapErrorf(NotFoundErr("Ecs:Snapshot", id), NotFoundWithResponse, response)
		}

		for _, v := range resp.([]interface{}) {
			if fmt.Sprint(v.(map[string]interface{})["SnapshotId"]) == id {
				idExist = true
				return v.(map[string]interface{}), nil
			}
		}

		if len(resp.([]interface{})) < request["PageSize"].(int) {
			break
		}

		request["PageNumber"] = request["PageNumber"].(int) + 1
	}

	if !idExist {
		return object, WrapErrorf(NotFoundErr("Ecs:Snapshot", id), NotFoundWithResponse, response)
	}

	return object, nil
}

func (s *EcsService) EcsSnapshotStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsSnapshot(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if object["Status"].(string) == failState {
				return object, object["Status"].(string), WrapError(Error(FailedToReachTargetStatus, object["Status"].(string)))
			}
		}

		return object, object["Status"].(string), nil
	}
}

func (s *EcsService) DescribeEcsKeyPair(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeKeyPairs"
	request := map[string]interface{}{
		"RegionId":    s.client.RegionId,
		"KeyPairName": id,
	}
	response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidKeyPair.NotFound"}) {
			err = WrapErrorf(NotFoundErr("EcsKeyPair", id), NotFoundMsg, ProviderERROR)
			return object, err
		}
		err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		return object, err
	}
	addDebug(action, response, request)
	v, err := jsonpath.Get("$.KeyPairs.KeyPair", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.KeyPairs.KeyPair", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["KeyPairName"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeEcsAutoSnapshotPolicyAttachment(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	parts, err := ParseResourceId(id, 2)
	if err != nil {
		return object, WrapError(err)
	}
	action := "DescribeDisks"
	request := map[string]interface{}{
		"RegionId":             s.client.RegionId,
		"DiskIds":              convertListToJsonString([]interface{}{parts[1]}),
		"autoSnapshotPolicyId": parts[0],
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		err = WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		return object, err
	}
	v, err := jsonpath.Get("$.Disks.Disk", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Disks.Disk", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeEcsDisk(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeDisks"
	request := map[string]interface{}{
		"RegionId": s.client.RegionId,
		"DiskIds":  convertListToJsonString([]interface{}{id}),
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidDiskChargeType.NotFound", "InvalidDiskIds.ValueNotSupported", "InvalidFilterKey.NotFound", "InvalidFilterValue", "InvalidLockReason.NotFound"}) {
			return object, WrapErrorf(NotFoundErr("ECS:Disk", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.Disks.Disk", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Disks.Disk", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS:Disk", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["DiskId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS:Disk", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) EcsDiskStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsDisk(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if fmt.Sprint(object["Status"]) == failState {
				return object, fmt.Sprint(object["Status"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["Status"])))
			}
		}
		return object, fmt.Sprint(object["Status"]), nil
	}
}

func (s *EcsService) EcsDiskPropertyRefreshFunc(id, property string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsDisk(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}
		return object, fmt.Sprint(object[property]), nil
	}
}

func (s *EcsService) DescribeEcsNetworkInterface(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeNetworkInterfaces"
	request := map[string]interface{}{
		"RegionId":           s.client.RegionId,
		"NetworkInterfaceId": []string{id},
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		if IsExpectedErrors(err, []string{"InvalidEcsId.NotFound", "InvalidEniId.NotFound", "InvalidSecurityGroupId.NotFound", "InvalidVSwitchId.NotFound"}) {
			return object, WrapErrorf(NotFoundErr("ECS:NetworkInterface", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.NetworkInterfaceSets.NetworkInterfaceSet", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.NetworkInterfaceSets.NetworkInterfaceSet", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["NetworkInterfaceId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) EcsNetworkInterfaceStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsNetworkInterface(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if fmt.Sprint(object["Status"]) == failState {
				return object, fmt.Sprint(object["Status"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["Status"])))
			}
		}

		return object, fmt.Sprint(object["Status"]), nil
	}
}

func (s *EcsService) DescribeEcsDeploymentSet(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client

	action := "DescribeDeploymentSets"
	request := map[string]interface{}{
		"RegionId":         s.client.RegionId,
		"DeploymentSetIds": convertListToJsonString([]interface{}{id}),
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.DeploymentSets.DeploymentSet", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.DeploymentSets.DeploymentSet", response)
	}

	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["DeploymentSetId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}

	object = v.([]interface{})[0].(map[string]interface{})

	return object, nil
}

func (s *EcsService) DescribeEcsDedicatedHostCluster(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeDedicatedHostClusters"
	request := map[string]interface{}{
		"RegionId":   s.client.RegionId,
		"PageNumber": 1,
		"PageSize":   PageSizeLarge,
	}
	idExist := false
	for {
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
			return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		}
		v, err := jsonpath.Get("$.DedicatedHostClusters.DedicatedHostCluster", response)
		if err != nil {
			return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.DedicatedHostClusters.DedicatedHostCluster", response)
		}
		if len(v.([]interface{})) < 1 {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
		for _, v := range v.([]interface{}) {
			if fmt.Sprint(v.(map[string]interface{})["DedicatedHostClusterId"]) == id {
				idExist = true
				return v.(map[string]interface{}), nil
			}
		}
		if len(v.([]interface{})) < request["PageSize"].(int) {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}
	if !idExist {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	}
	return
}

func (s *EcsService) DescribeEcsSessionManagerStatus(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	action := "DescribeUserBusinessBehavior"

	client := s.client

	request := map[string]interface{}{
		"statusKey": id,
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$", response)
	}

	object = v.(map[string]interface{})

	return object, nil
}

func (s *EcsService) EcsSessionManagerStatusStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsSessionManagerStatus(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if fmt.Sprint(object["StatusValue"]) == failState {
				return object, object["StatusValue"].(string), WrapError(Error(FailedToReachTargetStatus, object["StatusValue"].(string)))
			}
		}

		return object, fmt.Sprint(object["StatusValue"]), nil
	}
}

func (s *EcsService) DescribeEcsPrefixList(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribePrefixListAttributes"
	request := map[string]interface{}{
		"RegionId":     s.client.RegionId,
		"PrefixListId": id,
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidPrefixListId.NotFound"}) {
			return object, WrapErrorf(NotFoundErr("ECS:PrefixList", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$", response)
	}
	object = v.(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeEcsStorageCapacityUnit(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeStorageCapacityUnits"
	request := map[string]interface{}{
		"RegionId":              s.client.RegionId,
		"StorageCapacityUnitId": id,
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.StorageCapacityUnits.StorageCapacityUnit", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.StorageCapacityUnits.StorageCapacityUnit", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if fmt.Sprint(v.([]interface{})[0].(map[string]interface{})["StorageCapacityUnitId"]) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) EcsStorageCapacityUnitStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsStorageCapacityUnit(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if fmt.Sprint(object["Status"]) == failState {
				return object, fmt.Sprint(object["Status"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["Status"])))
			}
		}
		return object, fmt.Sprint(object["Status"]), nil
	}
}

func (s *EcsService) DescribeEcsImageComponent(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeImageComponents"
	request := map[string]interface{}{
		"RegionId":         s.client.RegionId,
		"ImageComponentId": []string{id},
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.ImageComponent.ImageComponentSet", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.ImageComponent.ImageComponentSet", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if fmt.Sprint(v.([]interface{})[0].(map[string]interface{})["ImageComponentId"]) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeEcsSnapshotGroup(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeSnapshotGroups"
	request := map[string]interface{}{
		"RegionId":   s.client.RegionId,
		"MaxResults": PageSizeLarge,
	}
	idExist := false
	for {
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
			return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
		}
		v, err := jsonpath.Get("$.SnapshotGroups.SnapshotGroup", response)
		if err != nil {
			return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.SnapshotGroups.SnapshotGroup", response)
		}
		if len(v.([]interface{})) < 1 {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
		for _, v := range v.([]interface{}) {
			if fmt.Sprint(v.(map[string]interface{})["SnapshotGroupId"]) == id {
				idExist = true
				return v.(map[string]interface{}), nil
			}
		}

		if nextToken, ok := response["NextToken"].(string); ok && nextToken != "" {
			request["NextToken"] = nextToken
		} else {
			break
		}
	}
	if !idExist {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	}
	return
}

func (s *EcsService) EcsSnapshotGroupStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsSnapshotGroup(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if fmt.Sprint(object["Status"]) == failState {
				return object, fmt.Sprint(object["Status"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["Status"])))
			}
		}
		return object, fmt.Sprint(object["Status"]), nil
	}
}

func (s *EcsService) DescribeEcsImagePipeline(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeImagePipelines"
	request := map[string]interface{}{
		"RegionId":        s.client.RegionId,
		"MaxResults":      PageSizeMedium,
		"ImagePipelineId": []string{id},
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.ImagePipeline.ImagePipelineSet", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.ImagePipeline.ImagePipelineSet", response)
	}

	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	}
	return v.([]interface{})[0].(map[string]interface{}), nil
}

func (s *EcsService) DescribeEcsNetworkInterfacePermission(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeNetworkInterfacePermissions"
	request := map[string]interface{}{
		"RegionId":                     s.client.RegionId,
		"NetworkInterfacePermissionId": []string{id},
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.NetworkInterfacePermissions.NetworkInterfacePermission", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.NetworkInterfacePermissions.NetworkInterfacePermission", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if fmt.Sprint(v.([]interface{})[0].(map[string]interface{})["NetworkInterfacePermissionId"]) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) EcsNetworkInterfacePermissionStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsNetworkInterfacePermission(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if fmt.Sprint(object["PermissionState"]) == failState {
				return object, fmt.Sprint(object["PermissionState"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["PermissionState"])))
			}
		}
		return object, fmt.Sprint(object["PermissionState"]), nil
	}
}

func (s *EcsService) DescribeEcsInvocation(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeInvocations"
	request := map[string]interface{}{
		"RegionId": s.client.RegionId,
		"InvokeId": id,
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.Invocations.Invocation", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Invocations.Invocation", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if fmt.Sprint(v.([]interface{})[0].(map[string]interface{})["InvokeId"]) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) EcsInvocationStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsInvocation(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if fmt.Sprint(object["InvocationStatus"]) == failState {
				return object, fmt.Sprint(object["InvocationStatus"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["InvocationStatus"])))
			}
		}
		return object, fmt.Sprint(object["InvocationStatus"]), nil
	}
}

func (s *EcsService) DescribeEcsSystemDisk(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeDisks"
	request := map[string]interface{}{
		"RegionId":   s.client.RegionId,
		"InstanceId": id,
		"DiskType":   "system",
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidDiskChargeType.NotFound", "InvalidDiskIds.ValueNotSupported", "InvalidFilterKey.NotFound", "InvalidFilterValue", "InvalidLockReason.NotFound"}) {
			return object, WrapErrorf(NotFoundErr("ECS:Disk", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.Disks.Disk", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Disks.Disk", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if v.([]interface{})[0].(map[string]interface{})["InstanceId"].(string) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) SetInstanceSetResourceTags(d *schema.ResourceData, resourceType string, instanceIds []string) error {

	if d.HasChange("tags") {
		added, removed := parsingTags(d)
		client := s.client

		removedTagKeys := make([]string, 0)
		for _, v := range removed {
			if !ignoredTags(v, "") {
				removedTagKeys = append(removedTagKeys, v)
			}
		}

		instanceIdChunks := SplitSlice(convertListStringToListInterface(instanceIds), 50)
		for _, instanceIdChunk := range instanceIdChunks {

			if len(removedTagKeys) > 0 {
				action := "UntagResources"
				request := map[string]interface{}{
					"RegionId":     s.client.RegionId,
					"ResourceType": resourceType,
				}
				for i, key := range instanceIdChunk {
					request[fmt.Sprintf("ResourceId.%d", i+1)] = key
				}
				for i, key := range removedTagKeys {
					request[fmt.Sprintf("TagKey.%d", i+1)] = key
				}
				wait := incrementalWait(2*time.Second, 1*time.Second)
				err := resource.Retry(10*time.Minute, func() *resource.RetryError {
					response, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
					if err != nil {
						if NeedRetry(err) {
							wait()
							return resource.RetryableError(err)

						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}
			}
			if len(added) > 0 {
				action := "TagResources"
				request := map[string]interface{}{
					"RegionId":     s.client.RegionId,
					"ResourceType": resourceType,
				}
				for i, key := range instanceIdChunk {
					request[fmt.Sprintf("ResourceId.%d", i+1)] = key
				}
				count := 1
				for key, value := range added {
					request[fmt.Sprintf("Tag.%d.Key", count)] = key
					request[fmt.Sprintf("Tag.%d.Value", count)] = value
					count++
				}

				wait := incrementalWait(2*time.Second, 1*time.Second)
				err := resource.Retry(10*time.Minute, func() *resource.RetryError {
					response, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, false)
					if err != nil {
						if NeedRetry(err) {
							wait()
							return resource.RetryableError(err)

						}
						return resource.NonRetryableError(err)
					}
					addDebug(action, response, request)
					return nil
				})
				if err != nil {
					return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
				}
			}
			d.SetPartial("tags")
		}
	}
	return nil
}

func (s *EcsService) DescribeEcsInstanceSet(id string) (objects []map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client

	instanceIds, err := decodeFromBase64String(id)
	if err != nil {
		return objects, WrapError(err)
	}

	action := "DescribeInstances"
	request := map[string]interface{}{
		"RegionId":    s.client.RegionId,
		"InstanceIds": convertListToJsonString(convertListStringToListInterface(instanceIds)),
		"PageNumber":  1,
		"PageSize":    PageSizeXLarge,
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidInstanceIds.NotFound"}) {
			return objects, WrapErrorf(NotFoundErr("ECS", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return objects, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	resp, err := jsonpath.Get("$.Instances.Instance", response)
	if err != nil {
		return objects, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Instances.Instance", response)
	}
	if len(resp.([]interface{})) < 1 {
		return objects, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	}

	result, _ := resp.([]interface{})
	for _, v := range result {
		item := v.(map[string]interface{})
		objects = append(objects, item)
	}

	return objects, nil
}

func (s *EcsService) EcsInstanceSetStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		objects, err := s.DescribeEcsInstanceSetCloudAssistantStatus(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		count := 0
		for _, object := range objects {
			if fmt.Sprint(object["CloudAssistantStatus"]) == "true" {
				count++
			}
		}

		if count == len(objects) {
			return objects, "Running", nil
		}

		return objects, "Starting", nil
	}
}

func (s *EcsService) DescribeEcsInstanceSetCloudAssistantStatus(id string) (objects []map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client

	instanceIds, err := decodeFromBase64String(id)
	if err != nil {
		return objects, WrapError(err)
	}

	action := "DescribeCloudAssistantStatus"
	request := map[string]interface{}{
		"RegionId":   s.client.RegionId,
		"PageNumber": 1,
		"PageSize":   PageSizeXLarge,
	}

	for i, instanceId := range instanceIds {
		request[fmt.Sprintf("InstanceId.%d", i+1)] = instanceId
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidInstanceIds.NotFound"}) {
			return objects, WrapErrorf(NotFoundErr("ECS", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return objects, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	resp, err := jsonpath.Get("$.InstanceCloudAssistantStatusSet.InstanceCloudAssistantStatus", response)
	if err != nil {
		return objects, WrapErrorf(err, FailedGetAttributeMsg, id, "$.InstanceCloudAssistantStatusSet.InstanceCloudAssistantStatus", response)
	}
	if len(resp.([]interface{})) < 1 {
		return objects, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	}

	result, _ := resp.([]interface{})
	for _, v := range result {
		item := v.(map[string]interface{})
		objects = append(objects, item)
	}

	return objects, nil
}

func (s *EcsService) EcsInstanceVmSetStateRefreshFuncWithoutOsCheck(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		objects, err := s.DescribeEcsInstanceVmSetStatus(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		count := 0
		for _, object := range objects {
			if fmt.Sprint(object["Status"]) == "Running" {
				count++
			}
		}

		if count == len(objects) {
			return objects, "Running", nil
		}

		return objects, "Starting", nil
	}
}

func (s *EcsService) DescribeEcsInstanceVmSetStatus(id string) (objects []map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client

	instanceIds, err := decodeFromBase64String(id)
	if err != nil {
		return objects, WrapError(err)
	}

	action := "DescribeInstances"
	request := map[string]interface{}{
		"RegionId":   s.client.RegionId,
		"PageNumber": 1,
		"PageSize":   PageSizeXLarge,
	}

	instanceIdsStr, _ := json.Marshal(instanceIds)
	request["InstanceIds"] = string(instanceIdsStr)

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		if IsExpectedErrors(err, []string{"InvalidInstanceIds.NotFound"}) {
			return objects, WrapErrorf(NotFoundErr("ECS", id), NotFoundMsg, ProviderERROR, fmt.Sprint(response["RequestId"]))
		}
		return objects, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	resp, err := jsonpath.Get("$.Instances.Instance", response)
	if err != nil {
		return objects, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Instances.Instance", response)
	}
	if len(resp.([]interface{})) < 1 {
		return objects, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	}

	result, _ := resp.([]interface{})
	for _, v := range result {
		item := v.(map[string]interface{})
		objects = append(objects, item)
	}

	return objects, nil
}

func (s *EcsService) DescribeEcsActivation(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeActivations"
	request := map[string]interface{}{
		"RegionId":     s.client.RegionId,
		"ActivationId": id,
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.ActivationList", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.ActivationList", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if fmt.Sprint(v.([]interface{})[0].(map[string]interface{})["ActivationId"]) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeInstanceMaintenanceAttribute(id string) (object map[string]interface{}, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeInstanceMaintenanceAttributes"
	request := map[string]interface{}{
		"RegionId":   s.client.RegionId,
		"InstanceId": []string{id},
	}
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.MaintenanceAttributes.MaintenanceAttribute", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.MaintenanceAttributes.MaintenanceAttribute", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
	} else {
		if fmt.Sprint(v.([]interface{})[0].(map[string]interface{})["InstanceId"]) != id {
			return object, WrapErrorf(NotFoundErr("ECS", id), NotFoundWithResponse, response)
		}
	}
	object = v.([]interface{})[0].(map[string]interface{})
	return object, nil
}

func (s *EcsService) DescribeSystemFailureDeleteEventInstanceIds(instanceIds []string) (ids []string, err error) {
	var response map[string]interface{}
	client := s.client
	action := "DescribeInstanceHistoryEvents"
	request := map[string]interface{}{
		"RegionId":     s.client.RegionId,
		"EventType":    "SystemFailure.Delete",
		"PageSize":     PageSizeXLarge,
		"PageNumber":   1,
		"ResourceType": "instance",
	}

	for i, instanceId := range instanceIds {
		request[fmt.Sprintf("ResourceId.%d", i)] = instanceId
	}

	for {
		wait := incrementalWait(3*time.Second, 3*time.Second)
		err = resource.Retry(5*time.Minute, func() *resource.RetryError {
			response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
			return ids, WrapErrorf(err, DataDefaultErrorMsg, action, AlibabaCloudSdkGoERROR)
		}
		resp, err := jsonpath.Get("$.InstanceSystemEventSet.InstanceSystemEventType", response)
		if err != nil {
			return ids, WrapErrorf(err, FailedGetAttributeMsg, action, "$.ImagePipeline.ImagePipelineSet", response)
		}
		result, _ := resp.([]interface{})
		for _, v := range result {
			item := v.(map[string]interface{})
			ids = append(ids, fmt.Sprint(item["InstanceId"]))
		}

		if len(result) < request["PageSize"].(int) {
			break
		}
		request["PageNumber"] = request["PageNumber"].(int) + 1
	}

	return ids, nil
}
func (s *EcsService) DescribeReservedInstanceAutoRenewAttribute(id string) (object map[string]interface{}, err error) {
	client := s.client
	if err != nil {
		return object, WrapError(err)
	}

	request := map[string]interface{}{
		"ReservedInstanceId": []string{id},
		"RegionId":           s.client.RegionId,
	}

	var response map[string]interface{}
	action := "DescribeReservedInstanceAutoRenewAttribute"
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.ReservedInstanceRenewAttributes.ReservedInstanceRenewAttribute", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.ReservedInstanceRenewAttributes.ReservedInstanceRenewAttribute", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ReservedInstance", id), NotFoundWithResponse, response)
	}
	return v.([]interface{})[0].(map[string]interface{}), nil
}
func (s *EcsService) DescribeEcsReservedInstance(id string) (object map[string]interface{}, err error) {
	client := s.client
	if err != nil {
		return object, WrapError(err)
	}

	request := map[string]interface{}{
		"ReservedInstanceId": id,
		"RegionId":           s.client.RegionId,
	}

	var response map[string]interface{}
	action := "DescribeReservedInstances"
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.ReservedInstances.ReservedInstance", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.ReservedInstances.ReservedInstance", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ReservedInstance", id), NotFoundWithResponse, response)
	}
	return v.([]interface{})[0].(map[string]interface{}), nil
}
func (s *EcsService) EcsReservedInstanceStateRefreshFunc(d *schema.ResourceData, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsReservedInstance(d.Id())
		if err != nil {
			if NotFoundError(err) {
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}
		for _, failState := range failStates {
			if fmt.Sprint(object["Status"]) == failState {
				return object, fmt.Sprint(object["Status"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["Status"])))
			}
		}
		return object, fmt.Sprint(object["Status"]), nil
	}
}

func (s *EcsService) WaitForModifyIpv6AddressCount(id string, count int, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	var ipv6SetList []interface{}
	for {
		object, err := s.DescribeEcsNetworkInterface(id)
		if err != nil {
			return WrapError(err)
		}

		if v, ok := object["Ipv6Sets"].(map[string]interface{})["Ipv6Set"]; ok {
			ipv6SetList = v.([]interface{})

			if len(ipv6SetList) == count {
				return nil
			}
		} else if count == 0 {
			return nil
		}

		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, len(ipv6SetList), count, ProviderERROR)
		}
	}
}

func (s *EcsService) WaitForModifyIpv6Address(id string, addressList []string, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)

	sort.Strings(addressList)
	for {
		ipv6SetList := make([]string, 0)
		object, err := s.DescribeEcsNetworkInterface(id)
		if err != nil {
			return WrapError(err)
		}

		if v, ok := object["Ipv6Sets"].(map[string]interface{})["Ipv6Set"]; ok {
			for _, v := range v.([]interface{}) {
				ipv6Set := v.(map[string]interface{})
				ipv6SetList = append(ipv6SetList, fmt.Sprint(ipv6Set["Ipv6Address"]))
			}
			sort.Strings(ipv6SetList)
			if len(ipv6SetList) != len(addressList) {
				continue
			}

			if len(ipv6SetList) == 0 {
				return nil
			}

			diff := false
			for key, value := range ipv6SetList {
				if value != addressList[key] {
					diff = true
					continue
				}
			}
			if !diff {
				return nil
			}

		} else if len(addressList) == 0 {
			return nil
		}

		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, ipv6SetList, addressList, ProviderERROR)
		}
	}
}

func (s *EcsService) WaitForModifyIpv4PrefixCount(id string, count int, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	var ipv4PrefixSetList []interface{}
	for {
		object, err := s.DescribeEcsNetworkInterface(id)
		if err != nil {
			return WrapError(err)
		}

		if v, ok := object["Ipv4PrefixSets"].(map[string]interface{})["Ipv4PrefixSet"]; ok {
			ipv4PrefixSetList = v.([]interface{})

			if len(ipv4PrefixSetList) == count {
				return nil
			}
		} else if count == 0 {
			return nil
		}

		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, len(ipv4PrefixSetList), count, ProviderERROR)
		}
	}
}

func (s *EcsService) WaitForModifyIpv4Prefix(id string, addressList []string, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)

	sort.Strings(addressList)
	for {
		ipv4PrefixSetList := make([]string, 0)
		object, err := s.DescribeEcsNetworkInterface(id)
		if err != nil {
			return WrapError(err)
		}

		if v, ok := object["Ipv4PrefixSets"].(map[string]interface{})["Ipv4PrefixSet"]; ok {
			for _, v := range v.([]interface{}) {
				ipv4PrefixSet := v.(map[string]interface{})
				ipv4PrefixSetList = append(ipv4PrefixSetList, fmt.Sprint(ipv4PrefixSet["Ipv4Prefix"]))
			}
			sort.Strings(ipv4PrefixSetList)
			if len(ipv4PrefixSetList) != len(addressList) {
				continue
			}

			if len(ipv4PrefixSetList) == 0 {
				return nil
			}

			diff := false
			for key, value := range ipv4PrefixSetList {
				if value != addressList[key] {
					diff = true
					continue
				}
			}
			if !diff {
				return nil
			}

		} else if len(addressList) == 0 {
			return nil
		}

		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, ipv4PrefixSetList, addressList, ProviderERROR)
		}
	}
}

func (s *EcsService) DescribeEcsCapacityReservation(id string) (object map[string]interface{}, err error) {
	client := s.client
	if err != nil {
		return object, WrapError(err)
	}

	request := map[string]interface{}{
		"PrivatePoolOptions.Ids": convertListToJsonString([]interface{}{id}),
		"RegionId":               s.client.RegionId,
		"Status":                 "Active",
	}

	var response map[string]interface{}
	action := "DescribeCapacityReservations"
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
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
		if IsExpectedErrors(err, []string{"Invalid.PrivatePoolOptions.Ids"}) {
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.CapacityReservationSet.CapacityReservationItem", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.CapacityReservationSet.CapacityReservationItem", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("CapacityReservation", id), NotFoundWithResponse, response)
	}
	return v.([]interface{})[0].(map[string]interface{}), nil
}

func (s *EcsService) EcsCapacityReservationStateRefreshFunc(d *schema.ResourceData, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsCapacityReservation(d.Id())
		if err != nil {
			if NotFoundError(err) {
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}
		for _, failState := range failStates {
			if fmt.Sprint(object["Status"]) == failState {
				return object, fmt.Sprint(object["Status"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["Status"])))
			}
		}
		return object, fmt.Sprint(object["Status"]), nil
	}
}

func (s *EcsService) DescribeEcsElasticityAssurance(id string) (object map[string]interface{}, err error) {
	client := s.client
	if err != nil {
		return object, WrapError(err)
	}

	request := map[string]interface{}{
		"PrivatePoolOptions.Ids": convertListToJsonString([]interface{}{id}),
		"RegionId":               s.client.RegionId,
	}

	var response map[string]interface{}
	action := "DescribeElasticityAssurances"
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, err := client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
		if err != nil {
			if NeedRetry(err) {
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
		if IsExpectedErrors(err, []string{"Invalid.PrivatePoolOptions.Id"}) {
			return object, WrapErrorf(err, NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}
	v, err := jsonpath.Get("$.ElasticityAssuranceSet.ElasticityAssuranceItem", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.ElasticityAssuranceSet.ElasticityAssuranceItem", response)
	}
	if len(v.([]interface{})) < 1 {
		return object, WrapErrorf(NotFoundErr("ElasticityAssurance", id), NotFoundWithResponse, response)
	}
	return v.([]interface{})[0].(map[string]interface{}), nil
}

func (s *EcsService) EcsElasticityAssuranceStateRefreshFunc(d *schema.ResourceData, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeEcsElasticityAssurance(d.Id())
		if err != nil {
			if NotFoundError(err) {
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}
		for _, failState := range failStates {
			if fmt.Sprint(object["Status"]) == failState {
				return object, fmt.Sprint(object["Status"]), WrapError(Error(FailedToReachTargetStatus, fmt.Sprint(object["Status"])))
			}
		}
		return object, fmt.Sprint(object["Status"]), nil
	}
}

func (s *EcsService) DescribeInstanceAttachmentAttribute(id string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["InstanceIds"] = "[\"" + id + "\"]"
	request["RegionId"] = client.RegionId
	action := "DescribeInstanceAttachmentAttributes"

	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, query, request, true)

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
		return object, WrapErrorf(err, DefaultErrorMsg, id, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Instances.Instance[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, id, "$.Instances.Instance[*]", response)
	}

	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(NotFoundErr("Instance", id), NotFoundMsg, response)
	}

	return v.([]interface{})[0].(map[string]interface{}), nil
}

func (s *EcsService) isSupportedNetworkCardIndex(instanceType string) (bool, error) {
	var response map[string]interface{}
	var err error
	action := "DescribeInstanceTypes"
	client := s.client
	request := map[string]interface{}{
		"InstanceTypes": []string{instanceType},
	}

	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		response, err = client.RpcPost("Ecs", "2014-05-26", action, nil, request, true)
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
		return false, WrapErrorf(err, DefaultErrorMsg, instanceType, action, AlibabaCloudSdkGoERROR)
	}

	resp, err := jsonpath.Get("$.InstanceTypes.InstanceType", response)
	if err != nil {
		return false, WrapErrorf(err, FailedGetAttributeMsg, instanceType, "$.InstanceTypes.InstanceType", response)
	}

	if v, ok := resp.([]interface{}); !ok || len(v) < 1 {
		return false, nil
	}

	for _, v := range resp.([]interface{}) {
		if fmt.Sprint(v.(map[string]interface{})["InstanceTypeId"]) == instanceType {
			if _, ok := v.(map[string]interface{})["NetworkCards"]; ok {
				return true, nil
			}
		}
	}

	return false, nil
}
