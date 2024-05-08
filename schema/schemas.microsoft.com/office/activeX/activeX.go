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

package activeX ;import (_c "encoding/xml";_d "fmt";_ac "github.com/unidoc/unioffice";_g "github.com/unidoc/unioffice/common/logger";);func NewCT_Picture ()*CT_Picture {_gae :=&CT_Picture {};return _gae };type CT_Ocx struct{ClassidAttr string ;LicenseAttr *string ;
IdAttr *string ;PersistenceAttr ST_Persistence ;OcxPr []*CT_OcxPr ;};func (_gg *CT_Ocx )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_gg .PersistenceAttr =ST_Persistence (1);for _ ,_fbe :=range start .Attr {if _fbe .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_fbe .Name .Local =="\u0069\u0064"||_fbe .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_fbe .Name .Local =="\u0069\u0064"{_fg ,_dg :=_fbe .Value ,error (nil );
if _dg !=nil {return _dg ;};_gg .IdAttr =&_fg ;continue ;};if _fbe .Name .Local =="\u0063l\u0061\u0073\u0073\u0069\u0064"{_gfd ,_fbf :=_fbe .Value ,error (nil );if _fbf !=nil {return _fbf ;};_gg .ClassidAttr =_gfd ;continue ;};if _fbe .Name .Local =="\u006ci\u0063\u0065\u006e\u0073\u0065"{_ad ,_gc :=_fbe .Value ,error (nil );
if _gc !=nil {return _gc ;};_gg .LicenseAttr =&_ad ;continue ;};if _fbe .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_gg .PersistenceAttr .UnmarshalXMLAttr (_fbe );continue ;};};_ec :for {_cg ,_dga :=d .Token ();if _dga !=nil {return _dga ;
};switch _edc :=_cg .(type ){case _c .StartElement :switch _edc .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_cgf :=NewCT_OcxPr ();
if _gfe :=d .DecodeElement (_cgf ,&_edc );_gfe !=nil {return _gfe ;};_gg .OcxPr =append (_gg .OcxPr ,_cgf );default:_g .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0020\u0025\u0076",_edc .Name );
if _ebcf :=d .Skip ();_ebcf !=nil {return _ebcf ;};};case _c .EndElement :break _ec ;case _c .CharData :};};return nil ;};type CT_Picture struct{IdAttr *string ;};

