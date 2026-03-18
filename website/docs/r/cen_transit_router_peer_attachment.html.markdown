---
subcategory: "Cloud Enterprise Network (CEN)"
layout: "alicloud"
page_title: "Alicloud: alicloud_cen_transit_router_peer_attachment"
description: |-
  Provides a Alicloud Cloud Enterprise Network (CEN) Transit Router Peer Attachment resource.
---

# alicloud_cen_transit_router_peer_attachment

Provides a Cloud Enterprise Network (CEN) Transit Router Peer Attachment resource.



For information about Cloud Enterprise Network (CEN) Transit Router Peer Attachment and how to use it, see [What is Transit Router Peer Attachment](https://next.api.alibabacloud.com/document/Cbn/2017-09-12/CreateTransitRouterPeerAttachment).

-> **NOTE:** Available since v1.128.0.

## Example Usage

Basic Usage

```terraform
variable "name" {
  default = "tf_example"
}
variable "region" {
  default = "cn-hangzhou"
}
variable "peer_region" {
  default = "cn-beijing"
}
provider "alicloud" {
  alias  = "hz"
  region = var.region
}
provider "alicloud" {
  alias  = "bj"
  region = var.peer_region
}

resource "alicloud_cen_instance" "example" {
  provider          = alicloud.bj
  cen_instance_name = var.name
  protection_level  = "REDUCED"
}

resource "alicloud_cen_bandwidth_package" "example" {
  provider                   = alicloud.bj
  bandwidth                  = 5
  cen_bandwidth_package_name = "tf_example"
  geographic_region_a_id     = "China"
  geographic_region_b_id     = "China"
}

resource "alicloud_cen_bandwidth_package_attachment" "example" {
  provider             = alicloud.bj
  instance_id          = alicloud_cen_instance.example.id
  bandwidth_package_id = alicloud_cen_bandwidth_package.example.id
}

resource "alicloud_cen_transit_router" "example" {
  provider = alicloud.hz
  cen_id   = alicloud_cen_bandwidth_package_attachment.example.instance_id
}

resource "alicloud_cen_transit_router" "peer" {
  provider = alicloud.bj
  cen_id   = alicloud_cen_transit_router.example.cen_id
}

resource "alicloud_cen_transit_router_peer_attachment" "example" {
  provider                              = alicloud.hz
  cen_id                                = alicloud_cen_instance.example.id
  transit_router_id                     = alicloud_cen_transit_router.example.transit_router_id
  peer_transit_router_region_id         = var.peer_region
  peer_transit_router_id                = alicloud_cen_transit_router.peer.transit_router_id
  cen_bandwidth_package_id              = alicloud_cen_bandwidth_package_attachment.example.bandwidth_package_id
  bandwidth                             = 5
  transit_router_attachment_description = var.name
  transit_router_attachment_name        = var.name
}
```

## Argument Reference

The following arguments are supported:
* `auto_publish_route_enabled` - (Optional) AutoPublishRouteEnabled
* `bandwidth` - (Optional, Int) Bandwidth
* `bandwidth_type` - (Optional, Computed, Available since v1.157.0) BandwidthType
* `cen_bandwidth_package_id` - (Optional) BandwidthPackageId
* `cen_id` - (Optional, ForceNew) CenId
* `default_link_type` - (Optional, Computed, Available since v1.223.1) DefaultLinkType
* `peer_transit_router_id` - (Required, ForceNew) PeerTransitRouterId
* `peer_transit_router_region_id` - (Optional, ForceNew) PeerTransitRouterRegionId
* `tags` - (Optional, Map, Available since v1.247.0) The tag of the resource
* `transit_router_attachment_description` - (Optional) TransitRouterAttachmentDescription
* `transit_router_id` - (Optional, ForceNew) TransitRouterId
* `transit_router_peer_attachment_name` - (Optional, Computed, Available since v1.247.0) TransitRouterAttachmentName

The following arguments will be discarded. Please use new fields as soon as possible:
* `transit_router_attachment_name` - (Deprecated since v1.274.0). Field 'transit_router_attachment_name' has been deprecated from provider version 1.274.0. New field 'transit_router_peer_attachment_name' instead.

## Attributes Reference

The following attributes are exported:
* `id` - The ID of the resource supplied above. 
* `create_time` - The creation time of the resource.
* `region_id` - The region ID of the resource.
* `status` - The status of the resource.
* `transit_router_attachment_id` - The first ID of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) for certain actions:
* `create` - (Defaults to 5 mins) Used when create the Transit Router Peer Attachment.
* `delete` - (Defaults to 5 mins) Used when delete the Transit Router Peer Attachment.
* `update` - (Defaults to 5 mins) Used when update the Transit Router Peer Attachment.

## Import

Cloud Enterprise Network (CEN) Transit Router Peer Attachment can be imported using the id, e.g.

```shell
$ terraform import alicloud_cen_transit_router_peer_attachment.example <transit_router_attachment_id>
```