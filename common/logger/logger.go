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

package logger ;import (_f "fmt";_c "io";_fc "os";_b "path/filepath";_ff "runtime";);

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _c .Writer ;};func (_cf ConsoleLogger )output (_ad _c .Writer ,_dbf string ,_bc string ,_aabc ...interface{}){_gag (_ad ,_dbf ,_bc ,_aabc ...);};

// Warning logs warning message.
func (_ae ConsoleLogger )Warning (format string ,args ...interface{}){if _ae .LogLevel >=LogLevelWarning {_ac :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ae .output (_fc .Stdout ,_ac ,format ,args ...);};};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// Info logs info message.
func (_ba ConsoleLogger )Info (format string ,args ...interface{}){if _ba .LogLevel >=LogLevelInfo {_bae :="\u005bI\u004e\u0046\u004f\u005d\u0020";_ba .output (_fc .Stdout ,_bae ,format ,args ...);};};func _gag (_cef _c .Writer ,_ag string ,_bb string ,_cag ...interface{}){_ ,_ccf ,_abg ,_aga :=_ff .Caller (3);
if !_aga {_ccf ="\u003f\u003f\u003f";_abg =0;}else {_ccf =_b .Base (_ccf );};_abga :=_f .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_ag ,_ccf ,_abg )+_bb +"\u000a";_f .Fprintf (_cef ,_abga ,_cag ...);};

// Notice logs notice message.
func (_gg ConsoleLogger )Notice (format string ,args ...interface{}){if _gg .LogLevel >=LogLevelNotice {_gc :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_gg .output (_fc .Stdout ,_gc ,format ,args ...);};};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };var Log Logger =DummyLogger {};

// DummyLogger does nothing.
type DummyLogger struct{};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);

// Trace logs trace message.
func (_aab ConsoleLogger )Trace (format string ,args ...interface{}){if _aab .LogLevel >=LogLevelTrace {_gac :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_aab .output (_fc .Stdout ,_gac ,format ,args ...);};};

// Debug logs debug message.
func (_aa ConsoleLogger )Debug (format string ,args ...interface{}){if _aa .LogLevel >=LogLevelDebug {_bd :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_aa .output (_fc .Stdout ,_bd ,format ,args ...);};};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// Notice logs notice message.
func (_eec WriterLogger )Notice (format string ,args ...interface{}){if _eec .LogLevel >=LogLevelNotice {_bf :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_eec .logToWriter (_eec .Output ,_bf ,format ,args ...);};};

// Trace logs trace message.
func (_ffg WriterLogger )Trace (format string ,args ...interface{}){if _ffg .LogLevel >=LogLevelTrace {_ea :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_ffg .logToWriter (_ffg .Output ,_ea ,format ,args ...);};};

// Error logs error message.
func (_ccc ConsoleLogger )Error (format string ,args ...interface{}){if _ccc .LogLevel >=LogLevelError {_ec :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_ccc .output (_fc .Stdout ,_ec ,format ,args ...);};};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ga ConsoleLogger )IsLogLevel (level LogLevel )bool {return _ga .LogLevel >=level };

// Error logs error message.
func (_gce WriterLogger )Error (format string ,args ...interface{}){if _gce .LogLevel >=LogLevelError {_gef :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_gce .logToWriter (_gce .Output ,_gef ,format ,args ...);};};func (_cb WriterLogger )logToWriter (_fce _c .Writer ,_gge string ,_de string ,_gaf ...interface{}){_gag (_fce ,_gge ,_de ,_gaf );
};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_ge WriterLogger )IsLogLevel (level LogLevel )bool {return _ge .LogLevel >=level };

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_be string ,_e ...interface{});Warning (_a string ,_ee ...interface{});Notice (_eb string ,_ab ...interface{});Info (_cc string ,_dg ...interface{});Debug (_ce string ,_db ...interface{});Trace (_ef string ,_cg ...interface{});
IsLogLevel (_g LogLevel )bool ;};

// Debug logs debug message.
func (_ebc WriterLogger )Debug (format string ,args ...interface{}){if _ebc .LogLevel >=LogLevelDebug {_ed :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_ebc .logToWriter (_ebc .Output ,_ed ,format ,args ...);};};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// Warning logs warning message.
func (_ca WriterLogger )Warning (format string ,args ...interface{}){if _ca .LogLevel >=LogLevelWarning {_aaf :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_ca .logToWriter (_ca .Output ,_aaf ,format ,args ...);};};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _c .Writer )*WriterLogger {logger :=WriterLogger {Output :writer ,LogLevel :logLevel };return &logger ;};

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// Info logs info message.
func (_fg WriterLogger )Info (format string ,args ...interface{}){if _fg .LogLevel >=LogLevelInfo {_bea :="\u005bI\u004e\u0046\u004f\u005d\u0020";_fg .logToWriter (_fg .Output ,_bea ,format ,args ...);};};