package jwt

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kaynetik/modular-monolith-example/tests"
)

// TestGetIssuer
//
// Placeholder for JWTSuite test example.
func (s *JWTSuite) TestGetIssuer() {
	var (
		issuerID string
		claims   jwt.MapClaims
	)

	tc := []tests.TestCase{
		{
			Name: "valid issuer in claims",
			PreRequisites: func() {
				issuerID = uuid.Must(uuid.NewV4()).String()
				claims = jwt.MapClaims{"iss": issuerID}
			},
			Assert: func() {
				issuer, err := GetIssuer(claims)
				s.NoError(err)
				s.NotEmpty(issuer)
				s.True(claims.VerifyIssuer(issuerID, true))
			},
		},
	}

	tests.RunTestCases(s, tc)
}
