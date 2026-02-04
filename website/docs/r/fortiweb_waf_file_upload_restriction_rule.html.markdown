---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_file_upload_restriction_rule"
description: |-
  Configure FortiWEB waf file upload restriction rule configuration.
---

# fortiweb_waf_file_upload_restriction_rule
Configure FortiWEB waf file upload restriction rule configuration.

## Example Usage
```hcl
resource "fortiweb_waf_file_upload_restriction_rule" "test" {
	mkey = "test"
	type = "Allow"
	host_status = "enable"
	host = "1.1.1.1"
	request_type = "plain"
	request_file = "/test"
	file_size_limit = 1024
	file_uncompress = "enable"
	json_file_support = "enable"
	json_key_for_filename = "test"
	json_key_field = "test"
	octet_stream_filename_position = "Header"
	octet_stream_filename_string = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf file upload restriction rule entry.
* `type` - Select to Allow or Block file types and custom file types.
* `host_status` - Enable to apply this exception only to HTTP requests for specific web hosts.
* `host` - Enter the name of a protected host that the Host: field of an HTTP request must be in order to match the rule.
* `request_type` - Select whether analyzer-policy will contain a literal URL (plain), or a regular expression designed to match multiple URLs (regular).
* `request_file` - The literal URL or a regular expression.
* `file_size_limit` - Optionally, enter a number to represent the maximum size in kilobytes for any individual file. Valid values: 1024, 2048, 3584.
* `file_uncompress` - Enable file unzip in CLI to verify file type and size in the compressed files.
* `json_file_support` - Enable JSON File Support if you want FortiWeb to further parse the file contained in JSON file.
* `json_key_for_filename` - FortiWeb will parse the JSON file to find the value of the filename parameter.
* `json_key_field` - FortiWeb will parse the JSON file to find the value of the content parameter.
* `octet_stream_filename_position` - Identify where to retrieve the filename of octet-stream type file. Valid values: Default, Header, Parameter, Resource.
 `octet_stream_filename_string` - Specify the header or parameter names to get the file name of octet-stream.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf File Upload Restriction Rule can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_file_upload_restriction_rule.labelname {mkey}
```
