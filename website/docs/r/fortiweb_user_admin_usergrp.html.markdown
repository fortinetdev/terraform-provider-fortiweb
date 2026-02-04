---
subcategory: "FortiWEB User"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_user_admin_usergrp"
description: |-
  Configure FortiWEB user admin usergrp configuration.
---

# fortiweb_user_admin_usergrp
Configure FortiWEB user admin usergrp configuration.

## Example Usage
```hcl
resource "fortiweb_user_admin_usergrp" "test" {
	mkey = "test"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the user admin usergrp entry.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

User Admin Usergrp can be imported using any of these accepted formats:
```
$ terraform import fortiweb_user_admin_usergrp.labelname {mkey}
```
