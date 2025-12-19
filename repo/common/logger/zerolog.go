package logger

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

// Struct Konkrit implementasi Zerolog
type zeroLogImpl struct {
	z zerolog.Logger
}

// SetupZeroLog: Panggil ini SEKALI di main.go
func SetupZeroLog() {
	// 1. Config Standard
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// 2. Init Zerolog (Output JSON ke Stdout)
	z := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	// 3. Masukin ke Global Engine
	engine = &zeroLogImpl{z: z}
}

// --- IMPLEMENTASI LOGIC ---

func (l *zeroLogImpl) Debug(ctx context.Context, msg string, args ...any) {
	l.enrich(ctx).Debug().Msgf(msg, args...)
}

func (l *zeroLogImpl) Info(ctx context.Context, msg string, args ...any) {
	l.enrich(ctx).Info().Msgf(msg, args...)
}

func (l *zeroLogImpl) Warn(ctx context.Context, msg string, args ...any) {
	l.enrich(ctx).Warn().Msgf(msg, args...)
}

func (l *zeroLogImpl) Error(ctx context.Context, err error, msg string, args ...any) {
	// Stack() penting buat New Relic Error Tracking
	l.enrich(ctx).Error().Stack().Err(err).Msgf(msg, args...)
}

// Helper buat ngambil Trace ID dari Context (PENTING BUAT NEW RELIC)
func (l *zeroLogImpl) enrich(ctx context.Context) *zerolog.Logger {
	// Sesuaikan Key string ini dengan Middleware lo.
	// Biasanya "trace_id", "request_id", atau "X-Request-ID"
	if traceID, ok := ctx.Value("trace_id").(string); ok {
		// Log ini otomatis nambahin field "trace_id": "xyz-123" di JSON-nya
		enriched := l.z.With().Str("trace_id", traceID).Logger()
		return &enriched
	}
	return &l.z
}
