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

		g.It("should create a unique struct", func() {
			ulid := unique.CreateULID()

			fmt.Println(ulid)

			g.Assert(len(ulid)).Equal(26)

			// should put this in a map
			for i := 0; i < 10; i++ {
				g.Assert(unique.CreateULID() != ulid).IsTrue()
			}
		})

	})
}
