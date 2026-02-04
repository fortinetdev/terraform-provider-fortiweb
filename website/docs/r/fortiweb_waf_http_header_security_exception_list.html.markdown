---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_http_header_security_exception_list"
description: |-
  Configure FortiWEB waf http header security exception list configuration.
---

# fortiweb_waf_http_header_security_exception_list
Configure FortiWEB waf http header security exception list configuration.

## Example Usage
```hcl
resource "fortiweb_waf_http_header_security_exception_list" "test" {
	mkey = "test"
	sub_mkey = "1"
	client_ip = "1.1.1.1"
	client_ip_status = "enable"
	request_url_pattern = "/test.htm"
	request_url_type = "plain"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf http header security.
* `sub_mkey` - The id of the waf http header security exception list entry.
* `client_ip_status` - Enable to set Client IP exception.
* `client_ip` - Client-Ip.
* `request_url_type` - Select 'plain' (Simple String)/ 'regular' (Regular Expression) to match the URL of requests with a literal URL/ a regular expression specified in Request URL.
* `request_url_pattern` - Request-Url-Pattern.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Http Header Security Exception List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_http_header_security_exception_list.labelname {mkey}
```
