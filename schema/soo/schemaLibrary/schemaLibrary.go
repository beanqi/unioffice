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

package schemaLibrary ;import (_f "encoding/xml";_d "fmt";_fb "github.com/unidoc/unioffice";_da "github.com/unidoc/unioffice/common/logger";);func (_c *CT_Schema )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {for _ ,_e :=range start .Attr {if _e .Name .Local =="\u0075\u0072\u0069"{_ge ,_dg :=_e .Value ,error (nil );
if _dg !=nil {return _dg ;};_c .UriAttr =&_ge ;continue ;};if _e .Name .Local =="\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"{_dgb ,_fa :=_e .Value ,error (nil );if _fa !=nil {return _fa ;};_c .ManifestLocationAttr =&_dgb ;
continue ;};if _e .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"{_dac ,_ce :=_e .Value ,error (nil );if _ce !=nil {return _ce ;};_c .SchemaLocationAttr =&_dac ;continue ;};if _e .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"{_gg ,_fg :=_e .Value ,error (nil );
if _fg !=nil {return _fg ;};_c .SchemaLanguageAttr =&_gg ;continue ;};};for {_b ,_bf :=d .Token ();if _bf !=nil {return _d .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073",_bf );};if _a ,_db :=_b .(_f .EndElement );
_db &&_a .Name ==start .Name {break ;};};return nil ;};func (_dgg *SchemaLibrary )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});
start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});
start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079";return _dgg .CT_SchemaLibrary .MarshalXML (e ,start );};func NewCT_Schema ()*CT_Schema {_ga :=&CT_Schema {};return _ga };func (_fad *CT_SchemaLibrary )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_fc :for {_bc ,_ba :=d .Token ();
if _ba !=nil {return _ba ;};switch _fd :=_bc .(type ){case _f .StartElement :switch _fd .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_cb :=NewCT_Schema ();
if _cgd :=d .DecodeElement (_cb ,&_fd );_cgd !=nil {return _cgd ;};_fad .Schema =append (_fad .Schema ,_cb );default:_da .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v",_fd .Name );
if _gc :=d .Skip ();_gc !=nil {return _gc ;};};case _f .EndElement :break _fc ;case _f .CharData :};};return nil ;};type SchemaLibrary struct{CT_SchemaLibrary };func (_ad *SchemaLibrary )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_ad .CT_SchemaLibrary =*NewCT_SchemaLibrary ();
_gb :for {_ea ,_bcb :=d .Token ();if _bcb !=nil {return _bcb ;};switch _bd :=_ea .(type ){case _f .StartElement :switch _bd .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_aa :=NewCT_Schema ();
if _dag :=d .DecodeElement (_aa ,&_bd );_dag !=nil {return _dag ;};_ad .Schema =append (_ad .Schema ,_aa );default:_da .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076",_bd .Name );
if _cab :=d .Skip ();_cab !=nil {return _cab ;};};case _f .EndElement :break _gb ;case _f .CharData :};};return nil ;};func (_ca *CT_SchemaLibrary )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {e .EncodeToken (start );if _ca .Schema !=nil {_cgg :=_f .StartElement {Name :_f .Name {Local :"\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}};
for _ ,_dc :=range _ca .Schema {e .EncodeElement (_dc ,_cgg );};};e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};func NewSchemaLibrary ()*SchemaLibrary {_ggf :=&SchemaLibrary {};_ggf .CT_SchemaLibrary =*NewCT_SchemaLibrary ();return _ggf ;
};

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_ff *SchemaLibrary )ValidateWithPath (path string )error {if _bb :=_ff .CT_SchemaLibrary .ValidateWithPath (path );_bb !=nil {return _bb ;};return nil ;};func NewCT_SchemaLibrary ()*CT_SchemaLibrary {_gdf :=&CT_SchemaLibrary {};return _gdf };

// Validate validates the CT_SchemaLibrary and its children
func (_be *CT_SchemaLibrary )Validate ()error {return _be .ValidateWithPath ("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};func (_dd *CT_Schema )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {if _dd .UriAttr !=nil {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u006d\u0061\u003a\u0075\u0072\u0069"},Value :_d .Sprintf ("\u0025\u0076",*_dd .UriAttr )});
};if _dd .ManifestLocationAttr !=nil {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"},Value :_d .Sprintf ("\u0025\u0076",*_dd .ManifestLocationAttr )});
};if _dd .SchemaLocationAttr !=nil {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"},Value :_d .Sprintf ("\u0025\u0076",*_dd .SchemaLocationAttr )});
};if _dd .SchemaLanguageAttr !=nil {start .Attr =append (start .Attr ,_f .Attr {Name :_f .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"},Value :_d .Sprintf ("\u0025\u0076",*_dd .SchemaLanguageAttr )});
};e .EncodeToken (start );e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_ee *CT_SchemaLibrary )ValidateWithPath (path string )error {for _eg ,_eb :=range _ee .Schema {if _ec :=_eb .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d",path ,_eg ));_ec !=nil {return _ec ;
};};return nil ;};type CT_Schema struct{UriAttr *string ;ManifestLocationAttr *string ;SchemaLocationAttr *string ;SchemaLanguageAttr *string ;};type CT_SchemaLibrary struct{Schema []*CT_Schema ;};

// Validate validates the SchemaLibrary and its children
func (_gf *SchemaLibrary )Validate ()error {return _gf .ValidateWithPath ("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_cg *CT_Schema )ValidateWithPath (path string )error {return nil };

// Validate validates the CT_Schema and its children
func (_dde *CT_Schema )Validate ()error {return _dde .ValidateWithPath ("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da");};func init (){_fb .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043T\u005f\u0053\u0063\u0068\u0065\u006da",NewCT_Schema );
_fb .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewCT_SchemaLibrary );
_fb .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewSchemaLibrary );
};