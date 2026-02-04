---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_url_access_url_access_rule"
description: |-
  Configure FortiWEB waf url access url access rule configuration.
---

# fortiweb_waf_url_access_url_access_rule
Configure FortiWEB waf url access url access rule configuration.

## Example Usage
```hcl
resource "fortiweb_waf_url_access_url_access_rule" "test" {
	mkey = "test"
	host_status = "enable"
	host = "test"
	action = "alert_deny"
	severity = "Low"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf url access url access rule entry.
* `host_status` - Enable to require that the Host: field of the HTTP request match a protected hosts entry in order to match the rule. Valid values: enable, disable.
* `host` - The name of a protected host that the Host: field of an HTTP request must be in order to match the rule.
* `action` - Select which action the FortiWeb appliance will take when a request matches the URL access rule. Valid values: alert_deny, continue, pass, deny_no_log.
* `severity` - Select which severity level the FortiWeb appliance will use when a blocklisted IP address attempts to connect to your web servers. Valid values: Informative, Low, Medium, High, Info.
* `trigger` - Select which trigger.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Url Access Url Access Rule can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_url_access_url_access_rule.labelname {mkey}
```
