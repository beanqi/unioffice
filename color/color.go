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

// Package color provides color handling structures and functions for use across
// all of the document types.
package color ;import (_e "fmt";_ge "github.com/unidoc/unioffice";);var LightSkyBlue =Color {0x87,0xCE,0xFA,255,false };var Aquamarine =Color {0x7F,0xFF,0xD4,255,false };var Cyan =Color {0x00,0xFF,0xFF,255,false };var Orchid =Color {0xDA,0x70,0xD6,255,false };
var Ivory =Color {0xFF,0xFF,0xF0,255,false };var Gainsboro =Color {0xDC,0xDC,0xDC,255,false };var WhiteSmoke =Color {0xF5,0xF5,0xF5,255,false };var SandyBrown =Color {0xF4,0xA4,0x60,255,false };var Sienna =Color {0xA0,0x52,0x2D,255,false };var LightYellow =Color {0xFF,0xFF,0xE0,255,false };
var LightSteelBlue =Color {0xB0,0xC4,0xDE,255,false };var LightBlue =Color {0xAD,0xD8,0xE6,255,false };var Lime =Color {0x00,0xFF,0x00,255,false };

// Color is a 24 bit color that can be converted to
// internal ECMA-376 formats as needed.
type Color struct{_d ,_f ,_a ,_gd uint8 ;_fd bool ;};var DeepPink =Color {0xFF,0x14,0x93,255,false };var OliveDrab =Color {0x6B,0x8E,0x23,255,false };var DarkCyan =Color {0x00,0x8B,0x8B,255,false };var DarkGoldenRod =Color {0xB8,0x86,0x0B,255,false };var MediumVioletRed =Color {0xC7,0x15,0x85,255,false };
var PaleGoldenRod =Color {0xEE,0xE8,0xAA,255,false };var AntiqueWhite =Color {0xFA,0xEB,0xD7,255,false };var DarkGreen =Color {0x00,0x64,0x00,255,false };var Gold =Color {0xFF,0xD7,0x00,255,false };var DarkBlue =Color {0x00,0x00,0x8B,255,false };var Maroon =Color {0x80,0x00,0x00,255,false };
var SteelBlue =Color {0x46,0x82,0xB4,255,false };var DarkViolet =Color {0x94,0x00,0xD3,255,false };var LightGray =Color {0xD3,0xD3,0xD3,255,false };var LightCyan =Color {0xE0,0xFF,0xFF,255,false };var MediumBlue =Color {0x00,0x00,0xCD,255,false };var YellowGreen =Color {0x9A,0xCD,0x32,255,false };
var PowderBlue =Color {0xB0,0xE0,0xE6,255,false };var Salmon =Color {0xFA,0x80,0x72,255,false };var MediumSlateBlue =Color {0x7B,0x68,0xEE,255,false };var LavenderBlush =Color {0xFF,0xF0,0xF5,255,false };var Wheat =Color {0xF5,0xDE,0xB3,255,false };var Coral =Color {0xFF,0x7F,0x50,255,false };
var BlueViolet =Color {0x8A,0x2B,0xE2,255,false };var DarkTurquoise =Color {0x00,0xCE,0xD1,255,false };var Navy =Color {0x00,0x00,0x80,255,false };var RebeccaPurple =Color {0x66,0x33,0x99,255,false };var Cornsilk =Color {0xFF,0xF8,0xDC,255,false };var MediumOrchid =Color {0xBA,0x55,0xD3,255,false };
var MintCream =Color {0xF5,0xFF,0xFA,255,false };var MistyRose =Color {0xFF,0xE4,0xE1,255,false };var Auto =Color {0,0,0,255,true };var MediumPurple =Color {0x93,0x70,0xDB,255,false };var OrangeRed =Color {0xFF,0x45,0x00,255,false };var Linen =Color {0xFA,0xF0,0xE6,255,false };
var Chocolate =Color {0xD2,0x69,0x1E,255,false };var Tan =Color {0xD2,0xB4,0x8C,255,false };var Crimson =Color {0xDC,0x14,0x3C,255,false };var CornflowerBlue =Color {0x64,0x95,0xED,255,false };var Yellow =Color {0xFF,0xFF,0x00,255,false };var DarkSlateGrey =Color {0x2F,0x4F,0x4F,255,false };
var Olive =Color {0x80,0x80,0x00,255,false };var PapayaWhip =Color {0xFF,0xEF,0xD5,255,false };var Gray =Color {0x80,0x80,0x80,255,false };var Azure =Color {0xF0,0xFF,0xFF,255,false };var FloralWhite =Color {0xFF,0xFA,0xF0,255,false };var Plum =Color {0xDD,0xA0,0xDD,255,false };
var LightPink =Color {0xFF,0xB6,0xC1,255,false };var GoldenRod =Color {0xDA,0xA5,0x20,255,false };var SpringGreen =Color {0x00,0xFF,0x7F,255,false };var SeaShell =Color {0xFF,0xF5,0xEE,255,false };var LightCoral =Color {0xF0,0x80,0x80,255,false };var DimGray =Color {0x69,0x69,0x69,255,false };
var PeachPuff =Color {0xFF,0xDA,0xB9,255,false };var SaddleBrown =Color {0x8B,0x45,0x13,255,false };var Snow =Color {0xFF,0xFA,0xFA,255,false };var PaleGreen =Color {0x98,0xFB,0x98,255,false };

