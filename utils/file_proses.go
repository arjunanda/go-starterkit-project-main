package utils

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
)

func UploadImageUser(c *fiber.Ctx, b *multipart.FileHeader) error {

	err := c.SaveFile(b, fmt.Sprintf("./public/images/%s", b.Filename))

	if err != nil {
		return err
	}

	return err
}

func DeleteImageUser(ImageName string) error {

	err := os.Remove(fmt.Sprintf("./public/images/%s", ImageName))

	if err != nil {
		return err
	}

	return err
}
