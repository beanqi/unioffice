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

package powerpoint ;import (_f "encoding/xml";_e "fmt";_d "github.com/unidoc/unioffice";);func (_ag *CT_Rel )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {for _ ,_gc :=range start .Attr {if _gc .Name .Local =="\u0069\u0064"{_ecf ,_ece :=_gc .Value ,error (nil );
if _ece !=nil {return _ece ;};_ag .IdAttr =&_ecf ;continue ;};};for {_cc ,_ad :=d .Token ();if _ad !=nil {return _e .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_ad );};if _af ,_eag :=_cc .(_f .EndElement );
_eag &&_af .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Empty and its children
func (_fgg *CT_Empty )Validate ()error {return _fgg .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};func (_ea *CT_Rel )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {if _ea .IdAttr !=nil {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0069\u0064"},Value :_e .Sprintf ("\u0025\u0076",*_ea .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};func NewCT_Empty ()*CT_Empty {_g :=&CT_Empty {};return _g };func (_dda *Iscomment )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_dda .CT_Empty =*NewCT_Empty ();
for {_ecfe ,_ac :=d .Token ();if _ac !=nil {return _e .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_ac );};if _fb ,_gd :=_ecfe .(_f .EndElement );_gd &&_fb .Name ==start .Name {break ;
};};return nil ;};

// Validate validates the Iscomment and its children
func (_bg *Iscomment )Validate ()error {return _bg .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_da *CT_Rel )ValidateWithPath (path string )error {return nil };func (_ef *CT_Empty )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {e .EncodeToken (start );e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};type CT_Rel struct{IdAttr *string ;
};type Iscomment struct{CT_Empty };func (_cb *Textdata )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _cb .CT_Rel .MarshalXML (e ,start );};func (_ecea *Iscomment )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});
start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _ecea .CT_Empty .MarshalXML (e ,start );};

// Validate validates the Textdata and its children
func (_ga *Textdata )Validate ()error {return _ga .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_ae *CT_Empty )ValidateWithPath (path string )error {return nil };func NewTextdata ()*Textdata {_fdf :=&Textdata {};_fdf .CT_Rel =*NewCT_Rel ();return _fdf };func NewIscomment ()*Iscomment {_bf :=&Iscomment {};_bf .CT_Empty =*NewCT_Empty ();return _bf };
func NewCT_Rel ()*CT_Rel {_dg :=&CT_Rel {};return _dg };func (_bbb *Textdata )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_bbb .CT_Rel =*NewCT_Rel ();for _ ,_aef :=range start .Attr {if _aef .Name .Local =="\u0069\u0064"{_gb ,_dgg :=_aef .Value ,error (nil );
if _dgg !=nil {return _dgg ;};_bbb .IdAttr =&_gb ;continue ;};};for {_gf ,_ddg :=d .Token ();if _ddg !=nil {return _e .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_ddg );};if _ce ,_ceg :=_gf .(_f .EndElement );
_ceg &&_ce .Name ==start .Name {break ;};};return nil ;};type CT_Empty struct{};

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_fd *Iscomment )ValidateWithPath (path string )error {if _agd :=_fd .CT_Empty .ValidateWithPath (path );_agd !=nil {return _agd ;};return nil ;};

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_cba *Textdata )ValidateWithPath (path string )error {if _gag :=_cba .CT_Rel .ValidateWithPath (path );_gag !=nil {return _gag ;};return nil ;};type Textdata struct{CT_Rel };func (_fg *CT_Empty )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {for {_a ,_c :=d .Token ();
if _c !=nil {return _e .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_c );};if _ec ,_bb :=_a .(_f .EndElement );_bb &&_ec .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Rel and its children
func (_dd *CT_Rel )Validate ()error {return _dd .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func init (){_d .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );
_d .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );
_d .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );
_d .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );
};