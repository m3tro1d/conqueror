package uuid

import (
	uuid "github.com/satori/go.uuid"
)

const uuidSize = 16

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

func FromString(input string) (u UUID, err error) {
	impl, err := uuid.FromString(input)
	if err != nil {
		return u, err
	}
	u = UUID(impl)
	return
}

func Generate() UUID {
	impl := uuid.NewV1()
	return UUID(impl)
}
