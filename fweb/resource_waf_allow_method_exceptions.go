// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/allow_method_exceptions.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafAllowMethodExceptions() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafAllowMethodExceptionsRead,
		Update: resourceWafAllowMethodExceptionsUpdate,
		Create: resourceWafAllowMethodExceptionsCreate,
		Delete: resourceWafAllowMethodExceptionsDelete,

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

func resourceWafAllowMethodExceptionsCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowMethodExceptions: type error")
	}

	obj, err := getObjectWafAllowMethodExceptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowMethodExceptions resource while getting object: %v", err)
	}

	_, err = c.CreateWafAllowMethodExceptions(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowMethodExceptions resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafAllowMethodExceptionsRead(d, m)
}

func resourceWafAllowMethodExceptionsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafAllowMethodExceptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowMethodExceptions resource while getting object: %v", err)
	}

	_, err = c.UpdateWafAllowMethodExceptions(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowMethodExceptions resource: %v", err)
	}

	return resourceWafAllowMethodExceptionsRead(d, m)
}

func resourceWafAllowMethodExceptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafAllowMethodExceptions(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafAllowMethodExceptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafAllowMethodExceptionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafAllowMethodExceptions(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowMethodExceptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafAllowMethodExceptions(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowMethodExceptions resource from API: %v", err)
	}

	return nil
}

func flattenWafAllowMethodExceptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafAllowMethodExceptions(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenWafAllowMethodExceptions(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafAllowMethodExceptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafAllowMethodExceptions(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafAllowMethodExceptions(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
