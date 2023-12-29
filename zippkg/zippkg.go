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

package zippkg ;import (_a "archive/zip";_g "bytes";_ef "encoding/xml";_fb "fmt";_dfd "github.com/unidoc/unioffice";_dd "github.com/unidoc/unioffice/algo";_ad "github.com/unidoc/unioffice/common/tempstorage";_ae "github.com/unidoc/unioffice/schema/soo/pkg/relationships";
_c "io";_fa "path";_df "path/filepath";_d "sort";_ac "strings";_f "time";);

// AddTarget allows documents to register decode targets. Path is a path that
// will be found in the zip file and ifc is an XML element that the file will be
// unmarshaled to.  filePath is the absolute path to the target, ifc is the
// object to decode into, sourceFileType is the type of file that the reference
// was discovered in, and index is the index of the source file type.
func (_ab *DecodeMap )AddTarget (filePath string ,ifc interface{},sourceFileType string ,idx uint32 )bool {if _ab ._eba ==nil {_ab ._eba =make (map[string ]Target );_ab ._dfdc =make (map[*_ae .Relationships ]string );_ab ._ca =make (map[string ]struct{});
_ab ._cb =make (map[string ]int );};if _fa .IsAbs (filePath ){filePath =_ac .TrimPrefix (filePath ,"\u002f");};_cc :=_fa .Clean (filePath );if _ ,_af :=_ab ._ca [_cc ];_af {return false ;};_ab ._ca [_cc ]=struct{}{};_ab ._eba [_cc ]=Target {Path :_cc ,Typ :sourceFileType ,Ifc :ifc ,Index :idx };
return true ;};func (_bg *DecodeMap )IndexFor (path string )int {return _bg ._cb [path ]};

// SelfClosingWriter wraps a writer and replaces XML tags of the
// type <foo></foo> with <foo/>
type SelfClosingWriter struct{W _c .Writer ;};

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes (z *_a .Writer ,zipPath string ,data []byte )error {_dfg ,_eea :=z .Create (zipPath );if _eea !=nil {return _fb .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_eea );
};_ ,_eea =_c .Copy (_dfg ,_g .NewReader (data ));return _eea ;};

// SetOnNewRelationshipFunc sets the function to be called when a new
// relationship has been discovered.
func (_efd *DecodeMap )SetOnNewRelationshipFunc (fn OnNewRelationshipFunc ){_efd ._fba =fn };const XMLHeader ="\u003c\u003f\u0078\u006d\u006c\u0020\u0076e\u0072\u0073\u0069o\u006e\u003d\u00221\u002e\u0030\"\u0020\u0065\u006e\u0063\u006f\u0064i\u006eg=\u0022\u0055\u0054\u0046\u002d\u0038\u0022\u0020\u0073\u0074\u0061\u006e\u0064\u0061\u006c\u006f\u006e\u0065\u003d\u0022\u0079\u0065\u0073\u0022\u003f\u003e"+"\u000a";


// OnNewRelationshipFunc is called when a new relationship has been discovered.
//
// target is a resolved path that takes into account the location of the
// relationships file source and should be the path in the zip file.
//
// files are passed so non-XML files that can't be handled by AddTarget can be
// decoded directly (e.g. images)
//
// rel is the actual relationship so its target can be modified if the source
// target doesn't match where unioffice will write the file (e.g. read in
// 'xl/worksheets/MyWorksheet.xml' and we'll write out
// 'xl/worksheets/sheet1.xml')
type OnNewRelationshipFunc func (_fbc *DecodeMap ,_gc ,_eb string ,_b []*_a .File ,_ag *_ae .Relationship ,_da Target )error ;

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp (f *_a .File ,path string )(string ,error ){_cf ,_gbc :=_ad .TempFile (path ,"\u007a\u007a");if _gbc !=nil {return "",_gbc ;};defer _cf .Close ();_ade ,_gbc :=f .Open ();if _gbc !=nil {return "",_gbc ;};defer _ade .Close ();_ ,_gbc =_c .Copy (_cf ,_ade );
if _gbc !=nil {return "",_gbc ;};return _cf .Name (),nil ;};