// RGBA constructs a new RGBA color with a given red, green, blue and alpha
// value.
func RGBA (r ,g ,b ,a uint8 )Color {return Color {r ,g ,b ,a ,false }};var GreenYellow =Color {0xAD,0xFF,0x2F,255,false };var DarkOrchid =Color {0x99,0x32,0xCC,255,false };var Bisque =Color {0xFF,0xE4,0xC4,255,false };var MidnightBlue =Color {0x19,0x19,0x70,255,false };
var LightSlateGrey =Color {0x77,0x88,0x99,255,false };var Chartreuse =Color {0x7F,0xFF,0x00,255,false };var RosyBrown =Color {0xBC,0x8F,0x8F,255,false };var LightGreen =Color {0x90,0xEE,0x90,255,false };var Green =Color {0x00,0x80,0x00,255,false };var Silver =Color {0xC0,0xC0,0xC0,255,false };
var SlateGrey =Color {0x70,0x80,0x90,255,false };var Red =Color {0xFF,0x00,0x00,255,false };var LightSlateGray =Color {0x77,0x88,0x99,255,false };var DeepSkyBlue =Color {0x00,0xBF,0xFF,255,false };var ForestGreen =Color {0x22,0x8B,0x22,255,false };var Lavender =Color {0xE6,0xE6,0xFA,255,false };
var MediumAquaMarine =Color {0x66,0xCD,0xAA,255,false };var DarkRed =Color {0x8B,0x00,0x00,255,false };var Fuchsia =Color {0xFF,0x00,0xFF,255,false };var Teal =Color {0x00,0x80,0x80,255,false };var Blue =Color {0x00,0x00,0xFF,255,false };var DarkGrey =Color {0xA9,0xA9,0xA9,255,false };


// RGB constructs a new RGB color with a given red, green and blue value.
func RGB (r ,g ,b uint8 )Color {return Color {r ,g ,b ,255,false }};var DarkGray =Color {0xA9,0xA9,0xA9,255,false };var Indigo =Color {0x4B,0x00,0x82,255,false };var OldLace =Color {0xFD,0xF5,0xE6,255,false };

// AsRGBString is used by the various wrappers to return a pointer
// to a string containing a six digit hex RGB value.
func (_df Color )AsRGBString ()*string {return _ge .Stringf ("\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078",_df ._d ,_df ._f ,_df ._a );};var DarkOrange =Color {0xFF,0x8C,0x00,255,false };var MediumSeaGreen =Color {0x3C,0xB3,0x71,255,false };
var AliceBlue =Color {0xF0,0xF8,0xFF,255,false };var LawnGreen =Color {0x7C,0xFC,0x00,255,false };var SlateGray =Color {0x70,0x80,0x90,255,false };var GhostWhite =Color {0xF8,0xF8,0xFF,255,false };var Tomato =Color {0xFF,0x63,0x47,255,false };var SlateBlue =Color {0x6A,0x5A,0xCD,255,false };
var LightGoldenRodYellow =Color {0xFA,0xFA,0xD2,255,false };var Aqua =Color {0x00,0xFF,0xFF,255,false };var NavajoWhite =Color {0xFF,0xDE,0xAD,255,false };var Orange =Color {0xFF,0xA5,0x00,255,false };var DodgerBlue =Color {0x1E,0x90,0xFF,255,false };var Brown =Color {0xA5,0x2A,0x2A,255,false };
var PaleTurquoise =Color {0xAF,0xEE,0xEE,255,false };var Beige =Color {0xF5,0xF5,0xDC,255,false };var Purple =Color {0x80,0x00,0x80,255,false };

