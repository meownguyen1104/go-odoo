package odoo

// PosSession represents pos.session model.
type PosSession struct {
	AccessToken                 *String    `xmlrpc:"access_token,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	BankPaymentIds              *Relation  `xmlrpc:"bank_payment_ids,omitempty"`
	CashControl                 *Bool      `xmlrpc:"cash_control,omitempty"`
	CashJournalId               *Many2One  `xmlrpc:"cash_journal_id,omitempty"`
	CashRealTransaction         *Float     `xmlrpc:"cash_real_transaction,omitempty"`
	CashRegisterBalanceEnd      *Float     `xmlrpc:"cash_register_balance_end,omitempty"`
	CashRegisterBalanceEndReal  *Float     `xmlrpc:"cash_register_balance_end_real,omitempty"`
	CashRegisterBalanceStart    *Float     `xmlrpc:"cash_register_balance_start,omitempty"`
	CashRegisterDifference      *Float     `xmlrpc:"cash_register_difference,omitempty"`
	ClosingNotes                *String    `xmlrpc:"closing_notes,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	ConfigId                    *Many2One  `xmlrpc:"config_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	FailedPickings              *Bool      `xmlrpc:"failed_pickings,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsInCompanyCurrency         *Bool      `xmlrpc:"is_in_company_currency,omitempty"`
	LoginNumber                 *Int       `xmlrpc:"login_number,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MoveId                      *Many2One  `xmlrpc:"move_id,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	OpeningNotes                *String    `xmlrpc:"opening_notes,omitempty"`
	OrderCount                  *Int       `xmlrpc:"order_count,omitempty"`
	OrderIds                    *Relation  `xmlrpc:"order_ids,omitempty"`
	PaymentMethodIds            *Relation  `xmlrpc:"payment_method_ids,omitempty"`
	PickingCount                *Int       `xmlrpc:"picking_count,omitempty"`
	PickingIds                  *Relation  `xmlrpc:"picking_ids,omitempty"`
	Rescue                      *Bool      `xmlrpc:"rescue,omitempty"`
	SequenceNumber              *Int       `xmlrpc:"sequence_number,omitempty"`
	StartAt                     *Time      `xmlrpc:"start_at,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	StatementLineIds            *Relation  `xmlrpc:"statement_line_ids,omitempty"`
	StopAt                      *Time      `xmlrpc:"stop_at,omitempty"`
	TotalPaymentsAmount         *Float     `xmlrpc:"total_payments_amount,omitempty"`
	UpdateStockAtClosing        *Bool      `xmlrpc:"update_stock_at_closing,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
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

// CreatePosSession creates a new pos.session model and returns its id.
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
