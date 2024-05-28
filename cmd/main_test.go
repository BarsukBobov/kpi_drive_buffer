package main

import (
	"encoding/json"
	"kpi_drive_buffer/internal/repository/kpidrive"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"testing"
)

// sendTestPackages отправляет 10 тестовых пакетов данных
func sendTestPackages(t *testing.T) {
	var wg sync.WaitGroup
	testPackages := 10
	wg.Add(testPackages)
	packagesIds := make([]int, testPackages)
	for i := 0; i < testPackages; i++ {
		go func(i int) {
			defer wg.Done()
			data := url.Values{}
			data.Set("period_start", "2024-05-01")
			data.Set("period_end", "2024-05-31")
			data.Set("period_key", "month")
			data.Set("indicator_to_mo_id", strconv.Itoa(227373))
			data.Set("indicator_to_mo_fact_id", strconv.Itoa(0))
			data.Set("value", strconv.Itoa(i))
			data.Set("fact_time", "2024-05-31")
			data.Set("is_plan", strconv.Itoa(0))
			data.Set("auth_user_id", strconv.Itoa(40))
			data.Set("comment", "")

			response, err := http.PostForm("http://localhost:3000/api/v1/buffer/save_fact", data)
			if err != nil {
				t.Errorf("Failed to send request: %v", err)
				return
			}
			defer response.Body.Close()

			var resp kpidrive.SaveFactResponse
			if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
				t.Errorf("Failed to decode response: %v", err)
			}
			packagesIds = append(packagesIds, resp.DATA.IndicatorToMoFactId)
		}(i)
	}

	wg.Wait()

	if !isSortedAscending(packagesIds) {
		t.Errorf("Incorect order: %v", packagesIds)
	}
}

// isSortedAscending проверяет что пакеты были отправлены в правильном порядке
func isSortedAscending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func TestSendTestPackages(t *testing.T) {
	sendTestPackages(t)
}
