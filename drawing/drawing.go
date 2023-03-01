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

package drawing ;import (_d "github.com/unidoc/unioffice";_a "github.com/unidoc/unioffice/color";_cc "github.com/unidoc/unioffice/measurement";_c "github.com/unidoc/unioffice/schema/soo/dml";);

// SetLevel sets the level of indentation of a paragraph.
func (_ae ParagraphProperties )SetLevel (idx int32 ){_ae ._ge .LvlAttr =_d .Int32 (idx )};

// SetFlipVertical controls if the shape is flipped vertically.
func (_abc ShapeProperties )SetFlipVertical (b bool ){_abc .ensureXfrm ();if !b {_abc ._df .Xfrm .FlipVAttr =nil ;}else {_abc ._df .Xfrm .FlipVAttr =_d .Bool (true );};};func (_ab ShapeProperties )clearFill (){_ab ._df .NoFill =nil ;_ab ._df .BlipFill =nil ;_ab ._df .GradFill =nil ;_ab ._df .GrpFill =nil ;_ab ._df .SolidFill =nil ;_ab ._df .PattFill =nil ;};type ShapeProperties struct{_df *_c .CT_ShapeProperties };

// X returns the inner wrapped XML type.
func (_ba Paragraph )X ()*_c .CT_TextParagraph {return _ba ._fg };func (_da LineProperties )SetNoFill (){_da .clearFill ();_da ._f .NoFill =_c .NewCT_NoFillProperties ()};func (_daf ShapeProperties )SetSolidFill (c _a .Color ){_daf .clearFill ();_daf ._df .SolidFill =_c .NewCT_SolidColorFillProperties ();_daf ._df .SolidFill .SrgbClr =_c .NewCT_SRgbColor ();_daf ._df .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// MakeRunProperties constructs a new RunProperties wrapper.
func MakeRunProperties (x *_c .CT_TextCharacterProperties )RunProperties {return RunProperties {x }};func (_cdd ShapeProperties )LineProperties ()LineProperties {if _cdd ._df .Ln ==nil {_cdd ._df .Ln =_c .NewCT_LineProperties ();};return LineProperties {_cdd ._df .Ln };};

// X returns the inner wrapped XML type.
func (_cff ShapeProperties )X ()*_c .CT_ShapeProperties {return _cff ._df };

// Properties returns the paragraph properties.
func (_cd Paragraph )Properties ()ParagraphProperties {if _cd ._fg .PPr ==nil {_cd ._fg .PPr =_c .NewCT_TextParagraphProperties ();};return MakeParagraphProperties (_cd ._fg .PPr );};

// SetFont controls the font of a run.
func (_de RunProperties )SetFont (s string ){_de ._bad .Latin =_c .NewCT_TextFont ();_de ._bad .Latin .TypefaceAttr =s ;};

// SetSize sets the width and height of the shape.
func (_fe ShapeProperties )SetSize (w ,h _cc .Distance ){_fe .SetWidth (w );_fe .SetHeight (h )};

// SetJoin sets the line join style.
func (_g LineProperties )SetJoin (e LineJoin ){_g ._f .Round =nil ;_g ._f .Miter =nil ;_g ._f .Bevel =nil ;switch e {case LineJoinRound :_g ._f .Round =_c .NewCT_LineJoinRound ();case LineJoinBevel :_g ._f .Bevel =_c .NewCT_LineJoinBevel ();case LineJoinMiter :_g ._f .Miter =_c .NewCT_LineJoinMiterProperties ();};};func (_dg LineProperties )SetSolidFill (c _a .Color ){_dg .clearFill ();_dg ._f .SolidFill =_c .NewCT_SolidColorFillProperties ();_dg ._f .SolidFill .SrgbClr =_c .NewCT_SRgbColor ();_dg ._f .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};

// X returns the inner wrapped XML type.
func (_b LineProperties )X ()*_c .CT_LineProperties {return _b ._f };

// SetNumbered controls if bullets are numbered or not.
func (_gg ParagraphProperties )SetNumbered (scheme _c .ST_TextAutonumberScheme ){if scheme ==_c .ST_TextAutonumberSchemeUnset {_gg ._ge .BuAutoNum =nil ;}else {_gg ._ge .BuAutoNum =_c .NewCT_TextAutonumberBullet ();_gg ._ge .BuAutoNum .TypeAttr =scheme ;};};

// MakeParagraphProperties constructs a new ParagraphProperties wrapper.
func MakeParagraphProperties (x *_c .CT_TextParagraphProperties )ParagraphProperties {return ParagraphProperties {x };};

// SetBulletFont controls the font for the bullet character.
func (_fd ParagraphProperties )SetBulletFont (f string ){if f ==""{_fd ._ge .BuFont =nil ;}else {_fd ._ge .BuFont =_c .NewCT_TextFont ();_fd ._ge .BuFont .TypefaceAttr =f ;};};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;LineJoinMiter ;);

