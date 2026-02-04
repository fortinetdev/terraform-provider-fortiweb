// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/http_header_security/http-header-security-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafHttpHeaderSecurityHttpHeaderSecurityList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafHttpHeaderSecurityHttpHeaderSecurityListRead,
		Update: resourceWafHttpHeaderSecurityHttpHeaderSecurityListUpdate,
		Create: resourceWafHttpHeaderSecurityHttpHeaderSecurityListCreate,
		Delete: resourceWafHttpHeaderSecurityHttpHeaderSecurityListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
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
			"request_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"nonce": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"reporting_endpoints_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"report_to_value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception": &schema.Schema{
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

func resourceWafHttpHeaderSecurityHttpHeaderSecurityListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityHttpHeaderSecurityList: type error")
	}

	obj, err := getObjectWafHttpHeaderSecurityHttpHeaderSecurityList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpHeaderSecurityHttpHeaderSecurityList resource while getting object: %v", err)
	}

	_, err = c.CreateWafHttpHeaderSecurityHttpHeaderSecurityList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpHeaderSecurityHttpHeaderSecurityList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafHttpHeaderSecurityHttpHeaderSecurityListRead(d, m)
}

func resourceWafHttpHeaderSecurityHttpHeaderSecurityListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityHttpHeaderSecurityList: type error")
	}

	obj, err := getObjectWafHttpHeaderSecurityHttpHeaderSecurityList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpHeaderSecurityHttpHeaderSecurityList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafHttpHeaderSecurityHttpHeaderSecurityList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpHeaderSecurityHttpHeaderSecurityList resource: %v", err)
	}

	return resourceWafHttpHeaderSecurityHttpHeaderSecurityListRead(d, m)
}

func resourceWafHttpHeaderSecurityHttpHeaderSecurityListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityHttpHeaderSecurityList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafHttpHeaderSecurityHttpHeaderSecurityList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafHttpHeaderSecurityHttpHeaderSecurityList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafHttpHeaderSecurityHttpHeaderSecurityListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityHttpHeaderSecurityList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafHttpHeaderSecurityHttpHeaderSecurityList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpHeaderSecurityHttpHeaderSecurityList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafHttpHeaderSecurityHttpHeaderSecurityList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpHeaderSecurityHttpHeaderSecurityList resource from API: %v", err)
	}

	return nil
}

func flattenWafHttpHeaderSecurityHttpHeaderSecurityList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafHttpHeaderSecurityHttpHeaderSecurityList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("name", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}
	if err = d.Set("request_type", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["request-type"], d, "request_type", sv)); err != nil {
		if !fortiAPIPatch(o["request-type"]) {
			return fmt.Errorf("Error reading request_type: %v", err)
		}
	}
	if err = d.Set("request_file", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["request-file"], d, "request_file", sv)); err != nil {
		if !fortiAPIPatch(o["request-file"]) {
			return fmt.Errorf("Error reading request_file: %v", err)
		}
	}
	if err = d.Set("request_status", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["request-status"], d, "request_status", sv)); err != nil {
		if !fortiAPIPatch(o["request-status"]) {
			return fmt.Errorf("Error reading request_status: %v", err)
		}
	}
	if err = d.Set("value", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["value"], d, "value", sv)); err != nil {
		if !fortiAPIPatch(o["value"]) {
			return fmt.Errorf("Error reading value: %v", err)
		}
	}
	if err = d.Set("nonce", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["nonce"], d, "nonce", sv)); err != nil {
		if !fortiAPIPatch(o["nonce"]) {
			return fmt.Errorf("Error reading nonce: %v", err)
		}
	}
	if err = d.Set("reporting_endpoints_value", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["reporting-endpoints-value"], d, "reporting_endpoints_value", sv)); err != nil {
		if !fortiAPIPatch(o["reporting-endpoints-value"]) {
			return fmt.Errorf("Error reading reporting_endpoints_value: %v", err)
		}
	}
	if err = d.Set("report_to_value", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["report-to-value"], d, "report_to_value", sv)); err != nil {
		if !fortiAPIPatch(o["report-to-value"]) {
			return fmt.Errorf("Error reading report_to_value: %v", err)
		}
	}
	if err = d.Set("exception", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["exception"], d, "exception", sv)); err != nil {
		if !fortiAPIPatch(o["exception"]) {
			return fmt.Errorf("Error reading exception: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafHttpHeaderSecurityHttpHeaderSecurityList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafHttpHeaderSecurityHttpHeaderSecurityList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafHttpHeaderSecurityHttpHeaderSecurityList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("name"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}
	if v, ok := d.GetOk("request_type"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "request_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-type"] = t
		}
	}
	if v, ok := d.GetOk("request_file"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "request_file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-file"] = t
		}
	}
	if v, ok := d.GetOk("request_status"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "request_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-status"] = t
		}
	}
	if v, ok := d.GetOk("value"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value"] = t
		}
	}
	if v, ok := d.GetOk("nonce"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "nonce", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nonce"] = t
		}
	}
	if v, ok := d.GetOk("reporting_endpoints_value"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "reporting_endpoints_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reporting-endpoints-value"] = t
		}
	}
	if v, ok := d.GetOk("report_to_value"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "report_to_value", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-to-value"] = t
		}
	}
	if v, ok := d.GetOk("exception"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "exception", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafHttpHeaderSecurityHttpHeaderSecurityList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
