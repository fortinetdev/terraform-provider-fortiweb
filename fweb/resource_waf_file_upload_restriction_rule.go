// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/file_upload_restriction_rule.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafFileUploadRestrictionRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafFileUploadRestrictionRuleRead,
		Update: resourceWafFileUploadRestrictionRuleUpdate,
		Create: resourceWafFileUploadRestrictionRuleCreate,
		Delete: resourceWafFileUploadRestrictionRuleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
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
			//(0-102400)(KBytes)
			"file_size_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"file_uncompress": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"json_file_support": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"json_key_for_filename": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"json_key_field": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"octet_stream_filename_position": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"octet_stream_filename_string": &schema.Schema{
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

func resourceWafFileUploadRestrictionRuleCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionRule: type error")
	}

	obj, err := getObjectWafFileUploadRestrictionRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafFileUploadRestrictionRule resource while getting object: %v", err)
	}

	_, err = c.CreateWafFileUploadRestrictionRule(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafFileUploadRestrictionRule resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafFileUploadRestrictionRuleRead(d, m)
}

func resourceWafFileUploadRestrictionRuleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafFileUploadRestrictionRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafFileUploadRestrictionRule resource while getting object: %v", err)
	}

	_, err = c.UpdateWafFileUploadRestrictionRule(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafFileUploadRestrictionRule resource: %v", err)
	}

	return resourceWafFileUploadRestrictionRuleRead(d, m)
}

func resourceWafFileUploadRestrictionRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafFileUploadRestrictionRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafFileUploadRestrictionRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafFileUploadRestrictionRuleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafFileUploadRestrictionRule(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafFileUploadRestrictionRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafFileUploadRestrictionRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafFileUploadRestrictionRule resource from API: %v", err)
	}

	return nil
}

func flattenWafFileUploadRestrictionRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafFileUploadRestrictionRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("host_status", flattenWafFileUploadRestrictionRule(o["host-status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host-status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}
	if err = d.Set("host", flattenWafFileUploadRestrictionRule(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}
	if err = d.Set("request_type", flattenWafFileUploadRestrictionRule(o["request-type"], d, "request_type", sv)); err != nil {
		if !fortiAPIPatch(o["request-type"]) {
			return fmt.Errorf("Error reading request_type: %v", err)
		}
	}
	if err = d.Set("request_file", flattenWafFileUploadRestrictionRule(o["request-file"], d, "request_file", sv)); err != nil {
		if !fortiAPIPatch(o["request-file"]) {
			return fmt.Errorf("Error reading request_file: %v", err)
		}
	}
	if err = d.Set("file_size_limit", flattenWafFileUploadRestrictionRule(o["file-size-limit"], d, "file_size_limit", sv)); err != nil {
		if !fortiAPIPatch(o["file-size-limit"]) {
			return fmt.Errorf("Error reading file_size_limit: %v", err)
		}
	}
	if err = d.Set("type", flattenWafFileUploadRestrictionRule(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}
	if err = d.Set("file_uncompress", flattenWafFileUploadRestrictionRule(o["file-uncompress"], d, "file_uncompress", sv)); err != nil {
		if !fortiAPIPatch(o["file-uncompress"]) {
			return fmt.Errorf("Error reading file_uncompress: %v", err)
		}
	}
	if err = d.Set("json_file_support", flattenWafFileUploadRestrictionRule(o["json-file-support"], d, "json_file_support", sv)); err != nil {
		if !fortiAPIPatch(o["json-file-support"]) {
			return fmt.Errorf("Error reading json_file_support: %v", err)
		}
	}
	if err = d.Set("json_key_for_filename", flattenWafFileUploadRestrictionRule(o["json-key-for-filename"], d, "json_key_for_filename", sv)); err != nil {
		if !fortiAPIPatch(o["json-key-for-filename"]) {
			return fmt.Errorf("Error reading json_key_for_filename: %v", err)
		}
	}
	if err = d.Set("json_key_field", flattenWafFileUploadRestrictionRule(o["json-key-field"], d, "json_key_field", sv)); err != nil {
		if !fortiAPIPatch(o["json-key-field"]) {
			return fmt.Errorf("Error reading json_key_field: %v", err)
		}
	}
	if err = d.Set("octet_stream_filename_string", flattenWafFileUploadRestrictionRule(o["octet-stream-filename-string"], d, "octet_stream_filename_string", sv)); err != nil {
		if !fortiAPIPatch(o["octet-stream-filename-string"]) {
			return fmt.Errorf("Error reading octet_stream_filename_string: %v", err)
		}
	}
	if err = d.Set("octet_stream_filename_position", flattenWafFileUploadRestrictionRule(o["octet-stream-filename-position"], d, "octet_stream_filename_position", sv)); err != nil {
		if !fortiAPIPatch(o["octet-stream-filename-position"]) {
			return fmt.Errorf("Error reading octet_stream_filename_position: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafFileUploadRestrictionRule(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafFileUploadRestrictionRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafFileUploadRestrictionRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-status"] = t
		}
	}
	if v, ok := d.GetOk("host"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}
	if v, ok := d.GetOk("request_type"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "request_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-type"] = t
		}
	}
	if v, ok := d.GetOk("request_file"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "request_file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-file"] = t
		}
	}
	if v, ok := d.GetOk("file_size_limit"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "file_size_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-size-limit"] = t
		}
	}
	if v, ok := d.GetOk("type"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}
	if v, ok := d.GetOk("file_uncompress"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "file_uncompress", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-uncompress"] = t
		}
	}
	if v, ok := d.GetOk("json_file_support"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "json_file_support", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["json-file-support"] = t
		}
	}
	if v, ok := d.GetOk("json_key_for_filename"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "json_key_for_filename", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["json-key-for-filename"] = t
		}
	}
	if v, ok := d.GetOk("json_key_field"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "json_key_field", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["json-key-field"] = t
		}
	}
	if v, ok := d.GetOk("octet_stream_filename_string"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "octet_stream_filename_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["octet-stream-filename-string"] = t
		}
	}
	if v, ok := d.GetOk("octet_stream_filename_position"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "octet_stream_filename_position", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["octet-stream-filename-position"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafFileUploadRestrictionRule(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
