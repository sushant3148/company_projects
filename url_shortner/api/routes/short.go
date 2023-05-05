package routes

import (
	"fmt"
	"time"
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
func ShortUrl(t *fiber.Ctx)error{
	body := new(request)
	if err := t.BodyParsing(&body); err != nil{
		return t.Status(fiber.StatusBadRequest).JSON(fiber.Map("error":"cannot parse this json"))
	}
	if !govalidator.IsURL(body.Url){
		return t.Status(fiber.StatusBadRequest).JSON(fiber.Map("error":"invalid url"))
	}

	if !helpers.DomainError(body.Url) {
		return t.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map("error":"domain is incorrect"))
	}
	body.Url = helpers.EnforceHttp(body.Url)


}
// storing 20000 url for per user


