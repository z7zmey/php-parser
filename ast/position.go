package ast

type Position struct {
	PS int `json:"pos_start"`
	PE int `json:"pos_end"`
	LS int `json:"line_start"`
	LE int `json:"line_end"`
}
