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

package terms ;import (_c "encoding/xml";_f "fmt";_dg "github.com/unidoc/unioffice";_d "github.com/unidoc/unioffice/common/logger";_fb "github.com/unidoc/unioffice/schema/purl.org/dc/elements";);func (_bafg *LCSH )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u004c\u0043\u0053\u0048";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_ccf *MESH )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_bda ,_aab :=d .Token ();if _aab !=nil {return _f .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004d\u0045\u0053\u0048\u003a\u0020\u0025\u0073",_aab );};if _age ,_fed :=_bda .(_c .EndElement );_fed &&_age .Name ==start .Name {break ;};};return nil ;};func NewUDC ()*UDC {_edgf :=&UDC {};return _edgf };

// Validate validates the IMT and its children
func (_abg *IMT )Validate ()error {return _abg .ValidateWithPath ("\u0049\u004d\u0054")};func (_fg *ISO3166 )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0049S\u004f\u0033\u0031\u0036\u0036";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// Validate validates the RFC1766 and its children
func (_gfff *RFC1766 )Validate ()error {return _gfff .ValidateWithPath ("\u0052F\u0043\u0031\u0037\u0036\u0036");};func (_fd *DCMIType )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_egg *RFC3066 )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0052F\u0043\u0033\u0030\u0036\u0036";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_da *ElementsAndRefinementsGroupChoice )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_ace :for {_eag ,_aff :=d .Token ();if _aff !=nil {return _aff ;};switch _dc :=_eag .(type ){case _c .StartElement :switch _dc .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_gg :=_fb .NewAny ();if _efg :=d .DecodeElement (_gg ,&_dc );_efg !=nil {return _efg ;};_da .Any =append (_da .Any ,_gg );default:_d .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0041\u006ed\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006fu\u0070\u0043\u0068o\u0069\u0063\u0065\u0020\u0025\u0076",_dc .Name );if _fba :=d .Skip ();_fba !=nil {return _fba ;};};case _c .EndElement :break _ace ;case _c .CharData :};};return nil ;};func (_ea *ElementsAndRefinementsGroupChoice )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _ea .Any !=nil {_fe :=_c .StartElement {Name :_c .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};for _ ,_gcc :=range _ea .Any {e .EncodeElement (_gcc ,_fe );};};return nil ;};func (_ccfg *W3CDTF )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_gdff ,_gfd :=d .Token ();if _gfd !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u00573\u0043\u0044T\u0046\u003a\u0020\u0025\u0073",_gfd );};if _cba ,_bgfa :=_gdff .(_c .EndElement );_bgfa &&_cba .Name ==start .Name {break ;};};return nil ;};

// Validate validates the DDC and its children
func (_df *DDC )Validate ()error {return _df .ValidateWithPath ("\u0044\u0044\u0043")};

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (_fbg *DDC )ValidateWithPath (path string )error {return nil };

// Validate validates the Period and its children
func (_dgd *Period )Validate ()error {return _dgd .ValidateWithPath ("\u0050\u0065\u0072\u0069\u006f\u0064");};

// Validate validates the Point and its children
func (_bbga *Point )Validate ()error {return _bbga .ValidateWithPath ("\u0050\u006f\u0069n\u0074")};func NewDCMIType ()*DCMIType {_g :=&DCMIType {};return _g };func NewElementOrRefinementContainer ()*ElementOrRefinementContainer {_ae :=&ElementOrRefinementContainer {};return _ae ;};type ISO3166 struct{};func (_gf *DDC )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0044\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_ffa *UDC )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0055\u0044\u0043";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_edg *LCSH )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_bgf ,_dbc :=d .Token ();if _dbc !=nil {return _f .Errorf ("\u0070\u0061r\u0073\u0069\u006eg\u0020\u004c\u0043\u0053\u0048\u003a\u0020\u0025\u0073",_dbc );};if _ead ,_dfaf :=_bgf .(_c .EndElement );_dfaf &&_ead .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the ISO639_2 and its children, prefixing error messages with path
func (_gaf *ISO639_2 )ValidateWithPath (path string )error {return nil };func (_gbde *TGN )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_fgf ,_ddca :=d .Token ();if _ddca !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0054\u0047\u004e\u003a\u0020\u0025\u0073",_ddca );};if _ebg ,_gdb :=_fgf .(_c .EndElement );_gdb &&_ebg .Name ==start .Name {break ;};};return nil ;};func NewPoint ()*Point {_ceb :=&Point {};return _ceb };func NewRFC3066 ()*RFC3066 {_bdc :=&RFC3066 {};return _bdc };func NewMESH ()*MESH {_edd :=&MESH {};return _edd };

