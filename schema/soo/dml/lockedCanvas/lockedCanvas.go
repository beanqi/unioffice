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

package lockedCanvas ;import (_d "encoding/xml";_fa "fmt";_b "github.com/unidoc/unioffice";_c "github.com/unidoc/unioffice/schema/soo/dml";);

// ValidateWithPath validates the LockedCanvas and its children, prefixing error messages with path
func (_gf *LockedCanvas )ValidateWithPath (path string )error {if _ec :=_gf .CT_GvmlGroupShape .ValidateWithPath (path );_ec !=nil {return _ec ;};return nil ;};func (_ga *LockedCanvas )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006c\u006f\u0063\u006b\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073"});
start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u006c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073";return _ga .CT_GvmlGroupShape .MarshalXML (e ,start );};func (_cf *LockedCanvas )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_cf .CT_GvmlGroupShape =*_c .NewCT_GvmlGroupShape ();
for {_e ,_faa :=d .Token ();if _faa !=nil {return _fa .Errorf ("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u004c\u006f\u0063\u006b\u0065d\u0043\u0061\u006e\u0076\u0061\u0073\u003a\u0020\u0025\u0073",_faa );};if _ea ,_a :=_e .(_d .EndElement );_a &&_ea .Name ==start .Name {break ;
};};return nil ;};type LockedCanvas struct{_c .CT_GvmlGroupShape };

// Validate validates the LockedCanvas and its children
func (_ef *LockedCanvas )Validate ()error {return _ef .ValidateWithPath ("\u004c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073");};func NewLockedCanvas ()*LockedCanvas {_g :=&LockedCanvas {};_g .CT_GvmlGroupShape =*_c .NewCT_GvmlGroupShape ();
return _g ;};func init (){_b .RegisterConstructor ("h\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061t\u0073.\u006f\u0072\u0067\u002fd\u0072\u0061w\u0069\u006e\u0067\u006d\u006c\u002f\u0032\u0030\u0030\u0036\u002f\u006c\u006f\u0063\u006b\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073","\u006c\u006f\u0063k\u0065\u0064\u0043\u0061\u006e\u0076\u0061\u0073",NewLockedCanvas );
};