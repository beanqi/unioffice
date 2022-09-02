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

package formatutils ;import (_ag "fmt";_gd "github.com/unidoc/unioffice/schema/soo/wml";_e "strconv";_g "strings";);var (_cf =[]string {"","\u0049","\u0049\u0049","\u0049\u0049\u0049","\u0049\u0056","\u0056","\u0056\u0049","\u0056\u0049\u0049","\u0056\u0049\u0049\u0049","\u0049\u0058"};_ea =[]string {"","\u0058","\u0058\u0058","\u0058\u0058\u0058","\u0058\u004c","\u004c","\u004c\u0058","\u004c\u0058\u0058","\u004c\u0058\u0058\u0058","\u0058\u0043"};_ce =[]string {"","\u0043","\u0043\u0043","\u0043\u0043\u0043","\u0043\u0044","\u0044","\u0044\u0043","\u0044\u0043\u0043","\u0044\u0043\u0043\u0043","\u0043\u004d","\u004d"};_bc =[]string {"","\u004d","\u004d\u004d","\u004d\u004d\u004d","\u004d\u004d\u004d\u004d","\u004d\u004d\u004dM\u004d","\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d\u004dM","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d"};_f =[]string {"\u006f\u006e\u0065","\u0074\u0077\u006f","\u0074\u0068\u0072e\u0065","\u0066\u006f\u0075\u0072","\u0066\u0069\u0076\u0065","\u0073\u0069\u0078","\u0073\u0065\u0076e\u006e","\u0065\u0069\u0067h\u0074","\u006e\u0069\u006e\u0065","\u0074\u0065\u006e","\u0065\u006c\u0065\u0076\u0065\u006e","\u0074\u0077\u0065\u006c\u0076\u0065","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e","\u0066i\u0066\u0074\u0065\u0065\u006e","\u0073i\u0078\u0074\u0065\u0065\u006e","\u0073e\u0076\u0065\u006e\u0074\u0065\u0065n","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e"};_aee =[]string {"\u0074\u0065\u006e","\u0074\u0077\u0065\u006e\u0074\u0079","\u0074\u0068\u0069\u0072\u0074\u0079","\u0066\u006f\u0072t\u0079","\u0066\u0069\u0066t\u0079","\u0073\u0069\u0078t\u0079","\u0073e\u0076\u0065\u006e\u0074\u0079","\u0065\u0069\u0067\u0068\u0074\u0079","\u006e\u0069\u006e\u0065\u0074\u0079"};_de =[]string {"\u0066\u0069\u0072s\u0074","\u0073\u0065\u0063\u006f\u006e\u0064","\u0074\u0068\u0069r\u0064","\u0066\u006f\u0075\u0072\u0074\u0068","\u0066\u0069\u0066t\u0068","\u0073\u0069\u0078t\u0068","\u0073e\u0076\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0068","\u006e\u0069\u006et\u0068","\u0074\u0065\u006et\u0068","\u0065\u006c\u0065\u0076\u0065\u006e\u0074\u0068","\u0074w\u0065\u006c\u0066\u0074\u0068","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066i\u0066\u0074\u0065\u0065\u006e\u0074h","\u0073i\u0078\u0074\u0065\u0065\u006e\u0074h","s\u0065\u0076\u0065\u006e\u0074\u0065\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e\u0074\u0068","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e\u0074\u0068"};_dec =[]string {"\u0074\u0065\u006et\u0068","\u0074w\u0065\u006e\u0074\u0069\u0065\u0074h","\u0074h\u0069\u0072\u0074\u0069\u0065\u0074h","\u0066\u006f\u0072\u0074\u0069\u0065\u0074\u0068","\u0066\u0069\u0066\u0074\u0069\u0065\u0074\u0068","\u0073\u0069\u0078\u0074\u0069\u0065\u0074\u0068","\u0073\u0065\u0076\u0065\u006e\u0074\u0069\u0065\u0074\u0068","\u0065i\u0067\u0068\u0074\u0069\u0065\u0074h","\u006ei\u006e\u0065\u0074\u0069\u0065\u0074h"};_abe ="\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004a\u004bL\u004d\u004e\u004f\u0050\u0051\u0052\u0053\u0054\u0055\u0056W\u0058\u0059\u005a";);func FormatNumberingText (currentNumber int64 ,ilvl int64 ,lvlText string ,numFmt *_gd .CT_NumFmt ,levelNumbers map[int64 ]int64 )string {_c :=_bg (lvlText );_agc :=_agb (currentNumber ,numFmt );_ab :=int64 (0);for _b ,_eb :=range _c {_gc :=_ag .Sprintf ("\u0025\u0025\u0025\u0064",_b +1);if len (_c )==1{_gc =_ag .Sprintf ("\u0025\u0025\u0025\u0064",ilvl +1);_c [_b ]=_g .Replace (_eb ,_gc ,_agc ,1);break ;};if ilvl > 0&&ilvl > _ab &&_b < (len (_c )-1){_ae :=_agb (levelNumbers [_ab ],numFmt );_c [_b ]=_g .Replace (_eb ,_gc ,_ae ,1);_ab ++;}else {_c [_b ]=_g .Replace (_eb ,_gc ,_agc ,1);};};return _g .Join (_c ,"");};func _agb (_cfc int64 ,_dc *_gd .CT_NumFmt )(_ec string ){if _dc ==nil {return ;};_ef :=_dc .ValAttr ;switch _ef {case _gd .ST_NumberFormatNone :_ec ="";case _gd .ST_NumberFormatDecimal :_ec =_e .Itoa (int (_cfc ));case _gd .ST_NumberFormatDecimalZero :_ec =_e .Itoa (int (_cfc ));if _cfc < 10{_ec ="\u0030"+_ec ;};case _gd .ST_NumberFormatUpperRoman :var (_af =_cfc %10;_gg =(_cfc %100)/10;_fe =(_cfc %1000)/100;_ecc =_cfc /1000;);_ec =_bc [_ecc ]+_ce [_fe ]+_ea [_gg ]+_cf [_af ];case _gd .ST_NumberFormatLowerRoman :var (_cc =_cfc %10;_dcf =(_cfc %100)/10;_afd =(_cfc %1000)/100;_gb =_cfc /1000;);_ec =_bc [_gb ]+_ce [_afd ]+_ea [_dcf ]+_cf [_cc ];_ec =_g .ToLower (_ec );case _gd .ST_NumberFormatUpperLetter :_ebg :=_cfc %780;if _ebg ==0{_ebg =780;};_ca :=(_ebg -1)/26;_fec :=(_ebg -1)%26;_ggg :=_abe [_ca +_fec ];_ec =string (_ggg );case _gd .ST_NumberFormatLowerLetter :_agf :=_cfc %780;if _agf ==0{_agf =780;};_bbc :=(_agf -1)/26;_fef :=(_agf -1)%26;_abd :=_abe [_bbc +_fef ];_ec =_g .ToLower (string (_abd ));default:_ec ="";};return ;};func _bg (_eba string )(_d []string ){for _bb :=0;_bb < len (_eba )-2;_bb ++{if string (_eba [_bb ])=="\u0025"{if !_g .Contains (string (_eba [_bb +2:]),"\u0025"){if _bb ==0{_d =append (_d ,_ag .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_eba [_bb ]),string (_eba [_bb +1]),string (_eba [_bb +2:])));}else {_d =append (_d ,_ag .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073\u0025\u0073",string (_eba [_bb -1]),string (_eba [_bb ]),string (_eba [_bb +1]),string (_eba [_bb +2:])));};}else {_d =append (_d ,_ag .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_eba [_bb ]),string (_eba [_bb +1]),string (_eba [_bb +2])));};};};return ;};