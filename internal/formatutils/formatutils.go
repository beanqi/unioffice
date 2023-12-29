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

package formatutils ;import (_e "fmt";_a "github.com/unidoc/unioffice/schema/soo/wml";_ef "strconv";_c "strings";);func _de (_cfb int64 ,_ea *_a .CT_NumFmt )(_gf string ){if _ea ==nil {return ;};_ebf :=_ea .ValAttr ;switch _ebf {case _a .ST_NumberFormatNone :_gf ="";
case _a .ST_NumberFormatDecimal :_gf =_ef .Itoa (int (_cfb ));case _a .ST_NumberFormatDecimalZero :_gf =_ef .Itoa (int (_cfb ));if _cfb < 10{_gf ="\u0030"+_gf ;};case _a .ST_NumberFormatUpperRoman :var (_cd =_cfb %10;_bgg =(_cfb %100)/10;_bab =(_cfb %1000)/100;
_bb =_cfb /1000;);_gf =_aa [_bb ]+_bg [_bab ]+_ae [_bgg ]+_eb [_cd ];case _a .ST_NumberFormatLowerRoman :var (_ecd =_cfb %10;_dbf =(_cfb %100)/10;_dec =(_cfb %1000)/100;_ga =_cfb /1000;);_gf =_aa [_ga ]+_bg [_dec ]+_ae [_dbf ]+_eb [_ecd ];_gf =_c .ToLower (_gf );
case _a .ST_NumberFormatUpperLetter :_fd :=_cfb %780;if _fd ==0{_fd =780;};_dd :=(_fd -1)/26;_af :=(_fd -1)%26;_ac :=_db [_dd +_af ];_gf =string (_ac );case _a .ST_NumberFormatLowerLetter :_fa :=_cfb %780;if _fa ==0{_fa =780;};_dfe :=(_fa -1)/26;_fb :=(_fa -1)%26;
_cff :=_db [_dfe +_fb ];_gf =_c .ToLower (string (_cff ));default:_gf ="";};return ;};func FormatNumberingText (currentNumber int64 ,ilvl int64 ,lvlText string ,numFmt *_a .CT_NumFmt ,levelNumbers map[int64 ]int64 )string {_ba :=_ff (lvlText );_bf :=_de (currentNumber ,numFmt );
_d :=int64 (0);for _ad ,_g :=range _ba {_f :=_e .Sprintf ("\u0025\u0025\u0025\u0064",_ad +1);if len (_ba )==1{_f =_e .Sprintf ("\u0025\u0025\u0025\u0064",ilvl +1);_ba [_ad ]=_c .Replace (_g ,_f ,_bf ,1);break ;};if ilvl > 0&&ilvl > _d &&_ad < (len (_ba )-1){_ec :=_de (levelNumbers [_d ],numFmt );
_ba [_ad ]=_c .Replace (_g ,_f ,_ec ,1);_d ++;}else {_ba [_ad ]=_c .Replace (_g ,_f ,_bf ,1);};};return _c .Join (_ba ,"");};var (_eb =[]string {"","\u0049","\u0049\u0049","\u0049\u0049\u0049","\u0049\u0056","\u0056","\u0056\u0049","\u0056\u0049\u0049","\u0056\u0049\u0049\u0049","\u0049\u0058"};
_ae =[]string {"","\u0058","\u0058\u0058","\u0058\u0058\u0058","\u0058\u004c","\u004c","\u004c\u0058","\u004c\u0058\u0058","\u004c\u0058\u0058\u0058","\u0058\u0043"};_bg =[]string {"","\u0043","\u0043\u0043","\u0043\u0043\u0043","\u0043\u0044","\u0044","\u0044\u0043","\u0044\u0043\u0043","\u0044\u0043\u0043\u0043","\u0043\u004d","\u004d"};
_aa =[]string {"","\u004d","\u004d\u004d","\u004d\u004d\u004d","\u004d\u004d\u004d\u004d","\u004d\u004d\u004dM\u004d","\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d","\u004dM\u004d\u004d\u004d\u004d\u004d\u004dM","\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d\u004d"};
_df =[]string {"\u006f\u006e\u0065","\u0074\u0077\u006f","\u0074\u0068\u0072e\u0065","\u0066\u006f\u0075\u0072","\u0066\u0069\u0076\u0065","\u0073\u0069\u0078","\u0073\u0065\u0076e\u006e","\u0065\u0069\u0067h\u0074","\u006e\u0069\u006e\u0065","\u0074\u0065\u006e","\u0065\u006c\u0065\u0076\u0065\u006e","\u0074\u0077\u0065\u006c\u0076\u0065","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e","\u0066i\u0066\u0074\u0065\u0065\u006e","\u0073i\u0078\u0074\u0065\u0065\u006e","\u0073e\u0076\u0065\u006e\u0074\u0065\u0065n","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e"};
_gd =[]string {"\u0074\u0065\u006e","\u0074\u0077\u0065\u006e\u0074\u0079","\u0074\u0068\u0069\u0072\u0074\u0079","\u0066\u006f\u0072t\u0079","\u0066\u0069\u0066t\u0079","\u0073\u0069\u0078t\u0079","\u0073e\u0076\u0065\u006e\u0074\u0079","\u0065\u0069\u0067\u0068\u0074\u0079","\u006e\u0069\u006e\u0065\u0074\u0079"};
_cc =[]string {"\u0066\u0069\u0072s\u0074","\u0073\u0065\u0063\u006f\u006e\u0064","\u0074\u0068\u0069r\u0064","\u0066\u006f\u0075\u0072\u0074\u0068","\u0066\u0069\u0066t\u0068","\u0073\u0069\u0078t\u0068","\u0073e\u0076\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0068","\u006e\u0069\u006et\u0068","\u0074\u0065\u006et\u0068","\u0065\u006c\u0065\u0076\u0065\u006e\u0074\u0068","\u0074w\u0065\u006c\u0066\u0074\u0068","\u0074\u0068\u0069\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066\u006f\u0075\u0072\u0074\u0065\u0065\u006e\u0074\u0068","\u0066i\u0066\u0074\u0065\u0065\u006e\u0074h","\u0073i\u0078\u0074\u0065\u0065\u006e\u0074h","s\u0065\u0076\u0065\u006e\u0074\u0065\u0065\u006e\u0074\u0068","\u0065\u0069\u0067\u0068\u0074\u0065\u0065\u006e\u0074\u0068","\u006e\u0069\u006e\u0065\u0074\u0065\u0065\u006e\u0074\u0068"};
_ca =[]string {"\u0074\u0065\u006et\u0068","\u0074w\u0065\u006e\u0074\u0069\u0065\u0074h","\u0074h\u0069\u0072\u0074\u0069\u0065\u0074h","\u0066\u006f\u0072\u0074\u0069\u0065\u0074\u0068","\u0066\u0069\u0066\u0074\u0069\u0065\u0074\u0068","\u0073\u0069\u0078\u0074\u0069\u0065\u0074\u0068","\u0073\u0065\u0076\u0065\u006e\u0074\u0069\u0065\u0074\u0068","\u0065i\u0067\u0068\u0074\u0069\u0065\u0074h","\u006ei\u006e\u0065\u0074\u0069\u0065\u0074h"};
_db ="\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004a\u004bL\u004d\u004e\u004f\u0050\u0051\u0052\u0053\u0054\u0055\u0056W\u0058\u0059\u005a";);func _ff (_cf string )(_ffc []string ){for _fe :=0;_fe < len (_cf )-2;_fe ++{if string (_cf [_fe ])=="\u0025"{if !_c .Contains (string (_cf [_fe +2:]),"\u0025"){if _fe ==0{_ffc =append (_ffc ,_e .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_cf [_fe ]),string (_cf [_fe +1]),string (_cf [_fe +2:])));
}else {_ffc =append (_ffc ,_e .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073\u0025\u0073",string (_cf [_fe -1]),string (_cf [_fe ]),string (_cf [_fe +1]),string (_cf [_fe +2:])));};}else {_ffc =append (_ffc ,_e .Sprintf ("\u0025\u0073\u0025\u0073\u0025\u0073",string (_cf [_fe ]),string (_cf [_fe +1]),string (_cf [_fe +2])));
};};};return ;};func StringToNumbers (str string )(int ,bool ){_deb :=0;_cdd :=false ;for _ ,_gda :=range []byte (str ){_gda -='0';if _gda > 9{continue ;};_deb =_deb *10+int (_gda );_cdd =true ;};return _deb ,_cdd ;};