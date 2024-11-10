package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization token is missing"})
	}

	// Token bisa ditambahkan dengan verifikasi JWT di sini
	// Misalnya: decode JWT dan cek peran pengguna (admin atau user)
	// Jika valid, lanjutkan ke rute berikutnya
	if !isValidToken(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	return c.Next()
}

// Fungsi validasi token
func isValidToken(token string) bool {
	// Logika untuk memverifikasi token JWT
	return strings.HasPrefix(token, "Bearer ")
}
