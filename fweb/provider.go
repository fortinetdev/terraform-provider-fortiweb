// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Provider for FortiOS

package fortiweb

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider creates and returns the FortiOS *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The hostname/IP address of the FortiOS to be connected",
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Username for FortiWeb authentication",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Password for FortiWeb authentication",
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},

			"cabundlefile": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "CA Bundle file",
			},

			"cabundlecontent": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "CA Bundle file content",
			},

			"http_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "HTTP proxy address",
			},

			"https_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "HTTPS proxy address",
			},

			"no_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "NO proxy address",
			},

			"peerauth": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "disable",
				Description: "Enable/disable peer authentication, can be 'enable' or 'disable'",
			},

			"cacert": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "CA certtificate(Optional)",
			},

			"clientcert": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "User certificate",
			},

			"clientkey": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "User private key",
			},

			"vdom": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"fmg_hostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Hostname/IP address of the FortiManager to connect to",
			},

			"fmg_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"fmg_passwd": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"fmg_insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "",
			},

			"fmg_cabundlefile": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "CA Bundle file",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			// manual path
		},

		ResourcesMap: map[string]*schema.Resource{
			"fortiweb_router_policy":                                                   resourceRouterPolicy(),
			"fortiweb_router_static":                                                   resourceRouterStatic(),
			"fortiweb_server_policy_persistence_policy":                                resourceServerPolicyPersistencePolicy(),
			"fortiweb_server_policy_server_pool":                                       resourceServerPolicyServerPool(),
			"fortiweb_server_policy_server_pool_pserver_list":                          resourceServerPolicyServerPoolPserverList(),
			"fortiweb_server_policy_service_custom":                                    resourceServerPolicyServiceCustom(),
			"fortiweb_server_policy_ssl_ciphers_custom":                                resourceServerPolicySslCiphersCustom(),
			"fortiweb_server_policy_vserver":                                           resourceServerPolicyVserver(),
			"fortiweb_server_policy_vserver_vip_list":                                  resourceServerPolicyVserverVipList(),
			"fortiweb_system_accprofile":                                               resourceSystemAccprofile(),
			"fortiweb_system_admin":                                                    resourceSystemAdmin(),
			"fortiweb_system_admin_settings":                                           resourceSystemAdminSettings(),
			"fortiweb_system_certificate_ca":                                           resourceSystemCertificateCa(),
			"fortiweb_system_certificate_ca_group":                                     resourceSystemCertificateCaGroup(),
			"fortiweb_system_certificate_ca_group_members":                             resourceSystemCertificateCaGroupMembers(),
			"fortiweb_system_certificate_crl":                                          resourceSystemCertificateCrl(),
			"fortiweb_system_certificate_crl_group":                                    resourceSystemCertificateCrlGroup(),
			"fortiweb_system_certificate_crl_group_members":                            resourceSystemCertificateCrlGroupMembers(),
			"fortiweb_system_certificate_hpkp":                                         resourceSystemCertificateHpkp(),
			"fortiweb_system_certificate_intermediate_certificate":                     resourceSystemCertificateIntermediateCertificate(),
			"fortiweb_system_certificate_intermediate_certificate_group":               resourceSystemCertificateIntermediateCertificateGroup(),
			"fortiweb_system_certificate_intermediate_certificate_group_members":       resourceSystemCertificateIntermediateCertificateGroupMembers(),
			"fortiweb_system_certificate_letsencrypt":                                  resourceSystemCertificateLetsencrypt(),
			"fortiweb_system_certificate_letsencrypt_san_list":                         resourceSystemCertificateLetsencryptSanList(),
			"fortiweb_system_certificate_local_generate_csr":                           resourceSystemCertificateLocalGenerateCsr(),
			"fortiweb_system_certificate_local_import_certificate":                     resourceSystemCertificateLocalImportCertificate(),
			"fortiweb_system_certificate_multi_local":                                  resourceSystemCertificateMultiLocal(),
			"fortiweb_system_certificate_ocsp_responder":                               resourceSystemCertificateOcspResponder(),
			"fortiweb_system_certificate_ocsp_signing_certs":                           resourceSystemCertificateOcspSigningCerts(),
			"fortiweb_system_certificate_ocsp_stapling":                                resourceSystemCertificateOcspStapling(),
			"fortiweb_system_certificate_offline_sni":                                  resourceSystemCertificateOfflineSni(),
			"fortiweb_system_certificate_offline_sni_members":                          resourceSystemCertificateOfflineSniMembers(),
			"fortiweb_system_certificate_server_certificate_verify":                    resourceSystemCertificateServerCertificateVerify(),
			"fortiweb_system_certificate_sign_ca":                                      resourceSystemCertificateSignCa(),
			"fortiweb_system_certificate_sni":                                          resourceSystemCertificateSni(),
			"fortiweb_system_certificate_sni_members":                                  resourceSystemCertificateSniMembers(),
			"fortiweb_system_certificate_tsl_ca":                                       resourceSystemCertificateTslCa(),
			"fortiweb_system_certificate_urlcert":                                      resourceSystemCertificateUrlcert(),
			"fortiweb_system_certificate_urlcert_list":                                 resourceSystemCertificateUrlcertList(),
			"fortiweb_system_certificate_verify":                                       resourceSystemCertificateVerify(),
			"fortiweb_system_certificate_xml_client_certificate":                       resourceSystemCertificateXmlClientCertificate(),
			"fortiweb_system_certificate_xml_client_certificate_group":                 resourceSystemCertificateXmlClientCertificateGroup(),
			"fortiweb_system_certificate_xml_client_certificate_group_members":         resourceSystemCertificateXmlClientCertificateGroupMembers(),
			"fortiweb_system_certificate_xml_server_certificate":                       resourceSystemCertificateXmlServerCertificate(),
			"fortiweb_system_dns":                                                      resourceSystemDns(),
			"fortiweb_system_ha":                                                       resourceSystemHa(),
			"fortiweb_system_maintenance_systemtime":                                   resourceSystemMaintenanceSystemTime(),
			"fortiweb_system_ntp":                                                      resourceSystemNtp(),
			"fortiweb_system_ntp_ntpserver":                                            resourceSystemNtpNtpServer(),
			"fortiweb_system_snmp_community":                                           resourceSystemSnmpCommunity(),
			"fortiweb_system_snmp_community_hosts":                                     resourceSystemSnmpCommunityHosts(),
			"fortiweb_system_snmp_sysinfo":                                             resourceSystemSnmpSysinfo(),
			"fortiweb_system_snmp_user":                                                resourceSystemSnmpUser(),
			"fortiweb_system_snmp_user_hosts":                                          resourceSystemSnmpUserHosts(),
			"fortiweb_system_sso_admin":                                                resourceSystemSsoAdmin(),
			"fortiweb_system_vip":                                                      resourceSystemVip(),
			"fortiweb_user_admin_usergrp":                                              resourceUserAdminUsergrp(),
			"fortiweb_user_admin_usergrp_members":                                      resourceUserAdminUsergrpMembers(),
			"fortiweb_user_kerberos_user":                                              resourceUserKerberosUser(),
			"fortiweb_user_kerberos_user_server_members":                               resourceUserKerberosUserServerMembers(),
			"fortiweb_user_ldap_user":                                                  resourceUserLdapUser(),
			"fortiweb_user_ldap_user_extracted_attributes":                             resourceUserLdapUserExtractedAttributes(),
			"fortiweb_user_ntlm_user":                                                  resourceUserNtlmUser(),
			"fortiweb_user_oauth_user_server":                                          resourceUserOauthUserServer(),
			"fortiweb_user_radius_user":                                                resourceUserRadiusUser(),
			"fortiweb_user_recaptcha_user":                                             resourceUserRecaptchaUser(),
			"fortiweb_user_tacacs_user":                                                resourceUserTacacsUser(),
			"fortiweb_waf_allow_method_exceptions":                                     resourceWafAllowMethodExceptions(),
			"fortiweb_waf_allow_method_exceptions_allow_method_exception_list":         resourceWafAllowMethodExceptionsAllowMethodExceptionList(),
			"fortiweb_waf_allow_method_policy":                                         resourceWafAllowMethodPolicy(),
			"fortiweb_waf_allowed_origins":                                             resourceWafAllowedOrigins(),
			"fortiweb_waf_allowed_origins_origin_list":                                 resourceWafAllowedOriginsOriginList(),
			"fortiweb_waf_cors_protection_rule":                                        resourceWafCorsProtectionRule(),
			"fortiweb_waf_cors_protection_rule_allowed_headers_list":                   resourceWafCorsProtectionRuleAllowedHeadersList(),
			"fortiweb_waf_cors_protection_rule_allowed_methods_list":                   resourceWafCorsProtectionRuleAllowedMethodsList(),
			"fortiweb_waf_cors_protection_rule_exposed_headers_list":                   resourceWafCorsProtectionRuleExposedHeadersList(),
			"fortiweb_waf_custom_protection_group":                                     resourceWafCustomProtectionGroup(),
			"fortiweb_waf_custom_protection_group_type_list":                           resourceWafCustomProtectionGroupTypeList(),
			"fortiweb_waf_custom_protection_rule":                                      resourceWafCustomProtectionRule(),
			"fortiweb_waf_custom_protection_rule_meet_condition":                       resourceWafCustomProtectionRuleMeetCondition(),
			"fortiweb_waf_file_upload_restriction_policy":                              resourceWafFileUploadRestrictionPolicy(),
			"fortiweb_waf_file_upload_restriction_policy_rule":                         resourceWafFileUploadRestrictionPolicyRule(),
			"fortiweb_waf_file_upload_restriction_rule":                                resourceWafFileUploadRestrictionRule(),
			"fortiweb_waf_file_upload_restriction_rule_file_types":                     resourceWafFileUploadRestrictionRuleFileTypes(),
			"fortiweb_waf_http_constraints_exceptions":                                 resourceWafHttpConstraintsExceptions(),
			"fortiweb_waf_http_constraints_exceptions_http_constraints_exception_list": resourceWafHttpConstraintsExceptionsHttpConstraintsExceptionList(),
			"fortiweb_waf_http_header_security":                                        resourceWafHttpHeaderSecurity(),
			"fortiweb_waf_http_header_security_exception":                              resourceWafHttpHeaderSecurityException(),
			"fortiweb_waf_http_header_security_exception_list":                         resourceWafHttpHeaderSecurityExceptionList(),
			"fortiweb_waf_http_header_security_http_header_security_list":              resourceWafHttpHeaderSecurityHttpHeaderSecurityList(),
			"fortiweb_waf_http_protocol_parameter_restriction":                         resourceWafHttpProtocolParameterRestriction(),
			"fortiweb_waf_signature":                                                   resourceWafSignature(),
			"fortiweb_waf_url_access_url_access_parameter":                             resourceWafUrlAccessUrlAccessParameter(),
			"fortiweb_waf_url_access_url_access_parameter_url_access_parameter_list":   resourceWafUrlAccessUrlAccessParameterUrlAccessParameterList(),
			"fortiweb_waf_url_access_url_access_policy":                                resourceWafUrlAccessUrlAccessPolicy(),
			"fortiweb_waf_url_access_url_access_policy_rule":                           resourceWafUrlAccessUrlAccessPolicyRule(),
			"fortiweb_waf_url_access_url_access_rule":                                  resourceWafUrlAccessUrlAccessRule(),
			"fortiweb_waf_url_access_url_access_rule_match_condition":                  resourceWafUrlAccessUrlAccessRuleMatchCondition(),
			"fortiweb_waf_url_encryption_url_encryption_policy":                        resourceWafUrlEncryptionUrlEncryptionPolicy(),
			"fortiweb_waf_url_encryption_url_encryption_policy_rule_list":              resourceWafUrlEncryptionUrlEncryptionPolicyRuleList(),
			"fortiweb_waf_url_encryption_url_encryption_rule":                          resourceWafUrlEncryptionUrlEncryptionRule(),
			"fortiweb_waf_url_encryption_url_encryption_rule_exceptions":               resourceWafUrlEncryptionUrlEncryptionRuleExceptions(),
			"fortiweb_waf_url_encryption_url_encryption_rule_url_list":                 resourceWafUrlEncryptionUrlEncryptionRuleUrlList(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	// Init client config with the values from TF files
	config := Config{
		Hostname:        d.Get("hostname").(string),
		Token:           d.Get("token").(string),
		Username:        d.Get("username").(string),
		Password:        d.Get("password").(string),
		CABundle:        d.Get("cabundlefile").(string),
		CABundleContent: d.Get("cabundlecontent").(string),
		Vdom:            d.Get("vdom").(string),
		HTTPProxy:       d.Get("http_proxy").(string),
		HTTPSProxy:      d.Get("https_proxy").(string),
		NOProxy:         d.Get("no_proxy").(string),

		PeerAuth:   d.Get("peerauth").(string),
		CaCert:     d.Get("cacert").(string),
		ClientCert: d.Get("clientcert").(string),
		ClientKey:  d.Get("clientkey").(string),
	}

	v1, ok1 := d.GetOkExists("insecure")
	if ok1 {
		insecure := v1.(bool)
		config.Insecure = &insecure
	}

	// Create Client for later connections
	return config.CreateClient()
}
