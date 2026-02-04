---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_url_encryption_url_encryption_rule"
description: |-
  Configure FortiWEB waf url encryption url encryption rule configuration.
---

# fortiweb_waf_url_encryption_url_encryption_rule
Configure FortiWEB waf url encryption url encryption rule configuration.

## Example Usage
```hcl
resource "fortiweb_waf_url_encryption_url_encryption_rule" "test" {
	mkey = "test"
	host_status = "enable"
	host = "test"
	allow_unencrypted = "disable"
	action = "alert"
	block_period = 600
	severity = "High"
	trigger = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf url encryption url encryption rule entry.
* `host_status` - Enable to require that the Host: field of the HTTP request match a protected host names entry in order to match the URL acceleration rule. Valid values: enable, disable.
* `host` - Select which protected host names entry (either a web host name or IP address) that the Host: field of the HTTP request must be in to match the URL acceleration rule.
* `allow_unencrypted` - Unencrypted URL requests are the valid requests from the client that FortiWeb failed to decrypt. Valid values: enable, disable.
* `action` - Select which action the FortiWeb appliance will take when it detects a violation. Valid values: alert, alert_deny, deny_no_log or block-period.
* `block_period` - The number of seconds that you want to block the requests. Valid values: 1-3600 seconds.
* `severity` - The severity level that FortiWeb will record when the rule is violated. Valid values: Low, Medium, High and Informative.
* `trigger` - Select the trigger, if any, that FortiWeb carries out when it logs and/or sends an alert email about a rule violation. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Url Encryption Url Encryption Rule can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_url_encryption_url_encryption_rule.labelname {mkey}
```
