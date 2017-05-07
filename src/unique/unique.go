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
    id, err := uuid.GenerateUUID()

    if err != nil {
        fmt.Println("error generating uuid: ", err);
    }

    return id
}

// CreateTSID generates a 12 character time-stamp / base 36 id
func CreateTSID() string {
    id := strconv.FormatInt(time.Now().UnixNano(), 36)

    return id
}
