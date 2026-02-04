---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_kerberos_user"
description: |-
  Configure FortiWEB user kerberos user configuration.
---

# fortiweb_user_kerberos_user
Configure FortiWEB user kerberos user configuration.

## Example Usage
```hcl
resource "fortiweb_user_kerberos_user" "test" {
	mkey = "test"
	realm = "test3"
	shortname = "testname"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user kerberos user entry.
* `realm` - The domain of the domain controller (DC) that the Key Distribution Center (KDC) belongs to.
* `shortname` - The shortname for the realm you specified (This is optional). 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Kerberos User can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_kerberos_user.labelname {mkey}
```
