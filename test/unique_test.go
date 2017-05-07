//
// unique tests
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-07 12:18:25
//

package unit

import (
	"fmt"
	"testing"

	"unique"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Unique", func() {

		g.It("should create a 26 character ULID", func() {
			ulid := unique.CreateULID()

			fmt.Println(ulid)

			g.Assert(len(ulid)).Equal(26)

			// should put this in a map
			for i := 0; i < 10; i++ {
				g.Assert(unique.CreateULID() != ulid).IsTrue()
			}
		})

		g.It("should create a 36 character UUID", func() {
            uuid := unique.CreateUUID()

            fmt.Println(uuid)

            g.Assert(len(uuid)).Equal(36)
        })

        g.It("should create a 12 chanacter TSID", func() {
            tsid := unique.CreateTSID()

            fmt.Println(tsid)

            g.Assert(len(tsid)).Equal(12)
        })
	})
}