// SetWidth sets the width of the shape.
func (_aa ShapeProperties )SetWidth (w _cc .Distance ){_aa .ensureXfrm ();if _aa ._df .Xfrm .Ext ==nil {_aa ._df .Xfrm .Ext =_c .NewCT_PositiveSize2D ();};_aa ._df .Xfrm .Ext .CxAttr =int64 (w /_cc .EMU );};

// Run is a run within a paragraph.
type Run struct{_fgc *_c .EG_TextRun };

// SetFlipHorizontal controls if the shape is flipped horizontally.
func (_fba ShapeProperties )SetFlipHorizontal (b bool ){_fba .ensureXfrm ();if !b {_fba ._df .Xfrm .FlipHAttr =nil ;}else {_fba ._df .Xfrm .FlipHAttr =_d .Bool (true );};};

// SetText sets the run's text contents.
func (_dgb Run )SetText (s string ){_dgb ._fgc .Br =nil ;_dgb ._fgc .Fld =nil ;if _dgb ._fgc .R ==nil {_dgb ._fgc .R =_c .NewCT_RegularTextRun ();};_dgb ._fgc .R .T =s ;};

// SetPosition sets the position of the shape.
func (_bb ShapeProperties )SetPosition (x ,y _cc .Distance ){_bb .ensureXfrm ();if _bb ._df .Xfrm .Off ==nil {_bb ._df .Xfrm .Off =_c .NewCT_Point2D ();};_bb ._df .Xfrm .Off .XAttr .ST_CoordinateUnqualified =_d .Int64 (int64 (x /_cc .EMU ));_bb ._df .Xfrm .Off .YAttr .ST_CoordinateUnqualified =_d .Int64 (int64 (y /_cc .EMU ));};func (_fb LineProperties )clearFill (){_fb ._f .NoFill =nil ;_fb ._f .GradFill =nil ;_fb ._f .SolidFill =nil ;_fb ._f .PattFill =nil ;};

// Paragraph is a paragraph within a document.
type Paragraph struct{_fg *_c .CT_TextParagraph };

// Properties returns the run's properties.
func (_ebb Run )Properties ()RunProperties {if _ebb ._fgc .R ==nil {_ebb ._fgc .R =_c .NewCT_RegularTextRun ();};if _ebb ._fgc .R .RPr ==nil {_ebb ._fgc .R .RPr =_c .NewCT_TextCharacterProperties ();};return RunProperties {_ebb ._fgc .R .RPr };};

// ParagraphProperties allows controlling paragraph properties.
type ParagraphProperties struct{_ge *_c .CT_TextParagraphProperties ;};

// X returns the inner wrapped XML type.
func (_eb Run )X ()*_c .EG_TextRun {return _eb ._fgc };

// MakeRun constructs a new Run wrapper.
func MakeRun (x *_c .EG_TextRun )Run {return Run {x }};

// SetBold controls the bolding of a run.
func (_eba RunProperties )SetBold (b bool ){_eba ._bad .BAttr =_d .Bool (b )};func MakeShapeProperties (x *_c .CT_ShapeProperties )ShapeProperties {return ShapeProperties {x }};type LineProperties struct{_f *_c .CT_LineProperties };func (_dagf ShapeProperties )ensureXfrm (){if _dagf ._df .Xfrm ==nil {_dagf ._df .Xfrm =_c .NewCT_Transform2D ();};};

