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

package relationships ;import (_gg "encoding/xml";_d "fmt";_e "github.com/unidoc/unioffice";_f "github.com/unidoc/unioffice/common/logger";);

// ValidateWithPath validates the CT_Relationships and its children, prefixing error messages with path
func (_fae *CT_Relationships )ValidateWithPath (path string )error {for _cb ,_bg :=range _fae .Relationship {if _fb :=_bg .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002f\u0052el\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u005b\u0025\u0064\u005d",path ,_cb ));
_fb !=nil {return _fb ;};};return nil ;};func (_ffg *ST_TargetMode )UnmarshalXMLAttr (attr _gg .Attr )error {switch attr .Value {case "":*_ffg =0;case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_ffg =1;case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_ffg =2;
};return nil ;};

// ValidateWithPath validates the Relationship and its children, prefixing error messages with path
func (_gdf *Relationship )ValidateWithPath (path string )error {if _gec :=_gdf .CT_Relationship .ValidateWithPath (path );_gec !=nil {return _gec ;};return nil ;};func (_cde *Relationship )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {return _cde .CT_Relationship .MarshalXML (e ,start );
};const (ST_TargetModeUnset ST_TargetMode =0;ST_TargetModeExternal ST_TargetMode =1;ST_TargetModeInternal ST_TargetMode =2;);func (_ga *CT_Relationship )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {for _ ,_gdd :=range start .Attr {if _gdd .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"{_ga .TargetModeAttr .UnmarshalXMLAttr (_gdd );
continue ;};if _gdd .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074"{_fd ,_daf :=_gdd .Value ,error (nil );if _daf !=nil {return _daf ;};_ga .TargetAttr =_fd ;continue ;};if _gdd .Name .Local =="\u0054\u0079\u0070\u0065"{_ec ,_faf :=_gdd .Value ,error (nil );
if _faf !=nil {return _faf ;};_ga .TypeAttr =_ec ;continue ;};if _gdd .Name .Local =="\u0049\u0064"{_ff ,_gf :=_gdd .Value ,error (nil );if _gf !=nil {return _gf ;};_ga .IdAttr =_ff ;continue ;};};for {_fdc ,_ce :=d .Token ();if _ce !=nil {return _d .Errorf ("p\u0061\u0072\u0073\u0069\u006e\u0067 \u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069o\u006e\u0073\u0068i\u0070:\u0020\u0025\u0073",_ce );
};if _eg ,_ac :=_fdc .(_gg .CharData );_ac {_ga .Content =string (_eg );};if _gag ,_aa :=_fdc .(_gg .EndElement );_aa &&_gag .Name ==start .Name {break ;};};return nil ;};func (_ece ST_TargetMode )ValidateWithPath (path string )error {switch _ece {case 0,1,2:default:return _d .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_ece ));
};return nil ;};func (_bd *ST_TargetMode )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_bgb ,_aca :=d .Token ();if _aca !=nil {return _aca ;};if _ea ,_bc :=_bgb .(_gg .EndElement );_bc &&_ea .Name ==start .Name {*_bd =1;return nil ;};
if _eba ,_faa :=_bgb .(_gg .CharData );!_faa {return _d .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_bgb );}else {switch string (_eba ){case "":*_bd =0;
case "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_bd =1;case "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_bd =2;};};_bgb ,_aca =d .Token ();if _aca !=nil {return _aca ;};if _acb ,_gb :=_bgb .(_gg .EndElement );_gb &&_acb .Name ==start .Name {return nil ;
};return _d .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_bgb );};func (_gca *Relationship )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_gca .CT_Relationship =*NewCT_Relationship ();
for _ ,_eca :=range start .Attr {if _eca .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"{_gca .TargetModeAttr .UnmarshalXMLAttr (_eca );continue ;};if _eca .Name .Local =="\u0054\u0061\u0072\u0067\u0065\u0074"{_eea ,_ae :=_eca .Value ,error (nil );
if _ae !=nil {return _ae ;};_gca .TargetAttr =_eea ;continue ;};if _eca .Name .Local =="\u0054\u0079\u0070\u0065"{_ffb ,_fc :=_eca .Value ,error (nil );if _fc !=nil {return _fc ;};_gca .TypeAttr =_ffb ;continue ;};if _eca .Name .Local =="\u0049\u0064"{_ef ,_fe :=_eca .Value ,error (nil );
if _fe !=nil {return _fe ;};_gca .IdAttr =_ef ;continue ;};};for {_gcg ,_gfe :=d .Token ();if _gfe !=nil {return _d .Errorf ("\u0070a\u0072\u0073\u0069\u006e\u0067\u0020\u0052\u0065\u006c\u0061\u0074i\u006f\u006e\u0073\u0068\u0069\u0070\u003a\u0020\u0025\u0073",_gfe );
};if _be ,_beb :=_gcg .(_gg .EndElement );_beb &&_be .Name ==start .Name {break ;};};return nil ;};func (_ddb ST_TargetMode )MarshalXMLAttr (name _gg .Name )(_gg .Attr ,error ){_ab :=_gg .Attr {};_ab .Name =name ;switch _ddb {case ST_TargetModeUnset :_ab .Value ="";
case ST_TargetModeExternal :_ab .Value ="\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c";case ST_TargetModeInternal :_ab .Value ="\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c";};return _ab ,nil ;};func NewCT_Relationship ()*CT_Relationship {_fa :=&CT_Relationship {};
return _fa };

