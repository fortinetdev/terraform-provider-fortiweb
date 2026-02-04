---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_allow_method_exceptions_allow_method_exception_list"
description: |-
  Configure FortiWEB waf allow method exceptions allow method exception list configuration.
---

# fortiweb_waf_allow_method_exceptions_allow_method_exception_list
Configure FortiWEB waf allow method exceptions allow method exception list configuration.

## Example Usage
```hcl
resource "fortiweb_waf_allow_method_exceptions_allow_method_exception_list" "test" {
	sub_mkey = "1"
	mkey = "test"
	host_status = "enable"
	host = "test123"
	request_type = "plain"
	request_file = "/test"
	allow_request = "get post put"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf allow method exceptions.
* `sub_mkey` - The id of the waf allow method exceptions allow method exception list entry.
* `host_status` - Enable to require that the Host: field of the HTTP request match a protected hosts entry in order to match the allowed method exception.
* `host` - The name of a protected host that the Host: field of an HTTP request must be in order to match the exception.
* `request_type` - Indicate whether request-file is a literal URL (plain) or a regular expression (regular).
* `request_file` - The literal URL or a regular expression.
* `allow_request` - Select one or more of the allowed HTTP request methods that are an exception for that combination of URL and host. Valid values: get, post, head, options, trace, connect, delete, put, patch, webdav, rpc, others.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Allow Method Exceptions Allow Method Exception List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_allow_method_exceptions_allow_method_exception_list.labelname {mkey}
```
