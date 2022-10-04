package types

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type GqlID struct {
	Prefix string
	Uuid   uuid.UUID
}

func (gid GqlID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(fmt.Sprintf("/%s/%s", gid.Prefix, gid.Uuid.String())))
	// io.WriteString(w, fmt.Sprintf("/%s/%s", gid.Prefix, gid.Uuid.String()))
}

func (gid *GqlID) UnmarshalGQLWithPrefix(prefix string, v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid type %T, expect string", v)
	}
	idPrefix := fmt.Sprintf("/%s/", prefix)
	if !strings.HasPrefix(s, idPrefix) {
		return fmt.Errorf("expected ID to start with prefix %s", idPrefix)
	}
	p, err := uuid.Parse(s[len(idPrefix):])
	*gid = GqlID{prefix, p}
	return err
}

// Value implements sql.Valuer so that UUIDs can be written to databases
// transparently. Currently, UUIDs map to strings. Please consult
// database-specific driver documentation for matching types.
func (gid GqlID) Value() (driver.Value, error) {
	return gid.Uuid.Value()
}

// Scan implements sql.Scanner so UUIDs can be read from databases transparently.
// Currently, database types that map to string and []byte are supported. Please
// consult database-specific driver documentation for matching types.
func (id *GqlID) ScanWithPrefix(prefix string, src interface{}) error {
	id.Prefix = prefix

	return id.Uuid.Scan(src)
}
