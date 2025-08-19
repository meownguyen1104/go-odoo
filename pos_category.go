package odoo

// PosCategory represents pos.category model.
type PosCategory struct {
	ChildIds    *Relation `xmlrpc:"child_ids,omitempty"`
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	HasImage    *Bool     `xmlrpc:"has_image,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Image128    *String   `xmlrpc:"image_128,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PosCategorys represents array of pos.category model.
type PosCategorys []PosCategory

// PosCategoryModel is the odoo model name.
const PosCategoryModel = "pos.category"

// Many2One convert PosCategory to *Many2One.
func (pc *PosCategory) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreatePosCategory creates a new pos.category model and returns its id.
func (c *Client) CreatePosCategory(pc *PosCategory) (int64, error) {
	ids, err := c.CreatePosCategorys([]*PosCategory{pc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosCategory creates a new pos.category model and returns its id.
func (c *Client) CreatePosCategorys(pcs []*PosCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcs {
		vv = append(vv, v)
	}
	return c.Create(PosCategoryModel, vv, nil)
}

// UpdatePosCategory updates an existing pos.category record.
func (c *Client) UpdatePosCategory(pc *PosCategory) error {
	return c.UpdatePosCategorys([]int64{pc.Id.Get()}, pc)
}

// UpdatePosCategorys updates existing pos.category records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdatePosCategorys(ids []int64, pc *PosCategory) error {
	return c.Update(PosCategoryModel, ids, pc, nil)
}

// DeletePosCategory deletes an existing pos.category record.
func (c *Client) DeletePosCategory(id int64) error {
	return c.DeletePosCategorys([]int64{id})
}

// DeletePosCategorys deletes existing pos.category records.
func (c *Client) DeletePosCategorys(ids []int64) error {
	return c.Delete(PosCategoryModel, ids)
}

// GetPosCategory gets pos.category existing record.
func (c *Client) GetPosCategory(id int64) (*PosCategory, error) {
	pcs, err := c.GetPosCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// GetPosCategorys gets pos.category existing records.
func (c *Client) GetPosCategorys(ids []int64) (*PosCategorys, error) {
	pcs := &PosCategorys{}
	if err := c.Read(PosCategoryModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosCategory finds pos.category record by querying it with criteria.
func (c *Client) FindPosCategory(criteria *Criteria) (*PosCategory, error) {
	pcs := &PosCategorys{}
	if err := c.SearchRead(PosCategoryModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	return &((*pcs)[0]), nil
}

// FindPosCategorys finds pos.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosCategorys(criteria *Criteria, options *Options) (*PosCategorys, error) {
	pcs := &PosCategorys{}
	if err := c.SearchRead(PosCategoryModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosCategoryModel, criteria, options)
}

// FindPosCategoryId finds record id by querying it with criteria.
func (c *Client) FindPosCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
