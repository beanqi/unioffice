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
package diskstore ;import (_aa "github.com/unidoc/unioffice/common/tempstorage";_e "io/ioutil";_a "os";_b "strings";);

// RemoveAll removes all files in the directory
func (_d diskStorage )RemoveAll (dir string )error {if _b .HasPrefix (dir ,_a .TempDir ()){return _a .RemoveAll (dir );};return nil ;};

// Add is not applicable in the diskstore implementation
func (_bga diskStorage )Add (path string )error {return nil };

// TempFile creates a new temp file by calling ioutil TempFile
func (_ce diskStorage )TempFile (dir ,pattern string )(_aa .File ,error ){return _e .TempFile (dir ,pattern );};

// TempFile creates a new temp directory by calling ioutil TempDir
func (_bg diskStorage )TempDir (pattern string )(string ,error ){return _e .TempDir ("",pattern )};type diskStorage struct{};

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage (){_f :=diskStorage {};_aa .SetAsStorage (&_f )};

// Open opens file from disk according to a path
func (_g diskStorage )Open (path string )(_aa .File ,error ){return _a .OpenFile (path ,_a .O_RDWR ,0644)};