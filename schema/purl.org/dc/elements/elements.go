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

package elements ;import (_f "encoding/xml";_d "fmt";_fb "github.com/unidoc/unioffice";_da "github.com/unidoc/unioffice/common/logger";);func NewElementsGroupChoice ()*ElementsGroupChoice {_ggb :=&ElementsGroupChoice {};return _ggb };

// Validate validates the SimpleLiteral and its children
func (_gee *SimpleLiteral )Validate ()error {return _gee .ValidateWithPath ("\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c");};

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (_gf *Any )ValidateWithPath (path string )error {if _ef :=_gf .SimpleLiteral .ValidateWithPath (path );_ef !=nil {return _ef ;};return nil ;};

// ValidateWithPath validates the ElementsGroup and its children, prefixing error messages with path
func (_ff *ElementsGroup )ValidateWithPath (path string )error {for _ce ,_fba :=range _ff .Choice {if _dac :=_fba .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_ce ));_dac !=nil {return _dac ;
};};return nil ;};

// Validate validates the Any and its children
func (_ga *Any )Validate ()error {return _ga .ValidateWithPath ("\u0041\u006e\u0079")};func (_gfc *ElementsGroup )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_aa :for {_dc ,_eec :=d .Token ();if _eec !=nil {return _eec ;};switch _ggc :=_dc .(type ){case _f .StartElement :switch _ggc .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_dd :=NewElementsGroupChoice ();
if _dae :=d .DecodeElement (&_dd .Any ,&_ggc );_dae !=nil {return _dae ;};_gfc .Choice =append (_gfc .Choice ,_dd );default:_da .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067 \u0075\u006e\u0073up\u0070\u006f\u0072\u0074\u0065\u0064 \u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006de\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070 \u0025\u0076",_ggc .Name );
if _bc :=d .Skip ();_bc !=nil {return _bc ;};};case _f .EndElement :break _aa ;case _f .CharData :};};return nil ;};

// ValidateWithPath validates the ElementContainer and its children, prefixing error messages with path
func (_cc *ElementContainer )ValidateWithPath (path string )error {for _a ,_ecb :=range _cc .Choice {if _fe :=_ecb .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002f\u0043\u0068\u006f\u0069\u0063\u0065\u005b\u0025\u0064\u005d",path ,_a ));_fe !=nil {return _fe ;
};};return nil ;};type ElementsGroupChoice struct{Any []*Any ;};func (_af *ElementsGroupChoice )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_afe :for {_gab ,_fff :=d .Token ();if _fff !=nil {return _fff ;};switch _aaa :=_gab .(type ){case _f .StartElement :switch _aaa .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_cf :=NewAny ();
if _bcg :=d .DecodeElement (_cf ,&_aaa );_bcg !=nil {return _bcg ;};_af .Any =append (_af .Any ,_cf );default:_da .Log .Debug ("\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070o\u0072\u0074\u0065\u0064\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020o\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072ou\u0070\u0043\u0068\u006f\u0069\u0063\u0065\u0020\u0025\u0076",_aaa .Name );
if _dgf :=d .Skip ();_dgf !=nil {return _dgf ;};};case _f .EndElement :break _afe ;case _f .CharData :};};return nil ;};func (_gfg *ElementsGroup )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {if _gfg .Choice !=nil {for _ ,_ge :=range _gfg .Choice {_ge .MarshalXML (e ,_f .StartElement {});
};};return nil ;};func NewElementsGroup ()*ElementsGroup {_bfc :=&ElementsGroup {};return _bfc };func NewAny ()*Any {_e :=&Any {};_e .SimpleLiteral =*NewSimpleLiteral ();return _e };type ElementsGroup struct{Choice []*ElementsGroupChoice ;};

// ValidateWithPath validates the SimpleLiteral and its children, prefixing error messages with path
func (_ebb *SimpleLiteral )ValidateWithPath (path string )error {return nil };type SimpleLiteral struct{};func (_cd *Any )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {return _cd .SimpleLiteral .MarshalXML (e ,start );};type Any struct{SimpleLiteral };
func (_b *ElementContainer )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {start .Name .Local ="\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072";e .EncodeToken (start );if _b .Choice !=nil {for _ ,_cb :=range _b .Choice {_cb .MarshalXML (e ,_f .StartElement {});
};};e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};func (_gbf *SimpleLiteral )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {for {_bfcd ,_ba :=d .Token ();if _ba !=nil {return _d .Errorf ("\u0070a\u0072\u0073\u0069\u006eg\u0020\u0053\u0069\u006d\u0070l\u0065L\u0069t\u0065\u0072\u0061\u006c\u003a\u0020\u0025s",_ba );
};if _bgb ,_gff :=_bfcd .(_f .EndElement );_gff &&_bgb .Name ==start .Name {break ;};};return nil ;};type ElementContainer struct{Choice []*ElementsGroupChoice ;};

