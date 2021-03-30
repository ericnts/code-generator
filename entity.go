package main

type EntityRecord struct {
	index int64
	count int
}

func (p *EntityRecord) CheckField(index int, field string) {
	switch field {
	case "id", "create_by", "create_date", "update_by", "update_date", "del_flag":
		p.index |= 1 << index
		p.count++
	default:
		return
	}
}

func (p *EntityRecord) IsCommonField(index int) bool {
	return p.index&(1<<index) > 0
}

func (p *EntityRecord) HasCommon() bool {
	return p.count == 6
}
