package extension

import (
	"database/sql/driver"
	"errors"
	"strconv"

	"github.com/speps/go-hashids"
)

type (
	// Hash is struct for hashids
	Hash struct {
		hash *hashids.HashID
	}

	// HashID64 is encoded by go-hashids
	HashID64 int64
)

var (
	_hash *Hash
)

// GetHash return *Hash
func GetHash() *Hash {
	return _hash
}

// NewHash is constructor of HashID
func NewHash(salt string, minLength int) (*Hash, error) {
	h := hashids.NewData()
	h.Salt = salt
	h.MinLength = minLength
	hd, err := hashids.NewWithData(h)

	_hash = &Hash{hash: hd}
	return _hash, err
}

// EncodeInt64 ...
func (h *Hash) EncodeInt64(id int64) (string, error) {
	return h.hash.EncodeInt64([]int64{id})
}

// DecodeInt64 ...
func (h *Hash) DecodeInt64(hash string) (int64, error) {
	id, err := h.hash.DecodeInt64WithError(hash)
	if len(id) == 0 {
		return 0, errors.New("hash string has not int64 value in DecodeInt64")
	}
	return id[0], err
}

// MarshalJSON mapping struct to json using HashID
func (h *HashID64) MarshalJSON() ([]byte, error) {
	hd := GetHash()
	encHashID, err := hd.EncodeInt64(int64(*h))
	return []byte(strconv.Quote(encHashID)), err
}

// UnmarshalJSON mapping json to struct using HashID64
func (h *HashID64) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	hd := GetHash()
	id, err := hd.DecodeInt64(s)
	if err != nil {
		return err
	}

	*h = HashID64(id)
	return nil
}

// Scan is implementation for sql driver interface
func (h *HashID64) Scan(value interface{}) error {
	i, ok := value.(int64)
	if !ok {
		return errors.New("do not cast int64 in HashID64's Scan")
	}
	*h = HashID64(i)
	return nil
}

// Value is implementation for sql driver interface
func (h HashID64) Value() (driver.Value, error) {
	return driver.Value(int64(h)), nil
}

// UnmarshalParam for Bind function of echo
func (h *HashID64) UnmarshalParam(src string) error {
	hd := GetHash()
	id, err := hd.DecodeInt64(src)
	if err != nil {
		return err
	}
	*h = HashID64(id)
	return nil
}