// Validate validates the ElementsGroupChoice and its children
func (_fee *ElementsGroupChoice )Validate ()error {return _fee .ValidateWithPath ("\u0045\u006c\u0065\u006den\u0074\u0073\u0047\u0072\u006f\u0075\u0070\u0043\u0068\u006f\u0069\u0063\u0065");};func (_g *Any )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_g .SimpleLiteral =*NewSimpleLiteral ();
for {_gb ,_fg :=d .Token ();if _fg !=nil {return _d .Errorf ("\u0070a\u0072s\u0069\u006e\u0067\u0020\u0041\u006e\u0079\u003a\u0020\u0025\u0073",_fg );};if _gg ,_eb :=_gb .(_f .EndElement );_eb &&_gg .Name ==start .Name {break ;};};return nil ;};func (_gd *ElementsGroupChoice )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {if _gd .Any !=nil {_efe :=_f .StartElement {Name :_f .Name {Local :"\u0064\u0063\u003a\u0061\u006e\u0079"}};
for _ ,_bffd :=range _gd .Any {e .EncodeElement (_bffd ,_efe );};};return nil ;};func (_bf *ElementContainer )UnmarshalXML (d *_f .Decoder ,start _f .StartElement )error {_df :for {_ed ,_eg :=d .Token ();if _eg !=nil {return _eg ;};switch _bg :=_ed .(type ){case _f .StartElement :switch _bg .Name {case _f .Name {Space :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f",Local :"\u0061\u006e\u0079"}:_eef :=NewElementsGroupChoice ();
if _dg :=d .DecodeElement (&_eef .Any ,&_bg );_dg !=nil {return _dg ;};_bf .Choice =append (_bf .Choice ,_eef );default:_da .Log .Debug ("\u0073\u006b\u0069\u0070\u0070i\u006e\u0067\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065d\u0020\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0020\u006f\u006e\u0020\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072\u0020\u0025v",_bg .Name );
if _bff :=d .Skip ();_bff !=nil {return _bff ;};};case _f .EndElement :break _df ;case _f .CharData :};};return nil ;};func NewSimpleLiteral ()*SimpleLiteral {_ae :=&SimpleLiteral {};return _ae };func (_cg *SimpleLiteral )MarshalXML (e *_f .Encoder ,start _f .StartElement )error {e .EncodeToken (start );
e .EncodeToken (_f .EndElement {Name :start .Name });return nil ;};

// Validate validates the ElementsGroup and its children
func (_gbc *ElementsGroup )Validate ()error {return _gbc .ValidateWithPath ("\u0045\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070");};

// Validate validates the ElementContainer and its children
func (_ggd *ElementContainer )Validate ()error {return _ggd .ValidateWithPath ("\u0045\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072");};func NewElementContainer ()*ElementContainer {_ec :=&ElementContainer {};return _ec };


// ValidateWithPath validates the ElementsGroupChoice and its children, prefixing error messages with path
func (_bb *ElementsGroupChoice )ValidateWithPath (path string )error {for _cbe ,_caf :=range _bb .Any {if _fgb :=_caf .ValidateWithPath (_d .Sprintf ("\u0025\u0073\u002f\u0041\u006e\u0079\u005b\u0025\u0064\u005d",path ,_cbe ));_fgb !=nil {return _fgb ;
};};return nil ;};func init (){_fb .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0053\u0069\u006d\u0070\u006c\u0065\u004c\u0069\u0074\u0065\u0072\u0061\u006c",NewSimpleLiteral );
_fb .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006ce\u006d\u0065\u006et\u0043\u006f\u006e\u0074\u0061\u0069\u006e\u0065\u0072",NewElementContainer );
_fb .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0061\u006e\u0079",NewAny );_fb .RegisterConstructor ("\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0070\u0075\u0072\u006c\u002e\u006f\u0072\u0067/\u0064c\u002f\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u002f\u0031\u002e\u0031\u002f","\u0065\u006c\u0065\u006d\u0065\u006e\u0074\u0073\u0047\u0072\u006f\u0075\u0070",NewElementsGroup );
};