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

// Package diskstore implements tempStorage interface
// by using disk as a storage
package diskstore ;import (_dc "github.com/unidoc/unioffice/common/tempstorage";_fb "io/ioutil";_d "os";_c "strings";);

// RemoveAll removes all files in the directory
func (_a diskStorage )RemoveAll (dir string )error {if _c .HasPrefix (dir ,_d .TempDir ()){return _d .RemoveAll (dir );};return nil ;};

// Add is not applicable in the diskstore implementation
func (_b diskStorage )Add (path string )error {return nil };

// TempFile creates a new temp file by calling ioutil TempFile
func (_g diskStorage )TempFile (dir ,pattern string )(_dc .File ,error ){return _fb .TempFile (dir ,pattern );};

// TempFile creates a new temp directory by calling ioutil TempDir
func (_e diskStorage )TempDir (pattern string )(string ,error ){return _fb .TempDir ("",pattern )};type diskStorage struct{};

// Open opens file from disk according to a path
func (_df diskStorage )Open (path string )(_dc .File ,error ){return _d .OpenFile (path ,_d .O_RDWR ,0644)};

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage (){_de :=diskStorage {};_dc .SetAsStorage (&_de )};