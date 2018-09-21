// Package api holds the structs used by the api.
package api

import "strings"

// ListHolder contains the struct used to respond to a list collection response.
type ListHolder struct {
	Data []Payment `json:"data"`
}

// Payment API type.
type Payment struct {
	Type           string `json:"type"`
	ID             string `json:"id"`
	Version        int    `json:"version"`
	OrganisationID string `json:"organisation_id"`
	Attributes     `json:"attributes"`
}

// Valid returns true if the payment passes validation. Otherwise it returns false and a message containing the
// detected errors.
func (payment *Payment) Valid() (valid bool, messages string) {
	valid = true
	msgBuf := &strings.Builder{}

	// example spot check validation only.
	if payment.Type == "" {
		valid = false
		msgBuf.WriteString("payment type is missing")
	}
	if payment.OrganisationID == "" {
		valid = false
		if msgBuf.Len() >0 {
			msgBuf.WriteString(", ")
		}
		msgBuf.WriteString( "payment organisation ID is missing")
	}

	if msgBuf.Len() > 0 {
		messages = msgBuf.String()
		strings.TrimSpace(messages)
	}

	return
}

// Attributes API type.
type Attributes struct {
	Amount               string `json:"amount"`
	BeneficiaryParty     `json:"beneficiary_party"`
	ChargesInformation   `json:"charges_information"`
	Currency             string `json:"currency"`
	DebtorParty          `json:"debtor_party"`
	EndToEndReference    string `json:"end_to_end_reference"`
	Fx                   `json:"fx"`
	NumericReference     string `json:"numeric_reference"`
	PaymentID            string `json:"payment_id"`
	PaymentPurpose       string `json:"payment_purpose"`
	PaymentScheme        string `json:"payment_scheme"`
	PaymentType          string `json:"payment_type"`
	ProcessingDate       string `json:"processing_date"`
	Reference            string `json:"reference"`
	SchemePaymentSubType string `json:"scheme_payment_sub_type"`
	SchemePaymentType    string `json:"scheme_payment_type"`
	SponsorParty         `json:"sponsor_party"`
}

// BeneficiaryParty API type.
type BeneficiaryParty struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

// ChargesInformation API type.
type ChargesInformation struct {
	BearerCode    string `json:"bearer_code"`
	SenderCharges []struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
	} `json:"sender_charges"`
	ReceiverChargesAmount   string `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string `json:"receiver_charges_currency"`
}

// DebtorParty API type.
type DebtorParty struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

// Fx API type.
type Fx struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

// SponsorParty API type.
type SponsorParty struct {
	AccountNumber string `json:"account_number"`
	BankID        string `json:"bank_id"`
	BankIDCode    string `json:"bank_id_code"`
}
