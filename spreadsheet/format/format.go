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

// Package format provides support for parsing and evaluating
// spreadsheetml/Excel number formats.
//
// Internally spreadsheets store numbers and dates values as a text
// representation of a floating point number (e.g. 1.2345).  This number is then
// displayed in Excel or another spreadsheet viewer differently depending on the
// number fornat of the cell style applied to the cell.
//
// As an example, the same value of 1.2345 can be displayed as:
// - "1" with format "0"
// - "1.2" with format "0.0"
// - "1.23" with format "0.00"
// - "1.235" with format "0.000"
// - "123%" with format "0%"
// - "1 23/100" with fornat "0 0/100"
// - "1.23E+00" with format "0.00E+00"
// - "29:37:41s" with format `[h]:mm:ss"s"`
package format ;import (_f "bytes";_aa "fmt";_aab "github.com/unidoc/unioffice/common/logger";_c "io";_e "math";_b "strconv";_cec "strings";_ce "time";);

// Value formats a value as a number or string depending on  if it appears to be
// a number or string.
func Value (v string ,f string )string {if IsNumber (v ){_dc ,_ :=_b .ParseFloat (v ,64);return Number (_dc ,f );};return String (v ,f );};func _gb (_ede []byte )[]byte {for _fde :=len (_ede )-1;_fde > 0;_fde --{if _ede [_fde ]=='9'+1{_ede [_fde ]='0';if _ede [_fde -1]=='.'{_fde --;};_ede [_fde -1]++;};};if _ede [0]=='9'+1{_ede [0]='0';copy (_ede [1:],_ede [0:]);_ede [0]='1';};return _ede ;};const (FmtTypeLiteral FmtType =iota ;FmtTypeDigit ;FmtTypeDigitOpt ;FmtTypeComma ;FmtTypeDecimal ;FmtTypePercent ;FmtTypeDollar ;FmtTypeDigitOptThousands ;FmtTypeUnderscore ;FmtTypeDate ;FmtTypeTime ;FmtTypeFraction ;FmtTypeText ;);const _cea ="\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u004c\u0069\u0074\u0065\u0072a\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0069\u0067\u0069\u0074\u0046\u006d\u0074\u0054y\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0043o\u006d\u006d\u0061\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0044\u0065\u0063\u0069\u006da\u006c\u0046\u006d\u0074\u0054\u0079\u0070\u0065Pe\u0072\u0063e\u006e\u0074\u0046\u006d\u0074\u0054\u0079\u0070e\u0044\u006f\u006c\u006c\u0061\u0072\u0046\u006d\u0074Ty\u0070\u0065\u0044i\u0067\u0069\u0074\u004f\u0070\u0074\u0054\u0068\u006f\u0075\u0073\u0061n\u0064\u0073\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0055n\u0064\u0065\u0072\u0073c\u006f\u0072\u0065\u0046\u006d\u0074T\u0079\u0070\u0065\u0044\u0061\u0074\u0065\u0046\u006d\u0074\u0054y\u0070e\u0054\u0069\u006d\u0065\u0046\u006d\u0074\u0054\u0079\u0070\u0065\u0046\u0072\u0061\u0063t\u0069\u006f\u006e\u0046\u006dt\u0054\u0079\u0070\u0065\u0054e\u0078\u0074";var _cf =[...]uint8 {0,14,26,41,53,67,81,94,118,135,146,157,172,183};func _dgf (_cae _ce .Time )_ce .Time {_cae =_cae .UTC ();return _ce .Date (_cae .Year (),_cae .Month (),_cae .Day (),_cae .Hour (),_cae .Minute (),_cae .Second (),_cae .Nanosecond (),_ce .Local );};func (_ecg *Lexer )nextFmt (){_ecg ._efc =append (_ecg ._efc ,_ecg ._daa );_ecg ._daa =Format {}};func _aafd (_cbd int64 )int64 {if _cbd < 0{return -_cbd ;};return _cbd ;};const _eeb int =-1;

