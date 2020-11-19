package queue

import (
	"fmt"
	"testing"
)

func TestCreatePriorityQueue(t *testing.T) {
	queue := CreatePriorityQueue()
	t.Logf("CreateQueue: %+v \n\n", queue)

	// Test priority_Insert code
	tests := []struct {
		id   int
		sort int
	}{
		{id: 5, sort: 10},
		{id: 20, sort: 18},
		{id: 21, sort: 1},
		{id: 22, sort: 4},
		{id: 23, sort: 7},
		{id: 24, sort: 58},
		{id: 25, sort: 98},
		{id: 60, sort: 28},
	}
	for key, value := range tests {
		t.Run(fmt.Sprintf("RunTest%d", key), func(t *testing.T) {
			err := queue.Insert(value.id, value.sort)
			if err != nil {
				t.Error(err)
				return
			}
			//t.Logf("TestInsert: id: %d sort: %d\n", value.id, value.sort)
		})
	}
	t.Logf("\nQueueContent: %+v\n", queue)
	// Test priority_Insert code end

	// Test Len
	wantLen := len(tests)
	result := queue.Len()
	if wantLen != result {
		t.Errorf("Wrong result from method Len: result: %d  wantLen: %d \n", result, wantLen)
	}

	// test GetMax
	id, sort := queue.GetMax()
	wantMax := struct {
		id   int
		sort int
	}{25, 98}
	if id != wantMax.id || sort != wantMax.sort {
		t.Errorf("Wrong result from method GetMax: resultId: %d  resultSort: %d want: %+v \n", id, sort, wantMax)
	}

	// test  GetSort
	theSort := queue.GetSort(25)
	wantSort := 98
	if wantLen != result {
		t.Errorf("Wrong result from method GetSort: result: %d  wantLen: %d \n", theSort, wantSort)
	}

	// test Contain
	if !queue.Contain(25) {
		t.Errorf("Wrong result from method Contain;")
	}

	// test change
	wantChange := struct {
		id   int
		sort int
	}{25, 19}
	err := queue.Change(wantChange.id, wantChange.sort)
	if err != nil {
		t.Error(err)
	}
	if queue.GetSort(wantChange.id) != wantChange.sort {
		t.Errorf("Wrong result from method Change;")
	}

	//test extract
	var resultList []int
	for !queue.IsEmpty() {
		_, itemSort, err := queue.Extract()
		if err != nil {
			t.Error(err)
			break
		}
		//t.Logf("%d\n", itemSort)
		resultList = append(resultList, itemSort)
	}
	// validate
	for i := 1; i < len(resultList); i++ {
		if resultList[i-1] < resultList[i] {
			t.Errorf("Wrong sort!")
			break
		}
	}

}
