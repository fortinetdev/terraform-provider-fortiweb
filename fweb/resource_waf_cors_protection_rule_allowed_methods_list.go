// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/cors_protection_rule/allowed-methods-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafCorsProtectionRuleAllowedMethodsList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafCorsProtectionRuleAllowedMethodsListRead,
		Update: resourceWafCorsProtectionRuleAllowedMethodsListUpdate,
		Create: resourceWafCorsProtectionRuleAllowedMethodsListCreate,
		Delete: resourceWafCorsProtectionRuleAllowedMethodsListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func resourceWafCorsProtectionRuleAllowedMethodsListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleAllowedMethodsList: type error")
	}

	obj, err := getObjectWafCorsProtectionRuleAllowedMethodsList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafCorsProtectionRuleAllowedMethodsList resource while getting object: %v", err)
	}

	_, err = c.CreateWafCorsProtectionRuleAllowedMethodsList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafCorsProtectionRuleAllowedMethodsList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafCorsProtectionRuleAllowedMethodsListRead(d, m)
}

func resourceWafCorsProtectionRuleAllowedMethodsListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleAllowedMethodsList: type error")
	}

	obj, err := getObjectWafCorsProtectionRuleAllowedMethodsList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafCorsProtectionRuleAllowedMethodsList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafCorsProtectionRuleAllowedMethodsList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafCorsProtectionRuleAllowedMethodsList resource: %v", err)
	}

	return resourceWafCorsProtectionRuleAllowedMethodsListRead(d, m)
}

func resourceWafCorsProtectionRuleAllowedMethodsListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleAllowedMethodsList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafCorsProtectionRuleAllowedMethodsList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafCorsProtectionRuleAllowedMethodsList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafCorsProtectionRuleAllowedMethodsListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRuleAllowedMethodsList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafCorsProtectionRuleAllowedMethodsList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafCorsProtectionRuleAllowedMethodsList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafCorsProtectionRuleAllowedMethodsList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafCorsProtectionRuleAllowedMethodsList resource from API: %v", err)
	}

	return nil
}

func flattenWafCorsProtectionRuleAllowedMethodsList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafCorsProtectionRuleAllowedMethodsList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafCorsProtectionRuleAllowedMethodsList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("method", flattenWafCorsProtectionRuleAllowedMethodsList(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafCorsProtectionRuleAllowedMethodsList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafCorsProtectionRuleAllowedMethodsList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafCorsProtectionRuleAllowedMethodsList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafCorsProtectionRuleAllowedMethodsList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("method"); ok {
		t, err := expandWafCorsProtectionRuleAllowedMethodsList(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafCorsProtectionRuleAllowedMethodsList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
