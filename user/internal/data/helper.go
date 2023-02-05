package data

import (
	"bytes"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type PageHelper struct {
	db      *gorm.DB
	Page    int64
	Size    int64
	MaxPage int64
	Count   int64
}

func NewPageHelper(db *gorm.DB, page int64, size int64) *PageHelper {
	return &PageHelper{db: db, Page: page, Size: size}
}

func (p *PageHelper) List(dest interface{}) (err error) {
	var c int64
	err = p.db.Count(&c).Error
	if err != nil {
		return
	}
	p.Count = c
	if c == 0 {
		return nil
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Size == 0 {
		p.Size = 10
	}
	ps := c / p.Size
	if c%p.Size == 0 {
		p.MaxPage = ps
	} else {
		p.MaxPage = ps + 1
	}
	off := int((p.Page - 1) * p.Size)
	err = p.db.Offset(off).Limit(int(p.Size)).Find(dest).Error
	return
}

func Where(db *gorm.DB, c map[string]interface{}, or ...string) *gorm.DB {
	var bf bytes.Buffer
	var sta = " %s "
	if len(or) == 0 {
		sta = fmt.Sprintf(sta, "and")
	} else {
		sta = fmt.Sprintf(sta, "or")
	}
	ln := len(sta)
	ps := make([]interface{}, 0)
	for k, v := range c {
		switch v.(type) {
		case string:
			if v.(string) == "" {
				continue
			}
		case int:
			if v.(int) == 0 {
				continue
			}
		case int32:
			if v.(int32) == 0 {
				continue
			}
		case int64:
			if v.(int64) == 0 {
				continue
			}
		case float64:
			if v.(float64) == 0 {
				continue
			}
		case float32:
			if v.(float32) == 0 {
				continue
			}
		case time.Time:
			if v.(time.Time).IsZero() {
				continue
			}
		case QueryBool:
			if v.(QueryBool) == -1 {
				continue
			}
		}
		bf.WriteString(k)
		bf.WriteString(sta)
		ps = append(ps, v)
	}
	if bf.Len() == 0 {
		return db
	}
	bf.Truncate(bf.Len() - ln)
	return db.Where(bf.String(), ps...)
}

type QueryBool int
