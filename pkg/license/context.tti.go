// Copyright © 2019 The Things Industries B.V.

package license

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"go.thethings.network/lorawan-stack/pkg/ttipb"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
)

type licenseContextKeyType struct{}

var licenseContextKey licenseContextKeyType

var now = time.Now()

var defaultLicense = ttipb.License{
	LicenseIdentifiers:      ttipb.LicenseIdentifiers{LicenseID: "development"},
	CreatedAt:               now,
	ValidFrom:               now,
	ValidUntil:              now.Add(time.Hour),
	ComponentAddressRegexps: []string{"localhost"},
	DevAddrPrefixes: []types.DevAddrPrefix{
		{DevAddr: types.DevAddr{0, 0, 0, 0}, Length: 7},
		{DevAddr: types.DevAddr{0, 0, 0, 2}, Length: 7},
	},
}

// NewContextWithLicense returns a context derived from parent that contains license.
func NewContextWithLicense(parent context.Context, license ttipb.License) context.Context {
	return context.WithValue(parent, licenseContextKey, license)
}

// FromContext returns the License from the context if present. Otherwise it returns default License.
func FromContext(ctx context.Context) ttipb.License {
	if license, ok := ctx.Value(licenseContextKey).(ttipb.License); ok {
		if license.Metering != nil {
			license = globalMetering.Apply(license)
		}
		return license
	}
	return defaultLicense
}

// RequireComponent requires components to be included in the license.
func RequireComponent(ctx context.Context, components ...ttnpb.ClusterRole) error {
	license := FromContext(ctx)
	if err := checkValidity(&license); err != nil {
		return err
	}
	if len(license.Components) == 0 {
		return nil
	}
nextComponent:
	for _, component := range components {
		for _, licensed := range license.Components {
			if component == licensed {
				continue nextComponent
			}
		}
		return fmt.Errorf("the %s component is not included in this license", strings.Title(strings.Replace(component.String(), "_", " ", -1)))
	}
	return nil
}

// RequireMultiTenancy requires multi-tenancy to be included in the license.
func RequireMultiTenancy(ctx context.Context) error {
	license := FromContext(ctx)
	if err := checkValidity(&license); err != nil {
		return err
	}
	if !license.MultiTenancy {
		return errors.New("multi-tenancy is not included in this license")
	}
	return nil
}
