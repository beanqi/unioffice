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

package terms ;import (_d "encoding/xml";_a "fmt";_da "github.com/unidoc/unioffice";_aa "github.com/unidoc/unioffice/common/logger";_e "github.com/unidoc/unioffice/schema/purl.org/dc/elements";);

// Validate validates the ISO639_2 and its children
func (_afd *ISO639_2 )Validate ()error {return _afd .ValidateWithPath ("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032");};func (_dbg *LCSH )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_eac ,_bge :=d .Token ();if _bge !=nil {return _a .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073",_bge );
};if _acb ,_eafc :=_eac .(_d .EndElement );_eafc &&_acb .Name ==start .Name {break ;};};return nil ;};func (_efc *ElementsAndRefinementsGroup )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_fcg :for {_aaa ,_adg :=d .Token ();if _adg !=nil {return _adg ;
};switch _fg :=_aaa .(type ){case _d .StartElement :switch _fg .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_eaf :=NewElementsAndRefinementsGroupChoice ();
if _dea :=d .DecodeElement (&_eaf .Any ,&_fg );_dea !=nil {return _dea ;};_efc .Choice =append (_efc .Choice ,_eaf );default:_aa .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076",_fg .Name );
if _dae :=d .Skip ();_dae !=nil {return _dae ;};};case _d .EndElement :break _fcg ;case _d .CharData :};};return nil ;};

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_ffg *TGN )ValidateWithPath (path string )error {return nil };func (_dac *DDC )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0044\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });
return nil ;};func (_gfg *Period )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0050\u0065\u0072\u0069\u006f\u0064";e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};type LCSH struct{};
type ISO3166 struct{};

// Validate validates the UDC and its children
func (_geb *UDC )Validate ()error {return _geb .ValidateWithPath ("\u0055\u0044\u0043")};

// Validate validates the ISO3166 and its children
func (_df *ISO3166 )Validate ()error {return _df .ValidateWithPath ("\u0049S\u004f\u0033\u0031\u0036\u0036");};type Box struct{};func NewURI ()*URI {_cggd :=&URI {};return _cggd };type TGN struct{};type ISO639_2 struct{};

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_ffd *UDC )ValidateWithPath (path string )error {return nil };

// Validate validates the URI and its children
func (_bceb *URI )Validate ()error {return _bceb .ValidateWithPath ("\u0055\u0052\u0049")};

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_bce *MESH )ValidateWithPath (path string )error {return nil };func (_dg *Box )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0042\u006f\u0078";e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });
return nil ;};func (_c *Box )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_cf ,_g :=d .Token ();if _g !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073",_g );};if _adb ,_ge :=_cf .(_d .EndElement );
_ge &&_adb .Name ==start .Name {break ;};};return nil ;};

// Validate validates the LCC and its children
func (_ffc *LCC )Validate ()error {return _ffc .ValidateWithPath ("\u004c\u0043\u0043")};

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_cfc *ISO639_2 )ValidateWithPath (path string )error {return nil };func (_cbg *IMT )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0049\u004d\u0054";e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });
return nil ;};func (_bd *LCC )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u004c\u0043\u0043";e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_gfa *ElementsAndRefinementsGroupChoice )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _gfa .Any !=nil {_fba :=_d .StartElement {Name :_d .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};
for _ ,_fef :=range _gfa .Any {e .EncodeElement (_fef ,_fba );};};return nil ;};func NewDCMIType ()*DCMIType {_ac :=&DCMIType {};return _ac };type UDC struct{};func (_dade *LCSH )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u004c\u0043\u0053\u0048";
e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func NewElementOrRefinementContainer ()*ElementOrRefinementContainer {_fc :=&ElementOrRefinementContainer {};return _fc ;};

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_adbd *DDC )ValidateWithPath (path string )error {return nil };type DDC struct{};func (_ced *RFC1766 )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_eag ,_bbc :=d .Token ();if _bbc !=nil {return _a .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073",_bbc );
};if _gcf ,_cbb :=_eag .(_d .EndElement );_cbb &&_gcf .Name ==start .Name {break ;};};return nil ;};func (_bab *RFC1766 )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0052F\u0043\u0031\u0037\u0036\u0036";e .EncodeToken (start );
e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_dad *ISO639_2 )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_cbc ,_gce :=d .Token ();if _gce !=nil {return _a .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073",_gce );
};if _feae ,_fbee :=_cbc .(_d .EndElement );_fbee &&_feae .Name ==start .Name {break ;};};return nil ;};type DCMIType struct{};func (_cc *IMT )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_fbe ,_gegb :=d .Token ();if _gegb !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073",_gegb );
};if _bfe ,_ccb :=_fbe .(_d .EndElement );_ccb &&_bfe .Name ==start .Name {break ;};};return nil ;};

