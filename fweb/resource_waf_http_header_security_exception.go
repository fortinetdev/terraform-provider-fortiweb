// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/http_header_security_exception.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafHttpHeaderSecurityException() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafHttpHeaderSecurityExceptionRead,
		Update: resourceWafHttpHeaderSecurityExceptionUpdate,
		Create: resourceWafHttpHeaderSecurityExceptionCreate,
		Delete: resourceWafHttpHeaderSecurityExceptionDelete,

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

func resourceWafHttpHeaderSecurityExceptionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityException: type error")
	}

	obj, err := getObjectWafHttpHeaderSecurityException(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpHeaderSecurityException resource while getting object: %v", err)
	}

	_, err = c.CreateWafHttpHeaderSecurityException(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpHeaderSecurityException resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafHttpHeaderSecurityExceptionRead(d, m)
}

func resourceWafHttpHeaderSecurityExceptionUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceWafHttpHeaderSecurityExceptionRead(d, m)
}

func resourceWafHttpHeaderSecurityExceptionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafHttpHeaderSecurityException(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafHttpHeaderSecurityException resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafHttpHeaderSecurityExceptionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafHttpHeaderSecurityException(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpHeaderSecurityException resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafHttpHeaderSecurityException(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpHeaderSecurityException resource from API: %v", err)
	}

	return nil
}

func flattenWafHttpHeaderSecurityException(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafHttpHeaderSecurityException(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenWafHttpHeaderSecurityException(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafHttpHeaderSecurityException(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafHttpHeaderSecurityException(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafHttpHeaderSecurityException(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
