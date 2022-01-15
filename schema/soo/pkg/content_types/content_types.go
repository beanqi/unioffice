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

package content_types ;import (_d "encoding/xml";_c "fmt";_a "github.com/unidoc/unioffice";_e "github.com/unidoc/unioffice/common/logger";_f "regexp";);

// ValidateWithPath validates the Types and its children, prefixing error messages with path
func (_bdg *Types )ValidateWithPath (path string )error {if _gbe :=_bdg .CT_Types .ValidateWithPath (path );_gbe !=nil {return _gbe ;};return nil ;};func (_eg *CT_Override )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_c .Sprintf ("\u0025\u0076",_eg .ContentTypeAttr )});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"},Value :_c .Sprintf ("\u0025\u0076",_eg .PartNameAttr )});e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};type Default struct{CT_Default };func (_fgd *CT_Types )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {e .EncodeToken (start );if _fgd .Default !=nil {_ee :=_d .StartElement {Name :_d .Name {Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}};for _ ,_cf :=range _fgd .Default {e .EncodeElement (_cf ,_ee );};};if _fgd .Override !=nil {_agf :=_d .StartElement {Name :_d .Name {Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}};for _ ,_ab :=range _fgd .Override {e .EncodeElement (_ab ,_agf );};};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_cdd *Override )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {return _cdd .CT_Override .MarshalXML (e ,start );};

// Validate validates the Override and its children
func (_cc *Override )Validate ()error {return _cc .ValidateWithPath ("\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};func NewDefault ()*Default {_ggg :=&Default {};_ggg .CT_Default =*NewCT_Default ();return _ggg };

// ValidateWithPath validates the CT_Types and its children, prefixing error messages with path
func (_badg *CT_Types )ValidateWithPath (path string )error {for _fea ,_gef :=range _badg .Default {if _egc :=_gef .ValidateWithPath (_c .Sprintf ("\u0025\u0073\u002f\u0044\u0065\u0066\u0061\u0075\u006ct\u005b\u0025\u0064\u005d",path ,_fea ));_egc !=nil {return _egc ;};};for _egcd ,_cbb :=range _badg .Override {if _bf :=_cbb .ValidateWithPath (_c .Sprintf ("\u0025s\u002fO\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u005b\u0025\u0064\u005d",path ,_egcd ));_bf !=nil {return _bf ;};};return nil ;};func (_def *CT_Override )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_def .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";for _ ,_bbfc :=range start .Attr {if _bbfc .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_ca ,_fbd :=_bbfc .Value ,error (nil );if _fbd !=nil {return _fbd ;};_def .ContentTypeAttr =_ca ;continue ;};if _bbfc .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_fec ,_dg :=_bbfc .Value ,error (nil );if _dg !=nil {return _dg ;};_def .PartNameAttr =_fec ;continue ;};};for {_cae ,_be :=d .Token ();if _be !=nil {return _c .Errorf ("\u0070\u0061\u0072si\u006e\u0067\u0020\u0043\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065\u003a\u0020\u0025\u0073",_be );};if _dge ,_bad :=_cae .(_d .EndElement );_bad &&_dge .Name ==start .Name {break ;};};return nil ;};func NewCT_Default ()*CT_Default {_ec :=&CT_Default {};_ec .ExtensionAttr ="\u0078\u006d\u006c";_ec .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _ec ;};func (_fc *Default )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {return _fc .CT_Default .MarshalXML (e ,start );};type CT_Default struct{ExtensionAttr string ;ContentTypeAttr string ;};func (_df *CT_Types )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_fee :for {_ecg ,_dfd :=d .Token ();if _dfd !=nil {return _dfd ;};switch _cb :=_ecg .(type ){case _d .StartElement :switch _cb .Name {case _d .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_aba :=NewDefault ();if _dca :=d .DecodeElement (_aba ,&_cb );_dca !=nil {return _dca ;};_df .Default =append (_df .Default ,_aba );case _d .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_ed :=NewOverride ();if _badf :=d .DecodeElement (_ed ,&_cb );_badf !=nil {return _badf ;};_df .Override =append (_df .Override ,_ed );default:_e .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073\u0020\u0025\u0076",_cb .Name );if _ef :=d .Skip ();_ef !=nil {return _ef ;};};case _d .EndElement :break _fee ;case _d .CharData :};};return nil ;};

// Validate validates the Default and its children
func (_edc *Default )Validate ()error {return _edc .ValidateWithPath ("\u0044e\u0066\u0061\u0075\u006c\u0074");};

