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

package logger ;import (_b "fmt";_de "io";_a "os";_d "path/filepath";_c "runtime";);

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_db WriterLogger )IsLogLevel (level LogLevel )bool {return _db .LogLevel >=level };

// Info does nothing for dummy logger.
func (DummyLogger )Info (format string ,args ...interface{}){};

// IsLogLevel returns true if log level is greater or equal than `level`.
// Can be used to avoid resource intensive calls to loggers.
func (_be ConsoleLogger )IsLogLevel (level LogLevel )bool {return _be .LogLevel >=level };

// Debug logs debug message.
func (_ca ConsoleLogger )Debug (format string ,args ...interface{}){if _ca .LogLevel >=LogLevelDebug {_ge :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_ca .output (_a .Stdout ,_ge ,format ,args ...);};};

// ConsoleLogger is a logger that writes logs to the 'os.Stdout'
type ConsoleLogger struct{LogLevel LogLevel ;};

// LogLevel is the verbosity level for logging.
type LogLevel int ;

// Trace does nothing for dummy logger.
func (DummyLogger )Trace (format string ,args ...interface{}){};

// SetLogger sets 'logger' to be used by the unidoc unipdf library.
func SetLogger (logger Logger ){Log =logger };

// Logger is the interface used for logging in the unipdf package.
type Logger interface{Error (_cd string ,_dd ...interface{});Warning (_ab string ,_ea ...interface{});Notice (_cf string ,_df ...interface{});Info (_f string ,_abd ...interface{});Debug (_ff string ,_ac ...interface{});Trace (_g string ,_fg ...interface{});IsLogLevel (_gc LogLevel )bool ;};

// Notice does nothing for dummy logger.
func (DummyLogger )Notice (format string ,args ...interface{}){};

// Warning logs warning message.
func (_adb WriterLogger )Warning (format string ,args ...interface{}){if _adb .LogLevel >=LogLevelWarning {_ec :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_adb .logToWriter (_adb .Output ,_ec ,format ,args ...);};};

// Error logs error message.
func (_gd ConsoleLogger )Error (format string ,args ...interface{}){if _gd .LogLevel >=LogLevelError {_da :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_gd .output (_a .Stdout ,_da ,format ,args ...);};};

// NewWriterLogger creates new 'writer' logger.
func NewWriterLogger (logLevel LogLevel ,writer _de .Writer )*WriterLogger {logger :=WriterLogger {Output :writer ,LogLevel :logLevel };return &logger ;};

// Debug logs debug message.
func (_dbc WriterLogger )Debug (format string ,args ...interface{}){if _dbc .LogLevel >=LogLevelDebug {_ade :="\u005b\u0044\u0045\u0042\u0055\u0047\u005d\u0020";_dbc .logToWriter (_dbc .Output ,_ade ,format ,args ...);};};

// Trace logs trace message.
func (_gf ConsoleLogger )Trace (format string ,args ...interface{}){if _gf .LogLevel >=LogLevelTrace {_bb :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_gf .output (_a .Stdout ,_bb ,format ,args ...);};};

// Info logs info message.
func (_gdg ConsoleLogger )Info (format string ,args ...interface{}){if _gdg .LogLevel >=LogLevelInfo {_ad :="\u005bI\u004e\u0046\u004f\u005d\u0020";_gdg .output (_a .Stdout ,_ad ,format ,args ...);};};

// Error logs error message.
func (_aba WriterLogger )Error (format string ,args ...interface{}){if _aba .LogLevel >=LogLevelError {_bc :="\u005b\u0045\u0052\u0052\u004f\u0052\u005d\u0020";_aba .logToWriter (_aba .Output ,_bc ,format ,args ...);};};func (_ded ConsoleLogger )output (_fb _de .Writer ,_fbc string ,_af string ,_eg ...interface{}){_acg (_fb ,_fbc ,_af ,_eg ...);};

// Notice logs notice message.
func (_fe ConsoleLogger )Notice (format string ,args ...interface{}){if _fe .LogLevel >=LogLevelNotice {_bg :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_fe .output (_a .Stdout ,_bg ,format ,args ...);};};

// Debug does nothing for dummy logger.
func (DummyLogger )Debug (format string ,args ...interface{}){};

// Warning does nothing for dummy logger.
func (DummyLogger )Warning (format string ,args ...interface{}){};

// DummyLogger does nothing.
type DummyLogger struct{};

// Notice logs notice message.
func (_cc WriterLogger )Notice (format string ,args ...interface{}){if _cc .LogLevel >=LogLevelNotice {_dg :="\u005bN\u004f\u0054\u0049\u0043\u0045\u005d ";_cc .logToWriter (_cc .Output ,_dg ,format ,args ...);};};

// IsLogLevel returns true from dummy logger.
func (DummyLogger )IsLogLevel (level LogLevel )bool {return true };func (_dc WriterLogger )logToWriter (_cdb _de .Writer ,_ef string ,_bcd string ,_fc ...interface{}){_acg (_cdb ,_ef ,_bcd ,_fc );};

// Info logs info message.
func (_gfa WriterLogger )Info (format string ,args ...interface{}){if _gfa .LogLevel >=LogLevelInfo {_egf :="\u005bI\u004e\u0046\u004f\u005d\u0020";_gfa .logToWriter (_gfa .Output ,_egf ,format ,args ...);};};func _acg (_cba _de .Writer ,_cdf string ,_ba string ,_ee ...interface{}){_ ,_cgf ,_dbf ,_eaf :=_c .Caller (3);if !_eaf {_cgf ="\u003f\u003f\u003f";_dbf =0;}else {_cgf =_d .Base (_cgf );};_dga :=_b .Sprintf ("\u0025s\u0020\u0025\u0073\u003a\u0025\u0064 ",_cdf ,_cgf ,_dbf )+_ba +"\u000a";_b .Fprintf (_cba ,_dga ,_ee ...);};

// Trace logs trace message.
func (_afb WriterLogger )Trace (format string ,args ...interface{}){if _afb .LogLevel >=LogLevelTrace {_cb :="\u005b\u0054\u0052\u0041\u0043\u0045\u005d\u0020";_afb .logToWriter (_afb .Output ,_cb ,format ,args ...);};};

// Error does nothing for dummy logger.
func (DummyLogger )Error (format string ,args ...interface{}){};

// WriterLogger is the logger that writes data to the Output writer
type WriterLogger struct{LogLevel LogLevel ;Output _de .Writer ;};

// Warning logs warning message.
func (_cg ConsoleLogger )Warning (format string ,args ...interface{}){if _cg .LogLevel >=LogLevelWarning {_ffd :="\u005b\u0057\u0041\u0052\u004e\u0049\u004e\u0047\u005d\u0020";_cg .output (_a .Stdout ,_ffd ,format ,args ...);};};

// NewConsoleLogger creates new console logger.
func NewConsoleLogger (logLevel LogLevel )*ConsoleLogger {return &ConsoleLogger {LogLevel :logLevel }};const (LogLevelTrace LogLevel =5;LogLevelDebug LogLevel =4;LogLevelInfo LogLevel =3;LogLevelNotice LogLevel =2;LogLevelWarning LogLevel =1;LogLevelError LogLevel =0;);var Log Logger =DummyLogger {};