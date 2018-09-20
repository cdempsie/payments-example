package persist_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/cdempsie/payments-example/api"
	"github.com/cdempsie/payments-example/persist"
	"github.com/cdempsie/payments-example/test"
	"github.com/google/uuid"
)

func TestCreate(t *testing.T) {
	store := persist.NewInMemoryStore()
	dec := json.NewDecoder(strings.NewReader(test.CreatePayment))
	payment := &api.Payment{}
	err := dec.Decode(payment)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}
	err = store.Create(payment)
	if err != nil {
		t.Fatalf("Failed to create payment in store: %v", err)
	}

	if payment.ID == "" {
		t.Fatalf("Payment ID was empty")
	}
}

func TestLoad(t *testing.T) {
	store := persist.NewInMemoryStore()
	payment := create(t, store)

	result, err := store.Load(payment.ID)
	if err != nil {
		t.Fatalf("Failed to load payment from store: %v", err)
	}
	if result == nil {
		t.Fatal("Expected payment struct but was nil")
	}

	//spot check
	if payment.ID != result.ID && payment.BeneficiaryParty.BankID != result.BeneficiaryParty.BankID {
		t.Fatalf("Expected payment structs to match but they didn't. Expected: %v\nGot %v\n", payment, result)
	}
}

func TestLoadNotFoundID(t *testing.T) {
	store := persist.NewInMemoryStore()
	create(t, store)

	testID := uuid.New().String()
	_, err := store.Load(testID)
	if err == nil {
		t.Fatalf("Expected error for unknown ID: %v", testID)
	}
}

func TestUpdate(t *testing.T) {
	store := persist.NewInMemoryStore()
	payment := create(t, store)

	payment.BeneficiaryParty.Address = "new address"
	err := store.Update(payment)
	if err != nil {
		t.Fatalf("Failed to load payment from store: %v", err)
	}
}

func TestUpdateNotFoundID(t *testing.T) {
	store := persist.NewInMemoryStore()
	payment := create(t, store)

	// set ID to a new val
	payment.ID = uuid.New().String()
	err := store.Update(payment)
	if err == nil {
		t.Fatalf("Expected error for unknown ID: %v", payment.ID)
	}
}

func TestDelete(t *testing.T) {
	store := persist.NewInMemoryStore()
	payment := create(t, store)

	err := store.Delete(payment.ID)
	if err != nil {
		t.Fatalf("Failed to delete payment from store: %v", err)
	}
}

func TestDeleteNotFoundID(t *testing.T) {
	store := persist.NewInMemoryStore()

	// set ID to a new val
	testID := uuid.New().String()
	err := store.Delete(testID)
	if err == nil {
		t.Fatalf("Expected error for unknown ID: %v", testID)
	}
}

func TestList(t *testing.T) {
	store := persist.NewInMemoryStore()
	payment := create(t, store)

	results, err := store.List()
	if err != nil {
		t.Fatalf("Failed to delete payment from store: %v", err)
	}

	if len(results.Data) > 1 {
		t.Fatalf("Got %d results when only 1 was expected", len(results.Data))
	}

	if results.Data[0].ID != payment.ID {
		t.Fatalf("Payment IDs did not match. Expected %s got: %s", payment.ID, results.Data[0].ID)
	}
}

func create(t *testing.T, store persist.PaymentStore) *api.Payment {
	dec := json.NewDecoder(strings.NewReader(test.CreatePayment))
	payment := &api.Payment{}
	err := dec.Decode(payment)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}
	err = store.Create(payment)
	if err != nil {
		t.Fatalf("Failed to create payment in store: %v", err)
	}

	return payment
}
