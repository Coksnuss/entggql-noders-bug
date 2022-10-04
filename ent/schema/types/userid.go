package types

import (
	"database/sql/driver"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type UserID GqlID

func (id UserID) MarshalGQL(w io.Writer) {
	GqlID(id).MarshalGQL(w)
}

func (id *UserID) UnmarshalGQL(v interface{}) error {
	var gid GqlID
	if err := gid.UnmarshalGQLWithPrefix("users", v); err != nil {
		return err
	}
	*id = UserID(gid)
	return nil
}

func (id UserID) Value() (driver.Value, error) {
	return GqlID(id).Value()
}

func (id *UserID) Scan(src interface{}) error {
	return (*GqlID)(id).ScanWithPrefix("users", src)
}

func NewUserID(id uuid.UUID) UserID {
	return UserID{"users", id}
}

func MarshalUserID(id UserID) graphql.Marshaler {
	return id
}

func UnmarshalUserID(v interface{}) (id UserID, err error) {
	err = id.UnmarshalGQL(v)
	return
}
