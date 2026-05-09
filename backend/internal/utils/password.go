package utils

import "github.com/alexedwards/argon2id"

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(
		password,
		argon2id.DefaultParams,
	)

	return hash, err
}

func CheckPassword(password, hash string) error {
	match, err := argon2id.ComparePasswordAndHash(password, hash)

	if err != nil {
		return err
	}

	if !match {
		return err
	}

	return nil
}