// SetHeight sets the height of the shape.
func (_fc ShapeProperties )SetHeight (h _cc .Distance ){_fc .ensureXfrm ();if _fc ._df .Xfrm .Ext ==nil {_fc ._df .Xfrm .Ext =_c .NewCT_PositiveSize2D ();};_fc ._df .Xfrm .Ext .CyAttr =int64 (h /_cc .EMU );};

// SetGeometry sets the shape type of the shape
func (_aef ShapeProperties )SetGeometry (g _c .ST_ShapeType ){if _aef ._df .PrstGeom ==nil {_aef ._df .PrstGeom =_c .NewCT_PresetGeometry2D ();};_aef ._df .PrstGeom .PrstAttr =g ;};

// SetWidth sets the line width, MS products treat zero as the minimum width
// that can be displayed.
func (_eg LineProperties )SetWidth (w _cc .Distance ){_eg ._f .WAttr =_d .Int32 (int32 (w /_cc .EMU ))};

// SetBulletChar sets the bullet character for the paragraph.
func (_fge ParagraphProperties )SetBulletChar (c string ){if c ==""{_fge ._ge .BuChar =nil ;}else {_fge ._ge .BuChar =_c .NewCT_TextCharBullet ();_fge ._ge .BuChar .CharAttr =c ;};};

// AddRun adds a new run to a paragraph.
func (_cb Paragraph )AddRun ()Run {_be :=MakeRun (_c .NewEG_TextRun ());_cb ._fg .EG_TextRun =append (_cb ._fg .EG_TextRun ,_be .X ());return _be ;};

// GetPosition gets the position of the shape in EMU.
func (_fea ShapeProperties )GetPosition ()(int64 ,int64 ){_fea .ensureXfrm ();if _fea ._df .Xfrm .Off ==nil {_fea ._df .Xfrm .Off =_c .NewCT_Point2D ();};return *_fea ._df .Xfrm .Off .XAttr .ST_CoordinateUnqualified ,*_fea ._df .Xfrm .Off .YAttr .ST_CoordinateUnqualified ;};

// LineJoin is the type of line join
type LineJoin byte ;

// AddBreak adds a new line break to a paragraph.
func (_ccf Paragraph )AddBreak (){_egg :=_c .NewEG_TextRun ();_egg .Br =_c .NewCT_TextLineBreak ();_ccf ._fg .EG_TextRun =append (_ccf ._fg .EG_TextRun ,_egg );};

// SetAlign controls the paragraph alignment
func (_gcd ParagraphProperties )SetAlign (a _c .ST_TextAlignType ){_gcd ._ge .AlgnAttr =a };

// RunProperties controls the run properties.
type RunProperties struct{_bad *_c .CT_TextCharacterProperties ;};

// MakeParagraph constructs a new paragraph wrapper.
func MakeParagraph (x *_c .CT_TextParagraph )Paragraph {return Paragraph {x }};

// X returns the inner wrapped XML type.
func (_gc ParagraphProperties )X ()*_c .CT_TextParagraphProperties {return _gc ._ge };func (_dc ShapeProperties )SetNoFill (){_dc .clearFill ();_dc ._df .NoFill =_c .NewCT_NoFillProperties ()};

// SetSize sets the font size of the run text
func (_cf RunProperties )SetSize (sz _cc .Distance ){_cf ._bad .SzAttr =_d .Int32 (int32 (sz /_cc .HundredthPoint ));};

// SetSolidFill controls the text color of a run.
func (_dag RunProperties )SetSolidFill (c _a .Color ){_dag ._bad .NoFill =nil ;_dag ._bad .BlipFill =nil ;_dag ._bad .GradFill =nil ;_dag ._bad .GrpFill =nil ;_dag ._bad .PattFill =nil ;_dag ._bad .SolidFill =_c .NewCT_SolidColorFillProperties ();_dag ._bad .SolidFill .SrgbClr =_c .NewCT_SRgbColor ();_dag ._bad .SolidFill .SrgbClr .ValAttr =*c .AsRGBString ();};