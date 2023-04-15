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

// Package memstore implements tempStorage interface
// by using memory as a storage
package memstore ;import (_gd "encoding/hex";_e "errors";_c "fmt";_eb "github.com/unidoc/unioffice/common/tempstorage";_g "io";_ge "io/ioutil";_f "math/rand";_gb "sync";);

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage (){_be :=memStorage {_bd :_gb .Map {}};_eb .SetAsStorage (&_be )};

// TempDir creates a name for a new temp directory using a pattern argument
func (_bb *memStorage )TempDir (pattern string )(string ,error ){return _da (pattern ),nil };func _aaf (_fca int )(string ,error ){_bcf :=make ([]byte ,_fca );if _ ,_cba :=_f .Read (_bcf );_cba !=nil {return "",_cba ;};return _gd .EncodeToString (_bcf ),nil ;};

// Add reads a file from a disk and adds it to the storage
func (_fc *memStorage )Add (path string )error {_ ,_ea :=_fc ._bd .Load (path );if _ea {return nil ;};_aad ,_ga :=_ge .ReadFile (path );if _ga !=nil {return _ga ;};_fc ._bd .Store (path ,&memDataCell {_aeg :path ,_cfc :_aad ,_ff :int64 (len (_aad ))});return nil ;};func _da (_fce string )string {_bfg ,_ :=_aaf (6);return _fce +_bfg };

// ReadAt reads from the underlying memDataCell at an offset provided in order to implement ReaderAt interface.
// It does not affect f.readOffset.
func (_bfa *memFile )ReadAt (p []byte ,readOffset int64 )(int ,error ){_d :=_bfa ._fb ._ff ;_gdd :=int64 (len (p ));if _gdd > _d {_gdd =_d ;p =p [:_gdd ];};if readOffset >=_d {return 0,_g .EOF ;};_fbb :=readOffset +_gdd ;if _fbb >=_d {_fbb =_d ;};_cd :=copy (p ,_bfa ._fb ._cfc [readOffset :_fbb ]);return _cd ,nil ;};type memDataCell struct{_aeg string ;_cfc []byte ;_ff int64 ;};

// Close is not applicable in this implementation
func (_ee *memFile )Close ()error {return nil };

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_b *memFile )Read (p []byte )(int ,error ){_ae :=_b ._fe ;_bf :=_b ._fb ._ff ;_cf :=int64 (len (p ));if _cf > _bf {_cf =_bf ;p =p [:_cf ];};if _ae >=_bf {return 0,_g .EOF ;};_cb :=_ae +_cf ;if _cb >=_bf {_cb =_bf ;};_ec :=copy (p ,_b ._fb ._cfc [_ae :_cb ]);_b ._fe =_cb ;return _ec ,nil ;};

// RemoveAll removes all files according to the dir argument prefix
func (_bfaa *memStorage )RemoveAll (dir string )error {_bfaa ._bd .Range (func (_dg ,_bc interface{})bool {_bfaa ._bd .Delete (_dg );return true });return nil ;};

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_ca *memFile )Write (p []byte )(int ,error ){_ca ._fb ._cfc =append (_ca ._fb ._cfc ,p ...);_ca ._fb ._ff +=int64 (len (p ));return len (p ),nil ;};type memStorage struct{_bd _gb .Map };

// TempFile creates a new empty file in the storage and returns it
func (_afg *memStorage )TempFile (dir ,pattern string )(_eb .File ,error ){_beb :=dir +"\u002f"+_da (pattern );_cfb :=&memDataCell {_aeg :_beb ,_cfc :[]byte {}};_aa :=&memFile {_fb :_cfb };_afg ._bd .Store (_beb ,_cfb );return _aa ,nil ;};

// Open returns tempstorage File object by name
func (_fbbc *memStorage )Open (path string )(_eb .File ,error ){_af ,_bfe :=_fbbc ._bd .Load (path );if !_bfe {return nil ,_e .New (_c .Sprintf ("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",path ));};return &memFile {_fb :_af .(*memDataCell )},nil ;};type memFile struct{_fb *memDataCell ;_fe int64 ;};

// Name returns the filename of the underlying memDataCell
func (_ac *memFile )Name ()string {return _ac ._fb ._aeg };