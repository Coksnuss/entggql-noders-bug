package types

import (
	"database/sql/driver"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type PostID GqlID

func (id PostID) MarshalGQL(w io.Writer) {
	GqlID(id).MarshalGQL(w)
}

func (id *PostID) UnmarshalGQL(v interface{}) error {
	var gid GqlID
	if err := gid.UnmarshalGQLWithPrefix("posts", v); err != nil {
		return err
	}
	*id = PostID(gid)
	return nil
}

func (id PostID) Value() (driver.Value, error) {
	return GqlID(id).Value()
}

func (id *PostID) Scan(src interface{}) error {
	return (*GqlID)(id).ScanWithPrefix("posts", src)
}

func NewPostID(id uuid.UUID) PostID {
	return PostID{"posts", id}
}

func MarshalPostID(id PostID) graphql.Marshaler {
	return id
}

func UnmarshalPostID(v interface{}) (id PostID, err error) {
	err = id.UnmarshalGQL(v)
	return
}
