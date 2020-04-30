// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttipb

var LicenseFieldPathsNested = []string{
	"component_address_regexps",
	"components",
	"created_at",
	"dev_addr_prefixes",
	"id",
	"id.license_id",
	"join_eui_prefixes",
	"license_issuer_ids",
	"license_issuer_ids.license_issuer_id",
	"limit_for",
	"max_applications",
	"max_clients",
	"max_end_devices",
	"max_gateways",
	"max_organizations",
	"max_users",
	"max_version",
	"metering",
	"metering.interval",
	"metering.metering",
	"metering.metering.aws",
	"metering.metering.aws.sku",
	"metering.metering.prometheus",
	"metering.metering.tenant_billing_server",
	"metering.on_success",
	"metering.on_success.extend_valid_until",
	"min_version",
	"multi_tenancy",
	"valid_from",
	"valid_until",
	"warn_for",
}

var LicenseFieldPathsTopLevel = []string{
	"component_address_regexps",
	"components",
	"created_at",
	"dev_addr_prefixes",
	"id",
	"join_eui_prefixes",
	"license_issuer_ids",
	"limit_for",
	"max_applications",
	"max_clients",
	"max_end_devices",
	"max_gateways",
	"max_organizations",
	"max_users",
	"max_version",
	"metering",
	"min_version",
	"multi_tenancy",
	"valid_from",
	"valid_until",
	"warn_for",
}
var LicenseUpdateFieldPathsNested = []string{
	"extend_valid_until",
}

var LicenseUpdateFieldPathsTopLevel = []string{
	"extend_valid_until",
}
var MeteringConfigurationFieldPathsNested = []string{
	"interval",
	"metering",
	"metering.aws",
	"metering.aws.sku",
	"metering.prometheus",
	"metering.tenant_billing_server",
	"on_success",
	"on_success.extend_valid_until",
}

var MeteringConfigurationFieldPathsTopLevel = []string{
	"interval",
	"metering",
	"on_success",
}
var LicenseKeyFieldPathsNested = []string{
	"license",
	"signatures",
}

var LicenseKeyFieldPathsTopLevel = []string{
	"license",
	"signatures",
}
var MeteringDataFieldPathsNested = []string{
	"tenants",
}

var MeteringDataFieldPathsTopLevel = []string{
	"tenants",
}
var MeteringConfiguration_AWSFieldPathsNested = []string{
	"sku",
}

var MeteringConfiguration_AWSFieldPathsTopLevel = []string{
	"sku",
}
var MeteringConfiguration_PrometheusFieldPathsNested []string
var MeteringConfiguration_PrometheusFieldPathsTopLevel []string
var MeteringConfiguration_TenantBillingServerFieldPathsNested []string
var MeteringConfiguration_TenantBillingServerFieldPathsTopLevel []string
var LicenseKey_SignatureFieldPathsNested = []string{
	"key_id",
	"signature",
}

var LicenseKey_SignatureFieldPathsTopLevel = []string{
	"key_id",
	"signature",
}
var MeteringData_TenantMeteringDataFieldPathsNested = []string{
	"tenant_id",
	"tenant_id.tenant_id",
	"totals",
	"totals.applications",
	"totals.clients",
	"totals.end_devices",
	"totals.gateways",
	"totals.organizations",
	"totals.users",
}

var MeteringData_TenantMeteringDataFieldPathsTopLevel = []string{
	"tenant_id",
	"totals",
}
