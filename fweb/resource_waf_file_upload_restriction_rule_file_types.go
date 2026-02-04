// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/file_upload_restriction_rule/file-types.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafFileUploadRestrictionRuleFileTypes() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafFileUploadRestrictionRuleFileTypesRead,
		Update: resourceWafFileUploadRestrictionRuleFileTypesUpdate,
		Create: resourceWafFileUploadRestrictionRuleFileTypesCreate,
		Delete: resourceWafFileUploadRestrictionRuleFileTypesDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"file_type_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"file_type_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"group": &schema.Schema{
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

func resourceWafFileUploadRestrictionRuleFileTypesCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionRuleFileTypes: type error")
	}

	obj, err := getObjectWafFileUploadRestrictionRuleFileTypes(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafFileUploadRestrictionRuleFileTypes resource while getting object: %v", err)
	}

	resultsArray, err := c.ReadWafFileSecurityFileTypes(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafFileSecurityFileTypes: %v", err)
	}

	group := getGroup(resultsArray, (*obj)["file-type-name"].(string))
	fileTypeValue := getFileTypeValue(resultsArray, (*obj)["file-type-name"].(string))
	(*obj)["group"] = group
	(*obj)["file-type-value"] = fileTypeValue

	_, err = c.CreateWafFileUploadRestrictionRuleFileTypes(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafFileUploadRestrictionRuleFileTypes resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafFileUploadRestrictionRuleFileTypesRead(d, m)
}

func resourceWafFileUploadRestrictionRuleFileTypesUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceWafFileUploadRestrictionRuleFileTypesRead(d, m)
}

func resourceWafFileUploadRestrictionRuleFileTypesDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionRuleFileTypes: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafFileUploadRestrictionRuleFileTypes(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafFileUploadRestrictionRuleFileTypes resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafFileUploadRestrictionRuleFileTypesRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionRuleFileTypes: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafFileUploadRestrictionRuleFileTypes(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafFileUploadRestrictionRuleFileTypes resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafFileUploadRestrictionRuleFileTypes(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafFileUploadRestrictionRuleFileTypes resource from API: %v", err)
	}

	return nil
}

func flattenWafFileUploadRestrictionRuleFileTypes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafFileUploadRestrictionRuleFileTypes(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafFileUploadRestrictionRuleFileTypes(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("file_type_name", flattenWafFileUploadRestrictionRuleFileTypes(o["file-type-name"], d, "file_type_name", sv)); err != nil {
		if !fortiAPIPatch(o["file-type-name"]) {
			return fmt.Errorf("Error reading file_type_name: %v", err)
		}
	}
	if err = d.Set("file_type_value", flattenWafFileUploadRestrictionRuleFileTypes(o["file-type-value"], d, "file_type_value", sv)); err != nil {
		if !fortiAPIPatch(o["file-type-value"]) {
			return fmt.Errorf("Error reading file_type_value: %v", err)
		}
	}
	if err = d.Set("group", flattenWafFileUploadRestrictionRuleFileTypes(o["group"], d, "group", sv)); err != nil {
		if !fortiAPIPatch(o["group"]) {
			return fmt.Errorf("Error reading group: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafFileUploadRestrictionRuleFileTypes(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafFileUploadRestrictionRuleFileTypes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getGroup(resultsArray []interface{}, fileTypeName string) string {
	for _, item := range resultsArray {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if name, ok := itemMap["file-type-name"].(string); ok && name == fileTypeName {
				if group, ok := itemMap["group"].(string); ok {
					return group
				}
			}
		}
	}
	return ""
}

func getFileTypeValue(resultsArray []interface{}, fileTypeName string) string {
	for _, item := range resultsArray {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if name, ok := itemMap["file-type-name"].(string); ok && name == fileTypeName {
				if value, ok := itemMap["file-type-value"].(string); ok {
					return value
				}
			}
		}
	}
	return ""
}

func getObjectWafFileUploadRestrictionRuleFileTypes(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("file_type_name"); ok {
		t, err := expandWafFileUploadRestrictionRuleFileTypes(d, v, "file_type_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type-name"] = t
		}
	}
	if v, ok := d.GetOk("file_type_value"); ok {
		t, err := expandWafFileUploadRestrictionRuleFileTypes(d, v, "file_type_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-type-value"] = t
		}
	}
	if v, ok := d.GetOk("group"); ok {
		t, err := expandWafFileUploadRestrictionRuleFileTypes(d, v, "group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group"] = t
		}
	}

	return &obj, nil
}
