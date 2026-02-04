// Copyright 2025 Fortinet, Inc. All rights reserved.
// Author: Wayne Chou
// Description: Configure waf/http_protocol_parameter_restriction.

package fortiweb

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceWafHttpProtocolParameterRestriction() *schema.Resource {
	return &schema.Resource{
		Read:   resourceWafHttpProtocolParameterRestrictionRead,
		Update: resourceWafHttpProtocolParameterRestrictionUpdate,
		Create: resourceWafHttpProtocolParameterRestrictionCreate,
		Delete: resourceWafHttpProtocolParameterRestrictionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"max_http_header_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_header_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_header_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_content_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_content_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_content_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_content_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_content_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_content_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_content_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_body_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_body_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_request_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_request_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_version_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_version_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_version_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_http_version_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_version_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_version_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_cookie_in_request_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_cookie_in_request": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_cookie_in_request_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_cookie_in_request_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_cookie_in_request_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_cookie_in_request_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_cookie_in_request_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_header_line_request_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_header_line_request": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_header_line_request_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_header_line_request_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_header_line_request_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_header_line_request_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_header_line_request_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_request_method_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_request_method_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_request_method_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_http_request_method_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_request_method_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_http_request_method_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_parameter_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_host_name_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_host_name_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_host_name_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_host_name_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_host_name_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_host_name_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"number_of_ranges_in_range_header_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"number_of_ranges_in_range_header": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"number_of_ranges_in_range_header_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"number_of_ranges_in_range_header_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"number_of_ranges_in_range_header_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"number_of_ranges_in_range_header_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"number_of_ranges_in_range_header_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2_max_requests_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2_max_requests": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"http2_max_requests_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2_max_requests_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"http2_max_requests_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2_max_requests_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"http2_max_requests_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"exception_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_malformed_request_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_malformed_request_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_malformed_request_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"block_malformed_request_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_malformed_request_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"block_malformed_request_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_length_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_length_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_content_length_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_length_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_length_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_type_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_type_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_type_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_content_type_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_type_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_content_type_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_response_code_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_response_code_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_response_code_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_response_code_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_response_code_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_response_code_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"post_request_ctype_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"post_request_ctype_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"post_request_ctype_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"post_request_ctype_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"post_request_ctype_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"post_request_ctype_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_name_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_name_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_header_name_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_name_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_header_name_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_name_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_name_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_value_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_value_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_header_value_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_value_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_header_value_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_value_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_header_value_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_name_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_name_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_name_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"parameter_name_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_name_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_name_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_value_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_value_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_value_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"parameter_value_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_value_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"parameter_value_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_name_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_name_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_name_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_header_name_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_name_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_name_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_value_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_value_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_value_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_header_value_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_value_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_header_value_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_parameter_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_parameter_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_body_parameter_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_parameter_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_body_parameter_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_parameter_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_body_parameter_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_filename_length_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_filename_length": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_request_filename_length_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_filename_length_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_http_request_filename_length_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_filename_length_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_http_request_filename_length_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_socket_protocol_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_socket_protocol_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_socket_protocol_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"web_socket_protocol_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_socket_protocol_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"web_socket_protocol_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_table_size_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_table_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_table_size_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_table_size_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_table_size_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_table_size_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_table_size_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_current_streams_num_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_current_streams_num": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_current_streams_num_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_current_streams_num_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_current_streams_num_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_current_streams_num_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_current_streams_num_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_initial_window_size_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_initial_window_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_initial_window_size_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_initial_window_size_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_initial_window_size_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_initial_window_size_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_initial_window_size_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_frame_size_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_frame_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_frame_size_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_frame_size_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_frame_size_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_frame_size_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_frame_size_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_list_size_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_list_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_list_size_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_list_size_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_list_size_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_list_size_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_setting_header_list_size_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_name_len_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_name_len": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_url_param_name_len_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_name_len_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_url_param_name_len_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_name_len_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_name_len_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_value_len_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_value_len": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_url_param_value_len_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_value_len_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"max_url_param_value_len_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_value_len_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"max_url_param_value_len_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_name_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_name_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_name_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"url_param_name_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_name_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_name_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_value_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_value_check_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_value_check_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"url_param_value_check_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_value_check_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"url_param_value_check_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"null_byte_in_url_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"null_byte_in_url_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"null_byte_in_url_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"null_byte_in_url_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"null_byte_in_url_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"null_byte_in_url_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_byte_in_url_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_byte_in_url_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_byte_in_url_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"illegal_byte_in_url_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_byte_in_url_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"illegal_byte_in_url_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"malformed_url_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"malformed_url_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"malformed_url_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"malformed_url_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"malformed_url_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"malformed_url_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redundant_header_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redundant_header_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redundant_header_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"redundant_header_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redundant_header_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"redundant_header_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"chunk_size_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"chunk_size_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"chunk_size_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"chunk_size_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"chunk_size_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"chunk_size_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internal_resource_limits_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internal_resource_limits_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internal_resource_limits_block_period": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internal_resource_limits_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internal_resource_limits_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"internal_resource_limits_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rpc_protocol_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rpc_protocol_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rpc_protocol_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"rpc_protocol_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rpc_protocol_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"rpc_protocol_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"duplicate_paramname_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"duplicate_paramname_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"duplicate_paramname_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"duplicate_paramname_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"duplicate_paramname_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"duplicate_paramname_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"odd_and_even_space_attack_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"odd_and_even_space_attack_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"odd_and_even_space_attack_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"odd_and_even_space_attack_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"odd_and_even_space_attack_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"odd_and_even_space_attack_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cl_te_coexist_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cl_te_coexist_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cl_te_coexist_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"cl_te_coexist_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cl_te_coexist_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"cl_te_coexist_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"inconsistent_cl_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"inconsistent_cl_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"inconsistent_cl_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"inconsistent_cl_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"inconsistent_cl_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"inconsistent_cl_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"missing_host_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"missing_host_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"missing_host_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"missing_host_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"missing_host_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"missing_host_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"range_overlapping_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"range_overlapping_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"range_overlapping_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"range_overlapping_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"range_overlapping_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"range_overlapping_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multipart_formdata_bad_request_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multipart_formdata_bad_request_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multipart_formdata_bad_request_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"multipart_formdata_bad_request_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multipart_formdata_bad_request_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"multipart_formdata_bad_request_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_freq_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_freq": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_freq_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_freq_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_freq_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_freq_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h2_rst_stream_freq_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_table_capacity_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_table_capacity": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_max_table_capacity_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_table_capacity_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_max_table_capacity_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_table_capacity_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_table_capacity_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_field_section_size_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_field_section_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_max_field_section_size_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_field_section_size_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_max_field_section_size_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_field_section_size_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_max_field_section_size_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_block_streams_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_block_streams": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_block_streams_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_block_streams_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_block_streams_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_block_streams_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_block_streams_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_bidir_concurrent_stream_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_bidir_concurrent_stream": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_bidir_concurrent_stream_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_bidir_concurrent_stream_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_bidir_concurrent_stream_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_bidir_concurrent_stream_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_bidir_concurrent_stream_threat_weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_unidir_concurrent_stream_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_unidir_concurrent_stream": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_unidir_concurrent_stream_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_unidir_concurrent_stream_block_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
				Optional: true,
			},
			"h3_unidir_concurrent_stream_severity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_unidir_concurrent_stream_trigger": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"h3_unidir_concurrent_stream_threat_weight": &schema.Schema{
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

func resourceWafHttpProtocolParameterRestrictionCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing WafHttpProtocolParameterRestriction: type error")
	}

	obj, err := getObjectWafHttpProtocolParameterRestriction(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpProtocolParameterRestriction resource while getting object: %v", err)
	}

	_, err = c.CreateWafHttpProtocolParameterRestriction(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating WafHttpProtocolParameterRestriction resource: %v", err)
	}

	d.SetId(mkey)

	return resourceWafHttpProtocolParameterRestrictionRead(d, m)
}

func resourceWafHttpProtocolParameterRestrictionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectWafHttpProtocolParameterRestriction(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpProtocolParameterRestriction resource while getting object: %v", err)
	}

	_, err = c.UpdateWafHttpProtocolParameterRestriction(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating WafHttpProtocolParameterRestriction resource: %v", err)
	}

	return resourceWafHttpProtocolParameterRestrictionRead(d, m)
}

func resourceWafHttpProtocolParameterRestrictionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteWafHttpProtocolParameterRestriction(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting WafHttpProtocolParameterRestriction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafHttpProtocolParameterRestrictionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadWafHttpProtocolParameterRestriction(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpProtocolParameterRestriction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafHttpProtocolParameterRestriction(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafHttpProtocolParameterRestriction resource from API: %v", err)
	}

	return nil
}

func flattenWafHttpProtocolParameterRestriction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafHttpProtocolParameterRestriction(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_http_header_length_check", flattenWafHttpProtocolParameterRestriction(o["max-http-header-length-check"], d, "max_http_header_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-length-check"]) {
			return fmt.Errorf("Error reading max_http_header_length_check: %v", err)
		}
	}
	if err = d.Set("max_http_header_length", flattenWafHttpProtocolParameterRestriction(o["max-http-header-length"], d, "max_http_header_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-length"]) {
			return fmt.Errorf("Error reading max_http_header_length: %v", err)
		}
	}
	if err = d.Set("max_http_header_length_action", flattenWafHttpProtocolParameterRestriction(o["max-http-header-length-action"], d, "max_http_header_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-length-action"]) {
			return fmt.Errorf("Error reading max_http_header_length_action: %v", err)
		}
	}
	if err = d.Set("max_http_header_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-http-header-length-block-period"], d, "max_http_header_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-length-block-period"]) {
			return fmt.Errorf("Error reading max_http_header_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_http_header_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-http-header-length-threat-weight"], d, "max_http_header_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_http_header_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_http_header_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-http-header-length-severity"], d, "max_http_header_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-length-severity"]) {
			return fmt.Errorf("Error reading max_http_header_length_severity: %v", err)
		}
	}
	if err = d.Set("max_http_header_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-http-header-length-trigger"], d, "max_http_header_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-length-trigger"]) {
			return fmt.Errorf("Error reading max_http_header_length_trigger: %v", err)
		}
	}
	if err = d.Set("max_http_content_length_check", flattenWafHttpProtocolParameterRestriction(o["max-http-content-length-check"], d, "max_http_content_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-content-length-check"]) {
			return fmt.Errorf("Error reading max_http_content_length_check: %v", err)
		}
	}
	if err = d.Set("max_http_content_length", flattenWafHttpProtocolParameterRestriction(o["max-http-content-length"], d, "max_http_content_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-content-length"]) {
			return fmt.Errorf("Error reading max_http_content_length: %v", err)
		}
	}
	if err = d.Set("max_http_content_length_action", flattenWafHttpProtocolParameterRestriction(o["max-http-content-length-action"], d, "max_http_content_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-content-length-action"]) {
			return fmt.Errorf("Error reading max_http_content_length_action: %v", err)
		}
	}
	if err = d.Set("max_http_content_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-http-content-length-block-period"], d, "max_http_content_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-content-length-block-period"]) {
			return fmt.Errorf("Error reading max_http_content_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_http_content_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-http-content-length-threat-weight"], d, "max_http_content_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-content-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_http_content_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_http_content_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-http-content-length-severity"], d, "max_http_content_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-content-length-severity"]) {
			return fmt.Errorf("Error reading max_http_content_length_severity: %v", err)
		}
	}
	if err = d.Set("max_http_content_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-http-content-length-trigger"], d, "max_http_content_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-content-length-trigger"]) {
			return fmt.Errorf("Error reading max_http_content_length_trigger: %v", err)
		}
	}
	if err = d.Set("max_http_body_length_check", flattenWafHttpProtocolParameterRestriction(o["max-http-body-length-check"], d, "max_http_body_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-length-check"]) {
			return fmt.Errorf("Error reading max_http_body_length_check: %v", err)
		}
	}
	if err = d.Set("max_http_body_length", flattenWafHttpProtocolParameterRestriction(o["max-http-body-length"], d, "max_http_body_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-length"]) {
			return fmt.Errorf("Error reading max_http_body_length: %v", err)
		}
	}
	if err = d.Set("max_http_body_length_action", flattenWafHttpProtocolParameterRestriction(o["max-http-body-length-action"], d, "max_http_body_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-length-action"]) {
			return fmt.Errorf("Error reading max_http_body_length_action: %v", err)
		}
	}
	if err = d.Set("max_http_body_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-http-body-length-block-period"], d, "max_http_body_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-length-block-period"]) {
			return fmt.Errorf("Error reading max_http_body_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_http_body_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-http-body-length-threat-weight"], d, "max_http_body_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_http_body_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_http_body_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-http-body-length-severity"], d, "max_http_body_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-length-severity"]) {
			return fmt.Errorf("Error reading max_http_body_length_severity: %v", err)
		}
	}
	if err = d.Set("max_http_body_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-http-body-length-trigger"], d, "max_http_body_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-length-trigger"]) {
			return fmt.Errorf("Error reading max_http_body_length_trigger: %v", err)
		}
	}
	if err = d.Set("max_http_request_length_check", flattenWafHttpProtocolParameterRestriction(o["max-http-request-length-check"], d, "max_http_request_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-length-check"]) {
			return fmt.Errorf("Error reading max_http_request_length_check: %v", err)
		}
	}
	if err = d.Set("max_http_request_length", flattenWafHttpProtocolParameterRestriction(o["max-http-request-length"], d, "max_http_request_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-length"]) {
			return fmt.Errorf("Error reading max_http_request_length: %v", err)
		}
	}
	if err = d.Set("max_http_request_length_action", flattenWafHttpProtocolParameterRestriction(o["max-http-request-length-action"], d, "max_http_request_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-length-action"]) {
			return fmt.Errorf("Error reading max_http_request_length_action: %v", err)
		}
	}
	if err = d.Set("max_http_request_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-http-request-length-block-period"], d, "max_http_request_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-length-block-period"]) {
			return fmt.Errorf("Error reading max_http_request_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_http_request_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-http-request-length-threat-weight"], d, "max_http_request_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_http_request_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_http_request_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-http-request-length-severity"], d, "max_http_request_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-length-severity"]) {
			return fmt.Errorf("Error reading max_http_request_length_severity: %v", err)
		}
	}
	if err = d.Set("max_http_request_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-http-request-length-trigger"], d, "max_http_request_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-length-trigger"]) {
			return fmt.Errorf("Error reading max_http_request_length_trigger: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_length_check", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-length-check"], d, "max_url_parameter_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-length-check"]) {
			return fmt.Errorf("Error reading max_url_parameter_length_check: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_length", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-length"], d, "max_url_parameter_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-length"]) {
			return fmt.Errorf("Error reading max_url_parameter_length: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_length_action", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-length-action"], d, "max_url_parameter_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-length-action"]) {
			return fmt.Errorf("Error reading max_url_parameter_length_action: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-length-block-period"], d, "max_url_parameter_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-length-block-period"]) {
			return fmt.Errorf("Error reading max_url_parameter_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-length-threat-weight"], d, "max_url_parameter_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_url_parameter_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-length-severity"], d, "max_url_parameter_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-length-severity"]) {
			return fmt.Errorf("Error reading max_url_parameter_length_severity: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-length-trigger"], d, "max_url_parameter_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-length-trigger"]) {
			return fmt.Errorf("Error reading max_url_parameter_length_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_http_version_check", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-version-check"], d, "illegal_http_version_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-version-check"]) {
			return fmt.Errorf("Error reading illegal_http_version_check: %v", err)
		}
	}
	if err = d.Set("illegal_http_version_check_action", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-version-check-action"], d, "illegal_http_version_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-version-check-action"]) {
			return fmt.Errorf("Error reading illegal_http_version_check_action: %v", err)
		}
	}
	if err = d.Set("illegal_http_version_check_block_period", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-version-check-block-period"], d, "illegal_http_version_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-version-check-block-period"]) {
			return fmt.Errorf("Error reading illegal_http_version_check_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_http_version_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-version-threat-weight"], d, "illegal_http_version_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-version-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_http_version_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_http_version_check_severity", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-version-check-severity"], d, "illegal_http_version_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-version-check-severity"]) {
			return fmt.Errorf("Error reading illegal_http_version_check_severity: %v", err)
		}
	}
	if err = d.Set("illegal_http_version_check_trigger", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-version-check-trigger"], d, "illegal_http_version_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-version-check-trigger"]) {
			return fmt.Errorf("Error reading illegal_http_version_check_trigger: %v", err)
		}
	}
	if err = d.Set("max_cookie_in_request_check", flattenWafHttpProtocolParameterRestriction(o["max-cookie-in-request-check"], d, "max_cookie_in_request_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-cookie-in-request-check"]) {
			return fmt.Errorf("Error reading max_cookie_in_request_check: %v", err)
		}
	}
	if err = d.Set("max_cookie_in_request", flattenWafHttpProtocolParameterRestriction(o["max-cookie-in-request"], d, "max_cookie_in_request", sv)); err != nil {
		if !fortiAPIPatch(o["max-cookie-in-request"]) {
			return fmt.Errorf("Error reading max_cookie_in_request: %v", err)
		}
	}
	if err = d.Set("max_cookie_in_request_action", flattenWafHttpProtocolParameterRestriction(o["max-cookie-in-request-action"], d, "max_cookie_in_request_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-cookie-in-request-action"]) {
			return fmt.Errorf("Error reading max_cookie_in_request_action: %v", err)
		}
	}
	if err = d.Set("max_cookie_in_request_block_period", flattenWafHttpProtocolParameterRestriction(o["max-cookie-in-request-block-period"], d, "max_cookie_in_request_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-cookie-in-request-block-period"]) {
			return fmt.Errorf("Error reading max_cookie_in_request_block_period: %v", err)
		}
	}
	if err = d.Set("max_cookie_in_request_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-cookie-in-request-threat-weight"], d, "max_cookie_in_request_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-cookie-in-request-threat-weight"]) {
			return fmt.Errorf("Error reading max_cookie_in_request_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_cookie_in_request_severity", flattenWafHttpProtocolParameterRestriction(o["max-cookie-in-request-severity"], d, "max_cookie_in_request_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-cookie-in-request-severity"]) {
			return fmt.Errorf("Error reading max_cookie_in_request_severity: %v", err)
		}
	}
	if err = d.Set("max_cookie_in_request_trigger", flattenWafHttpProtocolParameterRestriction(o["max-cookie-in-request-trigger"], d, "max_cookie_in_request_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-cookie-in-request-trigger"]) {
			return fmt.Errorf("Error reading max_cookie_in_request_trigger: %v", err)
		}
	}
	if err = d.Set("max_header_line_request_check", flattenWafHttpProtocolParameterRestriction(o["max-header-line-request-check"], d, "max_header_line_request_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-header-line-request-check"]) {
			return fmt.Errorf("Error reading max_header_line_request_check: %v", err)
		}
	}
	if err = d.Set("max_header_line_request", flattenWafHttpProtocolParameterRestriction(o["max-header-line-request"], d, "max_header_line_request", sv)); err != nil {
		if !fortiAPIPatch(o["max-header-line-request"]) {
			return fmt.Errorf("Error reading max_header_line_request: %v", err)
		}
	}
	if err = d.Set("max_header_line_request_action", flattenWafHttpProtocolParameterRestriction(o["max-header-line-request-action"], d, "max_header_line_request_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-header-line-request-action"]) {
			return fmt.Errorf("Error reading max_header_line_request_action: %v", err)
		}
	}
	if err = d.Set("max_header_line_request_block_period", flattenWafHttpProtocolParameterRestriction(o["max-header-line-request-block-period"], d, "max_header_line_request_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-header-line-request-block-period"]) {
			return fmt.Errorf("Error reading max_header_line_request_block_period: %v", err)
		}
	}
	if err = d.Set("max_header_line_request_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-header-line-request-threat-weight"], d, "max_header_line_request_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-header-line-request-threat-weight"]) {
			return fmt.Errorf("Error reading max_header_line_request_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_header_line_request_severity", flattenWafHttpProtocolParameterRestriction(o["max-header-line-request-severity"], d, "max_header_line_request_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-header-line-request-severity"]) {
			return fmt.Errorf("Error reading max_header_line_request_severity: %v", err)
		}
	}
	if err = d.Set("max_header_line_request_trigger", flattenWafHttpProtocolParameterRestriction(o["max-header-line-request-trigger"], d, "max_header_line_request_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-header-line-request-trigger"]) {
			return fmt.Errorf("Error reading max_header_line_request_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_http_request_method_check", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-request-method-check"], d, "illegal_http_request_method_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-request-method-check"]) {
			return fmt.Errorf("Error reading illegal_http_request_method_check: %v", err)
		}
	}
	if err = d.Set("illegal_http_request_method_action", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-request-method-action"], d, "illegal_http_request_method_action", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-request-method-action"]) {
			return fmt.Errorf("Error reading illegal_http_request_method_action: %v", err)
		}
	}
	if err = d.Set("illegal_http_request_method_block_period", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-request-method-block-period"], d, "illegal_http_request_method_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-request-method-block-period"]) {
			return fmt.Errorf("Error reading illegal_http_request_method_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_http_request_method_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-request-method-threat-weight"], d, "illegal_http_request_method_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-request-method-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_http_request_method_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_http_request_method_severity", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-request-method-severity"], d, "illegal_http_request_method_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-request-method-severity"]) {
			return fmt.Errorf("Error reading illegal_http_request_method_severity: %v", err)
		}
	}
	if err = d.Set("illegal_http_request_method_trigger", flattenWafHttpProtocolParameterRestriction(o["Illegal-http-request-method-trigger"], d, "illegal_http_request_method_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-http-request-method-trigger"]) {
			return fmt.Errorf("Error reading illegal_http_request_method_trigger: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_check", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-check"], d, "max_url_parameter_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-check"]) {
			return fmt.Errorf("Error reading max_url_parameter_check: %v", err)
		}
	}
	if err = d.Set("max_url_parameter", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter"], d, "max_url_parameter", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter"]) {
			return fmt.Errorf("Error reading max_url_parameter: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_action", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-action"], d, "max_url_parameter_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-action"]) {
			return fmt.Errorf("Error reading max_url_parameter_action: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_block_period", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-block-period"], d, "max_url_parameter_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-block-period"]) {
			return fmt.Errorf("Error reading max_url_parameter_block_period: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-threat-weight"], d, "max_url_parameter_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-threat-weight"]) {
			return fmt.Errorf("Error reading max_url_parameter_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_severity", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-severity"], d, "max_url_parameter_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-severity"]) {
			return fmt.Errorf("Error reading max_url_parameter_severity: %v", err)
		}
	}
	if err = d.Set("max_url_parameter_trigger", flattenWafHttpProtocolParameterRestriction(o["max-url-parameter-trigger"], d, "max_url_parameter_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-parameter-trigger"]) {
			return fmt.Errorf("Error reading max_url_parameter_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_host_name_check", flattenWafHttpProtocolParameterRestriction(o["Illegal-host-name-check"], d, "illegal_host_name_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-host-name-check"]) {
			return fmt.Errorf("Error reading illegal_host_name_check: %v", err)
		}
	}
	if err = d.Set("illegal_host_name_check_action", flattenWafHttpProtocolParameterRestriction(o["Illegal-host-name-check-action"], d, "illegal_host_name_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-host-name-check-action"]) {
			return fmt.Errorf("Error reading illegal_host_name_check_action: %v", err)
		}
	}
	if err = d.Set("illegal_host_name_check_block_period", flattenWafHttpProtocolParameterRestriction(o["Illegal-host-name-check-block-period"], d, "illegal_host_name_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-host-name-check-block-period"]) {
			return fmt.Errorf("Error reading illegal_host_name_check_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_host_name_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Illegal-host-name-check-threat-weight"], d, "illegal_host_name_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-host-name-check-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_host_name_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_host_name_check_severity", flattenWafHttpProtocolParameterRestriction(o["Illegal-host-name-check-severity"], d, "illegal_host_name_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-host-name-check-severity"]) {
			return fmt.Errorf("Error reading illegal_host_name_check_severity: %v", err)
		}
	}
	if err = d.Set("illegal_host_name_check_trigger", flattenWafHttpProtocolParameterRestriction(o["Illegal-host-name-check-trigger"], d, "illegal_host_name_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-host-name-check-trigger"]) {
			return fmt.Errorf("Error reading illegal_host_name_check_trigger: %v", err)
		}
	}
	if err = d.Set("number_of_ranges_in_range_header_check", flattenWafHttpProtocolParameterRestriction(o["number-of-ranges-in-range-header-check"], d, "number_of_ranges_in_range_header_check", sv)); err != nil {
		if !fortiAPIPatch(o["number-of-ranges-in-range-header-check"]) {
			return fmt.Errorf("Error reading number_of_ranges_in_range_header_check: %v", err)
		}
	}
	if err = d.Set("number_of_ranges_in_range_header", flattenWafHttpProtocolParameterRestriction(o["number-of-ranges-in-range-header"], d, "number_of_ranges_in_range_header", sv)); err != nil {
		if !fortiAPIPatch(o["number-of-ranges-in-range-header"]) {
			return fmt.Errorf("Error reading number_of_ranges_in_range_header: %v", err)
		}
	}
	if err = d.Set("number_of_ranges_in_range_header_action", flattenWafHttpProtocolParameterRestriction(o["number-of-ranges-in-range-header-action"], d, "number_of_ranges_in_range_header_action", sv)); err != nil {
		if !fortiAPIPatch(o["number-of-ranges-in-range-header-action"]) {
			return fmt.Errorf("Error reading number_of_ranges_in_range_header_action: %v", err)
		}
	}
	if err = d.Set("number_of_ranges_in_range_header_block_period", flattenWafHttpProtocolParameterRestriction(o["number-of-ranges-in-range-header-block-period"], d, "number_of_ranges_in_range_header_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["number-of-ranges-in-range-header-block-period"]) {
			return fmt.Errorf("Error reading number_of_ranges_in_range_header_block_period: %v", err)
		}
	}
	if err = d.Set("number_of_ranges_in_range_header_threat_weight", flattenWafHttpProtocolParameterRestriction(o["number-of-ranges-in-range-header-threat-weight"], d, "number_of_ranges_in_range_header_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["number-of-ranges-in-range-header-threat-weight"]) {
			return fmt.Errorf("Error reading number_of_ranges_in_range_header_threat_weight: %v", err)
		}
	}
	if err = d.Set("number_of_ranges_in_range_header_severity", flattenWafHttpProtocolParameterRestriction(o["number-of-ranges-in-range-header-severity"], d, "number_of_ranges_in_range_header_severity", sv)); err != nil {
		if !fortiAPIPatch(o["number-of-ranges-in-range-header-severity"]) {
			return fmt.Errorf("Error reading number_of_ranges_in_range_header_severity: %v", err)
		}
	}
	if err = d.Set("number_of_ranges_in_range_header_trigger", flattenWafHttpProtocolParameterRestriction(o["number-of-ranges-in-range-header-trigger"], d, "number_of_ranges_in_range_header_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["number-of-ranges-in-range-header-trigger"]) {
			return fmt.Errorf("Error reading number_of_ranges_in_range_header_trigger: %v", err)
		}
	}
	if err = d.Set("http2_max_requests_check", flattenWafHttpProtocolParameterRestriction(o["http2-max-requests-check"], d, "http2_max_requests_check", sv)); err != nil {
		if !fortiAPIPatch(o["http2-max-requests-check"]) {
			return fmt.Errorf("Error reading http2_max_requests_check: %v", err)
		}
	}
	if err = d.Set("http2_max_requests", flattenWafHttpProtocolParameterRestriction(o["http2-max-requests"], d, "http2_max_requests", sv)); err != nil {
		if !fortiAPIPatch(o["http2-max-requests"]) {
			return fmt.Errorf("Error reading http2_max_requests: %v", err)
		}
	}
	if err = d.Set("http2_max_requests_action", flattenWafHttpProtocolParameterRestriction(o["http2-max-requests-action"], d, "http2_max_requests_action", sv)); err != nil {
		if !fortiAPIPatch(o["http2-max-requests-action"]) {
			return fmt.Errorf("Error reading http2_max_requests_action: %v", err)
		}
	}
	if err = d.Set("http2_max_requests_block_period", flattenWafHttpProtocolParameterRestriction(o["http2-max-requests-block-period"], d, "http2_max_requests_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["http2-max-requests-block-period"]) {
			return fmt.Errorf("Error reading http2_max_requests_block_period: %v", err)
		}
	}
	if err = d.Set("http2_max_requests_severity", flattenWafHttpProtocolParameterRestriction(o["http2-max-requests-severity"], d, "http2_max_requests_severity", sv)); err != nil {
		if !fortiAPIPatch(o["http2-max-requests-severity"]) {
			return fmt.Errorf("Error reading http2_max_requests_severity: %v", err)
		}
	}
	if err = d.Set("http2_max_requests_trigger", flattenWafHttpProtocolParameterRestriction(o["http2-max-requests-trigger"], d, "http2_max_requests_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["http2-max-requests-trigger"]) {
			return fmt.Errorf("Error reading http2_max_requests_trigger: %v", err)
		}
	}
	if err = d.Set("http2_max_requests_threat_weight", flattenWafHttpProtocolParameterRestriction(o["http2-max-requests-threat-weight"], d, "http2_max_requests_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["http2-max-requests-threat-weight"]) {
			return fmt.Errorf("Error reading http2_max_requests_threat_weight: %v", err)
		}
	}
	if err = d.Set("exception_name", flattenWafHttpProtocolParameterRestriction(o["exception_name"], d, "exception_name", sv)); err != nil {
		if !fortiAPIPatch(o["exception_name"]) {
			return fmt.Errorf("Error reading exception_name: %v", err)
		}
	}
	if err = d.Set("block_malformed_request_check", flattenWafHttpProtocolParameterRestriction(o["block-malformed-request-check"], d, "block_malformed_request_check", sv)); err != nil {
		if !fortiAPIPatch(o["block-malformed-request-check"]) {
			return fmt.Errorf("Error reading block_malformed_request_check: %v", err)
		}
	}
	if err = d.Set("block_malformed_request_action", flattenWafHttpProtocolParameterRestriction(o["block-malformed-request-action"], d, "block_malformed_request_action", sv)); err != nil {
		if !fortiAPIPatch(o["block-malformed-request-action"]) {
			return fmt.Errorf("Error reading block_malformed_request_action: %v", err)
		}
	}
	if err = d.Set("block_malformed_request_block_period", flattenWafHttpProtocolParameterRestriction(o["block-malformed-request-block-period"], d, "block_malformed_request_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["block-malformed-request-block-period"]) {
			return fmt.Errorf("Error reading block_malformed_request_block_period: %v", err)
		}
	}
	if err = d.Set("block_malformed_request_threat_weight", flattenWafHttpProtocolParameterRestriction(o["block-malformed-request-threat-weight"], d, "block_malformed_request_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["block-malformed-request-threat-weight"]) {
			return fmt.Errorf("Error reading block_malformed_request_threat_weight: %v", err)
		}
	}
	if err = d.Set("block_malformed_request_severity", flattenWafHttpProtocolParameterRestriction(o["block-malformed-request-severity"], d, "block_malformed_request_severity", sv)); err != nil {
		if !fortiAPIPatch(o["block-malformed-request-severity"]) {
			return fmt.Errorf("Error reading block_malformed_request_severity: %v", err)
		}
	}
	if err = d.Set("block_malformed_request_trigger", flattenWafHttpProtocolParameterRestriction(o["block-malformed-request-trigger"], d, "block_malformed_request_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["block-malformed-request-trigger"]) {
			return fmt.Errorf("Error reading block_malformed_request_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_content_length_check", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-length-check"], d, "illegal_content_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-length-check"]) {
			return fmt.Errorf("Error reading illegal_content_length_check: %v", err)
		}
	}
	if err = d.Set("illegal_content_length_check_action", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-length-check-action"], d, "illegal_content_length_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-length-check-action"]) {
			return fmt.Errorf("Error reading illegal_content_length_check_action: %v", err)
		}
	}
	if err = d.Set("illegal_content_length_check_block_period", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-length-check-block-period"], d, "illegal_content_length_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-length-check-block-period"]) {
			return fmt.Errorf("Error reading illegal_content_length_check_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_content_length_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-length-check-threat-weight"], d, "illegal_content_length_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-length-check-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_content_length_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_content_length_check_severity", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-length-check-severity"], d, "illegal_content_length_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-length-check-severity"]) {
			return fmt.Errorf("Error reading illegal_content_length_check_severity: %v", err)
		}
	}
	if err = d.Set("illegal_content_length_check_trigger", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-length-check-trigger"], d, "illegal_content_length_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-length-check-trigger"]) {
			return fmt.Errorf("Error reading illegal_content_length_check_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_content_type_check", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-type-check"], d, "illegal_content_type_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-type-check"]) {
			return fmt.Errorf("Error reading illegal_content_type_check: %v", err)
		}
	}
	if err = d.Set("illegal_content_type_check_action", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-type-check-action"], d, "illegal_content_type_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-type-check-action"]) {
			return fmt.Errorf("Error reading illegal_content_type_check_action: %v", err)
		}
	}
	if err = d.Set("illegal_content_type_check_block_period", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-type-check-block-period"], d, "illegal_content_type_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-type-check-block-period"]) {
			return fmt.Errorf("Error reading illegal_content_type_check_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_content_type_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-type-check-threat-weight"], d, "illegal_content_type_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-type-check-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_content_type_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_content_type_check_severity", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-type-check-severity"], d, "illegal_content_type_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-type-check-severity"]) {
			return fmt.Errorf("Error reading illegal_content_type_check_severity: %v", err)
		}
	}
	if err = d.Set("illegal_content_type_check_trigger", flattenWafHttpProtocolParameterRestriction(o["Illegal-content-type-check-trigger"], d, "illegal_content_type_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-content-type-check-trigger"]) {
			return fmt.Errorf("Error reading illegal_content_type_check_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_response_code_check", flattenWafHttpProtocolParameterRestriction(o["Illegal-response-code-check"], d, "illegal_response_code_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-response-code-check"]) {
			return fmt.Errorf("Error reading illegal_response_code_check: %v", err)
		}
	}
	if err = d.Set("illegal_response_code_check_action", flattenWafHttpProtocolParameterRestriction(o["Illegal-response-code-check-action"], d, "illegal_response_code_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-response-code-check-action"]) {
			return fmt.Errorf("Error reading illegal_response_code_check_action: %v", err)
		}
	}
	if err = d.Set("illegal_response_code_check_block_period", flattenWafHttpProtocolParameterRestriction(o["Illegal-response-code-check-block-period"], d, "illegal_response_code_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-response-code-check-block-period"]) {
			return fmt.Errorf("Error reading illegal_response_code_check_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_response_code_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Illegal-response-code-check-threat-weight"], d, "illegal_response_code_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-response-code-check-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_response_code_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_response_code_check_severity", flattenWafHttpProtocolParameterRestriction(o["Illegal-response-code-check-severity"], d, "illegal_response_code_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-response-code-check-severity"]) {
			return fmt.Errorf("Error reading illegal_response_code_check_severity: %v", err)
		}
	}
	if err = d.Set("illegal_response_code_check_trigger", flattenWafHttpProtocolParameterRestriction(o["Illegal-response-code-check-trigger"], d, "illegal_response_code_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-response-code-check-trigger"]) {
			return fmt.Errorf("Error reading illegal_response_code_check_trigger: %v", err)
		}
	}
	if err = d.Set("post_request_ctype_check", flattenWafHttpProtocolParameterRestriction(o["Post-request-ctype-check"], d, "post_request_ctype_check", sv)); err != nil {
		if !fortiAPIPatch(o["Post-request-ctype-check"]) {
			return fmt.Errorf("Error reading post_request_ctype_check: %v", err)
		}
	}
	if err = d.Set("post_request_ctype_check_action", flattenWafHttpProtocolParameterRestriction(o["Post-request-ctype-check-action"], d, "post_request_ctype_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["Post-request-ctype-check-action"]) {
			return fmt.Errorf("Error reading post_request_ctype_check_action: %v", err)
		}
	}
	if err = d.Set("post_request_ctype_check_block_period", flattenWafHttpProtocolParameterRestriction(o["Post-request-ctype-check-block-period"], d, "post_request_ctype_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Post-request-ctype-check-block-period"]) {
			return fmt.Errorf("Error reading post_request_ctype_check_block_period: %v", err)
		}
	}
	if err = d.Set("post_request_ctype_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Post-request-ctype-check-threat-weight"], d, "post_request_ctype_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Post-request-ctype-check-threat-weight"]) {
			return fmt.Errorf("Error reading post_request_ctype_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("post_request_ctype_check_severity", flattenWafHttpProtocolParameterRestriction(o["Post-request-ctype-check-severity"], d, "post_request_ctype_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Post-request-ctype-check-severity"]) {
			return fmt.Errorf("Error reading post_request_ctype_check_severity: %v", err)
		}
	}
	if err = d.Set("post_request_ctype_check_trigger", flattenWafHttpProtocolParameterRestriction(o["Post-request-ctype-check-trigger"], d, "post_request_ctype_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Post-request-ctype-check-trigger"]) {
			return fmt.Errorf("Error reading post_request_ctype_check_trigger: %v", err)
		}
	}
	if err = d.Set("max_http_header_name_length_check", flattenWafHttpProtocolParameterRestriction(o["max-http-header-name-length-check"], d, "max_http_header_name_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-name-length-check"]) {
			return fmt.Errorf("Error reading max_http_header_name_length_check: %v", err)
		}
	}
	if err = d.Set("max_http_header_name_length", flattenWafHttpProtocolParameterRestriction(o["max-http-header-name-length"], d, "max_http_header_name_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-name-length"]) {
			return fmt.Errorf("Error reading max_http_header_name_length: %v", err)
		}
	}
	if err = d.Set("max_http_header_name_length_action", flattenWafHttpProtocolParameterRestriction(o["max-http-header-name-length-action"], d, "max_http_header_name_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-name-length-action"]) {
			return fmt.Errorf("Error reading max_http_header_name_length_action: %v", err)
		}
	}
	if err = d.Set("max_http_header_name_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-http-header-name-length-block-period"], d, "max_http_header_name_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-name-length-block-period"]) {
			return fmt.Errorf("Error reading max_http_header_name_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_http_header_name_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-http-header-name-length-threat-weight"], d, "max_http_header_name_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-name-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_http_header_name_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_http_header_name_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-http-header-name-length-severity"], d, "max_http_header_name_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-name-length-severity"]) {
			return fmt.Errorf("Error reading max_http_header_name_length_severity: %v", err)
		}
	}
	if err = d.Set("max_http_header_name_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-http-header-name-length-trigger"], d, "max_http_header_name_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-name-length-trigger"]) {
			return fmt.Errorf("Error reading max_http_header_name_length_trigger: %v", err)
		}
	}
	if err = d.Set("max_http_header_value_length_check", flattenWafHttpProtocolParameterRestriction(o["max-http-header-value-length-check"], d, "max_http_header_value_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-value-length-check"]) {
			return fmt.Errorf("Error reading max_http_header_value_length_check: %v", err)
		}
	}
	if err = d.Set("max_http_header_value_length", flattenWafHttpProtocolParameterRestriction(o["max-http-header-value-length"], d, "max_http_header_value_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-value-length"]) {
			return fmt.Errorf("Error reading max_http_header_value_length: %v", err)
		}
	}
	if err = d.Set("max_http_header_value_length_action", flattenWafHttpProtocolParameterRestriction(o["max-http-header-value-length-action"], d, "max_http_header_value_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-value-length-action"]) {
			return fmt.Errorf("Error reading max_http_header_value_length_action: %v", err)
		}
	}
	if err = d.Set("max_http_header_value_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-http-header-value-length-block-period"], d, "max_http_header_value_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-value-length-block-period"]) {
			return fmt.Errorf("Error reading max_http_header_value_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_http_header_value_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-http-header-value-length-threat-weight"], d, "max_http_header_value_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-value-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_http_header_value_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_http_header_value_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-http-header-value-length-severity"], d, "max_http_header_value_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-value-length-severity"]) {
			return fmt.Errorf("Error reading max_http_header_value_length_severity: %v", err)
		}
	}
	if err = d.Set("max_http_header_value_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-http-header-value-length-trigger"], d, "max_http_header_value_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-header-value-length-trigger"]) {
			return fmt.Errorf("Error reading max_http_header_value_length_trigger: %v", err)
		}
	}
	if err = d.Set("parameter_name_check", flattenWafHttpProtocolParameterRestriction(o["parameter-name-check"], d, "parameter_name_check", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-name-check"]) {
			return fmt.Errorf("Error reading parameter_name_check: %v", err)
		}
	}
	if err = d.Set("parameter_name_check_action", flattenWafHttpProtocolParameterRestriction(o["parameter-name-check-action"], d, "parameter_name_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-name-check-action"]) {
			return fmt.Errorf("Error reading parameter_name_check_action: %v", err)
		}
	}
	if err = d.Set("parameter_name_check_block_period", flattenWafHttpProtocolParameterRestriction(o["parameter-name-check-block-period"], d, "parameter_name_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-name-check-block-period"]) {
			return fmt.Errorf("Error reading parameter_name_check_block_period: %v", err)
		}
	}
	if err = d.Set("parameter_name_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["parameter-name-check-threat-weight"], d, "parameter_name_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-name-check-threat-weight"]) {
			return fmt.Errorf("Error reading parameter_name_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("parameter_name_check_severity", flattenWafHttpProtocolParameterRestriction(o["parameter-name-check-severity"], d, "parameter_name_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-name-check-severity"]) {
			return fmt.Errorf("Error reading parameter_name_check_severity: %v", err)
		}
	}
	if err = d.Set("parameter_name_check_trigger", flattenWafHttpProtocolParameterRestriction(o["parameter-name-check-trigger"], d, "parameter_name_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-name-check-trigger"]) {
			return fmt.Errorf("Error reading parameter_name_check_trigger: %v", err)
		}
	}
	if err = d.Set("parameter_value_check", flattenWafHttpProtocolParameterRestriction(o["parameter-value-check"], d, "parameter_value_check", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-value-check"]) {
			return fmt.Errorf("Error reading parameter_value_check: %v", err)
		}
	}
	if err = d.Set("parameter_value_check_action", flattenWafHttpProtocolParameterRestriction(o["parameter-value-check-action"], d, "parameter_value_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-value-check-action"]) {
			return fmt.Errorf("Error reading parameter_value_check_action: %v", err)
		}
	}
	if err = d.Set("parameter_value_check_block_period", flattenWafHttpProtocolParameterRestriction(o["parameter-value-check-block-period"], d, "parameter_value_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-value-check-block-period"]) {
			return fmt.Errorf("Error reading parameter_value_check_block_period: %v", err)
		}
	}
	if err = d.Set("parameter_value_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["parameter-value-check-threat-weight"], d, "parameter_value_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-value-check-threat-weight"]) {
			return fmt.Errorf("Error reading parameter_value_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("parameter_value_check_severity", flattenWafHttpProtocolParameterRestriction(o["parameter-value-check-severity"], d, "parameter_value_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-value-check-severity"]) {
			return fmt.Errorf("Error reading parameter_value_check_severity: %v", err)
		}
	}
	if err = d.Set("parameter_value_check_trigger", flattenWafHttpProtocolParameterRestriction(o["parameter-value-check-trigger"], d, "parameter_value_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["parameter-value-check-trigger"]) {
			return fmt.Errorf("Error reading parameter_value_check_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_header_name_check", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-name-check"], d, "illegal_header_name_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-name-check"]) {
			return fmt.Errorf("Error reading illegal_header_name_check: %v", err)
		}
	}
	if err = d.Set("illegal_header_name_check_action", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-name-check-action"], d, "illegal_header_name_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-name-check-action"]) {
			return fmt.Errorf("Error reading illegal_header_name_check_action: %v", err)
		}
	}
	if err = d.Set("illegal_header_name_check_block_period", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-name-check-block-period"], d, "illegal_header_name_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-name-check-block-period"]) {
			return fmt.Errorf("Error reading illegal_header_name_check_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_header_name_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-name-check-threat-weight"], d, "illegal_header_name_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-name-check-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_header_name_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_header_name_check_severity", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-name-check-severity"], d, "illegal_header_name_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-name-check-severity"]) {
			return fmt.Errorf("Error reading illegal_header_name_check_severity: %v", err)
		}
	}
	if err = d.Set("illegal_header_name_check_trigger", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-name-check-trigger"], d, "illegal_header_name_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-name-check-trigger"]) {
			return fmt.Errorf("Error reading illegal_header_name_check_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_header_value_check", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-value-check"], d, "illegal_header_value_check", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-value-check"]) {
			return fmt.Errorf("Error reading illegal_header_value_check: %v", err)
		}
	}
	if err = d.Set("illegal_header_value_check_action", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-value-check-action"], d, "illegal_header_value_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-value-check-action"]) {
			return fmt.Errorf("Error reading illegal_header_value_check_action: %v", err)
		}
	}
	if err = d.Set("illegal_header_value_check_block_period", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-value-check-block-period"], d, "illegal_header_value_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-value-check-block-period"]) {
			return fmt.Errorf("Error reading illegal_header_value_check_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_header_value_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-value-check-threat-weight"], d, "illegal_header_value_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-value-check-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_header_value_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_header_value_check_severity", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-value-check-severity"], d, "illegal_header_value_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-value-check-severity"]) {
			return fmt.Errorf("Error reading illegal_header_value_check_severity: %v", err)
		}
	}
	if err = d.Set("illegal_header_value_check_trigger", flattenWafHttpProtocolParameterRestriction(o["Illegal-header-value-check-trigger"], d, "illegal_header_value_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Illegal-header-value-check-trigger"]) {
			return fmt.Errorf("Error reading illegal_header_value_check_trigger: %v", err)
		}
	}
	if err = d.Set("max_http_body_parameter_length_check", flattenWafHttpProtocolParameterRestriction(o["max-http-body-parameter-length-check"], d, "max_http_body_parameter_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-parameter-length-check"]) {
			return fmt.Errorf("Error reading max_http_body_parameter_length_check: %v", err)
		}
	}
	if err = d.Set("max_http_body_parameter_length", flattenWafHttpProtocolParameterRestriction(o["max-http-body-parameter-length"], d, "max_http_body_parameter_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-parameter-length"]) {
			return fmt.Errorf("Error reading max_http_body_parameter_length: %v", err)
		}
	}
	if err = d.Set("max_http_body_parameter_length_action", flattenWafHttpProtocolParameterRestriction(o["max-http-body-parameter-length-action"], d, "max_http_body_parameter_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-parameter-length-action"]) {
			return fmt.Errorf("Error reading max_http_body_parameter_length_action: %v", err)
		}
	}
	if err = d.Set("max_http_body_parameter_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-http-body-parameter-length-block-period"], d, "max_http_body_parameter_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-parameter-length-block-period"]) {
			return fmt.Errorf("Error reading max_http_body_parameter_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_http_body_parameter_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-http-body-parameter-length-threat-weight"], d, "max_http_body_parameter_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-parameter-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_http_body_parameter_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_http_body_parameter_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-http-body-parameter-length-severity"], d, "max_http_body_parameter_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-parameter-length-severity"]) {
			return fmt.Errorf("Error reading max_http_body_parameter_length_severity: %v", err)
		}
	}
	if err = d.Set("max_http_body_parameter_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-http-body-parameter-length-trigger"], d, "max_http_body_parameter_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-body-parameter-length-trigger"]) {
			return fmt.Errorf("Error reading max_http_body_parameter_length_trigger: %v", err)
		}
	}
	if err = d.Set("max_http_request_filename_length_check", flattenWafHttpProtocolParameterRestriction(o["max-http-request-filename-length-check"], d, "max_http_request_filename_length_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-filename-length-check"]) {
			return fmt.Errorf("Error reading max_http_request_filename_length_check: %v", err)
		}
	}
	if err = d.Set("max_http_request_filename_length", flattenWafHttpProtocolParameterRestriction(o["max-http-request-filename-length"], d, "max_http_request_filename_length", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-filename-length"]) {
			return fmt.Errorf("Error reading max_http_request_filename_length: %v", err)
		}
	}
	if err = d.Set("max_http_request_filename_length_action", flattenWafHttpProtocolParameterRestriction(o["max-http-request-filename-length-action"], d, "max_http_request_filename_length_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-filename-length-action"]) {
			return fmt.Errorf("Error reading max_http_request_filename_length_action: %v", err)
		}
	}
	if err = d.Set("max_http_request_filename_length_block_period", flattenWafHttpProtocolParameterRestriction(o["max-http-request-filename-length-block-period"], d, "max_http_request_filename_length_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-filename-length-block-period"]) {
			return fmt.Errorf("Error reading max_http_request_filename_length_block_period: %v", err)
		}
	}
	if err = d.Set("max_http_request_filename_length_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-http-request-filename-length-threat-weight"], d, "max_http_request_filename_length_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-filename-length-threat-weight"]) {
			return fmt.Errorf("Error reading max_http_request_filename_length_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_http_request_filename_length_severity", flattenWafHttpProtocolParameterRestriction(o["max-http-request-filename-length-severity"], d, "max_http_request_filename_length_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-filename-length-severity"]) {
			return fmt.Errorf("Error reading max_http_request_filename_length_severity: %v", err)
		}
	}
	if err = d.Set("max_http_request_filename_length_trigger", flattenWafHttpProtocolParameterRestriction(o["max-http-request-filename-length-trigger"], d, "max_http_request_filename_length_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-http-request-filename-length-trigger"]) {
			return fmt.Errorf("Error reading max_http_request_filename_length_trigger: %v", err)
		}
	}
	if err = d.Set("web_socket_protocol_check", flattenWafHttpProtocolParameterRestriction(o["web-socket-protocol-check"], d, "web_socket_protocol_check", sv)); err != nil {
		if !fortiAPIPatch(o["web-socket-protocol-check"]) {
			return fmt.Errorf("Error reading web_socket_protocol_check: %v", err)
		}
	}
	if err = d.Set("web_socket_protocol_action", flattenWafHttpProtocolParameterRestriction(o["web-socket-protocol-action"], d, "web_socket_protocol_action", sv)); err != nil {
		if !fortiAPIPatch(o["web-socket-protocol-action"]) {
			return fmt.Errorf("Error reading web_socket_protocol_action: %v", err)
		}
	}
	if err = d.Set("web_socket_protocol_block_period", flattenWafHttpProtocolParameterRestriction(o["web-socket-protocol-block-period"], d, "web_socket_protocol_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["web-socket-protocol-block-period"]) {
			return fmt.Errorf("Error reading web_socket_protocol_block_period: %v", err)
		}
	}
	if err = d.Set("web_socket_protocol_severity", flattenWafHttpProtocolParameterRestriction(o["web-socket-protocol-severity"], d, "web_socket_protocol_severity", sv)); err != nil {
		if !fortiAPIPatch(o["web-socket-protocol-severity"]) {
			return fmt.Errorf("Error reading web_socket_protocol_severity: %v", err)
		}
	}
	if err = d.Set("web_socket_protocol_trigger", flattenWafHttpProtocolParameterRestriction(o["web-socket-protocol-trigger"], d, "web_socket_protocol_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["web-socket-protocol-trigger"]) {
			return fmt.Errorf("Error reading web_socket_protocol_trigger: %v", err)
		}
	}
	if err = d.Set("web_socket_protocol_threat_weight", flattenWafHttpProtocolParameterRestriction(o["web-socket-protocol-threat-weight"], d, "web_socket_protocol_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["web-socket-protocol-threat-weight"]) {
			return fmt.Errorf("Error reading web_socket_protocol_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_setting_header_table_size_check", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-table-size-check"], d, "max_setting_header_table_size_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-table-size-check"]) {
			return fmt.Errorf("Error reading max_setting_header_table_size_check: %v", err)
		}
	}
	if err = d.Set("max_setting_header_table_size", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-table-size"], d, "max_setting_header_table_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-table-size"]) {
			return fmt.Errorf("Error reading max_setting_header_table_size: %v", err)
		}
	}
	if err = d.Set("max_setting_header_table_size_action", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-table-size-action"], d, "max_setting_header_table_size_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-table-size-action"]) {
			return fmt.Errorf("Error reading max_setting_header_table_size_action: %v", err)
		}
	}
	if err = d.Set("max_setting_header_table_size_block_period", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-table-size-block-period"], d, "max_setting_header_table_size_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-table-size-block-period"]) {
			return fmt.Errorf("Error reading max_setting_header_table_size_block_period: %v", err)
		}
	}
	if err = d.Set("max_setting_header_table_size_severity", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-table-size-severity"], d, "max_setting_header_table_size_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-table-size-severity"]) {
			return fmt.Errorf("Error reading max_setting_header_table_size_severity: %v", err)
		}
	}
	if err = d.Set("max_setting_header_table_size_trigger", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-table-size-trigger"], d, "max_setting_header_table_size_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-table-size-trigger"]) {
			return fmt.Errorf("Error reading max_setting_header_table_size_trigger: %v", err)
		}
	}
	if err = d.Set("max_setting_header_table_size_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-table-size-threat-weight"], d, "max_setting_header_table_size_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-table-size-threat-weight"]) {
			return fmt.Errorf("Error reading max_setting_header_table_size_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_setting_current_streams_num_check", flattenWafHttpProtocolParameterRestriction(o["max-setting-current-streams-num-check"], d, "max_setting_current_streams_num_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-current-streams-num-check"]) {
			return fmt.Errorf("Error reading max_setting_current_streams_num_check: %v", err)
		}
	}
	if err = d.Set("max_setting_current_streams_num", flattenWafHttpProtocolParameterRestriction(o["max-setting-current-streams-num"], d, "max_setting_current_streams_num", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-current-streams-num"]) {
			return fmt.Errorf("Error reading max_setting_current_streams_num: %v", err)
		}
	}
	if err = d.Set("max_setting_current_streams_num_action", flattenWafHttpProtocolParameterRestriction(o["max-setting-current-streams-num-action"], d, "max_setting_current_streams_num_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-current-streams-num-action"]) {
			return fmt.Errorf("Error reading max_setting_current_streams_num_action: %v", err)
		}
	}
	if err = d.Set("max_setting_current_streams_num_block_period", flattenWafHttpProtocolParameterRestriction(o["max-setting-current-streams-num-block-period"], d, "max_setting_current_streams_num_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-current-streams-num-block-period"]) {
			return fmt.Errorf("Error reading max_setting_current_streams_num_block_period: %v", err)
		}
	}
	if err = d.Set("max_setting_current_streams_num_severity", flattenWafHttpProtocolParameterRestriction(o["max-setting-current-streams-num-severity"], d, "max_setting_current_streams_num_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-current-streams-num-severity"]) {
			return fmt.Errorf("Error reading max_setting_current_streams_num_severity: %v", err)
		}
	}
	if err = d.Set("max_setting_current_streams_num_trigger", flattenWafHttpProtocolParameterRestriction(o["max-setting-current-streams-num-trigger"], d, "max_setting_current_streams_num_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-current-streams-num-trigger"]) {
			return fmt.Errorf("Error reading max_setting_current_streams_num_trigger: %v", err)
		}
	}
	if err = d.Set("max_setting_current_streams_num_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-setting-current-streams-num-threat-weight"], d, "max_setting_current_streams_num_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-current-streams-num-threat-weight"]) {
			return fmt.Errorf("Error reading max_setting_current_streams_num_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_setting_initial_window_size_check", flattenWafHttpProtocolParameterRestriction(o["max-setting-initial-window-size-check"], d, "max_setting_initial_window_size_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-initial-window-size-check"]) {
			return fmt.Errorf("Error reading max_setting_initial_window_size_check: %v", err)
		}
	}
	if err = d.Set("max_setting_initial_window_size", flattenWafHttpProtocolParameterRestriction(o["max-setting-initial-window-size"], d, "max_setting_initial_window_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-initial-window-size"]) {
			return fmt.Errorf("Error reading max_setting_initial_window_size: %v", err)
		}
	}
	if err = d.Set("max_setting_initial_window_size_action", flattenWafHttpProtocolParameterRestriction(o["max-setting-initial-window-size-action"], d, "max_setting_initial_window_size_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-initial-window-size-action"]) {
			return fmt.Errorf("Error reading max_setting_initial_window_size_action: %v", err)
		}
	}
	if err = d.Set("max_setting_initial_window_size_block_period", flattenWafHttpProtocolParameterRestriction(o["max-setting-initial-window-size-block-period"], d, "max_setting_initial_window_size_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-initial-window-size-block-period"]) {
			return fmt.Errorf("Error reading max_setting_initial_window_size_block_period: %v", err)
		}
	}
	if err = d.Set("max_setting_initial_window_size_severity", flattenWafHttpProtocolParameterRestriction(o["max-setting-initial-window-size-severity"], d, "max_setting_initial_window_size_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-initial-window-size-severity"]) {
			return fmt.Errorf("Error reading max_setting_initial_window_size_severity: %v", err)
		}
	}
	if err = d.Set("max_setting_initial_window_size_trigger", flattenWafHttpProtocolParameterRestriction(o["max-setting-initial-window-size-trigger"], d, "max_setting_initial_window_size_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-initial-window-size-trigger"]) {
			return fmt.Errorf("Error reading max_setting_initial_window_size_trigger: %v", err)
		}
	}
	if err = d.Set("max_setting_initial_window_size_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-setting-initial-window-size-threat-weight"], d, "max_setting_initial_window_size_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-initial-window-size-threat-weight"]) {
			return fmt.Errorf("Error reading max_setting_initial_window_size_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_setting_frame_size_check", flattenWafHttpProtocolParameterRestriction(o["max-setting-frame-size-check"], d, "max_setting_frame_size_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-frame-size-check"]) {
			return fmt.Errorf("Error reading max_setting_frame_size_check: %v", err)
		}
	}
	if err = d.Set("max_setting_frame_size", flattenWafHttpProtocolParameterRestriction(o["max-setting-frame-size"], d, "max_setting_frame_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-frame-size"]) {
			return fmt.Errorf("Error reading max_setting_frame_size: %v", err)
		}
	}
	if err = d.Set("max_setting_frame_size_action", flattenWafHttpProtocolParameterRestriction(o["max-setting-frame-size-action"], d, "max_setting_frame_size_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-frame-size-action"]) {
			return fmt.Errorf("Error reading max_setting_frame_size_action: %v", err)
		}
	}
	if err = d.Set("max_setting_frame_size_block_period", flattenWafHttpProtocolParameterRestriction(o["max-setting-frame-size-block-period"], d, "max_setting_frame_size_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-frame-size-block-period"]) {
			return fmt.Errorf("Error reading max_setting_frame_size_block_period: %v", err)
		}
	}
	if err = d.Set("max_setting_frame_size_severity", flattenWafHttpProtocolParameterRestriction(o["max-setting-frame-size-severity"], d, "max_setting_frame_size_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-frame-size-severity"]) {
			return fmt.Errorf("Error reading max_setting_frame_size_severity: %v", err)
		}
	}
	if err = d.Set("max_setting_frame_size_trigger", flattenWafHttpProtocolParameterRestriction(o["max-setting-frame-size-trigger"], d, "max_setting_frame_size_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-frame-size-trigger"]) {
			return fmt.Errorf("Error reading max_setting_frame_size_trigger: %v", err)
		}
	}
	if err = d.Set("max_setting_frame_size_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-setting-frame-size-threat-weight"], d, "max_setting_frame_size_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-frame-size-threat-weight"]) {
			return fmt.Errorf("Error reading max_setting_frame_size_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_setting_header_list_size_check", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-list-size-check"], d, "max_setting_header_list_size_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-list-size-check"]) {
			return fmt.Errorf("Error reading max_setting_header_list_size_check: %v", err)
		}
	}
	if err = d.Set("max_setting_header_list_size", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-list-size"], d, "max_setting_header_list_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-list-size"]) {
			return fmt.Errorf("Error reading max_setting_header_list_size: %v", err)
		}
	}
	if err = d.Set("max_setting_header_list_size_action", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-list-size-action"], d, "max_setting_header_list_size_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-list-size-action"]) {
			return fmt.Errorf("Error reading max_setting_header_list_size_action: %v", err)
		}
	}
	if err = d.Set("max_setting_header_list_size_block_period", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-list-size-block-period"], d, "max_setting_header_list_size_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-list-size-block-period"]) {
			return fmt.Errorf("Error reading max_setting_header_list_size_block_period: %v", err)
		}
	}
	if err = d.Set("max_setting_header_list_size_severity", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-list-size-severity"], d, "max_setting_header_list_size_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-list-size-severity"]) {
			return fmt.Errorf("Error reading max_setting_header_list_size_severity: %v", err)
		}
	}
	if err = d.Set("max_setting_header_list_size_trigger", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-list-size-trigger"], d, "max_setting_header_list_size_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-list-size-trigger"]) {
			return fmt.Errorf("Error reading max_setting_header_list_size_trigger: %v", err)
		}
	}
	if err = d.Set("max_setting_header_list_size_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-setting-header-list-size-threat-weight"], d, "max_setting_header_list_size_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-setting-header-list-size-threat-weight"]) {
			return fmt.Errorf("Error reading max_setting_header_list_size_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_url_param_name_len_check", flattenWafHttpProtocolParameterRestriction(o["max-url-param-name-len-check"], d, "max_url_param_name_len_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-name-len-check"]) {
			return fmt.Errorf("Error reading max_url_param_name_len_check: %v", err)
		}
	}
	if err = d.Set("max_url_param_name_len", flattenWafHttpProtocolParameterRestriction(o["max-url-param-name-len"], d, "max_url_param_name_len", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-name-len"]) {
			return fmt.Errorf("Error reading max_url_param_name_len: %v", err)
		}
	}
	if err = d.Set("max_url_param_name_len_action", flattenWafHttpProtocolParameterRestriction(o["max-url-param-name-len-action"], d, "max_url_param_name_len_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-name-len-action"]) {
			return fmt.Errorf("Error reading max_url_param_name_len_action: %v", err)
		}
	}
	if err = d.Set("max_url_param_name_len_block_period", flattenWafHttpProtocolParameterRestriction(o["max-url-param-name-len-block-period"], d, "max_url_param_name_len_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-name-len-block-period"]) {
			return fmt.Errorf("Error reading max_url_param_name_len_block_period: %v", err)
		}
	}
	if err = d.Set("max_url_param_name_len_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-url-param-name-len-threat-weight"], d, "max_url_param_name_len_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-name-len-threat-weight"]) {
			return fmt.Errorf("Error reading max_url_param_name_len_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_url_param_name_len_severity", flattenWafHttpProtocolParameterRestriction(o["max-url-param-name-len-severity"], d, "max_url_param_name_len_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-name-len-severity"]) {
			return fmt.Errorf("Error reading max_url_param_name_len_severity: %v", err)
		}
	}
	if err = d.Set("max_url_param_name_len_trigger", flattenWafHttpProtocolParameterRestriction(o["max-url-param-name-len-trigger"], d, "max_url_param_name_len_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-name-len-trigger"]) {
			return fmt.Errorf("Error reading max_url_param_name_len_trigger: %v", err)
		}
	}
	if err = d.Set("max_url_param_value_len_check", flattenWafHttpProtocolParameterRestriction(o["max-url-param-value-len-check"], d, "max_url_param_value_len_check", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-value-len-check"]) {
			return fmt.Errorf("Error reading max_url_param_value_len_check: %v", err)
		}
	}
	if err = d.Set("max_url_param_value_len", flattenWafHttpProtocolParameterRestriction(o["max-url-param-value-len"], d, "max_url_param_value_len", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-value-len"]) {
			return fmt.Errorf("Error reading max_url_param_value_len: %v", err)
		}
	}
	if err = d.Set("max_url_param_value_len_action", flattenWafHttpProtocolParameterRestriction(o["max-url-param-value-len-action"], d, "max_url_param_value_len_action", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-value-len-action"]) {
			return fmt.Errorf("Error reading max_url_param_value_len_action: %v", err)
		}
	}
	if err = d.Set("max_url_param_value_len_block_period", flattenWafHttpProtocolParameterRestriction(o["max-url-param-value-len-block-period"], d, "max_url_param_value_len_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-value-len-block-period"]) {
			return fmt.Errorf("Error reading max_url_param_value_len_block_period: %v", err)
		}
	}
	if err = d.Set("max_url_param_value_len_threat_weight", flattenWafHttpProtocolParameterRestriction(o["max-url-param-value-len-threat-weight"], d, "max_url_param_value_len_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-value-len-threat-weight"]) {
			return fmt.Errorf("Error reading max_url_param_value_len_threat_weight: %v", err)
		}
	}
	if err = d.Set("max_url_param_value_len_severity", flattenWafHttpProtocolParameterRestriction(o["max-url-param-value-len-severity"], d, "max_url_param_value_len_severity", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-value-len-severity"]) {
			return fmt.Errorf("Error reading max_url_param_value_len_severity: %v", err)
		}
	}
	if err = d.Set("max_url_param_value_len_trigger", flattenWafHttpProtocolParameterRestriction(o["max-url-param-value-len-trigger"], d, "max_url_param_value_len_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["max-url-param-value-len-trigger"]) {
			return fmt.Errorf("Error reading max_url_param_value_len_trigger: %v", err)
		}
	}
	if err = d.Set("url_param_name_check", flattenWafHttpProtocolParameterRestriction(o["url-param-name-check"], d, "url_param_name_check", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-name-check"]) {
			return fmt.Errorf("Error reading url_param_name_check: %v", err)
		}
	}
	if err = d.Set("url_param_name_check_action", flattenWafHttpProtocolParameterRestriction(o["url-param-name-check-action"], d, "url_param_name_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-name-check-action"]) {
			return fmt.Errorf("Error reading url_param_name_check_action: %v", err)
		}
	}
	if err = d.Set("url_param_name_check_block_period", flattenWafHttpProtocolParameterRestriction(o["url-param-name-check-block-period"], d, "url_param_name_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-name-check-block-period"]) {
			return fmt.Errorf("Error reading url_param_name_check_block_period: %v", err)
		}
	}
	if err = d.Set("url_param_name_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["url-param-name-check-threat-weight"], d, "url_param_name_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-name-check-threat-weight"]) {
			return fmt.Errorf("Error reading url_param_name_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("url_param_name_check_severity", flattenWafHttpProtocolParameterRestriction(o["url-param-name-check-severity"], d, "url_param_name_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-name-check-severity"]) {
			return fmt.Errorf("Error reading url_param_name_check_severity: %v", err)
		}
	}
	if err = d.Set("url_param_name_check_trigger", flattenWafHttpProtocolParameterRestriction(o["url-param-name-check-trigger"], d, "url_param_name_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-name-check-trigger"]) {
			return fmt.Errorf("Error reading url_param_name_check_trigger: %v", err)
		}
	}
	if err = d.Set("url_param_value_check", flattenWafHttpProtocolParameterRestriction(o["url-param-value-check"], d, "url_param_value_check", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-value-check"]) {
			return fmt.Errorf("Error reading url_param_value_check: %v", err)
		}
	}
	if err = d.Set("url_param_value_check_action", flattenWafHttpProtocolParameterRestriction(o["url-param-value-check-action"], d, "url_param_value_check_action", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-value-check-action"]) {
			return fmt.Errorf("Error reading url_param_value_check_action: %v", err)
		}
	}
	if err = d.Set("url_param_value_check_block_period", flattenWafHttpProtocolParameterRestriction(o["url-param-value-check-block-period"], d, "url_param_value_check_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-value-check-block-period"]) {
			return fmt.Errorf("Error reading url_param_value_check_block_period: %v", err)
		}
	}
	if err = d.Set("url_param_value_check_threat_weight", flattenWafHttpProtocolParameterRestriction(o["url-param-value-check-threat-weight"], d, "url_param_value_check_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-value-check-threat-weight"]) {
			return fmt.Errorf("Error reading url_param_value_check_threat_weight: %v", err)
		}
	}
	if err = d.Set("url_param_value_check_severity", flattenWafHttpProtocolParameterRestriction(o["url-param-value-check-severity"], d, "url_param_value_check_severity", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-value-check-severity"]) {
			return fmt.Errorf("Error reading url_param_value_check_severity: %v", err)
		}
	}
	if err = d.Set("url_param_value_check_trigger", flattenWafHttpProtocolParameterRestriction(o["url-param-value-check-trigger"], d, "url_param_value_check_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["url-param-value-check-trigger"]) {
			return fmt.Errorf("Error reading url_param_value_check_trigger: %v", err)
		}
	}
	if err = d.Set("null_byte_in_url_check", flattenWafHttpProtocolParameterRestriction(o["null-byte-in-url-check"], d, "null_byte_in_url_check", sv)); err != nil {
		if !fortiAPIPatch(o["null-byte-in-url-check"]) {
			return fmt.Errorf("Error reading null_byte_in_url_check: %v", err)
		}
	}
	if err = d.Set("null_byte_in_url_action", flattenWafHttpProtocolParameterRestriction(o["null-byte-in-url-action"], d, "null_byte_in_url_action", sv)); err != nil {
		if !fortiAPIPatch(o["null-byte-in-url-action"]) {
			return fmt.Errorf("Error reading null_byte_in_url_action: %v", err)
		}
	}
	if err = d.Set("null_byte_in_url_block_period", flattenWafHttpProtocolParameterRestriction(o["null-byte-in-url-block-period"], d, "null_byte_in_url_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["null-byte-in-url-block-period"]) {
			return fmt.Errorf("Error reading null_byte_in_url_block_period: %v", err)
		}
	}
	if err = d.Set("null_byte_in_url_threat_weight", flattenWafHttpProtocolParameterRestriction(o["null-byte-in-url-threat-weight"], d, "null_byte_in_url_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["null-byte-in-url-threat-weight"]) {
			return fmt.Errorf("Error reading null_byte_in_url_threat_weight: %v", err)
		}
	}
	if err = d.Set("null_byte_in_url_severity", flattenWafHttpProtocolParameterRestriction(o["null-byte-in-url-severity"], d, "null_byte_in_url_severity", sv)); err != nil {
		if !fortiAPIPatch(o["null-byte-in-url-severity"]) {
			return fmt.Errorf("Error reading null_byte_in_url_severity: %v", err)
		}
	}
	if err = d.Set("null_byte_in_url_trigger", flattenWafHttpProtocolParameterRestriction(o["null-byte-in-url-trigger"], d, "null_byte_in_url_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["null-byte-in-url-trigger"]) {
			return fmt.Errorf("Error reading null_byte_in_url_trigger: %v", err)
		}
	}
	if err = d.Set("illegal_byte_in_url_check", flattenWafHttpProtocolParameterRestriction(o["illegal-byte-in-url-check"], d, "illegal_byte_in_url_check", sv)); err != nil {
		if !fortiAPIPatch(o["illegal-byte-in-url-check"]) {
			return fmt.Errorf("Error reading illegal_byte_in_url_check: %v", err)
		}
	}
	if err = d.Set("illegal_byte_in_url_action", flattenWafHttpProtocolParameterRestriction(o["illegal-byte-in-url-action"], d, "illegal_byte_in_url_action", sv)); err != nil {
		if !fortiAPIPatch(o["illegal-byte-in-url-action"]) {
			return fmt.Errorf("Error reading illegal_byte_in_url_action: %v", err)
		}
	}
	if err = d.Set("illegal_byte_in_url_block_period", flattenWafHttpProtocolParameterRestriction(o["illegal-byte-in-url-block-period"], d, "illegal_byte_in_url_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["illegal-byte-in-url-block-period"]) {
			return fmt.Errorf("Error reading illegal_byte_in_url_block_period: %v", err)
		}
	}
	if err = d.Set("illegal_byte_in_url_threat_weight", flattenWafHttpProtocolParameterRestriction(o["illegal-byte-in-url-threat-weight"], d, "illegal_byte_in_url_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["illegal-byte-in-url-threat-weight"]) {
			return fmt.Errorf("Error reading illegal_byte_in_url_threat_weight: %v", err)
		}
	}
	if err = d.Set("illegal_byte_in_url_severity", flattenWafHttpProtocolParameterRestriction(o["illegal-byte-in-url-severity"], d, "illegal_byte_in_url_severity", sv)); err != nil {
		if !fortiAPIPatch(o["illegal-byte-in-url-severity"]) {
			return fmt.Errorf("Error reading illegal_byte_in_url_severity: %v", err)
		}
	}
	if err = d.Set("illegal_byte_in_url_trigger", flattenWafHttpProtocolParameterRestriction(o["illegal-byte-in-url-trigger"], d, "illegal_byte_in_url_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["illegal-byte-in-url-trigger"]) {
			return fmt.Errorf("Error reading illegal_byte_in_url_trigger: %v", err)
		}
	}
	if err = d.Set("malformed_url_check", flattenWafHttpProtocolParameterRestriction(o["malformed-url-check"], d, "malformed_url_check", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-url-check"]) {
			return fmt.Errorf("Error reading malformed_url_check: %v", err)
		}
	}
	if err = d.Set("malformed_url_action", flattenWafHttpProtocolParameterRestriction(o["malformed-url-action"], d, "malformed_url_action", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-url-action"]) {
			return fmt.Errorf("Error reading malformed_url_action: %v", err)
		}
	}
	if err = d.Set("malformed_url_block_period", flattenWafHttpProtocolParameterRestriction(o["malformed-url-block-period"], d, "malformed_url_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-url-block-period"]) {
			return fmt.Errorf("Error reading malformed_url_block_period: %v", err)
		}
	}
	if err = d.Set("malformed_url_severity", flattenWafHttpProtocolParameterRestriction(o["malformed-url-severity"], d, "malformed_url_severity", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-url-severity"]) {
			return fmt.Errorf("Error reading malformed_url_severity: %v", err)
		}
	}
	if err = d.Set("malformed_url_trigger", flattenWafHttpProtocolParameterRestriction(o["malformed-url-trigger"], d, "malformed_url_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-url-trigger"]) {
			return fmt.Errorf("Error reading malformed_url_trigger: %v", err)
		}
	}
	if err = d.Set("malformed_url_weight", flattenWafHttpProtocolParameterRestriction(o["malformed-url-weight"], d, "malformed_url_weight", sv)); err != nil {
		if !fortiAPIPatch(o["malformed-url-weight"]) {
			return fmt.Errorf("Error reading malformed_url_weight: %v", err)
		}
	}
	if err = d.Set("redundant_header_check", flattenWafHttpProtocolParameterRestriction(o["redundant-header-check"], d, "redundant_header_check", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-header-check"]) {
			return fmt.Errorf("Error reading redundant_header_check: %v", err)
		}
	}
	if err = d.Set("redundant_header_action", flattenWafHttpProtocolParameterRestriction(o["redundant-header-action"], d, "redundant_header_action", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-header-action"]) {
			return fmt.Errorf("Error reading redundant_header_action: %v", err)
		}
	}
	if err = d.Set("redundant_header_block_period", flattenWafHttpProtocolParameterRestriction(o["redundant-header-block-period"], d, "redundant_header_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-header-block-period"]) {
			return fmt.Errorf("Error reading redundant_header_block_period: %v", err)
		}
	}
	if err = d.Set("redundant_header_threat_weight", flattenWafHttpProtocolParameterRestriction(o["redundant-header-threat-weight"], d, "redundant_header_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-header-threat-weight"]) {
			return fmt.Errorf("Error reading redundant_header_threat_weight: %v", err)
		}
	}
	if err = d.Set("redundant_header_severity", flattenWafHttpProtocolParameterRestriction(o["redundant-header-severity"], d, "redundant_header_severity", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-header-severity"]) {
			return fmt.Errorf("Error reading redundant_header_severity: %v", err)
		}
	}
	if err = d.Set("redundant_header_trigger", flattenWafHttpProtocolParameterRestriction(o["redundant-header-trigger"], d, "redundant_header_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-header-trigger"]) {
			return fmt.Errorf("Error reading redundant_header_trigger: %v", err)
		}
	}
	if err = d.Set("chunk_size_check", flattenWafHttpProtocolParameterRestriction(o["chunk-size-check"], d, "chunk_size_check", sv)); err != nil {
		if !fortiAPIPatch(o["chunk-size-check"]) {
			return fmt.Errorf("Error reading chunk_size_check: %v", err)
		}
	}
	if err = d.Set("chunk_size_action", flattenWafHttpProtocolParameterRestriction(o["chunk-size-action"], d, "chunk_size_action", sv)); err != nil {
		if !fortiAPIPatch(o["chunk-size-action"]) {
			return fmt.Errorf("Error reading chunk_size_action: %v", err)
		}
	}
	if err = d.Set("chunk_size_block_period", flattenWafHttpProtocolParameterRestriction(o["chunk-size-block-period"], d, "chunk_size_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["chunk-size-block-period"]) {
			return fmt.Errorf("Error reading chunk_size_block_period: %v", err)
		}
	}
	if err = d.Set("chunk_size_severity", flattenWafHttpProtocolParameterRestriction(o["chunk-size-severity"], d, "chunk_size_severity", sv)); err != nil {
		if !fortiAPIPatch(o["chunk-size-severity"]) {
			return fmt.Errorf("Error reading chunk_size_severity: %v", err)
		}
	}
	if err = d.Set("chunk_size_trigger", flattenWafHttpProtocolParameterRestriction(o["chunk-size-trigger"], d, "chunk_size_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["chunk-size-trigger"]) {
			return fmt.Errorf("Error reading chunk_size_trigger: %v", err)
		}
	}
	if err = d.Set("chunk_size_weight", flattenWafHttpProtocolParameterRestriction(o["chunk-size-weight"], d, "chunk_size_weight", sv)); err != nil {
		if !fortiAPIPatch(o["chunk-size-weight"]) {
			return fmt.Errorf("Error reading chunk_size_weight: %v", err)
		}
	}
	if err = d.Set("internal_resource_limits_check", flattenWafHttpProtocolParameterRestriction(o["Internal-resource-limits-check"], d, "internal_resource_limits_check", sv)); err != nil {
		if !fortiAPIPatch(o["Internal-resource-limits-check"]) {
			return fmt.Errorf("Error reading internal_resource_limits_check: %v", err)
		}
	}
	if err = d.Set("internal_resource_limits_action", flattenWafHttpProtocolParameterRestriction(o["Internal-resource-limits-action"], d, "internal_resource_limits_action", sv)); err != nil {
		if !fortiAPIPatch(o["Internal-resource-limits-action"]) {
			return fmt.Errorf("Error reading internal_resource_limits_action: %v", err)
		}
	}
	if err = d.Set("internal_resource_limits_block_period", flattenWafHttpProtocolParameterRestriction(o["Internal-resource-limits-block-period"], d, "internal_resource_limits_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["Internal-resource-limits-block-period"]) {
			return fmt.Errorf("Error reading internal_resource_limits_block_period: %v", err)
		}
	}
	if err = d.Set("internal_resource_limits_severity", flattenWafHttpProtocolParameterRestriction(o["Internal-resource-limits-severity"], d, "internal_resource_limits_severity", sv)); err != nil {
		if !fortiAPIPatch(o["Internal-resource-limits-severity"]) {
			return fmt.Errorf("Error reading internal_resource_limits_severity: %v", err)
		}
	}
	if err = d.Set("internal_resource_limits_trigger", flattenWafHttpProtocolParameterRestriction(o["Internal-resource-limits-trigger"], d, "internal_resource_limits_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["Internal-resource-limits-trigger"]) {
			return fmt.Errorf("Error reading internal_resource_limits_trigger: %v", err)
		}
	}
	if err = d.Set("internal_resource_limits_threat_weight", flattenWafHttpProtocolParameterRestriction(o["Internal-resource-limits-threat-weight"], d, "internal_resource_limits_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["Internal-resource-limits-threat-weight"]) {
			return fmt.Errorf("Error reading internal_resource_limits_threat_weight: %v", err)
		}
	}
	if err = d.Set("rpc_protocol_check", flattenWafHttpProtocolParameterRestriction(o["rpc-protocol-check"], d, "rpc_protocol_check", sv)); err != nil {
		if !fortiAPIPatch(o["rpc-protocol-check"]) {
			return fmt.Errorf("Error reading rpc_protocol_check: %v", err)
		}
	}
	if err = d.Set("rpc_protocol_action", flattenWafHttpProtocolParameterRestriction(o["rpc-protocol-action"], d, "rpc_protocol_action", sv)); err != nil {
		if !fortiAPIPatch(o["rpc-protocol-action"]) {
			return fmt.Errorf("Error reading rpc_protocol_action: %v", err)
		}
	}
	if err = d.Set("rpc_protocol_block_period", flattenWafHttpProtocolParameterRestriction(o["rpc-protocol-block-period"], d, "rpc_protocol_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["rpc-protocol-block-period"]) {
			return fmt.Errorf("Error reading rpc_protocol_block_period: %v", err)
		}
	}
	if err = d.Set("rpc_protocol_severity", flattenWafHttpProtocolParameterRestriction(o["rpc-protocol-severity"], d, "rpc_protocol_severity", sv)); err != nil {
		if !fortiAPIPatch(o["rpc-protocol-severity"]) {
			return fmt.Errorf("Error reading rpc_protocol_severity: %v", err)
		}
	}
	if err = d.Set("rpc_protocol_trigger", flattenWafHttpProtocolParameterRestriction(o["rpc-protocol-trigger"], d, "rpc_protocol_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["rpc-protocol-trigger"]) {
			return fmt.Errorf("Error reading rpc_protocol_trigger: %v", err)
		}
	}
	if err = d.Set("rpc_protocol_threat_weight", flattenWafHttpProtocolParameterRestriction(o["rpc-protocol-threat-weight"], d, "rpc_protocol_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["rpc-protocol-threat-weight"]) {
			return fmt.Errorf("Error reading rpc_protocol_threat_weight: %v", err)
		}
	}
	if err = d.Set("duplicate_paramname_check", flattenWafHttpProtocolParameterRestriction(o["duplicate-paramname-check"], d, "duplicate_paramname_check", sv)); err != nil {
		if !fortiAPIPatch(o["duplicate-paramname-check"]) {
			return fmt.Errorf("Error reading duplicate_paramname_check: %v", err)
		}
	}
	if err = d.Set("duplicate_paramname_action", flattenWafHttpProtocolParameterRestriction(o["duplicate-paramname-action"], d, "duplicate_paramname_action", sv)); err != nil {
		if !fortiAPIPatch(o["duplicate-paramname-action"]) {
			return fmt.Errorf("Error reading duplicate_paramname_action: %v", err)
		}
	}
	if err = d.Set("duplicate_paramname_block_period", flattenWafHttpProtocolParameterRestriction(o["duplicate-paramname-block-period"], d, "duplicate_paramname_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["duplicate-paramname-block-period"]) {
			return fmt.Errorf("Error reading duplicate_paramname_block_period: %v", err)
		}
	}
	if err = d.Set("duplicate_paramname_threat_weight", flattenWafHttpProtocolParameterRestriction(o["duplicate-paramname-threat-weight"], d, "duplicate_paramname_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["duplicate-paramname-threat-weight"]) {
			return fmt.Errorf("Error reading duplicate_paramname_threat_weight: %v", err)
		}
	}
	if err = d.Set("duplicate_paramname_severity", flattenWafHttpProtocolParameterRestriction(o["duplicate-paramname-severity"], d, "duplicate_paramname_severity", sv)); err != nil {
		if !fortiAPIPatch(o["duplicate-paramname-severity"]) {
			return fmt.Errorf("Error reading duplicate_paramname_severity: %v", err)
		}
	}
	if err = d.Set("duplicate_paramname_trigger", flattenWafHttpProtocolParameterRestriction(o["duplicate-paramname-trigger"], d, "duplicate_paramname_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["duplicate-paramname-trigger"]) {
			return fmt.Errorf("Error reading duplicate_paramname_trigger: %v", err)
		}
	}
	if err = d.Set("odd_and_even_space_attack_check", flattenWafHttpProtocolParameterRestriction(o["odd-and-even-space-attack-check"], d, "odd_and_even_space_attack_check", sv)); err != nil {
		if !fortiAPIPatch(o["odd-and-even-space-attack-check"]) {
			return fmt.Errorf("Error reading odd_and_even_space_attack_check: %v", err)
		}
	}
	if err = d.Set("odd_and_even_space_attack_action", flattenWafHttpProtocolParameterRestriction(o["odd-and-even-space-attack-action"], d, "odd_and_even_space_attack_action", sv)); err != nil {
		if !fortiAPIPatch(o["odd-and-even-space-attack-action"]) {
			return fmt.Errorf("Error reading odd_and_even_space_attack_action: %v", err)
		}
	}
	if err = d.Set("odd_and_even_space_attack_block_period", flattenWafHttpProtocolParameterRestriction(o["odd-and-even-space-attack-block-period"], d, "odd_and_even_space_attack_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["odd-and-even-space-attack-block-period"]) {
			return fmt.Errorf("Error reading odd_and_even_space_attack_block_period: %v", err)
		}
	}
	if err = d.Set("odd_and_even_space_attack_severity", flattenWafHttpProtocolParameterRestriction(o["odd-and-even-space-attack-severity"], d, "odd_and_even_space_attack_severity", sv)); err != nil {
		if !fortiAPIPatch(o["odd-and-even-space-attack-severity"]) {
			return fmt.Errorf("Error reading odd_and_even_space_attack_severity: %v", err)
		}
	}
	if err = d.Set("odd_and_even_space_attack_trigger", flattenWafHttpProtocolParameterRestriction(o["odd-and-even-space-attack-trigger"], d, "odd_and_even_space_attack_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["odd-and-even-space-attack-trigger"]) {
			return fmt.Errorf("Error reading odd_and_even_space_attack_trigger: %v", err)
		}
	}
	if err = d.Set("odd_and_even_space_attack_weight", flattenWafHttpProtocolParameterRestriction(o["odd-and-even-space-attack-weight"], d, "odd_and_even_space_attack_weight", sv)); err != nil {
		if !fortiAPIPatch(o["odd-and-even-space-attack-weight"]) {
			return fmt.Errorf("Error reading odd_and_even_space_attack_weight: %v", err)
		}
	}
	if err = d.Set("cl_te_coexist_check", flattenWafHttpProtocolParameterRestriction(o["cl-te-coexist-check"], d, "cl_te_coexist_check", sv)); err != nil {
		if !fortiAPIPatch(o["cl-te-coexist-check"]) {
			return fmt.Errorf("Error reading cl_te_coexist_check: %v", err)
		}
	}
	if err = d.Set("cl_te_coexist_action", flattenWafHttpProtocolParameterRestriction(o["cl-te-coexist-action"], d, "cl_te_coexist_action", sv)); err != nil {
		if !fortiAPIPatch(o["cl-te-coexist-action"]) {
			return fmt.Errorf("Error reading cl_te_coexist_action: %v", err)
		}
	}
	if err = d.Set("cl_te_coexist_block_period", flattenWafHttpProtocolParameterRestriction(o["cl-te-coexist-block-period"], d, "cl_te_coexist_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["cl-te-coexist-block-period"]) {
			return fmt.Errorf("Error reading cl_te_coexist_block_period: %v", err)
		}
	}
	if err = d.Set("cl_te_coexist_severity", flattenWafHttpProtocolParameterRestriction(o["cl-te-coexist-severity"], d, "cl_te_coexist_severity", sv)); err != nil {
		if !fortiAPIPatch(o["cl-te-coexist-severity"]) {
			return fmt.Errorf("Error reading cl_te_coexist_severity: %v", err)
		}
	}
	if err = d.Set("cl_te_coexist_trigger", flattenWafHttpProtocolParameterRestriction(o["cl-te-coexist-trigger"], d, "cl_te_coexist_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["cl-te-coexist-trigger"]) {
			return fmt.Errorf("Error reading cl_te_coexist_trigger: %v", err)
		}
	}
	if err = d.Set("cl_te_coexist_threat_weight", flattenWafHttpProtocolParameterRestriction(o["cl-te-coexist-threat-weight"], d, "cl_te_coexist_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["cl-te-coexist-threat-weight"]) {
			return fmt.Errorf("Error reading cl_te_coexist_threat_weight: %v", err)
		}
	}
	if err = d.Set("inconsistent_cl_check", flattenWafHttpProtocolParameterRestriction(o["inconsistent-cl-check"], d, "inconsistent_cl_check", sv)); err != nil {
		if !fortiAPIPatch(o["inconsistent-cl-check"]) {
			return fmt.Errorf("Error reading inconsistent_cl_check: %v", err)
		}
	}
	if err = d.Set("inconsistent_cl_action", flattenWafHttpProtocolParameterRestriction(o["inconsistent-cl-action"], d, "inconsistent_cl_action", sv)); err != nil {
		if !fortiAPIPatch(o["inconsistent-cl-action"]) {
			return fmt.Errorf("Error reading inconsistent_cl_action: %v", err)
		}
	}
	if err = d.Set("inconsistent_cl_block_period", flattenWafHttpProtocolParameterRestriction(o["inconsistent-cl-block-period"], d, "inconsistent_cl_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["inconsistent-cl-block-period"]) {
			return fmt.Errorf("Error reading inconsistent_cl_block_period: %v", err)
		}
	}
	if err = d.Set("inconsistent_cl_severity", flattenWafHttpProtocolParameterRestriction(o["inconsistent-cl-severity"], d, "inconsistent_cl_severity", sv)); err != nil {
		if !fortiAPIPatch(o["inconsistent-cl-severity"]) {
			return fmt.Errorf("Error reading inconsistent_cl_severity: %v", err)
		}
	}
	if err = d.Set("inconsistent_cl_trigger", flattenWafHttpProtocolParameterRestriction(o["inconsistent-cl-trigger"], d, "inconsistent_cl_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["inconsistent-cl-trigger"]) {
			return fmt.Errorf("Error reading inconsistent_cl_trigger: %v", err)
		}
	}
	if err = d.Set("inconsistent_cl_threat_weight", flattenWafHttpProtocolParameterRestriction(o["inconsistent-cl-threat-weight"], d, "inconsistent_cl_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["inconsistent-cl-threat-weight"]) {
			return fmt.Errorf("Error reading inconsistent_cl_threat_weight: %v", err)
		}
	}
	if err = d.Set("missing_host_check", flattenWafHttpProtocolParameterRestriction(o["missing-host-check"], d, "missing_host_check", sv)); err != nil {
		if !fortiAPIPatch(o["missing-host-check"]) {
			return fmt.Errorf("Error reading missing_host_check: %v", err)
		}
	}
	if err = d.Set("missing_host_action", flattenWafHttpProtocolParameterRestriction(o["missing-host-action"], d, "missing_host_action", sv)); err != nil {
		if !fortiAPIPatch(o["missing-host-action"]) {
			return fmt.Errorf("Error reading missing_host_action: %v", err)
		}
	}
	if err = d.Set("missing_host_block_period", flattenWafHttpProtocolParameterRestriction(o["missing-host-block-period"], d, "missing_host_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["missing-host-block-period"]) {
			return fmt.Errorf("Error reading missing_host_block_period: %v", err)
		}
	}
	if err = d.Set("missing_host_severity", flattenWafHttpProtocolParameterRestriction(o["missing-host-severity"], d, "missing_host_severity", sv)); err != nil {
		if !fortiAPIPatch(o["missing-host-severity"]) {
			return fmt.Errorf("Error reading missing_host_severity: %v", err)
		}
	}
	if err = d.Set("missing_host_trigger", flattenWafHttpProtocolParameterRestriction(o["missing-host-trigger"], d, "missing_host_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["missing-host-trigger"]) {
			return fmt.Errorf("Error reading missing_host_trigger: %v", err)
		}
	}
	if err = d.Set("missing_host_threat_weight", flattenWafHttpProtocolParameterRestriction(o["missing-host-threat-weight"], d, "missing_host_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["missing-host-threat-weight"]) {
			return fmt.Errorf("Error reading missing_host_threat_weight: %v", err)
		}
	}
	if err = d.Set("range_overlapping_check", flattenWafHttpProtocolParameterRestriction(o["range-overlapping-check"], d, "range_overlapping_check", sv)); err != nil {
		if !fortiAPIPatch(o["range-overlapping-check"]) {
			return fmt.Errorf("Error reading range_overlapping_check: %v", err)
		}
	}
	if err = d.Set("range_overlapping_action", flattenWafHttpProtocolParameterRestriction(o["range-overlapping-action"], d, "range_overlapping_action", sv)); err != nil {
		if !fortiAPIPatch(o["range-overlapping-action"]) {
			return fmt.Errorf("Error reading range_overlapping_action: %v", err)
		}
	}
	if err = d.Set("range_overlapping_block_period", flattenWafHttpProtocolParameterRestriction(o["range-overlapping-block-period"], d, "range_overlapping_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["range-overlapping-block-period"]) {
			return fmt.Errorf("Error reading range_overlapping_block_period: %v", err)
		}
	}
	if err = d.Set("range_overlapping_severity", flattenWafHttpProtocolParameterRestriction(o["range-overlapping-severity"], d, "range_overlapping_severity", sv)); err != nil {
		if !fortiAPIPatch(o["range-overlapping-severity"]) {
			return fmt.Errorf("Error reading range_overlapping_severity: %v", err)
		}
	}
	if err = d.Set("range_overlapping_trigger", flattenWafHttpProtocolParameterRestriction(o["range-overlapping-trigger"], d, "range_overlapping_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["range-overlapping-trigger"]) {
			return fmt.Errorf("Error reading range_overlapping_trigger: %v", err)
		}
	}
	if err = d.Set("range_overlapping_threat_weight", flattenWafHttpProtocolParameterRestriction(o["range-overlapping-threat-weight"], d, "range_overlapping_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["range-overlapping-threat-weight"]) {
			return fmt.Errorf("Error reading range_overlapping_threat_weight: %v", err)
		}
	}
	if err = d.Set("multipart_formdata_bad_request_check", flattenWafHttpProtocolParameterRestriction(o["multipart-formdata-bad-request-check"], d, "multipart_formdata_bad_request_check", sv)); err != nil {
		if !fortiAPIPatch(o["multipart-formdata-bad-request-check"]) {
			return fmt.Errorf("Error reading multipart_formdata_bad_request_check: %v", err)
		}
	}
	if err = d.Set("multipart_formdata_bad_request_action", flattenWafHttpProtocolParameterRestriction(o["multipart-formdata-bad-request-action"], d, "multipart_formdata_bad_request_action", sv)); err != nil {
		if !fortiAPIPatch(o["multipart-formdata-bad-request-action"]) {
			return fmt.Errorf("Error reading multipart_formdata_bad_request_action: %v", err)
		}
	}
	if err = d.Set("multipart_formdata_bad_request_block_period", flattenWafHttpProtocolParameterRestriction(o["multipart-formdata-bad-request-block-period"], d, "multipart_formdata_bad_request_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["multipart-formdata-bad-request-block-period"]) {
			return fmt.Errorf("Error reading multipart_formdata_bad_request_block_period: %v", err)
		}
	}
	if err = d.Set("multipart_formdata_bad_request_severity", flattenWafHttpProtocolParameterRestriction(o["multipart-formdata-bad-request-severity"], d, "multipart_formdata_bad_request_severity", sv)); err != nil {
		if !fortiAPIPatch(o["multipart-formdata-bad-request-severity"]) {
			return fmt.Errorf("Error reading multipart_formdata_bad_request_severity: %v", err)
		}
	}
	if err = d.Set("multipart_formdata_bad_request_trigger", flattenWafHttpProtocolParameterRestriction(o["multipart-formdata-bad-request-trigger"], d, "multipart_formdata_bad_request_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["multipart-formdata-bad-request-trigger"]) {
			return fmt.Errorf("Error reading multipart_formdata_bad_request_trigger: %v", err)
		}
	}
	if err = d.Set("multipart_formdata_bad_request_threat_weight", flattenWafHttpProtocolParameterRestriction(o["multipart-formdata-bad-request-threat-weight"], d, "multipart_formdata_bad_request_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["multipart-formdata-bad-request-threat-weight"]) {
			return fmt.Errorf("Error reading multipart_formdata_bad_request_threat_weight: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_check", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-check"], d, "h2_rst_stream_check", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-check"]) {
			return fmt.Errorf("Error reading h2_rst_stream_check: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream"], d, "h2_rst_stream", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream"]) {
			return fmt.Errorf("Error reading h2_rst_stream: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_action", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-action"], d, "h2_rst_stream_action", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-action"]) {
			return fmt.Errorf("Error reading h2_rst_stream_action: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_block_period", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-block-period"], d, "h2_rst_stream_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-block-period"]) {
			return fmt.Errorf("Error reading h2_rst_stream_block_period: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_severity", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-severity"], d, "h2_rst_stream_severity", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-severity"]) {
			return fmt.Errorf("Error reading h2_rst_stream_severity: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_trigger", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-trigger"], d, "h2_rst_stream_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-trigger"]) {
			return fmt.Errorf("Error reading h2_rst_stream_trigger: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_threat_weight", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-threat-weight"], d, "h2_rst_stream_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-threat-weight"]) {
			return fmt.Errorf("Error reading h2_rst_stream_threat_weight: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_freq_check", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-freq-check"], d, "h2_rst_stream_freq_check", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-freq-check"]) {
			return fmt.Errorf("Error reading h2_rst_stream_freq_check: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_freq", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-freq"], d, "h2_rst_stream_freq", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-freq"]) {
			return fmt.Errorf("Error reading h2_rst_stream_freq: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_freq_action", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-freq-action"], d, "h2_rst_stream_freq_action", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-freq-action"]) {
			return fmt.Errorf("Error reading h2_rst_stream_freq_action: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_freq_block_period", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-freq-block-period"], d, "h2_rst_stream_freq_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-freq-block-period"]) {
			return fmt.Errorf("Error reading h2_rst_stream_freq_block_period: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_freq_severity", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-freq-severity"], d, "h2_rst_stream_freq_severity", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-freq-severity"]) {
			return fmt.Errorf("Error reading h2_rst_stream_freq_severity: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_freq_trigger", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-freq-trigger"], d, "h2_rst_stream_freq_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-freq-trigger"]) {
			return fmt.Errorf("Error reading h2_rst_stream_freq_trigger: %v", err)
		}
	}
	if err = d.Set("h2_rst_stream_freq_threat_weight", flattenWafHttpProtocolParameterRestriction(o["h2-rst-stream-freq-threat-weight"], d, "h2_rst_stream_freq_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["h2-rst-stream-freq-threat-weight"]) {
			return fmt.Errorf("Error reading h2_rst_stream_freq_threat_weight: %v", err)
		}
	}
	if err = d.Set("h3_max_table_capacity_check", flattenWafHttpProtocolParameterRestriction(o["h3-max-table-capacity-check"], d, "h3_max_table_capacity_check", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-table-capacity-check"]) {
			return fmt.Errorf("Error reading h3_max_table_capacity_check: %v", err)
		}
	}
	if err = d.Set("h3_max_table_capacity", flattenWafHttpProtocolParameterRestriction(o["h3-max-table-capacity"], d, "h3_max_table_capacity", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-table-capacity"]) {
			return fmt.Errorf("Error reading h3_max_table_capacity: %v", err)
		}
	}
	if err = d.Set("h3_max_table_capacity_action", flattenWafHttpProtocolParameterRestriction(o["h3-max-table-capacity-action"], d, "h3_max_table_capacity_action", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-table-capacity-action"]) {
			return fmt.Errorf("Error reading h3_max_table_capacity_action: %v", err)
		}
	}
	if err = d.Set("h3_max_table_capacity_block_period", flattenWafHttpProtocolParameterRestriction(o["h3-max-table-capacity-block-period"], d, "h3_max_table_capacity_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-table-capacity-block-period"]) {
			return fmt.Errorf("Error reading h3_max_table_capacity_block_period: %v", err)
		}
	}
	if err = d.Set("h3_max_table_capacity_severity", flattenWafHttpProtocolParameterRestriction(o["h3-max-table-capacity-severity"], d, "h3_max_table_capacity_severity", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-table-capacity-severity"]) {
			return fmt.Errorf("Error reading h3_max_table_capacity_severity: %v", err)
		}
	}
	if err = d.Set("h3_max_table_capacity_trigger", flattenWafHttpProtocolParameterRestriction(o["h3-max-table-capacity-trigger"], d, "h3_max_table_capacity_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-table-capacity-trigger"]) {
			return fmt.Errorf("Error reading h3_max_table_capacity_trigger: %v", err)
		}
	}
	if err = d.Set("h3_max_table_capacity_threat_weight", flattenWafHttpProtocolParameterRestriction(o["h3-max-table-capacity-threat-weight"], d, "h3_max_table_capacity_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-table-capacity-threat-weight"]) {
			return fmt.Errorf("Error reading h3_max_table_capacity_threat_weight: %v", err)
		}
	}
	if err = d.Set("h3_max_field_section_size_check", flattenWafHttpProtocolParameterRestriction(o["h3-max-field-section-size-check"], d, "h3_max_field_section_size_check", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-field-section-size-check"]) {
			return fmt.Errorf("Error reading h3_max_field_section_size_check: %v", err)
		}
	}
	if err = d.Set("h3_max_field_section_size", flattenWafHttpProtocolParameterRestriction(o["h3-max-field-section-size"], d, "h3_max_field_section_size", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-field-section-size"]) {
			return fmt.Errorf("Error reading h3_max_field_section_size: %v", err)
		}
	}
	if err = d.Set("h3_max_field_section_size_action", flattenWafHttpProtocolParameterRestriction(o["h3-max-field-section-size-action"], d, "h3_max_field_section_size_action", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-field-section-size-action"]) {
			return fmt.Errorf("Error reading h3_max_field_section_size_action: %v", err)
		}
	}
	if err = d.Set("h3_max_field_section_size_block_period", flattenWafHttpProtocolParameterRestriction(o["h3-max-field-section-size-block-period"], d, "h3_max_field_section_size_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-field-section-size-block-period"]) {
			return fmt.Errorf("Error reading h3_max_field_section_size_block_period: %v", err)
		}
	}
	if err = d.Set("h3_max_field_section_size_severity", flattenWafHttpProtocolParameterRestriction(o["h3-max-field-section-size-severity"], d, "h3_max_field_section_size_severity", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-field-section-size-severity"]) {
			return fmt.Errorf("Error reading h3_max_field_section_size_severity: %v", err)
		}
	}
	if err = d.Set("h3_max_field_section_size_trigger", flattenWafHttpProtocolParameterRestriction(o["h3-max-field-section-size-trigger"], d, "h3_max_field_section_size_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-field-section-size-trigger"]) {
			return fmt.Errorf("Error reading h3_max_field_section_size_trigger: %v", err)
		}
	}
	if err = d.Set("h3_max_field_section_size_threat_weight", flattenWafHttpProtocolParameterRestriction(o["h3-max-field-section-size-threat-weight"], d, "h3_max_field_section_size_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["h3-max-field-section-size-threat-weight"]) {
			return fmt.Errorf("Error reading h3_max_field_section_size_threat_weight: %v", err)
		}
	}
	if err = d.Set("h3_block_streams_check", flattenWafHttpProtocolParameterRestriction(o["h3-block-streams-check"], d, "h3_block_streams_check", sv)); err != nil {
		if !fortiAPIPatch(o["h3-block-streams-check"]) {
			return fmt.Errorf("Error reading h3_block_streams_check: %v", err)
		}
	}
	if err = d.Set("h3_block_streams", flattenWafHttpProtocolParameterRestriction(o["h3-block-streams"], d, "h3_block_streams", sv)); err != nil {
		if !fortiAPIPatch(o["h3-block-streams"]) {
			return fmt.Errorf("Error reading h3_block_streams: %v", err)
		}
	}
	if err = d.Set("h3_block_streams_action", flattenWafHttpProtocolParameterRestriction(o["h3-block-streams-action"], d, "h3_block_streams_action", sv)); err != nil {
		if !fortiAPIPatch(o["h3-block-streams-action"]) {
			return fmt.Errorf("Error reading h3_block_streams_action: %v", err)
		}
	}
	if err = d.Set("h3_block_streams_block_period", flattenWafHttpProtocolParameterRestriction(o["h3-block-streams-block-period"], d, "h3_block_streams_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["h3-block-streams-block-period"]) {
			return fmt.Errorf("Error reading h3_block_streams_block_period: %v", err)
		}
	}
	if err = d.Set("h3_block_streams_severity", flattenWafHttpProtocolParameterRestriction(o["h3-block-streams-severity"], d, "h3_block_streams_severity", sv)); err != nil {
		if !fortiAPIPatch(o["h3-block-streams-severity"]) {
			return fmt.Errorf("Error reading h3_block_streams_severity: %v", err)
		}
	}
	if err = d.Set("h3_block_streams_trigger", flattenWafHttpProtocolParameterRestriction(o["h3-block-streams-trigger"], d, "h3_block_streams_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["h3-block-streams-trigger"]) {
			return fmt.Errorf("Error reading h3_block_streams_trigger: %v", err)
		}
	}
	if err = d.Set("h3_block_streams_threat_weight", flattenWafHttpProtocolParameterRestriction(o["h3-block-streams-threat-weight"], d, "h3_block_streams_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["h3-block-streams-threat-weight"]) {
			return fmt.Errorf("Error reading h3_block_streams_threat_weight: %v", err)
		}
	}
	if err = d.Set("h3_bidir_concurrent_stream_check", flattenWafHttpProtocolParameterRestriction(o["h3-bidir-concurrent-stream-check"], d, "h3_bidir_concurrent_stream_check", sv)); err != nil {
		if !fortiAPIPatch(o["h3-bidir-concurrent-stream-check"]) {
			return fmt.Errorf("Error reading h3_bidir_concurrent_stream_check: %v", err)
		}
	}
	if err = d.Set("h3_bidir_concurrent_stream", flattenWafHttpProtocolParameterRestriction(o["h3-bidir-concurrent-stream"], d, "h3_bidir_concurrent_stream", sv)); err != nil {
		if !fortiAPIPatch(o["h3-bidir-concurrent-stream"]) {
			return fmt.Errorf("Error reading h3_bidir_concurrent_stream: %v", err)
		}
	}
	if err = d.Set("h3_bidir_concurrent_stream_action", flattenWafHttpProtocolParameterRestriction(o["h3-bidir-concurrent-stream-action"], d, "h3_bidir_concurrent_stream_action", sv)); err != nil {
		if !fortiAPIPatch(o["h3-bidir-concurrent-stream-action"]) {
			return fmt.Errorf("Error reading h3_bidir_concurrent_stream_action: %v", err)
		}
	}
	if err = d.Set("h3_bidir_concurrent_stream_block_period", flattenWafHttpProtocolParameterRestriction(o["h3-bidir-concurrent-stream-block-period"], d, "h3_bidir_concurrent_stream_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["h3-bidir-concurrent-stream-block-period"]) {
			return fmt.Errorf("Error reading h3_bidir_concurrent_stream_block_period: %v", err)
		}
	}
	if err = d.Set("h3_bidir_concurrent_stream_severity", flattenWafHttpProtocolParameterRestriction(o["h3-bidir-concurrent-stream-severity"], d, "h3_bidir_concurrent_stream_severity", sv)); err != nil {
		if !fortiAPIPatch(o["h3-bidir-concurrent-stream-severity"]) {
			return fmt.Errorf("Error reading h3_bidir_concurrent_stream_severity: %v", err)
		}
	}
	if err = d.Set("h3_bidir_concurrent_stream_trigger", flattenWafHttpProtocolParameterRestriction(o["h3-bidir-concurrent-stream-trigger"], d, "h3_bidir_concurrent_stream_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["h3-bidir-concurrent-stream-trigger"]) {
			return fmt.Errorf("Error reading h3_bidir_concurrent_stream_trigger: %v", err)
		}
	}
	if err = d.Set("h3_bidir_concurrent_stream_threat_weight", flattenWafHttpProtocolParameterRestriction(o["h3-bidir-concurrent-stream-threat-weight"], d, "h3_bidir_concurrent_stream_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["h3-bidir-concurrent-stream-threat-weight"]) {
			return fmt.Errorf("Error reading h3_bidir_concurrent_stream_threat_weight: %v", err)
		}
	}
	if err = d.Set("h3_unidir_concurrent_stream_check", flattenWafHttpProtocolParameterRestriction(o["h3-unidir-concurrent-stream-check"], d, "h3_unidir_concurrent_stream_check", sv)); err != nil {
		if !fortiAPIPatch(o["h3-unidir-concurrent-stream-check"]) {
			return fmt.Errorf("Error reading h3_unidir_concurrent_stream_check: %v", err)
		}
	}
	if err = d.Set("h3_unidir_concurrent_stream", flattenWafHttpProtocolParameterRestriction(o["h3-unidir-concurrent-stream"], d, "h3_unidir_concurrent_stream", sv)); err != nil {
		if !fortiAPIPatch(o["h3-unidir-concurrent-stream"]) {
			return fmt.Errorf("Error reading h3_unidir_concurrent_stream: %v", err)
		}
	}
	if err = d.Set("h3_unidir_concurrent_stream_action", flattenWafHttpProtocolParameterRestriction(o["h3-unidir-concurrent-stream-action"], d, "h3_unidir_concurrent_stream_action", sv)); err != nil {
		if !fortiAPIPatch(o["h3-unidir-concurrent-stream-action"]) {
			return fmt.Errorf("Error reading h3_unidir_concurrent_stream_action: %v", err)
		}
	}
	if err = d.Set("h3_unidir_concurrent_stream_block_period", flattenWafHttpProtocolParameterRestriction(o["h3-unidir-concurrent-stream-block-period"], d, "h3_unidir_concurrent_stream_block_period", sv)); err != nil {
		if !fortiAPIPatch(o["h3-unidir-concurrent-stream-block-period"]) {
			return fmt.Errorf("Error reading h3_unidir_concurrent_stream_block_period: %v", err)
		}
	}
	if err = d.Set("h3_unidir_concurrent_stream_severity", flattenWafHttpProtocolParameterRestriction(o["h3-unidir-concurrent-stream-severity"], d, "h3_unidir_concurrent_stream_severity", sv)); err != nil {
		if !fortiAPIPatch(o["h3-unidir-concurrent-stream-severity"]) {
			return fmt.Errorf("Error reading h3_unidir_concurrent_stream_severity: %v", err)
		}
	}
	if err = d.Set("h3_unidir_concurrent_stream_trigger", flattenWafHttpProtocolParameterRestriction(o["h3-unidir-concurrent-stream-trigger"], d, "h3_unidir_concurrent_stream_trigger", sv)); err != nil {
		if !fortiAPIPatch(o["h3-unidir-concurrent-stream-trigger"]) {
			return fmt.Errorf("Error reading h3_unidir_concurrent_stream_trigger: %v", err)
		}
	}
	if err = d.Set("h3_unidir_concurrent_stream_threat_weight", flattenWafHttpProtocolParameterRestriction(o["h3-unidir-concurrent-stream-threat-weight"], d, "h3_unidir_concurrent_stream_threat_weight", sv)); err != nil {
		if !fortiAPIPatch(o["h3-unidir-concurrent-stream-threat-weight"]) {
			return fmt.Errorf("Error reading h3_unidir_concurrent_stream_threat_weight: %v", err)
		}
	}
	if err = d.Set("mkey", flattenWafHttpProtocolParameterRestriction(o["name"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandWafHttpProtocolParameterRestriction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafHttpProtocolParameterRestriction(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_http_header_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_http_content_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_content_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-content-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_content_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_content_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-content-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_content_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_content_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-content-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_http_content_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_content_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-content-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_http_content_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_content_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-content-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_http_content_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_content_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-content-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_http_content_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_content_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-content-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-length"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_version_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_version_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-version-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_version_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_version_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-version-check-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_version_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_version_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-version-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_version_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_version_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-version-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_version_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_version_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-version-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_version_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_version_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-version-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_cookie_in_request_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_cookie_in_request_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie-in-request-check"] = t
		}
	}
	if v, ok := d.GetOk("max_cookie_in_request"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_cookie_in_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie-in-request"] = t
		}
	}
	if v, ok := d.GetOk("max_cookie_in_request_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_cookie_in_request_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie-in-request-action"] = t
		}
	}
	if v, ok := d.GetOk("max_cookie_in_request_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_cookie_in_request_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie-in-request-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_cookie_in_request_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_cookie_in_request_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie-in-request-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_cookie_in_request_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_cookie_in_request_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie-in-request-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_cookie_in_request_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_cookie_in_request_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-cookie-in-request-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_header_line_request_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_header_line_request_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line-request-check"] = t
		}
	}
	if v, ok := d.GetOk("max_header_line_request"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_header_line_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line-request"] = t
		}
	}
	if v, ok := d.GetOk("max_header_line_request_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_header_line_request_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line-request-action"] = t
		}
	}
	if v, ok := d.GetOk("max_header_line_request_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_header_line_request_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line-request-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_header_line_request_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_header_line_request_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line-request-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_header_line_request_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_header_line_request_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line-request-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_header_line_request_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_header_line_request_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-header-line-request-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_request_method_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_request_method_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-request-method-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_request_method_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_request_method_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-request-method-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_request_method_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_request_method_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-request-method-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_request_method_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_request_method_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-request-method-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_request_method_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_request_method_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-request-method-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_http_request_method_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_http_request_method_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-http-request-method-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-check"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-action"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_url_parameter_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_parameter_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-parameter-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_host_name_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_host_name_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-host-name-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_host_name_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_host_name_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-host-name-check-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_host_name_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_host_name_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-host-name-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_host_name_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_host_name_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-host-name-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_host_name_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_host_name_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-host-name-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_host_name_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_host_name_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-host-name-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("number_of_ranges_in_range_header_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "number_of_ranges_in_range_header_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-of-ranges-in-range-header-check"] = t
		}
	}
	if v, ok := d.GetOk("number_of_ranges_in_range_header"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "number_of_ranges_in_range_header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-of-ranges-in-range-header"] = t
		}
	}
	if v, ok := d.GetOk("number_of_ranges_in_range_header_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "number_of_ranges_in_range_header_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-of-ranges-in-range-header-action"] = t
		}
	}
	if v, ok := d.GetOk("number_of_ranges_in_range_header_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "number_of_ranges_in_range_header_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-of-ranges-in-range-header-block-period"] = t
		}
	}
	if v, ok := d.GetOk("number_of_ranges_in_range_header_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "number_of_ranges_in_range_header_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-of-ranges-in-range-header-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("number_of_ranges_in_range_header_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "number_of_ranges_in_range_header_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-of-ranges-in-range-header-severity"] = t
		}
	}
	if v, ok := d.GetOk("number_of_ranges_in_range_header_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "number_of_ranges_in_range_header_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["number-of-ranges-in-range-header-trigger"] = t
		}
	}
	if v, ok := d.GetOk("http2_max_requests_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "http2_max_requests_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-max-requests-check"] = t
		}
	}
	if v, ok := d.GetOk("http2_max_requests"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "http2_max_requests", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-max-requests"] = t
		}
	}
	if v, ok := d.GetOk("http2_max_requests_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "http2_max_requests_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-max-requests-action"] = t
		}
	}
	if v, ok := d.GetOk("http2_max_requests_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "http2_max_requests_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-max-requests-block-period"] = t
		}
	}
	if v, ok := d.GetOk("http2_max_requests_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "http2_max_requests_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-max-requests-severity"] = t
		}
	}
	if v, ok := d.GetOk("http2_max_requests_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "http2_max_requests_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-max-requests-trigger"] = t
		}
	}
	if v, ok := d.GetOk("http2_max_requests_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "http2_max_requests_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http2-max-requests-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("exception_name"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "exception_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exception_name"] = t
		}
	}
	if v, ok := d.GetOk("block_malformed_request_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "block_malformed_request_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malformed-request-check"] = t
		}
	}
	if v, ok := d.GetOk("block_malformed_request_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "block_malformed_request_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malformed-request-action"] = t
		}
	}
	if v, ok := d.GetOk("block_malformed_request_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "block_malformed_request_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malformed-request-block-period"] = t
		}
	}
	if v, ok := d.GetOk("block_malformed_request_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "block_malformed_request_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malformed-request-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("block_malformed_request_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "block_malformed_request_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malformed-request-severity"] = t
		}
	}
	if v, ok := d.GetOk("block_malformed_request_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "block_malformed_request_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-malformed-request-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-length-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_length_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_length_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-length-check-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_length_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_length_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-length-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_length_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_length_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-length-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_length_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_length_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-length-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_length_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_length_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-length-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_type_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_type_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-type-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_type_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_type_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-type-check-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_type_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_type_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-type-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_type_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_type_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-type-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_type_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_type_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-type-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_content_type_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_content_type_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-content-type-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_response_code_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_response_code_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-response-code-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_response_code_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_response_code_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-response-code-check-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_response_code_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_response_code_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-response-code-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_response_code_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_response_code_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-response-code-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_response_code_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_response_code_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-response-code-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_response_code_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_response_code_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-response-code-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("post_request_ctype_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "post_request_ctype_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Post-request-ctype-check"] = t
		}
	}
	if v, ok := d.GetOk("post_request_ctype_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "post_request_ctype_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Post-request-ctype-check-action"] = t
		}
	}
	if v, ok := d.GetOk("post_request_ctype_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "post_request_ctype_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Post-request-ctype-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("post_request_ctype_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "post_request_ctype_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Post-request-ctype-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("post_request_ctype_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "post_request_ctype_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Post-request-ctype-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("post_request_ctype_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "post_request_ctype_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Post-request-ctype-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_name_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_name_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-name-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_name_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_name_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-name-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_name_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_name_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-name-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_name_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_name_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-name-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_name_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_name_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-name-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_name_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_name_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-name-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_name_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_name_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-name-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_value_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_value_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-value-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_value_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_value_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-value-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_value_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_value_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-value-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_value_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_value_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-value-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_value_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_value_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-value-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_value_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_value_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-value-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_http_header_value_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_header_value_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-header-value-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("parameter_name_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_name_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-name-check"] = t
		}
	}
	if v, ok := d.GetOk("parameter_name_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_name_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-name-check-action"] = t
		}
	}
	if v, ok := d.GetOk("parameter_name_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_name_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-name-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("parameter_name_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_name_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-name-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("parameter_name_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_name_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-name-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("parameter_name_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_name_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-name-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("parameter_value_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_value_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-value-check"] = t
		}
	}
	if v, ok := d.GetOk("parameter_value_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_value_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-value-check-action"] = t
		}
	}
	if v, ok := d.GetOk("parameter_value_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_value_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-value-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("parameter_value_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_value_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-value-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("parameter_value_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_value_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-value-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("parameter_value_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "parameter_value_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter-value-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_name_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_name_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-name-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_name_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_name_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-name-check-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_name_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_name_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-name-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_name_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_name_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-name-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_name_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_name_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-name-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_name_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_name_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-name-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_value_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_value_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-value-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_value_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_value_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-value-check-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_value_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_value_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-value-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_value_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_value_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-value-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_value_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_value_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-value-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_header_value_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_header_value_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Illegal-header-value-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_parameter_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_parameter_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-parameter-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_parameter_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_parameter_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-parameter-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_parameter_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_parameter_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-parameter-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_parameter_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_parameter_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-parameter-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_parameter_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_parameter_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-parameter-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_parameter_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_parameter_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-parameter-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_http_body_parameter_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_body_parameter_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-body-parameter-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_filename_length_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_filename_length_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-filename-length-check"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_filename_length"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_filename_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-filename-length"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_filename_length_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_filename_length_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-filename-length-action"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_filename_length_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_filename_length_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-filename-length-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_filename_length_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_filename_length_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-filename-length-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_filename_length_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_filename_length_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-filename-length-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_http_request_filename_length_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_http_request_filename_length_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-http-request-filename-length-trigger"] = t
		}
	}
	if v, ok := d.GetOk("web_socket_protocol_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "web_socket_protocol_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-socket-protocol-check"] = t
		}
	}
	if v, ok := d.GetOk("web_socket_protocol_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "web_socket_protocol_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-socket-protocol-action"] = t
		}
	}
	if v, ok := d.GetOk("web_socket_protocol_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "web_socket_protocol_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-socket-protocol-block-period"] = t
		}
	}
	if v, ok := d.GetOk("web_socket_protocol_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "web_socket_protocol_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-socket-protocol-severity"] = t
		}
	}
	if v, ok := d.GetOk("web_socket_protocol_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "web_socket_protocol_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-socket-protocol-trigger"] = t
		}
	}
	if v, ok := d.GetOk("web_socket_protocol_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "web_socket_protocol_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-socket-protocol-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_table_size_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_table_size_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-table-size-check"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_table_size"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_table_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-table-size"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_table_size_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_table_size_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-table-size-action"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_table_size_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_table_size_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-table-size-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_table_size_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_table_size_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-table-size-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_table_size_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_table_size_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-table-size-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_table_size_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_table_size_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-table-size-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_current_streams_num_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_current_streams_num_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-current-streams-num-check"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_current_streams_num"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_current_streams_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-current-streams-num"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_current_streams_num_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_current_streams_num_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-current-streams-num-action"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_current_streams_num_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_current_streams_num_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-current-streams-num-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_current_streams_num_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_current_streams_num_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-current-streams-num-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_current_streams_num_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_current_streams_num_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-current-streams-num-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_current_streams_num_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_current_streams_num_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-current-streams-num-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_initial_window_size_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_initial_window_size_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-initial-window-size-check"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_initial_window_size"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_initial_window_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-initial-window-size"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_initial_window_size_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_initial_window_size_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-initial-window-size-action"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_initial_window_size_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_initial_window_size_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-initial-window-size-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_initial_window_size_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_initial_window_size_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-initial-window-size-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_initial_window_size_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_initial_window_size_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-initial-window-size-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_initial_window_size_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_initial_window_size_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-initial-window-size-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_frame_size_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_frame_size_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-frame-size-check"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_frame_size"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_frame_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-frame-size"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_frame_size_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_frame_size_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-frame-size-action"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_frame_size_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_frame_size_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-frame-size-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_frame_size_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_frame_size_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-frame-size-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_frame_size_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_frame_size_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-frame-size-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_frame_size_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_frame_size_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-frame-size-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_list_size_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_list_size_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-list-size-check"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_list_size"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_list_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-list-size"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_list_size_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_list_size_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-list-size-action"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_list_size_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_list_size_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-list-size-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_list_size_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_list_size_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-list-size-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_list_size_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_list_size_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-list-size-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_setting_header_list_size_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_setting_header_list_size_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-setting-header-list-size-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_name_len_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_name_len_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-name-len-check"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_name_len"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_name_len", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-name-len"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_name_len_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_name_len_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-name-len-action"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_name_len_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_name_len_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-name-len-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_name_len_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_name_len_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-name-len-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_name_len_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_name_len_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-name-len-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_name_len_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_name_len_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-name-len-trigger"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_value_len_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_value_len_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-value-len-check"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_value_len"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_value_len", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-value-len"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_value_len_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_value_len_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-value-len-action"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_value_len_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_value_len_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-value-len-block-period"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_value_len_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_value_len_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-value-len-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_value_len_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_value_len_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-value-len-severity"] = t
		}
	}
	if v, ok := d.GetOk("max_url_param_value_len_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "max_url_param_value_len_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-url-param-value-len-trigger"] = t
		}
	}
	if v, ok := d.GetOk("url_param_name_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_name_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-name-check"] = t
		}
	}
	if v, ok := d.GetOk("url_param_name_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_name_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-name-check-action"] = t
		}
	}
	if v, ok := d.GetOk("url_param_name_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_name_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-name-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("url_param_name_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_name_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-name-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("url_param_name_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_name_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-name-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("url_param_name_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_name_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-name-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("url_param_value_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_value_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-value-check"] = t
		}
	}
	if v, ok := d.GetOk("url_param_value_check_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_value_check_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-value-check-action"] = t
		}
	}
	if v, ok := d.GetOk("url_param_value_check_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_value_check_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-value-check-block-period"] = t
		}
	}
	if v, ok := d.GetOk("url_param_value_check_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_value_check_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-value-check-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("url_param_value_check_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_value_check_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-value-check-severity"] = t
		}
	}
	if v, ok := d.GetOk("url_param_value_check_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "url_param_value_check_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-param-value-check-trigger"] = t
		}
	}
	if v, ok := d.GetOk("null_byte_in_url_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "null_byte_in_url_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-byte-in-url-check"] = t
		}
	}
	if v, ok := d.GetOk("null_byte_in_url_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "null_byte_in_url_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-byte-in-url-action"] = t
		}
	}
	if v, ok := d.GetOk("null_byte_in_url_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "null_byte_in_url_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-byte-in-url-block-period"] = t
		}
	}
	if v, ok := d.GetOk("null_byte_in_url_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "null_byte_in_url_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-byte-in-url-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("null_byte_in_url_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "null_byte_in_url_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-byte-in-url-severity"] = t
		}
	}
	if v, ok := d.GetOk("null_byte_in_url_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "null_byte_in_url_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["null-byte-in-url-trigger"] = t
		}
	}
	if v, ok := d.GetOk("illegal_byte_in_url_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_byte_in_url_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["illegal-byte-in-url-check"] = t
		}
	}
	if v, ok := d.GetOk("illegal_byte_in_url_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_byte_in_url_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["illegal-byte-in-url-action"] = t
		}
	}
	if v, ok := d.GetOk("illegal_byte_in_url_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_byte_in_url_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["illegal-byte-in-url-block-period"] = t
		}
	}
	if v, ok := d.GetOk("illegal_byte_in_url_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_byte_in_url_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["illegal-byte-in-url-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("illegal_byte_in_url_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_byte_in_url_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["illegal-byte-in-url-severity"] = t
		}
	}
	if v, ok := d.GetOk("illegal_byte_in_url_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "illegal_byte_in_url_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["illegal-byte-in-url-trigger"] = t
		}
	}
	if v, ok := d.GetOk("malformed_url_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "malformed_url_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-url-check"] = t
		}
	}
	if v, ok := d.GetOk("malformed_url_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "malformed_url_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-url-action"] = t
		}
	}
	if v, ok := d.GetOk("malformed_url_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "malformed_url_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-url-block-period"] = t
		}
	}
	if v, ok := d.GetOk("malformed_url_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "malformed_url_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-url-severity"] = t
		}
	}
	if v, ok := d.GetOk("malformed_url_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "malformed_url_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-url-trigger"] = t
		}
	}
	if v, ok := d.GetOk("malformed_url_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "malformed_url_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["malformed-url-weight"] = t
		}
	}
	if v, ok := d.GetOk("redundant_header_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "redundant_header_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-header-check"] = t
		}
	}
	if v, ok := d.GetOk("redundant_header_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "redundant_header_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-header-action"] = t
		}
	}
	if v, ok := d.GetOk("redundant_header_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "redundant_header_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-header-block-period"] = t
		}
	}
	if v, ok := d.GetOk("redundant_header_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "redundant_header_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-header-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("redundant_header_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "redundant_header_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-header-severity"] = t
		}
	}
	if v, ok := d.GetOk("redundant_header_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "redundant_header_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-header-trigger"] = t
		}
	}
	if v, ok := d.GetOk("chunk_size_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "chunk_size_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-size-check"] = t
		}
	}
	if v, ok := d.GetOk("chunk_size_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "chunk_size_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-size-action"] = t
		}
	}
	if v, ok := d.GetOk("chunk_size_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "chunk_size_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-size-block-period"] = t
		}
	}
	if v, ok := d.GetOk("chunk_size_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "chunk_size_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-size-severity"] = t
		}
	}
	if v, ok := d.GetOk("chunk_size_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "chunk_size_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-size-trigger"] = t
		}
	}
	if v, ok := d.GetOk("chunk_size_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "chunk_size_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-size-weight"] = t
		}
	}
	if v, ok := d.GetOk("internal_resource_limits_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "internal_resource_limits_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Internal-resource-limits-check"] = t
		}
	}
	if v, ok := d.GetOk("internal_resource_limits_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "internal_resource_limits_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Internal-resource-limits-action"] = t
		}
	}
	if v, ok := d.GetOk("internal_resource_limits_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "internal_resource_limits_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Internal-resource-limits-block-period"] = t
		}
	}
	if v, ok := d.GetOk("internal_resource_limits_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "internal_resource_limits_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Internal-resource-limits-severity"] = t
		}
	}
	if v, ok := d.GetOk("internal_resource_limits_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "internal_resource_limits_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Internal-resource-limits-trigger"] = t
		}
	}
	if v, ok := d.GetOk("internal_resource_limits_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "internal_resource_limits_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Internal-resource-limits-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("rpc_protocol_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "rpc_protocol_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-protocol-check"] = t
		}
	}
	if v, ok := d.GetOk("rpc_protocol_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "rpc_protocol_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-protocol-action"] = t
		}
	}
	if v, ok := d.GetOk("rpc_protocol_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "rpc_protocol_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-protocol-block-period"] = t
		}
	}
	if v, ok := d.GetOk("rpc_protocol_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "rpc_protocol_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-protocol-severity"] = t
		}
	}
	if v, ok := d.GetOk("rpc_protocol_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "rpc_protocol_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-protocol-trigger"] = t
		}
	}
	if v, ok := d.GetOk("rpc_protocol_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "rpc_protocol_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rpc-protocol-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("duplicate_paramname_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "duplicate_paramname_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-paramname-check"] = t
		}
	}
	if v, ok := d.GetOk("duplicate_paramname_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "duplicate_paramname_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-paramname-action"] = t
		}
	}
	if v, ok := d.GetOk("duplicate_paramname_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "duplicate_paramname_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-paramname-block-period"] = t
		}
	}
	if v, ok := d.GetOk("duplicate_paramname_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "duplicate_paramname_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-paramname-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("duplicate_paramname_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "duplicate_paramname_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-paramname-severity"] = t
		}
	}
	if v, ok := d.GetOk("duplicate_paramname_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "duplicate_paramname_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-paramname-trigger"] = t
		}
	}
	if v, ok := d.GetOk("odd_and_even_space_attack_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "odd_and_even_space_attack_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["odd-and-even-space-attack-check"] = t
		}
	}
	if v, ok := d.GetOk("odd_and_even_space_attack_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "odd_and_even_space_attack_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["odd-and-even-space-attack-action"] = t
		}
	}
	if v, ok := d.GetOk("odd_and_even_space_attack_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "odd_and_even_space_attack_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["odd-and-even-space-attack-block-period"] = t
		}
	}
	if v, ok := d.GetOk("odd_and_even_space_attack_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "odd_and_even_space_attack_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["odd-and-even-space-attack-severity"] = t
		}
	}
	if v, ok := d.GetOk("odd_and_even_space_attack_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "odd_and_even_space_attack_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["odd-and-even-space-attack-trigger"] = t
		}
	}
	if v, ok := d.GetOk("odd_and_even_space_attack_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "odd_and_even_space_attack_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["odd-and-even-space-attack-weight"] = t
		}
	}
	if v, ok := d.GetOk("cl_te_coexist_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "cl_te_coexist_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cl-te-coexist-check"] = t
		}
	}
	if v, ok := d.GetOk("cl_te_coexist_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "cl_te_coexist_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cl-te-coexist-action"] = t
		}
	}
	if v, ok := d.GetOk("cl_te_coexist_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "cl_te_coexist_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cl-te-coexist-block-period"] = t
		}
	}
	if v, ok := d.GetOk("cl_te_coexist_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "cl_te_coexist_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cl-te-coexist-severity"] = t
		}
	}
	if v, ok := d.GetOk("cl_te_coexist_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "cl_te_coexist_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cl-te-coexist-trigger"] = t
		}
	}
	if v, ok := d.GetOk("cl_te_coexist_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "cl_te_coexist_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cl-te-coexist-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("inconsistent_cl_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "inconsistent_cl_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inconsistent-cl-check"] = t
		}
	}
	if v, ok := d.GetOk("inconsistent_cl_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "inconsistent_cl_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inconsistent-cl-action"] = t
		}
	}
	if v, ok := d.GetOk("inconsistent_cl_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "inconsistent_cl_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inconsistent-cl-block-period"] = t
		}
	}
	if v, ok := d.GetOk("inconsistent_cl_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "inconsistent_cl_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inconsistent-cl-severity"] = t
		}
	}
	if v, ok := d.GetOk("inconsistent_cl_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "inconsistent_cl_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inconsistent-cl-trigger"] = t
		}
	}
	if v, ok := d.GetOk("inconsistent_cl_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "inconsistent_cl_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inconsistent-cl-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("missing_host_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "missing_host_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-host-check"] = t
		}
	}
	if v, ok := d.GetOk("missing_host_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "missing_host_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-host-action"] = t
		}
	}
	if v, ok := d.GetOk("missing_host_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "missing_host_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-host-block-period"] = t
		}
	}
	if v, ok := d.GetOk("missing_host_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "missing_host_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-host-severity"] = t
		}
	}
	if v, ok := d.GetOk("missing_host_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "missing_host_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-host-trigger"] = t
		}
	}
	if v, ok := d.GetOk("missing_host_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "missing_host_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["missing-host-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("range_overlapping_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "range_overlapping_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-overlapping-check"] = t
		}
	}
	if v, ok := d.GetOk("range_overlapping_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "range_overlapping_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-overlapping-action"] = t
		}
	}
	if v, ok := d.GetOk("range_overlapping_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "range_overlapping_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-overlapping-block-period"] = t
		}
	}
	if v, ok := d.GetOk("range_overlapping_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "range_overlapping_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-overlapping-severity"] = t
		}
	}
	if v, ok := d.GetOk("range_overlapping_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "range_overlapping_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-overlapping-trigger"] = t
		}
	}
	if v, ok := d.GetOk("range_overlapping_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "range_overlapping_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range-overlapping-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("multipart_formdata_bad_request_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "multipart_formdata_bad_request_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipart-formdata-bad-request-check"] = t
		}
	}
	if v, ok := d.GetOk("multipart_formdata_bad_request_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "multipart_formdata_bad_request_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipart-formdata-bad-request-action"] = t
		}
	}
	if v, ok := d.GetOk("multipart_formdata_bad_request_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "multipart_formdata_bad_request_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipart-formdata-bad-request-block-period"] = t
		}
	}
	if v, ok := d.GetOk("multipart_formdata_bad_request_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "multipart_formdata_bad_request_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipart-formdata-bad-request-severity"] = t
		}
	}
	if v, ok := d.GetOk("multipart_formdata_bad_request_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "multipart_formdata_bad_request_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipart-formdata-bad-request-trigger"] = t
		}
	}
	if v, ok := d.GetOk("multipart_formdata_bad_request_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "multipart_formdata_bad_request_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multipart-formdata-bad-request-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-check"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-action"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-block-period"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-severity"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-trigger"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_freq_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_freq_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-freq-check"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_freq"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_freq", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-freq"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_freq_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_freq_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-freq-action"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_freq_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_freq_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-freq-block-period"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_freq_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_freq_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-freq-severity"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_freq_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_freq_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-freq-trigger"] = t
		}
	}
	if v, ok := d.GetOk("h2_rst_stream_freq_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h2_rst_stream_freq_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h2-rst-stream-freq-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_table_capacity_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_table_capacity_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-table-capacity-check"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_table_capacity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_table_capacity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-table-capacity"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_table_capacity_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_table_capacity_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-table-capacity-action"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_table_capacity_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_table_capacity_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-table-capacity-block-period"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_table_capacity_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_table_capacity_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-table-capacity-severity"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_table_capacity_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_table_capacity_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-table-capacity-trigger"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_table_capacity_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_table_capacity_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-table-capacity-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_field_section_size_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_field_section_size_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-field-section-size-check"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_field_section_size"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_field_section_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-field-section-size"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_field_section_size_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_field_section_size_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-field-section-size-action"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_field_section_size_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_field_section_size_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-field-section-size-block-period"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_field_section_size_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_field_section_size_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-field-section-size-severity"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_field_section_size_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_field_section_size_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-field-section-size-trigger"] = t
		}
	}
	if v, ok := d.GetOk("h3_max_field_section_size_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_max_field_section_size_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-max-field-section-size-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("h3_block_streams_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_block_streams_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-block-streams-check"] = t
		}
	}
	if v, ok := d.GetOk("h3_block_streams"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_block_streams", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-block-streams"] = t
		}
	}
	if v, ok := d.GetOk("h3_block_streams_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_block_streams_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-block-streams-action"] = t
		}
	}
	if v, ok := d.GetOk("h3_block_streams_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_block_streams_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-block-streams-block-period"] = t
		}
	}
	if v, ok := d.GetOk("h3_block_streams_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_block_streams_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-block-streams-severity"] = t
		}
	}
	if v, ok := d.GetOk("h3_block_streams_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_block_streams_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-block-streams-trigger"] = t
		}
	}
	if v, ok := d.GetOk("h3_block_streams_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_block_streams_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-block-streams-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("h3_bidir_concurrent_stream_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_bidir_concurrent_stream_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-bidir-concurrent-stream-check"] = t
		}
	}
	if v, ok := d.GetOk("h3_bidir_concurrent_stream"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_bidir_concurrent_stream", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-bidir-concurrent-stream"] = t
		}
	}
	if v, ok := d.GetOk("h3_bidir_concurrent_stream_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_bidir_concurrent_stream_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-bidir-concurrent-stream-action"] = t
		}
	}
	if v, ok := d.GetOk("h3_bidir_concurrent_stream_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_bidir_concurrent_stream_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-bidir-concurrent-stream-block-period"] = t
		}
	}
	if v, ok := d.GetOk("h3_bidir_concurrent_stream_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_bidir_concurrent_stream_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-bidir-concurrent-stream-severity"] = t
		}
	}
	if v, ok := d.GetOk("h3_bidir_concurrent_stream_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_bidir_concurrent_stream_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-bidir-concurrent-stream-trigger"] = t
		}
	}
	if v, ok := d.GetOk("h3_bidir_concurrent_stream_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_bidir_concurrent_stream_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-bidir-concurrent-stream-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("h3_unidir_concurrent_stream_check"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_unidir_concurrent_stream_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-unidir-concurrent-stream-check"] = t
		}
	}
	if v, ok := d.GetOk("h3_unidir_concurrent_stream"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_unidir_concurrent_stream", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-unidir-concurrent-stream"] = t
		}
	}
	if v, ok := d.GetOk("h3_unidir_concurrent_stream_action"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_unidir_concurrent_stream_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-unidir-concurrent-stream-action"] = t
		}
	}
	if v, ok := d.GetOk("h3_unidir_concurrent_stream_block_period"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_unidir_concurrent_stream_block_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-unidir-concurrent-stream-block-period"] = t
		}
	}
	if v, ok := d.GetOk("h3_unidir_concurrent_stream_severity"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_unidir_concurrent_stream_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-unidir-concurrent-stream-severity"] = t
		}
	}
	if v, ok := d.GetOk("h3_unidir_concurrent_stream_trigger"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_unidir_concurrent_stream_trigger", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-unidir-concurrent-stream-trigger"] = t
		}
	}
	if v, ok := d.GetOk("h3_unidir_concurrent_stream_threat_weight"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "h3_unidir_concurrent_stream_threat_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3-unidir-concurrent-stream-threat-weight"] = t
		}
	}
	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandWafHttpProtocolParameterRestriction(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	return &obj, nil
}
