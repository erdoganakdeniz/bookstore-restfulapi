package controllers

import (
	"bookstoreapi/app/models"
	"bookstoreapi/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
	"bookstoreapi/pkg/utils"
)

func GetBooks(c *fiber.Ctx) error {
	db,err:=database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":err.Error(),
		})
	}

	books,err:=db.GetBooks()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "books were not found",
			"count": 0,
			"books": nil,
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(books),
		"books": books,
	})

}
func GetBook(c *fiber.Ctx) error {
	id,err:=uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":true,
			"msg": err.Error(),
		})
	}
	db,err:=database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":true,
			"msg": err.Error(),
		})
	}

	book,err:=db.GetBook(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with the given ID is not found",
			"book":  nil,
		})

	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})


}
func CreateBook(c *fiber.Ctx) error {
	now :=time.Now().Unix()
	claims,err:=utils.ExtractTokenMetaData(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":true,
			"msg":err.Error(),
		})
	}
	expires:=claims.Expires
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	book:=&models.Book{}
	if err:=c.BodyParser(book);err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	db,err:=database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	validate:=utils.NewValidator()
	book.ID=uuid.New()
	book.CreatedAt=time.Now()
	book.BookStatus=1
	if err:=validate.Struct(book);err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	if err:=db.CreateBook(book);err!=nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})

}
//Update Book Controller
func UpdateBook(c *fiber.Ctx) error {
	now :=time.Now().Unix()

	claims,err:=utils.ExtractTokenMetaData(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	expires:=claims.Expires
	if now>expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}
	book:=&models.Book{}
	if err:=c.BodyParser(book);err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	db,err:=database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	foundedBook , err :=db.GetBook(book.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with this ID not found",
		})
	}
	book.UpdatedAt=time.Now()

	validate:=utils.NewValidator()
	if err:=validate.Struct(book);err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	if err:=db.UpdateBook(foundedBook.ID,book) ;err!= nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})

	}
	if err:=db.DeleteBook(foundedBook.ID,book) ;err!= nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusCreated)

	
}


