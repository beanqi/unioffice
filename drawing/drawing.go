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

package drawing ;import (_g "github.com/unidoc/unioffice";_f "github.com/unidoc/unioffice/color";_dc "github.com/unidoc/unioffice/measurement";_b "github.com/unidoc/unioffice/schema/soo/dml";);

// Paragraph is a paragraph within a document.
type Paragraph struct{_e *_b .CT_TextParagraph };

// Properties returns the run's properties.
func (_eg Run )Properties ()RunProperties {if _eg ._bc .R ==nil {_eg ._bc .R =_b .NewCT_RegularTextRun ();};if _eg ._bc .R .RPr ==nil {_eg ._bc .R .RPr =_b .NewCT_TextCharacterProperties ();};return RunProperties {_eg ._bc .R .RPr };};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_fb *_b .CT_TextParagraphProperties ;};

// SetHeight sets the height of the shape.
func (_ba ShapeProperties )SetHeight (h _dc .Distance ){_ba .ensureXfrm ();if _ba ._ad .Xfrm .Ext ==nil {_ba ._ad .Xfrm .Ext =_b .NewCT_PositiveSize2D ();};_ba ._ad .Xfrm .Ext .CyAttr =int64 (h /_dc .EMU );};

// X returns the inner wrapped XML type.
func (_dad ShapeProperties )X ()*_b .CT_ShapeProperties {return _dad ._ad };

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_dg ShapeProperties )SetFlipHorizontal (b bool ){_dg .ensureXfrm ();if !b {_dg ._ad .Xfrm .FlipHAttr =nil ;}else {_dg ._ad .Xfrm .FlipHAttr =_g .Bool (true );};};

// SetGeometry sets the shape type of the shape
func (_dadf ShapeProperties )SetGeometry (g _b .ST_ShapeType ){if _dadf ._ad .PrstGeom ==nil {_dadf ._ad .PrstGeom =_b .NewCT_PresetGeometry2D ();};_dadf ._ad .PrstGeom .PrstAttr =g ;};

// X returns the inner wrapped XML type.
func (_cd Paragraph )X ()*_b .CT_TextParagraph {return _cd ._e };

// SetNumbered controls if bullets are numbered or not.
func (_ef ParagraphProperties )SetNumbered (scheme _b .ST_TextAutonumberScheme ){if scheme ==_b .ST_TextAutonumberSchemeUnset {_ef ._fb .BuAutoNum =nil ;}else {_ef ._fb .BuAutoNum =_b .NewCT_TextAutonumberBullet ();_ef ._fb .BuAutoNum .TypeAttr =scheme ;
};};

// SetBulletFont controls the font for the bullet character.
func (_ge ParagraphProperties )SetBulletFont (f string ){if f ==""{_ge ._fb .BuFont =nil ;}else {_ge ._fb .BuFont =_b .NewCT_TextFont ();_ge ._fb .BuFont .TypefaceAttr =f ;};};

// Run is a run within a paragraph.
type Run struct{_bc *_b .EG_TextRun };

// SetAlign controls the paragraph alignment
func (_gea ParagraphProperties )SetAlign (a _b .ST_TextAlignType ){_gea ._fb .AlgnAttr =a };

// SetBold controls the bolding of a run.
func (_cb RunProperties )SetBold (b bool ){_cb ._fe .BAttr =_g .Bool (b )};func (_fa ShapeProperties )ensureXfrm (){if _fa ._ad .Xfrm ==nil {_fa ._ad .Xfrm =_b .NewCT_Transform2D ();};};

// SetBulletChar sets the bullet character for the paragraph.
func (_daf ParagraphProperties )SetBulletChar (c string ){if c ==""{_daf ._fb .BuChar =nil ;}else {_daf ._fb .BuChar =_b .NewCT_TextCharBullet ();_daf ._fb .BuChar .CharAttr =c ;};};func (_da LineProperties )clearFill (){_da ._bg .NoFill =nil ;_da ._bg .GradFill =nil ;
_da ._bg .SolidFill =nil ;_da ._bg .PattFill =nil ;};func (_gee ShapeProperties )clearFill (){_gee ._ad .NoFill =nil ;_gee ._ad .BlipFill =nil ;_gee ._ad .GradFill =nil ;_gee ._ad .GrpFill =nil ;_gee ._ad .SolidFill =nil ;_gee ._ad .PattFill =nil ;};func MakeShapeProperties (x *_b .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};


// SetWidth sets the width of the shape.
func (_fc ShapeProperties )SetWidth (w _dc .Distance ){_fc .ensureXfrm ();if _fc ._ad .Xfrm .Ext ==nil {_fc ._ad .Xfrm .Ext =_b .NewCT_PositiveSize2D ();};_fc ._ad .Xfrm .Ext .CxAttr =int64 (w /_dc .EMU );};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;
LineJoinMiter ;);

// RunProperties controls the run properties.
type RunProperties struct{_fe *_b .CT_TextCharacterProperties ;};

// X returns the inner wrapped XML type.
func (_a LineProperties )X ()*_b .CT_LineProperties {return _a ._bg };

// SetJoin sets the line join style.
func (_ca LineProperties )SetJoin (e LineJoin ){_ca ._bg .Round =nil ;_ca ._bg .Miter =nil ;_ca ._bg .Bevel =nil ;switch e {case LineJoinRound :_ca ._bg .Round =_b .NewCT_LineJoinRound ();case LineJoinBevel :_ca ._bg .Bevel =_b .NewCT_LineJoinBevel ();
case LineJoinMiter :_ca ._bg .Miter =_b .NewCT_LineJoinMiterProperties ();};};

