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

package formatutils ;import (_g "fmt";_f "github.com/unidoc/unioffice/schema/soo/wml";_d "strconv";_e "strings";);var (_dgf =[]string {"","\u0049","\u0049\u0049","\u0049\u0049\u0049","\u0049\u0056","\u0056","\u0056\u0049","\u0056\u0049\u0049","\u0056\u0049\u0049\u0049","\u0049\u0058"};_eg =[]string {"","\u0058","\u0058\u0058","\u0058\u0058\u0058","\u0058\u004c","\u004c","\u004c\u0058","\u004c\u0058\u0058","\u004c\u0058\u0058\u0058","\u0058\u0043"};_agd =[]string {"","\u0043","\u0043\u0043","\u0043\u0043\u0043","\u0043\u0044","\u0044","\u0044\u0043","\u0044\u0043\u0043","\u0044\u0043\u0043\u0043","\u0043\u004d","\u004d"};_ae =[]string {"","\u004d","\u004d\u004d","\u004d\u004d\u004d","\u004d\u004d\u004d\u004d","\u004d\u004d\u004dM\u004d","\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d\u004dM","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d"};_egd =[]string {"\u006f\u006e\u0065","\u0074\u0077\u006f","\u0074\u0068\u0072e\u0065","\u0066\u006f\u0075\u0072","\u0066\u0069\u0076\u0065","\u0073\u0069\u0078","\u0073\u0065\u0076e\u006e","\u0065\u0069\u0067h\u0074","\u006e\u0069\u006e\u0065","\u0074\u0065\u006e","\u0065\u006c\u0065\u0076\u0065\u006e","\u0074\u0077\u0065\u006c\u0076\u0065","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e","\u0066i\u0066\u0074\u0065\u0065\u006e","\u0073i\u0078\u0074\u0065\u0065\u006e","\u0073e\u0076\u0065\u006e\u0074\u0065\u0065n","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e"};_ab =[]string {"\u0074\u0065\u006e","\u0074\u0077\u0065\u006e\u0074\u0079","\u0074\u0068\u0069\u0072\u0074\u0079","\u0066\u006f\u0072t\u0079","\u0066\u0069\u0066t\u0079","\u0073\u0069\u0078t\u0079","\u0073e\u0076\u0065\u006e\u0074\u0079","\u0065\u0069\u0067\u0068\u0074\u0079","\u006e\u0069\u006e\u0065\u0074\u0079"};_eb =[]string {"\u0066\u0069\u0072s\u0074","\u0073\u0065\u0063\u006f\u006e\u0064","\u0074\u0068\u0069r\u0064","\u0066\u006f\u0075\u0072\u0074\u0068","\u0066\u0069\u0066t\u0068","\u0073\u0069\u0078t\u0068","\u0073e\u0076\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0068","\u006e\u0069\u006et\u0068","\u0074\u0065\u006et\u0068","\u0065\u006c\u0065\u0076\u0065\u006e\u0074\u0068","\u0074w\u0065\u006c\u0066\u0074\u0068","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066i\u0066\u0074\u0065\u0065\u006e\u0074h","\u0073i\u0078\u0074\u0065\u0065\u006e\u0074h","s\u0065\u0076\u0065\u006e\u0074\u0065\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e\u0074\u0068","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e\u0074\u0068"};_agf =[]string {"\u0074\u0065\u006et\u0068","\u0074w\u0065\u006e\u0074\u0069\u0065\u0074h","\u0074h\u0069\u0072\u0074\u0069\u0065\u0074h","\u0066\u006f\u0072\u0074\u0069\u0065\u0074\u0068","\u0066\u0069\u0066\u0074\u0069\u0065\u0074\u0068","\u0073\u0069\u0078\u0074\u0069\u0065\u0074\u0068","\u0073\u0065\u0076\u0065\u006e\u0074\u0069\u0065\u0074\u0068","\u0065i\u0067\u0068\u0074\u0069\u0065\u0074h","\u006ei\u006e\u0065\u0074\u0069\u0065\u0074h"};_dgg ="\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004a\u004bL\u004d\u004e\u004f\u0050\u0051\u0052\u0053\u0054\u0055\u0056W\u0058\u0059\u005a";);func FormatNumberingText (currentNumber int64 ,ilvl int64 ,lvlText string ,numFmt *_f .CT_NumFmt ,levelNumbers map[int64 ]int64 )string {_de :=_df (lvlText );_ge :=_af (currentNumber ,numFmt );_b :=int64 (0);for _dg ,_be :=range _de {_fg :=_g .Sprintf ("\u0025\u0025\u0025\u0064",_dg +1);if len (_de )==1{_fg =_g .Sprintf ("\u0025\u0025\u0025\u0064",ilvl +1);_de [_dg ]=_e .Replace (_be ,_fg ,_ge ,1);break ;};if ilvl > 0&&ilvl > _b &&_dg < (len (_de )-1){_da :=_af (levelNumbers [_b ],numFmt );_de [_dg ]=_e .Replace (_be ,_fg ,_da ,1);_b ++;}else {_de [_dg ]=_e .Replace (_be ,_fg ,_ge ,1);};};return _e .Join (_de ,"");};func _df (_ac string )(_ag []string ){for _def :=0;_def < len (_ac )-2;_def ++{if string (_ac [_def ])=="\u0025"{if !_e .Contains (string (_ac [_def +2:]),"\u0025"){if _def ==0{_ag =append (_ag ,_g .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_ac [_def ]),string (_ac [_def +1]),string (_ac [_def +2:])));}else {_ag =append (_ag ,_g .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073\u0025\u0073",string (_ac [_def -1]),string (_ac [_def ]),string (_ac [_def +1]),string (_ac [_def +2:])));};}else {_ag =append (_ag ,_g .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_ac [_def ]),string (_ac [_def +1]),string (_ac [_def +2])));};};};return ;};func _af (_aga int64 ,_gg *_f .CT_NumFmt )(_ed string ){if _gg ==nil {return ;};_bb :=_gg .ValAttr ;switch _bb {case _f .ST_NumberFormatNone :_ed ="";case _f .ST_NumberFormatDecimal :_ed =_d .Itoa (int (_aga ));case _f .ST_NumberFormatDecimalZero :_ed =_d .Itoa (int (_aga ));if _aga < 10{_ed ="\u0030"+_ed ;};case _f .ST_NumberFormatUpperRoman :var (_edc =_aga %10;_fe =(_aga %100)/10;_bd =(_aga %1000)/100;_ea =_aga /1000;);_ed =_ae [_ea ]+_agd [_bd ]+_eg [_fe ]+_dgf [_edc ];case _f .ST_NumberFormatLowerRoman :var (_ec =_aga %10;_ef =(_aga %100)/10;_c =(_aga %1000)/100;_ff =_aga /1000;);_ed =_ae [_ff ]+_agd [_c ]+_eg [_ef ]+_dgf [_ec ];_ed =_e .ToLower (_ed );case _f .ST_NumberFormatUpperLetter :_agac :=_aga %780;if _agac ==0{_agac =780;};_agda :=(_agac -1)/26;_ded :=(_agac -1)%26;_dab :=_dgg [_agda +_ded ];_ed =string (_dab );case _f .ST_NumberFormatLowerLetter :_age :=_aga %780;if _age ==0{_age =780;};_ba :=(_age -1)/26;_fea :=(_age -1)%26;_dede :=_dgg [_ba +_fea ];_ed =_e .ToLower (string (_dede ));default:_ed ="";};return ;};