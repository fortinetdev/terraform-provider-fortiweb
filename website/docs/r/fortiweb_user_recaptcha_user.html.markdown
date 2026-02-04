---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_recaptcha_user"
description: |-
  Configure FortiWEB user recaptcha user configuration.
---

# fortiweb_user_recaptcha_user
Configure FortiWEB user recaptcha user configuration.

## Example Usage
```hcl
resource "fortiweb_user_recaptcha_user" "test" {
	mkey = "test"
	secret_key = "test2"
	site_key = "test4"
	type = "tickbox"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user recaptcha user entry.
* `type` - Select the type of the reCAPTCHA service you have registered in Google. Valid values: tickbox, invisible, reCAPTCHA-V3.
* `site_key` - The site key.
* `secret_key` - The secret key.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Recaptcha User can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_recaptcha_user.labelname {mkey}
```
