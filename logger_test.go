package sugaredgorm_test

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/erhanakp/sugaredgorm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

func TestLogMode(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.InfoLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(sugerLogger, config)

	l := logger.LogMode(gormlogger.Info)

	if l != logger {
		t.Errorf("Expected %v, got %v", logger, l)
	}
}

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.InfoLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(sugerLogger, config)

	expectedMsg := "Test info message"
	logger.Info(context.Background(), expectedMsg)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "info	Test info message") {
		t.Errorf("Expected %s, got %s", expectedMsg, capturedOutput)
	}
}

func TestLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.WarnLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(sugerLogger, config)

	expectedMsg := "Test warn message"
	logger.Warn(context.Background(), expectedMsg)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "warn	Test warn message") {
		t.Errorf("Expected %s, got %s", expectedMsg, capturedOutput)
	}
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.ErrorLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(sugerLogger, config)

	expectedMsg := "Test error message"
	logger.Error(context.Background(), expectedMsg)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "error	Test error message") {
		t.Errorf("Expected %s, got %s", expectedMsg, capturedOutput)
	}
}

func TestLogger_TraceNil(t *testing.T) {
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(nil, config)

	// Test Trace method
	begin := time.Now()
	fakeFunc := func() (string, int64) {
		return "SELECT * FROM table", 5
	}
	logger.Trace(context.Background(), begin, fakeFunc, nil)
}

func TestLogger_TraceError(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.ErrorLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(sugerLogger, config)

	// Test Trace method
	begin := time.Now()
	fakeFunc := func() (string, int64) {
		return "SELECT * FROM table", 5
	}
	err := errors.New("test error")
	logger.Trace(context.Background(), begin, fakeFunc, err)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "test error") {
		t.Errorf("Expected %s, got %s", "test error", capturedOutput)
	}

	if !strings.Contains(capturedOutput, "[rows:5] SELECT * FROM table") {
		t.Errorf("Expected %s, got %s", " [rows:5] SELECT * FROM table", capturedOutput)
	}
}

func TestLogger_TraceErrorNoRow(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.ErrorLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(sugerLogger, config)

	// Test Trace method
	begin := time.Now()
	fakeFunc := func() (string, int64) {
		return "SELECT * FROM table", -1
	}
	err := errors.New("test error")
	logger.Trace(context.Background(), begin, fakeFunc, err)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "test error") {
		t.Errorf("Expected %s, got %s", "test error", capturedOutput)
	}

	if !strings.Contains(capturedOutput, "[rows:-] SELECT * FROM table") {
		t.Errorf("Expected %s, got %s", " [rows:-] SELECT * FROM table", capturedOutput)
	}
}

func TestLogger_TraceWarn(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.WarnLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{
		SlowThreshold: time.Nanosecond,
	}
	logger := sugaredgorm.New(sugerLogger, config)

	// Test Trace method
	begin := time.Now()
	fakeFunc := func() (string, int64) {
		time.Sleep(1 * time.Microsecond)
		return "SELECT * FROM table", 5
	}
	logger.Trace(context.Background(), begin, fakeFunc, nil)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "SLOW SQL >= 1ns") {
		t.Errorf("Expected %s, got %s", " SLOW SQL >= 1ns", capturedOutput)
	}

	if !strings.Contains(capturedOutput, "[rows:5] SELECT * FROM table") {
		t.Errorf("Expected %s, got %s", " [rows:5] SELECT * FROM table", capturedOutput)
	}
}

func TestLogger_TraceWarnNoRow(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.WarnLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{
		SlowThreshold: time.Nanosecond,
	}
	logger := sugaredgorm.New(sugerLogger, config)

	// Test Trace method
	begin := time.Now()
	fakeFunc := func() (string, int64) {
		time.Sleep(1 * time.Microsecond)
		return "SELECT * FROM table", -1
	}
	logger.Trace(context.Background(), begin, fakeFunc, nil)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "SLOW SQL >= 1ns") {
		t.Errorf("Expected %s, got %s", " SLOW SQL >= 1ns", capturedOutput)
	}

	if !strings.Contains(capturedOutput, "[rows:-] SELECT * FROM table") {
		t.Errorf("Expected %s, got %s", " [rows:-] SELECT * FROM table", capturedOutput)
	}
}

func TestLogger_TraceInfo(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.InfoLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(sugerLogger, config)

	// Test Trace method
	begin := time.Now()
	fakeFunc := func() (string, int64) {
		return "SELECT * FROM table", 5
	}
	logger.Trace(context.Background(), begin, fakeFunc, nil)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "[rows:5] SELECT * FROM table") {
		t.Errorf("Expected %s, got %s", "[rows:5] SELECT * FROM table", capturedOutput)
	}
}

func TestLogger_TraceInfoNoRow(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.InfoLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{}
	logger := sugaredgorm.New(sugerLogger, config)

	// Test Trace method
	begin := time.Now()
	fakeFunc := func() (string, int64) {
		return "SELECT * FROM table", -1
	}
	logger.Trace(context.Background(), begin, fakeFunc, nil)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, "[rows:-] SELECT * FROM table") {
		t.Errorf("Expected %s, got %s", " [rows:-] SELECT * FROM table", capturedOutput)
	}
}

func TestLogger_TraceWithColorfulError(t *testing.T) {
	var buf bytes.Buffer
	sink := zapcore.AddSync(&buf)
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		sink,
		zapcore.ErrorLevel,
	)

	mockLogger := zap.New(core)
	sugerLogger := mockLogger.Sugar()
	config := sugaredgorm.Config{
		Colorful: true,
	}
	logger := sugaredgorm.New(sugerLogger, config)

	// Test Trace method
	begin := time.Now()
	fakeFunc := func() (string, int64) {
		return "SELECT * FROM table", 5
	}
	err := errors.New("test error")
	logger.Trace(context.Background(), begin, fakeFunc, err)

	capturedOutput := buf.String()
	if !strings.Contains(capturedOutput, sugaredgorm.MagentaBold+"test error") {
		t.Errorf("Expected %s, got %s", sugaredgorm.MagentaBold+"test error", capturedOutput)
	}
}