// Properties returns the paragraph properties.
func (_cf Paragraph )Properties ()ParagraphProperties {if _cf ._e .PPr ==nil {_cf ._e .PPr =_b .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_cf ._e .PPr );};

// SetText sets the run's text contents.
func (_ce Run )SetText (s string ){_ce ._bc .Br =nil ;_ce ._bc .Fld =nil ;if _ce ._bc .R ==nil {_ce ._bc .R =_b .NewCT_RegularTextRun ();};_ce ._bc .R .T =s ;};func (_bf LineProperties )SetSolidFill (c _f .Color ){_bf .clearFill ();_bf ._bg .SolidFill =_b .NewCT_SolidColorFillProperties ();
_bf ._bg .SolidFill .SrgbClr =_b .NewCT_SRgbColor ();_bf ._bg .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};type LineProperties struct{_bg *_b .CT_LineProperties };func (_fg ShapeProperties )SetSolidFill (c _f .Color ){_fg .clearFill ();_fg ._ad .SolidFill =_b .NewCT_SolidColorFillProperties ();
_fg ._ad .SolidFill .SrgbClr =_b .NewCT_SRgbColor ();_fg ._ad .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};type ShapeProperties struct{_ad *_b .CT_ShapeProperties };

// AddBreak adds a new line break to a paragraph.
func (_bga Paragraph )AddBreak (){_ga :=_b .NewEG_TextRun ();_ga .Br =_b .NewCT_TextLineBreak ();_bga ._e .EG_TextRun =append (_bga ._e .EG_TextRun ,_ga );};

// SetFlipVertical controls if the shape is flipped vertically.
func (_ceg ShapeProperties )SetFlipVertical (b bool ){_ceg .ensureXfrm ();if !b {_ceg ._ad .Xfrm .FlipVAttr =nil ;}else {_ceg ._ad .Xfrm .FlipVAttr =_g .Bool (true );};};

// X returns the inner wrapped XML type.
func (_af Run )X ()*_b .EG_TextRun {return _af ._bc };

// AddRun adds a new run to a paragraph.
func (_de Paragraph )AddRun ()Run {_dd :=MakeRun (_b .NewEG_TextRun ());_de ._e .EG_TextRun =append (_de ._e .EG_TextRun ,_dd .X ());return _dd ;};func (_c LineProperties )SetNoFill (){_c .clearFill ();_c ._bg .NoFill =_b .NewCT_NoFillProperties ()};

// SetFont controls the font of a run.
func (_afa RunProperties )SetFont (s string ){_afa ._fe .Latin =_b .NewCT_TextFont ();_afa ._fe .Latin .TypefaceAttr =s ;};

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_b .EG_TextRun )Run {return Run {x }};

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_b .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};func (_dcd ShapeProperties )SetNoFill (){_dcd .clearFill ();_dcd ._ad .NoFill =_b .NewCT_NoFillProperties ();};func (_ea ShapeProperties )LineProperties ()LineProperties {if _ea ._ad .Ln ==nil {_ea ._ad .Ln =_b .NewCT_LineProperties ();
};return LineProperties {_ea ._ad .Ln };};

// SetSize sets the width and height of the shape.
func (_ac ShapeProperties )SetSize (w ,h _dc .Distance ){_ac .SetWidth (w );_ac .SetHeight (h )};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_b .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};

// SetSolidFill controls the text color of a run.
func (_be RunProperties )SetSolidFill (c _f .Color ){_be ._fe .NoFill =nil ;_be ._fe .BlipFill =nil ;_be ._fe .GradFill =nil ;_be ._fe .GrpFill =nil ;_be ._fe .PattFill =nil ;_be ._fe .SolidFill =_b .NewCT_SolidColorFillProperties ();_be ._fe .SolidFill .SrgbClr =_b .NewCT_SRgbColor ();
_be ._fe .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// SetLevel sets the level of indentation of a paragraph.
func (_gc ParagraphProperties )SetLevel (idx int32 ){_gc ._fb .LvlAttr =_g .Int32 (idx )};

// GetPosition gets the position of the shape in EMU.
func (_feg ShapeProperties )GetPosition ()(int64 ,int64 ){_feg .ensureXfrm ();if _feg ._ad .Xfrm .Off ==nil {_feg ._ad .Xfrm .Off =_b .NewCT_Point2D ();};return *_feg ._ad .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_feg ._ad .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;
};

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_gb LineProperties )SetWidth (w _dc .Distance ){_gb ._bg .WAttr =_g .Int32 (int32 (w /_dc .EMU ))};

// LineJoin is the type of line join
type LineJoin byte ;

// SetPosition sets the position of the shape.
func (_feb ShapeProperties )SetPosition (x ,y _dc .Distance ){_feb .ensureXfrm ();if _feb ._ad .Xfrm .Off ==nil {_feb ._ad .Xfrm .Off =_b .NewCT_Point2D ();};_feb ._ad .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_g .Int64 (int64 (x /_dc .EMU ));_feb ._ad .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_g .Int64 (int64 (y /_dc .EMU ));
};

// X returns the inner wrapped XML type.
func (_bge ParagraphProperties )X ()*_b .CT_TextParagraphProperties {return _bge ._fb };

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_b .CT_TextParagraph )Paragraph {return Paragraph {x }};

// SetSize sets the font size of the run text
func (_bd RunProperties )SetSize (sz _dc .Distance ){_bd ._fe .SzAttr =_g .Int32 (int32 (sz /_dc .HundredthPoint ));};