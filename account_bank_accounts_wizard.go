package odoo

import (
	"fmt"
)

// AccountBankAccountsWizard represents account.bank.accounts.wizard model.
type AccountBankAccountsWizard struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	AccName       *String    `xmlrpc:"acc_name,omitempty"`
	AccountType   *Selection `xmlrpc:"account_type,omitempty"`
	BankAccountId *Many2One  `xmlrpc:"bank_account_id,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId    *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountBankAccountsWizards represents array of account.bank.accounts.wizard model.
type AccountBankAccountsWizards []AccountBankAccountsWizard

// AccountBankAccountsWizardModel is the odoo model name.
const AccountBankAccountsWizardModel = "account.bank.accounts.wizard"

// Many2One convert AccountBankAccountsWizard to *Many2One.
func (abaw *AccountBankAccountsWizard) Many2One() *Many2One {
	return NewMany2One(abaw.Id.Get(), "")
}

// CreateAccountBankAccountsWizard creates a new account.bank.accounts.wizard model and returns its id.
func (c *Client) CreateAccountBankAccountsWizard(abaw *AccountBankAccountsWizard) (int64, error) {
	return c.Create(AccountBankAccountsWizardModel, abaw)
}

// UpdateAccountBankAccountsWizard updates an existing account.bank.accounts.wizard record.
func (c *Client) UpdateAccountBankAccountsWizard(abaw *AccountBankAccountsWizard) error {
	return c.UpdateAccountBankAccountsWizards([]int64{abaw.Id.Get()}, abaw)
}

// UpdateAccountBankAccountsWizards updates existing account.bank.accounts.wizard records.
// All records (represented by ids) will be updated by abaw values.
func (c *Client) UpdateAccountBankAccountsWizards(ids []int64, abaw *AccountBankAccountsWizard) error {
	return c.Update(AccountBankAccountsWizardModel, ids, abaw)
}

// DeleteAccountBankAccountsWizard deletes an existing account.bank.accounts.wizard record.
func (c *Client) DeleteAccountBankAccountsWizard(id int64) error {
	return c.DeleteAccountBankAccountsWizards([]int64{id})
}

// DeleteAccountBankAccountsWizards deletes existing account.bank.accounts.wizard records.
func (c *Client) DeleteAccountBankAccountsWizards(ids []int64) error {
	return c.Delete(AccountBankAccountsWizardModel, ids)
}

// GetAccountBankAccountsWizard gets account.bank.accounts.wizard existing record.
func (c *Client) GetAccountBankAccountsWizard(id int64) (*AccountBankAccountsWizard, error) {
	abaws, err := c.GetAccountBankAccountsWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if abaws != nil && len(*abaws) > 0 {
		return &((*abaws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.accounts.wizard not found", id)
}

// GetAccountBankAccountsWizards gets account.bank.accounts.wizard existing records.
func (c *Client) GetAccountBankAccountsWizards(ids []int64) (*AccountBankAccountsWizards, error) {
	abaws := &AccountBankAccountsWizards{}
	if err := c.Read(AccountBankAccountsWizardModel, ids, nil, abaws); err != nil {
		return nil, err
	}
	return abaws, nil
}

// FindAccountBankAccountsWizard finds account.bank.accounts.wizard record by querying it with criteria.
func (c *Client) FindAccountBankAccountsWizard(criteria *Criteria) (*AccountBankAccountsWizard, error) {
	abaws := &AccountBankAccountsWizards{}
	if err := c.SearchRead(AccountBankAccountsWizardModel, criteria, NewOptions().Limit(1), abaws); err != nil {
		return nil, err
	}
	if abaws != nil && len(*abaws) > 0 {
		return &((*abaws)[0]), nil
	}
	return nil, fmt.Errorf("no account.bank.accounts.wizard was found with criteria %v", criteria)
}

// FindAccountBankAccountsWizards finds account.bank.accounts.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankAccountsWizards(criteria *Criteria, options *Options) (*AccountBankAccountsWizards, error) {
	abaws := &AccountBankAccountsWizards{}
	if err := c.SearchRead(AccountBankAccountsWizardModel, criteria, options, abaws); err != nil {
		return nil, err
	}
	return abaws, nil
}

// FindAccountBankAccountsWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankAccountsWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankAccountsWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankAccountsWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountBankAccountsWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankAccountsWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no account.bank.accounts.wizard was found with criteria %v and options %v", criteria, options)
}
