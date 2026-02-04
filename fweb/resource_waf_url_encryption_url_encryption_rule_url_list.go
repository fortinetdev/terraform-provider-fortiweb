// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_encryption.url_encryption_rule/url-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlEncryptionUrlEncryptionRuleUrlList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlEncryptionUrlEncryptionRuleUrlListRead,
		Update: resourceWafUrlEncryptionUrlEncryptionRuleUrlListUpdate,
		Create: resourceWafUrlEncryptionUrlEncryptionRuleUrlListCreate,
		Delete: resourceWafUrlEncryptionUrlEncryptionRuleUrlListDelete,

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

func resourceWafUrlEncryptionUrlEncryptionRuleUrlListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRuleUrlList: type error")
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionRuleUrlList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionRuleUrlList resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlEncryptionUrlEncryptionRuleUrlList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlEncryptionUrlEncryptionRuleUrlList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlEncryptionUrlEncryptionRuleUrlListRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionRuleUrlListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRuleUrlList: type error")
	}

	obj, err := getObjectWafUrlEncryptionUrlEncryptionRuleUrlList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionRuleUrlList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafUrlEncryptionUrlEncryptionRuleUrlList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlEncryptionUrlEncryptionRuleUrlList resource: %v", err)
	}

	return resourceWafUrlEncryptionUrlEncryptionRuleUrlListRead(d, m)
}

func resourceWafUrlEncryptionUrlEncryptionRuleUrlListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRuleUrlList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafUrlEncryptionUrlEncryptionRuleUrlList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlEncryptionUrlEncryptionRuleUrlList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlEncryptionUrlEncryptionRuleUrlListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlEncryptionUrlEncryptionRuleUrlList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafUrlEncryptionUrlEncryptionRuleUrlList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionRuleUrlList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlEncryptionUrlEncryptionRuleUrlList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlEncryptionUrlEncryptionRuleUrlList resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlEncryptionUrlEncryptionRuleUrlList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlEncryptionUrlEncryptionRuleUrlList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafUrlEncryptionUrlEncryptionRuleUrlList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("url_type", flattenWafUrlEncryptionUrlEncryptionRuleUrlList(o["url-type"], d, "url_type", sv)); err != nil {
		if !fortiAPIPatch(o["url-type"]) {
			return fmt.Errorf("Error reading url_type: %v", err)
		}
	}
	if err = d.Set("url_pattern", flattenWafUrlEncryptionUrlEncryptionRuleUrlList(o["url-pattern"], d, "url_pattern", sv)); err != nil {
		if !fortiAPIPatch(o["url-pattern"]) {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafUrlEncryptionUrlEncryptionRuleUrlList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafUrlEncryptionUrlEncryptionRuleUrlList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlEncryptionUrlEncryptionRuleUrlList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRuleUrlList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("url_type"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRuleUrlList(d, v, "url_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-type"] = t
		}
	}
	if v, ok := d.GetOk("url_pattern"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRuleUrlList(d, v, "url_pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafUrlEncryptionUrlEncryptionRuleUrlList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
