// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/file_upload_restriction_policy.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafFileUploadRestrictionPolicy() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafFileUploadRestrictionPolicyRead,
		Update: resourceWafFileUploadRestrictionPolicyUpdate,
		Create: resourceWafFileUploadRestrictionPolicyCreate,
		Delete: resourceWafFileUploadRestrictionPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"av_scan": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"signature_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"fortisandbox_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"hold_session_while_scanning_file": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"icap_server_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exchange_mail_detection": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"owa_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"activesync_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mapi_protocol": &schema.Schema{
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

func resourceWafFileUploadRestrictionPolicyCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafFileUploadRestrictionPolicy: type error")
	}

	obj, err := getObjectWafFileUploadRestrictionPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafFileUploadRestrictionPolicy resource while getting object: %v", err)
	}

	_, err = c.CreateWafFileUploadRestrictionPolicy(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafFileUploadRestrictionPolicy resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafFileUploadRestrictionPolicyRead(d, m)
}

func resourceWafFileUploadRestrictionPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafFileUploadRestrictionPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafFileUploadRestrictionPolicy resource while getting object: %v", err)
	}

	_, err = c.UpdateWafFileUploadRestrictionPolicy(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafFileUploadRestrictionPolicy resource: %v", err)
	}

	return resourceWafFileUploadRestrictionPolicyRead(d, m)
}

func resourceWafFileUploadRestrictionPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafFileUploadRestrictionPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafFileUploadRestrictionPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafFileUploadRestrictionPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafFileUploadRestrictionPolicy(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafFileUploadRestrictionPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafFileUploadRestrictionPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafFileUploadRestrictionPolicy resource from API: %v", err)
	}

	return nil
}

func flattenWafFileUploadRestrictionPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafFileUploadRestrictionPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("action", flattenWafFileUploadRestrictionPolicy(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}
	if err = d.Set("block_period", flattenWafFileUploadRestrictionPolicy(o["block-period"], d, "block_period", sv)); err != nil {
		if !fortiAPIPatch(o["block-period"]) {
			return fmt.Errorf("Error reading block_period: %v", err)
		}
	}
	if err = d.Set("severity", flattenWafFileUploadRestrictionPolicy(o["severity"], d, "severity", sv)); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}
	if err = d.Set("trigger", flattenWafFileUploadRestrictionPolicy(o["trigger"], d, "trigger", sv)); err != nil {
		if !fortiAPIPatch(o["trigger"]) {
			return fmt.Errorf("Error reading trigger: %v", err)
		}
	}
	if err = d.Set("av_scan", flattenWafFileUploadRestrictionPolicy(o["av-scan"], d, "av_scan", sv)); err != nil {
		if !fortiAPIPatch(o["av-scan"]) {
			return fmt.Errorf("Error reading av_scan: %v", err)
		}
	}
	if err = d.Set("signature_check", flattenWafFileUploadRestrictionPolicy(o["signature-check"], d, "signature_check", sv)); err != nil {
		if !fortiAPIPatch(o["signature-check"]) {
			return fmt.Errorf("Error reading signature_check: %v", err)
		}
	}
	if err = d.Set("fortisandbox_check", flattenWafFileUploadRestrictionPolicy(o["fortisandbox-check"], d, "fortisandbox_check", sv)); err != nil {
		if !fortiAPIPatch(o["fortisandbox-check"]) {
			return fmt.Errorf("Error reading fortisandbox_check: %v", err)
		}
	}
	if err = d.Set("hold_session_while_scanning_file", flattenWafFileUploadRestrictionPolicy(o["hold-session-while-scanning-file"], d, "hold_session_while_scanning_file", sv)); err != nil {
		if !fortiAPIPatch(o["hold-session-while-scanning-file"]) {
			return fmt.Errorf("Error reading hold_session_while_scanning_file: %v", err)
		}
	}
	if err = d.Set("icap_server_check", flattenWafFileUploadRestrictionPolicy(o["icap-server-check"], d, "icap_server_check", sv)); err != nil {
		if !fortiAPIPatch(o["icap-server-check"]) {
			return fmt.Errorf("Error reading icap_server_check: %v", err)
		}
	}
	if err = d.Set("exchange_mail_detection", flattenWafFileUploadRestrictionPolicy(o["exchange-mail-detection"], d, "exchange_mail_detection", sv)); err != nil {
		if !fortiAPIPatch(o["exchange-mail-detection"]) {
			return fmt.Errorf("Error reading exchange_mail_detection: %v", err)
		}
	}
	if err = d.Set("owa_protocol", flattenWafFileUploadRestrictionPolicy(o["owa-protocol"], d, "owa_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["owa-protocol"]) {
			return fmt.Errorf("Error reading owa_protocol: %v", err)
		}
	}
	if err = d.Set("activesync_protocol", flattenWafFileUploadRestrictionPolicy(o["activesync-protocol"], d, "activesync_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["activesync-protocol"]) {
			return fmt.Errorf("Error reading activesync_protocol: %v", err)
		}
	}
	if err = d.Set("mapi_protocol", flattenWafFileUploadRestrictionPolicy(o["mapi-protocol"], d, "mapi_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["mapi-protocol"]) {
			return fmt.Errorf("Error reading mapi_protocol: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafFileUploadRestrictionPolicy(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafFileUploadRestrictionPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafFileUploadRestrictionPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("action"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}
	if v, ok := d.GetOk("block_period"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-period"] = t
		}
	}
	if v, ok := d.GetOk("severity"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}
	if v, ok := d.GetOk("trigger"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trigger"] = t
		}
	}
	if v, ok := d.GetOk("av_scan"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "av_scan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-scan"] = t
		}
	}
	if v, ok := d.GetOk("signature_check"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "signature_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature-check"] = t
		}
	}
	if v, ok := d.GetOk("fortisandbox_check"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "fortisandbox_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortisandbox-check"] = t
		}
	}
	if v, ok := d.GetOk("hold_session_while_scanning_file"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "hold_session_while_scanning_file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hold-session-while-scanning-file"] = t
		}
	}
	if v, ok := d.GetOk("icap_server_check"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "icap_server_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-server-check"] = t
		}
	}
	if v, ok := d.GetOk("exchange_mail_detection"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "exchange_mail_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exchange-mail-detection"] = t
		}
	}
	if v, ok := d.GetOk("owa_protocol"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "owa_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owa-protocol"] = t
		}
	}
	if v, ok := d.GetOk("activesync_protocol"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "activesync_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activesync-protocol"] = t
		}
	}
	if v, ok := d.GetOk("mapi_protocol"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "mapi_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi-protocol"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafFileUploadRestrictionPolicy(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
