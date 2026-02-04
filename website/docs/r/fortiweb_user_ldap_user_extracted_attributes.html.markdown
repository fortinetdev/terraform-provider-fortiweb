---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_ldap_user_extracted_attributes"
description: |-
  Configure FortiWEB user ldap user extracted attributes configuration.
---

# fortiweb_user_ldap_user_extracted_attributes
Configure FortiWEB user ldap user extracted attributes configuration.

## Example Usage
```hcl
resource "fortiweb_user_ldap_user_extracted_attributes" "test" {
	mkey = "test"
	sub_mkey = "1"
	name = "attribute2"
	attribute = "test2"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user ldap user.
* `sub_mkey` - The id of the user ldap user extracted attributes entry.
* `name` - Name setting.
* `attribute` - Attribute setting.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Ldap User Extracted Attributes can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_ldap_user_extracted_attributes.labelname {mkey}
```