// Validate validates the MESH and its children
func (_acg *MESH )Validate ()error {return _acg .ValidateWithPath ("\u004d\u0045\u0053\u0048")};func (_de *Period )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0050\u0065\u0072\u0069\u006f\u0064";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type LCSH struct{};func (_bgc *ISO639_2 )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_fca ,_cg :=d .Token ();if _cg !=nil {return _f .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0049\u0053\u004f6\u0033\u0039\u005f\u0032: \u0025\u0073",_cg );};if _aef ,_ccg :=_fca .(_c .EndElement );_ccg &&_aef .Name ==start .Name {break ;};};return nil ;};

// Validate validates the DCMIType and its children
func (_ag *DCMIType )Validate ()error {return _ag .ValidateWithPath ("\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065");};func (_aad *URI )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_gac ,_fedb :=d .Token ();if _fedb !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0052\u0049\u003a\u0020\u0025\u0073",_fedb );};if _aaf ,_gcee :=_gac .(_c .EndElement );_gcee &&_aaf .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the RFC3066 and its children, prefixing error messages with path
func (_efgg *RFC3066 )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the LCC and its children, prefixing error messages with path
func (_fac *LCC )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the Point and its children, prefixing error messages with path
func (_dee *Point )ValidateWithPath (path string )error {return nil };func (_dbb *LCC )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u004c\u0043\u0043";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the TGN and its children, prefixing error messages with path
func (_ebd *TGN )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the ElementOrRefinementContainer and its children, prefixing error messages with path
func (_agg *ElementOrRefinementContainer )ValidateWithPath (path string )error {for _dgfb ,_eg :=range _agg .Choice {if _bf :=_eg .ValidateWithPath (_f .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_dgfb ));_bf !=nil {return _bf ;};};return nil ;};func (_cd *DDC )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_dge ,_bg :=d .Token ();if _bg !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0044\u0044\u0043\u003a\u0020\u0025\u0073",_bg );};if _dda ,_dgf :=_dge .(_c .EndElement );_dgf &&_dda .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the UDC and its children, prefixing error messages with path
func (_cfe *UDC )ValidateWithPath (path string )error {return nil };func NewElementsAndRefinementsGroupChoice ()*ElementsAndRefinementsGroupChoice {_ga :=&ElementsAndRefinementsGroupChoice {};return _ga ;};

// Validate validates the ISO3166 and its children
func (_cf *ISO3166 )Validate ()error {return _cf .ValidateWithPath ("\u0049S\u004f\u0033\u0031\u0036\u0036");};func (_bggb *ElementsAndRefinementsGroup )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _bggb .Choice !=nil {for _ ,_ab :=range _bggb .Choice {_ab .MarshalXML (e ,_c .StartElement {});};};return nil ;};type TGN struct{};func (_ad *ElementOrRefinementContainer )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072";e .EncodeToken (start );if _ad .Choice !=nil {for _ ,_ff :=range _ad .Choice {_ff .MarshalXML (e ,_c .StartElement {});};};e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_dgfc *Point )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0050\u006f\u0069n\u0074";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the RFC1766 and its children, prefixing error messages with path
func (_acc *RFC1766 )ValidateWithPath (path string )error {return nil };

// Validate validates the UDC and its children
func (_fdg *UDC )Validate ()error {return _fdg .ValidateWithPath ("\u0055\u0044\u0043")};type ISO639_2 struct{};type IMT struct{};func NewDDC ()*DDC {_fbd :=&DDC {};return _fbd };func (_fda *ElementOrRefinementContainer )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_bge :for {_dgc ,_bgg :=d .Token ();if _bgg !=nil {return _bgg ;};switch _gb :=_dgc .(type ){case _c .StartElement :switch _gb .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_bgb :=NewElementsAndRefinementsGroupChoice ();if _ffe :=d .DecodeElement (&_bgb .Any ,&_gb );_ffe !=nil {return _ffe ;};_fda .Choice =append (_fda .Choice ,_bgb );default:_d .Log .Debug ("\u0073k\u0069\u0070\u0070\u0069\u006e\u0067\u0020un\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006de\u006e\u0074 \u006f\u006e\u0020E\u006c\u0065\u006d\u0065\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065n\u0074\u0043on\u0074\u0061\u0069n\u0065\u0072\u0020\u0025\u0076",_gb .Name );if _ada :=d .Skip ();_ada !=nil {return _ada ;};};case _c .EndElement :break _bge ;case _c .CharData :};};return nil ;};

