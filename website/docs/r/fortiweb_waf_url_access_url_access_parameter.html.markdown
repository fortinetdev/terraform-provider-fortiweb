---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_url_access_url_access_parameter"
description: |-
  Configure FortiWEB waf url access url access parameter configuration.
---

# fortiweb_waf_url_access_url_access_parameter
Configure FortiWEB waf url access url access parameter configuration.

## Example Usage
```hcl
resource "fortiweb_waf_url_access_url_access_parameter" "test" {
	mkey = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf url access url access parameter entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Url Access Url Access Parameter can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_url_access_url_access_parameter.labelname {mkey}
```
