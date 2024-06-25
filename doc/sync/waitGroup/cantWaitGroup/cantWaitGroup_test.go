package cantWaitGroup

import (
	"errors"
	"testing"
	"time"
)

func TestCantWaitGroup_AllTasksComplete(t *testing.T) {
	cwg := NewCantWaitGroup(2 * time.Second)

	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		cwg.Add(1)
		go func(id int) {
			defer cwg.Done()
			time.Sleep(1 * time.Second)
		}(i)
	}

	if err := cwg.Wait(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCantWaitGroup_Timeout(t *testing.T) {
	cwg := NewCantWaitGroup(1 * time.Second)

	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		cwg.Add(1)
		go func(id int) {
			defer cwg.Done()
			time.Sleep(2 * time.Second)
		}(i)
	}

	if err := cwg.Wait(); !errors.Is(err, CantWaitTimeoutErr) {
		t.Errorf("Expected timeout error, got %v", err)
	}
}

func TestCantWaitGroup_NoTasks(t *testing.T) {
	cwg := NewCantWaitGroup(1 * time.Second)

	if err := cwg.Wait(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestCantWaitGroup_TimeoutAndComplete(t *testing.T) {
	cwg := NewCantWaitGroup(1 * time.Second)

	cwg.Add(1)
	go func() {
		defer cwg.Done()
		time.Sleep(2 * time.Second)
	}()

	err := cwg.Wait()
	if !errors.Is(err, CantWaitTimeoutErr) {
		t.Errorf("Expected timeout error, got %v", err)
	}

	// Ensure it can still complete properly after timeout
	cwg.Add(1)
	go func() {
		defer cwg.Done()
		time.Sleep(500 * time.Millisecond)
	}()

	if err := cwg.Wait(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
