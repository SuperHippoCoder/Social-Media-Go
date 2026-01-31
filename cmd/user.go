package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleLoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func (h *Handler) HandleLolGet(c *gin.Context) {
	c.HTML(http.StatusOK, "LOL", nil)
}

func (h *Handler) HandleLoginPost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required,min=3,max=32"`
		Password string `form:"password" binding:"required,min=1"`
	}

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println("Ошибка валидации:", err)
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"Error": "Проверьте данные"})
		return
	}

	_, err := h.users.AuthenticateUser(form.Username, form.Password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Неверный логин или пароль"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/lol")
}

func (h *Handler) HandleRegistrationGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func (h *Handler) HandleRegistrationPost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required,min=3,max=32"`
		Password string `form:"password" binding:"required,min=1"`
	}

	if err := c.ShouldBind(&form); err != nil {
		fmt.Println("Ошибка валидации:", err)
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"Error": "Проверьте данные"})
		return
	}

	err := h.users.RegisterUser(form.Username, form.Password)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "register.html", gin.H{"Error": err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/lol")
}
