package main

import (
	"monitor/service"
	"net/http"
	"strconv"

	"git.fuyu.moe/Fuyu/router"
)

func getServices(c *router.Context) error {
	services := service.GetAll(qdb)

	return c.JSON(http.StatusOK, services)
}

func getService(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	service, ok := service.GetByID(qdb, id)
	if !ok {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, service)
}

func subscribe(c *router.Context, serviceID int) error {
	service.CreateSubscription(qdb, serviceID)
	return c.NoContent(http.StatusOK)
}

func unsubscribe(c *router.Context) error {
	id, err := strconv.Atoi(c.Param(`id`))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	service.DeleteSubscription(qdb, id)
	return c.NoContent(http.StatusOK)
}

func getSubscriptions(c *router.Context) error {
	subscriptions := service.GetAllSubscriptions(qdb)

	return c.JSON(http.StatusOK, subscriptions)
}
