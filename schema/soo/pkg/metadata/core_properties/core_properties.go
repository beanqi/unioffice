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

package core_properties ;import (_c "encoding/xml";_f "fmt";_e "github.com/unidoc/unioffice";_ed "github.com/unidoc/unioffice/common/logger";_g "time";);func (_dbc *CoreProperties )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0063\u0070"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0064\u0063\u0074\u0065\u0072\u006d\u0073"},Value :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0063\u0070\u003a\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073";return _dbc .CT_CoreProperties .MarshalXML (e ,start );};func NewCT_Keyword ()*CT_Keyword {_de :=&CT_Keyword {};return _de };

// Validate validates the CT_Keyword and its children
func (_cbb *CT_Keyword )Validate ()error {return _cbb .ValidateWithPath ("\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064");};func (_gbac *CT_Keyword )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _gbac .LangAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_f .Sprintf ("\u0025\u0076",*_gbac .LangAttr )});
};e .EncodeElement (_gbac .Content ,start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the CT_CoreProperties and its children, prefixing error messages with path
func (_fe *CT_CoreProperties )ValidateWithPath (path string )error {if _fe .Keywords !=nil {if _feg :=_fe .Keywords .ValidateWithPath (path +"\u002fK\u0065\u0079\u0077\u006f\u0072\u0064s");_feg !=nil {return _feg ;};};return nil ;};func (_ecd *CT_Keywords )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _ecd .LangAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006c\u003a\u006c\u0061\u006e\u0067"},Value :_f .Sprintf ("\u0025\u0076",*_ecd .LangAttr )});
};e .EncodeToken (start );if _ecd .Value !=nil {_acb :=_c .StartElement {Name :_c .Name {Local :"\u0063\u0070\u003a\u0076\u0061\u006c\u0075\u0065"}};for _ ,_ff :=range _ecd .Value {e .EncodeElement (_ff ,_acb );};};e .EncodeToken (_c .EndElement {Name :start .Name });
return nil ;};

// Validate validates the CoreProperties and its children
func (_fdba *CoreProperties )Validate ()error {return _fdba .ValidateWithPath ("\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};type CT_Keywords struct{LangAttr *string ;Value []*CT_Keyword ;};type CoreProperties struct{CT_CoreProperties };
type CT_Keyword struct{LangAttr *string ;Content string ;};func (_fbb *CT_Keywords )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_fg :=range start .Attr {if _fg .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_fg .Name .Local =="\u006c\u0061\u006e\u0067"{_gbd ,_dbg :=_fg .Value ,error (nil );
if _dbg !=nil {return _dbg ;};_fbb .LangAttr =&_gbd ;continue ;};};_efg :for {_ddc ,_bgc :=d .Token ();if _bgc !=nil {return _bgc ;};switch _ge :=_ddc .(type ){case _c .StartElement :switch _ge .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076\u0061\u006cu\u0065"}:_cbf :=NewCT_Keyword ();
if _ce :=d .DecodeElement (_cbf ,&_ge );_ce !=nil {return _ce ;};_fbb .Value =append (_fbb .Value ,_cbf );default:_ed .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073\u0020\u0025\u0076",_ge .Name );
if _bfa :=d .Skip ();_bfa !=nil {return _bfa ;};};case _c .EndElement :break _efg ;case _c .CharData :};};return nil ;};func NewCT_Keywords ()*CT_Keywords {_gd :=&CT_Keywords {};return _gd };

