// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/allowed_origins.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafAllowedOrigins() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafAllowedOriginsRead,
		Update: resourceWafAllowedOriginsUpdate,
		Create: resourceWafAllowedOriginsCreate,
		Delete: resourceWafAllowedOriginsDelete,

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

func resourceWafAllowedOriginsCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowedOrigins: type error")
	}

	obj, err := getObjectWafAllowedOrigins(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowedOrigins resource while getting object: %v", err)
	}

	_, err = c.CreateWafAllowedOrigins(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowedOrigins resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafAllowedOriginsRead(d, m)
}

func resourceWafAllowedOriginsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafAllowedOrigins(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowedOrigins resource while getting object: %v", err)
	}

	_, err = c.UpdateWafAllowedOrigins(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowedOrigins resource: %v", err)
	}

	return resourceWafAllowedOriginsRead(d, m)
}

func resourceWafAllowedOriginsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafAllowedOrigins(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafAllowedOrigins resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafAllowedOriginsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafAllowedOrigins(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowedOrigins resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafAllowedOrigins(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowedOrigins resource from API: %v", err)
	}

	return nil
}

func flattenWafAllowedOrigins(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafAllowedOrigins(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenWafAllowedOrigins(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafAllowedOrigins(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafAllowedOrigins(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafAllowedOrigins(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