// ValidateWithPath validates the Relationships and its children, prefixing error messages with path
func (_fea *Relationships )ValidateWithPath (path string )error {if _age :=_fea .CT_Relationships .ValidateWithPath (path );_age !=nil {return _age ;};return nil ;};func (_c *CT_Relationship )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {if _c .TargetModeAttr !=ST_TargetModeUnset {_da ,_a :=_c .TargetModeAttr .MarshalXMLAttr (_gg .Name {Local :"\u0054\u0061\u0072\u0067\u0065\u0074\u004d\u006f\u0064\u0065"});
if _a !=nil {return _a ;};start .Attr =append (start .Attr ,_da );};start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0054\u0061\u0072\u0067\u0065\u0074"},Value :_d .Sprintf ("\u0025\u0076",_c .TargetAttr )});start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0054\u0079\u0070\u0065"},Value :_d .Sprintf ("\u0025\u0076",_c .TypeAttr )});
start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0049\u0064"},Value :_d .Sprintf ("\u0025\u0076",_c .IdAttr )});e .EncodeElement (_c .Content ,start );e .EncodeToken (_gg .EndElement {Name :start .Name });return nil ;};func (_ca *Relationships )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s"});
start .Attr =append (start .Attr ,_gg .Attr {Name :_gg .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073";return _ca .CT_Relationships .MarshalXML (e ,start );};type Relationship struct{CT_Relationship };func (_cd *CT_Relationships )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_ba :for {_dg ,_cda :=d .Token ();
if _cda !=nil {return _cda ;};switch _ecb :=_dg .(type ){case _gg .StartElement :switch _ecb .Name {case _gg .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s",Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:_af :=NewRelationship ();
if _afe :=d .DecodeElement (_af ,&_ecb );_afe !=nil {return _afe ;};_cd .Relationship =append (_cd .Relationship ,_af );default:_f .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073\u0020\u0025v",_ecb .Name );
if _cfd :=d .Skip ();_cfd !=nil {return _cfd ;};};case _gg .EndElement :break _ba ;case _gg .CharData :};};return nil ;};func NewRelationships ()*Relationships {_gaf :=&Relationships {};_gaf .CT_Relationships =*NewCT_Relationships ();return _gaf ;};

// Validate validates the CT_Relationships and its children
func (_ee *CT_Relationships )Validate ()error {return _ee .ValidateWithPath ("\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073");};type ST_TargetMode byte ;func (_efb *Relationships )UnmarshalXML (d *_gg .Decoder ,start _gg .StartElement )error {_efb .CT_Relationships =*NewCT_Relationships ();
_dc :for {_dge ,_bef :=d .Token ();if _bef !=nil {return _bef ;};switch _aee :=_dge .(type ){case _gg .StartElement :switch _aee .Name {case _gg .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s",Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}:_ad :=NewRelationship ();
if _cff :=d .DecodeElement (_ad ,&_aee );_cff !=nil {return _cff ;};_efb .Relationship =append (_efb .Relationship ,_ad );default:_f .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0052\u0065\u006c\u0061t\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073 \u0025\u0076",_aee .Name );
if _efbf :=d .Skip ();_efbf !=nil {return _efbf ;};};case _gg .EndElement :break _dc ;case _gg .CharData :};};return nil ;};type Relationships struct{CT_Relationships };type CT_Relationship struct{TargetModeAttr ST_TargetMode ;TargetAttr string ;TypeAttr string ;
IdAttr string ;Content string ;};func (_ge *CT_Relationships )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {e .EncodeToken (start );if _ge .Relationship !=nil {_eb :=_gg .StartElement {Name :_gg .Name {Local :"\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070"}};
for _ ,_cf :=range _ge .Relationship {e .EncodeElement (_cf ,_eb );};};e .EncodeToken (_gg .EndElement {Name :start .Name });return nil ;};func NewCT_Relationships ()*CT_Relationships {_b :=&CT_Relationships {};return _b };

// Validate validates the CT_Relationship and its children
func (_dd *CT_Relationship )Validate ()error {return _dd .ValidateWithPath ("\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070");};

// Validate validates the Relationships and its children
func (_cfde *Relationships )Validate ()error {return _cfde .ValidateWithPath ("\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073");};func (_db ST_TargetMode )MarshalXML (e *_gg .Encoder ,start _gg .StartElement )error {return e .EncodeElement (_db .String (),start );
};func (_bgd ST_TargetMode )String ()string {switch _bgd {case 0:return "";case 1:return "\u0045\u0078\u0074\u0065\u0072\u006e\u0061\u006c";case 2:return "\u0049\u006e\u0074\u0065\u0072\u006e\u0061\u006c";};return "";};func NewRelationship ()*Relationship {_ag :=&Relationship {};
_ag .CT_Relationship =*NewCT_Relationship ();return _ag ;};

// ValidateWithPath validates the CT_Relationship and its children, prefixing error messages with path
func (_fg *CT_Relationship )ValidateWithPath (path string )error {if _gc :=_fg .TargetModeAttr .ValidateWithPath (path +"\u002fT\u0061r\u0067\u0065\u0074\u004d\u006f\u0064\u0065\u0041\u0074\u0074\u0072");_gc !=nil {return _gc ;};return nil ;};func (_gfb ST_TargetMode )Validate ()error {return _gfb .ValidateWithPath ("")};
type CT_Relationships struct{Relationship []*Relationship ;};

// Validate validates the Relationship and its children
func (_aeb *Relationship )Validate ()error {return _aeb .ValidateWithPath ("\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070");};func init (){_e .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0043\u0054_\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073",NewCT_Relationships );
_e .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0043T\u005fR\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070",NewCT_Relationship );
_e .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0052\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073",NewRelationships );
_e .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006fn\u0073h\u0069\u0070s","\u0052\u0065\u006ca\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070",NewRelationship );
};