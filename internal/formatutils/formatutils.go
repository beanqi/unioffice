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

package formatutils ;import (_e "fmt";_gf "github.com/unidoc/unioffice/schema/soo/wml";_g "strconv";_f "strings";);func FormatNumberingText (currentNumber int64 ,ilvl int64 ,lvlText string ,numFmt *_gf .CT_NumFmt ,levelNumbers map[int64 ]int64 )string {_ac :=_df (lvlText );_gc :=_db (currentNumber ,numFmt );_b :=int64 (0);for _fg ,_d :=range _ac {_gca :=_e .Sprintf ("\u0025\u0025\u0025\u0064",_fg +1);if len (_ac )==1{_gca =_e .Sprintf ("\u0025\u0025\u0025\u0064",ilvl +1);_ac [_fg ]=_f .Replace (_d ,_gca ,_gc ,1);break ;};if ilvl > 0&&ilvl > _b &&_fg < (len (_ac )-1){_ee :=_db (levelNumbers [_b ],numFmt );_ac [_fg ]=_f .Replace (_d ,_gca ,_ee ,1);_b ++;}else {_ac [_fg ]=_f .Replace (_d ,_gca ,_gc ,1);};};return _f .Join (_ac ,"");};func _df (_de string )(_fd []string ){for _dd :=0;_dd < len (_de )-2;_dd ++{if string (_de [_dd ])=="\u0025"{if !_f .Contains (string (_de [_dd +2:]),"\u0025"){if _dd ==0{_fd =append (_fd ,_e .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_de [_dd ]),string (_de [_dd +1]),string (_de [_dd +2:])));}else {_fd =append (_fd ,_e .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073\u0025\u0073",string (_de [_dd -1]),string (_de [_dd ]),string (_de [_dd +1]),string (_de [_dd +2:])));};}else {_fd =append (_fd ,_e .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_de [_dd ]),string (_de [_dd +1]),string (_de [_dd +2])));};};};return ;};func _db (_gff int64 ,_cbd *_gf .CT_NumFmt )(_ad string ){if _cbd ==nil {return ;};_ec :=_cbd .ValAttr ;switch _ec {case _gf .ST_NumberFormatNone :_ad ="";case _gf .ST_NumberFormatDecimal :_ad =_g .Itoa (int (_gff ));case _gf .ST_NumberFormatDecimalZero :_ad =_g .Itoa (int (_gff ));if _gff < 10{_ad ="\u0030"+_ad ;};case _gf .ST_NumberFormatUpperRoman :var (_fe =_gff %10;_aa =(_gff %100)/10;_acc =(_gff %1000)/100;_eb =_gff /1000;);_ad =_c [_eb ]+_eg [_acc ]+_ea [_aa ]+_gcc [_fe ];case _gf .ST_NumberFormatLowerRoman :var (_ead =_gff %10;_eec =(_gff %100)/10;_deb =(_gff %1000)/100;_af =_gff /1000;);_ad =_c [_af ]+_eg [_deb ]+_ea [_eec ]+_gcc [_ead ];_ad =_f .ToLower (_ad );case _gf .ST_NumberFormatUpperLetter :_fdc :=_gff %780;if _fdc ==0{_fdc =780;};_fa :=(_fdc -1)/26;_aff :=(_fdc -1)%26;_ba :=_gcb [_fa +_aff ];_ad =string (_ba );case _gf .ST_NumberFormatLowerLetter :_afc :=_gff %780;if _afc ==0{_afc =780;};_ceg :=(_afc -1)/26;_ab :=(_afc -1)%26;_dfb :=_gcb [_ceg +_ab ];_ad =_f .ToLower (string (_dfb ));default:_ad ="";};return ;};var (_gcc =[]string {"","\u0049","\u0049\u0049","\u0049\u0049\u0049","\u0049\u0056","\u0056","\u0056\u0049","\u0056\u0049\u0049","\u0056\u0049\u0049\u0049","\u0049\u0058"};_ea =[]string {"","\u0058","\u0058\u0058","\u0058\u0058\u0058","\u0058\u004c","\u004c","\u004c\u0058","\u004c\u0058\u0058","\u004c\u0058\u0058\u0058","\u0058\u0043"};_eg =[]string {"","\u0043","\u0043\u0043","\u0043\u0043\u0043","\u0043\u0044","\u0044","\u0044\u0043","\u0044\u0043\u0043","\u0044\u0043\u0043\u0043","\u0043\u004d","\u004d"};_c =[]string {"","\u004d","\u004d\u004d","\u004d\u004d\u004d","\u004d\u004d\u004d\u004d","\u004d\u004d\u004dM\u004d","\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d\u004dM","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d"};_bd =[]string {"\u006f\u006e\u0065","\u0074\u0077\u006f","\u0074\u0068\u0072e\u0065","\u0066\u006f\u0075\u0072","\u0066\u0069\u0076\u0065","\u0073\u0069\u0078","\u0073\u0065\u0076e\u006e","\u0065\u0069\u0067h\u0074","\u006e\u0069\u006e\u0065","\u0074\u0065\u006e","\u0065\u006c\u0065\u0076\u0065\u006e","\u0074\u0077\u0065\u006c\u0076\u0065","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e","\u0066i\u0066\u0074\u0065\u0065\u006e","\u0073i\u0078\u0074\u0065\u0065\u006e","\u0073e\u0076\u0065\u006e\u0074\u0065\u0065n","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e"};_ce =[]string {"\u0074\u0065\u006e","\u0074\u0077\u0065\u006e\u0074\u0079","\u0074\u0068\u0069\u0072\u0074\u0079","\u0066\u006f\u0072t\u0079","\u0066\u0069\u0066t\u0079","\u0073\u0069\u0078t\u0079","\u0073e\u0076\u0065\u006e\u0074\u0079","\u0065\u0069\u0067\u0068\u0074\u0079","\u006e\u0069\u006e\u0065\u0074\u0079"};_cf =[]string {"\u0066\u0069\u0072s\u0074","\u0073\u0065\u0063\u006f\u006e\u0064","\u0074\u0068\u0069r\u0064","\u0066\u006f\u0075\u0072\u0074\u0068","\u0066\u0069\u0066t\u0068","\u0073\u0069\u0078t\u0068","\u0073e\u0076\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0068","\u006e\u0069\u006et\u0068","\u0074\u0065\u006et\u0068","\u0065\u006c\u0065\u0076\u0065\u006e\u0074\u0068","\u0074w\u0065\u006c\u0066\u0074\u0068","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066i\u0066\u0074\u0065\u0065\u006e\u0074h","\u0073i\u0078\u0074\u0065\u0065\u006e\u0074h","s\u0065\u0076\u0065\u006e\u0074\u0065\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e\u0074\u0068","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e\u0074\u0068"};_cb =[]string {"\u0074\u0065\u006et\u0068","\u0074w\u0065\u006e\u0074\u0069\u0065\u0074h","\u0074h\u0069\u0072\u0074\u0069\u0065\u0074h","\u0066\u006f\u0072\u0074\u0069\u0065\u0074\u0068","\u0066\u0069\u0066\u0074\u0069\u0065\u0074\u0068","\u0073\u0069\u0078\u0074\u0069\u0065\u0074\u0068","\u0073\u0065\u0076\u0065\u006e\u0074\u0069\u0065\u0074\u0068","\u0065i\u0067\u0068\u0074\u0069\u0065\u0074h","\u006ei\u006e\u0065\u0074\u0069\u0065\u0074h"};_gcb ="\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004a\u004bL\u004d\u004e\u004f\u0050\u0051\u0052\u0053\u0054\u0055\u0056W\u0058\u0059\u005a";);