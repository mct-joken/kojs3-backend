package handlers

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/server/responses"
)

type UserHandlers struct {
	controller.UserController
	auth   controller.AuthController
	logger *zap.Logger
}

func NewUserHandlers(
	userController controller.UserController,
	authController controller.AuthController,
	logger *zap.Logger,
) *UserHandlers {
	return &UserHandlers{userController, authController, logger}
}

func (h *UserHandlers) CreateUser(c echo.Context) error {
	req := model.CreateUserRequestJSON{}
	if err := c.Bind(&req); err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}

	res, err := h.UserController.Create(req)
	if err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandlers) FindAllUser(c echo.Context) error {
	res, err := h.UserController.FindAllUsers()
	if err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *UserHandlers) Login(c echo.Context) error {
	req := model.LoginRequestJSON{}
	if err := c.Bind(&req); err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}

	res, err := h.auth.Login(req)
	if err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusInternalServerError, responses.InternalServerErrorResponseJSON)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandlers) Verify(c echo.Context) error {
	t := c.Param("token")
	ok, err := h.auth.Verify(t)
	if err != nil {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}
	if !ok {
		h.logger.Sugar().Errorf("%s", err)
		return c.JSON(http.StatusBadRequest, responses.InvalidRequestErrorResponseJSON)
	}
	return c.NoContent(http.StatusOK)
}
