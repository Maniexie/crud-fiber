package controllers

import (
	"log"

	"github.com/Maniexie/crud-fiber/database"
	"github.com/Maniexie/crud-fiber/models/entity"
	"github.com/Maniexie/crud-fiber/models/req"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// func UserControllerShow(c *fiber.Ctx) error {
// 	return c.JSON(fiber.Map{
// 		"pesan": "hello dunia",
// 	})
// }

// Melihat Semua data
func UserControllerShow(c *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(users)
}

// Membuat data baru
func UserControllerCreate(c *fiber.Ctx) error {
	user := new(req.UserReq)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nama / Email tidak sesuai",
			"error":   err.Error(),
			"test":    "saya ingin ngetes kk",
		})
	}

	newUser := entity.User{
		Name:  user.Name,
		Email: user.Email,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menambahkan data baru",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil Membuat Data Baru",
		"data":    newUser,
	})

}

// Mencari data berdasarkan id
func UserControllerFind(c *fiber.Ctx) error {
	var user []entity.User
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID tidak Boleh Kosong",
		})
		return nil
	}

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data User tidak di temukan",
		})
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Berhasil mengambil data user",
		"data":    user,
	})
}

// update data
func UserControllerUpdate(c *fiber.Ctx) error {
	id := c.Params("id")
	user := new(entity.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.DB.Where("id = ?", id).Updates(&user)
	return c.JSON(user)

}

func UserControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	var user []entity.User

	result := database.DB.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.JSON(map[string]string{
			"message": "Data user tidak ada , Silahkan cek kembali",
		})
	}

	return c.Status(200).JSON(map[string]string{
		"message": "Data berhasil di hapus",
	})
}
