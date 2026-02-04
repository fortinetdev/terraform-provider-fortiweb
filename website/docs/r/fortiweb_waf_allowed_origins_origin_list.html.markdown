---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_allowed_origins_origin_list"
description: |-
  Configure FortiWEB waf allowed origins origin list configuration.
---

# fortiweb_waf_allowed_origins_origin_list
Configure FortiWEB waf allowed origins origin list configuration.

## Example Usage
```hcl
resource "fortiweb_waf_allowed_origins_origin_list" "test1" {
	sub_mkey = "1"
	protocol = "HTTPS"
	origin_name = "testname"
	port = 80
	include_sub_domains = "enable"
	mkey = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf allowed origins.
* `sub_mkey` - The id of the waf allowed origins origin list entry.
* `protocol` - Select the HTTP protocols.
* `origin_name` - Origin-Name setting.
* `port` - Port setting.
* `include_sub_domains` - Include-Sub-Domains setting.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Allowed Origins Origin List can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_allowed_origins_origin_list.labelname {mkey}
```
