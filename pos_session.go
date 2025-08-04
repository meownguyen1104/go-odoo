package odoo

// PosSession represents pos.session model.
type PosSession struct {
	Id                           *Int        `xmlrpc:"id,omitempty"`
	ConfigId                     *Many2One   `xmlrpc:"config_id,omitempty"`
	UserId                       *Many2One   `xmlrpc:"user_id,omitempty"`
	SequenceNumber               *Int        `xmlrpc:"sequence_number,omitempty"`
	LoginNumber                  *Int        `xmlrpc:"login_number,omitempty"`
	CashJournalId                *Many2One   `xmlrpc:"cash_journal_id,omitempty"`
	MoveId                       *Many2One   `xmlrpc:"move_id,omitempty"`
	CreateUid                    *Many2One   `xmlrpc:"create_uid,omitempty"`
	WriteUid                     *Many2One   `xmlrpc:"write_uid,omitempty"`
	AccessToken                  *String     `xmlrpc:"access_token,omitempty"`
	Name                         *String     `xmlrpc:"name,omitempty"`
	State                        *Selection  `xmlrpc:"state,omitempty"`
	OpeningNotes                 *String     `xmlrpc:"opening_notes,omitempty"`
	ClosingNotes                 *String     `xmlrpc:"closing_notes,omitempty"`
	CashRegisterBalanceEndReal   *Float      `xmlrpc:"cash_register_balance_end_real,omitempty"`
	CashRegisterBalanceStart     *Float      `xmlrpc:"cash_register_balance_start,omitempty"`
	CashRealTransaction          *Float      `xmlrpc:"cash_real_transaction,omitempty"`
	Rescue                       *Bool       `xmlrpc:"rescue,omitempty"`
	UpdateStockAtClosing         *Bool       `xmlrpc:"update_stock_at_closing,omitempty"`
	StartAt                      *Time       `xmlrpc:"start_at,omitempty"`
	StopAt                       *Time       `xmlrpc:"stop_at,omitempty"`
	CreateDate                   *Time       `xmlrpc:"create_date,omitempty"`
	WriteDate                    *Time       `xmlrpc:"write_date,omitempty"`
	EmployeeId                   *Many2One   `xmlrpc:"employee_id,omitempty"`
}

// PosSessions represents array of pos.session model.
type PosSessions []PosSession

// PosSessionModel is the odoo model name.
const PosSessionModel = "pos.session"

// Many2One convert PosSession to *Many2One.
func (ps *PosSession) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePosSession creates a new pos.session model and returns its id.
func (c *Client) CreatePosSession(ps *PosSession) (int64, error) {
	ids, err := c.CreatePosSessions([]*PosSession{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePosSessions creates a new pos.session model and returns its id.
func (c *Client) CreatePosSessions(pss []*PosSession) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(PosSessionModel, vv, nil)
}

// UpdatePosSession updates an existing pos.session record.
func (c *Client) UpdatePosSession(ps *PosSession) error {
	return c.UpdatePosSessions([]int64{ps.Id.Get()}, ps)
}

// UpdatePosSessions updates existing pos.session records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePosSessions(ids []int64, ps *PosSession) error {
	return c.Update(PosSessionModel, ids, ps, nil)
}

// DeletePosSession deletes an existing pos.session record.
func (c *Client) DeletePosSession(id int64) error {
	return c.DeletePosSessions([]int64{id})
}

// DeletePosSessions deletes existing pos.session records.
func (c *Client) DeletePosSessions(ids []int64) error {
	return c.Delete(PosSessionModel, ids)
}

// GetPosSession gets pos.session existing record.
func (c *Client) GetPosSession(id int64) (*PosSession, error) {
	pss, err := c.GetPosSessions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pss)[0]), nil
}

// GetPosSessions gets pos.session existing records.
func (c *Client) GetPosSessions(ids []int64) (*PosSessions, error) {
	pss := &PosSessions{}
	if err := c.Read(PosSessionModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosSession finds pos.session record by querying it with criteria.
func (c *Client) FindPosSession(criteria *Criteria) (*PosSession, error) {
	pss := &PosSessions{}
	if err := c.SearchRead(PosSessionModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	return &((*pss)[0]), nil
}

// FindPosSessions finds pos.session records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSessions(criteria *Criteria, options *Options) (*PosSessions, error) {
	pss := &PosSessions{}
	if err := c.SearchRead(PosSessionModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosSessionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSessionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PosSessionModel, criteria, options)
}

// FindPosSessionId finds record id by querying it with criteria.
func (c *Client) FindPosSessionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosSessionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
