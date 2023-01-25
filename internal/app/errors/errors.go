package errors

import (
	"github.com/gofiber/fiber/v2"
)

type FiberErr uint32

const (
	ErrBadRequest              FiberErr = 10400
	ErrUnauthorized            FiberErr = 10401
	ErrInactiveUser            FiberErr = 10604
	ErrForbidden               FiberErr = 10403
	ErrUnprocessableEntity     FiberErr = 10422
	ErrPayloadValidationFailed FiberErr = 10423
	ErrInternalServerError     FiberErr = 10500
)

// Map - Returns predefined fiber.Map based on the FiberErr constant that it was called upon.
func (fe FiberErr) Map() fiber.Map {
	switch fe {
	case ErrBadRequest:
		return fiber.Map{"message": "bad request"}
	case ErrUnauthorized:
		return fiber.Map{"message": "unauthorized"}
	case ErrForbidden:
		return fiber.Map{"message": "forbidden"}
	case ErrUnprocessableEntity:
		return fiber.Map{"message": "unprocessable entity"}
	case ErrPayloadValidationFailed:
		return fiber.Map{"message": "payload validation failed"}
	case ErrInactiveUser:
		return fiber.Map{"message": "the current user is inactive"}
	case ErrInternalServerError:
		return fiber.Map{"message": "server issue occurred"}
	}

	return nil
}

// Code - returns predefined custom code attached to the FiberErr constant that it was called upon.
func (fe FiberErr) Code() uint32 {
	return uint32(fe)
}
