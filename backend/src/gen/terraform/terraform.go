package terraform

import (
	"fmt"
	"strings"
)

func genIndent(indentationLevel int) string {
	if indentationLevel <= 0 {
		return ""
	}

	const spacesInTab = 2
	return strings.Repeat(" ", indentationLevel*spacesInTab)
}

type Object struct {
	fields    []*Field
	dependsOn string
}

type Field struct {
	name    string
	content *string
	object  *Object
	list    []*Object
}

func (obj *Object) addField(name, content string) *Object {
	obj.fields = append(obj.fields, &Field{
		name:    name,
		content: &content,
	})
	return obj
}

func (obj *Object) addObject(name string, child *Object) *Object {
	obj.fields = append(obj.fields, &Field{
		name:   name,
		object: child,
	})
	return obj
}

func (obj *Object) addListOfObjects(name string, list []*Object) *Object {
	obj.fields = append(obj.fields, &Field{
		name: name,
		list: list,
	})
	return obj
}

func (f *Field) string(indentLevel int) string {
	if f.content != nil {
		return fmt.Sprintf("%s%s = %s", genIndent(indentLevel), f.name, *f.content)
	}

	if f.object != nil {
		return fmt.Sprintf("%s%s %s", genIndent(indentLevel), f.name, f.object.string(indentLevel))
	}

	if f.list != nil {
		var listContent []string
		for _, obj := range f.list {
			listContent = append(listContent, genIndent(indentLevel+1)+obj.string(indentLevel+1))
		}

		return fmt.Sprintf(`%s%s: [
%s
%s]`, genIndent(indentLevel), f.name, strings.Join(listContent, ",\n"), genIndent(indentLevel))
	}

	return ""
}

func (obj *Object) String() string {
	return obj.string(0)
}

func (obj *Object) string(indentLevel int) string {
	var fields []string
	for _, f := range obj.fields {
		fields = append(fields, f.string(indentLevel+1))
	}

	if obj.dependsOn != "" {
		fields = append(fields, fmt.Sprintf("%s%s", genIndent(indentLevel+1), obj.dependsOn))
	}

	innerContent := strings.Join(fields, "\n")

	return fmt.Sprintf(`{
%s
%s}`, innerContent, genIndent(indentLevel))
}

func newObject() *Object {
	return &Object{}
}

func stringValue(val string) string {
	return fmt.Sprintf(`"%s"`, val)
}

func stringArray(values []string) string {
	return fmt.Sprintf(`["%s"]`, strings.Join(values, `", "`))
}