// Format is a parsed number format.
type Format struct{Whole []Token ;Fractional []Token ;Exponent []Token ;IsExponential bool ;_af bool ;_ac bool ;_ae bool ;_g bool ;_gc bool ;_acg bool ;_acgf int64 ;_eg int ;};const _bec int =34;func _cg (_cgd float64 ,_eb Format ,_bbcg bool )string {if _eb ._ae {return NumberGeneric (_cgd );};_gg :=make ([]byte ,0,20);_gcg :=_e .Signbit (_cgd );_fb :=_e .Abs (_cgd );_gfg :=int64 (0);_aac :=int64 (0);if _eb .IsExponential {for _fb >=10{_aac ++;_fb /=10;};for _fb < 1{_aac --;_fb *=10;};}else if _eb ._ac {_fb *=100;}else if _eb ._af {if _eb ._acgf ==0{_ga :=_e .Pow (10,float64 (_eb ._eg ));_ee ,_ed :=1.0,1.0;_ =_ee ;for _cag :=1.0;_cag < _ga ;_cag ++{_ ,_ag :=_e .Modf (_fb *float64 (_cag ));if _ag < _ed {_ed =_ag ;_ee =_cag ;if _ag ==0{break ;};};};_eb ._acgf =int64 (_ee );};_gfg =int64 (_fb *float64 (_eb ._acgf )+0.5);if len (_eb .Whole )> 0&&_gfg > _eb ._acgf {_gfg =int64 (_fb *float64 (_eb ._acgf ))%_eb ._acgf ;_fb -=float64 (_gfg )/float64 (_eb ._acgf );}else {_fb -=float64 (_gfg )/float64 (_eb ._acgf );if _e .Abs (_fb )< 1{_bfe :=true ;for _ ,_gce :=range _eb .Whole {if _gce .Type ==FmtTypeDigitOpt {continue ;};if _gce .Type ==FmtTypeLiteral &&_gce .Literal ==' '{continue ;};_bfe =false ;};if _bfe {_eb .Whole =nil ;};};};};_acgfb :=1;for _ ,_fdf :=range _eb .Fractional {if _fdf .Type ==FmtTypeDigit ||_fdf .Type ==FmtTypeDigitOpt {_acgfb ++;};};_fb +=5*_e .Pow10 (-_acgfb );_bcg ,_ebg :=_e .Modf (_fb );_gg =append (_gg ,_acga (_bcg ,_cgd ,_eb )...);_gg =append (_gg ,_dca (_ebg ,_cgd ,_eb )...);_gg =append (_gg ,_cdb (_aac ,_eb )...);if _eb ._af {_gg =_b .AppendInt (_gg ,_gfg ,10);_gg =append (_gg ,'/');_gg =_b .AppendInt (_gg ,_eb ._acgf ,10);};if !_bbcg &&_gcg {return "\u002d"+string (_gg );};return string (_gg );};const _ca =1e11;const _fc =1e-10;func _fec (_cdbe []byte )[]byte {_cagb :=len (_cdbe );_egd :=false ;_gff :=false ;for _add :=len (_cdbe )-1;_add >=0;_add --{if _cdbe [_add ]=='0'&&!_gff &&!_egd {_cagb =_add ;}else if _cdbe [_add ]=='.'{_egd =true ;}else {_gff =true ;};};if _egd &&_gff {if _cdbe [_cagb -1]=='.'{_cagb --;};return _cdbe [0:_cagb ];};return _cdbe ;};func (_ecga *Lexer )Lex (r _c .Reader ){_ffag ,_addb ,_dcac :=0,0,0;_cabb :=-1;_abb ,_bgf ,_acba :=0,0,0;_ =_bgf ;_ =_acba ;_fdg :=1;_ =_fdg ;_gfgf :=make ([]byte ,4096);_ade :=false ;for !_ade {_dde :=0;if _abb > 0{_dde =_addb -_abb ;};_addb =0;_gdag ,_dgg :=r .Read (_gfgf [_dde :]);if _gdag ==0||_dgg !=nil {_ade =true ;};_dcac =_gdag +_dde ;if _dcac < len (_gfgf ){_cabb =_dcac ;};{_ffag =_gda ;_abb =0;_bgf =0;_acba =0;};{if _addb ==_dcac {goto _cbdb ;};switch _ffag {case 34:goto _bdc ;case 35:goto _bac ;case 0:goto _bcb ;case 36:goto _gdab ;case 37:goto _fcba ;case 1:goto _afeb ;case 2:goto _eged ;case 38:goto _gef ;case 3:goto _gdad ;case 4:goto _cgab ;case 39:goto _beca ;case 5:goto _cbfe ;case 6:goto _feac ;case 7:goto _gdf ;case 8:goto _bdcg ;case 40:goto _egf ;case 9:goto _acge ;case 41:goto _gffa ;case 10:goto _agd ;case 42:goto _ecc ;case 11:goto _fcbb ;case 43:goto _begd ;case 44:goto _acf ;case 45:goto _fegd ;case 12:goto _geda ;case 46:goto _gdbg ;case 13:goto _fcd ;case 14:goto _fcfb ;case 15:goto _aage ;case 16:goto _ddf ;case 47:goto _bfd ;case 17:goto _fdfa ;case 48:goto _bfef ;case 18:goto _fgcg ;case 19:goto _bbaa ;case 20:goto _dgbf ;case 49:goto _aaga ;case 50:goto _gdg ;case 21:goto _gdbe ;case 22:goto _baca ;case 23:goto _gec ;case 24:goto _dgd ;case 25:goto _edcc ;case 51:goto _cabg ;case 26:goto _acbad ;case 52:goto _ddfc ;case 53:goto _agc ;case 54:goto _baa ;case 55:goto _gbdb ;case 56:goto _cfge ;case 57:goto _eeaf ;case 27:goto _gdge ;case 28:goto _ddbf ;case 29:goto _gagf ;case 30:goto _dgcd ;case 31:goto _fed ;case 58:goto _bbga ;case 32:goto _ead ;case 59:goto _dfe ;case 33:goto _fgb ;case 60:goto _ece ;case 61:goto _ffd ;case 62:goto _abf ;};goto _fba ;_gbd :switch _acba {case 2:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypeDigit ,nil );};case 3:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypeDigitOpt ,nil );};case 5:{_addb =(_bgf )-1;};case 8:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypePercent ,nil );};case 13:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypeFraction ,_gfgf [_abb :_bgf ]);};case 14:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypeDate ,_gfgf [_abb :_bgf ]);};case 15:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypeTime ,_gfgf [_abb :_bgf ]);};case 16:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypeTime ,_gfgf [_abb :_bgf ]);};case 18:{_addb =(_bgf )-1;};case 20:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypeLiteral ,_gfgf [_abb :_bgf ]);};case 21:{_addb =(_bgf )-1;_ecga ._daa .AddToken (FmtTypeLiteral ,_gfgf [_abb +1:_bgf -1]);};};goto _cbe ;_feg :_addb =(_bgf )-1;{_ecga ._daa .AddToken (FmtTypeFraction ,_gfgf [_abb :_bgf ]);};goto _cbe ;_ege :_addb =(_bgf )-1;{_ecga ._daa .AddToken (FmtTypeDigitOpt ,nil );};goto _cbe ;_dba :_bgf =_addb +1;{_ecga ._daa .AddToken (FmtTypeDigitOptThousands ,nil );};goto _cbe ;_gag :_addb =(_bgf )-1;{_ecga ._daa .AddToken (FmtTypePercent ,nil );};goto _cbe ;_fegc :_addb =(_bgf )-1;{_ecga ._daa .AddToken (FmtTypeDate ,_gfgf [_abb :_bgf ]);};goto _cbe ;_gfcg :_addb =(_bgf )-1;{_ecga ._daa .AddToken (FmtTypeDigit ,nil );};goto _cbe ;_bgg :_addb =(_bgf )-1;{_ecga ._daa .AddToken (FmtTypeTime ,_gfgf [_abb :_bgf ]);};goto _cbe ;_eeca :_addb =(_bgf )-1;{_ecga ._daa .AddToken (FmtTypeLiteral ,_gfgf [_abb :_bgf ]);};goto _cbe ;_bfa :_bgf =_addb +1;{_ecga ._daa ._ae =true ;};goto _cbe ;_cgb :_bgf =_addb +1;{_ecga ._daa .AddToken (FmtTypeLiteral ,_gfgf [_abb :_bgf ]);};goto _cbe ;_gcgf :_bgf =_addb +1;{_ecga ._daa .AddToken (FmtTypeDollar ,nil );};goto _cbe ;_aea :_bgf =_addb +1;{_ecga ._daa .AddToken (FmtTypeComma ,nil );};goto _cbe ;_dfd :_bgf =_addb +1;{_ecga ._daa .AddToken (FmtTypeDecimal ,nil );};goto _cbe ;_ebb :_bgf =_addb +1;{_ecga .nextFmt ();};goto _cbe ;_fcf :_bgf =_addb +1;{_ecga ._daa .AddToken (FmtTypeText ,nil );};goto _cbe ;_ecf :_bgf =_addb +1;{_ecga ._daa .AddToken (FmtTypeUnderscore ,nil );};goto _cbe ;_eag :_bgf =_addb ;_addb --;{_ecga ._daa .AddToken (FmtTypeLiteral ,_gfgf [_abb :_bgf ]);};goto _cbe ;_edc :_bgf =_addb ;_addb --;{_ecga ._daa .AddToken (FmtTypeLiteral ,_gfgf [_abb +1:_bgf -1]);};goto _cbe ;_efgd :_bgf =_addb ;_addb --;{_ecga ._daa .AddToken (FmtTypeDigitOpt ,nil );};goto _cbe ;_cfg :_bgf =_addb ;_addb --;{_ecga ._daa .AddToken (FmtTypeFraction ,_gfgf [_abb :_bgf ]);};goto _cbe ;_addg :_bgf =_addb ;_addb --;{_ecga ._daa .AddToken (FmtTypePercent ,nil );};goto _cbe ;_cgc :_bgf =_addb ;_addb --;{_ecga ._daa .AddToken (FmtTypeDate ,_gfgf [_abb :_bgf ]);};goto _cbe ;_cga :_bgf =_addb ;_addb --;{_ecga ._daa .AddToken (FmtTypeDigit ,nil );};goto _cbe ;_fega :_bgf =_addb ;_addb --;{_ecga ._daa .AddToken (FmtTypeTime ,_gfgf [_abb :_bgf ]);};goto _cbe ;_gcfc :_bgf =_addb ;_addb --;{};goto _cbe ;_ebc :_bgf =_addb +1;{_ecga ._daa .IsExponential =true ;};goto _cbe ;_gdef :_bgf =_addb +1;{_ecga ._daa .AddToken (FmtTypeLiteral ,_gfgf [_abb +1:_bgf ]);};goto _cbe ;_cbe :_abb =0;if _addb ++;_addb ==_dcac {goto _fbab ;};_bdc :_abb =_addb ;switch _gfgf [_addb ]{case 34:goto _gfgd ;case 35:goto _dea ;case 36:goto _gcgf ;case 37:goto _aabd ;case 44:goto _aea ;case 46:goto _dfd ;case 47:goto _fade ;case 48:goto _gbc ;case 58:goto _def ;case 59:goto _ebb ;case 63:goto _dgc ;case 64:goto _fcf ;case 65:goto _cbb ;case 69:goto _dbc ;case 71:goto _efe ;case 91:goto _gfdg ;case 92:goto _abg ;case 95:goto _ecf ;case 100:goto _fade ;case 104:goto _def ;case 109:goto _ebf ;case 115:goto _bgbde ;case 121:goto _ffae ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _ced ;};goto _cgb ;_gfgd :_bgf =_addb +1;_acba =20;goto _fded ;_fded :if _addb ++;_addb ==_dcac {goto _cfbg ;};_bac :if _gfgf [_addb ]==34{goto _bdbc ;};goto _beb ;_beb :if _addb ++;_addb ==_dcac {goto _gba ;};_bcb :if _gfgf [_addb ]==34{goto _bdbc ;};goto _beb ;_bdbc :_bgf =_addb +1;_acba =21;goto _afe ;_afe :if _addb ++;_addb ==_dcac {goto _gaf ;};_gdab :if _gfgf [_addb ]==34{goto _beb ;};goto _edc ;_dea :_bgf =_addb +1;_acba =3;goto _gdd ;_gdd :if _addb ++;_addb ==_dcac {goto _fbd ;};_fcba :switch _gfgf [_addb ]{case 35:goto _bbab ;case 37:goto _bbab ;case 44:goto _acaa ;case 47:goto _fcbd ;case 48:goto _bbab ;case 63:goto _bbab ;};goto _efgd ;_bbab :if _addb ++;_addb ==_dcac {goto _ecfa ;};_afeb :switch _gfgf [_addb ]{case 35:goto _bbab ;case 37:goto _bbab ;case 47:goto _fcbd ;case 48:goto _bbab ;case 63:goto _bbab ;};goto _gbd ;_fcbd :if _addb ++;_addb ==_dcac {goto _aefd ;};_eged :switch _gfgf [_addb ]{case 35:goto _dce ;case 37:goto _fddd ;case 48:goto _eaf ;case 63:goto _dce ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _ecgag ;};goto _gbd ;_dce :_bgf =_addb +1;goto _fdac ;_fdac :if _addb ++;_addb ==_dcac {goto _dbb ;};_gef :switch _gfgf [_addb ]{case 35:goto _dce ;case 37:goto _dce ;case 44:goto _dce ;case 46:goto _dce ;case 48:goto _dce ;case 63:goto _dce ;case 65:goto _bdfb ;};goto _cfg ;_bdfb :if _addb ++;_addb ==_dcac {goto _agb ;};_gdad :switch _gfgf [_addb ]{case 47:goto _fef ;case 77:goto _bae ;};goto _feg ;_fef :if _addb ++;_addb ==_dcac {goto _ecge ;};_cgab :if _gfgf [_addb ]==80{goto _gfcd ;};goto _feg ;_gfcd :_bgf =_addb +1;goto _gdac ;_gdac :if _addb ++;_addb ==_dcac {goto _bge ;};_beca :if _gfgf [_addb ]==65{goto _bdfb ;};goto _cfg ;_bae :if _addb ++;_addb ==_dcac {goto _dcdg ;};_cbfe :if _gfgf [_addb ]==47{goto _cgdg ;};goto _feg ;_cgdg :if _addb ++;_addb ==_dcac {goto _gbg ;};_feac :if _gfgf [_addb ]==80{goto _dae ;};goto _feg ;_dae :if _addb ++;_addb ==_dcac {goto _cba ;};_gdf :if _gfgf [_addb ]==77{goto _gfcd ;};goto _feg ;_fddd :if _addb ++;_addb ==_dcac {goto _agg ;};_bdcg :switch _gfgf [_addb ]{case 35:goto _cgac ;case 37:goto _gae ;case 63:goto _cgac ;};if 48<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _cge ;};goto _gbd ;_cgac :_bgf =_addb +1;goto _aeg ;_aeg :if _addb ++;_addb ==_dcac {goto _ecb ;};_egf :switch _gfgf [_addb ]{case 35:goto _dce ;case 37:goto _age ;case 44:goto _dce ;case 46:goto _dce ;case 48:goto _dce ;case 63:goto _dce ;case 65:goto _bdfb ;};goto _cfg ;_age :if _addb ++;_addb ==_dcac {goto _ceg ;};_acge :switch _gfgf [_addb ]{case 35:goto _dcacb ;case 44:goto _dcacb ;case 46:goto _dcacb ;case 48:goto _dcacb ;case 63:goto _dcacb ;};goto _feg ;_dcacb :_bgf =_addb +1;goto _eac ;_eac :if _addb ++;_addb ==_dcac {goto _becaa ;};_gffa :switch _gfgf [_addb ]{case 35:goto _dcacb ;case 44:goto _dcacb ;case 46:goto _dcacb ;case 48:goto _dcacb ;case 63:goto _dcacb ;case 65:goto _bdfb ;};goto _cfg ;_gae :if _addb ++;_addb ==_dcac {goto _eegd ;};_agd :if _gfgf [_addb ]==37{goto _gae ;};if 48<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _cge ;};goto _gbd ;_cge :_bgf =_addb +1;_acba =13;goto _ggdd ;_ggdd :if _addb ++;_addb ==_dcac {goto _efb ;};_ecc :switch _gfgf [_addb ]{case 35:goto _dce ;case 37:goto _efcf ;case 44:goto _dce ;case 46:goto _dce ;case 48:goto _dadc ;case 63:goto _dce ;case 65:goto _bdfb ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _cge ;};goto _cfg ;_efcf :if _addb ++;_addb ==_dcac {goto _ecbg ;};_fcbb :switch _gfgf [_addb ]{case 35:goto _dcacb ;case 37:goto _gae ;case 44:goto _dcacb ;case 46:goto _dcacb ;case 63:goto _dcacb ;};if 48<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _cge ;};goto _feg ;_dadc :_bgf =_addb +1;goto _bcaa ;_bcaa :if _addb ++;_addb ==_dcac {goto _bcd ;};_begd :switch _gfgf [_addb ]{case 35:goto _dce ;case 37:goto _dadc ;case 44:goto _dce ;case 46:goto _dce ;case 48:goto _dadc ;case 63:goto _dce ;case 65:goto _bdfb ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _cge ;};goto _cfg ;_eaf :_bgf =_addb +1;goto _acgag ;_acgag :if _addb ++;_addb ==_dcac {goto _gdda ;};_acf :switch _gfgf [_addb ]{case 35:goto _dce ;case 37:goto _dadc ;case 44:goto _dce ;case 46:goto _dce ;case 48:goto _eaf ;case 63:goto _dce ;case 65:goto _bdfb ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _ecgag ;};goto _cfg ;_ecgag :_bgf =_addb +1;goto _geff ;_geff :if _addb ++;_addb ==_dcac {goto _agfe ;};_fegd :switch _gfgf [_addb ]{case 35:goto _dce ;case 37:goto _cge ;case 44:goto _dce ;case 46:goto _dce ;case 48:goto _eaf ;case 63:goto _dce ;case 65:goto _bdfb ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _ecgag ;};goto _cfg ;_acaa :if _addb ++;_addb ==_dcac {goto _gabg ;};_geda :if _gfgf [_addb ]==35{goto _dba ;};goto _ege ;_aabd :_bgf =_addb +1;_acba =8;goto _egeg ;_egeg :if _addb ++;_addb ==_dcac {goto _gecf ;};_gdbg :switch _gfgf [_addb ]{case 35:goto _faa ;case 37:goto _ddb ;case 48:goto _fag ;case 63:goto _faa ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _bda ;};goto _addg ;_faa :if _addb ++;_addb ==_dcac {goto _egga ;};_fcd :switch _gfgf [_addb ]{case 35:goto _faa ;case 47:goto _fcbd ;case 48:goto _faa ;case 63:goto _faa ;};goto _gag ;_ddb :if _addb ++;_addb ==_dcac {goto _acc ;};_fcfb :if _gfgf [_addb ]==37{goto _ddb ;};if 48<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _bda ;};goto _gbd ;_bda :if _addb ++;_addb ==_dcac {goto _cbaa ;};_aage :switch _gfgf [_addb ]{case 37:goto _ddb ;case 47:goto _fcbd ;};if 48<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _bda ;};goto _gbd ;_fag :if _addb ++;_addb ==_dcac {goto _aae ;};_ddf :switch _gfgf [_addb ]{case 35:goto _faa ;case 37:goto _ddb ;case 47:goto _fcbd ;case 48:goto _fag ;case 63:goto _faa ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _bda ;};goto _gag ;_fade :_bgf =_addb +1;goto _gfb ;_gfb :if _addb ++;_addb ==_dcac {goto _cgbg ;};_bfd :switch _gfgf [_addb ]{case 47:goto _fade ;case 100:goto _fade ;case 109:goto _fade ;case 121:goto _aef ;};goto _cgc ;_aef :if _addb ++;_addb ==_dcac {goto _ega ;};_fdfa :if _gfgf [_addb ]==121{goto _fade ;};goto _fegc ;_gbc :_bgf =_addb +1;_acba =2;goto _bbdf ;_bbdf :if _addb ++;_addb ==_dcac {goto _facb ;};_bfef :switch _gfgf [_addb ]{case 35:goto _bbab ;case 37:goto _fae ;case 47:goto _fcbd ;case 48:goto _dgb ;case 63:goto _bbab ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _fefe ;};goto _cga ;_fae :if _addb ++;_addb ==_dcac {goto _dfb ;};_fgcg :switch _gfgf [_addb ]{case 35:goto _bbab ;case 37:goto _fae ;case 47:goto _fcbd ;case 48:goto _fae ;case 63:goto _bbab ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _bda ;};goto _gfcg ;_dgb :if _addb ++;_addb ==_dcac {goto _cfgf ;};_bbaa :switch _gfgf [_addb ]{case 35:goto _bbab ;case 37:goto _fae ;case 47:goto _fcbd ;case 48:goto _dgb ;case 63:goto _bbab ;};if 49<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _fefe ;};goto _gfcg ;_fefe :if _addb ++;_addb ==_dcac {goto _cadb ;};_dgbf :switch _gfgf [_addb ]{case 37:goto _bda ;case 47:goto _fcbd ;};if 48<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _fefe ;};goto _gbd ;_ced :_bgf =_addb +1;_acba =20;goto _acbd ;_acbd :if _addb ++;_addb ==_dcac {goto _ggc ;};_aaga :switch _gfgf [_addb ]{case 37:goto _bda ;case 47:goto _fcbd ;};if 48<=_gfgf [_addb ]&&_gfgf [_addb ]<=57{goto _fefe ;};goto _eag ;_def :_bgf =_addb +1;_acba =15;goto _cfbf ;_cfbf :if _addb ++;_addb ==_dcac {goto _aacf ;};_gdg :switch _gfgf [_addb ]{case 58:goto _def ;case 65:goto _aafda ;case 104:goto _def ;case 109:goto _def ;case 115:goto _bgbde ;};goto _fega ;_aafda :if _addb ++;_addb ==_dcac {goto _fgg ;};_gdbe :switch _gfgf [_addb ]{case 47:goto _afa ;case 77:goto _agfc ;};goto _gbd ;_afa :if _addb ++;_addb ==_dcac {goto _bgfb ;};_baca :if _gfgf [_addb ]==80{goto _def ;};goto _gbd ;_agfc :if _addb ++;_addb ==_dcac {goto _gdde ;};_gec :if _gfgf [_addb ]==47{goto _caf ;};goto _gbd ;_caf :if _addb ++;_addb ==_dcac {goto _bbdb ;};_dgd :if _gfgf [_addb ]==80{goto _fdgg ;};goto _gbd ;_fdgg :if _addb ++;_addb ==_dcac {goto _acbc ;};_edcc :if _gfgf [_addb ]==77{goto _def ;};goto _gbd ;_bgbde :_bgf =_addb +1;_acba =15;goto _gad ;_gad :if _addb ++;_addb ==_dcac {goto _cedf ;};_cabg :switch _gfgf [_addb ]{case 46:goto _aefe ;case 58:goto _def ;case 65:goto _aafda ;case 104:goto _def ;case 109:goto _def ;case 115:goto _bgbde ;};goto _fega ;_aefe :if _addb ++;_addb ==_dcac {goto _bdbe ;};_acbad :if _gfgf [_addb ]==48{goto _adg ;};goto _bgg ;_adg :_bgf =_addb +1;_acba =15;goto _fbeg ;_fbeg :if _addb ++;_addb ==_dcac {goto _bbcab ;};_ddfc :switch _gfgf [_addb ]{case 48:goto _ddfa ;case 58:goto _def ;case 65:goto _aafda ;case 104:goto _def ;case 109:goto _def ;case 115:goto _bgbde ;};goto _fega ;_ddfa :_bgf =_addb +1;_acba =15;goto _cdc ;_cdc :if _addb ++;_addb ==_dcac {goto _ffga ;};_agc :switch _gfgf [_addb ]{case 48:goto _def ;case 58:goto _def ;case 65:goto _aafda ;case 104:goto _def ;case 109:goto _def ;case 115:goto _bgbde ;};goto _fega ;_dgc :_bgf =_addb +1;_acba =5;goto _bfgg ;_bfgg :if _addb ++;_addb ==_dcac {goto _cafe ;};_baa :switch _gfgf [_addb ]{case 35:goto _bbab ;case 37:goto _bbab ;case 47:goto _fcbd ;case 48:goto _bbab ;case 63:goto _bbab ;};goto _gcfc ;_cbb :_bgf =_addb +1;_acba =20;goto _gdgb ;_gdgb :if _addb ++;_addb ==_dcac {goto _cca ;};_gbdb :switch _gfgf [_addb ]{case 47:goto _afa ;case 77:goto _agfc ;};goto _eag ;_dbc :if _addb ++;_addb ==_dcac {goto _adec ;};_cfge :switch _gfgf [_addb ]{case 43:goto _ebc ;case 45:goto _ebc ;};goto _eag ;_efe :_bgf =_addb +1;goto _bde ;_bde :if _addb ++;_addb ==_dcac {goto _eccd ;};_eeaf :if _gfgf [_addb ]==101{goto _bdae ;};goto _eag ;_bdae :if _addb ++;_addb ==_dcac {goto _efa ;};_gdge :if _gfgf [_addb ]==110{goto _bgfa ;};goto _eeca ;_bgfa :if _addb ++;_addb ==_dcac {goto _cgcdf ;};_ddbf :if _gfgf [_addb ]==101{goto _cedc ;};goto _eeca ;_cedc :if _addb ++;_addb ==_dcac {goto _gaa ;};_gagf :if _gfgf [_addb ]==114{goto _cad ;};goto _eeca ;_cad :if _addb ++;_addb ==_dcac {goto _acbdd ;};_dgcd :if _gfgf [_addb ]==97{goto _fdb ;};goto _eeca ;_fdb :if _addb ++;_addb ==_dcac {goto _ebcd ;};_fed :if _gfgf [_addb ]==108{goto _bfa ;};goto _eeca ;_gfdg :_bgf =_addb +1;_acba =20;goto _gdbd ;_gdbd :if _addb ++;_addb ==_dcac {goto _gac ;};_bbga :switch _gfgf [_addb ]{case 104:goto _eed ;case 109:goto _eed ;case 115:goto _eed ;};goto _eaa ;_eaa :if _addb ++;_addb ==_dcac {goto _ccd ;};_ead :if _gfgf [_addb ]==93{goto _bfdd ;};goto _eaa ;_bfdd :_bgf =_addb +1;_acba =18;goto _cgcd ;_aga :_bgf =_addb +1;_acba =16;goto _cgcd ;_cgcd :if _addb ++;_addb ==_dcac {goto _ffgc ;};_dfe :if _gfgf [_addb ]==93{goto _bfdd ;};goto _eaa ;_eed :if _addb ++;_addb ==_dcac {goto _cece ;};_fgb :if _gfgf [_addb ]==93{goto _aga ;};goto _eaa ;_abg :if _addb ++;_addb ==_dcac {goto _ceae ;};_ece :goto _gdef ;_ebf :_bgf =_addb +1;_acba =14;goto _fecd ;_fecd :if _addb ++;_addb ==_dcac {goto _gffc ;};_ffd :switch _gfgf [_addb ]{case 47:goto _fade ;case 58:goto _def ;case 65:goto _aafda ;case 100:goto _fade ;case 104:goto _def ;case 109:goto _ebf ;case 115:goto _bgbde ;case 121:goto _aef ;};goto _cgc ;_ffae :if _addb ++;_addb ==_dcac {goto _cagf ;};_abf :if _gfgf [_addb ]==121{goto _fade ;};goto _eag ;_fba :_fbab :_ffag =34;goto _cbdb ;_cfbg :_ffag =35;goto _cbdb ;_gba :_ffag =0;goto _cbdb ;_gaf :_ffag =36;goto _cbdb ;_fbd :_ffag =37;goto _cbdb ;_ecfa :_ffag =1;goto _cbdb ;_aefd :_ffag =2;goto _cbdb ;_dbb :_ffag =38;goto _cbdb ;_agb :_ffag =3;goto _cbdb ;_ecge :_ffag =4;goto _cbdb ;_bge :_ffag =39;goto _cbdb ;_dcdg :_ffag =5;goto _cbdb ;_gbg :_ffag =6;goto _cbdb ;_cba :_ffag =7;goto _cbdb ;_agg :_ffag =8;goto _cbdb ;_ecb :_ffag =40;goto _cbdb ;_ceg :_ffag =9;goto _cbdb ;_becaa :_ffag =41;goto _cbdb ;_eegd :_ffag =10;goto _cbdb ;_efb :_ffag =42;goto _cbdb ;_ecbg :_ffag =11;goto _cbdb ;_bcd :_ffag =43;goto _cbdb ;_gdda :_ffag =44;goto _cbdb ;_agfe :_ffag =45;goto _cbdb ;_gabg :_ffag =12;goto _cbdb ;_gecf :_ffag =46;goto _cbdb ;_egga :_ffag =13;goto _cbdb ;_acc :_ffag =14;goto _cbdb ;_cbaa :_ffag =15;goto _cbdb ;_aae :_ffag =16;goto _cbdb ;_cgbg :_ffag =47;goto _cbdb ;_ega :_ffag =17;goto _cbdb ;_facb :_ffag =48;goto _cbdb ;_dfb :_ffag =18;goto _cbdb ;_cfgf :_ffag =19;goto _cbdb ;_cadb :_ffag =20;goto _cbdb ;_ggc :_ffag =49;goto _cbdb ;_aacf :_ffag =50;goto _cbdb ;_fgg :_ffag =21;goto _cbdb ;_bgfb :_ffag =22;goto _cbdb ;_gdde :_ffag =23;goto _cbdb ;_bbdb :_ffag =24;goto _cbdb ;_acbc :_ffag =25;goto _cbdb ;_cedf :_ffag =51;goto _cbdb ;_bdbe :_ffag =26;goto _cbdb ;_bbcab :_ffag =52;goto _cbdb ;_ffga :_ffag =53;goto _cbdb ;_cafe :_ffag =54;goto _cbdb ;_cca :_ffag =55;goto _cbdb ;_adec :_ffag =56;goto _cbdb ;_eccd :_ffag =57;goto _cbdb ;_efa :_ffag =27;goto _cbdb ;_cgcdf :_ffag =28;goto _cbdb ;_gaa :_ffag =29;goto _cbdb ;_acbdd :_ffag =30;goto _cbdb ;_ebcd :_ffag =31;goto _cbdb ;_gac :_ffag =58;goto _cbdb ;_ccd :_ffag =32;goto _cbdb ;_ffgc :_ffag =59;goto _cbdb ;_cece :_ffag =33;goto _cbdb ;_ceae :_ffag =60;goto _cbdb ;_gffc :_ffag =61;goto _cbdb ;_cagf :_ffag =62;goto _cbdb ;_cbdb :{};if _addb ==_cabb {switch _ffag {case 35:goto _eag ;case 0:goto _gbd ;case 36:goto _edc ;case 37:goto _efgd ;case 1:goto _gbd ;case 2:goto _gbd ;case 38:goto _cfg ;case 3:goto _feg ;case 4:goto _feg ;case 39:goto _cfg ;case 5:goto _feg ;case 6:goto _feg ;case 7:goto _feg ;case 8:goto _gbd ;case 40:goto _cfg ;case 9:goto _feg ;case 41:goto _cfg ;case 10:goto _gbd ;case 42:goto _cfg ;case 11:goto _feg ;case 43:goto _cfg ;case 44:goto _cfg ;case 45:goto _cfg ;case 12:goto _ege ;case 46:goto _addg ;case 13:goto _gag ;case 14:goto _gbd ;case 15:goto _gbd ;case 16:goto _gag ;case 47:goto _cgc ;case 17:goto _fegc ;case 48:goto _cga ;case 18:goto _gfcg ;case 19:goto _gfcg ;case 20:goto _gbd ;case 49:goto _eag ;case 50:goto _fega ;case 21:goto _gbd ;case 22:goto _gbd ;case 23:goto _gbd ;case 24:goto _gbd ;case 25:goto _gbd ;case 51:goto _fega ;case 26:goto _bgg ;case 52:goto _fega ;case 53:goto _fega ;case 54:goto _gcfc ;case 55:goto _eag ;case 56:goto _eag ;case 57:goto _eag ;case 27:goto _eeca ;case 28:goto _eeca ;case 29:goto _eeca ;case 30:goto _eeca ;case 31:goto _eeca ;case 58:goto _eag ;case 32:goto _gbd ;case 59:goto _gbd ;case 33:goto _eeca ;case 60:goto _eag ;case 61:goto _cgc ;case 62:goto _eag ;};};};if _abb > 0{copy (_gfgf [0:],_gfgf [_abb :]);};};_ =_cabb ;if _ffag ==_cacg {_aab .Log .Debug ("\u0066o\u0072m\u0061\u0074\u0020\u0070\u0061r\u0073\u0065 \u0065\u0072\u0072\u006f\u0072");};};

