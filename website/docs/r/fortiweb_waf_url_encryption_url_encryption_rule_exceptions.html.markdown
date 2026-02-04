---
subcategory: "FortiWEB WAF"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_waf_url_encryption_url_encryption_rule_exceptions"
description: |-
  Configure FortiWEB waf url encryption url encryption rule exceptions configuration.
---

# fortiweb_waf_url_encryption_url_encryption_rule_exceptions
Configure FortiWEB waf url encryption url encryption rule exceptions configuration.

## Example Usage
```hcl
resource "fortiweb_waf_url_encryption_url_encryption_rule_exceptions" "test" {
	mkey = "test"
	sub_mkey = "1"
	url_pattern = "/test"
	url_type = "plain"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the waf url encryption url encryption rule.
* `sub_mkey` - The id of the waf url encryption url encryption rule exceptions list.
* `url_type` - Select whether the URL Pattern field will contain a literal URL (plain), or a regular expression designed to match multiple URLs (regular). Valid values: plain, regular.
* `url_pattern` - The literal URL or a regular expression. Depending on the url-type.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Waf Url Encryption Url Encryption Rule Exceptions can be imported using any of these accepted formats:
```
$ terraform import fortiweb_waf_url_encryption_url_encryption_rule_exceptions.labelname {mkey}
```
