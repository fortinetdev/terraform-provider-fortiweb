// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/http_header_security.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafHttpHeaderSecurity() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafHttpHeaderSecurityRead,
		Update: resourceWafHttpHeaderSecurityUpdate,
		Create: resourceWafHttpHeaderSecurityCreate,
		Delete: resourceWafHttpHeaderSecurityDelete,

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

func resourceWafHttpHeaderSecurityCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurity: type error")
	}

	obj, err := getObjectWafHttpHeaderSecurity(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpHeaderSecurity resource while getting object: %v", err)
	}

	_, err = c.CreateWafHttpHeaderSecurity(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpHeaderSecurity resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafHttpHeaderSecurityRead(d, m)
}

func resourceWafHttpHeaderSecurityUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceWafHttpHeaderSecurityRead(d, m)
}

func resourceWafHttpHeaderSecurityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafHttpHeaderSecurity(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafHttpHeaderSecurity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafHttpHeaderSecurityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafHttpHeaderSecurity(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpHeaderSecurity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafHttpHeaderSecurity(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpHeaderSecurity resource from API: %v", err)
	}

	return nil
}

func flattenWafHttpHeaderSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafHttpHeaderSecurity(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenWafHttpHeaderSecurity(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafHttpHeaderSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafHttpHeaderSecurity(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafHttpHeaderSecurity(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
