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

package logger ;import (_b "fmt";_be "io";_ad "os";_dg "path/filepath";_d "runtime";);

// Trace logs trace message.
func (_fb WriterLogger )Trace (format string ,args ...interface{}){if _fb .LogLevel >=LogLevelTrace {_ccg :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_fb .logToWriter (_fb .Output ,_ccg ,format ,args ...);};};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Trace logs trace message.
func (_ddf ConsoleLogger )Trace (format string ,args ...interface{}){if _ddf .LogLevel >=LogLevelTrace {_gg :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_ddf .output (_ad .Stdout ,_gg ,format ,args ...);};};

// LogLevel is the verbosity level for logging.
type LogLevel int ;func (_eab ConsoleLogger )output (_cdg _be .Writer ,_cg string ,_eb string ,_dab ...interface{}){_cfg (_cdg ,_cg ,_eb ,_dab ...);};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// Info logs info message.
func (_dd ConsoleLogger )Info (format string ,args ...interface{}){if _dd .LogLevel >=LogLevelInfo {_ga :="\u005bI\u004e\u0046\u004f\u005d\u0020";_dd .output (_ad .Stdout ,_ga ,format ,args ...);};};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// Error logs error message.
func (_ea ConsoleLogger )Error (format string ,args ...interface{}){if _ea .LogLevel >=LogLevelError {_gb :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ea .output (_ad .Stdout ,_gb ,format ,args ...);};};

// Notice logs notice message.
func (_bd ConsoleLogger )Notice (format string ,args ...interface{}){if _bd .LogLevel >=LogLevelNotice {_ed :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_bd .output (_ad .Stdout ,_ed ,format ,args ...);};};

// Warning logs warning message.
func (_ebf WriterLogger )Warning (format string ,args ...interface{}){if _ebf .LogLevel >=LogLevelWarning {_fgc :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ebf .logToWriter (_ebf .Output ,_fgc ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_dc WriterLogger )IsLogLevel (level LogLevel )bool {return _dc .LogLevel >=level };var Log Logger =DummyLogger {};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_c string ,_cd ...interface{});Warning (_da string ,_af ...interface{});Notice (_e string ,_bb ...interface{});Info (_cc string ,_f ...interface{});Debug (_cb string ,_eg ...interface{});Trace (_fg string ,_g ...interface{});IsLogLevel (_afg LogLevel )bool ;};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};func (_cdd WriterLogger )logToWriter (_faf _be .Writer ,_bee string ,_ag string ,_dge ...interface{}){_cfg (_faf ,_bee ,_ag ,_dge );};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _be .Writer )*WriterLogger {logger :=WriterLogger {Output :writer ,LogLevel :logLevel };return &logger ;};

// Error logs error message.
func (_gbe WriterLogger )Error (format string ,args ...interface{}){if _gbe .LogLevel >=LogLevelError {_gaf :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_gbe .logToWriter (_gbe .Output ,_gaf ,format ,args ...);};};func _cfg (_beef _be .Writer ,_ae string ,_bef string ,_ade ...interface{}){_ ,_daae ,_ac ,_bea :=_d .Caller (3);if !_bea {_daae ="\u003f\u003f\u003f";_ac =0;}else {_daae =_dg .Base (_daae );};_ge :=_b .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_ae ,_daae ,_ac )+_bef +"\u000a";_b .Fprintf (_beef ,_ge ,_ade ...);};

// Info logs info message.
func (_bbb WriterLogger )Info (format string ,args ...interface{}){if _bbb .LogLevel >=LogLevelInfo {_gba :="\u005bI\u004e\u0046\u004f\u005d\u0020";_bbb .logToWriter (_bbb .Output ,_gba ,format ,args ...);};};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Notice logs notice message.
func (_cf WriterLogger )Notice (format string ,args ...interface{}){if _cf .LogLevel >=LogLevelNotice {_bc :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_cf .logToWriter (_cf .Output ,_bc ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_dgg ConsoleLogger )IsLogLevel (level LogLevel )bool {return _dgg .LogLevel >=level };

// Debug logs debug message.
func (_ee WriterLogger )Debug (format string ,args ...interface{}){if _ee .LogLevel >=LogLevelDebug {_gc :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_ee .logToWriter (_ee .Output ,_gc ,format ,args ...);};};

// Warning logs warning message.
func (_ce ConsoleLogger )Warning (format string ,args ...interface{}){if _ce .LogLevel >=LogLevelWarning {_fa :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ce .output (_ad .Stdout ,_fa ,format ,args ...);};};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// DummyLogger does nothing.
type DummyLogger struct{};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _be .Writer ;};

// Debug logs debug message.
func (_cbg ConsoleLogger )Debug (format string ,args ...interface{}){if _cbg .LogLevel >=LogLevelDebug {_daa :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_cbg .output (_ad .Stdout ,_daa ,format ,args ...);};};