package env

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

const (
	emailMaxLen = 100
	emailRegex  = `^([a-zA-Z0-9.!#$%&'*+/=?^_\x60{|}~-]+)@([a-zA-Z0-9_\-\.]+)\.([a-zA-Z]{2,60})$`
)

type EnvironmentVariable string

func (e EnvironmentVariable) String() string {
	return string(e)
}

func Get(e EnvironmentVariable) string {
	return os.Getenv(e.String())
}

func GetOr(e EnvironmentVariable, fallback string) string {
	if value := Get(e); value != "" {
		return value
	}

	return fallback
}

func GetUint(e EnvironmentVariable) (uint, error) {
	intValue, err := GetInt(e)
	if err != nil {
		return 0, err
	}

	if intValue < 0 {
		return 0, errors.New("uint value must not be less than zero")
	}

	return uint(intValue), nil
}

func GetInt(e EnvironmentVariable) (int, error) {
	stringValue := Get(e)
	if stringValue == "" {
		// Not provided
		return 0, nil
	}

	result, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0, fmt.Errorf("failed to parse value: %w", err)
	}

	return result, nil
}

func GetDuration(e EnvironmentVariable) (time.Duration, error) {
	intValue, err := GetInt(e)
	if err != nil {
		return 0, fmt.Errorf("failed to get integer: %w", err)
	}

	if intValue == 0 {
		return 0, nil
	}

	return time.Duration(intValue) * time.Second, nil
}

func GetDurationOr(e EnvironmentVariable, fallback time.Duration) (time.Duration, error) {
	value, err := GetDuration(e)
	if err != nil {
		return 0, err
	}

	if value.Microseconds() == 0 {
		return fallback, nil
	}

	return value, nil
}

func GetBoolean(e EnvironmentVariable) (bool, error) {
	stringValue := Get(e)
	if stringValue == "" {
		// Not provided
		return false, nil
	}

	return strconv.ParseBool(stringValue)
}

func GetBooleanOr(e EnvironmentVariable, fallback bool) bool {
	val, err := strconv.ParseBool(Get(e))
	if err != nil {
		return fallback
	}

	return val
}

func GetEmail(e EnvironmentVariable) (string, error) {
	email := Get(e)
	if len(email) < 1 {
		return "", errors.New("email is empty")
	}

	if len(email) > emailMaxLen {
		return "", errors.New("email is too long string")
	}

	re := regexp.MustCompile(emailRegex)

	if ok := re.MatchString(email); !ok {
		return "", fmt.Errorf("'%s' value doesn't match the email regexp", email)
	}

	return email, nil
}

func GetUUID(e EnvironmentVariable) (id uuid.UUID, err error) {
	idString := Get(e)
	if idString == "" {
		return uuid.Nil, errors.New("uuid is empty")
	}

	if id, err = uuid.FromString(idString); err != nil {
		return uuid.Nil, fmt.Errorf("failed to pasre uuid from string: %w", err)
	}

	if id == uuid.Nil {
		return id, errors.New("uuid value must not be nil")
	}

	return id, nil
}