// Number is used to format a number with a format string.  If the format
// string is empty, then General number formatting is used which attempts to mimic
// Excel's general formatting.
func Number (v float64 ,f string )string {if f ==""||f =="\u0047e\u006e\u0065\u0072\u0061\u006c"||f =="\u0040"{return NumberGeneric (v );};_egg :=Parse (f );if len (_egg )==1{return _cg (v ,_egg [0],false );}else if len (_egg )> 1&&v < 0{return _cg (v ,_egg [1],true );}else if len (_egg )> 2&&v ==0{return _cg (v ,_egg [2],false );};return _cg (v ,_egg [0],false );};

// String returns the string formatted according to the type.  In format strings
// this is the fourth item, where '@' is used as a placeholder for text.
func String (v string ,f string )string {_bf :=Parse (f );var _bc Format ;if len (_bf )==1{_bc =_bf [0];}else if len (_bf )==4{_bc =_bf [3];};_bb :=false ;for _ ,_fa :=range _bc .Whole {if _fa .Type ==FmtTypeText {_bb =true ;};};if !_bb {return v ;};_cd :=_f .Buffer {};for _ ,_aee :=range _bc .Whole {switch _aee .Type {case FmtTypeLiteral :_cd .WriteByte (_aee .Literal );case FmtTypeText :_cd .WriteString (v );};};return _cd .String ();};const _gedd int =0;func _acga (_da ,_ad float64 ,_cac Format )[]byte {if len (_cac .Whole )==0{return nil ;};_ef :=_ce .Date (1899,12,30,0,0,0,0,_ce .UTC );_adc :=_ef .Add (_ce .Duration (_ad *float64 (24*_ce .Hour )));_adc =_dgf (_adc );_bbd :=_b .AppendFloat (nil ,_da ,'f',-1,64);_dd :=make ([]byte ,0,len (_bbd ));_ge :=0;_fgf :=1;_df :for _agf :=len (_cac .Whole )-1;_agf >=0;_agf --{_gfd :=len (_bbd )-1-_ge ;_dee :=_cac .Whole [_agf ];switch _dee .Type {case FmtTypeDigit :if _gfd >=0{_dd =append (_dd ,_bbd [_gfd ]);_ge ++;_fgf =_agf ;}else {_dd =append (_dd ,'0');};case FmtTypeDigitOpt :if _gfd >=0{_dd =append (_dd ,_bbd [_gfd ]);_ge ++;_fgf =_agf ;}else {for _dfc :=_agf ;_dfc >=0;_dfc --{_edd :=_cac .Whole [_dfc ];if _edd .Type ==FmtTypeLiteral {_dd =append (_dd ,_edd .Literal );};};break _df ;};case FmtTypeDollar :for _fbb :=_ge ;_fbb < len (_bbd );_fbb ++{_dd =append (_dd ,_bbd [len (_bbd )-1-_fbb ]);_ge ++;};_dd =append (_dd ,'$');case FmtTypeComma :if !_cac ._g {_dd =append (_dd ,',');};case FmtTypeLiteral :_dd =append (_dd ,_dee .Literal );case FmtTypeDate :_dd =append (_dd ,_gd (_ffb (_adc ,_dee .DateTime ))...);case FmtTypeTime :_dd =append (_dd ,_gd (_cdda (_adc ,_ad ,_dee .DateTime ))...);default:_aab .Log .Debug ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070e\u0020i\u006e\u0020\u0077\u0068\u006f\u006c\u0065 \u0025\u0076",_dee );};};_db :=_gd (_dd );if _ge < len (_bbd )&&(_ge !=0||_cac ._acg ){_bfc :=len (_bbd )-_ge ;_dcg :=make ([]byte ,len (_db )+_bfc );copy (_dcg ,_db [0:_fgf ]);copy (_dcg [_fgf :],_bbd [0:]);copy (_dcg [_fgf +_bfc :],_db [_fgf :]);_db =_dcg ;};if _cac ._g {_ba :=_f .Buffer {};_bfb :=0;for _cde :=len (_db )-1;_cde >=0;_cde --{if !(_db [_cde ]>='0'&&_db [_cde ]<='9'){_bfb ++;}else {break ;};};for _eeg :=0;_eeg < len (_db );_eeg ++{_bfg :=(len (_db )-_eeg -_bfb );if _bfg %3==0&&_bfg !=0&&_eeg !=0{_ba .WriteByte (',');};_ba .WriteByte (_db [_eeg ]);};_db =_ba .Bytes ();};return _db ;};func _cdb (_bca int64 ,_ded Format )[]byte {if !_ded .IsExponential ||len (_ded .Exponent )==0{return nil ;};_bffb :=_b .AppendInt (nil ,_aafd (_bca ),10);_fdfe :=make ([]byte ,0,len (_bffb )+2);_fdfe =append (_fdfe ,'E');if _bca >=0{_fdfe =append (_fdfe ,'+');}else {_fdfe =append (_fdfe ,'-');_bca *=-1;};_dg :=0;_afb :for _ffc :=len (_ded .Exponent )-1;_ffc >=0;_ffc --{_cdd :=len (_bffb )-1-_dg ;_gcee :=_ded .Exponent [_ffc ];switch _gcee .Type {case FmtTypeDigit :if _cdd >=0{_fdfe =append (_fdfe ,_bffb [_cdd ]);_dg ++;}else {_fdfe =append (_fdfe ,'0');};case FmtTypeDigitOpt :if _cdd >=0{_fdfe =append (_fdfe ,_bffb [_cdd ]);_dg ++;}else {for _ebe :=_ffc ;_ebe >=0;_ebe --{_fcg :=_ded .Exponent [_ebe ];if _fcg .Type ==FmtTypeLiteral {_fdfe =append (_fdfe ,_fcg .Literal );};};break _afb ;};case FmtTypeLiteral :_fdfe =append (_fdfe ,_gcee .Literal );default:_aab .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020\u0065\u0078p\u0020\u0025\u0076",_gcee );};};if _dg < len (_bffb ){_fdfe =append (_fdfe ,_bffb [len (_bffb )-_dg -1:_dg -1]...);};_gd (_fdfe [2:]);return _fdfe ;};

