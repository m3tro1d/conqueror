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

func (u *UUID) UnmarshalText(text []byte) error {
	var impl uuid.UUID
	err := impl.UnmarshalText(text)

	*u = UUID(impl)
	return err
}

func (u UUID) MarshalText() ([]byte, error) {
	impl := uuid.UUID(u)
	return impl.MarshalText()
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

func Equal(u1, u2 UUID) bool {
	return uuid.Equal(uuid.UUID(u1), uuid.UUID(u2))
}
