package schema

import (
	"geeORM/dialect"
	"go/ast"
	"reflect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{} // 结构体本身
	Name       string      // 结构体名字(表名)
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name]
}

func (s *Schema) RecordValues(object interface{}) []interface{} {
	objectValue := reflect.Indirect(reflect.ValueOf(object))
	var fieldValues []interface{}
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, objectValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

func Parse(model interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(model)).Type()
	schema := &Schema{
		Model:    model,
		Name:     modelType.Name(),
		fieldMap: map[string]*Field{},
	}
	for i := 0; i < modelType.NumField(); i++ {
		f := modelType.Field(i)
		if !f.Anonymous && ast.IsExported(f.Name) {
			field := &Field{
				Name: f.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(f.Type))),
			}
			if v, ok := f.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, f.Name)
			schema.fieldMap[f.Name] = field
		}
	}
	return schema
}