// Validate validates the CT_Keywords and its children
func (_gcc *CT_Keywords )Validate ()error {return _gcc .ValidateWithPath ("C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073");};func (_gg *CT_CoreProperties )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_gc :for {_bfg ,_ga :=d .Token ();
if _ga !=nil {return _ga ;};switch _fdd :=_bfg .(type ){case _c .StartElement :switch _fdd .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_gg .Category =new (string );
if _cg :=d .DecodeElement (_gg .Category ,&_fdd );_cg !=nil {return _cg ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_gg .ContentStatus =new (string );
if _ac :=d .DecodeElement (_gg .ContentStatus ,&_fdd );_ac !=nil {return _ac ;};case _c .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_gg .Created =new (_e .XSDAny );
if _ccb :=d .DecodeElement (_gg .Created ,&_fdd );_ccb !=nil {return _ccb ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_gg .Creator =new (_e .XSDAny );
if _bfe :=d .DecodeElement (_gg .Creator ,&_fdd );_bfe !=nil {return _bfe ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_gg .Description =new (_e .XSDAny );
if _edf :=d .DecodeElement (_gg .Description ,&_fdd );_edf !=nil {return _edf ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_gg .Identifier =new (_e .XSDAny );
if _fda :=d .DecodeElement (_gg .Identifier ,&_fdd );_fda !=nil {return _fda ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_gg .Keywords =NewCT_Keywords ();
if _cde :=d .DecodeElement (_gg .Keywords ,&_fdd );_cde !=nil {return _cde ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_gg .Language =new (_e .XSDAny );
if _ef :=d .DecodeElement (_gg .Language ,&_fdd );_ef !=nil {return _ef ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_gg .LastModifiedBy =new (string );
if _ebf :=d .DecodeElement (_gg .LastModifiedBy ,&_fdd );_ebf !=nil {return _ebf ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_gg .LastPrinted =new (_g .Time );
if _ae :=d .DecodeElement (_gg .LastPrinted ,&_fdd );_ae !=nil {return _ae ;};case _c .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_gg .Modified =new (_e .XSDAny );
if _aeg :=d .DecodeElement (_gg .Modified ,&_fdd );_aeg !=nil {return _aeg ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_gg .Revision =new (string );
if _ad :=d .DecodeElement (_gg .Revision ,&_fdd );_ad !=nil {return _ad ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_gg .Subject =new (_e .XSDAny );
if _bc :=d .DecodeElement (_gg .Subject ,&_fdd );_bc !=nil {return _bc ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_gg .Title =new (_e .XSDAny );
if _dd :=d .DecodeElement (_gg .Title ,&_fdd );_dd !=nil {return _dd ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_gg .Version =new (string );
if _ca :=d .DecodeElement (_gg .Version ,&_fdd );_ca !=nil {return _ca ;};default:_ed .Log .Debug ("\u0073\u006bi\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073\u0075\u0070p\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0069\u0065\u0073\u0020\u0025\u0076",_fdd .Name );
if _ebff :=d .Skip ();_ebff !=nil {return _ebff ;};};case _c .EndElement :break _gc ;case _c .CharData :};};return nil ;};

// ValidateWithPath validates the CT_Keywords and its children, prefixing error messages with path
func (_gbdc *CT_Keywords )ValidateWithPath (path string )error {for _gca ,_gec :=range _gbdc .Value {if _gab :=_gec .ValidateWithPath (_f .Sprintf ("\u0025\u0073\u002fV\u0061\u006c\u0075\u0065\u005b\u0025\u0064\u005d",path ,_gca ));_gab !=nil {return _gab ;
};};return nil ;};func (_fb *CT_CoreProperties )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {e .EncodeToken (start );if _fb .Category !=nil {_db :=_c .StartElement {Name :_c .Name {Local :"c\u0070\u003a\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}};
_e .AddPreserveSpaceAttr (&_db ,*_fb .Category );e .EncodeElement (_fb .Category ,_db );};if _fb .ContentStatus !=nil {_a :=_c .StartElement {Name :_c .Name {Local :"\u0063\u0070:\u0063\u006f\u006et\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}};
_e .AddPreserveSpaceAttr (&_a ,*_fb .ContentStatus );e .EncodeElement (_fb .ContentStatus ,_a );};if _fb .Created !=nil {_da :=_c .StartElement {Name :_c .Name {Local :"\u0064c\u0074e\u0072\u006d\u0073\u003a\u0063\u0072\u0065\u0061\u0074\u0065\u0064"}};
e .EncodeElement (_fb .Created ,_da );};if _fb .Creator !=nil {_dc :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063\u003a\u0063\u0072\u0065\u0061\u0074\u006f\u0072"}};e .EncodeElement (_fb .Creator ,_dc );};if _fb .Description !=nil {_eb :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063\u003a\u0064\u0065\u0073\u0063\u0072\u0069p\u0074\u0069\u006f\u006e"}};
e .EncodeElement (_fb .Description ,_eb );};if _fb .Identifier !=nil {_b :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063\u003a\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}};e .EncodeElement (_fb .Identifier ,_b );};if _fb .Keywords !=nil {_gb :=_c .StartElement {Name :_c .Name {Local :"c\u0070\u003a\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}};
e .EncodeElement (_fb .Keywords ,_gb );};if _fb .Language !=nil {_bf :=_c .StartElement {Name :_c .Name {Local :"d\u0063\u003a\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}};e .EncodeElement (_fb .Language ,_bf );};if _fb .LastModifiedBy !=nil {_cd :=_c .StartElement {Name :_c .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}};
_e .AddPreserveSpaceAttr (&_cd ,*_fb .LastModifiedBy );e .EncodeElement (_fb .LastModifiedBy ,_cd );};if _fb .LastPrinted !=nil {_cda :=_c .StartElement {Name :_c .Name {Local :"\u0063\u0070\u003a\u006c\u0061\u0073\u0074\u0050\u0072i\u006e\u0074\u0065\u0064"}};
e .EncodeElement (_fb .LastPrinted ,_cda );};if _fb .Modified !=nil {_ag :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063t\u0065\u0072\u006ds\u003a\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}};e .EncodeElement (_fb .Modified ,_ag );};if _fb .Revision !=nil {_cb :=_c .StartElement {Name :_c .Name {Local :"c\u0070\u003a\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}};
_e .AddPreserveSpaceAttr (&_cb ,*_fb .Revision );e .EncodeElement (_fb .Revision ,_cb );};if _fb .Subject !=nil {_cc :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063\u003a\u0073\u0075\u0062\u006a\u0065\u0063\u0074"}};e .EncodeElement (_fb .Subject ,_cc );
};if _fb .Title !=nil {_gba :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063\u003a\u0074\u0069\u0074\u006c\u0065"}};e .EncodeElement (_fb .Title ,_gba );};if _fb .Version !=nil {_fd :=_c .StartElement {Name :_c .Name {Local :"\u0063\u0070\u003a\u0076\u0065\u0072\u0073\u0069\u006f\u006e"}};
_e .AddPreserveSpaceAttr (&_fd ,*_fb .Version );e .EncodeElement (_fb .Version ,_fd );};e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type CT_CoreProperties struct{Category *string ;ContentStatus *string ;Created *_e .XSDAny ;Creator *_e .XSDAny ;
Description *_e .XSDAny ;Identifier *_e .XSDAny ;Keywords *CT_Keywords ;Language *_e .XSDAny ;LastModifiedBy *string ;LastPrinted *_g .Time ;Modified *_e .XSDAny ;Revision *string ;Subject *_e .XSDAny ;Title *_e .XSDAny ;Version *string ;};

// ValidateWithPath validates the CoreProperties and its children, prefixing error messages with path
func (_fea *CoreProperties )ValidateWithPath (path string )error {if _fca :=_fea .CT_CoreProperties .ValidateWithPath (path );_fca !=nil {return _fca ;};return nil ;};func (_ced *CoreProperties )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_ced .CT_CoreProperties =*NewCT_CoreProperties ();
_eba :for {_dee ,_bd :=d .Token ();if _bd !=nil {return _bd ;};switch _ccbd :=_dee .(type ){case _c .StartElement :switch _ccbd .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u0061\u0074\u0065\u0067\u006f\u0072\u0079"}:_ced .Category =new (string );
if _gac :=d .DecodeElement (_ced .Category ,&_ccbd );_gac !=nil {return _gac ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0063\u006f\u006e\u0074\u0065\u006e\u0074\u0053\u0074\u0061\u0074\u0075\u0073"}:_ced .ContentStatus =new (string );
if _cf :=d .DecodeElement (_ced .ContentStatus ,&_ccbd );_cf !=nil {return _cf ;};case _c .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u0063r\u0065\u0061\u0074\u0065\u0064"}:_ced .Created =new (_e .XSDAny );
if _gdb :=d .DecodeElement (_ced .Created ,&_ccbd );_gdb !=nil {return _gdb ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0063r\u0065\u0061\u0074\u006f\u0072"}:_ced .Creator =new (_e .XSDAny );
if _gf :=d .DecodeElement (_ced .Creator ,&_ccbd );_gf !=nil {return _gf ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"d\u0065\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e"}:_ced .Description =new (_e .XSDAny );
if _af :=d .DecodeElement (_ced .Description ,&_ccbd );_af !=nil {return _af ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0069\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072"}:_ced .Identifier =new (_e .XSDAny );
if _fdb :=d .DecodeElement (_ced .Identifier ,&_ccbd );_fdb !=nil {return _fdb ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"}:_ced .Keywords =NewCT_Keywords ();
if _ccbf :=d .DecodeElement (_ced .Keywords ,&_ccbd );_ccbf !=nil {return _ccbf ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u006c\u0061\u006e\u0067\u0075\u0061\u0067\u0065"}:_ced .Language =new (_e .XSDAny );
if _fga :=d .DecodeElement (_ced .Language ,&_ccbd );_fga !=nil {return _fga ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u006c\u0061\u0073\u0074\u004d\u006f\u0064\u0069\u0066i\u0065\u0064\u0042\u0079"}:_ced .LastModifiedBy =new (string );
if _ebc :=d .DecodeElement (_ced .LastModifiedBy ,&_ccbd );_ebc !=nil {return _ebc ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"l\u0061\u0073\u0074\u0050\u0072\u0069\u006e\u0074\u0065\u0064"}:_ced .LastPrinted =new (_g .Time );
if _aec :=d .DecodeElement (_ced .LastPrinted ,&_ccbd );_aec !=nil {return _aec ;};case _c .Name {Space :"\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/",Local :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0064"}:_ced .Modified =new (_e .XSDAny );
if _fdf :=d .DecodeElement (_ced .Modified ,&_ccbd );_fdf !=nil {return _fdf ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0072\u0065\u0076\u0069\u0073\u0069\u006f\u006e"}:_ced .Revision =new (string );
if _cab :=d .DecodeElement (_ced .Revision ,&_ccbd );_cab !=nil {return _cab ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0073u\u0062\u006a\u0065\u0063\u0074"}:_ced .Subject =new (_e .XSDAny );
if _agc :=d .DecodeElement (_ced .Subject ,&_ccbd );_agc !=nil {return _agc ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0074\u0069\u0074l\u0065"}:_ced .Title =new (_e .XSDAny );
if _ade :=d .DecodeElement (_ced .Title ,&_ccbd );_ade !=nil {return _ade ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073",Local :"\u0076e\u0072\u0073\u0069\u006f\u006e"}:_ced .Version =new (string );
if _fdae :=d .DecodeElement (_ced .Version ,&_ccbd );_fdae !=nil {return _fdae ;};default:_ed .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065\u0072t\u0069e\u0073\u0020\u0025\u0076",_ccbd .Name );
if _aef :=d .Skip ();_aef !=nil {return _aef ;};};case _c .EndElement :break _eba ;case _c .CharData :};};return nil ;};func NewCT_CoreProperties ()*CT_CoreProperties {_ec :=&CT_CoreProperties {};return _ec };func NewCoreProperties ()*CoreProperties {_gga :=&CoreProperties {};
_gga .CT_CoreProperties =*NewCT_CoreProperties ();return _gga ;};

// ValidateWithPath validates the CT_Keyword and its children, prefixing error messages with path
func (_bcc *CT_Keyword )ValidateWithPath (path string )error {return nil };func (_ebg *CT_Keyword )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_eg :=range start .Attr {if _eg .Name .Space =="\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"&&_eg .Name .Local =="\u006c\u0061\u006e\u0067"{_df ,_fae :=_eg .Value ,error (nil );
if _fae !=nil {return _fae ;};_ebg .LangAttr =&_df ;continue ;};};for {_dca ,_ee :=d .Token ();if _ee !=nil {return _f .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u003a\u0020%\u0073",_ee );
};if _faa ,_dcf :=_dca .(_c .CharData );_dcf {_ebg .Content =string (_faa );};if _fc ,_bg :=_dca .(_c .EndElement );_bg &&_fc .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_CoreProperties and its children
func (_fa *CT_CoreProperties )Validate ()error {return _fa .ValidateWithPath ("\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073");};func init (){_e .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u0043\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCT_CoreProperties );
_e .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","C\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",NewCT_Keywords );
_e .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0043\u0054\u005f\u004b\u0065\u0079\u0077\u006f\u0072\u0064",NewCT_Keyword );
_e .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073\u0063h\u0065\u006d\u0061\u0073\u002e\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0070\u0061c\u006b\u0061g\u0065\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0065\u0074\u0061\u0064\u0061ta\u002f\u0063\u006f\u0072\u0065\u002d\u0070\u0072\u006fp\u0065\u0072\u0074\u0069\u0065\u0073","\u0063\u006f\u0072\u0065\u0050\u0072\u006f\u0070\u0065r\u0074\u0069\u0065\u0073",NewCoreProperties );
};