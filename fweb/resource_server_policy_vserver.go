// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure server_policy/vserver.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerPolicyVserver() *schema.Resource {
	return &schema.Resource{
		Read:   resourceServerPolicyVserverRead,
		Update: resourceServerPolicyVserverRead,
		Create: resourceServerPolicyVserverCreate,
		Delete: resourceServerPolicyVserverDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
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

func resourceServerPolicyVserverCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing ServerPolicyVserver: type error")
	}

	obj, err := getObjectServerPolicyVserver(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyVserver resource while getting object: %v", err)
	}

	_, err = c.CreateServerPolicyVserver(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating ServerPolicyVserver resource: %v", err)
	}

	d.SetId(mkey)

	return resourceServerPolicyVserverRead(d, m)
}

func resourceServerPolicyVserverUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectServerPolicyVserver(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyVserver resource while getting object: %v", err)
	}

	_, err = c.UpdateServerPolicyVserver(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating ServerPolicyVserver resource: %v", err)
	}

	return resourceServerPolicyVserverRead(d, m)
}

func resourceServerPolicyVserverDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteServerPolicyVserver(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting ServerPolicyVserver resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceServerPolicyVserverRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadServerPolicyVserver(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyVserver resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectServerPolicyVserver(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ServerPolicyVserver resource from API: %v", err)
	}

	return nil
}

func flattenServerPolicyVserver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectServerPolicyVserver(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenServerPolicyVserver(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandServerPolicyVserver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectServerPolicyVserver(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandServerPolicyVserver(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