// Validate validates the TGN and its children
func (_efb *TGN )Validate ()error {return _efb .ValidateWithPath ("\u0054\u0047\u004e")};type W3CDTF struct{};

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_bcb *ElementsAndRefinementsGroupChoice )ValidateWithPath (path string )error {for _adbf ,_edg :=range _bcb .Any {if _ded :=_edg .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_adbf ));_ded !=nil {return _ded ;
};};return nil ;};func (_dd *ISO3166 )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_ffb ,_ada :=d .Token ();if _ada !=nil {return _a .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073",_ada );
};if _fge ,_bbb :=_ffb .(_d .EndElement );_bbb &&_fge .Name ==start .Name {break ;};};return nil ;};type RFC3066 struct{};func NewRFC1766 ()*RFC1766 {_afc :=&RFC1766 {};return _afc };func NewMESH ()*MESH {_eab :=&MESH {};return _eab };

// Validate validates the LCSH and its children
func (_gge *LCSH )Validate ()error {return _gge .ValidateWithPath ("\u004c\u0043\u0053\u0048")};type URI struct{};type ElementsAndRefinementsGroup struct{Choice []*ElementsAndRefinementsGroupChoice ;};

// Validate validates the W3CDTF and its children
func (_cff *W3CDTF )Validate ()error {return _cff .ValidateWithPath ("\u0057\u0033\u0043\u0044\u0054\u0046");};

// Validate validates the ElementsAndRefinementsGroup and its children
func (_edc *ElementsAndRefinementsGroup )Validate ()error {return _edc .ValidateWithPath ("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070");};

// Validate validates the Point and its children
func (_bcbb *Point )Validate ()error {return _bcbb .ValidateWithPath ("\u0050\u006f\u0069n\u0074")};func NewLCSH ()*LCSH {_cgg :=&LCSH {};return _cgg };func (_fac *LCC )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_dfc ,_beb :=d .Token ();
if _beb !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073",_beb );};if _bad ,_eda :=_dfc .(_d .EndElement );_eda &&_bad .Name ==start .Name {break ;};};return nil ;};

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_gd *ElementsAndRefinementsGroupChoice )Validate ()error {return _gd .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065");
};func NewElementsAndRefinementsGroup ()*ElementsAndRefinementsGroup {_bbg :=&ElementsAndRefinementsGroup {};return _bbg ;};

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_ggef *RFC3066 )ValidateWithPath (path string )error {return nil };func NewBox ()*Box {_ad :=&Box {};return _ad };func (_fab *URI )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0055\u0052\u0049";e .EncodeToken (start );
e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_fee *Point )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0050\u006f\u0069n\u0074";e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });
return nil ;};func (_edgf *TGN )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0054\u0047\u004e";e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};type ElementsAndRefinementsGroupChoice struct{Any []*_e .Any ;
};

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_cbf *LCSH )ValidateWithPath (path string )error {return nil };type ElementOrRefinementContainer struct{Choice []*ElementsAndRefinementsGroupChoice ;};type Period struct{};

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_bfd *W3CDTF )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_fae *RFC1766 )ValidateWithPath (path string )error {return nil };func (_dce *Period )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_fffg ,_efce :=d .Token ();if _efce !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073",_efce );
};if _cca ,_afa :=_fffg .(_d .EndElement );_afa &&_cca .Name ==start .Name {break ;};};return nil ;};func NewPoint ()*Point {_aaf :=&Point {};return _aaf };type LCC struct{};func NewRFC3066 ()*RFC3066 {_fce :=&RFC3066 {};return _fce };func (_faf *ISO639_2 )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032";
e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_bf *ElementsAndRefinementsGroup )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {if _bf .Choice !=nil {for _ ,_fb :=range _bf .Choice {_fb .MarshalXML (e ,_d .StartElement {});
};};return nil ;};

// Validate validates the MESH and its children
func (_efe *MESH )Validate ()error {return _efe .ValidateWithPath ("\u004d\u0045\u0053\u0048")};

