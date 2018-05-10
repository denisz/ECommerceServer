package api

import (
	. "store/models"
	"github.com/asdine/storm/q"
	"github.com/asdine/storm"
)
type ControllerBatches struct {
	Controller
}


func (p *ControllerBatches) GetAllBatches(pagination Pagination) (*PageBatches, error) {
	matcher := q.Gte("ID", 0)

	var batches []Batch
	err := p.GetStore().
		From(NodeNamedBatches).
		Select(matcher).
		Limit(pagination.Limit).
		Skip(pagination.Offset).
		OrderBy("CreatedAt").
		Reverse().
		Find(&batches)

	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	total, err := p.GetStore().
		From(NodeNamedBatches).
		Select(matcher).
		Count(new(Batch))

	if err != nil {
		return nil, err
	}

	return &PageBatches{
		Content: batches,
		Cursor: Cursor{
			Total:  total,
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
		},
	}, nil
}

func (p *ControllerBatches) FormsForBatches() {

}
