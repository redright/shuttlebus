package common

import (
	"github.com/redright/shuttlebus/db"
)

type PM struct {
}

func (p *PM) GetParameters(groupCode string) []Parameter {
	r := db.Query("select ParameterCode,ParameterValue from parameter where GroupCode=?", groupCode)
	var result []Parameter
	for r.Next() {
		var row Parameter
		r.Scan(&row.ParameterCode, &row.ParameterValue)
	}
	return result
}

type Parameter struct {
	ParameterCode  string
	ParameterValue string
}
