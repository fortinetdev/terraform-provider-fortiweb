---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_ha"
description: |-
  Configure FortiWEB system ha configuration.
---

# fortiweb_system_ha
Configure FortiWEB system ha configuration.

## Example Usage
```hcl
resource "fortiweb_system_ha" "test1" {
	mode = "standalone"
}

resource "fortiweb_system_ha" "test2" {
	mode = "active-passive"
	priority = 5
	override = "enable"
	network_type = "upd-tunnel"
	group_name = "group1"
	group_id = 0
	tunnel_local = "1.1.1.1"
	tunnel_peer = "2.2.2.2 3.3.3.3"
	l7_persistence_sync = "enable"
	hlck_sync = "enable"
	monitor = "port8 port9"
	ha_mgmt_status = "enable"
	ha_mgmt_interface = "port5 port6"
}

resource "fortiweb_system_ha" "test3" {
        mode = "active-active-standard"
        priority = 6
        override = "enable"
        group_name = "group1"
        group_id = 0
        hlck_sync = "enable"
	session_pickup = "enable"
        monitor = "port8 port9"
	hbdev = "port3"
	hbdev_backup = "port4"
        ha_mgmt_status = "enable"
        ha_mgmt_interface = "port6 port7"
	sever_policy_hlck = "enable"
}

resource "fortiweb_system_ha" "test4" {
        mode = "active-active-high-volume"
        priority = 6
        override = "enable"
        network_type = "upd-tunnel"
        group_name = "group1"
        group_id = 0
	tunnel_local = "1.1.1.1"
	tunnel_peer = "2.2.2.2"
        monitor = "port8 port9"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mode` - Mode setting. Valid values: active-passive, active-active-standard, active-active-high-volume, standalone.
* `priority` - The priority of the appliance when electing the primary appliance in the HA pair. Valid values: 0-9.
* `override` - Enable to make priority a more important factor than uptime when selecting the primary appliance. Valid values: enable, disable.
* `network_type` - Select the common HA mode flat or udp-tunnel mode on OpenStack platform. Valid values: flat, udp-tunnel.
* `group_name` - A name to identify the HA pair if you have more than one.
* `group_id` - Both members of the HA pair must have the same group ID. Valid values: 0-63.
* `tunnel_local` - The local IP address on OpenStack platform.
* `tunnel_peer` - The peer IP address on OpenStack platform.
* `l7_persistence_sync` - When FortiWeb is operating in HA Active-Passive mode, you can enable Layer 7 Persistence Synchronization. Valid values: enable, disable.
* `hlck_sync` - Enable to synchronize the back-end servers' health check status from the primary to the secondary nodes. Valid values: enable, disable.
* `session_pickup` - Enable so that the primary unit in the HA cluster synchronizes the session table with all cluster units. Valid values: enable, disable.
* `monitor` - The name of one or more network interfaces that each directly correlate with a physical link.
* `hbdev` - Select which port on this appliance that the main and standby appliances will use to send heartbeat signals and synchronization data between each other.
* `hbdev_backup` - Select a secondary, standby port on this appliance that the main and standby appliances will use to send heartbeat signals and synchronization data between each other.
* `ha_mgmt_status` - Specifies whether the network interface you select provides administrative access to this appliance when it is a member of the HA cluster. Valid values: enable, disable.
* `ha_mgmt_interface` - Specifies the network interface that provides administrative access to this appliance when it is a member of the HA cluster.
* `sever_policy_hlck` - Enable HA AA server policy health check. Valid values: enable, disable.
* `schedule` - Specifies the load-balancing algorithm used by the primary appliance to distribute received traffic over the available cluster members. Valid values: ip, leastconnection, round-robin.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Ha can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_ha.labelname {mkey}
```
