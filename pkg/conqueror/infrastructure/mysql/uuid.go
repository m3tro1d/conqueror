package mysql

import (
	"database/sql/driver"
	"encoding/hex"
	"strings"

	"conqueror/pkg/common/uuid"
)

type nullBinaryUUID struct {
	UUID  uuid.UUID
	Valid bool
}

type binaryUUID uuid.UUID

func (uid binaryUUID) Value() (driver.Value, error) {
	return uuid.UUID(uid).Bytes(), nil
}

func (uid binaryUUID) Hex() string {
	return hex.EncodeToString(uid[:])
}

func (uid *binaryUUID) Scan(src interface{}) error {
	var result uuid.UUID
	err := result.Scan(src)
	*uid = binaryUUID(result)
	return err
}

func (uid nullBinaryUUID) Value() (driver.Value, error) {
	if !uid.Valid {
		return nil, nil
	}
	// Delegate to UUID Value function
	return uid.UUID.Bytes(), nil
}

func (uid *nullBinaryUUID) Scan(src interface{}) error {
	if src == nil {
		uid.UUID = uuid.UUID{}
		uid.Valid = false
		return nil
	}

	// Delegate to UUID Scan function
	uid.Valid = true
	return uid.UUID.Scan(src)
}

func (uid *nullBinaryUUID) ToOptionalUUID() *uuid.UUID {
	if uid.Valid {
		id := uid.UUID
		return &id
	}
	return nil
}

func makeNullBinaryUUID(id *uuid.UUID) nullBinaryUUID {
	result := nullBinaryUUID{}
	if id != nil {
		result.Valid = true
		result.UUID = *id
	}
	return result
}

func getIDsFilterQuery(ids []uuid.UUID) string {
	var q strings.Builder
	q.WriteString("IN (")

	for i, id := range ids {
		q.WriteString("0x")
		q.WriteString(binaryUUID(id).Hex())

		if i < len(ids)-1 {
			q.WriteRune(',')
		}
	}

	q.WriteRune(')')
	return q.String()
}
