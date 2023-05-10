package routes

import(
	"assigment2/api/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

)
func OriginalUrl(c * fiber.Ctx) error{
	url := c.Params("url")
	// intiating database
	t := database.CreateClient(0)

	defer t.Close()
	value,err := t.Get(database.Ctx,url).Result()
	if err == redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"url not found"})
	}else if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"internal server occured"})
	}

	rIncrement := database.CreateClient(1)
	rIncrement.Close()
	_= rIncrement.Incr(database.Ctx,"counter")
	return c.Redirect(value,301)

}