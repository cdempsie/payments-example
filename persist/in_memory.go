package persist

import (
	"fmt"
	"sync"

	"github.com/cdempsie/payments-example/api"
	"github.com/google/uuid"
)

// InMemoryStore provides an entirely in memory payment store.
// This is an example implementation of the PaymentStore interface and won't survive server restarts!
type InMemoryStore struct {
	data map[string]*api.Payment
	lock sync.RWMutex
}

// NewInMemoryStore return a newly initialised memory store.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{data: make(map[string]*api.Payment)}
}

// Create creates a new payment in the store, assigning a UUID in the process.
func (store *InMemoryStore) Create(payment *api.Payment) error {
	id := uuid.New().String()
	store.lock.Lock()
	defer store.lock.Unlock()

	payment.ID = id
	store.data[id] = payment

	return nil
}

// Update updates the given payment in the store.
// An error is returned if the payment with the given ID could not be found.
func (store *InMemoryStore) Update(payment *api.Payment) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	id := payment.ID
	if _, ok := store.data[id]; !ok {
		return fmt.Errorf("payment with ID: %s not found", id)
	}

	store.data[id] = payment

	return nil
}

func (store *InMemoryStore) Delete(paymentUID string) error {
	store.lock.Lock()
	defer store.lock.Unlock()

	if _, ok := store.data[paymentUID]; !ok {
		return fmt.Errorf("payment with ID: %s not found", paymentUID)
	}

	delete(store.data, paymentUID)

	return nil
}

// Load loads the payment with the given ID.
// If the payment is not found an error is returned.
func (store *InMemoryStore) Load(paymentUID string) (payment *api.Payment, err error) {
	store.lock.RLock()
	defer store.lock.RUnlock()

	if payment, ok := store.data[paymentUID]; ok {
		return payment, nil
	}

	return nil, fmt.Errorf("payment with ID: %s not found", paymentUID)
}

// List lists all the payments currently in the store.
func (store *InMemoryStore) List() (results *api.ListHolder, err error) {
	store.lock.RLock()
	defer store.lock.RUnlock()

	result := &api.ListHolder{}

	for _, payment := range store.data {
		result.Data = append(result.Data, *payment)
	}

	return result, nil
}
