//
// unique -
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-07 12:22:45
//

package unique

import (
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/oklog/ulid"
	"io"
	"math/rand"
	"strconv"
	"time"
)

var (
	entropy io.Reader = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func genulid(entropy io.Reader, ts uint64) (ulid.ULID, error) {
	value, err := ulid.New(ts, entropy)
	return value, err
}

// CreateRawULID create a raw ulid
func CreateRawULID() ulid.ULID {
	ts := uint64(time.Now().UnixNano() / 1000000)
	v, _ := genulid(entropy, ts)

	return v
}

// CreateULID generate and return a ulid as a string
func CreateULID() string {
	return CreateRawULID().String()
}

// IDType 26 bytes
type IDType [26]byte

// CreateUUID generates and returns a uuid as a string
func CreateUUID() string {
	id, _ := uuid.GenerateUUID()

	return id
}

// CreateGUID generates and returns a uuid as a string
func CreateGUID() string {
	buf, _ := uuid.GenerateRandomBytes(16)

	return fmt.Sprintf("%x", buf)
}

// CreateTSID generates a 12 character time-stamp / base 36 id
func CreateTSID() string {
	id := strconv.FormatInt(time.Now().UnixNano(), 36)

	return id
}

// CreateTXID generates a 16 character time-stamp / base 36 id
func CreateTXID() string {
	id := strconv.FormatInt(time.Now().UnixNano(), 36)
	buf, _ := uuid.GenerateRandomBytes(2)
	str := fmt.Sprintf("%s%x", id, buf)

	return str
}
