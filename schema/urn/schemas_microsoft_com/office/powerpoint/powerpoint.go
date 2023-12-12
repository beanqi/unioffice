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

package powerpoint ;import (_d "encoding/xml";_ad "fmt";_b "github.com/unidoc/unioffice";);

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_dgc *Iscomment )ValidateWithPath (path string )error {if _ag :=_dgc .CT_Empty .ValidateWithPath (path );_ag !=nil {return _ag ;};return nil ;};type CT_Empty struct{};type Textdata struct{CT_Rel };func (_bf *CT_Empty )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {e .EncodeToken (start );
e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func NewCT_Empty ()*CT_Empty {_c :=&CT_Empty {};return _c };func (_bb *CT_Empty )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_cc ,_g :=d .Token ();if _g !=nil {return _ad .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_g );
};if _bc ,_dd :=_cc .(_d .EndElement );_dd &&_bc .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_cbc *Textdata )ValidateWithPath (path string )error {if _db :=_cbc .CT_Rel .ValidateWithPath (path );_db !=nil {return _db ;};return nil ;};

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_ca *CT_Rel )ValidateWithPath (path string )error {return nil };func NewIscomment ()*Iscomment {_ef :=&Iscomment {};_ef .CT_Empty =*NewCT_Empty ();return _ef };

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_fd *CT_Empty )ValidateWithPath (path string )error {return nil };

// Validate validates the Textdata and its children
func (_cfc *Textdata )Validate ()error {return _cfc .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};type CT_Rel struct{IdAttr *string ;};type Iscomment struct{CT_Empty };func (_df *CT_Rel )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for _ ,_fde :=range start .Attr {if _fde .Name .Local =="\u0069\u0064"{_bfd ,_fg :=_fde .Value ,error (nil );
if _fg !=nil {return _fg ;};_df .IdAttr =&_bfd ;continue ;};};for {_e ,_gb :=d .Token ();if _gb !=nil {return _ad .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_gb );};if _gbb ,_fdd :=_e .(_d .EndElement );
_fdd &&_gbb .Name ==start .Name {break ;};};return nil ;};func NewTextdata ()*Textdata {_fa :=&Textdata {};_fa .CT_Rel =*NewCT_Rel ();return _fa };

// Validate validates the CT_Empty and its children
func (_f *CT_Empty )Validate ()error {return _f .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};func (_bed *Textdata )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _bed .CT_Rel .MarshalXML (e ,start );};func (_ace *Iscomment )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_ace .CT_Empty =*NewCT_Empty ();for {_cd ,_cb :=d .Token ();
if _cb !=nil {return _ad .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_cb );};if _dee ,_bff :=_cd .(_d .EndElement );_bff &&_dee .Name ==start .Name {break ;};};return nil ;};func (_dg *CT_Rel )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _dg .IdAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0069\u0064"},Value :_ad .Sprintf ("\u0025\u0076",*_dg .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_da *Iscomment )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _da .CT_Empty .MarshalXML (e ,start );};

// Validate validates the Iscomment and its children
func (_be *Iscomment )Validate ()error {return _be .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};func NewCT_Rel ()*CT_Rel {_ff :=&CT_Rel {};return _ff };func (_cbe *Textdata )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_cbe .CT_Rel =*NewCT_Rel ();
for _ ,_fb :=range start .Attr {if _fb .Name .Local =="\u0069\u0064"{_ba ,_ea :=_fb .Value ,error (nil );if _ea !=nil {return _ea ;};_cbe .IdAttr =&_ba ;continue ;};};for {_cf ,_gd :=d .Token ();if _gd !=nil {return _ad .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_gd );
};if _cbg ,_ae :=_cf .(_d .EndElement );_ae &&_cbg .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Rel and its children
func (_de *CT_Rel )Validate ()error {return _de .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func init (){_b .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );
_b .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );
_b .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );
_b .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );
};