// Package secure contains support for application security
package secure

import (
	"hash"

	zxcvbn "github.com/nbutton23/zxcvbn-go"
	"golang.org/x/crypto/bcrypt"
)

// New initalizes security service
func New(minPWStr int, h hash.Hash) *Service {
	return &Service{minPWStr: minPWStr, h: h}
}

// Service holds security related methods
type Service struct {
	minPWStr int
	h        hash.Hash
}

// Password checks whether password is secure enough using zxcvbn library
func (s *Service) Password(pass string, inputs ...string) bool {
	pwStrength := zxcvbn.PasswordStrength(pass, inputs)
	return pwStrength.Score >= s.minPWStr
}

// Hash hashes the password using bcrypt
func (*Service) Hash(password string) string {
	hashedPW, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPW)
}

func (*Service) HashMatchesPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
