// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_encryption.url_encryption_policy/rule-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlEncryptionUrlEncryptionPolicyRuleList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlEncryptionUrlEncryptionPolicyRuleListRead,
		Update: resourceWafUrlEncryptionUrlEncryptionPolicyRuleListUpdate,
		Create: resourceWafUrlEncryptionUrlEncryptionPolicyRuleListCreate,
		Delete: resourceWafUrlEncryptionUrlEncryptionPolicyRuleListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rule": &schema.Schema{
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

func resourceWafUrlEncryptionUrlEncryptionPolicyRuleListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionPolicyRuleList: type error")
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionPolicyRuleList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionPolicyRuleList resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlEncryptionUrlEncryptionPolicyRuleList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionPolicyRuleList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlEncryptionUrlEncryptionPolicyRuleListRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionPolicyRuleListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionPolicyRuleList: type error")
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionPolicyRuleList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionPolicyRuleList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafUrlEncryptionUrlEncryptionPolicyRuleList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionPolicyRuleList resource: %v", err)
	}

	return resourceWafUrlEncryptionUrlEncryptionPolicyRuleListRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionPolicyRuleListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionPolicyRuleList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafUrlEncryptionUrlEncryptionPolicyRuleList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlEncryptionUrlEncryptionPolicyRuleList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlEncryptionUrlEncryptionPolicyRuleListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionPolicyRuleList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafUrlEncryptionUrlEncryptionPolicyRuleList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionPolicyRuleList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlEncryptionUrlEncryptionPolicyRuleList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionPolicyRuleList resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlEncryptionUrlEncryptionPolicyRuleList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlEncryptionUrlEncryptionPolicyRuleList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafUrlEncryptionUrlEncryptionPolicyRuleList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("rule", flattenWafUrlEncryptionUrlEncryptionPolicyRuleList(o["rule"], d, "rule", sv)); err != nil {
		if !fortiAPIPatch(o["rule"]) {
			return fmt.Errorf("Error reading rule: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafUrlEncryptionUrlEncryptionPolicyRuleList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlEncryptionUrlEncryptionPolicyRuleList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlEncryptionUrlEncryptionPolicyRuleList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionPolicyRuleList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("rule"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionPolicyRuleList(d, v, "rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rule"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionPolicyRuleList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
