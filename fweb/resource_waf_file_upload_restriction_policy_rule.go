// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/file_upload_restriction_policy/rule.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafFileUploadRestrictionPolicyRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafFileUploadRestrictionPolicyRuleRead,
		Update: resourceWafFileUploadRestrictionPolicyRuleUpdate,
		Create: resourceWafFileUploadRestrictionPolicyRuleCreate,
		Delete: resourceWafFileUploadRestrictionPolicyRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"file_upload_restriction_rule": &schema.Schema{
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

func resourceWafFileUploadRestrictionPolicyRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionPolicyRule: type error")
	}

	obj, err := getObjectWafFileUploadRestrictionPolicyRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafFileUploadRestrictionPolicyRule resource while getting object: %v", err)
	}

	_, err = c.CreateWafFileUploadRestrictionPolicyRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafFileUploadRestrictionPolicyRule resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafFileUploadRestrictionPolicyRuleRead(d, m)
}

func resourceWafFileUploadRestrictionPolicyRuleUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionPolicyRule: type error")
	}

	obj, err := getObjectWafFileUploadRestrictionPolicyRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafFileUploadRestrictionPolicyRule resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafFileUploadRestrictionPolicyRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafFileUploadRestrictionPolicyRule resource: %v", err)
	}

	return resourceWafFileUploadRestrictionPolicyRuleRead(d, m)
}

func resourceWafFileUploadRestrictionPolicyRuleDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionPolicyRule: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafFileUploadRestrictionPolicyRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafFileUploadRestrictionPolicyRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafFileUploadRestrictionPolicyRuleRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionPolicyRule: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafFileUploadRestrictionPolicyRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafFileUploadRestrictionPolicyRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafFileUploadRestrictionPolicyRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafFileUploadRestrictionPolicyRule resource from API: %v", err)
	}

	return nil
}

func flattenWafFileUploadRestrictionPolicyRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafFileUploadRestrictionPolicyRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafFileUploadRestrictionPolicyRule(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("file_upload_restriction_rule", flattenWafFileUploadRestrictionPolicyRule(o["file-upload-restriction-rule"], d, "file_upload_restriction_rule", sv)); err != nil {
		if !fortiAPIPatch(o["file-upload-restriction-rule"]) {
			return fmt.Errorf("Error reading file_upload_restriction_rule: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafFileUploadRestrictionPolicyRule(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafFileUploadRestrictionPolicyRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafFileUploadRestrictionPolicyRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafFileUploadRestrictionPolicyRule(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("file_upload_restriction_rule"); ok {
		t, err := expandWafFileUploadRestrictionPolicyRule(d, v, "file_upload_restriction_rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-upload-restriction-rule"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafFileUploadRestrictionPolicyRule(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
