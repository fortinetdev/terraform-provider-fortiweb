// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/http_header_security_exception/list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafHttpHeaderSecurityExceptionList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafHttpHeaderSecurityExceptionListRead,
		Update: resourceWafHttpHeaderSecurityExceptionListUpdate,
		Create: resourceWafHttpHeaderSecurityExceptionListCreate,
		Delete: resourceWafHttpHeaderSecurityExceptionListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_ip_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"request_url_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"request_url_pattern": &schema.Schema{
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

func resourceWafHttpHeaderSecurityExceptionListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityExceptionList: type error")
	}

	obj, err := getObjectWafHttpHeaderSecurityExceptionList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpHeaderSecurityExceptionList resource while getting object: %v", err)
	}

	_, err = c.CreateWafHttpHeaderSecurityExceptionList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpHeaderSecurityExceptionList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafHttpHeaderSecurityExceptionListRead(d, m)
}

func resourceWafHttpHeaderSecurityExceptionListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityExceptionList: type error")
	}

	obj, err := getObjectWafHttpHeaderSecurityExceptionList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpHeaderSecurityExceptionList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafHttpHeaderSecurityExceptionList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpHeaderSecurityExceptionList resource: %v", err)
	}

	return resourceWafHttpHeaderSecurityExceptionListRead(d, m)
}

func resourceWafHttpHeaderSecurityExceptionListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityExceptionList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafHttpHeaderSecurityExceptionList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafHttpHeaderSecurityExceptionList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafHttpHeaderSecurityExceptionListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpHeaderSecurityExceptionList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafHttpHeaderSecurityExceptionList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpHeaderSecurityExceptionList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafHttpHeaderSecurityExceptionList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpHeaderSecurityExceptionList resource from API: %v", err)
	}

	return nil
}

func flattenWafHttpHeaderSecurityExceptionList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafHttpHeaderSecurityExceptionList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafHttpHeaderSecurityExceptionList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("client_ip_status", flattenWafHttpHeaderSecurityExceptionList(o["client-ip-status"], d, "client_ip_status", sv)); err != nil {
		if !fortiAPIPatch(o["client-ip-status"]) {
			return fmt.Errorf("Error reading client_ip_status: %v", err)
		}
	}
	if err = d.Set("client_ip", flattenWafHttpHeaderSecurityExceptionList(o["client-ip"], d, "client_ip", sv)); err != nil {
		if !fortiAPIPatch(o["client-ip"]) {
			return fmt.Errorf("Error reading client_ip: %v", err)
		}
	}
	if err = d.Set("request_url_type", flattenWafHttpHeaderSecurityExceptionList(o["request-url-type"], d, "request_url_type", sv)); err != nil {
		if !fortiAPIPatch(o["request-url-type"]) {
			return fmt.Errorf("Error reading request_url_type: %v", err)
		}
	}
	if err = d.Set("request_url_pattern", flattenWafHttpHeaderSecurityExceptionList(o["request-url-pattern"], d, "request_url_pattern", sv)); err != nil {
		if !fortiAPIPatch(o["request-url-pattern"]) {
			return fmt.Errorf("Error reading request_url_pattern: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafHttpHeaderSecurityExceptionList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafHttpHeaderSecurityExceptionList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafHttpHeaderSecurityExceptionList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafHttpHeaderSecurityExceptionList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("client_ip_status"); ok {
		t, err := expandWafHttpHeaderSecurityExceptionList(d, v, "client_ip_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-ip-status"] = t
		}
	}
	if v, ok := d.GetOk("client_ip"); ok {
		t, err := expandWafHttpHeaderSecurityExceptionList(d, v, "client_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-ip"] = t
		}
	}
	if v, ok := d.GetOk("request_url_type"); ok {
		t, err := expandWafHttpHeaderSecurityExceptionList(d, v, "request_url_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-url-type"] = t
		}
	}
	if v, ok := d.GetOk("request_url_pattern"); ok {
		t, err := expandWafHttpHeaderSecurityExceptionList(d, v, "request_url_pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-url-pattern"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafHttpHeaderSecurityExceptionList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
