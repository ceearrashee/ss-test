package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	irisJWT "github.com/kataras/iris/v12/middleware/jwt"

	"solid-software.test-task/pkg/framework/config"
)

type (
	// Service is an interface that has methods for working with JSON Web Tokens(JWT).
	Service interface {
		// GetToken creates a new token with the provided claim.
		GetToken(claim SampleClaim) ([]byte, error)
		// VerifyTokenAndReturnClaim verifies the token and returns the claim if the token is valid.
		VerifyTokenAndReturnClaim(ctx iris.Context) (*SampleClaim, error)
		// GetHandler returns a handler that verifies the incoming JWT tokens.
		GetHandler() iris.Handler
	}
	// SampleClaim represents the claim information embedded in the JWT token.
	SampleClaim struct {
		Username string `json:"username"`
	}
	jwt struct {
		signer   *irisJWT.Signer
		verifier *irisJWT.Verifier
		secret   string
	}
)

var (
	// ErrTokenInvalid is returned when the JWT token is not found or invalid.
	ErrTokenInvalid = errors.New("token not found or invalid")
	// ErrTokenExpired is returned when the JWT token has expired.
	ErrTokenExpired = errors.New("token expired")
)

// NewJWT constructs a JWTService with the provided configurations.
func NewJWT(conf config.Config) Service {
	secret := conf.GetString("webService.jwt.secret")
	tokenExpirationTime := time.Minute * time.Duration(conf.GetInt64("webService.jwt.tokenExpirationTimeInMinutes"))
	signer := irisJWT.NewSigner(irisJWT.HS256, secret, tokenExpirationTime)
	verifier := irisJWT.NewVerifier(irisJWT.HS256, secret)

	return &jwt{
		signer:   signer,
		verifier: verifier,
		secret:   secret,
	}
}

// VerifyTokenAndReturnClaim verifies the token and returns the claim if the token is valid.
func (*jwt) VerifyTokenAndReturnClaim(irisCtx iris.Context) (*SampleClaim, error) {
	verifiedToken := irisJWT.GetVerifiedToken(irisCtx)
	if verifiedToken == nil {
		return nil, ErrTokenInvalid
	}

	timeLeft := verifiedToken.StandardClaims.Timeleft()
	if timeLeft.Microseconds() <= 0 {
		return nil, ErrTokenExpired
	}

	var claim SampleClaim

	err := verifiedToken.Claims(claim)
	if err != nil {
		return nil, fmt.Errorf("get SampleClaim: %w", err)
	}

	return &claim, nil
}

// GetToken creates a new token with the provided claim.
func (j *jwt) GetToken(claim SampleClaim) ([]byte, error) {
	tokenBytes, err := j.signer.Sign(&claim)
	if err != nil {
		return nil, fmt.Errorf("sign claim: %w", err)
	}

	return tokenBytes, nil
}

// GetHandler returns a handler that verifies the incoming JWT tokens.
func (j *jwt) GetHandler() iris.Handler {
	return j.verifier.Verify(
		func() any {
			return &SampleClaim{}
		},
	)
}
