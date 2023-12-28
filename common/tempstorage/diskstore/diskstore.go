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
package diskstore ;import (_df "github.com/unidoc/unioffice/common/tempstorage";_f "io/ioutil";_e "os";_d "strings";);

// Open opens file from disk according to a path
func (_eg diskStorage )Open (path string )(_df .File ,error ){return _e .OpenFile (path ,_e .O_RDWR ,0644)};

// Add is not applicable in the diskstore implementation
func (_a diskStorage )Add (path string )error {return nil };

// SetAsStorage sets temp storage as a disk storage
func SetAsStorage (){_dd :=diskStorage {};_df .SetAsStorage (&_dd )};

// TempFile creates a new temp file by calling ioutil TempFile
func (_dg diskStorage )TempFile (dir ,pattern string )(_df .File ,error ){return _f .TempFile (dir ,pattern );};

// RemoveAll removes all files in the directory
func (_egd diskStorage )RemoveAll (dir string )error {if _d .HasPrefix (dir ,_e .TempDir ()){return _e .RemoveAll (dir );};return nil ;};type diskStorage struct{};

// TempFile creates a new temp directory by calling ioutil TempDir
func (_da diskStorage )TempDir (pattern string )(string ,error ){return _f .TempDir ("",pattern )};