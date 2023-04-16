package controllers

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllHarvests(c *fiber.Ctx) error {
	resultHarvest, err := services.GetAllHarvests()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	if len(resultHarvest) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Harvest list not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Harvest list successfully",
		Data:    resultHarvest,
	})
}

func GetOneHarvest(c *fiber.Ctx) error {
	HarvestID := c.Params("id")
	Harvest, err := services.GetOneHarvest(HarvestID)

	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Harvest not found, " + err.Error(),
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Harvest successfully",
		Data:    Harvest,
	})
}

func GetHistoricHarvestEsimation(c *fiber.Ctx) error {
	FarmLotID := c.Params("idFarmLot")
	HistoricHarvest, err := services.GetHistoricHarvestEsimation(FarmLotID)

	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(models.Response{
			Success: false,
			Message: "Historic Harvest not found, " + err.Error(),
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Success: true,
		Message: "Historic Harvest successfully",
		Data:    HistoricHarvest,
	})
}