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

package testutils ;import (_ca "crypto/md5";_be "encoding/hex";_d "encoding/json";_dg "errors";_ea "fmt";_bc "github.com/stretchr/testify/require";_ad "golang.org/x/image/font";_eg "golang.org/x/image/font/opentype";_dcb "golang.org/x/image/math/fixed";
_e "image";_de "image/color";_cg "image/draw";_ag "image/png";_dc "io";_c "io/ioutil";_a "math";_agb "os";_b "os/exec";_ce "path/filepath";_db "strings";_ba "sync";_gd "testing";_gc "time";);type ReferenceMap struct{_ba .Mutex ;_aga map[string ]ReferenceEntry ;
};func CombinePNGFiles (file1 ,file2 string )(bool ,error ){_cb ,_dd :=ReadPNG (file1 );if _dd !=nil {return false ,_dd ;};_ec ,_dd :=ReadPNG (file2 );if _dd !=nil {return false ,_dd ;};_fgc :=_e .Point {_cb .Bounds ().Dx (),0};_eff :=_e .Rectangle {_fgc ,_fgc .Add (_ec .Bounds ().Size ())};
_dfg :=_e .Rectangle {_e .Point {0,0},_eff .Max };_baa :=_e .NewRGBA (_dfg );_cg .Draw (_baa ,_cb .Bounds (),_cb ,_e .Point {0,0},_cg .Src );_cg .Draw (_baa ,_eff ,_ec ,_e .Point {0,0},_cg .Src );_cbc :=_ce .Dir (file1 );_ae :=_db .TrimSuffix (_ce .Base (file1 ),_ce .Ext (file1 ));
_fe ,_dd :=_agb .Create (_cbc +"\u002f"+_ae +"\u002d\u0063\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u002e\u0070\u006e\u0067");if _dd !=nil {return false ,_dd ;};defer _fe .Close ();if _dba :=_ag .Encode (_fe ,_baa );_dba !=nil {return false ,_dba ;};return true ,nil ;
};type ReferenceFile struct{Timestamp _gc .Time `json:"timestamp"`;Map *ReferenceMap `json:"map,omitempty"`;_gb string ;};func CompareImages (img1 ,img2 _e .Image )(bool ,error ){_gdg :=img1 .Bounds ();_acf :=0;for _adab :=0;_adab < _gdg .Size ().X ;_adab ++{for _ed :=0;
_ed < _gdg .Size ().Y ;_ed ++{_dbe ,_ead ,_fa ,_ :=img1 .At (_adab ,_ed ).RGBA ();_dbb ,_dbc ,_dca ,_ :=img2 .At (_adab ,_ed ).RGBA ();if _dbe !=_dbb ||_ead !=_dbc ||_fa !=_dca {_acf ++;};};};_gfe :=float64 (_acf )/float64 (_gdg .Dx ()*_gdg .Dy ());if _gfe > 0.0001{_ea .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_gfe ,_acf );
return false ,nil ;};return true ,nil ;};func _adg (_gba *_gd .T ,_gbf string )int64 {_cgc ,_eca :=_agb .Stat (_gbf );_bc .NoError (_gba ,_eca );return _cgc .Size ();};func (_aad *ReferenceFile )updateMap (_cd *ReferenceMap )int {var _bd int ;if _aad .Map ._aga ==nil {_aad .Map ._aga =map[string ]ReferenceEntry {};
};for _da ,_bf :=range _cd ._aga {_f ,_gf :=_aad .Map ._aga [_da ];if !_gf {_aad .Map ._aga [_da ]=_bf ;_bd ++;continue ;};if string (_f .Value )!=string (_bf .Value ){_aad .Map ._aga [_da ]=_bf ;_bd ++;};};for _cee :=range _aad .Map ._aga {if _ ,_ge :=_cd ._aga [_cee ];
!_ge {delete (_aad .Map ._aga ,_cee );_bd ++;};};return _bd ;};func (_ac *ReferenceMap )Write (key string ,entry ReferenceEntry ){_ac .Lock ();defer _ac .Unlock ();if _ac ._aga ==nil {_ac ._aga =map[string ]ReferenceEntry {};};_ac ._aga [key ]=entry ;};
func CopyFile (src ,dst string )error {_ade ,_ggg :=_agb .Open (src );if _ggg !=nil {return _ggg ;};defer _ade .Close ();_gfc ,_ggg :=_agb .Create (dst );if _ggg !=nil {return _ggg ;};defer _gfc .Close ();_ ,_ggg =_dc .Copy (_gfc ,_ade );return _ggg ;};
var (ErrRenderNotSupported =_dg .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m");
ErrImageSizeNotMatch =_dg .New ("\u0069\u006d\u0061ge\u0020\u0073\u0069\u007a\u0065\u0073\u0020\u0064\u006f\u006e\u0027\u0074\u0020\u006d\u0061\u0074\u0063\u0068"););func (_bcg *ReferenceMap )Copy ()*ReferenceMap {_ged :=ReferenceMap {_aga :make (map[string ]ReferenceEntry ,len (_bcg ._aga ))};
for _cac ,_cc :=range _bcg ._aga {_ged ._aga [_cac ]=_cc ;};return &_ged ;};func (_dgb *ReferenceMap )MarshalJSON ()([]byte ,error ){return _d .Marshal (_dgb ._aga )};func (_bb *ReferenceMap )Keys ()(_gg []string ){_gg =make ([]string ,len (_bb ._aga ));
var _df int ;for _dgd :=range _bb ._aga {_gg [_df ]=_dgd ;_df ++;};return _gg ;};func _dfc (_efg *_e .RGBA ,_dag string ,_dfcd string ,_fdd ,_edge int )error {_dgg ,_dbeg :=_c .ReadFile (_dag );if _dbeg !=nil {return _dbeg ;};_dda ,_dbeg :=_eg .Parse (_dgg );
if _dbeg !=nil {return _dbeg ;};_ggc ,_dbeg :=_eg .NewFace (_dda ,&_eg .FaceOptions {Size :18,DPI :72,Hinting :_ad .HintingNone });if _dbeg !=nil {return _dbeg ;};_deb :=&_ad .Drawer {Dst :_efg ,Src :_e .NewUniform (_de .RGBA {200,100,0,255}),Face :_ggc ,Dot :_dcb .P (_fdd ,_edge )};
_deb .DrawString (_dfcd );return nil ;};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;};if _ ,_dgde :=_b .LookPath ("\u0067\u0073");_dgde !=nil {return ErrRenderNotSupported ;};return _b .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_ea .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();
};func (_gfb *ReferenceMap )Read (key string )(ReferenceEntry ,bool ){_gfb .Lock ();defer _gfb .Unlock ();if _gfb ._aga ==nil {return ReferenceEntry {},false ;};_ada ,_beb :=_gfb ._aga [key ];return _ada ,_beb ;};func _cagd (_abd ,_effg _e .Rectangle )bool {return _abd .Min .X ==_effg .Min .X &&_abd .Min .Y ==_effg .Min .Y &&_abd .Max .X ==_effg .Max .X &&_abd .Max .Y ==_effg .Max .Y ;
};func HashFile (file string )(string ,error ){_ga ,_bff :=_agb .Open (file );if _bff !=nil {return "",_bff ;};defer _ga .Close ();_fg :=_ca .New ();if _ ,_bff =_dc .Copy (_fg ,_ga );_bff !=nil {return "",_bff ;};return _be .EncodeToString (_fg .Sum (nil )),nil ;
};func (_fc *ReferenceMap )Len ()int {return len (_fc ._aga )};func CreatePNGDiff (src ,dst string )(bool ,string ,float64 ,float64 ,error ){_cdcc ,_gcb :=ReadPNG (src );if _gcb !=nil {return false ,"",0,0,_gcb ;};_ffa ,_gcb :=ReadPNG (dst );if _gcb !=nil {return false ,"",0,0,_gcb ;
};_cff :=_cdcc .Bounds ();_gde :=_ffa .Bounds ();if !_cagd (_cff ,_gde ){return false ,"",0,0,ErrImageSizeNotMatch ;};_eac :=_e .NewRGBA (_e .Rect (0,0,_cff .Max .X ,_cff .Max .Y ));var (_dfgg float64 ;_gfg float64 ;);for _edb :=_cff .Min .Y ;_edb < _cff .Max .Y ;
_edb ++{for _bfe :=_cff .Min .X ;_bfe < _cff .Max .X ;_bfe ++{_dbd ,_dbf ,_fae ,_ecc :=_ffa .At (_bfe ,_edb ).RGBA ();_eac .Set (_bfe ,_edb ,_de .RGBA {uint8 (_dbd ),uint8 (_dbf ),uint8 (_fae ),64});_dga ,_beegb ,_cag ,_bae :=_cdcc .At (_bfe ,_edb ).RGBA ();
if !_bdb (_cdcc .At (_bfe ,_edb ),_ffa .At (_bfe ,_edb )){_eac .Set (_bfe ,_edb ,_de .RGBA {uint8 (_dga ),uint8 (_beegb ),uint8 (_cag ),uint8 (_bae )});_fad :=float64 (_dga )+float64 (_beegb )+float64 (_cag )+float64 (_bae )-float64 (_dbd )+float64 (_dbf )+float64 (_fae )+float64 (_ecc );
_fac :=_a .Sqrt (_a .Pow (_fad /float64 (_cff .Dx ()*_cff .Dy ()),2));_gfg +=_fac ;_dfgg ++;};};};_fdc :=_dfgg /float64 (_cff .Dx ()*_cff .Dy ())*100;_gfbe :=_ce .Dir (src );_fb :=_db .TrimSuffix (_ce .Base (src ),_ce .Ext (src ));_ffb ,_gcb :=_agb .Create (_gfbe +"\u002f"+_fb +"\u002dd\u0069\u0066\u0066\u002e\u0070\u006eg");
if _gcb !=nil {return false ,"",0,0,_gcb ;};defer _ffb .Close ();_def :=_db .Replace (_gfbe ,"\u0072\u0065\u006e\u0064\u0065\u0072","\u0066\u006f\u006et\u0073",1)+"\u002f\u0043\u0061l\u0069\u0062\u0072\u0069\u002e\u0074\u0074\u0066";_caf :=_ea .Sprintf ("\u0044\u0069f\u0066\u0065\u0072e\u006e\u0063\u0065\u003a\u0020\u0025\u0066\u0025\u0025",_fdc );
_gcb =_dfc (_eac ,_def ,_caf ,15,22);if _gcb !=nil {return false ,"",0,0,_gcb ;};_caf =_ea .Sprintf ("T\u006ft\u0061\u006c\u0020\u0044\u0069\u0066\u0066\u0065r\u0065\u006e\u0063\u0065: \u0025\u0066",_gfg );_gcb =_dfc (_eac ,_def ,_caf ,15,44);if _gcb !=nil {return false ,"",0,0,_gcb ;
};if _bdc :=_ag .Encode (_ffb ,_eac );_bdc !=nil {return false ,"",0,0,_bdc ;};return true ,_ffb .Name (),_fdc ,_gfg ,nil ;};func RunRenderTest (t *_gd .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ,currentHashMap *ReferenceMap ,refFile *ReferenceFile ){RunRenderOfficeTest (t ,pdfPath ,outputDir ,baselineRenderPath ,saveBaseline ,currentHashMap ,refFile ,"\u002em\u0073\u0077\u006f\u0072\u0064");
};type ReferenceEntry struct{Timestamp int64 `json:"timestamp"`;Value string `json:"value"`;ResultSize int64 `json:"resultSize,omitempty"`;DiffPercent float64 `json:"diffPercent,omitempty"`;DiffTotal float64 `json:"diffValue,omitempty"`;Invalid bool `json:"markedInvalid,omitempty"`;
};func ReadFile (dirPath ,testName string ,createEmpty bool )(*ReferenceFile ,error ){if dirPath ==""&&createEmpty {return &ReferenceFile {Map :&ReferenceMap {}},nil ;};if dirPath ==""{return nil ,_agb .ErrNotExist ;};_dgc :=_ce .Join (dirPath ,testName +"\u005fr\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u002e\u006a\u0073\u006f\u006e");
_fca :=&ReferenceFile {_gb :_dgc };_gef ,_agd :=_agb .Open (_dgc );if _dg .Is (_agd ,_agb .ErrNotExist )&&createEmpty {_fca .Timestamp =_gc .Now ().UTC ();_fca .Map =&ReferenceMap {};return _fca ,nil ;};if _agd !=nil {return nil ,_agd ;};defer _gef .Close ();
if _aada :=_d .NewDecoder (_gef ).Decode (_fca );_aada !=nil {if _aada .Error ()=="i\u006c\u006c\u0065\u0067\u0061\u006c \u0062\u0061\u0073\u0065\u0036\u0034 \u0064\u0061\u0074\u0061\u0020\u0061\u0074 \u0069\u006e\u0070\u0075\u0074\u0020\u0062\u0079\u0074\u0065 \u0030"&&createEmpty {return _fca ,nil ;
};return nil ,_aada ;};return _fca ,nil ;};func _bdb (_aaf ,_bga _de .Color )bool {_gce ,_dbg ,_dfe ,_edg :=_aaf .RGBA ();_bcdf ,_cfe ,_egb ,_fgg :=_bga .RGBA ();return _gce ==_bcdf &&_dbg ==_cfe &&_dfe ==_egb &&_edg ==_fgg ;};func (_bg *ReferenceFile )Update (currentMap *ReferenceMap )error {if _bg ._gb ==""{return nil ;
};_bcd :=_bg .updateMap (currentMap );if _bcd ==0{return nil ;};_aa ,_af :=_agb .OpenFile (_bg ._gb ,_agb .O_CREATE |_agb .O_TRUNC |_agb .O_WRONLY ,0664);if _af !=nil {return _af ;};defer _aa .Close ();_bg .Timestamp =_gc .Now ().UTC ();_aff :=_d .NewEncoder (_aa );
_aff .SetIndent ("","\u0009");return _aff .Encode (_bg );};func ReadPNG (file string )(_e .Image ,error ){_agf ,_cdc :=_agb .Open (file );if _cdc !=nil {return nil ,_cdc ;};defer _agf .Close ();return _ag .Decode (_agf );};func (_fd *ReferenceMap )UnmarshalJSON (data []byte )error {return _d .Unmarshal (data ,&_fd ._aga )};
func RunRenderOfficeTest (t *_gd .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ,currentHashMap *ReferenceMap ,refFile *ReferenceFile ,postfix string ){_eb :=_db .TrimSuffix (_ce .Base (pdfPath ),_ce .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_dcd *_gd .T ){_eed :=_ce .Join (outputDir ,_eb );
_eda :=_eed +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _ffe :=RenderPDFToPNGs (pdfPath ,0,_eda );_ffe !=nil {_dcd .Skip (_ffe );};_faf :=_eb +postfix ;_fef :=_ce .Join (outputDir ,_faf );_afa :=_fef +"\u002d%\u0064\u002e\u0070\u006e\u0067";_dae :=false ;
if saveBaseline {_ffg :=_ce .Dir (pdfPath );_dbbe :=_ce .Join (_ffg ,_faf +"\u002e\u0070\u0064\u0066");if _ ,_aab :=_agb .Stat (_dbbe );_aab ==nil {_dcd .Logf ("\u0052e\u006e\u0064\u0065\u0072\u0020\u004d\u0053\u0020\u004f\u0066\u0066i\u0063\u0065\u0020\u0050\u0044\u0046\u003a\u0020\u0025\u0076",_dbbe );
if _agac :=RenderPDFToPNGs (_dbbe ,0,_afa );_agac !=nil {_dcd .Skip (_agac );};_dae =true ;};};for _bdbc :=1;true ;_bdbc ++{_defa :=_ea .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_eed ,_bdbc );_aeg :=_ce .Join (baselineRenderPath ,_ea .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_eb ,_bdbc ));
if _ ,_ebb :=_agb .Stat (_defa );_ebb !=nil {break ;};_dcd .Logf ("\u0025\u0073",_aeg );if saveBaseline {_dcd .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_defa ,_aeg );_bfb :=CopyFile (_defa ,_aeg );if _bfb !=nil {_dcd .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_aeg ,_bfb );
};if _dae {_cde :=_ea .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_fef ,_bdbc );_cbe :=_ce .Join (baselineRenderPath ,_ea .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_faf ,_bdbc ));_dcd .Logf ("\u0043\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u004d\u0053\u0020\u0057\u006f\u0072\u0064 \u0072e\u0073\u0075\u006c\u0074\u0073\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_cde ,_cbe );
_bca :=CopyFile (_cde ,_cbe );if _bca !=nil {_dcd .Logf ("\u0045\u0052RO\u0052\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076\u002c \u0068\u0061\u0076\u0069n\u0067\u0020\u0064\u0069ff\u0065\u0072e\u006et\u0020\u0070\u0061\u0067\u0065\u0020s\u0069\u007a\u0065 \u0072\u0065s\u0075\u006c\u0074\u0073\u002c\u0020\u0075\u0073\u0065\u0020\u0070\u0072\u0065\u0076i\u006f\u0075\u0073\u0020\u0070\u0061\u0067\u0065",_cbe ,_bca );
_cde =_ea .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_fef ,_bdbc -1);_cbe =_ce .Join (baselineRenderPath ,_ea .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_faf ,_bdbc -1));if _bffb :=CopyFile (_cde ,_cbe );
_bffb !=nil {_dcd .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_cbe ,_bffb );};};_dcd .Logf ("\u0043\u006fm\u0062\u0069\u006e\u0069\u006eg\u0020\u0055\u006e\u0069\u004ff\u0066\u0069\u0063\u0065\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u0020\u0077\u0069\u0074\u0068\u0020\u004d\u0053\u0020\u004f\u0066\u0066\u0069\u0063\u0065\u0020\u0025\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0073",_aeg ,_cbe );
_fec ,_bca :=CombinePNGFiles (_aeg ,_cbe );if _agb .IsNotExist (_bca ){_dcd .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_fec {_dcd .Fatal ("\u0055n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063\u006f\u006db\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0073");
};_dcd .Logf ("\u0043\u0072\u0065\u0061\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0064\u0069\u0066f \u0055n\u0069\u004f\u0066\u0066\u0069\u0063\u0065\u0020\u0072\u0065\u0073\u0075l\u0074\u0073\u0020\u0077\u0069\u0074\u0068\u0020\u004d\u0053\u0020\u004f\u0066\u0066\u0069\u0063\u0065 \u0025\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0073",_aeg ,_cbe );
_fec ,_cge ,_feg ,_dcf ,_bca :=CreatePNGDiff (_aeg ,_cbe );if _bca !=nil &&_bca !=ErrImageSizeNotMatch {_dcd .Fatalf ("\u0045\u0072\u0072\u006fr\u0020\u006f\u006e\u0020\u0063\u0072\u0065\u0061\u0074\u0065 \u0050N\u0047\u0020\u0044\u0069\u0066\u0066\u003a \u0025\u0076",_bca );
};if _fec {_dcd .Logf ("\u0049\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073\u000a",_cge );_dcd .Logf ("D\u0069\u0066\u0066\u0020Pe\u0072c\u0065\u006e\u0074\u003a\u0020%\u0032\u002e\u0066\u0025\u0025\u000a",_feg );_dcd .Logf ("\u0044i\u0066f\u0020\u0054\u006f\u0074\u0061\u006c\u003a\u0020\u0025\u0066\u000a",_dcf );
_fbb :=_ce .Base (_cge );if _faeg ,_gda :=currentHashMap .Read (_fbb );_gda {if _faeg .DiffPercent < _feg ||_faeg .DiffTotal < _dcf {_dcd .Fatalf ("\u004e\u0065\u0077\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u0020\u0068\u0061\u0076\u0069\u006e\u0067 h\u0069g\u0068\u0065\u0072\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0070\u0065\u0072\u0063\u0065\u006e\u0074\u003a\u0020\u0025\u0066\u0020\u0061\u006e\u0064 \u0074\u006f\u0074\u0061\u006c\u0020\u0025\u0066\u000a",_feg ,_dcf );
};};_ggb ,_dbce :=HashFile (_cge );_bc .NoError (_dcd ,_dbce );_bda :=ReferenceEntry {Timestamp :_gc .Now ().UTC ().Unix (),Value :_ggb ,ResultSize :_adg (_dcd ,_cge ),DiffPercent :_feg ,DiffTotal :_dcf };currentHashMap .Write (_fbb ,_bda );if _dbce =refFile .Update (currentHashMap );
_dbce !=nil {_dcd .Logf ("\u0055\u0070\u0064\u0061\u0074\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020f\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_dbce );};};};continue ;};_dcd .Run (_ea .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_bdbc ),func (_fda *_gd .T ){_fda .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_defa ,_aeg );
_gdae ,_gad :=ComparePNGFiles (_defa ,_aeg );if _agb .IsNotExist (_gad ){_fda .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_gdae {_fda .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");
};});};});};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_beeg ,_ab :=HashFile (file1 );if _ab !=nil {return false ,_ab ;};_ded ,_ab :=HashFile (file2 );if _ab !=nil {return false ,_ab ;};if _beeg ==_ded {return true ,nil ;};_gfd ,_ab :=ReadPNG (file1 );
if _ab !=nil {return false ,_ab ;};_ff ,_ab :=ReadPNG (file2 );if _ab !=nil {return false ,_ab ;};if _gfd .Bounds ()!=_ff .Bounds (){return false ,nil ;};return CompareImages (_gfd ,_ff );};