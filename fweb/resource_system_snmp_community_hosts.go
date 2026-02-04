// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/snmp.community/hosts.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSnmpCommunityHosts() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemSnmpCommunityHostsRead,
		Update: resourceSystemSnmpCommunityHostsUpdate,
		Create: resourceSystemSnmpCommunityHostsCreate,
		Delete: resourceSystemSnmpCommunityHostsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
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

func resourceSystemSnmpCommunityHostsCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpCommunityHosts: type error")
	}

	obj, err := getObjectSystemSnmpCommunityHosts(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunityHosts resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpCommunityHosts(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpCommunityHosts resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemSnmpCommunityHostsRead(d, m)
}

func resourceSystemSnmpCommunityHostsUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpCommunityHosts: type error")
	}

	obj, err := getObjectSystemSnmpCommunityHosts(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunityHosts resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemSnmpCommunityHosts(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpCommunityHosts resource: %v", err)
	}

	return resourceSystemSnmpCommunityHostsRead(d, m)
}

func resourceSystemSnmpCommunityHostsDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpCommunityHosts: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteSystemSnmpCommunityHosts(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpCommunityHosts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpCommunityHostsRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpCommunityHosts: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadSystemSnmpCommunityHosts(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunityHosts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpCommunityHosts(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpCommunityHosts resource from API: %v", err)
	}

	return nil
}

func flattenSystemSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpCommunityHosts(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemSnmpCommunityHosts(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("ip", flattenSystemSnmpCommunityHosts(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemSnmpCommunityHosts(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpCommunityHosts(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemSnmpCommunityHosts(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemSnmpCommunityHosts(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemSnmpCommunityHosts(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
