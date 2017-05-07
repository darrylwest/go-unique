//
// unique -
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-07 12:22:45
//

package unique

import (
	crand "crypto/rand"
	"fmt"
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

// RandomBytes generates a byte buffer of the specified size and populates it with crypo-strength random bytes
func RandomBytes(size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := crand.Read(buf)

	return buf, err
}

func v4uuid() []byte {
	buf, _ := RandomBytes(16)

	return buf
}

// CreateUUID generates and returns a uuid as a string
func CreateUUID() string {
	buf := v4uuid()
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:16])
}

// CreateGUID generates and returns a uuid as a string
func CreateGUID() string {
	return fmt.Sprintf("%x", v4uuid())
}

// CreateTSID generates a 12 character time-stamp / base 36 id
func CreateTSID() string {
	id := strconv.FormatInt(time.Now().UnixNano(), 36)

	return id
}

// CreateTXID generates a 16 character time-stamp / base 36 id
func CreateTXID() string {
	id := strconv.FormatInt(time.Now().UnixNano(), 36)
	buf, _ := RandomBytes(2)
	str := fmt.Sprintf("%s%x", id, buf)

	return str
}
