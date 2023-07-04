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
package memstore ;import (_ba "encoding/hex";_ad "errors";_f "fmt";_ff "github.com/unidoc/unioffice/common/tempstorage";_b "io";_cd "io/ioutil";_a "math/rand";_c "sync";);

// Name returns the filename of the underlying memDataCell
func (_fe *memFile )Name ()string {return _fe ._ac ._cdb };

// Add reads a file from a disk and adds it to the storage
func (_ca *memStorage )Add (path string )error {_ ,_bf :=_ca ._gc .Load (path );if _bf {return nil ;};_ged ,_ea :=_cd .ReadFile (path );if _ea !=nil {return _ea ;};_ca ._gc .Store (path ,&memDataCell {_cdb :path ,_aa :_ged ,_cb :int64 (len (_ged ))});return nil ;};type memStorage struct{_gc _c .Map };

// Close is not applicable in this implementation
func (_gfa *memFile )Close ()error {return nil };

// TempFile creates a new empty file in the storage and returns it
func (_ffb *memStorage )TempFile (dir ,pattern string )(_ff .File ,error ){_cf :=dir +"\u002f"+_adf (pattern );_ed :=&memDataCell {_cdb :_cf ,_aa :[]byte {}};_db :=&memFile {_ac :_ed };_ffb ._gc .Store (_cf ,_ed );return _db ,nil ;};type memDataCell struct{_cdb string ;_aa []byte ;_cb int64 ;};

// RemoveAll removes all files according to the dir argument prefix
func (_bcb *memStorage )RemoveAll (dir string )error {_bcb ._gc .Range (func (_gcg ,_ef interface{})bool {_bcb ._gc .Delete (_gcg );return true });return nil ;};

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_ae *memFile )Write (p []byte )(int ,error ){_ae ._ac ._aa =append (_ae ._ac ._aa ,p ...);_ae ._ac ._cb +=int64 (len (p ));return len (p ),nil ;};

// ReadAt reads from the underlying memDataCell at an offset provided in order to implement ReaderAt interface.
// It does not affect f.readOffset.
func (_fb *memFile )ReadAt (p []byte ,readOffset int64 )(int ,error ){_fd :=_fb ._ac ._cb ;_fg :=int64 (len (p ));if _fg > _fd {_fg =_fd ;p =p [:_fg ];};if readOffset >=_fd {return 0,_b .EOF ;};_acb :=readOffset +_fg ;if _acb >=_fd {_acb =_fd ;};_cge :=copy (p ,_fb ._ac ._aa [readOffset :_acb ]);return _cge ,nil ;};

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_d *memFile )Read (p []byte )(int ,error ){_cg :=_d ._bc ;_gf :=_d ._ac ._cb ;_e :=int64 (len (p ));if _e > _gf {_e =_gf ;p =p [:_e ];};if _cg >=_gf {return 0,_b .EOF ;};_gfe :=_cg +_e ;if _gfe >=_gf {_gfe =_gf ;};_ee :=copy (p ,_d ._ac ._aa [_cg :_gfe ]);_d ._bc =_gfe ;return _ee ,nil ;};func _adf (_fdd string )string {_adff ,_ :=_bb (6);return _fdd +_adff };func _bb (_cac int )(string ,error ){_bd :=make ([]byte ,_cac );if _ ,_baa :=_a .Read (_bd );_baa !=nil {return "",_baa ;};return _ba .EncodeToString (_bd ),nil ;};

// Open returns tempstorage File object by name
func (_fgb *memStorage )Open (path string )(_ff .File ,error ){_fdf ,_aea :=_fgb ._gc .Load (path );if !_aea {return nil ,_ad .New (_f .Sprintf ("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",path ));};return &memFile {_ac :_fdf .(*memDataCell )},nil ;};

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage (){_eb :=memStorage {_gc :_c .Map {}};_ff .SetAsStorage (&_eb )};type memFile struct{_ac *memDataCell ;_bc int64 ;};

// TempDir creates a name for a new temp directory using a pattern argument
func (_ge *memStorage )TempDir (pattern string )(string ,error ){return _adf (pattern ),nil };