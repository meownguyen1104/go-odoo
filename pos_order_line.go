package odoo

// PosOrderLine represents pos.order.line model.
type PosOrderLine struct {
	AttributeValueIds         *Relation  `xmlrpc:"attribute_value_ids,omitempty"`
	ComboItemId               *Many2One  `xmlrpc:"combo_item_id,omitempty"`
	ComboLineIds              *Relation  `xmlrpc:"combo_line_ids,omitempty"`
	ComboParentId             *Many2One  `xmlrpc:"combo_parent_id,omitempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                *Many2One  `xmlrpc:"currency_id,omitempty"`
	CustomAttributeValueIds   *Relation  `xmlrpc:"custom_attribute_value_ids,omitempty"`
	CustomerNote              *String    `xmlrpc:"customer_note,omitempty"`
	Discount                  *Float     `xmlrpc:"discount,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	FullProductName           *String    `xmlrpc:"full_product_name,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	IsEdited                  *Bool      `xmlrpc:"is_edited,omitempty"`
	IsTotalCostComputed       *Bool      `xmlrpc:"is_total_cost_computed,omitempty"`
	Margin                    *Float     `xmlrpc:"margin,omitempty"`
	MarginPercent             *Float     `xmlrpc:"margin_percent,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	Note                      *String    `xmlrpc:"note,omitempty"`
	Notice                    *String    `xmlrpc:"notice,omitempty"`
	OrderId                   *Many2One  `xmlrpc:"order_id,omitempty"`
	PackLotIds                *Relation  `xmlrpc:"pack_lot_ids,omitempty"`
	PriceExtra                *Float     `xmlrpc:"price_extra,omitempty"`
	PriceSubtotal             *Float     `xmlrpc:"price_subtotal,omitempty"`
	PriceSubtotalIncl         *Float     `xmlrpc:"price_subtotal_incl,omitempty"`
	PriceType                 *Selection `xmlrpc:"price_type,omitempty"`
	PriceUnit                 *Float     `xmlrpc:"price_unit,omitempty"`
	ProductId                 *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductUomId              *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	Qty                       *Float     `xmlrpc:"qty,omitempty"`
	RefundOrderlineIds        *Relation  `xmlrpc:"refund_orderline_ids,omitempty"`
	RefundedOrderlineId       *Many2One  `xmlrpc:"refunded_orderline_id,omitempty"`
	RefundedQty               *Float     `xmlrpc:"refunded_qty,omitempty"`
	SkipChange                *Bool      `xmlrpc:"skip_change,omitempty"`
	TaxIds                    *Relation  `xmlrpc:"tax_ids,omitempty"`
	TaxIdsAfterFiscalPosition *Relation  `xmlrpc:"tax_ids_after_fiscal_position,omitempty"`
	TotalCost                 *Float     `xmlrpc:"total_cost,omitempty"`
	Uuid                      *String    `xmlrpc:"uuid,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PosOrderLines represents array of pos.order.line model.
type PosOrderLines []PosOrderLine

// PosOrderLineModel is the odoo model name.
const PosOrderLineModel = "pos.order.line"

// Many2One convert PosOrderLine to *Many2One.
func (pol *PosOrderLine) Many2One() *Many2One {
	return NewMany2One(pol.Id.Get(), "")
}

// CreatePosOrderLine creates a new pos.order.line model and returns its id.
func (c *Client) CreatePosOrderLine(pol *PosOrderLine) (int64, error) {
	ids, err := c.CreatePosOrderLines([]*PosOrderLine{pol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosOrderLine creates a new pos.order.line model and returns its id.
func (c *Client) CreatePosOrderLines(pols []*PosOrderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range pols {
		vv = append(vv, v)
	}
	return c.Create(PosOrderLineModel, vv, nil)
}

// UpdatePosOrderLine updates an existing pos.order.line record.
func (c *Client) UpdatePosOrderLine(pol *PosOrderLine) error {
	return c.UpdatePosOrderLines([]int64{pol.Id.Get()}, pol)
}

// UpdatePosOrderLines updates existing pos.order.line records.
// All records (represented by ids) will be updated by pol values.
func (c *Client) UpdatePosOrderLines(ids []int64, pol *PosOrderLine) error {
	return c.Update(PosOrderLineModel, ids, pol, nil)
}

// DeletePosOrderLine deletes an existing pos.order.line record.
func (c *Client) DeletePosOrderLine(id int64) error {
	return c.DeletePosOrderLines([]int64{id})
}

// DeletePosOrderLines deletes existing pos.order.line records.
func (c *Client) DeletePosOrderLines(ids []int64) error {
	return c.Delete(PosOrderLineModel, ids)
}

// GetPosOrderLine gets pos.order.line existing record.
func (c *Client) GetPosOrderLine(id int64) (*PosOrderLine, error) {
	pols, err := c.GetPosOrderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pols)[0]), nil
}

// GetPosOrderLines gets pos.order.line existing records.
func (c *Client) GetPosOrderLines(ids []int64) (*PosOrderLines, error) {
	pols := &PosOrderLines{}
	if err := c.Read(PosOrderLineModel, ids, nil, pols); err != nil {
		return nil, err
	}
	return pols, nil
}

// FindPosOrderLine finds pos.order.line record by querying it with criteria.
func (c *Client) FindPosOrderLine(criteria *Criteria) (*PosOrderLine, error) {
	pols := &PosOrderLines{}
	if err := c.SearchRead(PosOrderLineModel, criteria, NewOptions().Limit(1), pols); err != nil {
		return nil, err
	}
	return &((*pols)[0]), nil
}

// FindPosOrderLines finds pos.order.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOrderLines(criteria *Criteria, options *Options) (*PosOrderLines, error) {
	pols := &PosOrderLines{}
	if err := c.SearchRead(PosOrderLineModel, criteria, options, pols); err != nil {
		return nil, err
	}
	return pols, nil
}

// FindPosOrderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOrderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosOrderLineModel, criteria, options)
}

// FindPosOrderLineId finds record id by querying it with criteria.
func (c *Client) FindPosOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
