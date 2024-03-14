package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"htmxwebsite.dev/matheus-alpe/internal/templates"
)

func countExample(e *echo.Echo) {
	count := struct{ Count int }{
		Count: 0,
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", count)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count)
	})
}

type Contact struct {
	Name  string
	Email string
}

type Contacts []*Contact

func (c Contacts) hasEmail(email string) bool {
	for _, contact := range c {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type Data struct {
	Contacts Contacts
}

func newData() Data {
	return Data{
		Contacts: Contacts{
			{Name: "Matheus", Email: "matt.alpe.dev@gmail.com"},
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func newFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func formExample(e *echo.Echo) {
	page := newPage()

	e.GET("/form", func(c echo.Context) error {
		return c.Render(200, "form-page", page)
	})

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if page.Data.Contacts.hasEmail(email) {
			formData := newFormData()
			formData.Values["name"] = name
			formData.Values["email"] = email
			formData.Errors["email"] = "Email already exists"

			return c.Render(422, "form", formData)
		}

		contact := &Contact{
			Name:  name,
			Email: email,
		}

		page.Data.Contacts = append(page.Data.Contacts, contact)

		c.Render(200, "form", newFormData())
		return c.Render(200, "oob-contact", contact)
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = templates.New()

	countExample(e)
	formExample(e)

	e.Logger.Fatal(e.Start(":42069"))
}
