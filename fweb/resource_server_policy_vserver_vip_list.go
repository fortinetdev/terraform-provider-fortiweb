// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure server_policy/vserver/vip-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerPolicyVserverVipList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceServerPolicyVserverVipListRead,
		Update: resourceServerPolicyVserverVipListUpdate,
		Create: resourceServerPolicyVserverVipListCreate,
		Delete: resourceServerPolicyVserverVipListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"use_interface_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceServerPolicyVserverVipListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing ServerPolicyVserverVipList: type error")
	}

	obj, err := getObjectServerPolicyVserverVipList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyVserverVipList resource while getting object: %v", err)
	}

	_, err = c.CreateServerPolicyVserverVipList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyVserverVipList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceServerPolicyVserverVipListRead(d, m)
}

func resourceServerPolicyVserverVipListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing ServerPolicyVserverVipList: type error")
	}

	obj, err := getObjectServerPolicyVserverVipList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyVserverVipList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateServerPolicyVserverVipList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyVserverVipList resource: %v", err)
	}

	return resourceServerPolicyVserverVipListRead(d, m)
}

func resourceServerPolicyVserverVipListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing ServerPolicyVserverVipList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteServerPolicyVserverVipList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting ServerPolicyVserverVipList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceServerPolicyVserverVipListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	sub_mkey := "1"
	t := d.Get("sub_mkey")
	if v, ok := t.(string); ok {
		sub_mkey = v
	} else {
		return fmt.Errorf("Error describing ServerPolicyVserverVipList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadServerPolicyVserverVipList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyVserverVipList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectServerPolicyVserverVipList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyVserverVipList resource from API: %v", err)
	}

	return nil
}

func flattenServerPolicyVserverVipList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectServerPolicyVserverVipList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenServerPolicyVserverVipList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("interface", flattenServerPolicyVserverVipList(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}
	if err = d.Set("status", flattenServerPolicyVserverVipList(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}
	if err = d.Set("use_interface_ip", flattenServerPolicyVserverVipList(o["use-interface-ip"], d, "use_interface_ip", sv)); err != nil {
		if !fortiAPIPatch(o["use-interface-ip"]) {
			return fmt.Errorf("Error reading use_interface_ip: %v", err)
		}
	}

	return nil
}

func expandServerPolicyVserverVipList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectServerPolicyVserverVipList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandServerPolicyVserverVipList(d, v, "id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("interface"); ok {
		t, err := expandServerPolicyVserverVipList(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}
	if v, ok := d.GetOk("status"); ok {
		t, err := expandServerPolicyVserverVipList(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}
	if v, ok := d.GetOk("use_interface_ip"); ok {
		t, err := expandServerPolicyVserverVipList(d, v, "use_interface_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-interface-ip"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandServerPolicyVserverVipList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
