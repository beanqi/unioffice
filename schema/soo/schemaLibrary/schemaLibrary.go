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

package schemaLibrary ;import (_d "encoding/xml";_dc "fmt";_c "github.com/unidoc/unioffice";_de "github.com/unidoc/unioffice/common/logger";);func (_dg *SchemaLibrary )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_dg .CT_SchemaLibrary =*NewCT_SchemaLibrary ();_deg :for {_ba ,_ff :=d .Token ();if _ff !=nil {return _ff ;};switch _ac :=_ba .(type ){case _d .StartElement :switch _ac .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_beg :=NewCT_Schema ();if _fdg :=d .DecodeElement (_beg ,&_ac );_fdg !=nil {return _fdg ;};_dg .Schema =append (_dg .Schema ,_beg );default:_de .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0053\u0063\u0068\u0065m\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079 \u0025\u0076",_ac .Name );if _gbc :=d .Skip ();_gbc !=nil {return _gbc ;};};case _d .EndElement :break _deg ;case _d .CharData :};};return nil ;};

// ValidateWithPath validates the SchemaLibrary and its children, prefixing error messages with path
func (_gf *SchemaLibrary )ValidateWithPath (path string )error {if _gde :=_gf .CT_SchemaLibrary .ValidateWithPath (path );_gde !=nil {return _gde ;};return nil ;};

// ValidateWithPath validates the CT_SchemaLibrary and its children, prefixing error messages with path
func (_bg *CT_SchemaLibrary )ValidateWithPath (path string )error {for _aag ,_fa :=range _bg .Schema {if _eg :=_fa .ValidateWithPath (_dc .Sprintf ("\u0025\u0073\u002f\u0053\u0063\u0068\u0065\u006d\u0061\u005b\u0025\u0064\u005d",path ,_aag ));_eg !=nil {return _eg ;};};return nil ;};func (_bff *CT_SchemaLibrary )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_ga :for {_fc ,_gg :=d .Token ();if _gg !=nil {return _gg ;};switch _ee :=_fc .(type ){case _d .StartElement :switch _ee .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e",Local :"\u0073\u0063\u0068\u0065\u006d\u0061"}:_bb :=NewCT_Schema ();if _bfe :=d .DecodeElement (_bb ,&_ee );_bfe !=nil {return _bfe ;};_bff .Schema =append (_bff .Schema ,_bb );default:_de .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u0020\u0025v",_ee .Name );if _da :=d .Skip ();_da !=nil {return _da ;};};case _d .EndElement :break _ga ;case _d .CharData :};};return nil ;};

// ValidateWithPath validates the CT_Schema and its children, prefixing error messages with path
func (_cb *CT_Schema )ValidateWithPath (path string )error {return nil };func (_db *CT_Schema )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _db .UriAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u006d\u0061\u003a\u0075\u0072\u0069"},Value :_dc .Sprintf ("\u0025\u0076",*_db .UriAttr )});};if _db .ManifestLocationAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u006d\u0061\u003a\u006dan\u0069\u0066\u0065\u0073\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"},Value :_dc .Sprintf ("\u0025\u0076",*_db .ManifestLocationAttr )});};if _db .SchemaLocationAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"},Value :_dc .Sprintf ("\u0025\u0076",*_db .SchemaLocationAttr )});};if _db .SchemaLanguageAttr !=nil {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u006d\u0061\u003a\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"},Value :_dc .Sprintf ("\u0025\u0076",*_db .SchemaLanguageAttr )});};e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func NewCT_Schema ()*CT_Schema {_a :=&CT_Schema {};return _a };type CT_Schema struct{UriAttr *string ;ManifestLocationAttr *string ;SchemaLocationAttr *string ;SchemaLanguageAttr *string ;};func (_af *SchemaLibrary )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u006d\u0061"},Value :"\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u006d\u0061:\u0073\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079";return _af .CT_SchemaLibrary .MarshalXML (e ,start );};func NewSchemaLibrary ()*SchemaLibrary {_fb :=&SchemaLibrary {};_fb .CT_SchemaLibrary =*NewCT_SchemaLibrary ();return _fb ;};

// Validate validates the CT_SchemaLibrary and its children
func (_ag *CT_SchemaLibrary )Validate ()error {return _ag .ValidateWithPath ("\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};func (_fd *CT_SchemaLibrary )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {e .EncodeToken (start );if _fd .Schema !=nil {_fdd :=_d .StartElement {Name :_d .Name {Local :"\u006da\u003a\u0073\u0063\u0068\u0065\u006da"}};for _ ,_cd :=range _fd .Schema {e .EncodeElement (_cd ,_fdd );};};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};

// Validate validates the SchemaLibrary and its children
func (_bd *SchemaLibrary )Validate ()error {return _bd .ValidateWithPath ("\u0053\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079");};type CT_SchemaLibrary struct{Schema []*CT_Schema ;};func (_b *CT_Schema )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for _ ,_dbg :=range start .Attr {if _dbg .Name .Local =="\u0075\u0072\u0069"{_f ,_dbe :=_dbg .Value ,error (nil );if _dbe !=nil {return _dbe ;};_b .UriAttr =&_f ;continue ;};if _dbg .Name .Local =="\u006d\u0061n\u0069\u0066\u0065s\u0074\u004c\u006f\u0063\u0061\u0074\u0069\u006f\u006e"{_ea ,_dbd :=_dbg .Value ,error (nil );if _dbd !=nil {return _dbd ;};_b .ManifestLocationAttr =&_ea ;continue ;};if _dbg .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u006f\u0063a\u0074\u0069\u006f\u006e"{_aa ,_be :=_dbg .Value ,error (nil );if _be !=nil {return _be ;};_b .SchemaLocationAttr =&_aa ;continue ;};if _dbg .Name .Local =="\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0061\u006eg\u0075\u0061\u0067\u0065"{_gd ,_gdg :=_dbg .Value ,error (nil );if _gdg !=nil {return _gdg ;};_b .SchemaLanguageAttr =&_gd ;continue ;};};for {_ab ,_ae :=d .Token ();if _ae !=nil {return _dc .Errorf ("p\u0061\u0072\u0073\u0069ng\u0020C\u0054\u005f\u0053\u0063\u0068e\u006d\u0061\u003a\u0020\u0025\u0073",_ae );};if _ef ,_bf :=_ab .(_d .EndElement );_bf &&_ef .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_Schema and its children
func (_eb *CT_Schema )Validate ()error {return _eb .ValidateWithPath ("\u0043T\u005f\u0053\u0063\u0068\u0065\u006da");};type SchemaLibrary struct{CT_SchemaLibrary };func NewCT_SchemaLibrary ()*CT_SchemaLibrary {_gb :=&CT_SchemaLibrary {};return _gb };func init (){_c .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043T\u005f\u0053\u0063\u0068\u0065\u006da",NewCT_Schema );_c .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0043\u0054_\u0053\u0063\u0068e\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewCT_SchemaLibrary );_c .RegisterConstructor ("\u0068\u0074\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067/\u0073\u0063\u0068\u0065\u006da\u004c\u0069\u0062\u0072\u0061\u0072\u0079\u002f\u0032\u0030\u0030\u0036\u002f\u006d\u0061\u0069\u006e","\u0073\u0063\u0068\u0065\u006d\u0061\u004c\u0069\u0062\u0072\u0061\u0072\u0079",NewSchemaLibrary );};