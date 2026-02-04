---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_custom_protection_rule"
description: |-
  Configure FortiWEB waf custom protection rule configuration.
---

# fortiweb_waf_custom_protection_rule
Configure FortiWEB waf custom protection rule configuration.

## Example Usage
```hcl
resource "fortiweb_waf_custom_protection_rule" "test" {
	mkey = "test1"
	type = "request"
	action = "alert"
	block_period = 1000
	severity = "Medium"
	trigger = "test"
	threat_weight = "substantial"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf custom protection rule entry.
* `type` - Specify the type of regular expression. Valid values: request, response.
* `action` - Select the specific action to be taken when the request matches the this signature. Valid values: alert, alert_deny, alert_erase, redirect, block-period, send_HTTP_response, only_erase, deny_no_log.
* `block_period` - The number of seconds that you want to block subsequent requests from the client after the FortiWeb appliance detects that the client has violated the rule. Valid values: 1–3600 seconds.
* `severity` - Select which severity level the FortiWeb appliance will use when it logs a violation of the rule. Valid values: 1–3600 seconds. Valid values: High, Medium, Low, Info.
* `trigger` - Select which trigger policy, if any, that the FortiWeb appliance will use when it logs and/or sends an alert email about a violation of the rule.
* `threat_weight` - Set the threat weight. Valid values: low, critical, informational, moderate, substantial, severe.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Custom Protection Rule can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_custom_protection_rule.labelname {mkey}
```