// ValidateWithPath validates the Override and its children, prefixing error messages with path
func (_gf *Override )ValidateWithPath (path string )error {if _db :=_gf .CT_Override .ValidateWithPath (path );_db !=nil {return _db ;};return nil ;};func (_g *CT_Default )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_g .ExtensionAttr ="\u0078\u006d\u006c";_g .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";for _ ,_ge :=range start .Attr {if _ge .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_ac ,_fe :=_ge .Value ,error (nil );if _fe !=nil {return _fe ;};_g .ExtensionAttr =_ac ;continue ;};if _ge .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_acd ,_fb :=_ge .Value ,error (nil );if _fb !=nil {return _fb ;};_g .ContentTypeAttr =_acd ;continue ;};};for {_bb ,_bbf :=d .Token ();if _bbf !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020%\u0073",_bbf );};if _ag ,_gd :=_bb .(_d .EndElement );_gd &&_ag .Name ==start .Name {break ;};};return nil ;};type CT_Override struct{ContentTypeAttr string ;PartNameAttr string ;};const ST_ExtensionPattern ="\u0028\u005b\u0021\u0024\u0026\u0027\\\u0028\u005c\u0029\u005c\u002a\\\u002b\u002c\u003a\u003d\u005d\u007c(\u0025\u005b\u0030\u002d\u0039a\u002d\u0066\u0041\u002d\u0046\u005d\u005b\u0030\u002d\u0039\u0061\u002df\u0041\u002d\u0046\u005d\u0029\u007c\u005b\u003a\u0040\u005d\u007c\u005b\u0061\u002d\u007aA\u002d\u005a\u0030\u002d\u0039\u005c\u002d\u005f~\u005d\u0029\u002b";

// Validate validates the Types and its children
func (_gbd *Types )Validate ()error {return _gbd .ValidateWithPath ("\u0054\u0079\u0070e\u0073")};type Types struct{CT_Types };

// Validate validates the CT_Types and its children
func (_eaa *CT_Types )Validate ()error {return _eaa .ValidateWithPath ("\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073");};func NewCT_Types ()*CT_Types {_fg :=&CT_Types {};return _fg };func NewOverride ()*Override {_cd :=&Override {};_cd .CT_Override =*NewCT_Override ();return _cd };func NewCT_Override ()*CT_Override {_ba :=&CT_Override {};_ba .ContentTypeAttr ="\u0061p\u0070l\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u002f\u0078\u006d\u006c";return _ba ;};func (_cg *CT_Default )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"},Value :_c .Sprintf ("\u0025\u0076",_cg .ExtensionAttr )});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"},Value :_c .Sprintf ("\u0025\u0076",_cg .ContentTypeAttr )});e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_dcg *Types )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_dcg .CT_Types =*NewCT_Types ();_bbfb :for {_geff ,_bbb :=d .Token ();if _bbb !=nil {return _bbb ;};switch _fdb :=_geff .(type ){case _d .StartElement :switch _fdb .Name {case _d .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u0044e\u0066\u0061\u0075\u006c\u0074"}:_dfb :=NewDefault ();if _acc :=d .DecodeElement (_dfb ,&_fdb );_acc !=nil {return _acc ;};_dcg .Default =append (_dcg .Default ,_dfb );case _d .Name {Space :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s",Local :"\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065"}:_ggc :=NewOverride ();if _fa :=d .DecodeElement (_ggc ,&_fdb );_fa !=nil {return _fa ;};_dcg .Override =append (_dcg .Override ,_ggc );default:_e .Log .Debug ("s\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006fn \u0054\u0079\u0070e\u0073 \u0025\u0076",_fdb .Name );if _fcee :=d .Skip ();_fcee !=nil {return _fcee ;};};case _d .EndElement :break _bbfb ;case _d .CharData :};};return nil ;};func NewTypes ()*Types {_da :=&Types {};_da .CT_Types =*NewCT_Types ();return _da };func (_fce *Default )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_fce .CT_Default =*NewCT_Default ();for _ ,_dfa :=range start .Attr {if _dfa .Name .Local =="\u0045x\u0074\u0065\u006e\u0073\u0069\u006fn"{_abac ,_fd :=_dfa .Value ,error (nil );if _fd !=nil {return _fd ;};_fce .ExtensionAttr =_abac ;continue ;};if _dfa .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_bfa ,_eec :=_dfa .Value ,error (nil );if _eec !=nil {return _eec ;};_fce .ContentTypeAttr =_bfa ;continue ;};};for {_bd ,_bc :=d .Token ();if _bc !=nil {return _c .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0044\u0065\u0066\u0061\u0075\u006c\u0074\u003a\u0020\u0025\u0073",_bc );};if _eee ,_aa :=_bd .(_d .EndElement );_aa &&_eee .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the CT_Override and its children, prefixing error messages with path
func (_gg *CT_Override )ValidateWithPath (path string )error {if !ST_ContentTypePatternRe .MatchString (_gg .ContentTypeAttr ){return _c .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_gg .ContentTypeAttr );};return nil ;};func (_agb *Types )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s"});start .Attr =append (start .Attr ,_d .Attr {Name :_d .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0054\u0079\u0070e\u0073";return _agb .CT_Types .MarshalXML (e ,start );};const ST_ContentTypePattern ="\u005e\\\u0070{\u004c\u0061\u0074\u0069\u006e\u007d\u002b\u002f\u002e\u002a\u0024";type Override struct{CT_Override };

