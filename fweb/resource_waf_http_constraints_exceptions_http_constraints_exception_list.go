// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/http_constraints_exceptions/http_constraints-exception-list.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListRead,
		Update: resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListUpdate,
		Create: resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListCreate,
		Delete: resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sub_mkey": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
			"request_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"host_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_ip_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_content_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_filename_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_cookie_in_request": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_header_line_request": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_request_method_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_host_name_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"number_of_ranges_in_range_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2_max_requests": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_malformed_request": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_type_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"post_request_ctype_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_name_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_value_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_name_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_value_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_name_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_value_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_parameter_length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_name_len": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_value_len": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_name_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_value_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redundant_header_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"duplicate_paramname_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"null_byte_in_url_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_byte_in_url_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_socket_protocol_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"odd_and_even_space_attack_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internal_resource_limits_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rpc_protocol_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cl_te_coexist_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"inconsistent_cl_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"missing_host_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"range_overlapping_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multipart_formdata_bad_request_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_freq_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_bidir_concurrent_stream_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_unidir_concurrent_stream_check": &schema.Schema{
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

func resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpConstraintsExceptionsHttpConstraintsExceptionList: type error")
	}

	obj, err := getObjectWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpConstraintsExceptionsHttpConstraintsExceptionList resource while getting object: %v", err)
	}

	_, err = c.CreateWafHttpConstraintsExceptionsHttpConstraintsExceptionList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpConstraintsExceptionsHttpConstraintsExceptionList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListRead(d, m)
}

func resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListUpdate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpConstraintsExceptionsHttpConstraintsExceptionList: type error")
	}

	obj, err := getObjectWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpConstraintsExceptionsHttpConstraintsExceptionList resource while getting object: %v", err)
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	_, err = c.UpdateWafHttpConstraintsExceptionsHttpConstraintsExceptionList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpConstraintsExceptionsHttpConstraintsExceptionList resource: %v", err)
	}

	return resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListRead(d, m)
}

func resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListDelete(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpConstraintsExceptionsHttpConstraintsExceptionList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	err := c.DeleteWafHttpConstraintsExceptionsHttpConstraintsExceptionList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafHttpConstraintsExceptionsHttpConstraintsExceptionList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionListRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpConstraintsExceptionsHttpConstraintsExceptionList: type error")
	}

	mkey = mkey + "&sub_mkey=" + sub_mkey
	o, err := c.ReadWafHttpConstraintsExceptionsHttpConstraintsExceptionList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpConstraintsExceptionsHttpConstraintsExceptionList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpConstraintsExceptionsHttpConstraintsExceptionList resource from API: %v", err)
	}

	return nil
}

func flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sub_mkey", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["id"], d, "sub_mkey", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading sub_mkey: %v", err)
		}
	}
	if err = d.Set("host", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}
	if err = d.Set("request_file", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["request-file"], d, "request_file", sv)); err != nil {
		if !fortiAPIPatch(o["request-file"]) {
			return fmt.Errorf("Error reading request_file: %v", err)
		}
	}
	if err = d.Set("request_type", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["request-type"], d, "request_type", sv)); err != nil {
		if !fortiAPIPatch(o["request-type"]) {
			return fmt.Errorf("Error reading request_type: %v", err)
		}
	}
	if err = d.Set("host_status", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["host-status"], d, "host_status", sv)); err != nil {
		if !fortiAPIPatch(o["host-status"]) {
			return fmt.Errorf("Error reading host_status: %v", err)
		}
	}
	if err = d.Set("source_ip_status", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["source-ip-status"], d, "source_ip_status", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip-status"]) {
			return fmt.Errorf("Error reading source_ip_status: %v", err)
		}
	}
	if err = d.Set("source_ip", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}
	if err = d.Set("max_http_header_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-http-header-length"], d, "max_http_header_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-length"]) {
			return fmt.Errorf("Error reading max_http_header_length: %v", err)
		}
	}
	if err = d.Set("max_http_content_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-http-content-length"], d, "max_http_content_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-content-length"]) {
			return fmt.Errorf("Error reading max_http_content_length: %v", err)
		}
	}
	if err = d.Set("max_http_body_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-http-body-length"], d, "max_http_body_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-length"]) {
			return fmt.Errorf("Error reading max_http_body_length: %v", err)
		}
	}
	if err = d.Set("max_http_request_filename_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-http-request-filename-length"], d, "max_http_request_filename_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-filename-length"]) {
			return fmt.Errorf("Error reading max_http_request_filename_length: %v", err)
		}
	}
	if err = d.Set("max_http_request_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-http-request-length"], d, "max_http_request_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-length"]) {
			return fmt.Errorf("Error reading max_http_request_length: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-url-parameter-length"], d, "max_url_parameter_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-length"]) {
			return fmt.Errorf("Error reading max_url_parameter_length: %v", err)
		}
	}
	if err = d.Set("max_cookie_in_request", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-cookie-in-request"], d, "max_cookie_in_request", sv)); err != nil {
		if !fortiAPIPatch(o["max-cookie-in-request"]) {
			return fmt.Errorf("Error reading max_cookie_in_request: %v", err)
		}
	}
	if err = d.Set("max_header_line_request", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-header-line-request"], d, "max_header_line_request", sv)); err != nil {
		if !fortiAPIPatch(o["max-header-line-request"]) {
			return fmt.Errorf("Error reading max_header_line_request: %v", err)
		}
	}
	if err = d.Set("illegal_http_request_method_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Illegal-http-request-method-check"], d, "illegal_http_request_method_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-request-method-check"]) {
			return fmt.Errorf("Error reading illegal_http_request_method_check: %v", err)
		}
	}
	if err = d.Set("max_url_parameter", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-url-parameter"], d, "max_url_parameter", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter"]) {
			return fmt.Errorf("Error reading max_url_parameter: %v", err)
		}
	}
	if err = d.Set("illegal_host_name_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Illegal-host-name-check"], d, "illegal_host_name_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-host-name-check"]) {
			return fmt.Errorf("Error reading illegal_host_name_check: %v", err)
		}
	}
	if err = d.Set("number_of_ranges_in_range_header", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["number-of-ranges-in-range-header"], d, "number_of_ranges_in_range_header", sv)); err != nil {
		if !fortiAPIPatch(o["number-of-ranges-in-range-header"]) {
			return fmt.Errorf("Error reading number_of_ranges_in_range_header: %v", err)
		}
	}
	if err = d.Set("http2_max_requests", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["http2-max-requests"], d, "http2_max_requests", sv)); err != nil {
		if !fortiAPIPatch(o["http2-max-requests"]) {
			return fmt.Errorf("Error reading http2_max_requests: %v", err)
		}
	}
	if err = d.Set("block_malformed_request", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["block-malformed-request"], d, "block_malformed_request", sv)); err != nil {
		if !fortiAPIPatch(o["block-malformed-request"]) {
			return fmt.Errorf("Error reading block_malformed_request: %v", err)
		}
	}
	if err = d.Set("illegal_content_length_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Illegal-content-length-check"], d, "illegal_content_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-length-check"]) {
			return fmt.Errorf("Error reading illegal_content_length_check: %v", err)
		}
	}
	if err = d.Set("illegal_content_type_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Illegal-content-type-check"], d, "illegal_content_type_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-type-check"]) {
			return fmt.Errorf("Error reading illegal_content_type_check: %v", err)
		}
	}
	if err = d.Set("post_request_ctype_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Post-request-ctype-check"], d, "post_request_ctype_check", sv)); err != nil {
		if !fortiAPIPatch(o["Post-request-ctype-check"]) {
			return fmt.Errorf("Error reading post_request_ctype_check: %v", err)
		}
	}
	if err = d.Set("max_http_header_name_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-http-header-name-length"], d, "max_http_header_name_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-name-length"]) {
			return fmt.Errorf("Error reading max_http_header_name_length: %v", err)
		}
	}
	if err = d.Set("max_http_header_value_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-http-header-value-length"], d, "max_http_header_value_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-value-length"]) {
			return fmt.Errorf("Error reading max_http_header_value_length: %v", err)
		}
	}
	if err = d.Set("parameter_name_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["parameter-name-check"], d, "parameter_name_check", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-name-check"]) {
			return fmt.Errorf("Error reading parameter_name_check: %v", err)
		}
	}
	if err = d.Set("parameter_value_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["parameter-value-check"], d, "parameter_value_check", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-value-check"]) {
			return fmt.Errorf("Error reading parameter_value_check: %v", err)
		}
	}
	if err = d.Set("illegal_header_name_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Illegal-header-name-check"], d, "illegal_header_name_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-name-check"]) {
			return fmt.Errorf("Error reading illegal_header_name_check: %v", err)
		}
	}
	if err = d.Set("illegal_header_value_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Illegal-header-value-check"], d, "illegal_header_value_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-value-check"]) {
			return fmt.Errorf("Error reading illegal_header_value_check: %v", err)
		}
	}
	if err = d.Set("max_http_body_parameter_length", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-http-body-parameter-length"], d, "max_http_body_parameter_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-parameter-length"]) {
			return fmt.Errorf("Error reading max_http_body_parameter_length: %v", err)
		}
	}
	if err = d.Set("max_url_param_name_len", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-url-param-name-len"], d, "max_url_param_name_len", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-name-len"]) {
			return fmt.Errorf("Error reading max_url_param_name_len: %v", err)
		}
	}
	if err = d.Set("max_url_param_value_len", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["max-url-param-value-len"], d, "max_url_param_value_len", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-value-len"]) {
			return fmt.Errorf("Error reading max_url_param_value_len: %v", err)
		}
	}
	if err = d.Set("url_param_name_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["url-param-name-check"], d, "url_param_name_check", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-name-check"]) {
			return fmt.Errorf("Error reading url_param_name_check: %v", err)
		}
	}
	if err = d.Set("url_param_value_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["url-param-value-check"], d, "url_param_value_check", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-value-check"]) {
			return fmt.Errorf("Error reading url_param_value_check: %v", err)
		}
	}
	if err = d.Set("redundant_header_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["redundant-header-check"], d, "redundant_header_check", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-header-check"]) {
			return fmt.Errorf("Error reading redundant_header_check: %v", err)
		}
	}
	if err = d.Set("duplicate_paramname_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["duplicate-paramname-check"], d, "duplicate_paramname_check", sv)); err != nil {
		if !fortiAPIPatch(o["duplicate-paramname-check"]) {
			return fmt.Errorf("Error reading duplicate_paramname_check: %v", err)
		}
	}
	if err = d.Set("null_byte_in_url_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["null-byte-in-url-check"], d, "null_byte_in_url_check", sv)); err != nil {
		if !fortiAPIPatch(o["null-byte-in-url-check"]) {
			return fmt.Errorf("Error reading null_byte_in_url_check: %v", err)
		}
	}
	if err = d.Set("illegal_byte_in_url_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Illegal-byte-in-url-check"], d, "illegal_byte_in_url_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-byte-in-url-check"]) {
			return fmt.Errorf("Error reading illegal_byte_in_url_check: %v", err)
		}
	}
	if err = d.Set("web_socket_protocol_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["web-socket-protocol-check"], d, "web_socket_protocol_check", sv)); err != nil {
		if !fortiAPIPatch(o["web-socket-protocol-check"]) {
			return fmt.Errorf("Error reading web_socket_protocol_check: %v", err)
		}
	}
	if err = d.Set("odd_and_even_space_attack_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["odd-and-even-space-attack-check"], d, "odd_and_even_space_attack_check", sv)); err != nil {
		if !fortiAPIPatch(o["odd-and-even-space-attack-check"]) {
			return fmt.Errorf("Error reading odd_and_even_space_attack_check: %v", err)
		}
	}
	if err = d.Set("internal_resource_limits_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["Internal-resource-limits-check"], d, "internal_resource_limits_check", sv)); err != nil {
		if !fortiAPIPatch(o["Internal-resource-limits-check"]) {
			return fmt.Errorf("Error reading internal_resource_limits_check: %v", err)
		}
	}
	if err = d.Set("rpc_protocol_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["rpc-protocol-check"], d, "rpc_protocol_check", sv)); err != nil {
		if !fortiAPIPatch(o["rpc-protocol-check"]) {
			return fmt.Errorf("Error reading rpc_protocol_check: %v", err)
		}
	}
	if err = d.Set("cl_te_coexist_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["cl-te-coexist-check"], d, "cl_te_coexist_check", sv)); err != nil {
		if !fortiAPIPatch(o["cl-te-coexist-check"]) {
			return fmt.Errorf("Error reading cl_te_coexist_check: %v", err)
		}
	}
	if err = d.Set("inconsistent_cl_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["inconsistent-cl-check"], d, "inconsistent_cl_check", sv)); err != nil {
		if !fortiAPIPatch(o["inconsistent-cl-check"]) {
			return fmt.Errorf("Error reading inconsistent_cl_check: %v", err)
		}
	}
	if err = d.Set("missing_host_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["missing-host-check"], d, "missing_host_check", sv)); err != nil {
		if !fortiAPIPatch(o["missing-host-check"]) {
			return fmt.Errorf("Error reading missing_host_check: %v", err)
		}
	}
	if err = d.Set("range_overlapping_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["range-overlapping-check"], d, "range_overlapping_check", sv)); err != nil {
		if !fortiAPIPatch(o["range-overlapping-check"]) {
			return fmt.Errorf("Error reading range_overlapping_check: %v", err)
		}
	}
	if err = d.Set("multipart_formdata_bad_request_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["multipart-formdata-bad-request-check"], d, "multipart_formdata_bad_request_check", sv)); err != nil {
		if !fortiAPIPatch(o["multipart-formdata-bad-request-check"]) {
			return fmt.Errorf("Error reading multipart_formdata_bad_request_check: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["h2-rst-stream-check"], d, "h2_rst_stream_check", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-check"]) {
			return fmt.Errorf("Error reading h2_rst_stream_check: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_freq_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["h2-rst-stream-freq-check"], d, "h2_rst_stream_freq_check", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-freq-check"]) {
			return fmt.Errorf("Error reading h2_rst_stream_freq_check: %v", err)
		}
	}
	if err = d.Set("h3_bidir_concurrent_stream_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["h3-bidir-concurrent-stream-check"], d, "h3_bidir_concurrent_stream_check", sv)); err != nil {
		if !fortiAPIPatch(o["h3-bidir-concurrent-stream-check"]) {
			return fmt.Errorf("Error reading h3_bidir_concurrent_stream_check: %v", err)
		}
	}
	if err = d.Set("h3_unidir_concurrent_stream_check", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["h3-unidir-concurrent-stream-check"], d, "h3_unidir_concurrent_stream_check", sv)); err != nil {
		if !fortiAPIPatch(o["h3-unidir-concurrent-stream-check"]) {
			return fmt.Errorf("Error reading h3_unidir_concurrent_stream_check: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafHttpConstraintsExceptionsHttpConstraintsExceptionList(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sub_mkey"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "sub_mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}
	if v, ok := d.GetOk("host"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}
	if v, ok := d.GetOk("request_file"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "request_file", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-file"] = t
		}
	}
	if v, ok := d.GetOk("request_type"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "request_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-type"] = t
		}
	}
	if v, ok := d.GetOk("host_status"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "host_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-status"] = t
		}
	}
	if v, ok := d.GetOk("source_ip_status"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "source_ip_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-status"] = t
		}
	}
	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_http_header_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_content_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_http_content_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-content-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_http_body_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_filename_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_http_request_filename_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-filename-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_http_request_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-length"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_url_parameter_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-length"] = t
		}
	}
	if v, ok := d.GetOk("max_cookie_in_request"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_cookie_in_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie-in-request"] = t
		}
	}
	if v, ok := d.GetOk("max_header_line_request"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_header_line_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line-request"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_request_method_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "illegal_http_request_method_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-request-method-check"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_url_parameter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter"] = t
		}
	}
	if v, ok := d.GetOk("illegal_host_name_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "illegal_host_name_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-host-name-check"] = t
		}
	}
	if v, ok := d.GetOk("number_of_ranges_in_range_header"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "number_of_ranges_in_range_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-of-ranges-in-range-header"] = t
		}
	}
	if v, ok := d.GetOk("http2_max_requests"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "http2_max_requests", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-max-requests"] = t
		}
	}
	if v, ok := d.GetOk("block_malformed_request"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "block_malformed_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malformed-request"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_length_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "illegal_content_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-length-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_type_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "illegal_content_type_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-type-check"] = t
		}
	}
	if v, ok := d.GetOk("post_request_ctype_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "post_request_ctype_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Post-request-ctype-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_name_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_http_header_name_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-name-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_value_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_http_header_value_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-value-length"] = t
		}
	}
	if v, ok := d.GetOk("parameter_name_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "parameter_name_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-name-check"] = t
		}
	}
	if v, ok := d.GetOk("parameter_value_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "parameter_value_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-value-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_name_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "illegal_header_name_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-name-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_value_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "illegal_header_value_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-value-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_parameter_length"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_http_body_parameter_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-parameter-length"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_name_len"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_url_param_name_len", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-name-len"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_value_len"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "max_url_param_value_len", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-value-len"] = t
		}
	}
	if v, ok := d.GetOk("url_param_name_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "url_param_name_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-name-check"] = t
		}
	}
	if v, ok := d.GetOk("url_param_value_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "url_param_value_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-value-check"] = t
		}
	}
	if v, ok := d.GetOk("redundant_header_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "redundant_header_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-header-check"] = t
		}
	}
	if v, ok := d.GetOk("duplicate_paramname_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "duplicate_paramname_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-paramname-check"] = t
		}
	}
	if v, ok := d.GetOk("null_byte_in_url_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "null_byte_in_url_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-byte-in-url-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_byte_in_url_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "illegal_byte_in_url_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-byte-in-url-check"] = t
		}
	}
	if v, ok := d.GetOk("web_socket_protocol_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "web_socket_protocol_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-socket-protocol-check"] = t
		}
	}
	if v, ok := d.GetOk("odd_and_even_space_attack_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "odd_and_even_space_attack_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["odd-and-even-space-attack-check"] = t
		}
	}
	if v, ok := d.GetOk("internal_resource_limits_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "internal_resource_limits_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Internal-resource-limits-check"] = t
		}
	}
	if v, ok := d.GetOk("rpc_protocol_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "rpc_protocol_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-protocol-check"] = t
		}
	}
	if v, ok := d.GetOk("cl_te_coexist_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "cl_te_coexist_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cl-te-coexist-check"] = t
		}
	}
	if v, ok := d.GetOk("inconsistent_cl_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "inconsistent_cl_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inconsistent-cl-check"] = t
		}
	}
	if v, ok := d.GetOk("missing_host_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "missing_host_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-host-check"] = t
		}
	}
	if v, ok := d.GetOk("range_overlapping_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "range_overlapping_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-overlapping-check"] = t
		}
	}
	if v, ok := d.GetOk("multipart_formdata_bad_request_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "multipart_formdata_bad_request_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipart-formdata-bad-request-check"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "h2_rst_stream_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-check"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_freq_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "h2_rst_stream_freq_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-freq-check"] = t
		}
	}
	if v, ok := d.GetOk("h3_bidir_concurrent_stream_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "h3_bidir_concurrent_stream_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-bidir-concurrent-stream-check"] = t
		}
	}
	if v, ok := d.GetOk("h3_unidir_concurrent_stream_check"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "h3_unidir_concurrent_stream_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-unidir-concurrent-stream-check"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafHttpConstraintsExceptionsHttpConstraintsExceptionList(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
