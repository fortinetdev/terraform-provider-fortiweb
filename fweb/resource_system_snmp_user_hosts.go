// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure system/snmp.user/hosts.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemSnmpUserHosts() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemSnmpUserHostsRead,
		Update: resourceSystemSnmpUserHostsUpdate,
		Create: resourceSystemSnmpUserHostsCreate,
		Delete: resourceSystemSnmpUserHostsDelete,

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

func resourceSystemSnmpUserHostsCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpUserHosts: type error")
	}

	obj, err := getObjectSystemSnmpUserHosts(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpUserHosts resource while getting object: %v", err)
	}

	_, err = c.CreateSystemSnmpUserHosts(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpUserHosts resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemSnmpUserHostsRead(d, m)
}

func resourceSystemSnmpUserHostsUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpUserHosts: type error")
	}

	obj, err := getObjectSystemSnmpUserHosts(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpUserHosts resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateSystemSnmpUserHosts(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpUserHosts resource: %v", err)
	}

	return resourceSystemSnmpUserHostsRead(d, m)
}

func resourceSystemSnmpUserHostsDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpUserHosts: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteSystemSnmpUserHosts(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpUserHosts resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpUserHostsRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSnmpUserHosts: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadSystemSnmpUserHosts(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpUserHosts resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpUserHosts(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpUserHosts resource from API: %v", err)
	}

	return nil
}

func flattenSystemSnmpUserHosts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpUserHosts(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenSystemSnmpUserHosts(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("ip", flattenSystemSnmpUserHosts(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}
	if err = d.Set("mkey", flattenSystemSnmpUserHosts(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemSnmpUserHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpUserHosts(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandSystemSnmpUserHosts(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemSnmpUserHosts(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemSnmpUserHosts(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
