package pagination

import (
	"math"
	"github.com/astaxie/beego/logs"
)

type Paginator struct {
	Rows  int //每页的行数
	Total int //总页数
	Page  int //当前页
	Nums  int //总行数
}

func (this *Paginator) SetRows(rows int) {
	this.Rows = rows
}

func (this *Paginator) GetTotalPages() int {
	if this.Total != 0 {
		return this.Total
	}
	logs.Debug("the Nums is %d,the rows is %d", this.Nums, this.Rows)
	pageNums := math.Ceil(float64(this.Nums) / float64(this.Rows))
	this.Total = int(pageNums)
	return this.Total
}

func (this *Paginator) SetNums(nums int) {
	this.Nums = nums
}

func (this *Paginator) Offset() int {
	return (this.Page - 1) * this.Rows
}

func NewPaginator(page int, pageSize int) *Paginator {
	p := Paginator{}
	if pageSize <= 0 {
		p.Rows = 10
	}
	p.Rows = pageSize
	p.Page = page
	return &p
}

