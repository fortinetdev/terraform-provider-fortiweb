// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/http_constraints_exceptions.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafHttpConstraintsExceptions() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafHttpConstraintsExceptionsRead,
		Update: resourceWafHttpConstraintsExceptionsUpdate,
		Create: resourceWafHttpConstraintsExceptionsCreate,
		Delete: resourceWafHttpConstraintsExceptionsDelete,

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

func resourceWafHttpConstraintsExceptionsCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpConstraintsExceptions: type error")
	}

	obj, err := getObjectWafHttpConstraintsExceptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpConstraintsExceptions resource while getting object: %v", err)
	}

	_, err = c.CreateWafHttpConstraintsExceptions(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpConstraintsExceptions resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafHttpConstraintsExceptionsRead(d, m)
}

func resourceWafHttpConstraintsExceptionsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafHttpConstraintsExceptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpConstraintsExceptions resource while getting object: %v", err)
	}

	_, err = c.UpdateWafHttpConstraintsExceptions(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpConstraintsExceptions resource: %v", err)
	}

	return resourceWafHttpConstraintsExceptionsRead(d, m)
}

func resourceWafHttpConstraintsExceptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafHttpConstraintsExceptions(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafHttpConstraintsExceptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafHttpConstraintsExceptionsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafHttpConstraintsExceptions(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpConstraintsExceptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafHttpConstraintsExceptions(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpConstraintsExceptions resource from API: %v", err)
	}

	return nil
}

func flattenWafHttpConstraintsExceptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafHttpConstraintsExceptions(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenWafHttpConstraintsExceptions(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafHttpConstraintsExceptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafHttpConstraintsExceptions(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafHttpConstraintsExceptions(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