// Validate validates the RFC3066 and its children
func (_edc *RFC3066 )Validate ()error {return _edc .ValidateWithPath ("\u0052F\u0043\u0033\u0030\u0036\u0036");};func (_ddb *Point )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_bce ,_aeag :=d .Token ();if _aeag !=nil {return _f .Errorf ("\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0050\u006f\u0069\u006et\u003a\u0020\u0025\u0073",_aeag );};if _fag ,_cdc :=_bce .(_c .EndElement );_cdc &&_fag .Name ==start .Name {break ;};};return nil ;};func NewBox ()*Box {_fc :=&Box {};return _fc };

// ValidateWithPath validates the ISO3166 and its children, prefixing error messages with path
func (_dbf *ISO3166 )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the DCMIType and its children, prefixing error messages with path
func (_gc *DCMIType )ValidateWithPath (path string )error {return nil };

// ValidateWithPath validates the LCSH and its children, prefixing error messages with path
func (_ddc *LCSH )ValidateWithPath (path string )error {return nil };

// Validate validates the TGN and its children
func (_ceag *TGN )Validate ()error {return _ceag .ValidateWithPath ("\u0054\u0047\u004e")};func (_fff *TGN )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0054\u0047\u004e";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type DCMIType struct{};

// ValidateWithPath validates the IMT and its children, prefixing error messages with path
func (_ge *IMT )ValidateWithPath (path string )error {return nil };func (_gcce *UDC )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_aagd ,_ddag :=d .Token ();if _ddag !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0055\u0044\u0043\u003a\u0020\u0025\u0073",_ddag );};if _afb ,_ded :=_aagd .(_c .EndElement );_ded &&_afb .Name ==start .Name {break ;};};return nil ;};type ElementOrRefinementContainer struct{Choice []*ElementsAndRefinementsGroupChoice ;};func (_gbdg *RFC1766 )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0052F\u0043\u0031\u0037\u0036\u0036";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type Box struct{};func NewPeriod ()*Period {_edgb :=&Period {};return _edgb };type UDC struct{};func NewIMT ()*IMT {_dgbd :=&IMT {};return _dgbd };