// Validate validates the CT_Default and its children
func (_cgf *CT_Default )Validate ()error {return _cgf .ValidateWithPath ("\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074");};

// Validate validates the CT_Override and its children
func (_ea *CT_Override )Validate ()error {return _ea .ValidateWithPath ("C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065");};var ST_ContentTypePatternRe =_f .MustCompile (ST_ContentTypePattern );type CT_Types struct{Default []*Default ;Override []*Override ;};

// ValidateWithPath validates the Default and its children, prefixing error messages with path
func (_abg *Default )ValidateWithPath (path string )error {if _fbdg :=_abg .CT_Default .ValidateWithPath (path );_fbdg !=nil {return _fbdg ;};return nil ;};

// ValidateWithPath validates the CT_Default and its children, prefixing error messages with path
func (_de *CT_Default )ValidateWithPath (path string )error {if !ST_ExtensionPatternRe .MatchString (_de .ExtensionAttr ){return _c .Errorf ("\u0025s\u002f\u006d.\u0045\u0078\u0074\u0065n\u0073\u0069\u006fn\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074 m\u0061\u0074\u0063h\u0020\u0027%\u0073\u0027\u0020\u0028\u0068\u0061v\u0065\u0020%\u0076\u0029",path ,ST_ExtensionPatternRe ,_de .ExtensionAttr );};if !ST_ContentTypePatternRe .MatchString (_de .ContentTypeAttr ){return _c .Errorf ("\u0025\u0073/\u006d\u002e\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065\u0041\u0074\u0074\u0072\u0020\u006d\u0075\u0073\u0074\u0020\u006d\u0061\u0074\u0063\u0068\u0020\u0027\u0025\u0073\u0027\u0020\u0028\u0068\u0061\u0076\u0065\u0020\u0025\u0076\u0029",path ,ST_ContentTypePatternRe ,_de .ContentTypeAttr );};return nil ;};func (_cad *Override )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_cad .CT_Override =*NewCT_Override ();for _ ,_fecf :=range start .Attr {if _fecf .Name .Local =="C\u006f\u006e\u0074\u0065\u006e\u0074\u0054\u0079\u0070\u0065"{_ga ,_fead :=_fecf .Value ,error (nil );if _fead !=nil {return _fead ;};_cad .ContentTypeAttr =_ga ;continue ;};if _fecf .Name .Local =="\u0050\u0061\u0072\u0074\u004e\u0061\u006d\u0065"{_fbdge ,_dff :=_fecf .Value ,error (nil );if _dff !=nil {return _dff ;};_cad .PartNameAttr =_fbdge ;continue ;};};for {_caf ,_cbf :=d .Token ();if _cbf !=nil {return _c .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u004f\u0076\u0065r\u0072\u0069\u0064\u0065: \u0025\u0073",_cbf );};if _ce ,_ged :=_caf .(_d .EndElement );_ged &&_ce .Name ==start .Name {break ;};};return nil ;};var ST_ExtensionPatternRe =_f .MustCompile (ST_ExtensionPattern );func init (){_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0054\u0079\u0070\u0065\u0073",NewCT_Types );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0043\u0054\u005f\u0044\u0065\u0066\u0061\u0075\u006c\u0074",NewCT_Default );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","C\u0054\u005f\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewCT_Override );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0054\u0079\u0070e\u0073",NewTypes );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u0044e\u0066\u0061\u0075\u006c\u0074",NewDefault );_a .RegisterConstructor ("ht\u0074\u0070:\u002f\u002f\u0073\u0063\u0068\u0065\u006d\u0061\u0073.\u006f\u0070\u0065\u006e\u0078\u006d\u006c\u0066\u006f\u0072\u006d\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u0070\u0061\u0063\u006b\u0061\u0067\u0065\u002f\u00320\u00306\u002f\u0063\u006f\u006e\u0074\u0065\u006e\u0074-\u0074y\u0070\u0065s","\u004f\u0076\u0065\u0072\u0072\u0069\u0064\u0065",NewOverride );};