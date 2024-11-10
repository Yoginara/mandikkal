package controllers

import (
	"be-stepup/config"
	"be-stepup/models"
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllProducts fetches all products from the database
func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Product
	collection := config.GetCollection("products")

	// Gunakan variabel untuk context
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching products"})
	}
	defer func() {
		if err := cursor.Close(ctx); err != nil {
			fmt.Println("Error closing cursor:", err)
		}
	}()

	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error decoding product"})
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Cursor error"})
	}

	return c.JSON(products)
}

// CreateProduct creates a new product
func CreateProduct(c *fiber.Ctx) error {
	collection := config.GetCollection("products")
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	_, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}

	return c.JSON(product)
}

// UploadProductImage mengunggah gambar untuk produk dan mengganti nama file
func UploadProductImage(c *fiber.Ctx) error {
	// Mendapatkan file dari form
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to upload image"})
	}

	// Menggunakan UUID untuk nama file yang unik
	uniqueName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))

	// Set path untuk menyimpan file dengan nama yang unik
	filePath := fmt.Sprintf("uploads/%s", uniqueName)

	// Log untuk debugging path dan error
	fmt.Println("File Path:", filePath)

	// Simpan file ke path yang sudah ditentukan
	err = c.SaveFile(file, filePath)
	if err != nil {
		fmt.Println("Save File Error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save image"})
	}

	// Return path dari gambar yang berhasil diunggah
	return c.JSON(fiber.Map{"imagePath": filePath})
}

// UpdateProduct updates an existing product
func UpdateProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	productID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	var productData struct {
		Name        string `json:"name"`
		Brand       string `json:"brand"`
		Category    string `json:"category"`
		Price       int    `json:"price"`
		Stock       int    `json:"stock"`
		Description string `json:"description"`
		Image       string `json:"image"` // Field untuk path gambar
	}

	if err := c.BodyParser(&productData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	collection := config.GetCollection("products")
	filter := bson.M{"_id": productID}
	update := bson.M{
		"$set": bson.M{
			"name":        productData.Name,
			"brand":       productData.Brand,
			"category":    productData.Category,
			"price":       productData.Price,
			"stock":       productData.Stock,
			"description": productData.Description,
			"image":       productData.Image,
		},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
	}

	return c.JSON(fiber.Map{"message": "Product updated successfully"})
}

// DeleteProduct deletes a product from the database
func DeleteProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	productID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	collection := config.GetCollection("products")
	filter := bson.M{"_id": productID}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}
