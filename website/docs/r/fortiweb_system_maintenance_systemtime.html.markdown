---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_maintenance_systemtime"
description: |-
  Configure FortiWEB system maintenance systemtime configuration.
---

# fortiweb_system_maintenance_systemtime
Configure FortiWEB system maintenance systemtime configuration.

## Example Usage
```hcl
resource "fortiweb_system_maintenance_systemtime" "test" {
	daylightsaving = 0
	timezone = 4
	mode = "ntpServer"
	systemtime = "Mon Dec 8 01:31:25 2025"
	time = "12/8/2025 1:31:25"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `daylightsaving` - Enable daylight saving time. Valid value: 0-disable, 1-enable. 
* `timezone` - The index number of the time zone in which the FortiWeb appliance is located.
* `mode` - Valid value: ntpServer, setTime.
* `systemtime` - The time for the time zone in which the FortiWeb appliance is located according to a 24-hour clock, formatted as hh:mm:ss.
* `time` - The date for the time zone in which the FortiWeb appliance is located, formatted as yyyy-mm-dd.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Maintenance Systemtime can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_maintenance_systemtime.labelname {mkey}
```
