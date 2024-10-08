/**
 ******************************************************************************
 * @file           : dao_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package dao

import (
	"gorm.io/gen"
	"strings"
	"testing"
)

func TestGenGormMutiple(ts *testing.T) {
	for _, t := range types {
		name := t.Name()
		println(name)
		g := gen.NewGenerator(gen.Config{
			OutPath:       "dao/query/" + strings.ToLower(name) + "model",
			Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
			FieldNullable: true,
		})

		g.ApplyBasic(getStructInstance(t))
		if v, ok := mps[getStructName(t)]; ok {
			for _, m := range v {
				g.ApplyInterface(m, getStructInstance(t))
			}

		}
		//g.ApplyInterface(func(model.Method) {}, model.User{})
		//g.ApplyInterface(func(model.UserMethod) {}, model.User{})

		g.Execute()
	}
}
