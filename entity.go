package main

type EntityRecord struct {
	FieldIndex []int
}

func (p *EntityRecord) AddField(field string, index int) {
	switch field {
	case "id", "create_by", "create_date", "update_by", "update_date", "del_flag":
		p.FieldIndex = append(p.FieldIndex, index)
	default:
		return
	}
}

func (p *EntityRecord) HasBase() bool {
	return len(p.FieldIndex) == 6
}