// ValidateWithPath validates the CT_Ocx and its children, prefixing error messages with path
func (_fd *CT_Ocx )ValidateWithPath (path string )error {if _fd .PersistenceAttr ==ST_PersistenceUnset {return _d .Errorf ("\u0025\u0073\u002f\u0050\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063e\u0041\u0074\u0074\u0072\u0020\u0069s\u0020\u0061\u0020\u006d\u0061\u006e\u0064\u0061\u0074\u006f\u0072\u0079\u0020f\u0069\u0065\u006c\u0064",path );
};if _edg :=_fd .PersistenceAttr .ValidateWithPath (path +"\u002f\u0050e\u0072\u0073\u0069s\u0074\u0065\u006e\u0063\u0065\u0041\u0074\u0074\u0072");_edg !=nil {return _edg ;};for _cc ,_bf :=range _fd .OcxPr {if _bac :=_bf .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002fO\u0063\u0078\u0050\u0072\u005b\u0025\u0064\u005d",path ,_cc ));
_bac !=nil {return _bac ;};};return nil ;};func (_bfc *CT_OcxPr )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_bdb :=range start .Attr {if _bdb .Name .Local =="\u006e\u0061\u006d\u0065"{_fda ,_cegf :=_bdb .Value ,error (nil );if _cegf !=nil {return _cegf ;
};_bfc .NameAttr =_fda ;continue ;};if _bdb .Name .Local =="\u0076\u0061\u006cu\u0065"{_cf ,_ege :=_bdb .Value ,error (nil );if _ege !=nil {return _ege ;};_bfc .ValueAttr =&_cf ;continue ;};};_gb :for {_aaf ,_gba :=d .Token ();if _gba !=nil {return _gba ;
};switch _dda :=_aaf .(type ){case _c .StartElement :switch _dda .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0066\u006f\u006e\u0074"}:_bfc .Choice =NewCT_OcxPrChoice ();
if _bdf :=d .DecodeElement (&_bfc .Choice .Font ,&_dda );_bdf !=nil {return _bdf ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0070i\u0063\u0074\u0075\u0072\u0065"}:_bfc .Choice =NewCT_OcxPrChoice ();
if _gd :=d .DecodeElement (&_bfc .Choice .Picture ,&_dda );_gd !=nil {return _gd ;};default:_g .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006eg\u0020\u0075\u006es\u0075\u0070\u0070\u006fr\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0020\u0025\u0076",_dda .Name );
if _gea :=d .Skip ();_gea !=nil {return _gea ;};};case _c .EndElement :break _gb ;case _c .CharData :};};return nil ;};func (_edcd *CT_Picture )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_ddaa :=range start .Attr {if _ddaa .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_ddaa .Name .Local =="\u0069\u0064"||_ddaa .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_ddaa .Name .Local =="\u0069\u0064"{_add ,_gcf :=_ddaa .Value ,error (nil );
if _gcf !=nil {return _gcf ;};_edcd .IdAttr =&_add ;continue ;};};for {_bea ,_faea :=d .Token ();if _faea !=nil {return _d .Errorf ("\u0070\u0061\u0072\u0073in\u0067\u0020\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065\u003a\u0020%\u0073",_faea );
};if _bef ,_gbab :=_bea .(_c .EndElement );_gbab &&_bef .Name ==start .Name {break ;};};return nil ;};

// Validate validates the CT_OcxPr and its children
func (_fe *CT_OcxPr )Validate ()error {return _fe .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072");};func NewCT_Ocx ()*CT_Ocx {_ef :=&CT_Ocx {};_ef .PersistenceAttr =ST_Persistence (1);return _ef };const (ST_PersistenceUnset ST_Persistence =0;
ST_PersistencePersistPropertyBag ST_Persistence =1;ST_PersistencePersistStream ST_Persistence =2;ST_PersistencePersistStreamInit ST_Persistence =3;ST_PersistencePersistStorage ST_Persistence =4;);func (_adab ST_Persistence )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {return e .EncodeElement (_adab .String (),start );
};func (_aa *CT_Ocx )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0061\u0078\u003a\u0063\u006c\u0061\u0073\u0073\u0069\u0064"},Value :_d .Sprintf ("\u0025\u0076",_aa .ClassidAttr )});
if _aa .LicenseAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0061\u0078\u003a\u006c\u0069\u0063\u0065\u006e\u0073\u0065"},Value :_d .Sprintf ("\u0025\u0076",*_aa .LicenseAttr )});};if _aa .IdAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_d .Sprintf ("\u0025\u0076",*_aa .IdAttr )});
};_ebc ,_bbf :=_aa .PersistenceAttr .MarshalXMLAttr (_c .Name {Local :"\u0061\u0078\u003a\u0070\u0065\u0072\u0073\u0069\u0073t\u0065\u006e\u0063\u0065"});if _bbf !=nil {return _bbf ;};start .Attr =append (start .Attr ,_ebc );e .EncodeToken (start );if _aa .OcxPr !=nil {_ce :=_c .StartElement {Name :_c .Name {Local :"\u0061\u0078\u003a\u006f\u0063\u0078\u0050\u0072"}};
for _ ,_ceg :=range _aa .OcxPr {e .EncodeElement (_ceg ,_ce );};};e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};

