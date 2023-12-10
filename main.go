package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"
	"time"
)

type RequestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

type ResponsePayload struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNS       int64   `json:"time_ns"`
}

func sortSequential(arrays [][]int) ([][]int, int64) {
	startTime := time.Now()
	sortedArrays := make([][]int, len(arrays))
	for i, subArray := range arrays {
		sorted := make([]int, len(subArray))
		copy(sorted, subArray)
		sort.Ints(sorted)
		sortedArrays[i] = sorted
	}
	elapsedTime := time.Since(startTime).Nanoseconds()
	return sortedArrays, elapsedTime
}

func sortConcurrent(arrays [][]int) ([][]int, int64) {
	startTime := time.Now()
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	sortedArrays := make([][]int, len(arrays))

	for i, subArray := range arrays {
		wg.Add(1)
		go func(index int, arr []int) {
			defer wg.Done()
			sorted := make([]int, len(arr))
			copy(sorted, arr)
			sort.Ints(sorted)
			mutex.Lock()
			sortedArrays[index] = sorted
			mutex.Unlock()
		}(i, subArray)
	}

	wg.Wait()
	elapsedTime := time.Since(startTime).Nanoseconds()
	return sortedArrays, elapsedTime
}

func processSingleHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	sortedArrays, timeNS := sortSequential(requestPayload.ToSort)

	responsePayload := ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNS:       timeNS,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responsePayload)
}

func processConcurrentHandler(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	sortedArrays, timeNS := sortConcurrent(requestPayload.ToSort)

	responsePayload := ResponsePayload{
		SortedArrays: sortedArrays,
		TimeNS:       timeNS,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responsePayload)
}

func main() {
	http.HandleFunc("/process-single", processSingleHandler)
	http.HandleFunc("/process-concurrent", processConcurrentHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
