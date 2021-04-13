package main

type EntityRecord struct {
	showIndex int64
	index     int64
	count     int
}

func (p *EntityRecord) CheckField(index int, field string) {
	switch field {
	case "id", "create_by", "create_date", "update_by", "update_date", "del_flag":
		if field == "id" || field == "update_date" {
			p.showIndex |= 1 << index
		}
		p.index |= 1 << index
		p.count++
	default:
		return
	}
}

func (p *EntityRecord) IsCommonField(index int) bool {
	return p.index&(1<<index) > 0
}

func (p *EntityRecord) IsShow(index int) bool {
	return p.showIndex&(1<<index) > 0
}

func (p *EntityRecord) HasCommon() bool {
	return p.count == 6
}
