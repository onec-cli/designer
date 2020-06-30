package designer

import (
	"github.com/v8platform/marshaler"
)

// /DumpIB <имя файла>
//— выгрузка информационной базы в командном режиме.
type DumpIBOptions struct {
	Designer `v8:",inherit" json:"designer"`

	File string `v8:"/DumpIB" json:"file"`
}

func (d DumpIBOptions) Values() []string {

	v, _ := marshaler.Marshal(d)
	return v
}

// /RestoreIB <имя файла>
// — загрузка информационной базы в командном режиме.
// Если файл информационной базы отсутствует в указанном каталоге, будет создана новая информационная база.
type RestoreIBOptions struct {
	Designer `v8:",inherit" json:"designer"`

	File string `v8:"/RestoreIB" json:"file"`
}

func (d RestoreIBOptions) Values() []string {

	v, _ := marshaler.Marshal(d)
	return v
}