// ValidateWithPath validates the CT_OcxPr and its children, prefixing error messages with path
func (_ccb *CT_OcxPr )ValidateWithPath (path string )error {if _ccb .Choice !=nil {if _ada :=_ccb .Choice .ValidateWithPath (path +"\u002fC\u0068\u006f\u0069\u0063\u0065");_ada !=nil {return _ada ;};};return nil ;};func NewCT_OcxPr ()*CT_OcxPr {_cga :=&CT_OcxPr {};
return _cga };func (_dce ST_Persistence )ValidateWithPath (path string )error {switch _dce {case 0,1,2,3,4:default:return _d .Errorf ("\u0025s\u003a\u0020\u006f\u0075t\u0020\u006f\u0066\u0020\u0072a\u006eg\u0065 \u0076\u0061\u006c\u0075\u0065\u0020\u0025d",path ,int (_dce ));
};return nil ;};func NewCT_OcxPrChoice ()*CT_OcxPrChoice {_ff :=&CT_OcxPrChoice {};return _ff };func (_fef ST_Persistence )MarshalXMLAttr (name _c .Name )(_c .Attr ,error ){_dgg :=_c .Attr {};_dgg .Name =name ;switch _fef {case ST_PersistenceUnset :_dgg .Value ="";
case ST_PersistencePersistPropertyBag :_dgg .Value ="\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067";case ST_PersistencePersistStream :_dgg .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d";
case ST_PersistencePersistStreamInit :_dgg .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074";case ST_PersistencePersistStorage :_dgg .Value ="\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065";
};return _dgg ,nil ;};

// Validate validates the CT_Ocx and its children
func (_ebf *CT_Ocx )Validate ()error {return _ebf .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078");};

// Validate validates the CT_OcxPrChoice and its children
func (_bbd *CT_OcxPrChoice )Validate ()error {return _bbd .ValidateWithPath ("\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0043h\u006f\u0069\u0063\u0065");};type CT_OcxPr struct{NameAttr string ;ValueAttr *string ;Choice *CT_OcxPrChoice ;};func (_fc *CT_OcxPrChoice )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _fc .Font !=nil {_fce :=_c .StartElement {Name :_c .Name {Local :"\u0061x\u003a\u0066\u006f\u006e\u0074"}};
e .EncodeElement (_fc .Font ,_fce );};if _fc .Picture !=nil {_bcb :=_c .StartElement {Name :_c .Name {Local :"\u0061\u0078\u003a\u0070\u0069\u0063\u0074\u0075\u0072\u0065"}};e .EncodeElement (_fc .Picture ,_bcb );};return nil ;};func (_cb *Ocx )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_cb .CT_Ocx =*NewCT_Ocx ();
for _ ,_gdb :=range start .Attr {if _gdb .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_gdb .Name .Local =="\u0069\u0064"||_gdb .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_gdb .Name .Local =="\u0069\u0064"{_dge ,_acd :=_gdb .Value ,error (nil );
if _acd !=nil {return _acd ;};_cb .IdAttr =&_dge ;continue ;};if _gdb .Name .Local =="\u0063l\u0061\u0073\u0073\u0069\u0064"{_bcc ,_bgg :=_gdb .Value ,error (nil );if _bgg !=nil {return _bgg ;};_cb .ClassidAttr =_bcc ;continue ;};if _gdb .Name .Local =="\u006ci\u0063\u0065\u006e\u0073\u0065"{_fed ,_fbff :=_gdb .Value ,error (nil );
if _fbff !=nil {return _fbff ;};_cb .LicenseAttr =&_fed ;continue ;};if _gdb .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_cb .PersistenceAttr .UnmarshalXMLAttr (_gdb );continue ;};};_bcbe :for {_da ,_dca :=d .Token ();
if _dca !=nil {return _dca ;};switch _eggc :=_da .(type ){case _c .StartElement :switch _eggc .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_efa :=NewCT_OcxPr ();
if _ae :=d .DecodeElement (_efa ,&_eggc );_ae !=nil {return _ae ;};_cb .OcxPr =append (_cb .OcxPr ,_efa );default:_g .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065m\u0065\u006e\u0074\u0020\u006fn\u0020\u004fc\u0078\u0020\u0025\u0076",_eggc .Name );
if _acba :=d .Skip ();_acba !=nil {return _acba ;};};case _c .EndElement :break _bcbe ;case _c .CharData :};};return nil ;};func (_dfe ST_Persistence )String ()string {switch _dfe {case 0:return "";case 1:return "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067";
case 2:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d";case 3:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074";case 4:return "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065";
};return "";};func NewOcx ()*Ocx {_acb :=&Ocx {};_acb .CT_Ocx =*NewCT_Ocx ();return _acb };func (_de *CT_Font )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _de .PersistenceAttr !=ST_PersistenceUnset {_af ,_b :=_de .PersistenceAttr .MarshalXMLAttr (_c .Name {Local :"\u0061\u0078\u003a\u0070\u0065\u0072\u0073\u0069\u0073t\u0065\u006e\u0063\u0065"});
if _b !=nil {return _b ;};start .Attr =append (start .Attr ,_af );};if _de .IdAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_d .Sprintf ("\u0025\u0076",*_de .IdAttr )});};e .EncodeToken (start );
if _de .OcxPr !=nil {_cd :=_c .StartElement {Name :_c .Name {Local :"\u0061\u0078\u003a\u006f\u0063\u0078\u0050\u0072"}};for _ ,_f :=range _de .OcxPr {e .EncodeElement (_f ,_cd );};};e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type CT_Font struct{PersistenceAttr ST_Persistence ;
IdAttr *string ;OcxPr []*CT_OcxPr ;};

// ValidateWithPath validates the CT_OcxPrChoice and its children, prefixing error messages with path
func (_gac *CT_OcxPrChoice )ValidateWithPath (path string )error {if _gac .Font !=nil {if _bg :=_gac .Font .ValidateWithPath (path +"\u002f\u0046\u006fn\u0074");_bg !=nil {return _bg ;};};if _gac .Picture !=nil {if _fgf :=_gac .Picture .ValidateWithPath (path +"\u002f\u0050\u0069\u0063\u0074\u0075\u0072\u0065");
_fgf !=nil {return _fgf ;};};return nil ;};type ST_Persistence byte ;type CT_OcxPrChoice struct{Font *CT_Font ;Picture *CT_Picture ;};

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (_fdc *CT_Picture )ValidateWithPath (path string )error {return nil };func (_fff *ST_Persistence )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_gcd ,_aadd :=d .Token ();if _aadd !=nil {return _aadd ;};if _cbf ,_fdaa :=_gcd .(_c .EndElement );
_fdaa &&_cbf .Name ==start .Name {*_fff =1;return nil ;};if _gab ,_gfda :=_gcd .(_c .CharData );!_gfda {return _d .Errorf ("\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0063\u0068a\u0072\u0020\u0064\u0061\u0074\u0061\u002c\u0020\u0067\u006ft\u0020\u0025\u0054",_gcd );
}else {switch string (_gab ){case "":*_fff =0;case "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067":*_fff =1;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d":*_fff =2;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074":*_fff =3;
case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065":*_fff =4;};};_gcd ,_aadd =d .Token ();if _aadd !=nil {return _aadd ;};if _bab ,_bfb :=_gcd .(_c .EndElement );_bfb &&_bab .Name ==start .Name {return nil ;};return _d .Errorf ("\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0065\u006e\u0064\u0020\u0065\u006ce\u006de\u006e\u0074\u002c\u0020\u0067\u006f\u0074 \u0025\u0076",_gcd );
};func (_ca ST_Persistence )Validate ()error {return _ca .ValidateWithPath ("")};

// Validate validates the CT_Font and its children
func (_ga *CT_Font )Validate ()error {return _ga .ValidateWithPath ("\u0043T\u005f\u0046\u006f\u006e\u0074");};

// ValidateWithPath validates the Ocx and its children, prefixing error messages with path
func (_cfb *Ocx )ValidateWithPath (path string )error {if _bde :=_cfb .CT_Ocx .ValidateWithPath (path );_bde !=nil {return _bde ;};return nil ;};func NewCT_Font ()*CT_Font {_gf :=&CT_Font {};return _gf };

// Validate validates the Ocx and its children
func (_adb *Ocx )Validate ()error {return _adb .ValidateWithPath ("\u004f\u0063\u0078")};func (_egec *ST_Persistence )UnmarshalXMLAttr (attr _c .Attr )error {switch attr .Value {case "":*_egec =0;case "\u0070e\u0072s\u0069\u0073\u0074\u0050\u0072o\u0070\u0065r\u0074\u0079\u0042\u0061\u0067":*_egec =1;
case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061\u006d":*_egec =2;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074\u0072\u0065\u0061m\u0049\u006e\u0069\u0074":*_egec =3;case "\u0070\u0065\u0072\u0073\u0069\u0073\u0074\u0053\u0074o\u0072\u0061\u0067\u0065":*_egec =4;
};return nil ;};func (_feg *CT_OcxPrChoice )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {_bcg :for {_dc ,_aae :=d .Token ();if _aae !=nil {return _aae ;};switch _adc :=_dc .(type ){case _c .StartElement :switch _adc .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0066\u006f\u006e\u0074"}:_feg .Font =NewCT_Font ();
if _egg :=d .DecodeElement (_feg .Font ,&_adc );_egg !=nil {return _egg ;};case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u0070i\u0063\u0074\u0075\u0072\u0065"}:_feg .Picture =NewCT_Picture ();
if _fab :=d .DecodeElement (_feg .Picture ,&_adc );_fab !=nil {return _fab ;};default:_g .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069n\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006et\u0020\u006f\u006e\u0020\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072\u0043\u0068o\u0069c\u0065\u0020\u0025\u0076",_adc .Name );
if _gbaa :=d .Skip ();_gbaa !=nil {return _gbaa ;};};case _c .EndElement :break _bcg ;case _c .CharData :};};return nil ;};