// IsAuto returns true if the color is the 'Auto' type.  If the
// field doesn't support an Auto color, then black is used.
func (_dd Color )IsAuto ()bool {return _dd ._fd };var LemonChiffon =Color {0xFF,0xFA,0xCD,255,false };var SkyBlue =Color {0x87,0xCE,0xEB,255,false };var DarkSlateBlue =Color {0x48,0x3D,0x8B,255,false };func FromHex (s string )Color {if len (s )==0{return Auto ;
};if s [0]=='#'{s =s [1:];};var _ff ,_de ,_fg uint8 ;_gc ,_ :=_e .Sscanf (s ,"\u0025\u0030\u0032x\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078",&_ff ,&_de ,&_fg );if _gc ==3{return RGB (_ff ,_de ,_fg );};return Auto ;};var DarkSalmon =Color {0xE9,0x96,0x7A,255,false };
var DarkSlateGray =Color {0x2F,0x4F,0x4F,255,false };var Pink =Color {0xFF,0xC0,0xCB,255,false };var White =Color {0xFF,0xFF,0xFF,255,false };var DarkMagenta =Color {0x8B,0x00,0x8B,255,false };var MediumSpringGreen =Color {0x00,0xFA,0x9A,255,false };var SuccessGreen =Color {0x00,0xCC,0x00,255,false };
var IndianRed =Color {0xCD,0x5C,0x5C,255,false };var LightGrey =Color {0xD3,0xD3,0xD3,255,false };var HotPink =Color {0xFF,0x69,0xB4,255,false };var DimGrey =Color {0x69,0x69,0x69,255,false };var CadetBlue =Color {0x5F,0x9E,0xA0,255,false };var SeaGreen =Color {0x2E,0x8B,0x57,255,false };
var Turquoise =Color {0x40,0xE0,0xD0,255,false };var LightSeaGreen =Color {0x20,0xB2,0xAA,255,false };var FireBrick =Color {0xB2,0x22,0x22,255,false };var Black =Color {0x00,0x00,0x00,255,false };var MediumTurquoise =Color {0x48,0xD1,0xCC,255,false };var LimeGreen =Color {0x32,0xCD,0x32,255,false };
var DarkOliveGreen =Color {0x55,0x6B,0x2F,255,false };var HoneyDew =Color {0xF0,0xFF,0xF0,255,false };var Magenta =Color {0xFF,0x00,0xFF,255,false };var Moccasin =Color {0xFF,0xE4,0xB5,255,false };var DarkKhaki =Color {0xBD,0xB7,0x6B,255,false };var Khaki =Color {0xF0,0xE6,0x8C,255,false };
var RoyalBlue =Color {0x41,0x69,0xE1,255,false };var LightSalmon =Color {0xFF,0xA0,0x7A,255,false };

// AsRGBAString is used by the various wrappers to return a pointer
// to a string containing a six digit hex RGB value.
func (_gf Color )AsRGBAString ()*string {return _ge .Stringf ("\u0025\u00302\u0078\u0025\u00302\u0078\u0025\u0030\u0032\u0078\u0025\u0030\u0032\u0078",_gf ._gd ,_gf ._d ,_gf ._f ,_gf ._a );};var Violet =Color {0xEE,0x82,0xEE,255,false };var BurlyWood =Color {0xDE,0xB8,0x87,255,false };
var Peru =Color {0xCD,0x85,0x3F,255,false };var Thistle =Color {0xD8,0xBF,0xD8,255,false };var PaleVioletRed =Color {0xDB,0x70,0x93,255,false };var DarkSeaGreen =Color {0x8F,0xBC,0x8F,255,false };var BlanchedAlmond =Color {0xFF,0xEB,0xCD,255,false };