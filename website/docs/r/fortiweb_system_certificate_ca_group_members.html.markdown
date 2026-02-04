---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_ca_group_members"
description: |-
  Configure FortiWEB system certificate ca group members configuration.
---

# fortiweb_system_certificate_ca_group_members
Configure FortiWEB system certificate ca group members configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_ca_group_members" "test1" {
	mkey = "test"
	sub_mkey = "1"
	name = "CA_Cert_1"
	publish_dn = "enable"
	type = "CA"
}

resource "fortiweb_system_certificate_ca_group_members" "test2" {
	mkey = "test"
	sub_mkey = "2"
	tsl = "test"
	type = "TSL"
	publish_dn = "disable"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate ca group.
* `sub_mkey` - The ID of the system certificate ca group members entry.
* `name` - The name of the system certificate ca.
* `type` - Select to upload CA certificate or TSL.
* `publish_dn` - Enable to list only certificates related to the specified CA Group. Valid values: enable, disable.
* `tsl` - The name of the TSL.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Ca Group Members can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_ca_group_members.labelname {mkey}
```
