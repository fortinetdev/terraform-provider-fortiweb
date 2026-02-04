---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_cors_protection_rule"
description: |-
  Configure FortiWEB waf cors protection rule configuration.
---

# fortiweb_waf_cors_protection_rule
Configure FortiWEB waf cors protection rule configuration.

## Example Usage
```hcl
resource "fortiweb_waf_cors_protection_rule" "test" {
	mkey = "test"
	host_status = "enable"
	host = "test123"
	request_type = "plain"
	request_file = "/test"
	block_cors_traffic = "disable"
	allowed_origins_list = "test"
	allowed_credentials = "false"
	allowed_maximum_age = 3
	allowed_methods = "enable"
	allowed_headers = "enable"
	exposed_headers = "enable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf cors protection rule entry.
* `host_status` - Host-Status setting.
* `host` - Host setting.
* `request_type` - Request-Type setting.
* `request_file` - Request-File setting.
* `block_cors_traffic` - Block-Cors-Traffic setting.
* `allowed_origins_list` - Allowed-Origins-List setting.
* `allowed_methods` - Allowed-Methods setting.
* `allowed_headers` - Allowed-Headers setting.
* `exposed_headers` - Exposed-Headers setting.
* `remove_other_headers` - Remove-Other-Headers setting.
* `allowed_credentials` - Allowed-Credentials setting.
* `allowed_maximum_age` - Allowed-Maximum-Age setting.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Cors Protection Rule can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_cors_protection_rule.labelname {mkey}
```
