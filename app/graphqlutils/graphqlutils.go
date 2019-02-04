package graphqlutils

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
)

// encode cursor encodes the cursot position in base64
func EncodeCursor(i int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i))))
}

// decode cursor decodes the base 64 encoded cursor and resturns the integer
func DecodeCursor(s string) (int, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
	if err != nil {
		return 0, err
	}

	return i, nil
}
