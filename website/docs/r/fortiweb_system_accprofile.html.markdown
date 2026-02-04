---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_accprofile"
description: |-
  Configure FortiWEB system accprofile configuration.
---

# fortiweb_system_accprofile
Configure FortiWEB system accprofile configuration.

## Example Usage
```hcl
resource "fortiweb_system_accprofile" "test" {
	mkey = "test"
	mntgrp = "rw"
	admingrp = "rw"
	sysgrp = "none"
	netgrp = "none"
	loggrp = "r"
	authusergrp = "r"
	traroutegrp = "rw"
	wafgrp = "rw"
	wadgrp = "none"
	wvsgrp = "none"
	mlgrp = "r"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system accprofile entry.
* `mntgrp` - Enter the degree of access that administrator accounts using this access profile will have to maintenance commands. Valid values: none, r, rw, w.
* `admingrp` - Enter the degree of access that administrator accounts using this access profile will have to the system administrator configuration. Valid values: none, r, rw, w.
* `sysgrp` - Enter the degree of access that administrator accounts using this access profile will have to the basic system configuration. Valid values: none, r, rw, w.
* `netgrp` - Enter the degree of access that administrator accounts using this access profile will have to the network interface and routing configuration. Valid values: none, r, rw, w.
* `loggrp` - Enter the degree of access that administrator accounts using this access profile will have to the logging and alert email configuration. Valid values: none, r, rw, w.
* `authusergrp` - Enter the degree of access that administrator accounts using this access profile will have to the HTTP authentication user configuration. Valid values: none, r, rw, w.
* `traroutegrp` - Enter the degree of access that administrator accounts using this access profile will have to the server policy configuration. Valid values: none, r, rw, w.
* `wafgrp` - Enter the degree of access that administrator accounts using this access profile will have to the without affecting traffic configuration. Valid values: none, r, rw, w.
* `wadgrp` - Enter the degree of access that administrator accounts using this access profile will have to the web anti-defacement configuration. Valid values: none, r, rw, w.
* `wvsgrp` - Enter the degree of access that administrator accounts using this access profile will have to the web protection profile configuration. Valid values: none, r, rw, w.
* `mlgrp` - Enter the degree of access that administrator accounts using this access profile will have to the machine learning configuration. Valid values: none, r, rw, w.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Accprofile can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_accprofile.labelname {mkey}
```
