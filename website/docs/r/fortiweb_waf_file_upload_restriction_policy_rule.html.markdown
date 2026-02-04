---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_file_upload_restriction_policy_rule"
description: |-
  Configure FortiWEB waf file upload restriction policy rule configuration.
---

# fortiweb_waf_file_upload_restriction_policy_rule
Configure FortiWEB waf file upload restriction policy rule configuration.

## Example Usage
```hcl
resource "fortiweb_waf_file_upload_restriction_policy_rule" "test" {
	mkey = "test"
	sub_mkey = "1"
	file_upload_restriction_rule = "WebShell-Uploading"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf file upload restriction policy.
* `sub_mkey` - The id of the waf file upload restriction policy rule entry.
* `file_upload_restriction_rule` - The name of an upload restriction rule to use with the policy, if any.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf File Upload Restriction Policy Rule can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_file_upload_restriction_policy_rule.labelname {mkey}
```
