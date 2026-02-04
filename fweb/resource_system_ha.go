// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/ha.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemHaRead,
		Update: resourceSystemHaUpdate,
		Create: resourceSystemHaUpdate,
		Delete: resourceSystemHaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			//standalone, active-passive, active-active-standard, active-active-high-volume
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//0-9
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			//flat, udp-tunnel
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"tunnel_local": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"tunnel_peer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"l7_persistence_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hlck_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ha_mgmt_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ha_mgmt_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"session_pickup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hbdev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hbdev_backup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"server_policy_hlck": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"boot_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"arps": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"arp_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lacp_ha_slave": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"session_sync_dev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"session_sync_broadcast": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"session_warm_up": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"weight_1": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"weight_2": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"weight_3": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"weight_4": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"weight_5": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"weight_6": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"weight_7": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"weight_8": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"link_failed_signal": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"eip_addr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"eip_aid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ha_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hc_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"l2ep_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multi_cluster": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multi_cluster_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multi_cluster_switch_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multi_cluster_move_primary_cluster": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cluster_arp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lb_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lb_ocid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"lb_gcp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hlck_period_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hlck_period_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"distribution": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemHaCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemHa: type error")
	}

	obj, err := getObjectSystemHa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemHa resource while getting object: %v", err)
	}

	_, err = c.CreateSystemHa(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemHa resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemHaRead(d, m)
}

func resourceSystemHaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := "SystemHa"
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHa(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemHaRead(d, m)
}

func resourceSystemHaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemHa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	(*obj)["mode"] = "standalone"
	_, err = c.UpdateSystemHa(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemHa(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHa(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource from API: %v", err)
	}

	return nil
}

func flattenSystemHa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mode", flattenSystemHa(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}
	if err = d.Set("group_id", flattenSystemHa(o["group-id"], d, "group_id", sv)); err != nil {
		if !fortiAPIPatch(o["group-id"]) {
			return fmt.Errorf("Error reading group_id: %v", err)
		}
	}
	if err = d.Set("group_name", flattenSystemHa(o["group-name"], d, "group_name", sv)); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}
	if err = d.Set("priority", flattenSystemHa(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}
	if err = d.Set("override", flattenSystemHa(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}
	if err = d.Set("network_type", flattenSystemHa(o["network-type"], d, "network_type", sv)); err != nil {
		if !fortiAPIPatch(o["network-type"]) {
			return fmt.Errorf("Error reading network_type: %v", err)
		}
	}
	if err = d.Set("tunnel_local", flattenSystemHa(o["tunnel-local"], d, "tunnel_local", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-local"]) {
			return fmt.Errorf("Error reading tunnel_local: %v", err)
		}
	}
	if err = d.Set("tunnel_peer", flattenSystemHa(o["tunnel-peer"], d, "tunnel_peer", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-peer"]) {
			return fmt.Errorf("Error reading tunnel_peer: %v", err)
		}
	}
	if err = d.Set("hbdev", flattenSystemHa(o["hbdev"], d, "hbdev", sv)); err != nil {
		if !fortiAPIPatch(o["hbdev"]) {
			return fmt.Errorf("Error reading hbdev: %v", err)
		}
	}
	if err = d.Set("hbdev_backup", flattenSystemHa(o["hbdev-backup"], d, "hbdev_backup", sv)); err != nil {
		if !fortiAPIPatch(o["hbdev-backup"]) {
			return fmt.Errorf("Error reading hbdev_backup: %v", err)
		}
	}
	if err = d.Set("boot_time", flattenSystemHa(o["boot-time"], d, "boot_time", sv)); err != nil {
		if !fortiAPIPatch(o["boot-time"]) {
			return fmt.Errorf("Error reading boot_time: %v", err)
		}
	}
	if err = d.Set("hb_interval", flattenSystemHa(o["hb-interval"], d, "hb_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hb-interval"]) {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}
	if err = d.Set("hb_lost_threshold", flattenSystemHa(o["hb-lost-threshold"], d, "hb_lost_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}
	if err = d.Set("arps", flattenSystemHa(o["arps"], d, "arps", sv)); err != nil {
		if !fortiAPIPatch(o["arps"]) {
			return fmt.Errorf("Error reading arps: %v", err)
		}
	}
	if err = d.Set("arp_interval", flattenSystemHa(o["arp-interval"], d, "arp_interval", sv)); err != nil {
		if !fortiAPIPatch(o["arp-interval"]) {
			return fmt.Errorf("Error reading arp_interval: %v", err)
		}
	}
	if err = d.Set("monitor", flattenSystemHa(o["monitor"], d, "monitor", sv)); err != nil {
		if !fortiAPIPatch(o["monitor"]) {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}
	if err = d.Set("key", flattenSystemHa(o["key"], d, "key", sv)); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}
	if err = d.Set("lacp_ha_slave", flattenSystemHa(o["lacp-ha-slave"], d, "lacp_ha_slave", sv)); err != nil {
		if !fortiAPIPatch(o["lacp-ha-slave"]) {
			return fmt.Errorf("Error reading lacp_ha_slave: %v", err)
		}
	}
	if err = d.Set("ha_mgmt_status", flattenSystemHa(o["ha-mgmt-status"], d, "ha_mgmt_status", sv)); err != nil {
		if !fortiAPIPatch(o["ha-mgmt-status"]) {
			return fmt.Errorf("Error reading ha_mgmt_status: %v", err)
		}
	}
	if err = d.Set("ha_mgmt_interface", flattenSystemHa(o["ha-mgmt-interface"], d, "ha_mgmt_interface", sv)); err != nil {
		if !fortiAPIPatch(o["ha-mgmt-interface"]) {
			return fmt.Errorf("Error reading ha_mgmt_interface: %v", err)
		}
	}
	if err = d.Set("session_pickup", flattenSystemHa(o["session-pickup"], d, "session_pickup", sv)); err != nil {
		if !fortiAPIPatch(o["session-pickup"]) {
			return fmt.Errorf("Error reading session_pickup: %v", err)
		}
	}
	if err = d.Set("session_sync_dev", flattenSystemHa(o["session-sync-dev"], d, "session_sync_dev", sv)); err != nil {
		if !fortiAPIPatch(o["session-sync-dev"]) {
			return fmt.Errorf("Error reading session_sync_dev: %v", err)
		}
	}
	if err = d.Set("session_sync_broadcast", flattenSystemHa(o["session-sync-broadcast"], d, "session_sync_broadcast", sv)); err != nil {
		if !fortiAPIPatch(o["session-sync-broadcast"]) {
			return fmt.Errorf("Error reading session_sync_broadcast: %v", err)
		}
	}
	if err = d.Set("session_warm_up", flattenSystemHa(o["session-warm-up"], d, "session_warm_up", sv)); err != nil {
		if !fortiAPIPatch(o["session-warm-up"]) {
			return fmt.Errorf("Error reading session_warm_up: %v", err)
		}
	}
	if err = d.Set("schedule", flattenSystemHa(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}
	if err = d.Set("weight_1", flattenSystemHa(o["weight-1"], d, "weight_1", sv)); err != nil {
		if !fortiAPIPatch(o["weight-1"]) {
			return fmt.Errorf("Error reading weight_1: %v", err)
		}
	}
	if err = d.Set("weight_2", flattenSystemHa(o["weight-2"], d, "weight_2", sv)); err != nil {
		if !fortiAPIPatch(o["weight-2"]) {
			return fmt.Errorf("Error reading weight_2: %v", err)
		}
	}
	if err = d.Set("weight_3", flattenSystemHa(o["weight-3"], d, "weight_3", sv)); err != nil {
		if !fortiAPIPatch(o["weight-3"]) {
			return fmt.Errorf("Error reading weight_3: %v", err)
		}
	}
	if err = d.Set("weight_4", flattenSystemHa(o["weight-4"], d, "weight_4", sv)); err != nil {
		if !fortiAPIPatch(o["weight-4"]) {
			return fmt.Errorf("Error reading weight_4: %v", err)
		}
	}
	if err = d.Set("weight_5", flattenSystemHa(o["weight-5"], d, "weight_5", sv)); err != nil {
		if !fortiAPIPatch(o["weight-5"]) {
			return fmt.Errorf("Error reading weight_5: %v", err)
		}
	}
	if err = d.Set("weight_6", flattenSystemHa(o["weight-6"], d, "weight_6", sv)); err != nil {
		if !fortiAPIPatch(o["weight-6"]) {
			return fmt.Errorf("Error reading weight_6: %v", err)
		}
	}
	if err = d.Set("weight_7", flattenSystemHa(o["weight-7"], d, "weight_7", sv)); err != nil {
		if !fortiAPIPatch(o["weight-7"]) {
			return fmt.Errorf("Error reading weight_7: %v", err)
		}
	}
	if err = d.Set("weight_8", flattenSystemHa(o["weight-8"], d, "weight_8", sv)); err != nil {
		if !fortiAPIPatch(o["weight-8"]) {
			return fmt.Errorf("Error reading weight_8: %v", err)
		}
	}
	if err = d.Set("link_failed_signal", flattenSystemHa(o["link-failed-signal"], d, "link_failed_signal", sv)); err != nil {
		if !fortiAPIPatch(o["link-failed-signal"]) {
			return fmt.Errorf("Error reading link_failed_signal: %v", err)
		}
	}
	if err = d.Set("l7_persistence_sync", flattenSystemHa(o["l7-persistence-sync"], d, "l7_persistence_sync", sv)); err != nil {
		if !fortiAPIPatch(o["l7-persistence-sync"]) {
			return fmt.Errorf("Error reading l7_persistence_sync: %v", err)
		}
	}
	if err = d.Set("eip_addr", flattenSystemHa(o["eip-addr"], d, "eip_addr", sv)); err != nil {
		if !fortiAPIPatch(o["eip-addr"]) {
			return fmt.Errorf("Error reading eip_addr: %v", err)
		}
	}
	if err = d.Set("eip_aid", flattenSystemHa(o["eip-aid"], d, "eip_aid", sv)); err != nil {
		if !fortiAPIPatch(o["eip-aid"]) {
			return fmt.Errorf("Error reading eip_aid: %v", err)
		}
	}
	if err = d.Set("ha_eth_type", flattenSystemHa(o["ha-eth-type"], d, "ha_eth_type", sv)); err != nil {
		if !fortiAPIPatch(o["ha-eth-type"]) {
			return fmt.Errorf("Error reading ha_eth_type: %v", err)
		}
	}
	if err = d.Set("hc_eth_type", flattenSystemHa(o["hc-eth-type"], d, "hc_eth_type", sv)); err != nil {
		if !fortiAPIPatch(o["hc-eth-type"]) {
			return fmt.Errorf("Error reading hc_eth_type: %v", err)
		}
	}
	if err = d.Set("l2ep_eth_type", flattenSystemHa(o["l2ep-eth-type"], d, "l2ep_eth_type", sv)); err != nil {
		if !fortiAPIPatch(o["l2ep-eth-type"]) {
			return fmt.Errorf("Error reading l2ep_eth_type: %v", err)
		}
	}
	if err = d.Set("server_policy_hlck", flattenSystemHa(o["server-policy-hlck"], d, "server_policy_hlck", sv)); err != nil {
		if !fortiAPIPatch(o["server-policy-hlck"]) {
			return fmt.Errorf("Error reading server_policy_hlck: %v", err)
		}
	}
	if err = d.Set("multi_cluster", flattenSystemHa(o["multi-cluster"], d, "multi_cluster", sv)); err != nil {
		if !fortiAPIPatch(o["multi-cluster"]) {
			return fmt.Errorf("Error reading multi_cluster: %v", err)
		}
	}
	if err = d.Set("multi_cluster_group", flattenSystemHa(o["multi-cluster-group"], d, "multi_cluster_group", sv)); err != nil {
		if !fortiAPIPatch(o["multi-cluster-group"]) {
			return fmt.Errorf("Error reading multi_cluster_group: %v", err)
		}
	}
	if err = d.Set("multi_cluster_switch_by", flattenSystemHa(o["multi-cluster-switch-by"], d, "multi_cluster_switch_by", sv)); err != nil {
		if !fortiAPIPatch(o["multi-cluster-switch-by"]) {
			return fmt.Errorf("Error reading multi_cluster_switch_by: %v", err)
		}
	}
	if err = d.Set("multi_cluster_move_primary_cluster", flattenSystemHa(o["multi-cluster-move-primary-cluster"], d, "multi_cluster_move_primary_cluster", sv)); err != nil {
		if !fortiAPIPatch(o["multi-cluster-move-primary-cluster"]) {
			return fmt.Errorf("Error reading multi_cluster_move_primary_cluster: %v", err)
		}
	}
	if err = d.Set("encryption", flattenSystemHa(o["encryption"], d, "encryption", sv)); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}
	if err = d.Set("cluster_arp", flattenSystemHa(o["cluster-arp"], d, "cluster_arp", sv)); err != nil {
		if !fortiAPIPatch(o["cluster-arp"]) {
			return fmt.Errorf("Error reading cluster_arp: %v", err)
		}
	}
	if err = d.Set("sdn_connector", flattenSystemHa(o["sdn-connector"], d, "sdn_connector", sv)); err != nil {
		if !fortiAPIPatch(o["sdn-connector"]) {
			return fmt.Errorf("Error reading sdn_connector: %v", err)
		}
	}
	if err = d.Set("lb_name", flattenSystemHa(o["lb-name"], d, "lb_name", sv)); err != nil {
		if !fortiAPIPatch(o["lb-name"]) {
			return fmt.Errorf("Error reading lb_name: %v", err)
		}
	}
	if err = d.Set("lb_ocid", flattenSystemHa(o["lb-ocid"], d, "lb_ocid", sv)); err != nil {
		if !fortiAPIPatch(o["lb-ocid"]) {
			return fmt.Errorf("Error reading lb_ocid: %v", err)
		}
	}
	if err = d.Set("lb_gcp", flattenSystemHa(o["lb-gcp"], d, "lb_gcp", sv)); err != nil {
		if !fortiAPIPatch(o["lb-gcp"]) {
			return fmt.Errorf("Error reading lb_gcp: %v", err)
		}
	}
	if err = d.Set("hlck_sync", flattenSystemHa(o["hlck-sync"], d, "hlck_sync", sv)); err != nil {
		if !fortiAPIPatch(o["hlck-sync"]) {
			return fmt.Errorf("Error reading hlck_sync: %v", err)
		}
	}
	if err = d.Set("hlck_period_sync", flattenSystemHa(o["hlck-period-sync"], d, "hlck_period_sync", sv)); err != nil {
		if !fortiAPIPatch(o["hlck-period-sync"]) {
			return fmt.Errorf("Error reading hlck_period_sync: %v", err)
		}
	}
	if err = d.Set("hlck_period_timeout", flattenSystemHa(o["hlck-period-timeout"], d, "hlck_period_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["hlck-period-timeout"]) {
			return fmt.Errorf("Error reading hlck_period_timeout: %v", err)
		}
	}
	if err = d.Set("distribution", flattenSystemHa(o["distribution"], d, "distribution", sv)); err != nil {
		if !fortiAPIPatch(o["distribution"]) {
			return fmt.Errorf("Error reading distribution: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemHa(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemHa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHa(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemHa(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}
	if v, ok := d.GetOk("group_id"); ok {
		t, err := expandSystemHa(d, v, "group_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-id"] = t
		}
	}
	if v, ok := d.GetOk("group_name"); ok {
		t, err := expandSystemHa(d, v, "group_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}
	if v, ok := d.GetOk("priority"); ok {
		t, err := expandSystemHa(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}
	if v, ok := d.GetOk("override"); ok {
		t, err := expandSystemHa(d, v, "override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}
	if v, ok := d.GetOk("network_type"); ok {
		t, err := expandSystemHa(d, v, "network_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-type"] = t
		}
	}
	if v, ok := d.GetOk("tunnel_local"); ok {
		t, err := expandSystemHa(d, v, "tunnel_local", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-local"] = t
		}
	}
	if v, ok := d.GetOk("tunnel_peer"); ok {
		t, err := expandSystemHa(d, v, "tunnel_peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-peer"] = t
		}
	}
	if v, ok := d.GetOk("hbdev"); ok {
		t, err := expandSystemHa(d, v, "hbdev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hbdev"] = t
		}
	}
	if v, ok := d.GetOk("hbdev_backup"); ok {
		t, err := expandSystemHa(d, v, "hbdev_backup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hbdev-backup"] = t
		}
	}
	if v, ok := d.GetOk("boot_time"); ok {
		t, err := expandSystemHa(d, v, "boot_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["boot-time"] = t
		}
	}
	if v, ok := d.GetOk("hb_interval"); ok {
		t, err := expandSystemHa(d, v, "hb_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}
	if v, ok := d.GetOk("hb_lost_threshold"); ok {
		t, err := expandSystemHa(d, v, "hb_lost_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}
	if v, ok := d.GetOk("arps"); ok {
		t, err := expandSystemHa(d, v, "arps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arps"] = t
		}
	}
	if v, ok := d.GetOk("arp_interval"); ok {
		t, err := expandSystemHa(d, v, "arp_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arp-interval"] = t
		}
	}
	if v, ok := d.GetOk("monitor"); ok {
		t, err := expandSystemHa(d, v, "monitor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}
	if v, ok := d.GetOk("key"); ok {
		t, err := expandSystemHa(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}
	if v, ok := d.GetOk("lacp_ha_slave"); ok {
		t, err := expandSystemHa(d, v, "lacp_ha_slave", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-ha-slave"] = t
		}
	}
	if v, ok := d.GetOk("ha_mgmt_status"); ok {
		t, err := expandSystemHa(d, v, "ha_mgmt_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-mgmt-status"] = t
		}
	}
	if v, ok := d.GetOk("ha_mgmt_interface"); ok {
		t, err := expandSystemHa(d, v, "ha_mgmt_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-mgmt-interface"] = t
		}
	}
	if v, ok := d.GetOk("session_pickup"); ok {
		t, err := expandSystemHa(d, v, "session_pickup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup"] = t
		}
	}
	if v, ok := d.GetOk("session_sync_dev"); ok {
		t, err := expandSystemHa(d, v, "session_sync_dev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-dev"] = t
		}
	}
	if v, ok := d.GetOk("session_sync_broadcast"); ok {
		t, err := expandSystemHa(d, v, "session_sync_broadcast", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-broadcast"] = t
		}
	}
	if v, ok := d.GetOk("session_warm_up"); ok {
		t, err := expandSystemHa(d, v, "session_warm_up", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-warm-up"] = t
		}
	}
	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandSystemHa(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}
	if v, ok := d.GetOk("weight_1"); ok {
		t, err := expandSystemHa(d, v, "weight_1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-1"] = t
		}
	}
	if v, ok := d.GetOk("weight_2"); ok {
		t, err := expandSystemHa(d, v, "weight_2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-2"] = t
		}
	}
	if v, ok := d.GetOk("weight_3"); ok {
		t, err := expandSystemHa(d, v, "weight_3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-3"] = t
		}
	}
	if v, ok := d.GetOk("weight_4"); ok {
		t, err := expandSystemHa(d, v, "weight_4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-4"] = t
		}
	}
	if v, ok := d.GetOk("weight_5"); ok {
		t, err := expandSystemHa(d, v, "weight_5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-5"] = t
		}
	}
	if v, ok := d.GetOk("weight_6"); ok {
		t, err := expandSystemHa(d, v, "weight_6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-6"] = t
		}
	}
	if v, ok := d.GetOk("weight_7"); ok {
		t, err := expandSystemHa(d, v, "weight_7", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-7"] = t
		}
	}
	if v, ok := d.GetOk("weight_8"); ok {
		t, err := expandSystemHa(d, v, "weight_8", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight-8"] = t
		}
	}
	if v, ok := d.GetOk("link_failed_signal"); ok {
		t, err := expandSystemHa(d, v, "link_failed_signal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-failed-signal"] = t
		}
	}
	if v, ok := d.GetOk("l7_persistence_sync"); ok {
		t, err := expandSystemHa(d, v, "l7_persistence_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l7-persistence-sync"] = t
		}
	}
	if v, ok := d.GetOk("eip_addr"); ok {
		t, err := expandSystemHa(d, v, "eip_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eip-addr"] = t
		}
	}
	if v, ok := d.GetOk("eip_aid"); ok {
		t, err := expandSystemHa(d, v, "eip_aid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eip-aid"] = t
		}
	}
	if v, ok := d.GetOk("ha_eth_type"); ok {
		t, err := expandSystemHa(d, v, "ha_eth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-eth-type"] = t
		}
	}
	if v, ok := d.GetOk("hc_eth_type"); ok {
		t, err := expandSystemHa(d, v, "hc_eth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hc-eth-type"] = t
		}
	}
	if v, ok := d.GetOk("l2ep_eth_type"); ok {
		t, err := expandSystemHa(d, v, "l2ep_eth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2ep-eth-type"] = t
		}
	}
	if v, ok := d.GetOk("server_policy_hlck"); ok {
		t, err := expandSystemHa(d, v, "server_policy_hlck", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-policy-hlck"] = t
		}
	}
	if v, ok := d.GetOk("multi_cluster"); ok {
		t, err := expandSystemHa(d, v, "multi_cluster", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-cluster"] = t
		}
	}
	if v, ok := d.GetOk("multi_cluster_group"); ok {
		t, err := expandSystemHa(d, v, "multi_cluster_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-cluster-group"] = t
		}
	}
	if v, ok := d.GetOk("multi_cluster_switch_by"); ok {
		t, err := expandSystemHa(d, v, "multi_cluster_switch_by", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-cluster-switch-by"] = t
		}
	}
	if v, ok := d.GetOk("multi_cluster_move_primary_cluster"); ok {
		t, err := expandSystemHa(d, v, "multi_cluster_move_primary_cluster", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-cluster-move-primary-cluster"] = t
		}
	}
	if v, ok := d.GetOk("encryption"); ok {
		t, err := expandSystemHa(d, v, "encryption", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}
	if v, ok := d.GetOk("cluster_arp"); ok {
		t, err := expandSystemHa(d, v, "cluster_arp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cluster-arp"] = t
		}
	}
	if v, ok := d.GetOk("sdn_connector"); ok {
		t, err := expandSystemHa(d, v, "sdn_connector", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-connector"] = t
		}
	}
	if v, ok := d.GetOk("lb_name"); ok {
		t, err := expandSystemHa(d, v, "lb_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lb-name"] = t
		}
	}
	if v, ok := d.GetOk("lb_ocid"); ok {
		t, err := expandSystemHa(d, v, "lb_ocid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lb-ocid"] = t
		}
	}
	if v, ok := d.GetOk("lb_gcp"); ok {
		t, err := expandSystemHa(d, v, "lb_gcp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lb-gcp"] = t
		}
	}
	if v, ok := d.GetOk("hlck_sync"); ok {
		t, err := expandSystemHa(d, v, "hlck_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hlck-sync"] = t
		}
	}
	if v, ok := d.GetOk("hlck_period_sync"); ok {
		t, err := expandSystemHa(d, v, "hlck_period_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hlck-period-sync"] = t
		}
	}
	if v, ok := d.GetOk("hlck_period_timeout"); ok {
		t, err := expandSystemHa(d, v, "hlck_period_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hlck-period-timeout"] = t
		}
	}
	if v, ok := d.GetOk("distribution"); ok {
		t, err := expandSystemHa(d, v, "distribution", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribution"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemHa(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
