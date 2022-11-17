package odoo

import (
	"fmt"
)

// ProductAttributePrice represents product.attribute.price model.
type ProductAttributePrice struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	PriceExtra    *Float    `xmlrpc:"price_extra,omitempty"`
	ProductTmplId *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	ValueId       *Many2One `xmlrpc:"value_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductAttributePrices represents array of product.attribute.price model.
type ProductAttributePrices []ProductAttributePrice

// ProductAttributePriceModel is the odoo model name.
const ProductAttributePriceModel = "product.attribute.price"

// Many2One convert ProductAttributePrice to *Many2One.
func (pap *ProductAttributePrice) Many2One() *Many2One {
	return NewMany2One(pap.Id.Get(), "")
}

// CreateProductAttributePrice creates a new product.attribute.price model and returns its id.
func (c *Client) CreateProductAttributePrice(pap *ProductAttributePrice) (int64, error) {
	return c.Create(ProductAttributePriceModel, pap)
}

// UpdateProductAttributePrice updates an existing product.attribute.price record.
func (c *Client) UpdateProductAttributePrice(pap *ProductAttributePrice) error {
	return c.UpdateProductAttributePrices([]int64{pap.Id.Get()}, pap)
}

// UpdateProductAttributePrices updates existing product.attribute.price records.
// All records (represented by ids) will be updated by pap values.
func (c *Client) UpdateProductAttributePrices(ids []int64, pap *ProductAttributePrice) error {
	return c.Update(ProductAttributePriceModel, ids, pap)
}

// DeleteProductAttributePrice deletes an existing product.attribute.price record.
func (c *Client) DeleteProductAttributePrice(id int64) error {
	return c.DeleteProductAttributePrices([]int64{id})
}

// DeleteProductAttributePrices deletes existing product.attribute.price records.
func (c *Client) DeleteProductAttributePrices(ids []int64) error {
	return c.Delete(ProductAttributePriceModel, ids)
}

// GetProductAttributePrice gets product.attribute.price existing record.
func (c *Client) GetProductAttributePrice(id int64) (*ProductAttributePrice, error) {
	paps, err := c.GetProductAttributePrices([]int64{id})
	if err != nil {
		return nil, err
	}
	if paps != nil && len(*paps) > 0 {
		return &((*paps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.attribute.price not found", id)
}

// GetProductAttributePrices gets product.attribute.price existing records.
func (c *Client) GetProductAttributePrices(ids []int64) (*ProductAttributePrices, error) {
	paps := &ProductAttributePrices{}
	if err := c.Read(ProductAttributePriceModel, ids, nil, paps); err != nil {
		return nil, err
	}
	return paps, nil
}

// FindProductAttributePrice finds product.attribute.price record by querying it with criteria.
func (c *Client) FindProductAttributePrice(criteria *Criteria) (*ProductAttributePrice, error) {
	paps := &ProductAttributePrices{}
	if err := c.SearchRead(ProductAttributePriceModel, criteria, NewOptions().Limit(1), paps); err != nil {
		return nil, err
	}
	if paps != nil && len(*paps) > 0 {
		return &((*paps)[0]), nil
	}
	return nil, fmt.Errorf("no product.attribute.price was found with criteria %v", criteria)
}

// FindProductAttributePrices finds product.attribute.price records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributePrices(criteria *Criteria, options *Options) (*ProductAttributePrices, error) {
	paps := &ProductAttributePrices{}
	if err := c.SearchRead(ProductAttributePriceModel, criteria, options, paps); err != nil {
		return nil, err
	}
	return paps, nil
}

// FindProductAttributePriceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductAttributePriceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductAttributePriceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductAttributePriceId finds record id by querying it with criteria.
func (c *Client) FindProductAttributePriceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductAttributePriceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no product.attribute.price was found with criteria %v and options %v", criteria, options)
}
