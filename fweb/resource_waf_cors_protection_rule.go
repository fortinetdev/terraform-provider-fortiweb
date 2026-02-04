// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/cors_protection_rule.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafCorsProtectionRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafCorsProtectionRuleRead,
		Update: resourceWafCorsProtectionRuleUpdate,
		Create: resourceWafCorsProtectionRuleCreate,
		Delete: resourceWafCorsProtectionRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"request_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"request_file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_cors_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_origins_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_credentials": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_maximum_age": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"allowed_methods": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allowed_headers": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exposed_headers": &schema.Schema{
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

func resourceWafCorsProtectionRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafCorsProtectionRule: type error")
	}

	obj, err := getObjectWafCorsProtectionRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafCorsProtectionRule resource while getting object: %v", err)
	}

	_, err = c.CreateWafCorsProtectionRule(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafCorsProtectionRule resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafCorsProtectionRuleRead(d, m)
}

func resourceWafCorsProtectionRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafCorsProtectionRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafCorsProtectionRule resource while getting object: %v", err)
	}

	_, err = c.UpdateWafCorsProtectionRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafCorsProtectionRule resource: %v", err)
	}

	return resourceWafCorsProtectionRuleRead(d, m)
}

func resourceWafCorsProtectionRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafCorsProtectionRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafCorsProtectionRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafCorsProtectionRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafCorsProtectionRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafCorsProtectionRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafCorsProtectionRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafCorsProtectionRule resource from API: %v", err)
	}

	return nil
}

func flattenWafCorsProtectionRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafCorsProtectionRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("host_status", flattenWafCorsProtectionRule(o["host-status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host-status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}
	if err = d.Set("host", flattenWafCorsProtectionRule(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}
	if err = d.Set("request_type", flattenWafCorsProtectionRule(o["request-type"], d, "request_type", sv)); err != nil {
		if !fortiAPIPatch(o["request-type"]) {
			return fmt.Errorf("Error reading request_type: %v", err)
		}
	}
	if err = d.Set("request_file", flattenWafCorsProtectionRule(o["request-file"], d, "request_file", sv)); err != nil {
		if !fortiAPIPatch(o["request-file"]) {
			return fmt.Errorf("Error reading request_file: %v", err)
		}
	}
	if err = d.Set("block_cors_traffic", flattenWafCorsProtectionRule(o["block-cors-traffic"], d, "block_cors_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["block-cors-traffic"]) {
			return fmt.Errorf("Error reading block_cors_traffic: %v", err)
		}
	}
	if err = d.Set("allowed_origins_list", flattenWafCorsProtectionRule(o["allowed-origins-list"], d, "allowed_origins_list", sv)); err != nil {
		if !fortiAPIPatch(o["allowed-origins-list"]) {
			return fmt.Errorf("Error reading allowed_origins_list: %v", err)
		}
	}
	if err = d.Set("allowed_methods", flattenWafCorsProtectionRule(o["allowed-methods"], d, "allowed_methods", sv)); err != nil {
		if !fortiAPIPatch(o["allowed-methods"]) {
			return fmt.Errorf("Error reading allowed_methods: %v", err)
		}
	}
	if err = d.Set("method", flattenWafCorsProtectionRule(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}
	if err = d.Set("allowed_headers", flattenWafCorsProtectionRule(o["allowed-headers"], d, "allowed_headers", sv)); err != nil {
		if !fortiAPIPatch(o["allowed-headers"]) {
			return fmt.Errorf("Error reading allowed_headers: %v", err)
		}
	}
	if err = d.Set("allowed_credentials", flattenWafCorsProtectionRule(o["allowed-credentials"], d, "allowed_credentials", sv)); err != nil {
		if !fortiAPIPatch(o["allowed-credentials"]) {
			return fmt.Errorf("Error reading allowed_credentials: %v", err)
		}
	}
	if err = d.Set("allowed_maximum_age", flattenWafCorsProtectionRule(o["allowed-maximum-age"], d, "allowed_maximum_age", sv)); err != nil {
		if !fortiAPIPatch(o["allowed-maximum-age"]) {
			return fmt.Errorf("Error reading allowed_maximum_age: %v", err)
		}
	}
	if err = d.Set("exposed_headers", flattenWafCorsProtectionRule(o["exposed-headers"], d, "exposed_headers", sv)); err != nil {
		if !fortiAPIPatch(o["exposed-headers"]) {
			return fmt.Errorf("Error reading exposed_headers: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafCorsProtectionRule(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafCorsProtectionRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafCorsProtectionRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-status"] = t
		}
	}
	if v, ok := d.GetOk("host"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}
	if v, ok := d.GetOk("request_type"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "request_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-type"] = t
		}
	}
	if v, ok := d.GetOk("request_file"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "request_file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-file"] = t
		}
	}
	if v, ok := d.GetOk("block_cors_traffic"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "block_cors_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-cors-traffic"] = t
		}
	}
	if v, ok := d.GetOk("allowed_origins_list"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "allowed_origins_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-origins-list"] = t
		}
	}
	if v, ok := d.GetOk("allowed_methods"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "allowed_methods", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-methods"] = t
		}
	}
	if v, ok := d.GetOk("method"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}
	if v, ok := d.GetOk("allowed_headers"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "allowed_headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-headers"] = t
		}
	}
	if v, ok := d.GetOk("allowed_credentials"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "allowed_credentials", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-credentials"] = t
		}
	}
	if v, ok := d.GetOk("allowed_maximum_age"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "allowed_maximum_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-maximum-age"] = t
		}
	}
	if v, ok := d.GetOk("exposed_headers"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "exposed_headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exposed-headers"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafCorsProtectionRule(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
