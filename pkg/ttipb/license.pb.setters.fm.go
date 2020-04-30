// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttipb

import (
	fmt "fmt"
	time "time"
)

func (dst *License) SetFields(src *License, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "id":
			if len(subs) > 0 {
				var newDst, newSrc *LicenseIdentifiers
				if src != nil {
					newSrc = &src.LicenseIdentifiers
				}
				newDst = &dst.LicenseIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.LicenseIdentifiers = src.LicenseIdentifiers
				} else {
					var zero LicenseIdentifiers
					dst.LicenseIdentifiers = zero
				}
			}
		case "license_issuer_ids":
			if len(subs) > 0 {
				var newDst, newSrc *LicenseIssuerIdentifiers
				if src != nil {
					newSrc = &src.LicenseIssuerIdentifiers
				}
				newDst = &dst.LicenseIssuerIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.LicenseIssuerIdentifiers = src.LicenseIssuerIdentifiers
				} else {
					var zero LicenseIssuerIdentifiers
					dst.LicenseIssuerIdentifiers = zero
				}
			}
		case "created_at":
			if len(subs) > 0 {
				return fmt.Errorf("'created_at' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.CreatedAt = src.CreatedAt
			} else {
				var zero time.Time
				dst.CreatedAt = zero
			}
		case "valid_from":
			if len(subs) > 0 {
				return fmt.Errorf("'valid_from' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ValidFrom = src.ValidFrom
			} else {
				var zero time.Time
				dst.ValidFrom = zero
			}
		case "valid_until":
			if len(subs) > 0 {
				return fmt.Errorf("'valid_until' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ValidUntil = src.ValidUntil
			} else {
				var zero time.Time
				dst.ValidUntil = zero
			}
		case "warn_for":
			if len(subs) > 0 {
				return fmt.Errorf("'warn_for' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.WarnFor = src.WarnFor
			} else {
				var zero time.Duration
				dst.WarnFor = zero
			}
		case "limit_for":
			if len(subs) > 0 {
				return fmt.Errorf("'limit_for' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.LimitFor = src.LimitFor
			} else {
				var zero time.Duration
				dst.LimitFor = zero
			}
		case "min_version":
			if len(subs) > 0 {
				return fmt.Errorf("'min_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MinVersion = src.MinVersion
			} else {
				var zero string
				dst.MinVersion = zero
			}
		case "max_version":
			if len(subs) > 0 {
				return fmt.Errorf("'max_version' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxVersion = src.MaxVersion
			} else {
				var zero string
				dst.MaxVersion = zero
			}
		case "components":
			if len(subs) > 0 {
				return fmt.Errorf("'components' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Components = src.Components
			} else {
				dst.Components = nil
			}
		case "component_address_regexps":
			if len(subs) > 0 {
				return fmt.Errorf("'component_address_regexps' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ComponentAddressRegexps = src.ComponentAddressRegexps
			} else {
				dst.ComponentAddressRegexps = nil
			}
		case "dev_addr_prefixes":
			if len(subs) > 0 {
				return fmt.Errorf("'dev_addr_prefixes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.DevAddrPrefixes = src.DevAddrPrefixes
			} else {
				dst.DevAddrPrefixes = nil
			}
		case "join_eui_prefixes":
			if len(subs) > 0 {
				return fmt.Errorf("'join_eui_prefixes' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.JoinEUIPrefixes = src.JoinEUIPrefixes
			} else {
				dst.JoinEUIPrefixes = nil
			}
		case "multi_tenancy":
			if len(subs) > 0 {
				return fmt.Errorf("'multi_tenancy' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MultiTenancy = src.MultiTenancy
			} else {
				var zero bool
				dst.MultiTenancy = zero
			}
		case "max_applications":
			if len(subs) > 0 {
				return fmt.Errorf("'max_applications' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxApplications = src.MaxApplications
			} else {
				dst.MaxApplications = nil
			}
		case "max_clients":
			if len(subs) > 0 {
				return fmt.Errorf("'max_clients' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxClients = src.MaxClients
			} else {
				dst.MaxClients = nil
			}
		case "max_end_devices":
			if len(subs) > 0 {
				return fmt.Errorf("'max_end_devices' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxEndDevices = src.MaxEndDevices
			} else {
				dst.MaxEndDevices = nil
			}
		case "max_gateways":
			if len(subs) > 0 {
				return fmt.Errorf("'max_gateways' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxGateways = src.MaxGateways
			} else {
				dst.MaxGateways = nil
			}
		case "max_organizations":
			if len(subs) > 0 {
				return fmt.Errorf("'max_organizations' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxOrganizations = src.MaxOrganizations
			} else {
				dst.MaxOrganizations = nil
			}
		case "max_users":
			if len(subs) > 0 {
				return fmt.Errorf("'max_users' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.MaxUsers = src.MaxUsers
			} else {
				dst.MaxUsers = nil
			}
		case "metering":
			if len(subs) > 0 {
				var newDst, newSrc *MeteringConfiguration
				if (src == nil || src.Metering == nil) && dst.Metering == nil {
					continue
				}
				if src != nil {
					newSrc = src.Metering
				}
				if dst.Metering != nil {
					newDst = dst.Metering
				} else {
					newDst = &MeteringConfiguration{}
					dst.Metering = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Metering = src.Metering
				} else {
					dst.Metering = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *LicenseUpdate) SetFields(src *LicenseUpdate, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "extend_valid_until":
			if len(subs) > 0 {
				return fmt.Errorf("'extend_valid_until' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.ExtendValidUntil = src.ExtendValidUntil
			} else {
				dst.ExtendValidUntil = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *MeteringConfiguration) SetFields(src *MeteringConfiguration, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "interval":
			if len(subs) > 0 {
				return fmt.Errorf("'interval' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Interval = src.Interval
			} else {
				var zero time.Duration
				dst.Interval = zero
			}
		case "on_success":
			if len(subs) > 0 {
				var newDst, newSrc *LicenseUpdate
				if (src == nil || src.OnSuccess == nil) && dst.OnSuccess == nil {
					continue
				}
				if src != nil {
					newSrc = src.OnSuccess
				}
				if dst.OnSuccess != nil {
					newDst = dst.OnSuccess
				} else {
					newDst = &LicenseUpdate{}
					dst.OnSuccess = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.OnSuccess = src.OnSuccess
				} else {
					dst.OnSuccess = nil
				}
			}

		case "metering":
			if len(subs) == 0 && src == nil {
				dst.Metering = nil
				continue
			} else if len(subs) == 0 {
				dst.Metering = src.Metering
				continue
			}

			subPathMap := _processPaths(subs)
			if len(subPathMap) > 1 {
				return fmt.Errorf("more than one field specified for oneof field '%s'", name)
			}
			for oneofName, oneofSubs := range subPathMap {
				switch oneofName {
				case "aws":
					_, srcOk := src.Metering.(*MeteringConfiguration_AWS_)
					if !srcOk && src.Metering != nil {
						return fmt.Errorf("attempt to set oneof 'aws', while different oneof is set in source")
					}
					_, dstOk := dst.Metering.(*MeteringConfiguration_AWS_)
					if !dstOk && dst.Metering != nil {
						return fmt.Errorf("attempt to set oneof 'aws', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *MeteringConfiguration_AWS
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Metering.(*MeteringConfiguration_AWS_).AWS
						}
						if dstOk {
							newDst = dst.Metering.(*MeteringConfiguration_AWS_).AWS
						} else {
							newDst = &MeteringConfiguration_AWS{}
							dst.Metering = &MeteringConfiguration_AWS_{AWS: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Metering = src.Metering
						} else {
							dst.Metering = nil
						}
					}
				case "prometheus":
					_, srcOk := src.Metering.(*MeteringConfiguration_Prometheus_)
					if !srcOk && src.Metering != nil {
						return fmt.Errorf("attempt to set oneof 'prometheus', while different oneof is set in source")
					}
					_, dstOk := dst.Metering.(*MeteringConfiguration_Prometheus_)
					if !dstOk && dst.Metering != nil {
						return fmt.Errorf("attempt to set oneof 'prometheus', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *MeteringConfiguration_Prometheus
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Metering.(*MeteringConfiguration_Prometheus_).Prometheus
						}
						if dstOk {
							newDst = dst.Metering.(*MeteringConfiguration_Prometheus_).Prometheus
						} else {
							newDst = &MeteringConfiguration_Prometheus{}
							dst.Metering = &MeteringConfiguration_Prometheus_{Prometheus: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Metering = src.Metering
						} else {
							dst.Metering = nil
						}
					}
				case "tenant_billing_server":
					_, srcOk := src.Metering.(*MeteringConfiguration_TenantBillingServer_)
					if !srcOk && src.Metering != nil {
						return fmt.Errorf("attempt to set oneof 'tenant_billing_server', while different oneof is set in source")
					}
					_, dstOk := dst.Metering.(*MeteringConfiguration_TenantBillingServer_)
					if !dstOk && dst.Metering != nil {
						return fmt.Errorf("attempt to set oneof 'tenant_billing_server', while different oneof is set in destination")
					}
					if len(oneofSubs) > 0 {
						var newDst, newSrc *MeteringConfiguration_TenantBillingServer
						if !srcOk && !dstOk {
							continue
						}
						if srcOk {
							newSrc = src.Metering.(*MeteringConfiguration_TenantBillingServer_).TenantBillingServer
						}
						if dstOk {
							newDst = dst.Metering.(*MeteringConfiguration_TenantBillingServer_).TenantBillingServer
						} else {
							newDst = &MeteringConfiguration_TenantBillingServer{}
							dst.Metering = &MeteringConfiguration_TenantBillingServer_{TenantBillingServer: newDst}
						}
						if err := newDst.SetFields(newSrc, oneofSubs...); err != nil {
							return err
						}
					} else {
						if src != nil {
							dst.Metering = src.Metering
						} else {
							dst.Metering = nil
						}
					}

				default:
					return fmt.Errorf("invalid oneof field: '%s.%s'", name, oneofName)
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *LicenseKey) SetFields(src *LicenseKey, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "license":
			if len(subs) > 0 {
				return fmt.Errorf("'license' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.License = src.License
			} else {
				dst.License = nil
			}
		case "signatures":
			if len(subs) > 0 {
				return fmt.Errorf("'signatures' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Signatures = src.Signatures
			} else {
				dst.Signatures = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *MeteringData) SetFields(src *MeteringData, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "tenants":
			if len(subs) > 0 {
				return fmt.Errorf("'tenants' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Tenants = src.Tenants
			} else {
				dst.Tenants = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *MeteringConfiguration_AWS) SetFields(src *MeteringConfiguration_AWS, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "sku":
			if len(subs) > 0 {
				return fmt.Errorf("'sku' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.SKU = src.SKU
			} else {
				var zero string
				dst.SKU = zero
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *MeteringConfiguration_Prometheus) SetFields(src *MeteringConfiguration_Prometheus, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message MeteringConfiguration_Prometheus has no fields, but paths %s were specified", paths)
	}
	if src != nil {
		*dst = *src
	}
	return nil
}

func (dst *MeteringConfiguration_TenantBillingServer) SetFields(src *MeteringConfiguration_TenantBillingServer, paths ...string) error {
	if len(paths) != 0 {
		return fmt.Errorf("message MeteringConfiguration_TenantBillingServer has no fields, but paths %s were specified", paths)
	}
	if src != nil {
		*dst = *src
	}
	return nil
}

func (dst *LicenseKey_Signature) SetFields(src *LicenseKey_Signature, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "key_id":
			if len(subs) > 0 {
				return fmt.Errorf("'key_id' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.KeyID = src.KeyID
			} else {
				var zero string
				dst.KeyID = zero
			}
		case "signature":
			if len(subs) > 0 {
				return fmt.Errorf("'signature' has no subfields, but %s were specified", subs)
			}
			if src != nil {
				dst.Signature = src.Signature
			} else {
				dst.Signature = nil
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}

func (dst *MeteringData_TenantMeteringData) SetFields(src *MeteringData_TenantMeteringData, paths ...string) error {
	for name, subs := range _processPaths(paths) {
		switch name {
		case "tenant_id":
			if len(subs) > 0 {
				var newDst, newSrc *TenantIdentifiers
				if src != nil {
					newSrc = &src.TenantIdentifiers
				}
				newDst = &dst.TenantIdentifiers
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.TenantIdentifiers = src.TenantIdentifiers
				} else {
					var zero TenantIdentifiers
					dst.TenantIdentifiers = zero
				}
			}
		case "totals":
			if len(subs) > 0 {
				var newDst, newSrc *TenantRegistryTotals
				if (src == nil || src.Totals == nil) && dst.Totals == nil {
					continue
				}
				if src != nil {
					newSrc = src.Totals
				}
				if dst.Totals != nil {
					newDst = dst.Totals
				} else {
					newDst = &TenantRegistryTotals{}
					dst.Totals = newDst
				}
				if err := newDst.SetFields(newSrc, subs...); err != nil {
					return err
				}
			} else {
				if src != nil {
					dst.Totals = src.Totals
				} else {
					dst.Totals = nil
				}
			}

		default:
			return fmt.Errorf("invalid field: '%s'", name)
		}
	}
	return nil
}
