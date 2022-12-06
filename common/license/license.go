//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package license helps manage commercial licenses and check if they
// are valid for the version of UniOffice used.
package license ;import _c "github.com/unidoc/unioffice/internal/license";

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey ()*LicenseKey {return _c .MakeUnlicensedKey ()};

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey (s string )error {return _c .SetLegacyLicenseKey (s )};

// LicenseKey represents a loaded license key.
type LicenseKey =_c .LicenseKey ;

// LegacyLicense holds the old-style unioffice license information.
type LegacyLicense =_c .LegacyLicense ;const (LicenseTierUnlicensed =_c .LicenseTierUnlicensed ;LicenseTierCommunity =_c .LicenseTierCommunity ;LicenseTierIndividual =_c .LicenseTierIndividual ;LicenseTierBusiness =_c .LicenseTierBusiness ;);

// LegacyLicenseType is the type of license
type LegacyLicenseType =_c .LegacyLicenseType ;

// SetMeteredKey sets the metered License API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey (apiKey string )error {return _c .SetMeteredKey (apiKey )};

// SetLicenseKey sets and validates the license key.
func SetLicenseKey (content string ,customerName string )error {return _c .SetLicenseKey (content ,customerName );};

// SetMeteredKeyPersistentCache sets the metered License API Key persistent cache.
// Default value `true`, set to `false` will report the usage immediately to license server,
// this can be used when there's no access to persistent data storage.
func SetMeteredKeyPersistentCache (val bool ){_c .SetMeteredKeyPersistentCache (val )};

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState ()(_c .MeteredStatus ,error ){return _c .GetMeteredState ()};

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey ()*LicenseKey {return _c .GetLicenseKey ()};