package routes

import (
	"be-stepup/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Route untuk mendapatkan semua produk
	app.Get("/products", controllers.GetAllProducts)

	// Route untuk membuat produk baru
	app.Post("/products", controllers.CreateProduct)

	// Route untuk memperbarui produk berdasarkan ID
	app.Put("/products/:id", controllers.UpdateProduct)

	// Route untuk mengunggah gambar produk
	app.Post("/products/upload", controllers.UploadProductImage)

	// Route untuk melayani gambar yang diunggah (memungkinkan mengakses gambar melalui URL)
	app.Static("/images", "./be-stepup/uploads")

	// Routes untuk pemesanan
	app.Post("/api/orders", controllers.CreateOrder) // Buat pemesanan

	// Routes untuk status pengiriman
	app.Put("/api/orders/:order_id/shipping", controllers.UpdateShippingStatus) // Update status pengiriman

	app.Post("/api/users", controllers.CreateUser) // Menambah user baru
}
