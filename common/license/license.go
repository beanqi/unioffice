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
package license ;import _d "github.com/unidoc/unioffice/internal/license";

// SetMeteredKey sets the metered License API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey (apiKey string )error {return _d .SetMeteredKey (apiKey )};

// LicenseKey represents a loaded license key.
type LicenseKey =_d .LicenseKey ;

// SetLicenseKey sets and validates the license key.
func SetLicenseKey (content string ,customerName string )error {return _d .SetLicenseKey (content ,customerName );};

// LegacyLicense holds the old-style unioffice license information.
type LegacyLicense =_d .LegacyLicense ;const (LicenseTierUnlicensed =_d .LicenseTierUnlicensed ;LicenseTierCommunity =_d .LicenseTierCommunity ;LicenseTierIndividual =_d .LicenseTierIndividual ;LicenseTierBusiness =_d .LicenseTierBusiness ;);

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey ()*LicenseKey {return _d .GetLicenseKey ()};

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState ()(_d .MeteredStatus ,error ){return _d .GetMeteredState ()};

// LegacyLicenseType is the type of license
type LegacyLicenseType =_d .LegacyLicenseType ;

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey (s string )error {return _d .SetLegacyLicenseKey (s )};

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey ()*LicenseKey {return _d .MakeUnlicensedKey ()};