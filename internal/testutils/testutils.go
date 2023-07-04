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

package testutils ;import (_fdf "crypto/md5";_bc "encoding/hex";_fa "encoding/json";_gf "errors";_fd "fmt";_bb "github.com/stretchr/testify/require";_gd "golang.org/x/image/font";_fab "golang.org/x/image/font/opentype";_dee "golang.org/x/image/math/fixed";_a "image";_ef "image/color";_bg "image/draw";_gb "image/png";_b "io";_e "io/ioutil";_gc "math";_c "os";_de "os/exec";_ee "path/filepath";_g "strings";_af "sync";_d "testing";_fb "time";);type ReferenceEntry struct{Timestamp int64 `json:"timestamp"`;Value string `json:"value"`;ResultSize int64 `json:"resultSize,omitempty"`;DiffPercent float64 `json:"diffPercent,omitempty"`;DiffTotal float64 `json:"diffValue,omitempty"`;Invalid bool `json:"markedInvalid,omitempty"`;};func (_da *ReferenceMap )UnmarshalJSON (data []byte )error {return _fa .Unmarshal (data ,&_da ._cgf )};func CompareImages (img1 ,img2 _a .Image )(bool ,error ){_cfc :=img1 .Bounds ();_aeg :=0;for _fcb :=0;_fcb < _cfc .Size ().X ;_fcb ++{for _eee :=0;_eee < _cfc .Size ().Y ;_eee ++{_edb ,_eg ,_egf ,_ :=img1 .At (_fcb ,_eee ).RGBA ();_ege ,_fef ,_afg ,_ :=img2 .At (_fcb ,_eee ).RGBA ();if _edb !=_ege ||_eg !=_fef ||_egf !=_afg {_aeg ++;};};};_add :=float64 (_aeg )/float64 (_cfc .Dx ()*_cfc .Dy ());if _add > 0.0001{_fd .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_add ,_aeg );return false ,nil ;};return true ,nil ;};func (_cf *ReferenceMap )Len ()int {return len (_cf ._cgf )};func (_aag *ReferenceMap )Write (key string ,entry ReferenceEntry ){_aag .Lock ();defer _aag .Unlock ();if _aag ._cgf ==nil {_aag ._cgf =map[string ]ReferenceEntry {};};_aag ._cgf [key ]=entry ;};func CreatePNGDiff (src ,dst string )(bool ,string ,float64 ,float64 ,error ){_abc ,_fca :=ReadPNG (src );if _fca !=nil {return false ,"",0,0,_fca ;};_eacg ,_fca :=ReadPNG (dst );if _fca !=nil {return false ,"",0,0,_fca ;};_fbb :=_abc .Bounds ();_aca :=_eacg .Bounds ();if !_ccc (_fbb ,_aca ){return false ,"",0,0,ErrImageSizeNotMatch ;};_ecc :=_a .NewRGBA (_a .Rect (0,0,_fbb .Max .X ,_fbb .Max .Y ));var (_cc float64 ;_gec float64 ;);for _cedc :=_fbb .Min .Y ;_cedc < _fbb .Max .Y ;_cedc ++{for _abcc :=_fbb .Min .X ;_abcc < _fbb .Max .X ;_abcc ++{_gg ,_dce ,_aad ,_fac :=_eacg .At (_abcc ,_cedc ).RGBA ();_ecc .Set (_abcc ,_cedc ,_ef .RGBA {uint8 (_gg ),uint8 (_dce ),uint8 (_aad ),64});_efc ,_fgc ,_fgb ,_gga :=_abc .At (_abcc ,_cedc ).RGBA ();if !_fcc (_abc .At (_abcc ,_cedc ),_eacg .At (_abcc ,_cedc )){_ecc .Set (_abcc ,_cedc ,_ef .RGBA {uint8 (_efc ),uint8 (_fgc ),uint8 (_fgb ),uint8 (_gga )});_gcg :=float64 (_efc )+float64 (_fgc )+float64 (_fgb )+float64 (_gga )-float64 (_gg )+float64 (_dce )+float64 (_aad )+float64 (_fac );_bgd :=_gc .Sqrt (_gc .Pow (_gcg /float64 (_fbb .Dx ()*_fbb .Dy ()),2));_gec +=_bgd ;_cc ++;};};};_deb :=_cc /float64 (_fbb .Dx ()*_fbb .Dy ())*100;_gfd :=_ee .Dir (src );_feb :=_g .TrimSuffix (_ee .Base (src ),_ee .Ext (src ));_dbae ,_fca :=_c .Create (_gfd +"\u002f"+_feb +"\u002dd\u0069\u0066\u0066\u002e\u0070\u006eg");if _fca !=nil {return false ,"",0,0,_fca ;};defer _dbae .Close ();_be :=_g .Replace (_gfd ,"\u0072\u0065\u006e\u0064\u0065\u0072","\u0066\u006f\u006et\u0073",1)+"\u002f\u0043\u0061l\u0069\u0062\u0072\u0069\u002e\u0074\u0074\u0066";_acaf :=_fd .Sprintf ("\u0044\u0069f\u0066\u0065\u0072e\u006e\u0063\u0065\u003a\u0020\u0025\u0066\u0025\u0025",_deb );_fca =_abe (_ecc ,_be ,_acaf ,15,22);if _fca !=nil {return false ,"",0,0,_fca ;};_acaf =_fd .Sprintf ("T\u006ft\u0061\u006c\u0020\u0044\u0069\u0066\u0066\u0065r\u0065\u006e\u0063\u0065: \u0025\u0066",_gec );_fca =_abe (_ecc ,_be ,_acaf ,15,44);if _fca !=nil {return false ,"",0,0,_fca ;};if _aedb :=_gb .Encode (_dbae ,_ecc );_aedb !=nil {return false ,"",0,0,_aedb ;};return true ,_dbae .Name (),_deb ,_gec ,nil ;};func ReadFile (dirPath ,testName string ,createEmpty bool )(*ReferenceFile ,error ){if dirPath ==""&&createEmpty {return &ReferenceFile {Map :&ReferenceMap {}},nil ;};if dirPath ==""{return nil ,_c .ErrNotExist ;};_agc :=_ee .Join (dirPath ,testName +"\u005fr\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u002e\u006a\u0073\u006f\u006e");_gcbf :=&ReferenceFile {_ed :_agc };_ad ,_eac :=_c .Open (_agc );if _gf .Is (_eac ,_c .ErrNotExist )&&createEmpty {_gcbf .Timestamp =_fb .Now ().UTC ();_gcbf .Map =&ReferenceMap {};return _gcbf ,nil ;};if _eac !=nil {return nil ,_eac ;};defer _ad .Close ();if _adb :=_fa .NewDecoder (_ad ).Decode (_gcbf );_adb !=nil {if _adb .Error ()=="i\u006c\u006c\u0065\u0067\u0061\u006c \u0062\u0061\u0073\u0065\u0036\u0034 \u0064\u0061\u0074\u0061\u0020\u0061\u0074 \u0069\u006e\u0070\u0075\u0074\u0020\u0062\u0079\u0074\u0065 \u0030"&&createEmpty {return _gcbf ,nil ;};return nil ,_adb ;};return _gcbf ,nil ;};func CopyFile (src ,dst string )error {_daa ,_ec :=_c .Open (src );if _ec !=nil {return _ec ;};defer _daa .Close ();_dbag ,_ec :=_c .Create (dst );if _ec !=nil {return _ec ;};defer _dbag .Close ();_ ,_ec =_b .Copy (_dbag ,_daa );return _ec ;};func (_fe *ReferenceFile )updateMap (_db *ReferenceMap )int {var _ag int ;if _fe .Map ._cgf ==nil {_fe .Map ._cgf =map[string ]ReferenceEntry {};};for _afb ,_bbf :=range _db ._cgf {_aa ,_cg :=_fe .Map ._cgf [_afb ];if !_cg {_fe .Map ._cgf [_afb ]=_bbf ;_ag ++;continue ;};if string (_aa .Value )!=string (_bbf .Value ){_fe .Map ._cgf [_afb ]=_bbf ;_ag ++;};};for _ga :=range _fe .Map ._cgf {if _ ,_fg :=_db ._cgf [_ga ];!_fg {delete (_fe .Map ._cgf ,_ga );_ag ++;};};return _ag ;};func HashFile (file string )(string ,error ){_aab ,_edf :=_c .Open (file );if _edf !=nil {return "",_edf ;};defer _aab .Close ();_ge :=_fdf .New ();if _ ,_edf =_b .Copy (_ge ,_aab );_edf !=nil {return "",_edf ;};return _bc .EncodeToString (_ge .Sum (nil )),nil ;};func (_gcb *ReferenceMap )MarshalJSON ()([]byte ,error ){return _fa .Marshal (_gcb ._cgf )};func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_cfd ,_cbc :=HashFile (file1 );if _cbc !=nil {return false ,_cbc ;};_ecb ,_cbc :=HashFile (file2 );if _cbc !=nil {return false ,_cbc ;};if _cfd ==_ecb {return true ,nil ;};_dd ,_cbc :=ReadPNG (file1 );if _cbc !=nil {return false ,_cbc ;};_bbc ,_cbc :=ReadPNG (file2 );if _cbc !=nil {return false ,_cbc ;};if _dd .Bounds ()!=_bbc .Bounds (){return false ,nil ;};return CompareImages (_dd ,_bbc );};func (_bbe *ReferenceMap )Read (key string )(ReferenceEntry ,bool ){_bbe .Lock ();defer _bbe .Unlock ();if _bbe ._cgf ==nil {return ReferenceEntry {},false ;};_bgf ,_dba :=_bbe ._cgf [key ];return _bgf ,_dba ;};func (_ce *ReferenceMap )Copy ()*ReferenceMap {_gff :=ReferenceMap {_cgf :make (map[string ]ReferenceEntry ,len (_ce ._cgf ))};for _dc ,_fed :=range _ce ._cgf {_gff ._cgf [_dc ]=_fed ;};return &_gff ;};var (ErrRenderNotSupported =_gf .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m");ErrImageSizeNotMatch =_gf .New ("\u0069\u006d\u0061ge\u0020\u0073\u0069\u007a\u0065\u0073\u0020\u0064\u006f\u006e\u0027\u0074\u0020\u006d\u0061\u0074\u0063\u0068"););func _fde (_dca *_d .T ,_dea string )int64 {_edbb ,_abf :=_c .Stat (_dea );_bb .NoError (_dca ,_abf );return _edbb .Size ();};func _ccc (_cfe ,_ba _a .Rectangle )bool {return _cfe .Min .X ==_ba .Min .X &&_cfe .Min .Y ==_ba .Min .Y &&_cfe .Max .X ==_ba .Max .X &&_cfe .Max .Y ==_ba .Max .Y ;};func ReadPNG (file string )(_a .Image ,error ){_fc ,_aed :=_c .Open (file );if _aed !=nil {return nil ,_aed ;};defer _fc .Close ();return _gb .Decode (_fc );};func _abe (_ecec *_a .RGBA ,_fgbe string ,_dg string ,_dge ,_acad int )error {_aee ,_cef :=_e .ReadFile (_fgbe );if _cef !=nil {return _cef ;};_ecbb ,_cef :=_fab .Parse (_aee );if _cef !=nil {return _cef ;};_aga ,_cef :=_fab .NewFace (_ecbb ,&_fab .FaceOptions {Size :18,DPI :72,Hinting :_gd .HintingNone });if _cef !=nil {return _cef ;};_efa :=&_gd .Drawer {Dst :_ecec ,Src :_a .NewUniform (_ef .RGBA {200,100,0,255}),Face :_aga ,Dot :_dee .P (_dge ,_acad )};_efa .DrawString (_dg );return nil ;};func (_bd *ReferenceFile )Update (currentMap *ReferenceMap )error {if _bd ._ed ==""{return nil ;};_eeg :=_bd .updateMap (currentMap );if _eeg ==0{return nil ;};_gdf ,_ae :=_c .OpenFile (_bd ._ed ,_c .O_CREATE |_c .O_TRUNC |_c .O_WRONLY ,0664);if _ae !=nil {return _ae ;};defer _gdf .Close ();_bd .Timestamp =_fb .Now ().UTC ();_cd :=_fa .NewEncoder (_gdf );_cd .SetIndent ("","\u0009");return _cd .Encode (_bd );};func RunRenderTest (t *_d .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ,currentHashMap *ReferenceMap ,refFile *ReferenceFile ){_fabc :=_g .TrimSuffix (_ee .Base (pdfPath ),_ee .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_aadf *_d .T ){_ggf :=_ee .Join (outputDir ,_fabc );_geg :=_ggf +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _gbfa :=RenderPDFToPNGs (pdfPath ,0,_geg );_gbfa !=nil {_aadf .Skip (_gbfa );};_bf :=_fabc +"\u002em\u0073\u0077\u006f\u0072\u0064";_dde :=_ee .Join (outputDir ,_bf );_bec :=_dde +"\u002d%\u0064\u002e\u0070\u006e\u0067";_dfg :=false ;if saveBaseline {_cgb :=_ee .Dir (pdfPath );_cff :=_ee .Join (_cgb ,_bf +"\u002e\u0070\u0064\u0066");if _ ,_abd :=_c .Stat (_cff );_abd ==nil {_aadf .Logf ("\u0052\u0065\u006e\u0064er\u0020\u004d\u0053\u0020\u0057\u006f\u0072\u0064\u0020\u0050\u0044\u0046\u003a\u0020%\u0076",_cff );if _ccb :=RenderPDFToPNGs (_cff ,0,_bec );_ccb !=nil {_aadf .Skip (_ccb );};_dfg =true ;};};for _bad :=1;true ;_bad ++{_ecg :=_fd .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_ggf ,_bad );_gge :=_ee .Join (baselineRenderPath ,_fd .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_fabc ,_bad ));if _ ,_gca :=_c .Stat (_ecg );_gca !=nil {break ;};_aadf .Logf ("\u0025\u0073",_gge );if saveBaseline {_aadf .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_ecg ,_gge );_eff :=CopyFile (_ecg ,_gge );if _eff !=nil {_aadf .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_gge ,_eff );};if _dfg {_gdd :=_fd .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_dde ,_bad );_bae :=_ee .Join (baselineRenderPath ,_fd .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_bf ,_bad ));_aadf .Logf ("\u0043\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u004d\u0053\u0020\u0057\u006f\u0072\u0064 \u0072e\u0073\u0075\u006c\u0074\u0073\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_gdd ,_bae );_dfgd :=CopyFile (_gdd ,_bae );if _dfgd !=nil {_aadf .Logf ("\u0045\u0052RO\u0052\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076\u002c \u0068\u0061\u0076\u0069n\u0067\u0020\u0064\u0069ff\u0065\u0072e\u006et\u0020\u0070\u0061\u0067\u0065\u0020s\u0069\u007a\u0065 \u0072\u0065s\u0075\u006c\u0074\u0073\u002c\u0020\u0075\u0073\u0065\u0020\u0070\u0072\u0065\u0076i\u006f\u0075\u0073\u0020\u0070\u0061\u0067\u0065",_bae ,_dfgd );_gdd =_fd .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_dde ,_bad -1);_bae =_ee .Join (baselineRenderPath ,_fd .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_bf ,_bad -1));if _cfff :=CopyFile (_gdd ,_bae );_cfff !=nil {_aadf .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_bae ,_cfff );};};_aadf .Logf ("\u0043\u006f\u006d\u0062\u0069\u006e\u0069\u006e\u0067\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063e\u0020\u0072\u0065\u0073\u0075\u006c\u0074s\u0020\u0077\u0069\u0074\u0068\u0020\u004d\u0053\u0020\u0057\u006fr\u0064\u0020\u0025\u0073\u0020\u0061\u006e\u0064\u0020\u0025\u0073",_gge ,_bae );_ebg ,_dfgd :=CombinePNGFiles (_gge ,_bae );if _c .IsNotExist (_dfgd ){_aadf .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_ebg {_aadf .Fatal ("\u0055n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0063\u006f\u006db\u0069\u006e\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0073");};_aadf .Logf ("Cr\u0065\u0061t\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065 \u0064\u0069\u0066\u0066\u0020\u0055\u006e\u0069\u004f\u0066\u0066\u0069\u0063\u0065\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u0020\u0077\u0069\u0074\u0068\u0020M\u0053 \u0057\u006f\u0072\u0064\u0020\u0025\u0073\u0020a\u006ed\u0020\u0025s",_gge ,_bae );_ebg ,_ebd ,_gea ,_edgc ,_dfgd :=CreatePNGDiff (_gge ,_bae );if _dfgd !=nil &&_dfgd !=ErrImageSizeNotMatch {_aadf .Fatalf ("\u0045\u0072\u0072\u006fr\u0020\u006f\u006e\u0020\u0063\u0072\u0065\u0061\u0074\u0065 \u0050N\u0047\u0020\u0044\u0069\u0066\u0066\u003a \u0025\u0076",_dfgd );};if _ebg {_aadf .Logf ("\u0049\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073\u000a",_ebd );_aadf .Logf ("D\u0069\u0066\u0066\u0020Pe\u0072c\u0065\u006e\u0074\u003a\u0020%\u0032\u002e\u0066\u0025\u0025\u000a",_gea );_aadf .Logf ("\u0044i\u0066f\u0020\u0054\u006f\u0074\u0061\u006c\u003a\u0020\u0025\u0066\u000a",_edgc );_eaca :=_ee .Base (_ebd );if _cfeb ,_ddf :=currentHashMap .Read (_eaca );_ddf {if _cfeb .DiffPercent < _gea ||_cfeb .DiffTotal < _edgc {_aadf .Fatalf ("\u004e\u0065\u0077\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0073\u0020\u0068\u0061\u0076\u0069\u006e\u0067 h\u0069g\u0068\u0065\u0072\u0020\u0064\u0069\u0066\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0070\u0065\u0072\u0063\u0065\u006e\u0074\u003a\u0020\u0025\u0066\u0020\u0061\u006e\u0064 \u0074\u006f\u0074\u0061\u006c\u0020\u0025\u0066\u000a",_gea ,_edgc );};};_gfc ,_gbfd :=HashFile (_ebd );_bb .NoError (_aadf ,_gbfd );_gee :=ReferenceEntry {Timestamp :_fb .Now ().UTC ().Unix (),Value :_gfc ,ResultSize :_fde (_aadf ,_ebd ),DiffPercent :_gea ,DiffTotal :_edgc };currentHashMap .Write (_eaca ,_gee );if _gbfd =refFile .Update (currentHashMap );_gbfd !=nil {_aadf .Logf ("\u0055\u0070\u0064\u0061\u0074\u0065\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063e\u0020f\u0069\u006c\u0065\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_gbfd );};};};continue ;};_aadf .Run (_fd .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_bad ),func (_cfebg *_d .T ){_cfebg .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_ecg ,_gge );_fdd ,_ecca :=ComparePNGFiles (_ecg ,_gge );if _c .IsNotExist (_ecca ){_cfebg .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_fdd {_cfebg .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");};});};});};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;};if _ ,_bdg :=_de .LookPath ("\u0067\u0073");_bdg !=nil {return ErrRenderNotSupported ;};return _de .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_fd .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();};func (_ac *ReferenceMap )Keys ()(_cb []string ){_cb =make ([]string ,len (_ac ._cgf ));var _acc int ;for _ea :=range _ac ._cgf {_cb [_acc ]=_ea ;_acc ++;};return _cb ;};func CombinePNGFiles (file1 ,file2 string )(bool ,error ){_gad ,_fbd :=ReadPNG (file1 );if _fbd !=nil {return false ,_fbd ;};_dae ,_fbd :=ReadPNG (file2 );if _fbd !=nil {return false ,_fbd ;};_gba :=_a .Point {_gad .Bounds ().Dx (),0};_ceb :=_a .Rectangle {_gba ,_gba .Add (_dae .Bounds ().Size ())};_edg :=_a .Rectangle {_a .Point {0,0},_ceb .Max };_gbg :=_a .NewRGBA (_edg );_bg .Draw (_gbg ,_gad .Bounds (),_gad ,_a .Point {0,0},_bg .Src );_bg .Draw (_gbg ,_ceb ,_dae ,_a .Point {0,0},_bg .Src );_ced :=_ee .Dir (file1 );_dab :=_g .TrimSuffix (_ee .Base (file1 ),_ee .Ext (file1 ));_df ,_fbd :=_c .Create (_ced +"\u002f"+_dab +"\u002d\u0063\u006f\u006d\u0062\u0069\u006e\u0065\u0064\u002e\u0070\u006e\u0067");if _fbd !=nil {return false ,_fbd ;};defer _df .Close ();if _gbf :=_gb .Encode (_df ,_gbg );_gbf !=nil {return false ,_gbf ;};return true ,nil ;};type ReferenceMap struct{_af .Mutex ;_cgf map[string ]ReferenceEntry ;};func _fcc (_dfa ,_aff _ef .Color )bool {_edc ,_cebc ,_baa ,_fga :=_dfa .RGBA ();_eeec ,_eea ,_eb ,_ece :=_aff .RGBA ();return _edc ==_eeec &&_cebc ==_eea &&_baa ==_eb &&_fga ==_ece ;};type ReferenceFile struct{Timestamp _fb .Time `json:"timestamp"`;Map *ReferenceMap `json:"map,omitempty"`;_ed string ;};