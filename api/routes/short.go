package routes

import (
	"os"
	"strconv"
	"time"
	"assigment2/api/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
    "github.com/asaskevich/govalidator"
    "assigment2/api/helper"
)
type request struct{
	Url string `json:"url"`
	ShortUrl string `json:"short_url"`
	Expiry time.Duration `json:"expiry"`
}
type response struct{
	Url string `json:"url"`
	ShortUrl string `json:"short_url"`
	Expiry time.Duration `json:"expiry"`
	ShortLinkRemaining int `json:"short_link_remaining"`
}
func ShortUrl(c *fiber.Ctx) error{
	body := new(request)
	if err := c.BodyParser(&body); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse this json"})
	}
	r2 := database.CreateClient(0)
	defer r2.Close()
	value,err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil{
		_ =r2.Set(database.Ctx,c.IP(),os.Getenv("MAX_API"),1440*60*time.Second).Err()
	}else{
		value, _ =  r2.Get(database.Ctx, c.IP()).Result()
		vlInt,_ := strconv.Atoi(value)
		if vlInt <= 0{
			// limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"Rate limit excedded"})
		}
	}
// validating the url
	if !govalidator.IsURL(body.Url){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"invalid url"})
	}
// incountered all the domain error
	if !helpers.DomainError(body.Url) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"domain is incorrect"})
	}
	body.Url = helpers.EnforceHttp(body.Url)
	r := database.CreateClient(0)
	var id string
	defer r.Close()
	val,err := r.Get(database.Ctx,id).Result()
	if val != ""{
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error":"id already exist"})
	}
	if body.Expiry == 0{
		body.Expiry = 24
	}
	r.Decr(database.Ctx,c.IP())
    // passing the response
	resp := response{
		Url: body.Url,
		ShortUrl : "",
		Expiry : body.Expiry,
		ShortLinkRemaining : 2000,
	}
    
	val1,_ := r2.Get(database.Ctx,c.IP()).Result()
	resp.ShortLinkRemaining,_ = strconv.Atoi(val1) 
	resp.ShortUrl = os.Getenv("DOMAIN") + "/" + id
    return  c.Status(fiber.StatusOK).JSON(resp)
	

}