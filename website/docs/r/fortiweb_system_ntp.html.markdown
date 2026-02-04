---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_ntp"
description: |-
  Configure FortiWEB system ntp configuration.
---

# fortiweb_system_ntp
Configure FortiWEB system ntp configuration.

## Example Usage
```hcl
resource "fortiweb_system_ntp" "test" {
	ntpsync = "enable"
	syncinterval = 90
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `ntpsync` - Enable/disable use of NTP.
* `syncinterval` - Specify how often the system synchronizes its time with the NTP server. Valid values: 1-1440.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Ntp can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_ntp.labelname {mkey}
```