// Decode loops decoding targets registered with AddTarget and calling th
func (_efb *DecodeMap )Decode (files []*_a .File )error {_efbb :=1;for _efbb > 0{for len (_efb ._de )> 0{_add :=_efb ._de [0];_efb ._de =_efb ._de [1:];_bd :=_add .Ifc .(*_ae .Relationships );for _ ,_db :=range _bd .Relationship {_efdb :=_efb ._dfdc [_bd ];
if _df .IsAbs (_db .TargetAttr ){_db .TargetAttr =_ac .TrimPrefix (_db .TargetAttr ,"\u002f");if _ac .HasPrefix (_db .TargetAttr ,_efdb ){_efdb ="";};};_efb ._fba (_efb ,_efdb +_db .TargetAttr ,_db .TypeAttr ,files ,_db ,_add );};};for _gf ,_efa :=range files {if _efa ==nil {continue ;
};if _dad ,_ff :=_efb ._eba [_efa .Name ];_ff {delete (_efb ._eba ,_efa .Name );if _ea :=Decode (_efa ,_dad .Ifc );_ea !=nil {return _ea ;};files [_gf ]=nil ;if _ee ,_ec :=_dad .Ifc .(*_ae .Relationships );_ec {_efb ._de =append (_efb ._de ,_dad );_bde ,_ :=_fa .Split (_fa .Clean (_efa .Name +"\u002f\u002e\u002e\u002f"));
_efb ._dfdc [_ee ]=_bde ;_efbb ++;};};};_efbb --;};return nil ;};var _bge =[]byte {'\r','\n'};type Target struct{Path string ;Typ string ;Ifc interface{};Index uint32 ;};func (_eg *DecodeMap )RecordIndex (path string ,idx int ){_eg ._cb [path ]=idx };var _dbg =[]byte {'/','>'};


// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode (f *_a .File ,dest interface{})error {_bgd ,_ba :=f .Open ();if _ba !=nil {return _fb .Errorf ("e\u0072r\u006f\u0072\u0020\u0072\u0065\u0061\u0064\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",f .Name ,_ba );};defer _bgd .Close ();_gdb :=_ef .NewDecoder (_bgd );
if _ffd :=_gdb .Decode (dest );_ffd !=nil {return _fb .Errorf ("e\u0072\u0072\u006f\u0072 d\u0065c\u006f\u0064\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",f .Name ,_ffd );};if _bf ,_bgf :=dest .(*_ae .Relationships );_bgf {for _gb ,_gfg :=range _bf .Relationship {switch _gfg .TypeAttr {case _dfd .OfficeDocumentTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .OfficeDocumentType ;
case _dfd .StylesTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .StylesType ;case _dfd .ThemeTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .ThemeType ;case _dfd .ControlTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .ControlType ;case _dfd .SettingsTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .SettingsType ;
case _dfd .ImageTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .ImageType ;case _dfd .CommentsTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .CommentsType ;case _dfd .ThumbnailTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .ThumbnailType ;
case _dfd .DrawingTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .DrawingType ;case _dfd .ChartTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .ChartType ;case _dfd .ExtendedPropertiesTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .ExtendedPropertiesType ;
case _dfd .CustomXMLTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .CustomXMLType ;case _dfd .WorksheetTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .WorksheetType ;case _dfd .SharedStringsTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .SharedStringsType ;
case _dfd .TableTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .TableType ;case _dfd .HeaderTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .HeaderType ;case _dfd .FooterTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .FooterType ;case _dfd .NumberingTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .NumberingType ;
case _dfd .FontTableTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .FontTableType ;case _dfd .WebSettingsTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .WebSettingsType ;case _dfd .FootNotesTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .FootNotesType ;
case _dfd .EndNotesTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .EndNotesType ;case _dfd .SlideTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .SlideType ;case _dfd .VMLDrawingTypeStrict :_bf .Relationship [_gb ].TypeAttr =_dfd .VMLDrawingType ;
};};_d .Slice (_bf .Relationship ,func (_bgc ,_dag int )bool {_gdg :=_bf .Relationship [_bgc ];_gdd :=_bf .Relationship [_dag ];return _dd .NaturalLess (_gdg .IdAttr ,_gdd .IdAttr );});};return nil ;};func MarshalXMLByTypeIndex (z *_a .Writer ,dt _dfd .DocType ,typ string ,idx int ,v interface{})error {_cbd :=_dfd .AbsoluteFilename (dt ,typ ,idx );
return MarshalXML (z ,_cbd ,v );};func MarshalXMLByType (z *_a .Writer ,dt _dfd .DocType ,typ string ,v interface{})error {_dcd :=_dfd .AbsoluteFilename (dt ,typ ,0);return MarshalXML (z ,_dcd ,v );};

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor (path string )string {_gg :=_ac .Split (path ,"\u002f");_gd :=_ac .Join (_gg [0:len (_gg )-1],"\u002f");_ed :=_gg [len (_gg )-1];_gd +="\u002f_\u0072\u0065\u006c\u0073\u002f";_ed +="\u002e\u0072\u0065l\u0073";return _gd +_ed ;};

