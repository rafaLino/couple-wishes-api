package valueObjects

import "golang.org/x/crypto/bcrypt"

type Password struct {
	hash  []byte
	value string
	valid bool
}

func NewPassword(value string) *Password {
	hashedPassword, err := hashPassword(value)
	return &Password{
		hash:  hashedPassword,
		value: value,
		valid: err == nil,
	}
}

func (p *Password) Verify(password []byte) bool {
	err := bcrypt.CompareHashAndPassword(password, []byte(p.value))
	return err == nil
}

func (p *Password) Value() []byte {
	return p.hash
}

func hashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return bytes, err
}
