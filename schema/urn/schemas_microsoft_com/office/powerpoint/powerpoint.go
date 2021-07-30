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

package powerpoint ;import (_c "encoding/xml";_f "fmt";_fa "github.com/unidoc/unioffice";);

// ValidateWithPath validates the CT_Empty and its children, prefixing error messages with path
func (_ff *CT_Empty )ValidateWithPath (path string )error {return nil };func (_bf *Iscomment )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0069s\u0063\u006f\u006d\u006d\u0065\u006et";return _bf .CT_Empty .MarshalXML (e ,start );};type Iscomment struct{CT_Empty };func NewCT_Rel ()*CT_Rel {_dd :=&CT_Rel {};return _dd };

// Validate validates the CT_Empty and its children
func (_dg *CT_Empty )Validate ()error {return _dg .ValidateWithPath ("\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079");};type CT_Rel struct{IdAttr *string ;};

// ValidateWithPath validates the Iscomment and its children, prefixing error messages with path
func (_dfa *Iscomment )ValidateWithPath (path string )error {if _ag :=_dfa .CT_Empty .ValidateWithPath (path );_ag !=nil {return _ag ;};return nil ;};func (_df *CT_Rel )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _df .IdAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0069\u0064"},Value :_f .Sprintf ("\u0025\u0076",*_df .IdAttr )});};e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_bd *CT_Rel )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_ce :=range start .Attr {if _ce .Name .Local =="\u0069\u0064"{_ad ,_db :=_ce .Value ,error (nil );if _db !=nil {return _db ;};_bd .IdAttr =&_ad ;continue ;};};for {_bcc ,_fgd :=d .Token ();if _fgd !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0043T\u005f\u0052e\u006c\u003a\u0020\u0025\u0073",_fgd );};if _aa ,_e :=_bcc .(_c .EndElement );_e &&_aa .Name ==start .Name {break ;};};return nil ;};func NewCT_Empty ()*CT_Empty {_fe :=&CT_Empty {};return _fe };

// ValidateWithPath validates the CT_Rel and its children, prefixing error messages with path
func (_cg *CT_Rel )ValidateWithPath (path string )error {return nil };func (_cd *Textdata )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074"});start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061";return _cd .CT_Rel .MarshalXML (e ,start );};

// Validate validates the CT_Rel and its children
func (_ga *CT_Rel )Validate ()error {return _ga .ValidateWithPath ("\u0043\u0054\u005f\u0052\u0065\u006c");};func NewTextdata ()*Textdata {_cc :=&Textdata {};_cc .CT_Rel =*NewCT_Rel ();return _cc };type Textdata struct{CT_Rel };

// ValidateWithPath validates the Textdata and its children, prefixing error messages with path
func (_eg *Textdata )ValidateWithPath (path string )error {if _cf :=_eg .CT_Rel .ValidateWithPath (path );_cf !=nil {return _cf ;};return nil ;};func (_agf *Textdata )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_agf .CT_Rel =*NewCT_Rel ();for _ ,_gff :=range start .Attr {if _gff .Name .Local =="\u0069\u0064"{_cgb ,_ccc :=_gff .Value ,error (nil );if _ccc !=nil {return _ccc ;};_agf .IdAttr =&_cgb ;continue ;};};for {_aaa ,_bb :=d .Token ();if _bb !=nil {return _f .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0054\u0065\u0078t\u0064\u0061\u0074\u0061: \u0025\u0073",_bb );};if _beb ,_ea :=_aaa .(_c .EndElement );_ea &&_beb .Name ==start .Name {break ;};};return nil ;};func (_a *CT_Empty )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func NewIscomment ()*Iscomment {_fgf :=&Iscomment {};_fgf .CT_Empty =*NewCT_Empty ();return _fgf };type CT_Empty struct{};func (_bce *Iscomment )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_bce .CT_Empty =*NewCT_Empty ();for {_acb ,_fc :=d .Token ();if _fc !=nil {return _f .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020I\u0073\u0063\u006f\u006d\u006de\u006e\u0074\u003a\u0020\u0025\u0073",_fc );};if _da ,_ffe :=_acb .(_c .EndElement );_ffe &&_da .Name ==start .Name {break ;};};return nil ;};func (_bc *CT_Empty )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_fg ,_be :=d .Token ();if _be !=nil {return _f .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0043\u0054\u005fE\u006d\u0070\u0074\u0079: \u0025\u0073",_be );};if _d ,_ac :=_fg .(_c .EndElement );_ac &&_d .Name ==start .Name {break ;};};return nil ;};

// Validate validates the Iscomment and its children
func (_cb *Iscomment )Validate ()error {return _cb .ValidateWithPath ("\u0049s\u0063\u006f\u006d\u006d\u0065\u006et");};

// Validate validates the Textdata and its children
func (_ca *Textdata )Validate ()error {return _ca .ValidateWithPath ("\u0054\u0065\u0078\u0074\u0064\u0061\u0074\u0061");};func init (){_fa .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0045\u006d\u0070\u0074\u0079",NewCT_Empty );_fa .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0043\u0054\u005f\u0052\u0065\u006c",NewCT_Rel );_fa .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0069s\u0063\u006f\u006d\u006d\u0065\u006et",NewIscomment );_fa .RegisterConstructor ("\u0075\u0072\u006e\u003a\u0073\u0063\u0068e\u006d\u0061\u0073-\u006d\u0069\u0063\u0072o\u0073\u006f\u0066\u0074\u002d\u0063\u006f\u006d\u003a\u006f\u0066\u0066\u0069\u0063\u0065\u003a\u0070\u006f\u0077\u0065\u0072\u0070\u006f\u0069\u006e\u0074","\u0074\u0065\u0078\u0074\u0064\u0061\u0074\u0061",NewTextdata );};