// AddToken adds a format token to the format.
func (_d *Format )AddToken (t FmtType ,l []byte ){if _d ._gc {_d ._gc =false ;return ;};switch t {case FmtTypeDecimal :_d ._acg =true ;case FmtTypeUnderscore :_d ._gc =true ;case FmtTypeText :_d .Whole =append (_d .Whole ,Token {Type :t });case FmtTypeDate ,FmtTypeTime :_d .Whole =append (_d .Whole ,Token {Type :t ,DateTime :string (l )});case FmtTypePercent :_d ._ac =true ;t =FmtTypeLiteral ;l =[]byte {'%'};fallthrough;case FmtTypeDigitOpt :fallthrough;case FmtTypeLiteral ,FmtTypeDigit ,FmtTypeDollar ,FmtTypeComma :if l ==nil {l =[]byte {0};};for _ ,_aaf :=range l {if _d .IsExponential {_d .Exponent =append (_d .Exponent ,Token {Type :t ,Literal :_aaf });}else if !_d ._acg {_d .Whole =append (_d .Whole ,Token {Type :t ,Literal :_aaf });}else {_d .Fractional =append (_d .Fractional ,Token {Type :t ,Literal :_aaf });};};case FmtTypeDigitOptThousands :_d ._g =true ;case FmtTypeFraction :_cb :=_cec .Split (string (l ),"\u002f");if len (_cb )==2{_d ._af =true ;_d ._acgf ,_ =_b .ParseInt (_cb [1],10,64);for _ ,_afg :=range _cb [1]{if _afg =='?'||_afg =='0'{_d ._eg ++;};};};default:_aab .Log .Debug ("\u0075\u006e\u0073u\u0070\u0070\u006f\u0072t\u0065\u0064\u0020\u0070\u0068\u0020\u0074y\u0070\u0065\u0020\u0069\u006e\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0025\u0076",t );};};func _bdb (_fad float64 )string {_ec :=_b .FormatFloat (_fad ,'E',-1,64);_dad :=_b .FormatFloat (_fad ,'E',5,64);if len (_ec )< len (_dad ){return _b .FormatFloat (_fad ,'E',2,64);};return _dad ;};const _gda int =34;func (_fg FmtType )String ()string {if _fg >=FmtType (len (_cf )-1){return _aa .Sprintf ("F\u006d\u0074\u0054\u0079\u0070\u0065\u0028\u0025\u0064\u0029",_fg );};return _cea [_cf [_fg ]:_cf [_fg +1]];};

