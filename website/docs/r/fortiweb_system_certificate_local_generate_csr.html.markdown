---
subcategory: "FortiWEB System"
layout: "fortiweb"
page_title: "FortiWEB: fortiweb_system_certificate_local_generate_csr"
description: |-
  Configure FortiWEB system certificate local generate csr configuration.
---

# fortiweb_system_certificate_local_generate_csr
Configure FortiWEB system certificate local generate csr configuration.

## Example Usage
```hcl
resource "fortiweb_system_certificate_local_generate_csr" "test" {
	mkey = "test"
	idtype = "hostIp"
	ip = "0.0.0.0"
	organizationunit_1 = "orgunit"
	organization = "org"
	localitycity = "city1"
	stateprovince = "state1"
	countryregion = "US"
	email = "test@test.com"
	subject_alternative_type_1 = "domainname"
	subject_alternative_text_1 = "aaaa.com"
	keysize = "2048"
	digestalgorithm = "SHA256"
	enrollmentmethod = "file"
}

```

## Argument Reference

The following arguments are supported:

* `vdom` - Specifies the vdom to which the data source will be applied when the FortiWEB unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `mkey` - The name of the system certificate local entry.
* `idtype` - The type of identifier to use in the certificate to identify the FortiWeb appliance. Valid values: hostIp, domainName, email.
* `ip` - The static IP address of the FortiWeb appliance.
* `domainName` - The fully qualified domain name of the FortiWeb appliance.
* `subjectEmail` - The email address of the owner of the FortiWeb appliance.
* `organizationunit_1` - The name of your organizational unit, such as the name of your department.
* `organizationunit_2` - The name of your organizational unit, such as the name of your department.
* `organization` - The legal name of your organization. 
* `localitycity` - The name of the city or town where the FortiWeb appliance is located.
* `stateprovince` - The name of the state or province where the FortiWeb appliance is located.
* `countryregion` - The name of the country where the FortiWeb appliance is located.
* `email` - An email address that may be used for contact purposes.
* `subject_alternative_type_1` - The Subject Alternative type to specify additional host names. Valid values: hostIp, domainName, email.
* `subject_alternative_text_1` - The Subject Alternative Names to specify additional host names.
* `subject_alternative_type_2` - The Subject Alternative type to specify additional host names. Valid values: hostIp, domainName, email.
* `subject_alternative_text_2` - The Subject Alternative Names to specify additional host names.
* `subject_alternative_type_3` - The Subject Alternative type to specify additional host names. Valid values: hostIp, domainName, email.
* `subject_alternative_text_3` - The Subject Alternative Names to specify additional host names.
* `keysize` - Select a secure key size of 1024 Bit, 1536 Bit or 2048 Bit. Valid values: 1024, 1536, 2048.
* `digestalgorithm` - Select whether to use SHA1 or SHA256 algorithm to generate the certificate signing request. Valid values: SHA1, SHA256.
* `enrollmentmethod` - Valid values: file (File Based), scep (Online SCEP).
* `caserverurl` - Enter the CA Server URL.
* `challengepassword` - Enter the Challenge Password.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mkey}}.

## Import

System Certificate Local can be imported using any of these accepted formats:
```
$ terraform import fortiweb_system_certificate_local_generate_csr.labelname {mkey}
```
