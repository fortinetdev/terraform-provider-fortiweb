---
subcategory: "FortiWEB WAF"
layout: "fortiadc"
page_title: "FortiWEB: fortiweb_waf_http_constraints_exceptions_http_constraints_exception_list"
description: |-
  Configure FortiWEB waf http constraints exceptions http constraints exception list configuration.
---

# fortiweb_waf_http_constraints_exceptions_http_constraints_exception_list
Configure FortiWEB waf http constraints exceptions http constraints exception list configuration.

## Example Usage
```hcl
resource "fortiweb_waf_http_constraints_exceptions_http_constraints_exception_list" "test" {
	mkey = "test"
	sub_mkey = "1"
	host_status = "enable"
	host = "test"
	source_ip_status = "enable"
	source_ip = "1.2.3.4"
	request_type = "plain"
	request_file = "/test"
	max_http_content_length = "enable"
	illegal_content_length_check = "enable"
	cl_te_coexist_check = "enable"
	inconsistent_cl_check = "enable"
	max_http_header_length = "enable"
	max_http_header_name_length = "enable"
	max_http_header_value_length = "enable"
	illegal_header_name_check = "enable"
	illegal_header_value_check = "enable"
	redundant_header_check = "enable"
	max_url_parameter_length = "enable"
	max_http_body_parameter_length = "enable"
	max_url_parameter = "enable"
	parameter_name_check = "enable"
	parameter_value_check = "enable"
	max_url_param_name_len = "enable"
	max_url_param_value_len = "enable"
	url_param_name_check = "enable"
	url_param_value_check = "enable"
	duplicate_paramname_check = "enable"
	illegal_http_request_method_check = "enable"
	max_http_request_filename_length = "enable"
	max_http_request_length = "enable"
	max_header_line_request = "enable"
	post_request_ctype_check = "enable"
	missing_host_check = "enable"
	null_byte_in_url_check = "enable"
	illegal_byte_in_url_check = "enable"
	odd_and_even_space_attack_check = "enable"
	http2_max_requests = "enable"
	h2_rst_stream_check = "enable"
	h2_rst_stream_freq_check = "enable"
	h3_bidir_concurrent_stream_check = "enable"
	h3_unidir_concurrent_stream_check = "enable"
	illegal_content_type_check = "enable"
	illegal_host_name_check = "enable"
	max_http_body_length = "enable"
	max_cookie_in_request = "enable"
	number_of_ranges_in_range_header = "enable"
	block_malformed_request = "enable"
	rpc_protocol_check = "enable"
	web_socket_protocol_check = "enable"
	range_overlapping_check = "enable"
	multipart_formdata_bad_request_check = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf http constraints.
* `sub_mkey` - The ID of the waf http constraints exceptions http constraints exception list entry.
* `host_status` - Enable to apply this exception only to HTTP requests for specific web hosts. Valid values: enable, disable.
* `host` - Enter the name of specific host domain or ip address.
* `source_ip_status` - Enable to check requests for matching the HTTP constraint exceptions rule by their source IP addresses. Valid values: enable, disable.
* `source_ip` - Enter the source IP of the protected requests to which this exception applies.
* `request_type` - Enter either plain or regular (for a regular expression) to match the string entered in request-file. Valid values: plain, regular.
* `request_file` - Enter the literal URL or a regular expression.
* `max_http_content_length` - Enable to omit the constraint on the maximum HTTP content length. Valid values: enable, disable.
* `illegal_content_length_check` - Enable to omit the constraint on the maximum acceptable size in bytes of the request body. Valid values: enable, disable.
* `cl_te_coexist_check` - Enable to omit the constraint on content-length and transfer-encoding coexist. Valid values: enable, disable.
* `inconsistent_cl_check` - Enable to omit the constraint on the response has redundant body than the content-length specified. Valid values: enable, disable.
* `max_http_header_length` - Enable to omit the constraint on the maximum HTTP header length. Valid values: enable, disable.
* `max_http_header_name_length` - Enable to omit the constraint on the maximum acceptable size in bytes of a single HTTP header name. Valid values: enable, disable.
* `max_http_header_value_length` - Enable to omit the constraint on the maximum acceptable size in bytes of a single HTTP header value. Valid values: enable, disable.
* `illegal_header_name_check` - Enable to omit the constraint on whether the HTTP header name contains illegal characters. Valid values: enable, disable.
* `illegal_header_value_check` -Enable to omit the constraint on whether the HTTP header value contains illegal characters. Valid values: enable, disable.
* `redundant_header_check` - Enable to omit the constraint on the redundant instances of Content-Length, Content-Type and Host herder fields. Valid values: enable, disable.
* `max_url_parameter_length` - Enable to omit the constraint on the maximum length of parameters in the URL. Valid values: enable, disable.
* `max_http_body_parameter_length` - Enable to omit the constraint on the maximum acceptable size in bytes of all parameters in the HTTP body of HTTP POST requests. Valid values: enable, disable.
* `max_url_parameter` - Enable to omit the constraint on the maximum number of parameters in the URL. Valid values: enable, disable.
* `parameter_name_check` - Enable to omit the constraint on null characters in parameter names. Valid values: enable, disable.
* `parameter_value_check` - Enable to omit the constraint on null characters in parameter values. Valid values: enable, disable.
* `max_url_param_name_len` - Enable to omit the constraint on the maximum acceptable length in bytes of the parameter name. Valid values: enable, disable.
* `max_url_param_value_len` - Enable to omit the constraint on the maximum acceptable length in bytes of the parameter value. Valid values: enable, disable.
* `url_param_name_check` - Enable to omit the constraint on illegal characters in the parameter name. Valid values: enable, disable.
* `url_param_value_check` - Enable to omit the constraint on illegal characters in the parameter value. Valid values: enable, disable.
* `duplicate_paramname_check` - Enable to omit the constraint on duplicate parameter names. Valid values: enable, disable.
* `illegal_http_request_method_check` - Enable to omit the constraint on illegal HTTP request methods. Valid values: enable, disable.
* `max_http_request_filename_length` - Enable to omit the constraint on the maximum HTTP request filename length. Valid values: enable, disable.
* `max_http_request_length` - Enable to omit the constraint on the maximum HTTP request length. Valid values: enable, disable.
* `max_header_line_request` - Enable to omit the constraint on the maximum number of HTTP header lines. Valid values: enable, disable.
* `post_request_ctype_check` - Enable to omit the constraint on whether the Content-Type: header is available. Valid values: enable, disable.
* `missing_host_check` - Enable to omit the constraint on the Host header is missing. Valid values: enable, disable.
* `null_byte_in_url_check` - Enable to omit the constraint on null bytes in URL. Valid values: enable, disable.
* `illegal_byte_in_url_check` - Enable to omit the constraint on illegal bytes in URL. Valid values: enable, disable.
* `odd_and_even_space_attack_check` - Enable to omit the constraint on detecting Odd and Even Space Attack. Valid values: enable, disable.
* `http2_max_requests` - Enable to omit the constraint on the maximum acceptable number of requests in an HTTP/2 connection. Valid values: enable, disable.
* `h2_rst_stream_check` - Enable to omit the constraint on the maximum acceptable number of HTTP/2 RST Streams in an HTTP/2 connection. Valid values: enable, disable.
* `h2_rst_stream_freq_check` - Enable to omit the constraint on the maximum occurrences of the HTTP/2 RST Stream per second. Valid values: enable, disable.
* `h3_bidir_concurrent_stream_check` - Enable to omit the constraint on the maximum number of bidirectional concurrent streams in an HTTP/3 connection. Valid values: enable, disable.
* `h3_unidir_concurrent_stream_check` - Enable to omit the constraint on the maximum number of unidirectional concurrent streams in an HTTP/3 connection. Valid values: enable, disable.
* `illegal_content_type_check` - Enable to omit the constraint on whether the Content Type: value uses the format. Valid values: enable, disable.
* `illegal_host_name_check` - Enable to omit the constraint on host names with illegal characters. Valid values: enable, disable.
* `max_http_body_length` - Enable to omit the constraint on the maximum HTTP body length. Valid values: enable, disable.
* `max_cookie_in_request` - Enable to omit the constraint on the maximum number of cookies per request. Valid values: enable, disable.
* `number_of_ranges_in_range_header` - Enable to omit the constraint on the maximum acceptable number of Range: fields of an HTTP header. Valid values: enable, disable.
* `block_malformed_request` - Enable to omit the constraint on syntax and FortiWeb parsing errors. Valid values: enable, disable.
* `rpc_protocol_check` - Enable to omit detecting traffic that uses the PRC protocol. Valid values: enable, disable.
* `web_socket_protocol_check` - Enable to omit detecting traffic that uses the WebSocket TCP-based protocol. Valid values: enable, disable.
* `range_overlapping_check` - Enable to omit detecting RangeAmp Overlapping Byte Ranges attacks. Valid values: enable, disable.
* `multipart_formdata_bad_request_check` - Enable to omit detecting whether the multipart request chunk contains the strings "Content-Disposition" and "Name". Valid values: enable, disable.
* `internal_resource_limits_check` - Enable to omit the constraint on the maximum number of limits allowed by HTTP parser. Valid values: enable, disable.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Http Constraints Exceptions Http Constraints Exception List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_http_constraints_exceptions_http_constraints_exception_list.labelname {mkey}
```
