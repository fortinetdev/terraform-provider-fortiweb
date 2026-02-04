---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_signature"
description: |-
  Configure FortiWEB waf signature configuration.
---

# fortiweb_waf_signature
Configure FortiWEB waf signature configuration.

## Example Usage
```hcl
resource "fortiweb_waf_signature" "test" {
	mkey = "test"
	comment = "test"
	credit_card_detection_threshold = 1
	custom_signature_group = "test"
	sensitivity_level = 4
	main_class_list {	
		action = "alert_deny"
		block_period = 600	
		id = "010000000"
		main_class_id = "010000000"
		name = "Cross Site Scripting"
		severity = "High"
		status = "enable"
		trigger = ""
		fpm_cap = 0
		fpm_status = "disable"
	}
	main_class_list {
		action = "alert"
                block_period = 600
                id = "020000000"
                main_class_id = "020000000"
                name = "Cross Site Scripting (Extended)"
                severity = "Medium"
                status = "disalbe"
                trigger = ""
     		fpm_cap = 0
		fpm_status = "disable"             
        }
	main_class_list {
		action = "alert_deny"
                block_period = 600
                id = "030000000"
                main_class_id = "030000000"
                name = "SQL Injection"
                status = "enable"
                severity = "High"
                trigger = ""
                fpm_cap = 1
		fpm_status = "enable"
        }
	main_class_list {
		action = "alert"
                block_period = 600
                id = "040000000"
                main_class_id = "040000000"
                name = "SQL Injection (Extended)"
                severity = "Medium"
                status = "disable"
                trigger = ""
                fpm_cap = 0
		fpm_status = "disable"
        }
	main_class_list {
		action = "alert_deny"
                block_period = 600
                id = "050000000"
                main_class_id = "050000000"
                name = "Generic Attacks"
                severity = "High"
                status = "enable"
                trigger = ""
                fpm_cap = 0
		fpm_status = "disable"           
        }
	main_class_list {
		action = "alert"
                block_period = 600
                id = "060000000"
                main_class_id = "060000000"
                name = "Generic Attacks(Extended)"
                severity = "Medium"
                status = "disable"
                trigger = ""
                fpm_cap = 0
		fpm_status = "disable"
        }
	main_class_list {
		action = "alert_deny"
                block_period = 600
                id = "090000000"
                main_class_id = "090000000"
                name = "Known Exploits"
                severity = "High"
                status = "enable"
                trigger = ""
                fpm_cap = 0
		fpm_status = "disable"
        }
	main_class_list {
		action = "alert"
		block_period = 600
                id = "070000000"
                main_class_id = "070000000"
                name = "Trojans"
                severity = "Medium"
                status = "enable"
                trigger = ""
                fpm_cap = 0
		fpm_status = "disable"
        }
	main_class_list {
		action = "alert"
		block_period = 600
                id = "080000000"
                main_class_id = "080000000"
                name = "Information Disclosure"
                severity = "Low"
                status = "enable"
 		trigger = ""
                fpm_cap = 0
		fpm_status = "disable"          
        }
	main_class_list {
		action = "alert"
		block_period = 600
                id = "100000000"
                main_class_id = "100000000"
                name = "Personally Identifiable Information"
                severity = "High"
                status = "enable"
		trigger = ""
		fpm_cap = 0
		fpm_status = "disable"
        }
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf signature entry.
* `comment` -  	Enter a description or other comment.
* `credit_card_detection_threshold` - Enter the number of credit cards that triggers the credit card number detection feature. Valid values: 1-128.
* `custom_protection_group` - The name of the custom signature group to be used, if any.
* `sensitivity_level` - Increasing the level adds additional signatures but also adds the chance of blocking legitimate traffic. Valid values: 1, 2, 3, 4.
* `action` - Select which action the FortiWeb appliance will take when it detects a signature match. Valid values: alert, alert_deny, block-period, only_erase, send_HTTP_response, alert_erase, redirect, deny_no_log.
* `block_period` - The number of seconds that you want to block subsequent requests from the client after the FortiWeb appliance detects that the client has violated the rule. Valid values: 1–3,600 seconds.
* `id` - The ID of a signature class (or, for subclass overrides, the subclass ID). Valid values: 010000000, 020000000, 030000000, 040000000, 050000000, 060000000, 070000000, 080000000, 090000000, 100000000, 110000000, 120000000.
* `main_class_id` - The ID of a signature class.
* `name` - The name of a parameter or cookie to match.
* `severity` - Select which severity level the FortiWeb appliance will use when it logs a violation of the rule. Valid values: Low, Medium, High, Info.
* `status` - Enable to specify whether matching requests match a specified parameter or cookie value as well as the specified parameter or cookie name.
* `flag_operation` - Flag Operation setting.
* `status` - Status setting.
* `trigger` - The name of the trigger, if any, to apply when a protection rule is violated.
* `fpm_cap` - Enter the ID of a specific signature for which false positive mitigation is disabled.
* `fpm_status` - The status of false positive mitigation.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Signature can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_signature.labelname {mkey}
```
