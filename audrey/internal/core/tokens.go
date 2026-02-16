// Copyright 2026 Evie. P.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package core

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fernet/fernet-go"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Server) RandomBites(size int) []byte {
	bites := make([]byte, size)

	if _, err := rand.Read(bites); err != nil {
		s.Log.Fatalf("Unable to generate new session auth: %s", err)
	}

	return bites
}

func (s *Server) NewWebToken(userId string) (string, error) {
	claims := jwt.MapClaims{"iss": "mystly", "sub": userId}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.Config.JWTSecret))

	return signed, err
}

func (s *Server) ParseWebToken(tokenString string) (string, error) {
	// TODO: Logging...

	token, err := jwt.Parse(
		tokenString,
		func(t *jwt.Token) (any, error) { return []byte(s.Config.JWTSecret), nil },
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
	)

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["sub"].(string), nil
	} else {
		return "", err
	}
}

func (s *Server) GetFernetKey() *fernet.Key {
	cwd, err := os.Getwd()
	if err != nil {
		s.Log.Fatal("Unable to generate or get Fernet Key. %s", err)
	}

	fp := filepath.Clean(fmt.Sprintf("%s/.fkey", cwd))
	fkey, err := os.ReadFile(fp)

	if err == nil {
		key, err := fernet.DecodeKey(string(fkey))
		if err != nil {
			s.Log.Fatal("Unable to generate or get Fernet Key. %s", err)
		}

		return key
	}

	var ekey fernet.Key
	err = ekey.Generate()
	if err != nil {
		s.Log.Fatal("Unable to generate or get Fernet Key. %s", err)
	}

	encKey := ekey.Encode()
	err = os.WriteFile(fp, []byte(encKey), 0644)

	if err != nil {
		s.Log.Fatal("Unable to generate or get Fernet Key. %s", err)
	}

	return &ekey
}
