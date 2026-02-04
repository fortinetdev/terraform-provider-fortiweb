// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_encryption.url_encryption_policy.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlEncryptionUrlEncryptionPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlEncryptionUrlEncryptionPolicyRead,
		Update: resourceWafUrlEncryptionUrlEncryptionPolicyUpdate,
		Create: resourceWafUrlEncryptionUrlEncryptionPolicyCreate,
		Delete: resourceWafUrlEncryptionUrlEncryptionPolicyDelete,

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

func resourceWafUrlEncryptionUrlEncryptionPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionPolicy: type error")
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlEncryptionUrlEncryptionPolicy(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionPolicy resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlEncryptionUrlEncryptionPolicyRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateWafUrlEncryptionUrlEncryptionPolicy(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionPolicy resource: %v", err)
	}

	return resourceWafUrlEncryptionUrlEncryptionPolicyRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafUrlEncryptionUrlEncryptionPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlEncryptionUrlEncryptionPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlEncryptionUrlEncryptionPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafUrlEncryptionUrlEncryptionPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlEncryptionUrlEncryptionPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionPolicy resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlEncryptionUrlEncryptionPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlEncryptionUrlEncryptionPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenWafUrlEncryptionUrlEncryptionPolicy(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlEncryptionUrlEncryptionPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlEncryptionUrlEncryptionPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionPolicy(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
