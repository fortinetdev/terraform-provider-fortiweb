// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/allow_method_exceptions/allow-method-exception-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafAllowMethodExceptionsAllowMethodExceptionList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafAllowMethodExceptionsAllowMethodExceptionListRead,
		Update: resourceWafAllowMethodExceptionsAllowMethodExceptionListUpdate,
		Create: resourceWafAllowMethodExceptionsAllowMethodExceptionListCreate,
		Delete: resourceWafAllowMethodExceptionsAllowMethodExceptionListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"request_file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"request_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"allow_request": &schema.Schema{
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

func resourceWafAllowMethodExceptionsAllowMethodExceptionListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowMethodExceptionsAllowMethodExceptionList: type error")
	}

	obj, err := getObjectWafAllowMethodExceptionsAllowMethodExceptionList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowMethodExceptionsAllowMethodExceptionList resource while getting object: %v", err)
	}

	_, err = c.CreateWafAllowMethodExceptionsAllowMethodExceptionList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafAllowMethodExceptionsAllowMethodExceptionList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafAllowMethodExceptionsAllowMethodExceptionListRead(d, m)
}

func resourceWafAllowMethodExceptionsAllowMethodExceptionListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowMethodExceptionsAllowMethodExceptionList: type error")
	}

	obj, err := getObjectWafAllowMethodExceptionsAllowMethodExceptionList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowMethodExceptionsAllowMethodExceptionList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafAllowMethodExceptionsAllowMethodExceptionList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafAllowMethodExceptionsAllowMethodExceptionList resource: %v", err)
	}

	return resourceWafAllowMethodExceptionsAllowMethodExceptionListRead(d, m)
}

func resourceWafAllowMethodExceptionsAllowMethodExceptionListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowMethodExceptionsAllowMethodExceptionList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafAllowMethodExceptionsAllowMethodExceptionList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafAllowMethodExceptionsAllowMethodExceptionList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafAllowMethodExceptionsAllowMethodExceptionListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafAllowMethodExceptionsAllowMethodExceptionList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafAllowMethodExceptionsAllowMethodExceptionList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowMethodExceptionsAllowMethodExceptionList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafAllowMethodExceptionsAllowMethodExceptionList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafAllowMethodExceptionsAllowMethodExceptionList resource from API: %v", err)
	}

	return nil
}

func flattenWafAllowMethodExceptionsAllowMethodExceptionList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafAllowMethodExceptionsAllowMethodExceptionList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafAllowMethodExceptionsAllowMethodExceptionList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("host", flattenWafAllowMethodExceptionsAllowMethodExceptionList(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}
	if err = d.Set("request_file", flattenWafAllowMethodExceptionsAllowMethodExceptionList(o["request-file"], d, "request_file", sv)); err != nil {
		if !fortiAPIPatch(o["request-file"]) {
			return fmt.Errorf("Error reading request_file: %v", err)
		}
	}
	if err = d.Set("host_status", flattenWafAllowMethodExceptionsAllowMethodExceptionList(o["host-status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host-status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}
	if err = d.Set("request_type", flattenWafAllowMethodExceptionsAllowMethodExceptionList(o["request-type"], d, "request_type", sv)); err != nil {
		if !fortiAPIPatch(o["request-type"]) {
			return fmt.Errorf("Error reading request_type: %v", err)
		}
	}
	if err = d.Set("allow_request", flattenWafAllowMethodExceptionsAllowMethodExceptionList(o["allow-request"], d, "allow_request", sv)); err != nil {
		if !fortiAPIPatch(o["allow-request"]) {
			return fmt.Errorf("Error reading allow_request: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafAllowMethodExceptionsAllowMethodExceptionList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafAllowMethodExceptionsAllowMethodExceptionList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafAllowMethodExceptionsAllowMethodExceptionList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafAllowMethodExceptionsAllowMethodExceptionList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("host"); ok {
		t, err := expandWafAllowMethodExceptionsAllowMethodExceptionList(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}
	if v, ok := d.GetOk("request_file"); ok {
		t, err := expandWafAllowMethodExceptionsAllowMethodExceptionList(d, v, "request_file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-file"] = t
		}
	}
	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandWafAllowMethodExceptionsAllowMethodExceptionList(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-status"] = t
		}
	}
	if v, ok := d.GetOk("request_type"); ok {
		t, err := expandWafAllowMethodExceptionsAllowMethodExceptionList(d, v, "request_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-type"] = t
		}
	}
	if v, ok := d.GetOk("allow_request"); ok {
		t, err := expandWafAllowMethodExceptionsAllowMethodExceptionList(d, v, "allow_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-request"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafAllowMethodExceptionsAllowMethodExceptionList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
