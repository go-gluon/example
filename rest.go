package main

// import "github.com/gofiber/fiber/v2"

// //gluon:rest path=/user consumer=application/json producer=application/json
// type UserRestController struct {
// 	Name string
// }

// //gluon:init
// func (u *UserRestController) Init() {
// 	u.Name = "Test"
// }

// //gluon:rest method=GET
// func (u *UserRestController) GetUser() string {
// 	return "User"
// }

// //gluon:rest method=GET path=:id
// func (u *UserRestController) GetInfo(id string) string {
// 	return "Info:" + id
// }

// func NewUserRestController() *fiber.App {
// 	c := &UserRestController{}
// 	c.Init()

// 	app := fiber.New()
// 	app.Add(fiber.MethodGet, "/", func(ctx *fiber.Ctx) error {
// 		return ctx.SendString(c.GetUser())
// 	})
// 	app.Add(fiber.MethodGet, "/:id", func(ctx *fiber.Ctx) error {
// 		id := ctx.Params("id", "")
// 		return ctx.SendString(c.GetInfo(id))
// 	})
// 	return app
// }