// Validate validates the RFC1766 and its children
func (_bde *RFC1766 )Validate ()error {return _bde .ValidateWithPath ("\u0052F\u0043\u0031\u0037\u0036\u0036");};func (_gda *TGN )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_dcc ,_cd :=d .Token ();if _cd !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073",_cd );
};if _bfa ,_edce :=_dcc .(_d .EndElement );_edce &&_bfa .Name ==start .Name {break ;};};return nil ;};type IMT struct{};

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_daf *Point )ValidateWithPath (path string )error {return nil };func (_fadd *RFC3066 )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_agd ,_eaa :=d .Token ();if _eaa !=nil {return _a .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073",_eaa );
};if _dcd ,_cge :=_agd .(_d .EndElement );_cge &&_dcd .Name ==start .Name {break ;};};return nil ;};func (_cba *Point )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_ca ,_gea :=d .Token ();if _gea !=nil {return _a .Errorf ("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073",_gea );
};if _eedd ,_aed :=_ca .(_d .EndElement );_aed &&_eedd .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_ggg *Period )ValidateWithPath (path string )error {return nil };func (_feg *UDC )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0055\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });
return nil ;};

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_fad *IMT )ValidateWithPath (path string )error {return nil };func (_ef *DCMIType )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_gf ,_ee :=d .Token ();if _ee !=nil {return _a .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073",_ee );
};if _gef ,_f :=_gf .(_d .EndElement );_f &&_gef .Name ==start .Name {break ;};};return nil ;};func NewPeriod ()*Period {_ecf :=&Period {};return _ecf };

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_dga *URI )ValidateWithPath (path string )error {return nil };func NewDDC ()*DDC {_ff :=&DDC {};return _ff };

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_ab *DCMIType )ValidateWithPath (path string )error {return nil };

// Validate validates the ElementOrRefinementContainer and its children
func (_de *ElementOrRefinementContainer )Validate ()error {return _de .ValidateWithPath ("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072");};

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_bcf *ISO3166 )ValidateWithPath (path string )error {return nil };type Point struct{};func (_ece *ISO3166 )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0049S\u004f\u0033\u0031\u0036\u0036";e .EncodeToken (start );
e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_bga *MESH )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_fcge ,_agag :=d .Token ();if _agag !=nil {return _a .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073",_agag );
};if _dacd ,_dab :=_fcge .(_d .EndElement );_dab &&_dacd .Name ==start .Name {break ;};};return nil ;};func (_af *ElementOrRefinementContainer )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072";
e .EncodeToken (start );if _af .Choice !=nil {for _ ,_ag :=range _af .Choice {_ag .MarshalXML (e ,_d .StartElement {});};};e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};

// Validate validates the Box and its children
func (_gc *Box )Validate ()error {return _gc .ValidateWithPath ("\u0042\u006f\u0078")};

// Validate validates the RFC3066 and its children
func (_dcg *RFC3066 )Validate ()error {return _dcg .ValidateWithPath ("\u0052F\u0043\u0033\u0030\u0036\u0036");};func NewLCC ()*LCC {_fff :=&LCC {};return _fff };

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_acf *LCC )ValidateWithPath (path string )error {return nil };func (_dc *DDC )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_fe ,_db :=d .Token ();if _db !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073",_db );
};if _cb ,_eff :=_fe .(_d .EndElement );_eff &&_cb .Name ==start .Name {break ;};};return nil ;};

// Validate validates the DDC and its children
func (_ae *DDC )Validate ()error {return _ae .ValidateWithPath ("\u0044\u0044\u0043")};

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_efd *ElementsAndRefinementsGroup )ValidateWithPath (path string )error {for _aga ,_ga :=range _efd .Choice {if _eeb :=_ga .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_aga ));
_eeb !=nil {return _eeb ;};};return nil ;};func NewElementsAndRefinementsGroupChoice ()*ElementsAndRefinementsGroupChoice {_bc :=&ElementsAndRefinementsGroupChoice {};return _bc ;};func NewISO3166 ()*ISO3166 {_ccg :=&ISO3166 {};return _ccg };func NewTGN ()*TGN {_ead :=&TGN {};
return _ead };

