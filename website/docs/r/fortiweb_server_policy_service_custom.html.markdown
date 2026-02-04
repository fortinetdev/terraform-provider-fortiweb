---
subcategory: "FortiWEB Server Policy"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_server_policy_service_custom"
description: |-
  Configure FortiWEB server policy service custom configuration.
---

# fortiweb_server_policy_service_custom
Configure FortiWEB server policy service custom configuration.

## Example Usage
```hcl
resource "fortiweb_server_policy_service_custom" "test" {
	mkey = "test"
	port = "443"
	protocol = "TCP"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the server policy service custom entry.
* `port` - Enter the port number on which a virtual server will receive TCP/IP connections for HTTP or HTTPS requests. Valid values: 1–65535.
* `protocol` - TCP.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

Server Policy Service Custom can be imported using any of these accepted formats:
```
$ terraform import fortiweb_server_policy_service_custom.labelname {mkey}
```
