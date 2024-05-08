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

package wildcard ;func Match (pattern ,name string )(_bf bool ){if pattern ==""{return name ==pattern ;};if pattern =="\u002a"{return true ;};_de :=make ([]rune ,0,len (name ));_deg :=make ([]rune ,0,len (pattern ));for _ ,_gg :=range name {_de =append (_de ,_gg );
};for _ ,_bc :=range pattern {_deg =append (_deg ,_bc );};_a :=false ;return _gc (_de ,_deg ,_a );};func _gc (_ab ,_ed []rune ,_cb bool )bool {for len (_ed )> 0{switch _ed [0]{default:if len (_ab )==0||_ab [0]!=_ed [0]{return false ;};case '?':if len (_ab )==0&&!_cb {return false ;
};case '*':return _gc (_ab ,_ed [1:],_cb )||(len (_ab )> 0&&_gc (_ab [1:],_ed ,_cb ));};_ab =_ab [1:];_ed =_ed [1:];};return len (_ab )==0&&len (_ed )==0;};func Index (pattern ,name string )(_bd int ){if pattern ==""||pattern =="\u002a"{return 0;};_ae :=make ([]rune ,0,len (name ));
_gf :=make ([]rune ,0,len (pattern ));for _ ,_ga :=range name {_ae =append (_ae ,_ga );};for _ ,_ac :=range pattern {_gf =append (_gf ,_ac );};return _dd (_ae ,_gf ,0);};func MatchSimple (pattern ,name string )bool {if pattern ==""{return name ==pattern ;
};if pattern =="\u002a"{return true ;};_b :=make ([]rune ,0,len (name ));_e :=make ([]rune ,0,len (pattern ));for _ ,_ge :=range name {_b =append (_b ,_ge );};for _ ,_cc :=range pattern {_e =append (_e ,_cc );};_bg :=true ;return _gc (_b ,_e ,_bg );};func _dd (_f ,_cbd []rune ,_dc int )int {for len (_cbd )> 0{switch _cbd [0]{default:if len (_f )==0{return -1;
};if _f [0]!=_cbd [0]{return _dd (_f [1:],_cbd ,_dc +1);};case '?':if len (_f )==0{return -1;};case '*':if len (_f )==0{return -1;};_fg :=_dd (_f ,_cbd [1:],_dc );if _fg !=-1{return _dc ;}else {_fg =_dd (_f [1:],_cbd ,_dc );if _fg !=-1{return _dc ;}else {return -1;
};};};_f =_f [1:];_cbd =_cbd [1:];};return _dc ;};