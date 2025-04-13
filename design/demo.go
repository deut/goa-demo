package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("hamster", func() {
	Description("Hamsters distribution API")
	Method("list", func() {
		Description("List hamsters with optional pagination.")
		Result(ArrayOf(Hamster))

		HTTP(func() {
			GET("/hamsters")
			Response(StatusOK)
		})
	})

	Method("create", func() {
		Description("Add newborn hamster")

		Payload(HamsterPayload)
		Result(Hamster)

		HTTP(func() {
			POST("/haster")
			Response(StatusCreated)
		})
	})

	Method("show", func() {
		Description("Choose your hamster")
		Payload(func() {
			Attribute("hamsterID", String, "Hamster UUID", func() {
				Format(FormatUUID)
			})
			Required("hamsterID")
		})

		Result(Hamster)
		Error("not_found")

		HTTP(func() {
			GET("/hamster/{hamsterID}")
			Response(StatusOK)
			Response("not_found", StatusNotFound)
		})
	})
})

// Data Types
var HamsterPayload = Type("HamsterPayload", func() {
	Description("A hamster")

	Attribute("name", String, "Name of hamster", func() {
		MinLength(1)
		Example("Fluffy")
	})
	Attribute("colors", ArrayOf(String), "fur colors", func() {
		Description("List of colors the hamster can have.")
		Example([]string{"brown", "white"})
		MaxLength(1)
	})
})

var Hamster = Type("Hamster", func() {
	Description("A hamster")
	Extend(HamsterPayload)

	Attribute("id", String, "Unique concert ID", func() {
		Format(FormatUUID)
		Example("123e4567-e89b-12d3-a456-426614174000")
	})
	Required("id", "name", "colors")
})