// Validate validates the DCMIType and its children
func (_ed *DCMIType )Validate ()error {return _ed .ValidateWithPath ("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065");};func NewIMT ()*IMT {_acc :=&IMT {};return _acc };func NewUDC ()*UDC {_bbd :=&UDC {};return _bbd };func (_edaa *MESH )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u004d\u0045\u0053\u0048";
e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_bb *DCMIType )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065";e .EncodeToken (start );
e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_agb *UDC )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_abf ,_fgc :=d .Token ();if _fgc !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073",_fgc );
};if _dcf ,_ebd :=_abf .(_d .EndElement );_ebd &&_dcf .Name ==start .Name {break ;};};return nil ;};func NewW3CDTF ()*W3CDTF {_fggg :=&W3CDTF {};return _fggg };

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_bgf *ElementOrRefinementContainer )ValidateWithPath (path string )error {for _fea ,_geg :=range _bgf .Choice {if _eed :=_geg .ValidateWithPath (_a .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_fea ));
_eed !=nil {return _eed ;};};return nil ;};

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_eb *Box )ValidateWithPath (path string )error {return nil };type MESH struct{};func (_eg *URI )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_fgg ,_efa :=d .Token ();if _efa !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073",_efa );
};if _eade ,_dcce :=_fgg .(_d .EndElement );_dcce &&_eade .Name ==start .Name {break ;};};return nil ;};func (_bfb *RFC3066 )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0052F\u0043\u0033\u0030\u0036\u0036";e .EncodeToken (start );
e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};type RFC1766 struct{};func (_fa *ElementOrRefinementContainer )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_gb :for {_bg ,_fag :=d .Token ();if _fag !=nil {return _fag ;
};switch _be :=_bg .(type ){case _d .StartElement :switch _be .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_aad :=NewElementsAndRefinementsGroupChoice ();
if _ea :=d .DecodeElement (&_aad .Any ,&_be );_ea !=nil {return _ea ;};_fa .Choice =append (_fa .Choice ,_aad );default:_aa .Log .Debug ("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076",_be .Name );
if _dbd :=d .Skip ();_dbd !=nil {return _dbd ;};};case _d .EndElement :break _gb ;case _d .CharData :};};return nil ;};

// Validate validates the IMT and its children
func (_gg *IMT )Validate ()error {return _gg .ValidateWithPath ("\u0049\u004d\u0054")};func (_dba *ElementsAndRefinementsGroupChoice )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {_ec :for {_ce ,_cg :=d .Token ();if _cg !=nil {return _cg ;
};switch _aca :=_ce .(type ){case _d .StartElement :switch _aca .Name {case _d .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_eba :=_e .NewAny ();
if _ba :=d .DecodeElement (_eba ,&_aca );_ba !=nil {return _ba ;};_dba .Any =append (_dba .Any ,_eba );default:_aa .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076",_aca .Name );
if _eca :=d .Skip ();_eca !=nil {return _eca ;};};case _d .EndElement :break _ec ;case _d .CharData :};};return nil ;};func NewISO639_2 ()*ISO639_2 {_fgf :=&ISO639_2 {};return _fgf };func (_gab *W3CDTF )MarshalXML (e *_d .Encoder ,start _d .StartElement )error {start .Name .Local ="\u0057\u0033\u0043\u0044\u0054\u0046";
e .EncodeToken (start );e .EncodeToken (_d .EndElement {Name :start .Name });return nil ;};func (_dbdb *W3CDTF )UnmarshalXML (d *_d .Decoder ,start _d .StartElement )error {for {_eacg ,_bcff :=d .Token ();if _bcff !=nil {return _a .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073",_bcff );
};if _dca ,_fbb :=_eacg .(_d .EndElement );_fbb &&_dca .Name ==start .Name {break ;};};return nil ;};

// Validate validates the Period and its children
func (_bfea *Period )Validate ()error {return _bfea .ValidateWithPath ("\u0050\u0065\u0072\u0069\u006f\u0064");};func init (){_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0053\u0048",NewLCSH );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004d\u0045\u0053\u0048",NewMESH );_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0044\u0043",NewDDC );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0043",NewLCC );_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0044\u0043",NewUDC );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u0065\u0072\u0069\u006f\u0064",NewPeriod );_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0057\u0033\u0043\u0044\u0054\u0046",NewW3CDTF );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065",NewDCMIType );_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u004d\u0054",NewIMT );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0052\u0049",NewURI );_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032",NewISO639_2 );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0031\u0037\u0036\u0036",NewRFC1766 );_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0033\u0030\u0036\u0036",NewRFC3066 );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u006f\u0069n\u0074",NewPoint );_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049S\u004f\u0033\u0031\u0036\u0036",NewISO3166 );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0042\u006f\u0078",NewBox );_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0054\u0047\u004e",NewTGN );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072",NewElementOrRefinementContainer );
_da .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070",NewElementsAndRefinementsGroup );
};