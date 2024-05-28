package kpidrive

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"kpi_drive_buffer/pkg/misc"
	"net/http"
	"net/url"
	"strconv"
)

var logger = misc.GetLogger()

type Adapter struct {
	*misc.HttpAdapter
	Token string
}

func NewAdapter(baseUrl string, token string) (*Adapter, error) {
	httpAdapter := misc.NewHttpAdapter(baseUrl)
	//Здесь должен быть ping внешнего API
	//httpAdapter.Ping()

	return &Adapter{
		HttpAdapter: httpAdapter,
		Token:       token,
	}, nil
}

func (k *Adapter) SaveFact(saveFactForm *SaveFactForm) (*SaveFactResponseWithStatus, error) {
	client := new(http.Client)

	var params = map[string]string{
		"period_start":            saveFactForm.PeriodStart,
		"period_end":              saveFactForm.PeriodEnd,
		"period_key":              saveFactForm.PeriodKey,
		"indicator_to_mo_id":      strconv.Itoa(saveFactForm.IndicatorToMoID),
		"indicator_to_mo_fact_id": strconv.Itoa(saveFactForm.IndicatorToMoFactID),
		"value":                   strconv.Itoa(saveFactForm.Value),
		"fact_time":               saveFactForm.FactTime,
		"is_plan":                 strconv.Itoa(saveFactForm.IsPlan),
		"auth_user_id":            strconv.Itoa(saveFactForm.AuthUserID),
		"comment":                 saveFactForm.Comment,
	}
	formData := url.Values{}
	for key, value := range params {
		formData.Set(key, value)
	}
	body := bytes.NewBufferString(formData.Encode())

	req, err := http.NewRequest("POST", k.DefaultUrl, body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при создании запроса: %s", err.Error()))
	}

	// Установка типа запроса
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Авторизация
	req.Header.Set("Authorization", "Bearer "+k.Token)

	// Отправка запроса
	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при выполнении запроса: %s", err.Error()))
	}
	defer response.Body.Close()

	var resp SaveFactResponse
	if err = json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return &SaveFactResponseWithStatus{
		Body:       &resp,
		StatusCode: response.StatusCode,
	}, nil
}
