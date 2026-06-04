package alicloud

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAlicloudCenTransitRouterMulticastDomainPeerMemberDataSource(t *testing.T) {
	rand := acctest.RandIntRange(1000000, 9999999)
	checkoutSupportedRegions(t, true, connectivity.CENTransitRouterMulticastDomainPeerMemberSupportRegions)

	allConf := dataSourceTestAccConfig{
		existConfig: testAccCheckAlicloudCenTransitRouterMulticastDomainPeerMemberSourceConfig(rand, map[string]string{
			"ids":                                `["${alicloud_cen_transit_router_multicast_domain_peer_member.default.id}"]`,
			"transit_router_multicast_domain_id": `"${alicloud_cen_transit_router_multicast_domain_peer_member.default.transit_router_multicast_domain_id}"`,
		}),
		fakeConfig: testAccCheckAlicloudCenTransitRouterMulticastDomainPeerMemberSourceConfig(rand, map[string]string{
			"ids":                                `["${alicloud_cen_transit_router_multicast_domain_peer_member.default.id}_fake"]`,
			"transit_router_multicast_domain_id": `"${alicloud_cen_transit_router_multicast_domain_peer_member.default.transit_router_multicast_domain_id}"`,
		}),
	}
	steps := allConf.buildDataSourceSteps(t, &CenTransitRouterMulticastDomainPeerMemberCheckInfo, rand)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			testAccPreCheckWithAccountSiteType(t, DomesticSite)
		},
		ProviderFactories: testAccProviderFactoriesAlternate(),
		CheckDestroy:      nil,
		Steps:             steps,
	})
}

var existCenTransitRouterMulticastDomainPeerMemberMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"members.#":                  "1",
		"members.0.id":               CHECKSET,
		"members.0.group_ip_address": CHECKSET,
		"members.0.peer_transit_router_multicast_domain_id": CHECKSET,
		"members.0.transit_router_multicast_domain_id":      CHECKSET,
		"members.0.status": CHECKSET,
	}
}

var fakeCenTransitRouterMulticastDomainPeerMemberMapFunc = func(rand int) map[string]string {
	return map[string]string{
		"members.#": "0",
	}
}

var CenTransitRouterMulticastDomainPeerMemberCheckInfo = dataSourceAttr{
	resourceId:   "data.alicloud_cen_transit_router_multicast_domain_peer_members.default",
	existMapFunc: existCenTransitRouterMulticastDomainPeerMemberMapFunc,
	fakeMapFunc:  fakeCenTransitRouterMulticastDomainPeerMemberMapFunc,
}

func testAccCheckAlicloudCenTransitRouterMulticastDomainPeerMemberSourceConfig(rand int, attrMap map[string]string) string {
	var pairs []string
	for k, v := range attrMap {
		pairs = append(pairs, k+" = "+v)
	}
	config := fmt.Sprintf(`
variable "name" {
  default = "tf-testAccCenTransitRouterMulticastDomainPeerMember%d"
}
%s
data "alicloud_cen_instances" "default" {
  name_regex = "no-deleting-cen"
}
data "alicloud_cen_transit_routers" "default" {
  cen_id     = data.alicloud_cen_instances.default.instances.0.id
  name_regex = "no-deleting-cen"
}
data "alicloud_cen_transit_routers" "peer" {
  provider   = alicloudalt
  cen_id     = data.alicloud_cen_instances.default.instances.0.id
  name_regex = "qingdao-no-delete-cen"
}
data "alicloud_cen_transit_router_multicast_domains" "default" {
  transit_router_id = data.alicloud_cen_transit_routers.default.transit_routers.0.transit_router_id
  name_regex        = "no-deleting-cen"
}
data "alicloud_cen_transit_router_multicast_domains" "peer" {
  provider          = alicloudalt
  transit_router_id = data.alicloud_cen_transit_routers.peer.transit_routers.0.transit_router_id
  name_regex        = "no-deleting-cen"
}
resource "alicloud_cen_transit_router_multicast_domain_peer_member" "default" {
  transit_router_multicast_domain_id      = data.alicloud_cen_transit_router_multicast_domains.default.ids.0
  peer_transit_router_multicast_domain_id = data.alicloud_cen_transit_router_multicast_domains.peer.ids.0
  group_ip_address                        = "239.1.1.1"
}

data "alicloud_cen_transit_router_multicast_domain_peer_members" "default" {
%s
}
`, rand, configAlternateRegionProvider("cn-qingdao"), strings.Join(pairs, "\n   "))
	return config
}
