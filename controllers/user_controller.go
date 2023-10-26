package controllers

import (
	"github.com/bitebait/cupcakestore/models"
	"github.com/bitebait/cupcakestore/services"
	"github.com/bitebait/cupcakestore/utils"
	"github.com/bitebait/cupcakestore/views"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	RenderCreate(ctx *fiber.Ctx) error
	HandlerCreate(ctx *fiber.Ctx) error
	RenderUsers(ctx *fiber.Ctx) error
	RenderUser(ctx *fiber.Ctx) error
	HandlerUpdate(ctx *fiber.Ctx) error
	RenderDelete(ctx *fiber.Ctx) error
	HandlerDelete(ctx *fiber.Ctx) error
}

type userController struct {
	userService    services.UserService
	profileService services.ProfileService
}

func NewUserController(u services.UserService, p services.ProfileService) UserController {
	return &userController{
		userService:    u,
		profileService: p,
	}
}

func (c *userController) RenderCreate(ctx *fiber.Ctx) error {
	return views.Render(ctx, "users/create", nil, "", baseLayout)
}

func (c *userController) HandlerCreate(ctx *fiber.Ctx) error {
	user := &models.User{}
	if err := ctx.BodyParser(user); err != nil {
		return views.Render(ctx, "users/create", nil,
			"Dados de usuário inválidos: "+err.Error(), baseLayout)
	}

	if err := c.userService.Create(user); err != nil {
		return views.Render(ctx, "users/create", nil,
			"Falha ao criar usuário: "+err.Error(), baseLayout)
	}

	profile := &models.Profile{
		UserID: user.ID,
	}

	if err := c.profileService.Create(profile); err != nil {
		return views.Render(ctx, "users/create", nil,
			"Falha ao criar Perfil: "+err.Error(), baseLayout)
	}

	return ctx.Redirect("/users")
}

func (c *userController) RenderUsers(ctx *fiber.Ctx) error {
	query := ctx.Query("q", "")

	pagination := models.NewPagination(ctx.QueryInt("page"), ctx.QueryInt("limit"))
	users := c.userService.FindAll(pagination, query)
	data := fiber.Map{
		"Users":      users,
		"Pagination": pagination,
	}

	return views.Render(ctx, "users/users", data, "", baseLayout)
}

func (c *userController) RenderUser(ctx *fiber.Ctx) error {
	userID, err := utils.StringToId(ctx.Params("id"))
	if err != nil {
		return ctx.Redirect("/users")
	}

	user, err := c.userService.FindById(uint(userID))
	if err != nil {
		return ctx.Redirect("/users")
	}

	profile, err := c.profileService.FindByUserId(uint(userID))
	if err != nil {
		return ctx.Redirect("/users")
	}

	data := fiber.Map{
		"User":    user,
		"Profile": profile,
	}

	return views.Render(ctx, "users/user", data, "", baseLayout)
}

func (c *userController) HandlerUpdate(ctx *fiber.Ctx) error {
	id, err := utils.StringToId(ctx.Params("id"))
	if err != nil {
		return ctx.Redirect("/users")
	}

	user, err := c.userService.FindById(id)
	if err != nil {
		return ctx.Redirect("/users")
	}

	profile, err := c.profileService.FindByUserId(id)
	if err != nil {
		return ctx.Redirect("/users")
	}

	data := fiber.Map{
		"User":    user,
		"Profile": profile,
	}

	if err := ctx.BodyParser(&user); err != nil {
		return views.Render(ctx, "users/user", data, err.Error(), baseLayout)
	}

	oldPassword := ctx.FormValue("oldPassword")
	newPassword := ctx.FormValue("newPassword")
	if oldPassword != "" && newPassword != "" {
		if err := user.UpdatePassword(oldPassword, newPassword); err != nil {
			return views.Render(ctx, "users/user", data,
				"Falha ao atualizar a senha do usuário. Por favor, verifique os dados.", baseLayout)
		}
	}

	user.IsStaff = ctx.FormValue("isStaff") == "on"
	user.IsActive = ctx.FormValue("isActive") == "on"
	if err := c.userService.Update(&user); err != nil {
		return views.Render(ctx, "users/user", data,
			"Falha ao atualizar usuário.", baseLayout)
	}

	return ctx.Redirect("/users")
}

func (c *userController) RenderDelete(ctx *fiber.Ctx) error {
	id, err := utils.StringToId(ctx.Params("id"))
	if err != nil {
		return ctx.Redirect("/users")
	}

	user, err := c.userService.FindById(id)
	if err != nil {
		return ctx.Redirect("/users")
	}

	return views.Render(ctx, "users/delete", user, "", baseLayout)
}

func (c *userController) HandlerDelete(ctx *fiber.Ctx) error {
	id, err := utils.StringToId(ctx.Params("id"))
	if err != nil {
		return ctx.Redirect("/users")
	}

	err = c.userService.Delete(id)
	if err != nil {
		return ctx.Redirect("/users")
	}

	return ctx.Redirect("/users")
}
