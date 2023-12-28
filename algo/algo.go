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

package algo ;import _a "strconv";func _f (_af byte )bool {return _af >='0'&&_af <='9'};

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_fb ,_dg :=0,0;for _fb < len (lhs )&&_dg < len (rhs ){_dd :=lhs [_fb ];_fa :=rhs [_dg ];_db :=_f (_dd );_fd :=_f (_fa );switch {case _db &&!_fd :return true ;case !_db &&_fd :return false ;case !_db &&!_fd :if _dd !=_fa {return _dd < _fa ;
};_fb ++;_dg ++;default:_b :=_fb +1;_fdf :=_dg +1;for _b < len (lhs )&&_f (lhs [_b ]){_b ++;};for _fdf < len (rhs )&&_f (rhs [_fdf ]){_fdf ++;};_c ,_ :=_a .ParseUint (lhs [_fb :_b ],10,64);_fad ,_ :=_a .ParseUint (rhs [_fb :_fdf ],10,64);if _c !=_fad {return _c < _fad ;
};_fb =_b ;_dg =_fdf ;};};return len (lhs )< len (rhs );};func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_aa :=make ([]byte ,len (s )*cnt );_g :=[]byte (s );for _fdfg :=0;_fdfg < cnt ;_fdfg ++{copy (_aa [_fdfg :],_g );};return string (_aa );
};