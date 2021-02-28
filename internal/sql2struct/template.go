package sql2struct

import (
	"NanjoFan/Tour/internal/word"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
)

const structTpl = `package model

type {{.TableName | ToCamelCase}} struct {
{{- range .Columns}}{{ $typeLen := len .Type }} {{- if gt $typeLen 0 }}
	{{.Name | ToCamelCase}} {{.Type}}    {{.Tag}}{{ else }}{{.Name}}{{- end }}{{ $length := len .Comment}} {{- if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{- end}}
    {{- end}} 
}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

// 存储转换后的Go结构体中的所有字段信息
type StructColumn struct {
	Name string
	Type string
	Tag string
	Comment string
}

// 存储最终用于渲染的模板对象信息
type StructTemplateDB struct {
	TableName string
	Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

// 装配列 TableColumn 从数据库读取的Table列的信息
func (t *StructTemplate)AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     fmt.Sprintf("`json:" + "%s"+"`", column.ColumnName),
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

// Generate 生成
func (t *StructTemplate)Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.
		FuncMap{"ToCamelCase": word.UnderscoreToUpperCamelCase,}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}

	file, err := os.Create("file.go")
	if err != nil {
		log.Fatalln("os.Create failed, err is", err)
	}
	defer file.Close()

	err = tpl.Execute(file, tplDB)
	if err != nil {
		return err
	}

	cmd := exec.Command("go","fmt", "file.go")
	cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