// Token is a format token in the Excel format string.
type Token struct{Type FmtType ;Literal byte ;DateTime string ;};func Parse (s string )[]Format {_fadb :=Lexer {};_fadb .Lex (_cec .NewReader (s ));_fadb ._efc =append (_fadb ._efc ,_fadb ._daa );return _fadb ._efc ;};

// FmtType is the type of a format token.
//go:generate stringer -type=FmtType
type FmtType byte ;func _ffb (_ggb _ce .Time ,_bg string )[]byte {_bdf :=[]byte {};_cbf :=0;for _edg :=0;_edg < len (_bg );_edg ++{var _bba string ;if _bg [_edg ]=='/'{_bba =string (_bg [_cbf :_edg ]);_cbf =_edg +1;}else if _edg ==len (_bg )-1{_bba =string (_bg [_cbf :_edg +1]);}else {continue ;};switch _bba {case "\u0079\u0079":_bdf =_ggb .AppendFormat (_bdf ,"\u0030\u0036");case "\u0079\u0079\u0079\u0079":_bdf =_ggb .AppendFormat (_bdf ,"\u0032\u0030\u0030\u0036");case "\u006d":_bdf =_ggb .AppendFormat (_bdf ,"\u0031");case "\u006d\u006d":_bdf =_ggb .AppendFormat (_bdf ,"\u0030\u0031");case "\u006d\u006d\u006d":_bdf =_ggb .AppendFormat (_bdf ,"\u004a\u0061\u006e");case "\u006d\u006d\u006d\u006d":_bdf =_ggb .AppendFormat (_bdf ,"\u004aa\u006e\u0075\u0061\u0072\u0079");case "\u006d\u006d\u006dm\u006d":switch _ggb .Month (){case _ce .January ,_ce .July ,_ce .June :_bdf =append (_bdf ,'J');case _ce .February :_bdf =append (_bdf ,'M');case _ce .March ,_ce .May :_bdf =append (_bdf ,'M');case _ce .April ,_ce .August :_bdf =append (_bdf ,'A');case _ce .September :_bdf =append (_bdf ,'S');case _ce .October :_bdf =append (_bdf ,'O');case _ce .November :_bdf =append (_bdf ,'N');case _ce .December :_bdf =append (_bdf ,'D');};case "\u0064":_bdf =_ggb .AppendFormat (_bdf ,"\u0032");case "\u0064\u0064":_bdf =_ggb .AppendFormat (_bdf ,"\u0030\u0032");case "\u0064\u0064\u0064":_bdf =_ggb .AppendFormat (_bdf ,"\u004d\u006f\u006e");case "\u0064\u0064\u0064\u0064":_bdf =_ggb .AppendFormat (_bdf ,"\u004d\u006f\u006e\u0064\u0061\u0079");default:_aab .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0064\u0061\u0074\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073",_bba );};if _bg [_edg ]=='/'{_bdf =append (_bdf ,'/');};};return _bdf ;};const _caa int =34;const _cacg int =-1;func IsNumber (data string )(_ffa bool ){_gde ,_bgbd ,_bbce :=0,0,len (data );_dfa :=len (data );_eea ,_eda ,_gca :=0,0,0;_ =_eda ;_ =_gca ;_ =_eea ;{_gde =_gedd ;_eea =0;_eda =0;_gca =0;};{if _bgbd ==_bbce {goto _aaa ;};switch _gde {case 0:goto _bfgb ;case 1:goto _ea ;case 2:goto _gdb ;case 3:goto _dedc ;case 4:goto _fgc ;case 5:goto _dac ;case 6:goto _fcb ;case 7:goto _dcaf ;};goto _ebed ;_ggd :_eda =_bgbd ;_bgbd --;{_ffa =false ;};goto _dfaf ;_bbeg :_eda =_bgbd ;_bgbd --;{_ffa =_eda ==len (data );};goto _dfaf ;_deg :_eda =_bgbd ;_bgbd --;{_ffa =_eda ==len (data );};goto _dfaf ;_ffbf :switch _gca {case 2:{_bgbd =(_eda )-1;_ffa =_eda ==len (data );};case 3:{_bgbd =(_eda )-1;_ffa =false ;};};goto _dfaf ;_dfaf :_eea =0;if _bgbd ++;_bgbd ==_bbce {goto _fda ;};_bfgb :_eea =_bgbd ;switch data [_bgbd ]{case 43:goto _dedb ;case 45:goto _dedb ;};if 48<=data [_bgbd ]&&data [_bgbd ]<=57{goto _fea ;};goto _cc ;_cc :if _bgbd ++;_bgbd ==_bbce {goto _fee ;};_ea :goto _cc ;_dedb :if _bgbd ++;_bgbd ==_bbce {goto _afd ;};_gdb :if 48<=data [_bgbd ]&&data [_bgbd ]<=57{goto _fea ;};goto _cc ;_fea :if _bgbd ++;_bgbd ==_bbce {goto _beg ;};_dedc :if data [_bgbd ]==46{goto _fbe ;};if 48<=data [_bgbd ]&&data [_bgbd ]<=57{goto _fea ;};goto _cc ;_fbe :if _bgbd ++;_bgbd ==_bbce {goto _bbca ;};_fgc :if 48<=data [_bgbd ]&&data [_bgbd ]<=57{goto _eab ;};goto _cc ;_eab :if _bgbd ++;_bgbd ==_bbce {goto _eec ;};_dac :if data [_bgbd ]==69{goto _acb ;};if 48<=data [_bgbd ]&&data [_bgbd ]<=57{goto _eab ;};goto _cc ;_acb :if _bgbd ++;_bgbd ==_bbce {goto _gffd ;};_fcb :switch data [_bgbd ]{case 43:goto _gee ;case 45:goto _gee ;};goto _cc ;_gee :_eda =_bgbd +1;_gca =3;goto _bbg ;_fcga :_eda =_bgbd +1;_gca =2;goto _bbg ;_bbg :if _bgbd ++;_bgbd ==_bbce {goto _aag ;};_dcaf :if 48<=data [_bgbd ]&&data [_bgbd ]<=57{goto _fcga ;};goto _cc ;_ebed :_fda :_gde =0;goto _aaa ;_fee :_gde =1;goto _aaa ;_afd :_gde =2;goto _aaa ;_beg :_gde =3;goto _aaa ;_bbca :_gde =4;goto _aaa ;_eec :_gde =5;goto _aaa ;_gffd :_gde =6;goto _aaa ;_aag :_gde =7;goto _aaa ;_aaa :{};if _bgbd ==_dfa {switch _gde {case 1:goto _ggd ;case 2:goto _ggd ;case 3:goto _bbeg ;case 4:goto _ggd ;case 5:goto _deg ;case 6:goto _ggd ;case 7:goto _ffbf ;};};};if _gde ==_cacg {return false ;};return ;};const _fdd int =0;func _dca (_gge ,_adb float64 ,_cef Format )[]byte {if len (_cef .Fractional )==0{return nil ;};_ff :=_b .AppendFloat (nil ,_gge ,'f',-1,64);if len (_ff )> 2{_ff =_ff [2:];}else {_ff =nil ;};_dec :=make ([]byte ,0,len (_ff ));_dec =append (_dec ,'.');_fac :=0;_efd :for _ab :=0;_ab < len (_cef .Fractional );_ab ++{_efg :=_ab ;_edf :=_cef .Fractional [_ab ];switch _edf .Type {case FmtTypeDigit :if _efg < len (_ff ){_dec =append (_dec ,_ff [_efg ]);_fac ++;}else {_dec =append (_dec ,'0');};case FmtTypeDigitOpt :if _efg >=0{_dec =append (_dec ,_ff [_efg ]);_fac ++;}else {break _efd ;};case FmtTypeLiteral :_dec =append (_dec ,_edf .Literal );default:_aab .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0074\u0079\u0070\u0065\u0020\u0069\u006e\u0020f\u0072\u0061\u0063\u0074\u0069\u006f\u006ea\u006c\u0020\u0025\u0076",_edf );};};return _dec ;};type Lexer struct{_daa Format ;_efc []Format ;};func _gd (_cab []byte )[]byte {for _bff :=0;_bff < len (_cab )/2;_bff ++{_gf :=len (_cab )-1-_bff ;_cab [_bff ],_cab [_gf ]=_cab [_gf ],_cab [_bff ];};return _cab ;};func _cdda (_cdea _ce .Time ,_bgb float64 ,_bad string )[]byte {_dff :=[]byte {};_aca :=0;for _acd :=0;_acd < len (_bad );_acd ++{var _ged string ;if _bad [_acd ]==':'{_ged =string (_bad [_aca :_acd ]);_aca =_acd +1;}else if _acd ==len (_bad )-1{_ged =string (_bad [_aca :_acd +1]);}else {continue ;};switch _ged {case "\u0064":_dff =_cdea .AppendFormat (_dff ,"\u0032");case "\u0068":_dff =_cdea .AppendFormat (_dff ,"\u0033");case "\u0068\u0068":_dff =_cdea .AppendFormat (_dff ,"\u0031\u0035");case "\u006d":_dff =_cdea .AppendFormat (_dff ,"\u0034");case "\u006d\u006d":_dff =_cdea .AppendFormat (_dff ,"\u0030\u0034");case "\u0073":_dff =_cdea .Round (_ce .Second ).AppendFormat (_dff ,"\u0035");case "\u0073\u002e\u0030":_dff =_cdea .Round (_ce .Second /10).AppendFormat (_dff ,"\u0035\u002e\u0030");case "\u0073\u002e\u0030\u0030":_dff =_cdea .Round (_ce .Second /100).AppendFormat (_dff ,"\u0035\u002e\u0030\u0030");case "\u0073\u002e\u00300\u0030":_dff =_cdea .Round (_ce .Second /1000).AppendFormat (_dff ,"\u0035\u002e\u00300\u0030");case "\u0073\u0073":_dff =_cdea .Round (_ce .Second ).AppendFormat (_dff ,"\u0030\u0035");case "\u0073\u0073\u002e\u0030":_dff =_cdea .Round (_ce .Second /10).AppendFormat (_dff ,"\u0030\u0035\u002e\u0030");case "\u0073\u0073\u002e0\u0030":_dff =_cdea .Round (_ce .Second /100).AppendFormat (_dff ,"\u0030\u0035\u002e0\u0030");case "\u0073\u0073\u002e\u0030\u0030\u0030":_dff =_cdea .Round (_ce .Second /1000).AppendFormat (_dff ,"\u0030\u0035\u002e\u0030\u0030\u0030");case "\u0041\u004d\u002fP\u004d":_dff =_cdea .AppendFormat (_dff ,"\u0050\u004d");case "\u005b\u0068\u005d":_dff =_b .AppendInt (_dff ,int64 (_bgb *24),10);case "\u005b\u006d\u005d":_dff =_b .AppendInt (_dff ,int64 (_bgb *24*60),10);case "\u005b\u0073\u005d":_dff =_b .AppendInt (_dff ,int64 (_bgb *24*60*60),10);case "":default:_aab .Log .Debug ("\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064 \u0074\u0069\u006d\u0065\u0020\u0066\u006f\u0072\u006d\u0061t\u0020\u0025\u0073",_ged );};if _bad [_acd ]==':'{_dff =append (_dff ,':');};};return _dff ;};

// NumberGeneric formats the number with the generic format which attemps to
// mimic Excel's general formatting.
func NumberGeneric (v float64 )string {if _e .Abs (v )>=_ca ||_e .Abs (v )<=_fc &&v !=0{return _bdb (v );};_be :=make ([]byte ,0,15);_be =_b .AppendFloat (_be ,v ,'f',-1,64);if len (_be )> 11{_cbc :=_be [11]-'0';if _cbc >=5&&_cbc <=9{_be [10]++;_be =_be [0:11];_be =_gb (_be );};_be =_be [0:11];}else if len (_be )==11{if _be [len (_be )-1]=='9'{_be [len (_be )-1]++;_be =_gb (_be );};};_be =_fec (_be );return string (_be );};const _bgd int =0;