// Validate validates the CT_Picture and its children
func (_ddaaa *CT_Picture )Validate ()error {return _ddaaa .ValidateWithPath ("\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065");};func (_ade *Ocx )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006cn\u0073"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078\u006d\u006c\u006e\u0073\u003a\u0061\u0078"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0072"},Value :"\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"});
start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0078m\u006c\u006e\u0073\u003a\u0078\u006dl"},Value :"\u0068\u0074tp\u003a\u002f\u002fw\u0077\u0077\u002e\u00773.o\u0072g/\u0058\u004d\u004c\u002f\u0031\u0039\u00398/\u006e\u0061\u006d\u0065\u0073\u0070\u0061c\u0065"});
start .Name .Local ="\u0061\u0078\u003a\u006f\u0063\u0078";return _ade .CT_Ocx .MarshalXML (e ,start );};func (_ace *CT_Picture )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {if _ace .IdAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0072\u003a\u0069\u0064"},Value :_d .Sprintf ("\u0025\u0076",*_ace .IdAttr )});
};e .EncodeToken (start );e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};type Ocx struct{CT_Ocx };func (_ab *CT_OcxPr )MarshalXML (e *_c .Encoder ,start _c .StartElement )error {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0061x\u003a\u006e\u0061\u006d\u0065"},Value :_d .Sprintf ("\u0025\u0076",_ab .NameAttr )});
if _ab .ValueAttr !=nil {start .Attr =append (start .Attr ,_c .Attr {Name :_c .Name {Local :"\u0061\u0078\u003a\u0076\u0061\u006c\u0075\u0065"},Value :_d .Sprintf ("\u0025\u0076",*_ab .ValueAttr )});};e .EncodeToken (start );if _ab .Choice !=nil {_ab .Choice .MarshalXML (e ,_c .StartElement {});
};e .EncodeToken (_c .EndElement {Name :start .Name });return nil ;};func (_fa *CT_Font )UnmarshalXML (d *_c .Decoder ,start _c .StartElement )error {for _ ,_fb :=range start .Attr {if _fb .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f/\u0073\u0063\u0068\u0065\u006da\u0073\u002e\u006f\u0070\u0065\u006ex\u006d\u006c\u0066\u006f\u0072m\u0061\u0074\u0073\u002e\u006f\u0072\u0067\u002f\u006f\u0066\u0066\u0069c\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0032\u0030\u0030\u0036\u002fr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073h\u0069\u0070\u0073"&&_fb .Name .Local =="\u0069\u0064"||_fb .Name .Space =="\u0068\u0074\u0074\u0070\u003a\u002f\u002fp\u0075\u0072\u006c.\u006f\u0063\u006cc\u002e\u006fr\u0067\u002f\u006f\u006f\u0078\u006dl\u002fof\u0066\u0069\u0063\u0065\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u002f\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0073\u0068\u0069\u0070\u0073"&&_fb .Name .Local =="\u0069\u0064"{_df ,_e :=_fb .Value ,error (nil );
if _e !=nil {return _e ;};_fa .IdAttr =&_df ;continue ;};if _fb .Name .Local =="p\u0065\u0072\u0073\u0069\u0073\u0074\u0065\u006e\u0063\u0065"{_fa .PersistenceAttr .UnmarshalXMLAttr (_fb );continue ;};};_ed :for {_cdc ,_bc :=d .Token ();if _bc !=nil {return _bc ;
};switch _ea :=_cdc .(type ){case _c .StartElement :switch _ea .Name {case _c .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058",Local :"\u006f\u0063\u0078P\u0072"}:_ddb :=NewCT_OcxPr ();
if _ba :=d .DecodeElement (_ddb ,&_ea );_ba !=nil {return _ba ;};_fa .OcxPr =append (_fa .OcxPr ,_ddb );default:_g .Log .Debug ("\u0073\u006b\u0069p\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0043T\u005f\u0046\u006f\u006e\u0074\u0020\u0025\u0076",_ea .Name );
if _bd :=d .Skip ();_bd !=nil {return _bd ;};};case _c .EndElement :break _ed ;case _c .CharData :};};return nil ;};

