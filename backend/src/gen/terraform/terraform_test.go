package terraform

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObject_String(t *testing.T) {
	const expected = `{
  framework {
    name = "cinema-payment"
    public = false
    access {
      append = ["shipa-team", "shipa-team2"]
    }
  }
}`

	framework := newObject().
		addField("name", stringValue("cinema-payment")).
		addField("public", strconv.FormatBool(false)).
		addObject("access",
			newObject().addField("append", stringArray([]string{"shipa-team", "shipa-team2"})))

	root := newObject()
	root.addObject("framework", framework)

	assert.Equal(t, expected, root.String())
}
