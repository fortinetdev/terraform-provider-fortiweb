// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/url_access.url_access_parameter/url-access-parameter-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafUrlAccessUrlAccessParameterUrlAccessParameterList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListRead,
		Update: resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListUpdate,
		Create: resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListCreate,
		Delete: resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"argument_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"argument_name_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"argument_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"type_checked": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"data_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"argument_expression": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"custom_data_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"optional": &schema.Schema{
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

func resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessParameterUrlAccessParameterList: type error")
	}

	obj, err := getObjectWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessParameterUrlAccessParameterList resource while getting object: %v", err)
	}

	_, err = c.CreateWafUrlAccessUrlAccessParameterUrlAccessParameterList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafUrlAccessUrlAccessParameterUrlAccessParameterList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListRead(d, m)
}

func resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessParameterUrlAccessParameterList: type error")
	}

	obj, err := getObjectWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessParameterUrlAccessParameterList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafUrlAccessUrlAccessParameterUrlAccessParameterList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafUrlAccessUrlAccessParameterUrlAccessParameterList resource: %v", err)
	}

	return resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListRead(d, m)
}

func resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessParameterUrlAccessParameterList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafUrlAccessUrlAccessParameterUrlAccessParameterList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafUrlAccessUrlAccessParameterUrlAccessParameterList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafUrlAccessUrlAccessParameterUrlAccessParameterListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafUrlAccessUrlAccessParameterUrlAccessParameterList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafUrlAccessUrlAccessParameterUrlAccessParameterList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessParameterUrlAccessParameterList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafUrlAccessUrlAccessParameterUrlAccessParameterList resource from API: %v", err)
	}

	return nil
}

func flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafUrlAccessUrlAccessParameterUrlAccessParameterList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading id: %v", err)
		}
	}
	if err = d.Set("argument_type", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["argument-type"], d, "argument_type", sv)); err != nil {
		if !fortiAPIPatch(o["argument-type"]) {
			return fmt.Errorf("Error reading argument_type: %v", err)
		}
	}
	if err = d.Set("argument_name_type", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["argument-name-type"], d, "argument_name_type", sv)); err != nil {
		if !fortiAPIPatch(o["argument-name-type"]) {
			return fmt.Errorf("Error reading argument_name_type: %v", err)
		}
	}
	if err = d.Set("argument_name", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["argument-name"], d, "argument_name", sv)); err != nil {
		if !fortiAPIPatch(o["argument-name"]) {
			return fmt.Errorf("Error reading argument_name: %v", err)
		}
	}
	if err = d.Set("type_checked", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["type-checked"], d, "type_checked", sv)); err != nil {
		if !fortiAPIPatch(o["type-checked"]) {
			return fmt.Errorf("Error reading type_checked: %v", err)
		}
	}
	if err = d.Set("data_type", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["data-type"], d, "data_type", sv)); err != nil {
		if !fortiAPIPatch(o["data-type"]) {
			return fmt.Errorf("Error reading data_type: %v", err)
		}
	}
	if err = d.Set("argument_expression", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["argument-expression"], d, "argument_expression", sv)); err != nil {
		if !fortiAPIPatch(o["argument-expression"]) {
			return fmt.Errorf("Error reading argument_expression: %v", err)
		}
	}
	if err = d.Set("custom_data_type", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["custom-data-type"], d, "custom_data_type", sv)); err != nil {
		if !fortiAPIPatch(o["custom-data-type"]) {
			return fmt.Errorf("Error reading custom_data_type: %v", err)
		}
	}
	if err = d.Set("optional", flattenWafUrlAccessUrlAccessParameterUrlAccessParameterList(o["optional"], d, "optional", sv)); err != nil {
		if !fortiAPIPatch(o["optional"]) {
			return fmt.Errorf("Error reading optional: %v", err)
		}
	}

	return nil
}

func expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafUrlAccessUrlAccessParameterUrlAccessParameterList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("argument_type"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "argument_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["argument-type"] = t
		}
	}
	if v, ok := d.GetOk("argument_name_type"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "argument_name_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["argument-name-type"] = t
		}
	}
	if v, ok := d.GetOk("argument_name"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "argument_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["argument-name"] = t
		}
	}
	if v, ok := d.GetOk("type_checked"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "type_checked", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type-checked"] = t
		}
	}
	if v, ok := d.GetOk("data_type"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "data_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["data-type"] = t
		}
	}
	if v, ok := d.GetOk("argument_expression"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "argument_expression", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["argument-expression"] = t
		}
	}
	if v, ok := d.GetOk("custom_data_type"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "custom_data_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-data-type"] = t
		}
	}
	if v, ok := d.GetOk("optional"); ok {
		t, err := expandWafUrlAccessUrlAccessParameterUrlAccessParameterList(d, v, "optional", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["optional"] = t
		}
	}

	return &obj, nil
}
