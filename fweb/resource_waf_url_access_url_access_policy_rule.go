// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_access.url_access_policy/rule.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlAccessUrlAccessPolicyRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlAccessUrlAccessPolicyRuleRead,
		Update: resourceWafUrlAccessUrlAccessPolicyRuleUpdate,
		Create: resourceWafUrlAccessUrlAccessPolicyRuleCreate,
		Delete: resourceWafUrlAccessUrlAccessPolicyRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url_access_rule_name": &schema.Schema{
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

func resourceWafUrlAccessUrlAccessPolicyRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessPolicyRule: type error")
	}

	obj, err := getObjectWafUrlAccessUrlAccessPolicyRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessPolicyRule resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlAccessUrlAccessPolicyRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessPolicyRule resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlAccessUrlAccessPolicyRuleRead(d, m)
}

func resourceWafUrlAccessUrlAccessPolicyRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessPolicyRule: type error")
	}

	obj, err := getObjectWafUrlAccessUrlAccessPolicyRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessPolicyRule resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafUrlAccessUrlAccessPolicyRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessPolicyRule resource: %v", err)
	}

	return resourceWafUrlAccessUrlAccessPolicyRuleRead(d, m)
}

func resourceWafUrlAccessUrlAccessPolicyRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessPolicyRule: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafUrlAccessUrlAccessPolicyRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlAccessUrlAccessPolicyRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlAccessUrlAccessPolicyRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessPolicyRule: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafUrlAccessUrlAccessPolicyRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessPolicyRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlAccessUrlAccessPolicyRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessPolicyRule resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlAccessUrlAccessPolicyRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlAccessUrlAccessPolicyRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafUrlAccessUrlAccessPolicyRule(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("url_access_rule_name", flattenWafUrlAccessUrlAccessPolicyRule(o["url-access-rule-name"], d, "url_access_rule_name", sv)); err != nil {
		if !fortiAPIPatch(o["url-access-rule-name"]) {
			return fmt.Errorf("Error reading url_access_rule_name: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafUrlAccessUrlAccessPolicyRule(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlAccessUrlAccessPolicyRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlAccessUrlAccessPolicyRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafUrlAccessUrlAccessPolicyRule(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("url_access_rule_name"); ok {
		t, err := expandWafUrlAccessUrlAccessPolicyRule(d, v, "url_access_rule_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-access-rule-name"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlAccessUrlAccessPolicyRule(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
