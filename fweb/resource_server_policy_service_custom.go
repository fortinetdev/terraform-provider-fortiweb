// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure server_policy/service.custom.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerPolicyServiceCustom() *schema.Resource {
	return &schema.Resource{
		Read:   resourceServerPolicyServiceCustomRead,
		Update: resourceServerPolicyServiceCustomUpdate,
		Create: resourceServerPolicyServiceCustomCreate,
		Delete: resourceServerPolicyServiceCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"protocol": &schema.Schema{
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

func resourceServerPolicyServiceCustomCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing ServerPolicyServiceCustom: type error")
	}

	obj, err := getObjectServerPolicyServiceCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyServiceCustom resource while getting object: %v", err)
	}

	_, err = c.CreateServerPolicyServiceCustom(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyServiceCustom resource: %v", err)
	}

	d.SetId(mkey)

	return resourceServerPolicyServiceCustomRead(d, m)
}

func resourceServerPolicyServiceCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectServerPolicyServiceCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyServiceCustom resource while getting object: %v", err)
	}

	_, err = c.UpdateServerPolicyServiceCustom(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyServiceCustom resource: %v", err)
	}

	return resourceServerPolicyServiceCustomRead(d, m)
}

func resourceServerPolicyServiceCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteServerPolicyServiceCustom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting ServerPolicyServiceCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceServerPolicyServiceCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadServerPolicyServiceCustom(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyServiceCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectServerPolicyServiceCustom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyServiceCustom resource from API: %v", err)
	}

	return nil
}

func flattenServerPolicyServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectServerPolicyServiceCustom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("port", flattenServerPolicyServiceCustom(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}
	if err = d.Set("protocol", flattenServerPolicyServiceCustom(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}
	if err = d.Set("mkey", flattenServerPolicyServiceCustom(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandServerPolicyServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectServerPolicyServiceCustom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port"); ok {
		t, err := expandServerPolicyServiceCustom(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}
	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandServerPolicyServiceCustom(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandServerPolicyServiceCustom(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
