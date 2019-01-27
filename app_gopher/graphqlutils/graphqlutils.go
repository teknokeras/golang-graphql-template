package graphqlutils

import (
	"fmt"
	"strconv"
	"strings"
	"encoding/base64"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

func gqlIDToUint(i graphql.ID) (uint, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, errors.Wrap(err, "GqlIDToUint")
	}

	return uint(r), nil
}

func int32P(i uint) *int32 {
	r := int32(i)
	return &r
}

func boolP(b bool) *bool {
	return &b
}

func gqlIDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}

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