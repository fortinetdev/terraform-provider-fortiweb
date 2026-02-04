---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_allow_method_policy"
description: |-
  Configure FortiWEB waf allow method policy configuration.
---

# fortiweb_waf_allow_method_policy
Configure FortiWEB waf allow method policy configuration.

## Example Usage
```hcl
resource "fortiweb_waf_allow_method_policy" "test" {
	mkey = "test"
	override_header = "enable"
	override_parameter = "enable"
	allow_method = "get post put"
	severity = "Low"
	triggered_action = "test"
	allow_method_exception = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf allow method policy entry.
* `allow_method` - Select one or more HTTP request methods that you want to allow for this specific policy. Valid values: get, post, head, options, trace, connect, delete, put, patch, webdav, rpc.
* `override_header` - When Override Header or Override Parameter settings are enabled, FortiWeb should check methods from these headers or parameters as well as the HTTP method used in the actual request.
* `override_parameter` - When Override Header or Override Parameter settings are enabled, FortiWeb should check methods from these headers or parameters as well as the HTTP method used in the actual request.
* `severity` - Select the severity level to use in logs and reports generated when a violation of the policy occurs. Valid values: High, Medium, Low, Info.
* `triggered_action` - The name of the trigger policy you want FortiWeb to apply when a violation of the HTTP request method policy occurs.
* `allow_method_exception` - The name of an existing HTTP request method exception, if any, to apply to it. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Allow Method Policy can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_allow_method_policy.labelname {mkey}
```