// Validate validates the URI and its children
func (_affd *URI )Validate ()error {return _affd .ValidateWithPath ("\u0055\u0052\u0049")};func (_gff *RFC1766 )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_gdfc ,_fgc :=d .Token ();if _fgc !=nil {return _f .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0031\u0037\u0036\u0036\u003a\u0020\u0025\u0073",_fgc );};if _eda ,_gaa :=_gdfc .(_c .EndElement );_gaa &&_eda .Name ==start .Name {break ;};};return nil ;};func NewTGN ()*TGN {_gga :=&TGN {};return _gga };

// Validate validates the ISO639_2 and its children
func (_cedf *ISO639_2 )Validate ()error {return _cedf .ValidateWithPath ("\u0049\u0053\u004f\u0036\u0033\u0039\u005f\u0032");};func NewRFC1766 ()*RFC1766 {_dgfcf :=&RFC1766 {};return _dgfcf };func NewLCC ()*LCC {_baf :=&LCC {};return _baf };

// Validate validates the W3CDTF and its children
func (_cbe *W3CDTF )Validate ()error {return _cbe .ValidateWithPath ("\u0057\u0033\u0043\u0044\u0054\u0046");};

// Validate validates the ElementsAndRefinementsGroupChoice and its children
func (_ddf *ElementsAndRefinementsGroupChoice )Validate ()error {return _ddf .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006et\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0043h\u006f\u0069\u0063\u0065");};func NewURI ()*URI {_gab :=&URI {};return _gab };type RFC1766 struct{};

// ValidateWithPath validates the W3CDTF and its children, prefixing error messages with path
func (_aaa *W3CDTF )ValidateWithPath (path string )error {return nil };type Point struct{};

// Validate validates the ElementOrRefinementContainer and its children
func (_agd *ElementOrRefinementContainer )Validate ()error {return _agd .ValidateWithPath ("\u0045\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072");};func NewISO639_2 ()*ISO639_2 {_acd :=&ISO639_2 {};return _acd };type ElementsAndRefinementsGroupChoice struct{Any []*_fb .Any ;};type ElementsAndRefinementsGroup struct{Choice []*ElementsAndRefinementsGroupChoice ;};

// ValidateWithPath validates the Box and its children, prefixing error messages with path
func (_dd *Box )ValidateWithPath (path string )error {return nil };type LCC struct{};func (_bbg *LCC )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_afg ,_dfa :=d .Token ();if _dfa !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u004c\u0043\u0043\u003a\u0020\u0025\u0073",_dfa );};if _gcg ,_ccgc :=_afg .(_c .EndElement );_ccgc &&_gcg .Name ==start .Name {break ;};};return nil ;};func (_bb *Box )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_ce ,_bba :=d .Token ();if _bba !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0042\u006f\u0078\u003a\u0020\u0025\u0073",_bba );};if _e ,_db :=_ce .(_c .EndElement );_db &&_e .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the Period and its children, prefixing error messages with path
func (_cea *Period )ValidateWithPath (path string )error {return nil };func (_ac *ElementsAndRefinementsGroup )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_aa :for {_bbe ,_cc :=d .Token ();if _cc !=nil {return _cc ;};switch _cdg :=_bbe .(type ){case _c .StartElement :switch _cdg .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_fcc :=NewElementsAndRefinementsGroupChoice ();if _fbf :=d .DecodeElement (&_fcc .Any ,&_cdg );_fbf !=nil {return _fbf ;};_ac .Choice =append (_ac .Choice ,_fcc );default:_d .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074ed\u0020e\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006ce\u006d\u0065\u006e\u0074\u0073\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065\u006d\u0065\u006et\u0073\u0047\u0072\u006f\u0075\u0070\u0020\u0025\u0076",_cdg .Name );if _fdc :=d .Skip ();_fdc !=nil {return _fdc ;};};case _c .EndElement :break _aa ;case _c .CharData :};};return nil ;};func (_bc *IMT )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_aag ,_gdf :=d .Token ();if _gdf !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0049\u004d\u0054\u003a\u0020\u0025\u0073",_gdf );};if _dcg ,_bcb :=_aag .(_c .EndElement );_bcb &&_dcg .Name ==start .Name {break ;};};return nil ;};

// Validate validates the ElementsAndRefinementsGroup and its children
func (_egd *ElementsAndRefinementsGroup )Validate ()error {return _egd .ValidateWithPath ("E\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070");};func (_gge *Period )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_bgce ,_dca :=d .Token ();if _dca !=nil {return _f .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0050e\u0072\u0069o\u0064\u003a\u0020\u0025\u0073",_dca );};if _fdd ,_gae :=_bgce .(_c .EndElement );_gae &&_fdd .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the ElementsAndRefinementsGroupChoice and its children, prefixing error messages with path
func (_dgfg *ElementsAndRefinementsGroupChoice )ValidateWithPath (path string )error {for _gbf ,_gd :=range _dgfg .Any {if _ec :=_gd .ValidateWithPath (_f .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_gbf ));_ec !=nil {return _ec ;};};return nil ;};func NewElementsAndRefinementsGroup ()*ElementsAndRefinementsGroup {_gbd :=&ElementsAndRefinementsGroup {};return _gbd ;};func (_gcb *W3CDTF )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0057\u0033\u0043\u0044\u0054\u0046";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// Validate validates the LCSH and its children
func (_dgcd *LCSH )Validate ()error {return _dgcd .ValidateWithPath ("\u004c\u0043\u0053\u0048")};

// Validate validates the LCC and its children
func (_fgd *LCC )Validate ()error {return _fgd .ValidateWithPath ("\u004c\u0043\u0043")};

// Validate validates the Box and its children
func (_ee *Box )Validate ()error {return _ee .ValidateWithPath ("\u0042\u006f\u0078")};type RFC3066 struct{};func (_ca *Box )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0042\u006f\u0078";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_efd *URI )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0055\u0052\u0049";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the MESH and its children, prefixing error messages with path
func (_ccd *MESH )ValidateWithPath (path string )error {return nil };func NewW3CDTF ()*W3CDTF {_ged :=&W3CDTF {};return _ged };type Period struct{};

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (_af *ElementsAndRefinementsGroup )ValidateWithPath (path string )error {for _gce ,_gfa :=range _af .Choice {if _dfg :=_gfa .ValidateWithPath (_f .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_gce ));_dfg !=nil {return _dfg ;};};return nil ;};func NewLCSH ()*LCSH {_aeae :=&LCSH {};return _aeae };func (_feb *ISO639_2 )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type W3CDTF struct{};func (_fec *MESH )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u004d\u0045\u0053\u0048";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_aea *ISO3166 )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_daf ,_bff :=d .Token ();if _bff !=nil {return _f .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0049\u0053\u004f\u0033\u0031\u0036\u0036\u003a\u0020\u0025\u0073",_bff );};if _bd ,_ced :=_daf .(_c .EndElement );_ced &&_bd .Name ==start .Name {break ;};};return nil ;};type DDC struct{};type URI struct{};func (_dgb *DCMIType )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_a ,_ef :=d .Token ();if _ef !=nil {return _f .Errorf ("p\u0061r\u0073\u0069\u006e\u0067\u0020\u0044\u0043\u004dI\u0054\u0079\u0070\u0065: \u0025\u0073",_ef );};if _ba ,_fa :=_a .(_c .EndElement );_fa &&_ba .Name ==start .Name {break ;};};return nil ;};type MESH struct{};func (_ggc *RFC3066 )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for {_fcd ,_bgbg :=d .Token ();if _bgbg !=nil {return _f .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0052\u0046\u0043\u0033\u0030\u0036\u0036\u003a\u0020\u0025\u0073",_bgbg );};if _eb ,_baa :=_fcd .(_c .EndElement );_baa &&_eb .Name ==start .Name {break ;};};return nil ;};

// ValidateWithPath validates the URI and its children, prefixing error messages with path
func (_bdd *URI )ValidateWithPath (path string )error {return nil };func (_dbd *IMT )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Name .Local ="\u0049\u004d\u0054";e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func NewISO3166 ()*ISO3166 {_cb :=&ISO3166 {};return _cb };func init (){_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0053\u0048",NewLCSH );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004d\u0045\u0053\u0048",NewMESH );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0044\u0043",NewDDC );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u004c\u0043\u0043",NewLCC );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0044\u0043",NewUDC );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u0065\u0072\u0069\u006f\u0064",NewPeriod );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0057\u0033\u0043\u0044\u0054\u0046",NewW3CDTF );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0044\u0043\u004d\u0049\u0054\u0079\u0070\u0065",NewDCMIType );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u004d\u0054",NewIMT );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0055\u0052\u0049",NewURI );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049\u0053\u004f\u0036\u0033\u0039\u002d\u0032",NewISO639_2 );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0031\u0037\u0036\u0036",NewRFC1766 );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0052F\u0043\u0033\u0030\u0036\u0036",NewRFC3066 );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0050\u006f\u0069n\u0074",NewPoint );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0049S\u004f\u0033\u0031\u0036\u0036",NewISO3166 );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0042\u006f\u0078",NewBox );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0054\u0047\u004e",NewTGN );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","\u0065\u006c\u0065\u006de\u006e\u0074\u004f\u0072\u0052\u0065\u0066\u0069\u006e\u0065m\u0065n\u0074\u0043\u006f\u006e\u0074\u0061\u0069n\u0065\u0072",NewElementOrRefinementContainer );_dg .RegisterConstructor ("\u0068t\u0074\u0070\u003a\u002f/\u0070\u0075\u0072\u006c\u002eo\u0072g\u002fd\u0063\u002f\u0074\u0065\u0072\u006d\u0073/","e\u006c\u0065\u006d\u0065\u006e\u0074s\u0041\u006e\u0064\u0052\u0065\u0066\u0069\u006e\u0065m\u0065\u006e\u0074s\u0047r\u006f\u0075\u0070",NewElementsAndRefinementsGroup );};