---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_xml_client_certificate_group_members"
description: |-
  Configure FortiWEB system certificate xml client certificate group members configuration.
---

# fortiweb_system_certificate_xml_client_certificate_group_members
Configure FortiWEB system certificate xml client certificate group members configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_xml_client_certificate_group_members" "test" {
	mkey = "test"
	sub_mkey = "1"
	name = "sign-server"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate xml client certificate group.
* `sub_mkey` - The ID of the system certificate xml client certificate group members entry.
* `name` - The name of a client that you want to include in the group. 

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Xml Client Certificate Group Members can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_xml_client_certificate_group_members.labelname {mkey}
```
