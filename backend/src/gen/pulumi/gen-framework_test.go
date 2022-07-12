package pulumi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObject_SimpleBuilder(t *testing.T) {
	const expected = `{
    name: "pulumi-fr-1",
    provisioner: "kubernetes"
}`

	root := newObject()
	root.addField("name", `"pulumi-fr-1"`)
	root.addField("provisioner", `"kubernetes"`)

	assert.Equal(t, expected, root.string(0))
}

func TestObject_BuilderWithNestedObject(t *testing.T) {
	const expected = `{
    framework: {
        name: "pulumi-fr-1",
        provisioner: "kubernetes"
    }
}`

	root := newObject()
	framework := newObject()

	root.addObject("framework", framework)

	framework.addField("name", `"pulumi-fr-1"`)
	framework.addField("provisioner", `"kubernetes"`)

	assert.Equal(t, expected, root.string(0))
}

func TestObject_BuilderWithNestedObjects(t *testing.T) {
	const expected = `{
    framework: {
        name: "pulumi-fr-1",
        provisioner: "kubernetes",
        resources: {
            general: {
                setup: {
                    public: true,
                    default: false
                },
                plan: {
                    name: "shipa-plan"
                }
            }
        }
    }
}`

	root := newObject()
	framework := newObject()

	root.addObject("framework", framework)
	framework.addField("name", `"pulumi-fr-1"`)
	framework.addField("provisioner", `"kubernetes"`)

	resources := newObject()
	framework.addObject("resources", resources)
	general := newObject()
	resources.addObject("general", general)

	general.addObject("setup",
		newObject().
			addField("public", "true").
			addField("default", "false"))

	general.addObject("plan", newObject().addField("name", `"shipa-plan"`))

	assert.Equal(t, expected, root.String())
}

func TestObject_BuilderWithListOfNestedObjects(t *testing.T) {
	const expected = `{
    framework: {
        name: "pulumi-fr-1",
        provisioner: "kubernetes",
        resources: [
            {
                name: "res1"
            },
            {
                name: "res2"
            }
        ]
    }
}`

	root := newObject()
	framework := newObject()

	root.addObject("framework", framework)
	framework.addField("name", `"pulumi-fr-1"`)
	framework.addField("provisioner", `"kubernetes"`)

	framework.addListOfObjects("resources", []*Object{
		newObject().addField("name", `"res1"`),
		newObject().addField("name", `"res2"`),
	})

	assert.Equal(t, expected, root.String())
}
