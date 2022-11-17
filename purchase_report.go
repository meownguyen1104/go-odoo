package odoo

import (
	"fmt"
)

// PurchaseReport represents purchase.report model.
type PurchaseReport struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	AccountAnalyticId   *Many2One  `xmlrpc:"account_analytic_id,omitempty"`
	CategoryId          *Many2One  `xmlrpc:"category_id,omitempty"`
	CommercialPartnerId *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omitempty"`
	CurrencyId          *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateApprove         *Time      `xmlrpc:"date_approve,omitempty"`
	DateOrder           *Time      `xmlrpc:"date_order,omitempty"`
	Delay               *Float     `xmlrpc:"delay,omitempty"`
	DelayPass           *Float     `xmlrpc:"delay_pass,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	FiscalPositionId    *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	NbrLines            *Int       `xmlrpc:"nbr_lines,omitempty"`
	Negociation         *Float     `xmlrpc:"negociation,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	PickingTypeId       *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	PriceAverage        *Float     `xmlrpc:"price_average,omitempty"`
	PriceStandard       *Float     `xmlrpc:"price_standard,omitempty"`
	PriceTotal          *Float     `xmlrpc:"price_total,omitempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTmplId       *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUom          *Many2One  `xmlrpc:"product_uom,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	UnitQuantity        *Float     `xmlrpc:"unit_quantity,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty"`
	Volume              *Float     `xmlrpc:"volume,omitempty"`
	Weight              *Float     `xmlrpc:"weight,omitempty"`
}

// PurchaseReports represents array of purchase.report model.
type PurchaseReports []PurchaseReport

// PurchaseReportModel is the odoo model name.
const PurchaseReportModel = "purchase.report"

// Many2One convert PurchaseReport to *Many2One.
func (pr *PurchaseReport) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreatePurchaseReport creates a new purchase.report model and returns its id.
func (c *Client) CreatePurchaseReport(pr *PurchaseReport) (int64, error) {
	return c.Create(PurchaseReportModel, pr)
}

// UpdatePurchaseReport updates an existing purchase.report record.
func (c *Client) UpdatePurchaseReport(pr *PurchaseReport) error {
	return c.UpdatePurchaseReports([]int64{pr.Id.Get()}, pr)
}

// UpdatePurchaseReports updates existing purchase.report records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdatePurchaseReports(ids []int64, pr *PurchaseReport) error {
	return c.Update(PurchaseReportModel, ids, pr)
}

// DeletePurchaseReport deletes an existing purchase.report record.
func (c *Client) DeletePurchaseReport(id int64) error {
	return c.DeletePurchaseReports([]int64{id})
}

// DeletePurchaseReports deletes existing purchase.report records.
func (c *Client) DeletePurchaseReports(ids []int64) error {
	return c.Delete(PurchaseReportModel, ids)
}

// GetPurchaseReport gets purchase.report existing record.
func (c *Client) GetPurchaseReport(id int64) (*PurchaseReport, error) {
	prs, err := c.GetPurchaseReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of purchase.report not found", id)
}

// GetPurchaseReports gets purchase.report existing records.
func (c *Client) GetPurchaseReports(ids []int64) (*PurchaseReports, error) {
	prs := &PurchaseReports{}
	if err := c.Read(PurchaseReportModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPurchaseReport finds purchase.report record by querying it with criteria.
func (c *Client) FindPurchaseReport(criteria *Criteria) (*PurchaseReport, error) {
	prs := &PurchaseReports{}
	if err := c.SearchRead(PurchaseReportModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("no purchase.report was found with criteria %v", criteria)
}

// FindPurchaseReports finds purchase.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseReports(criteria *Criteria, options *Options) (*PurchaseReports, error) {
	prs := &PurchaseReports{}
	if err := c.SearchRead(PurchaseReportModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPurchaseReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PurchaseReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPurchaseReportId finds record id by querying it with criteria.
func (c *Client) FindPurchaseReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no purchase.report was found with criteria %v and options %v", criteria, options)
}
