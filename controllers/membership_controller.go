package controllers

import (
	"miniproject/configs"
	"miniproject/models"
	"miniproject/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMembershipsController(c echo.Context) error {
	var membership []models.Membership
	if err := configs.DB.Find(&membership).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	var membershipResponse []models.MembershipResponse
	for _, member := range membership {
		membershipResponse = append(membershipResponse, member.ToMembershipResponse())
	}

	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Retrieved Memberships",
		},
		Results: membershipResponse,
	}
	return c.JSON(http.StatusOK, response)

}

func GetMembershipController(c echo.Context) error {
	var membership models.Membership
	if err := configs.DB.First(&membership, c.Param("id")).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	membershipResponse := membership.ToMembershipResponse()
	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Get Membership",
		},
		Results: membershipResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func CreateMembershipController(c echo.Context) error {
	membership := new(models.Membership)
	if err := c.Bind(membership); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if membership.Name == "" || membership.Speed == "" || membership.Price == 0 {
		return utils.ErrorResponse(c, http.StatusBadRequest, "All fields are required")
	}

	if err := configs.DB.Create(membership).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	membershipResponse := membership.ToMembershipResponse()

	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Created Memberships",
		},
		Results: membershipResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteMembershipController(c echo.Context) error {
	id := c.Param("id")
	var membership models.Membership

	if err := configs.DB.First(&membership, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Membership not found")

	}

	if err := configs.DB.Delete(&membership).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Deleted Memberships",
		},
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateMembershipController(c echo.Context) error {
	id := c.Param("id")
	membership := models.Membership{}
	if err := configs.DB.First(&membership, id).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Membership not found")

	}
	if err := c.Bind(&membership); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := configs.DB.Save(&membership).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	membershipResponse := membership.ToMembershipResponse()
	response := utils.TSuccessResponse{
		Meta: utils.TResponseMeta{
			Success: true,
			Message: "Success! Updated Memberships",
		},
		Results: membershipResponse,
	}
	return c.JSON(http.StatusOK, response)
}