// AddFileFromDisk reads a file from internal storage and adds it at a given path to a zip file.
// TODO: Rename to AddFileFromStorage in next major version release (v2).
// NOTE: If disk storage cannot be used, memory storage can be used instead by calling memstore.SetAsStorage().
func AddFileFromDisk (z *_a .Writer ,zipPath ,storagePath string )error {_ddf ,_dc :=z .Create (zipPath );if _dc !=nil {return _fb .Errorf ("e\u0072\u0072\u006f\u0072 c\u0072e\u0061\u0074\u0069\u006e\u0067 \u0025\u0073\u003a\u0020\u0025\u0073",zipPath ,_dc );
};_ga ,_dc :=_ad .Open (storagePath );if _dc !=nil {return _fb .Errorf ("e\u0072r\u006f\u0072\u0020\u006f\u0070\u0065\u006e\u0069n\u0067\u0020\u0025\u0073: \u0025\u0073",storagePath ,_dc );};defer _ga .Close ();_ ,_dc =_c .Copy (_ddf ,_ga );return _dc ;
};

// DecodeMap is used to walk a tree of relationships, decoding files and passing
// control back to the document.
type DecodeMap struct{_eba map[string ]Target ;_dfdc map[*_ae .Relationships ]string ;_de []Target ;_fba OnNewRelationshipFunc ;_ca map[string ]struct{};_cb map[string ]int ;};func (_bc SelfClosingWriter )Write (b []byte )(int ,error ){_cde :=0;_cag :=0;
for _aga :=0;_aga < len (b )-2;_aga ++{if b [_aga ]=='>'&&b [_aga +1]=='<'&&b [_aga +2]=='/'{_bac :=[]byte {};_acf :=_aga ;for _cce :=_aga ;_cce >=0;_cce --{if b [_cce ]==' '{_acf =_cce ;}else if b [_cce ]=='<'{_bac =b [_cce +1:_acf ];break ;};};_aeg :=[]byte {};
for _aee :=_aga +3;_aee < len (b );_aee ++{if b [_aee ]=='>'{_aeg =b [_aga +3:_aee ];break ;};};if !_g .Equal (_bac ,_aeg ){continue ;};_gdbg ,_cg :=_bc .W .Write (b [_cde :_aga ]);if _cg !=nil {return _cag +_gdbg ,_cg ;};_cag +=_gdbg ;_ ,_cg =_bc .W .Write (_dbg );
if _cg !=nil {return _cag ,_cg ;};_cag +=3;for _efe :=_aga +2;_efe < len (b )&&b [_efe ]!='>';_efe ++{_cag ++;_cde =_efe +2;_aga =_cde ;};};};_fd ,_gbf :=_bc .W .Write (b [_cde :]);return _fd +_cag ,_gbf ;};

// MarshalXML creates a file inside of a zip and marshals an object as xml, prefixing it
// with a standard XML header.
func MarshalXML (z *_a .Writer ,filename string ,v interface{})error {_edd :=&_a .FileHeader {};_edd .Method =_a .Deflate ;_edd .Name =filename ;_edd .SetModTime (_f .Now ());_gdc ,_gcc :=z .CreateHeader (_edd );if _gcc !=nil {return _fb .Errorf ("\u0063\u0072\u0065\u0061ti\u006e\u0067\u0020\u0025\u0073\u0020\u0069\u006e\u0020\u007a\u0069\u0070\u003a\u0020%\u0073",filename ,_gcc );
};_ ,_gcc =_gdc .Write ([]byte (XMLHeader ));if _gcc !=nil {return _fb .Errorf ("\u0063\u0072e\u0061\u0074\u0069\u006e\u0067\u0020\u0078\u006d\u006c\u0020\u0068\u0065\u0061\u0064\u0065\u0072\u0020\u0074\u006f\u0020\u0025\u0073: \u0025\u0073",filename ,_gcc );
};if _gcc =_ef .NewEncoder (SelfClosingWriter {_gdc }).Encode (v );_gcc !=nil {return _fb .Errorf ("\u006d\u0061\u0072\u0073\u0068\u0061\u006c\u0069\u006e\u0067\u0020\u0025s\u003a\u0020\u0025\u0073",filename ,_gcc );};_ ,_gcc =_gdc .Write (_bge );return _gcc ;
};