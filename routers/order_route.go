package routers

import (
	"github.com/bitebait/cupcakestore/controllers"
	"github.com/bitebait/cupcakestore/database"
	"github.com/bitebait/cupcakestore/middlewares"
	"github.com/bitebait/cupcakestore/repositories"
	"github.com/bitebait/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type OrderRouter struct {
	orderController controllers.OrderController
}

func NewOrderRouter() *OrderRouter {
	orderRepository := repositories.NewOrderRepository(database.DB)
	storeConfigRepository := repositories.NewStoreConfigRepository(database.DB)
	storeConfigService := services.NewStoreConfigService(storeConfigRepository)
	orderService := services.NewOrderService(orderRepository, storeConfigService)
	orderController := controllers.NewOrderController(orderService, storeConfigService)

	return &OrderRouter{
		orderController: orderController,
	}
}

func (r *OrderRouter) InstallRouters(app *fiber.App) {
	order := app.Group("/orders")
	order.Get("/checkout/:id", r.orderController.Checkout)
	order.Post("/payment/:id", r.orderController.Payment)
	order.Get("/payment/:id", r.orderController.Payment)
	order.Get("/cancel/:id", r.orderController.RenderCancel)
	order.Post("/cancel/:id", r.orderController.Cancel)
	order.Get("/myorders", r.orderController.RenderUserOrders)
	order.Get("/myorders/:id", r.orderController.RenderOrder)
	order.Get("/admin", r.orderController.RenderAdminOrders).Use(middlewares.LoginAndStaffRequired())
}
