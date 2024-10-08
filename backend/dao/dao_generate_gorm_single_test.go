/**
 ******************************************************************************
 * @file           : dao_generate_gorm_single_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package dao

import (
	"gorm.io/gen"
	"testing"
)

func TestGormGenSingle(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "dao/query",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	for _, t := range types {
		name := t.Name()
		println(name)

		g.ApplyBasic(getStructInstance(t))
		if v, ok := mps[getStructName(t)]; ok {
			for _, m := range v {
				g.ApplyInterface(m, getStructInstance(t))
			}

		}
	}

	//g.ApplyInterface(func(model.Method) {}, model.User{})
	//g.ApplyInterface(func(model.UserMethod) {}, model.User{})

	g.Execute()
}
