package odoo

import "fmt"

// PosSession represents pos.session model with a comprehensive list of fields.
type PosSession struct {
	// --- Trường cơ bản ---
	Id                   *Int       `xmlrpc:"id,omitempty"`
	Name                 *String    `xmlrpc:"name,omitempty"`
	State                *Selection `xmlrpc:"state,omitempty"`
	AccessToken          *String    `xmlrpc:"access_token,omitempty"`
	OpeningNotes         *String    `xmlrpc:"opening_notes,omitempty"`
	ClosingNotes         *String    `xmlrpc:"closing_notes,omitempty"`
	Rescue               *Bool      `xmlrpc:"rescue,omitempty"`
	UpdateStockAtClosing *Bool      `xmlrpc:"update_stock_at_closing,omitempty"`
	StartAt              *Time      `xmlrpc:"start_at,omitempty"`
	StopAt               *Time      `xmlrpc:"stop_at,omitempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omitempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omitempty"`

	// --- Trường quan hệ (Many2One) ---
	ConfigId      *Many2One `xmlrpc:"config_id,omitempty"`
	UserId        *Many2One `xmlrpc:"user_id,omitempty"`
	MoveId        *Many2One `xmlrpc:"move_id,omitempty"`
	CashJournalId *Many2One `xmlrpc:"cash_journal_id,omitempty"`
	EmployeeId    *Many2One `xmlrpc:"employee_id,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`

	// --- Trường số (Integer & Float) ---
	SequenceNumber             *Int   `xmlrpc:"sequence_number,omitempty"`
	LoginNumber                *Int   `xmlrpc:"login_number,omitempty"`
	CashRegisterBalanceEndReal *Float `xmlrpc:"cash_register_balance_end_real,omitempty"`
	CashRegisterBalanceStart   *Float `xmlrpc:"cash_register_balance_start,omitempty"`
	CashRealTransaction        *Float `xmlrpc:"cash_real_transaction,omitempty"`
}

// PosSessions represents array of pos.session model.
type PosSessions []PosSession

// PosSessionModel is the odoo model name.
const PosSessionModel = "pos.session"

// GetPosSession gets a single pos.session record from Odoo by its ID.
func (c *Client) GetPosSession(id int64) (*PosSession, error) {
	pss, err := c.GetPosSessions([]int64{id})
	if err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("không tìm thấy phiên POS với ID %d", id)
}

// GetPosSessions gets a list of pos.session records from Odoo by their IDs.
func (c *Client) GetPosSessions(ids []int64) (*PosSessions, error) {
	pss := &PosSessions{}
	if err := c.Read(PosSessionModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosSessions finds pos.session records by querying it
// with criteria and options.
func (c *Client) FindPosSessions(criteria *Criteria, options *Options) (*PosSessions, error) {
	pss := &PosSessions{}
	if err := c.SearchRead(PosSessionModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosSession finds a single pos.session record by querying it with criteria.
func (c *Client) FindPosSession(criteria *Criteria) (*PosSession, error) {
	pss, err := c.FindPosSessions(criteria, NewOptions().Limit(1))
	if err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("không tìm thấy phiên POS nào khớp với điều kiện")
}
