// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_encryption.url_encryption_rule/exceptions.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlEncryptionUrlEncryptionRuleExceptions() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlEncryptionUrlEncryptionRuleExceptionsRead,
		Update: resourceWafUrlEncryptionUrlEncryptionRuleExceptionsUpdate,
		Create: resourceWafUrlEncryptionUrlEncryptionRuleExceptionsCreate,
		Delete: resourceWafUrlEncryptionUrlEncryptionRuleExceptionsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_pattern": &schema.Schema{
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

func resourceWafUrlEncryptionUrlEncryptionRuleExceptionsCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRuleExceptions: type error")
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionRuleExceptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionRuleExceptions resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlEncryptionUrlEncryptionRuleExceptions(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionRuleExceptions resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlEncryptionUrlEncryptionRuleExceptionsRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionRuleExceptionsUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRuleExceptions: type error")
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionRuleExceptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionRuleExceptions resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafUrlEncryptionUrlEncryptionRuleExceptions(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionRuleExceptions resource: %v", err)
	}

	return resourceWafUrlEncryptionUrlEncryptionRuleExceptionsRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionRuleExceptionsDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRuleExceptions: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafUrlEncryptionUrlEncryptionRuleExceptions(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlEncryptionUrlEncryptionRuleExceptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlEncryptionUrlEncryptionRuleExceptionsRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRuleExceptions: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafUrlEncryptionUrlEncryptionRuleExceptions(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionRuleExceptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlEncryptionUrlEncryptionRuleExceptions(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionRuleExceptions resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlEncryptionUrlEncryptionRuleExceptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlEncryptionUrlEncryptionRuleExceptions(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafUrlEncryptionUrlEncryptionRuleExceptions(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("url_type", flattenWafUrlEncryptionUrlEncryptionRuleExceptions(o["url-type"], d, "url_type", sv)); err != nil {
		if !fortiAPIPatch(o["url-type"]) {
			return fmt.Errorf("Error reading url_type: %v", err)
		}
	}
	if err = d.Set("url_pattern", flattenWafUrlEncryptionUrlEncryptionRuleExceptions(o["url-pattern"], d, "url_pattern", sv)); err != nil {
		if !fortiAPIPatch(o["url-pattern"]) {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafUrlEncryptionUrlEncryptionRuleExceptions(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlEncryptionUrlEncryptionRuleExceptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlEncryptionUrlEncryptionRuleExceptions(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRuleExceptions(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("url_type"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRuleExceptions(d, v, "url_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-type"] = t
		}
	}
	if v, ok := d.GetOk("url_pattern"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRuleExceptions(d, v, "url_pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRuleExceptions(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