// ValidateWithPath validates the CT_Font and its children, prefixing error messages with path
func (_ge *CT_Font )ValidateWithPath (path string )error {if _eb :=_ge .PersistenceAttr .ValidateWithPath (path +"\u002f\u0050e\u0072\u0073\u0069s\u0074\u0065\u006e\u0063\u0065\u0041\u0074\u0074\u0072");_eb !=nil {return _eb ;};for _bb ,_eg :=range _ge .OcxPr {if _aff :=_eg .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002fO\u0063\u0078\u0050\u0072\u005b\u0025\u0064\u005d",path ,_bb ));
_aff !=nil {return _aff ;};};return nil ;};func init (){_ac .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u004f\u0063\u0078",NewCT_Ocx );
_ac .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u004f\u0063\u0078\u0050\u0072",NewCT_OcxPr );
_ac .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043T\u005f\u0046\u006f\u006e\u0074",NewCT_Font );
_ac .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u0043\u0054\u005f\u0050\u0069\u0063\u0074\u0075\u0072\u0065",NewCT_Picture );
_ac .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0073c\u0068\u0065\u006das\u002e\u006d\u0069\u0063\u0072\u006fs\u006f\u0066\u0074\u002e\u0063\u006f\u006d\u002f\u006f\u0066\u0066\u0069\u0063\u0065\u002f2\u0030\u0030\u0036\u002f\u0061\u0063\u0074\u0069v\u0065\u0058","\u006f\u0063\u0078",NewOcx );
};