package logger

import (
	"context"
)

// 1. KONTRAK (Interface)
// Ini bentuk mesin logger kita.
type LoggerEngine interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, err error, msg string, args ...any)
}

// 2. GLOBAL VARIABLE (The Singleton)
// Defaultnya nil, nanti diisi pas Init.
var engine LoggerEngine

// 3. STATIC WRAPPER (Ini yang lo panggil di Service)
// Lo gak perlu inject-inject ke struct service. Cukup import package ini.

func Debug(ctx context.Context, msg string, args ...any) {
	if engine != nil {
		engine.Debug(ctx, msg, args...)
	}
}

func Info(ctx context.Context, msg string, args ...any) {
	if engine != nil {
		engine.Info(ctx, msg, args...)
	}
}

func Warn(ctx context.Context, msg string, args ...any) {
	if engine != nil {
		engine.Warn(ctx, msg, args...)
	}
}

func Error(ctx context.Context, err error, msg string, args ...any) {
	if engine != nil {
		engine.Error(ctx, err, msg, args...)
	}
}

// 4. TEST HOOK (Buat Mocking)
// Panggil ini di func Test lo buat ganti mesin logger jadi Mock.
func SetLogger(l LoggerEngine) {
	engine = l
}