package url

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/semicolon-ina/semicolon-url-shortener/repo/common/config"
	"github.com/semicolon-ina/semicolon-url-shortener/repo/infra/inmem"
)

type URLService struct {
	cfg    config.URLConfig
	redisM inmem.RedisItf
}

type URLInterface interface {
	ShortenURL(ctx context.Context, url string) (string, error)
	GetOriginalURL(ctx context.Context, code string) (URLModel, error)
}

// VARIABLES
const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length   = uint64(len(alphabet))
)

func New(uc config.URLConfig, redis inmem.RedisItf) URLInterface {
	return URLService{
		cfg:    uc,
		redisM: redis,
	}
}

func (s URLService) ShortenURL(ctx context.Context, url string) (string, error) {
	// 1. Validasi Input
	if url == "" {
		return "", errors.New("URL cannot be empty")
	}

	// Auto-add protocol kalau user lupa (biar redirect lancar)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	// 2. Generate ID (Logic Base62)
	// Kita pake Nano Time + Random Number biar unik & collision-free
	// Ini trik MVP biar ga perlu cek database dulu (Optimistic strategy)
	seed := uint64(time.Now().UnixNano()) + uint64(rand.Intn(1000))
	code := base62Encode(seed)

	// Potong jadi max 7 karakter biar pendek (optional)
	if len(code) > 7 {
		code = code[len(code)-7:]
	}

  saveToRedis := URLModel{
		Code:        code,
		OriginalURL: url,
		ShortedAt:   time.Now().Unix(),
		ExpiresAt:   time.Now().Add(24 * time.Hour).Unix(), // Contoh: Expire dalam 24 jam
	}

	// 3. Simpan ke Redis (Lewat Repo)
	// Default TTL bisa ambil dari config kalau mau, disini kita hardcode 24 jam dulu
	err := s.redisM.Set(ctx, code, saveToRedis, 24*time.Hour)
	if err != nil {
		return "", err
	}

	// 4. Return Full Short URL
	// Gabungin BaseURL dari config + Code
	// Contoh: "http://localhost:8080" + "/" + "AbC"
	// Pastiin BaseURL ga ada slash di belakang, atau handle manual
	baseURL := strings.TrimRight(s.cfg.BaseURL, "/")
	fullShortURL := baseURL + "/" + code

	return fullShortURL, nil
}

func (s URLService) GetOriginalURL(ctx context.Context, code string) (URLModel, error) {
	result := URLModel{}

	urlString, err := s.redisM.Get(ctx, code)
	if err != nil {
		return URLModel{}, err
	}

	if urlString == "" {
		return URLModel{}, errors.New("URL not found")
	}

	err = json.Unmarshal([]byte(urlString), &result)
	if err != nil {
		return URLModel{}, err
	}

	return result, err
}

// --- Helper: Base62 Encoder (Math Logic) ---
func base62Encode(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}

	chars := make([]byte, 0)
	for number > 0 {
		remainder := number % length
		chars = append(chars, alphabet[remainder])
		number = number / length
	}

	// Reverse string
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}
