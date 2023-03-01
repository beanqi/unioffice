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

package formatutils ;import (_g "fmt";_be "github.com/unidoc/unioffice/schema/soo/wml";_bf "strconv";_gf "strings";);func FormatNumberingText (currentNumber int64 ,ilvl int64 ,lvlText string ,numFmt *_be .CT_NumFmt ,levelNumbers map[int64 ]int64 )string {_f :=_ee (lvlText );_c :=_af (currentNumber ,numFmt );_ff :=int64 (0);for _e ,_gff :=range _f {_a :=_g .Sprintf ("\u0025\u0025\u0025\u0064",_e +1);if len (_f )==1{_a =_g .Sprintf ("\u0025\u0025\u0025\u0064",ilvl +1);_f [_e ]=_gf .Replace (_gff ,_a ,_c ,1);break ;};if ilvl > 0&&ilvl > _ff &&_e < (len (_f )-1){_ag :=_af (levelNumbers [_ff ],numFmt );_f [_e ]=_gf .Replace (_gff ,_a ,_ag ,1);_ff ++;}else {_f [_e ]=_gf .Replace (_gff ,_a ,_c ,1);};};return _gf .Join (_f ,"");};var (_bed =[]string {"","\u0049","\u0049\u0049","\u0049\u0049\u0049","\u0049\u0056","\u0056","\u0056\u0049","\u0056\u0049\u0049","\u0056\u0049\u0049\u0049","\u0049\u0058"};_aa =[]string {"","\u0058","\u0058\u0058","\u0058\u0058\u0058","\u0058\u004c","\u004c","\u004c\u0058","\u004c\u0058\u0058","\u004c\u0058\u0058\u0058","\u0058\u0043"};_cg =[]string {"","\u0043","\u0043\u0043","\u0043\u0043\u0043","\u0043\u0044","\u0044","\u0044\u0043","\u0044\u0043\u0043","\u0044\u0043\u0043\u0043","\u0043\u004d","\u004d"};_bef =[]string {"","\u004d","\u004d\u004d","\u004d\u004d\u004d","\u004d\u004d\u004d\u004d","\u004d\u004d\u004dM\u004d","\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d\u004dM","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d"};_fb =[]string {"\u006f\u006e\u0065","\u0074\u0077\u006f","\u0074\u0068\u0072e\u0065","\u0066\u006f\u0075\u0072","\u0066\u0069\u0076\u0065","\u0073\u0069\u0078","\u0073\u0065\u0076e\u006e","\u0065\u0069\u0067h\u0074","\u006e\u0069\u006e\u0065","\u0074\u0065\u006e","\u0065\u006c\u0065\u0076\u0065\u006e","\u0074\u0077\u0065\u006c\u0076\u0065","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e","\u0066i\u0066\u0074\u0065\u0065\u006e","\u0073i\u0078\u0074\u0065\u0065\u006e","\u0073e\u0076\u0065\u006e\u0074\u0065\u0065n","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e"};_gfc =[]string {"\u0074\u0065\u006e","\u0074\u0077\u0065\u006e\u0074\u0079","\u0074\u0068\u0069\u0072\u0074\u0079","\u0066\u006f\u0072t\u0079","\u0066\u0069\u0066t\u0079","\u0073\u0069\u0078t\u0079","\u0073e\u0076\u0065\u006e\u0074\u0079","\u0065\u0069\u0067\u0068\u0074\u0079","\u006e\u0069\u006e\u0065\u0074\u0079"};_ae =[]string {"\u0066\u0069\u0072s\u0074","\u0073\u0065\u0063\u006f\u006e\u0064","\u0074\u0068\u0069r\u0064","\u0066\u006f\u0075\u0072\u0074\u0068","\u0066\u0069\u0066t\u0068","\u0073\u0069\u0078t\u0068","\u0073e\u0076\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0068","\u006e\u0069\u006et\u0068","\u0074\u0065\u006et\u0068","\u0065\u006c\u0065\u0076\u0065\u006e\u0074\u0068","\u0074w\u0065\u006c\u0066\u0074\u0068","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066i\u0066\u0074\u0065\u0065\u006e\u0074h","\u0073i\u0078\u0074\u0065\u0065\u006e\u0074h","s\u0065\u0076\u0065\u006e\u0074\u0065\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e\u0074\u0068","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e\u0074\u0068"};_fg =[]string {"\u0074\u0065\u006et\u0068","\u0074w\u0065\u006e\u0074\u0069\u0065\u0074h","\u0074h\u0069\u0072\u0074\u0069\u0065\u0074h","\u0066\u006f\u0072\u0074\u0069\u0065\u0074\u0068","\u0066\u0069\u0066\u0074\u0069\u0065\u0074\u0068","\u0073\u0069\u0078\u0074\u0069\u0065\u0074\u0068","\u0073\u0065\u0076\u0065\u006e\u0074\u0069\u0065\u0074\u0068","\u0065i\u0067\u0068\u0074\u0069\u0065\u0074h","\u006ei\u006e\u0065\u0074\u0069\u0065\u0074h"};_cgc ="\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004a\u004bL\u004d\u004e\u004f\u0050\u0051\u0052\u0053\u0054\u0055\u0056W\u0058\u0059\u005a";);func _ee (_ea string )(_fa []string ){for _fc :=0;_fc < len (_ea )-2;_fc ++{if string (_ea [_fc ])=="\u0025"{if !_gf .Contains (string (_ea [_fc +2:]),"\u0025"){if _fc ==0{_fa =append (_fa ,_g .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_ea [_fc ]),string (_ea [_fc +1]),string (_ea [_fc +2:])));}else {_fa =append (_fa ,_g .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073\u0025\u0073",string (_ea [_fc -1]),string (_ea [_fc ]),string (_ea [_fc +1]),string (_ea [_fc +2:])));};}else {_fa =append (_fa ,_g .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_ea [_fc ]),string (_ea [_fc +1]),string (_ea [_fc +2])));};};};return ;};func _af (_fgg int64 ,_bc *_be .CT_NumFmt )(_d string ){if _bc ==nil {return ;};_fggb :=_bc .ValAttr ;switch _fggb {case _be .ST_NumberFormatNone :_d ="";case _be .ST_NumberFormatDecimal :_d =_bf .Itoa (int (_fgg ));case _be .ST_NumberFormatDecimalZero :_d =_bf .Itoa (int (_fgg ));if _fgg < 10{_d ="\u0030"+_d ;};case _be .ST_NumberFormatUpperRoman :var (_bd =_fgg %10;_db =(_fgg %100)/10;_afa =(_fgg %1000)/100;_dc =_fgg /1000;);_d =_bef [_dc ]+_cg [_afa ]+_aa [_db ]+_bed [_bd ];case _be .ST_NumberFormatLowerRoman :var (_bg =_fgg %10;_ga =(_fgg %100)/10;_ca =(_fgg %1000)/100;_dcb =_fgg /1000;);_d =_bef [_dcb ]+_cg [_ca ]+_aa [_ga ]+_bed [_bg ];_d =_gf .ToLower (_d );case _be .ST_NumberFormatUpperLetter :_bb :=_fgg %780;if _bb ==0{_bb =780;};_ed :=(_bb -1)/26;_ad :=(_bb -1)%26;_fcg :=_cgc [_ed +_ad ];_d =string (_fcg );case _be .ST_NumberFormatLowerLetter :_bba :=_fgg %780;if _bba ==0{_bba =780;};_ce :=(_bba -1)/26;_eg :=(_bba -1)%26;_gfe :=_cgc [_ce +_eg ];_d =_gf .ToLower (string (_gfe ));default:_d ="";};return ;};