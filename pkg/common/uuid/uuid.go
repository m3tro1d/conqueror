package uuid

import (
	stderrors "errors"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

const uuidSize = 16

var ErrInvalidUUID = stderrors.New("invalid UUID")

type UUID [uuidSize]byte

func (u *UUID) Scan(src interface{}) error {
	var impl uuid.UUID
	err := impl.Scan(src)

	*u = UUID(impl)
	return err
}

func (u UUID) String() string {
	impl := uuid.UUID(u)
	return impl.String()
}

func (u UUID) Bytes() []byte {
	impl := uuid.UUID(u)
	return impl.Bytes()
}

func Generate() UUID {
	impl := uuid.NewV1()
	return UUID(impl)
}

func FromString(input string) (u UUID, err error) {
	impl, err := uuid.FromString(input)
	if err != nil {
		return u, errors.WithStack(ErrInvalidUUID)
	}
	u = UUID(impl)
	return
}

func OptionalFromString(input *string) (*UUID, error) {
	if input == nil {
		return nil, nil
	}
	u, err := FromString(*input)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func OptionalToString(id *UUID) string {
	if id == nil {
		return ""
	}
	return id.String()
}

func FromStrings(ids []string) ([]UUID, error) {
	result := make([]UUID, 0, len(ids))
	for _, strId := range ids {
		id, err := FromString(strId)
		if err != nil {
			return nil, err
		}

		result = append(result, id)
	}

	return result, nil
}
