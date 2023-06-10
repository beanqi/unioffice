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

package powerpoint ;import (_a "encoding/xml";_af "fmt";_aa "github.com/unidoc/unioffice";);

// Validate validates the CT_Rel and its children
func (_cd *CT_Rel )Validate ()error {return _cd .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func (_dda *Iscomment )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _dda .CT_Empty .MarshalXML (e ,start );};

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_ea *Iscomment )ValidateWithPath (path string )error {if _agf :=_ea .CT_Empty .ValidateWithPath (path );_agf !=nil {return _agf ;};return nil ;};func NewIscomment ()*Iscomment {_afc :=&Iscomment {};_afc .CT_Empty =*NewCT_Empty ();return _afc };func NewTextdata ()*Textdata {_ed :=&Textdata {};_ed .CT_Rel =*NewCT_Rel ();return _ed };func (_ae *CT_Rel )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {if _ae .IdAttr !=nil {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0069\u0064"},Value :_af .Sprintf ("\u0025\u0076",*_ae .IdAttr )});};e .EncodeToken (start );e .EncodeToken (_a .EndElement {Name :start .Name });return nil ;};func (_d *CT_Empty )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {for {_dd ,_ce :=d .Token ();if _ce !=nil {return _af .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_ce );};if _aad ,_bg :=_dd .(_a .EndElement );_bg &&_aad .Name ==start .Name {break ;};};return nil ;};func (_de *CT_Rel )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {for _ ,_deb :=range start .Attr {if _deb .Name .Local =="\u0069\u0064"{_bf ,_da :=_deb .Value ,error (nil );if _da !=nil {return _da ;};_de .IdAttr =&_bf ;continue ;};};for {_fg ,_e :=d .Token ();if _e !=nil {return _af .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_e );};if _bgf ,_aac :=_fg .(_a .EndElement );_aac &&_bgf .Name ==start .Name {break ;};};return nil ;};func (_cf *Textdata )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_cf .CT_Rel =*NewCT_Rel ();for _ ,_eb :=range start .Attr {if _eb .Name .Local =="\u0069\u0064"{_ef ,_cbg :=_eb .Value ,error (nil );if _cbg !=nil {return _cbg ;};_cf .IdAttr =&_ef ;continue ;};};for {_cec ,_ebc :=d .Token ();if _ebc !=nil {return _af .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_ebc );};if _aed ,_db :=_cec .(_a .EndElement );_db &&_aed .Name ==start .Name {break ;};};return nil ;};func (_b *CT_Empty )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {e .EncodeToken (start );e .EncodeToken (_a .EndElement {Name :start .Name });return nil ;};func NewCT_Empty ()*CT_Empty {_c :=&CT_Empty {};return _c };

// Validate validates the Iscomment and its children
func (_fgb *Iscomment )Validate ()error {return _fgb .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};func (_ddf *Textdata )MarshalXML (e *_a .Encoder ,start _a .StartElement )error {start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});start .Attr =append (start .Attr ,_a .Attr {Name :_a .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _ddf .CT_Rel .MarshalXML (e ,start );};

// Validate validates the Textdata and its children
func (_fe *Textdata )Validate ()error {return _fe .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};type Textdata struct{CT_Rel };type CT_Empty struct{};

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_ag *CT_Empty )ValidateWithPath (path string )error {return nil };type Iscomment struct{CT_Empty };func (_ba *Iscomment )UnmarshalXML (d *_a .Decoder ,start _a .StartElement )error {_ba .CT_Empty =*NewCT_Empty ();for {_dde ,_ca :=d .Token ();if _ca !=nil {return _af .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_ca );};if _cc ,_cab :=_dde .(_a .EndElement );_cab &&_cc .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_cfe *Textdata )ValidateWithPath (path string )error {if _dea :=_cfe .CT_Rel .ValidateWithPath (path );_dea !=nil {return _dea ;};return nil ;};type CT_Rel struct{IdAttr *string ;};

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_cb *CT_Rel )ValidateWithPath (path string )error {return nil };func NewCT_Rel ()*CT_Rel {_gf :=&CT_Rel {};return _gf };

// Validate validates the CT_Empty and its children
func (_g *CT_Empty )Validate ()error {return _g .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};func init (){_aa .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );_aa .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );_aa .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );_aa .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );};