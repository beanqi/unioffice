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

package activeX ;import (_g "encoding/xml";_fc "fmt";_b "github.com/unidoc/unioffice";_fe "github.com/unidoc/unioffice/common/logger";);func (_ff *CT_Font )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {if _ff .PersistenceAttr !=ST_PersistenceUnset {_a ,_gab :=_ff .PersistenceAttr .MarshalXMLAttr (_g .Name {Local :"\u0061\u0078\u003a\u0070\u0065\u0072\u0073\u0069\u0073t\u0065\u006e\u0063\u0065"});if _gab !=nil {return _gab ;};start .Attr =append (start .Attr ,_a );};if _ff .IdAttr !=nil {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_fc .Sprintf ("\u0025\u0076",*_ff .IdAttr )});};e .EncodeToken (start );if _ff .OcxPr !=nil {_ffc :=_g .StartElement {Name :_g .Name {Local :"\u0061\u0078\u003a\u006f\u0063\u0078\u0050\u0072"}};for _ ,_c :=range _ff .OcxPr {e .EncodeElement (_c ,_ffc );};};e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};

// Validate validates the CT_OcxPr and its children
func (_ffg *CT_OcxPr )Validate ()error {return _ffg .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072");};func (_gd *CT_Ocx )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0061\u0078\u003a\u0063\u006c\u0061\u0073\u0073\u0069\u0064"},Value :_fc .Sprintf ("\u0025\u0076",_gd .ClassidAttr )});if _gd .LicenseAttr !=nil {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0061\u0078\u003a\u006c\u0069\u0063\u0065\u006e\u0073\u0065"},Value :_fc .Sprintf ("\u0025\u0076",*_gd .LicenseAttr )});};if _gd .IdAttr !=nil {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_fc .Sprintf ("\u0025\u0076",*_gd .IdAttr )});};_edaa ,_bcb :=_gd .PersistenceAttr .MarshalXMLAttr (_g .Name {Local :"\u0061\u0078\u003a\u0070\u0065\u0072\u0073\u0069\u0073t\u0065\u006e\u0063\u0065"});if _bcb !=nil {return _bcb ;};start .Attr =append (start .Attr ,_edaa );e .EncodeToken (start );if _gd .OcxPr !=nil {_gce :=_g .StartElement {Name :_g .Name {Local :"\u0061\u0078\u003a\u006f\u0063\u0078\u0050\u0072"}};for _ ,_bf :=range _gd .OcxPr {e .EncodeElement (_bf ,_gce );};};e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the Ocx and its children, prefixing error messages with path
func (_bgf *Ocx )ValidateWithPath (path string )error {if _feae :=_bgf .CT_Ocx .ValidateWithPath (path );_feae !=nil {return _feae ;};return nil ;};func NewCT_Picture ()*CT_Picture {_dc :=&CT_Picture {};return _dc };type CT_OcxPr struct{NameAttr string ;ValueAttr *string ;Choice *CT_OcxPrChoice ;};func (_db *CT_Picture )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {if _db .IdAttr !=nil {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_fc .Sprintf ("\u0025\u0076",*_db .IdAttr )});};e .EncodeToken (start );e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};func (_bae *CT_OcxPrChoice )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_eeb :for {_fce ,_eddc :=d .Token ();if _eddc !=nil {return _eddc ;};switch _dga :=_fce .(type ){case _g .StartElement :switch _dga .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0066\u006f\u006e\u0074"}:_bae .Font =NewCT_Font ();if _ad :=d .DecodeElement (_bae .Font ,&_dga );_ad !=nil {return _ad ;};case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0070i\u0063\u0074\u0075\u0072\u0065"}:_bae .Picture =NewCT_Picture ();if _fea :=d .DecodeElement (_bae .Picture ,&_dga );_fea !=nil {return _fea ;};default:_fe .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0043\u0068o\u0069c\u0065\u0020\u0025\u0076",_dga .Name );if _cc :=d .Skip ();_cc !=nil {return _cc ;};};case _g .EndElement :break _eeb ;case _g .CharData :};};return nil ;};func NewCT_OcxPrChoice ()*CT_OcxPrChoice {_ag :=&CT_OcxPrChoice {};return _ag };func (_gag *CT_OcxPr )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0061x\u003a\u006e\u0061\u006d\u0065"},Value :_fc .Sprintf ("\u0025\u0076",_gag .NameAttr )});if _gag .ValueAttr !=nil {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0061\u0078\u003a\u0076\u0061\u006c\u0075\u0065"},Value :_fc .Sprintf ("\u0025\u0076",*_gag .ValueAttr )});};e .EncodeToken (start );if _gag .Choice !=nil {_gag .Choice .MarshalXML (e ,_g .StartElement {});};e .EncodeToken (_g .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the CT_OcxPr and its children, prefixing error messages with path
func (_dfb *CT_OcxPr )ValidateWithPath (path string )error {if _dfb .Choice !=nil {if _abd :=_dfb .Choice .ValidateWithPath (path +"\u002fC\u0068\u006f\u0069\u0063\u0065");_abd !=nil {return _abd ;};};return nil ;};

// Validate validates the CT_Ocx and its children
func (_edd *CT_Ocx )Validate ()error {return _edd .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078");};func NewCT_OcxPr ()*CT_OcxPr {_fca :=&CT_OcxPr {};return _fca };func (_fb *CT_OcxPr )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for _ ,_dd :=range start .Attr {if _dd .Name .Local =="\u006e\u0061\u006d\u0065"{_dgd ,_dde :=_dd .Value ,error (nil );if _dde !=nil {return _dde ;};_fb .NameAttr =_dgd ;continue ;};if _dd .Name .Local =="\u0076\u0061\u006cu\u0065"{_caf ,_gfc :=_dd .Value ,error (nil );if _gfc !=nil {return _gfc ;};_fb .ValueAttr =&_caf ;continue ;};};_ea :for {_fbc ,_ded :=d .Token ();if _ded !=nil {return _ded ;};switch _bb :=_fbc .(type ){case _g .StartElement :switch _bb .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0066\u006f\u006e\u0074"}:_fb .Choice =NewCT_OcxPrChoice ();if _efe :=d .DecodeElement (&_fb .Choice .Font ,&_bb );_efe !=nil {return _efe ;};case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0070i\u0063\u0074\u0075\u0072\u0065"}:_fb .Choice =NewCT_OcxPrChoice ();if _gcg :=d .DecodeElement (&_fb .Choice .Picture ,&_bb );_gcg !=nil {return _gcg ;};default:_fe .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0020\u0025\u0076",_bb .Name );if _gdc :=d .Skip ();_gdc !=nil {return _gdc ;};};case _g .EndElement :break _ea ;case _g .CharData :};};return nil ;};func NewCT_Ocx ()*CT_Ocx {_bc :=&CT_Ocx {};_bc .PersistenceAttr =ST_Persistence (1);return _bc };func (_abb *ST_Persistence )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_ddf ,_aaab :=d .Token ();if _aaab !=nil {return _aaab ;};if _fff ,_ddeb :=_ddf .(_g .EndElement );_ddeb &&_fff .Name ==start .Name {*_abb =1;return nil ;};if _eff ,_aag :=_ddf .(_g .CharData );!_aag {return _fc .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_ddf );}else {switch string (_eff ){case "":*_abb =0;case "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067":*_abb =1;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d":*_abb =2;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074":*_abb =3;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065":*_abb =4;};};_ddf ,_aaab =d .Token ();if _aaab !=nil {return _aaab ;};if _cfc ,_dede :=_ddf .(_g .EndElement );_dede &&_cfc .Name ==start .Name {return nil ;};return _fc .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_ddf );};type ST_Persistence byte ;

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_gfd *CT_Picture )ValidateWithPath (path string )error {return nil };func (_deab *CT_OcxPrChoice )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {if _deab .Font !=nil {_aecd :=_g .StartElement {Name :_g .Name {Local :"\u0061x\u003a\u0066\u006f\u006e\u0074"}};e .EncodeElement (_deab .Font ,_aecd );};if _deab .Picture !=nil {_egb :=_g .StartElement {Name :_g .Name {Local :"\u0061\u0078\u003a\u0070\u0069\u0063\u0074\u0075\u0072\u0065"}};e .EncodeElement (_deab .Picture ,_egb );};return nil ;};const (ST_PersistenceUnset ST_Persistence =0;ST_PersistencePersistPropertyBag ST_Persistence =1;ST_PersistencePersistStream ST_Persistence =2;ST_PersistencePersistStreamInit ST_Persistence =3;ST_PersistencePersistStorage ST_Persistence =4;);type CT_Ocx struct{ClassidAttr string ;LicenseAttr *string ;IdAttr *string ;PersistenceAttr ST_Persistence ;OcxPr []*CT_OcxPr ;};

// Validate validates the CT_OcxPrChoice and its children
func (_eee *CT_OcxPrChoice )Validate ()error {return _eee .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0043h\u006f\u0069\u0063\u0065");};func (_cb *CT_Font )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for _ ,_fa :=range start .Attr {if _fa .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_fa .Name .Local =="\u0069\u0064"||_fa .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_fa .Name .Local =="\u0069\u0064"{_fg ,_d :=_fa .Value ,error (nil );if _d !=nil {return _d ;};_cb .IdAttr =&_fg ;continue ;};if _fa .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_cb .PersistenceAttr .UnmarshalXMLAttr (_fa );continue ;};};_e :for {_fd ,_cba :=d .Token ();if _cba !=nil {return _cba ;};switch _gg :=_fd .(type ){case _g .StartElement :switch _gg .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_ef :=NewCT_OcxPr ();if _ac :=d .DecodeElement (_ef ,&_gg );_ac !=nil {return _ac ;};_cb .OcxPr =append (_cb .OcxPr ,_ef );default:_fe .Log .Debug ("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0046\u006f\u006e\u0074\u0020\u0025\u0076",_gg .Name );if _ed :=d .Skip ();_ed !=nil {return _ed ;};};case _g .EndElement :break _e ;case _g .CharData :};};return nil ;};

// ValidateWithPath validates the CT_Font and its children, prefixing error messages with path
func (_cg *CT_Font )ValidateWithPath (path string )error {if _ee :=_cg .PersistenceAttr .ValidateWithPath (path +"\u002f\u0050e\u0072\u0073\u0069s\u0074\u0065\u006e\u0063\u0065\u0041\u0074\u0074\u0072");_ee !=nil {return _ee ;};for _gc ,_ca :=range _cg .OcxPr {if _gf :=_ca .ValidateWithPath (_fc .Sprintf ("\u0025\u0073\u002fO\u0063\u0078\u0050\u0072\u005b\u0025\u0064\u005d",path ,_gc ));_gf !=nil {return _gf ;};};return nil ;};

// Validate validates the Ocx and its children
func (_bbd *Ocx )Validate ()error {return _bbd .ValidateWithPath ("\u004f\u0063\u0078")};

// Validate validates the CT_Font and its children
func (_eda *CT_Font )Validate ()error {return _eda .ValidateWithPath ("\u0043T\u005f\u0046\u006f\u006e\u0074");};type Ocx struct{CT_Ocx };func (_gef ST_Persistence )String ()string {switch _gef {case 0:return "";case 1:return "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067";case 2:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d";case 3:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074";case 4:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065";};return "";};

// ValidateWithPath validates the CT_OcxPrChoice and its children, prefixing error messages with path
func (_abc *CT_OcxPrChoice )ValidateWithPath (path string )error {if _abc .Font !=nil {if _af :=_abc .Font .ValidateWithPath (path +"\u002f\u0046\u006fn\u0074");_af !=nil {return _af ;};};if _abc .Picture !=nil {if _ec :=_abc .Picture .ValidateWithPath (path +"\u002f\u0050\u0069\u0063\u0074\u0075\u0072\u0065");_ec !=nil {return _ec ;};};return nil ;};type CT_OcxPrChoice struct{Font *CT_Font ;Picture *CT_Picture ;};func NewOcx ()*Ocx {_dbf :=&Ocx {};_dbf .CT_Ocx =*NewCT_Ocx ();return _dbf };

// Validate validates the CT_Picture and its children
func (_eb *CT_Picture )Validate ()error {return _eb .ValidateWithPath ("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065");};func (_cgc ST_Persistence )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {return e .EncodeElement (_cgc .String (),start );};func (_dcd *Ocx )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_dcd .CT_Ocx =*NewCT_Ocx ();for _ ,_gec :=range start .Attr {if _gec .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_gec .Name .Local =="\u0069\u0064"||_gec .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_gec .Name .Local =="\u0069\u0064"{_dfd ,_cf :=_gec .Value ,error (nil );if _cf !=nil {return _cf ;};_dcd .IdAttr =&_dfd ;continue ;};if _gec .Name .Local =="\u0063l\u0061\u0073\u0073\u0069\u0064"{_gdd ,_dca :=_gec .Value ,error (nil );if _dca !=nil {return _dca ;};_dcd .ClassidAttr =_gdd ;continue ;};if _gec .Name .Local =="\u006ci\u0063\u0065\u006e\u0073\u0065"{_gage ,_fee :=_gec .Value ,error (nil );if _fee !=nil {return _fee ;};_dcd .LicenseAttr =&_gage ;continue ;};if _gec .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_dcd .PersistenceAttr .UnmarshalXMLAttr (_gec );continue ;};};_afg :for {_acd ,_efa :=d .Token ();if _efa !=nil {return _efa ;};switch _gga :=_acd .(type ){case _g .StartElement :switch _gga .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_gde :=NewCT_OcxPr ();if _eba :=d .DecodeElement (_gde ,&_gga );_eba !=nil {return _eba ;};_dcd .OcxPr =append (_dcd .OcxPr ,_gde );default:_fe .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u004fc\u0078\u0020\u0025\u0076",_gga .Name );if _gca :=d .Skip ();_gca !=nil {return _gca ;};};case _g .EndElement :break _afg ;case _g .CharData :};};return nil ;};func (_eaf *ST_Persistence )UnmarshalXMLAttr (attr _g .Attr )error {switch attr .Value {case "":*_eaf =0;case "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067":*_eaf =1;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d":*_eaf =2;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074":*_eaf =3;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065":*_eaf =4;};return nil ;};func (_bfg *Ocx )MarshalXML (e *_g .Encoder ,start _g .StartElement )error {start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058"});start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0061\u0078"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058"});start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0072"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"});start .Attr =append (start .Attr ,_g .Attr {Name :_g .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});start .Name .Local ="\u0061\u0078\u003a\u006f\u0063\u0078";return _bfg .CT_Ocx .MarshalXML (e ,start );};

// ValidateWithPath validates the CT_Ocx and its children, prefixing error messages with path
func (_aaa *CT_Ocx )ValidateWithPath (path string )error {if _aaa .PersistenceAttr ==ST_PersistenceUnset {return _fc .Errorf ("\u0025\u0073\u002f\u0050\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063e\u0041\u0074\u0074\u0072\u0020\u0069s\u0020\u0061\u0020\u006d\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020f\u0069\u0065\u006c\u0064",path );};if _df :=_aaa .PersistenceAttr .ValidateWithPath (path +"\u002f\u0050e\u0072\u0073\u0069s\u0074\u0065\u006e\u0063\u0065\u0041\u0074\u0074\u0072");_df !=nil {return _df ;};for _dea ,_cbd :=range _aaa .OcxPr {if _fgcf :=_cbd .ValidateWithPath (_fc .Sprintf ("\u0025\u0073\u002fO\u0063\u0078\u0050\u0072\u005b\u0025\u0064\u005d",path ,_dea ));_fgcf !=nil {return _fgcf ;};};return nil ;};func NewCT_Font ()*CT_Font {_ga :=&CT_Font {};return _ga };func (_fcc ST_Persistence )MarshalXMLAttr (name _g .Name )(_g .Attr ,error ){_agg :=_g .Attr {};_agg .Name =name ;switch _fcc {case ST_PersistenceUnset :_agg .Value ="";case ST_PersistencePersistPropertyBag :_agg .Value ="\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067";case ST_PersistencePersistStream :_agg .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d";case ST_PersistencePersistStreamInit :_agg .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074";case ST_PersistencePersistStorage :_agg .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065";};return _agg ,nil ;};type CT_Font struct{PersistenceAttr ST_Persistence ;IdAttr *string ;OcxPr []*CT_OcxPr ;};func (_aga *CT_Picture )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {for _ ,_bd :=range start .Attr {if _bd .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_bd .Name .Local =="\u0069\u0064"||_bd .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_bd .Name .Local =="\u0069\u0064"{_bcf ,_bfd :=_bd .Value ,error (nil );if _bfd !=nil {return _bfd ;};_aga .IdAttr =&_bcf ;continue ;};};for {_agf ,_efb :=d .Token ();if _efb !=nil {return _fc .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065\u003a\u0020%\u0073",_efb );};if _bcd ,_gad :=_agf .(_g .EndElement );_gad &&_bcd .Name ==start .Name {break ;};};return nil ;};func (_dbd ST_Persistence )Validate ()error {return _dbd .ValidateWithPath ("")};func (_fdc *CT_Ocx )UnmarshalXML (d *_g .Decoder ,start _g .StartElement )error {_fdc .PersistenceAttr =ST_Persistence (1);for _ ,_dg :=range start .Attr {if _dg .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_dg .Name .Local =="\u0069\u0064"||_dg .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_dg .Name .Local =="\u0069\u0064"{_fgc ,_gcb :=_dg .Value ,error (nil );if _gcb !=nil {return _gcb ;};_fdc .IdAttr =&_fgc ;continue ;};if _dg .Name .Local =="\u0063l\u0061\u0073\u0073\u0069\u0064"{_aa ,_ba :=_dg .Value ,error (nil );if _ba !=nil {return _ba ;};_fdc .ClassidAttr =_aa ;continue ;};if _dg .Name .Local =="\u006ci\u0063\u0065\u006e\u0073\u0065"{_gff ,_ffd :=_dg .Value ,error (nil );if _ffd !=nil {return _ffd ;};_fdc .LicenseAttr =&_gff ;continue ;};if _dg .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_fdc .PersistenceAttr .UnmarshalXMLAttr (_dg );continue ;};};_eg :for {_de ,_ggg :=d .Token ();if _ggg !=nil {return _ggg ;};switch _aec :=_de .(type ){case _g .StartElement :switch _aec .Name {case _g .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_dae :=NewCT_OcxPr ();if _bg :=d .DecodeElement (_dae ,&_aec );_bg !=nil {return _bg ;};_fdc .OcxPr =append (_fdc .OcxPr ,_dae );default:_fe .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0020\u0025\u0076",_aec .Name );if _ge :=d .Skip ();_ge !=nil {return _ge ;};};case _g .EndElement :break _eg ;case _g .CharData :};};return nil ;};func (_aca ST_Persistence )ValidateWithPath (path string )error {switch _aca {case 0,1,2,3,4:default:return _fc .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_aca ));};return nil ;};type CT_Picture struct{IdAttr *string ;};func init (){_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u004f\u0063\u0078",NewCT_Ocx );_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072",NewCT_OcxPr );_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043T\u005f\u0046\u006f\u006e\u0074",NewCT_Font );_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065",NewCT_Picture );_b .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u006f\u0063\u0078",NewOcx );};