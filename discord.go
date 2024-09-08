package discord

// BSD 3-Clause License | Copyright (c) 2024 Latte Softworks <https://latte.to>

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

var httpClient = &http.Client{}

func encodeMessage(message *Message) (payload *bytes.Buffer, contentType string, err error) {
	payload = new(bytes.Buffer)

	// Check for multipart/form-data upload stuff
	if message.Files != nil {
		writer := multipart.NewWriter(payload)

		partWriter, err := writer.CreateFormField("payload_json")
		if err != nil {
			return nil, "", err
		}

		err = json.NewEncoder(partWriter).Encode(message)
		if err != nil {
			return nil, "", err
		}

		for index, file := range *message.Files {
			partWriter, err := writer.CreateFormFile(fmt.Sprintf("files[%v]", index), file.Name)
			if err != nil {
				return nil, "", err
			}

			partWriter.Write(*file.Data)
		}

		writer.Close()

		return payload, "multipart/form-data; boundary=" + writer.Boundary(), nil
	} else {
		err = json.NewEncoder(payload).Encode(message)
		return payload, "application/json", err
	}
}

func returnMessageIdFromResp(resp *http.Response) (messageId *string, err error) {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// If wait=false it should be 204 (No Content)
	if resp.StatusCode == 204 {
		return nil, nil
	}

	// The status code MUST be 200 from here
	if resp.StatusCode != 200 {
		return nil, errors.New("Bad Discord response (Expected status code 200, got " + fmt.Sprint(resp.StatusCode) + "): " + string(respBody))
	}

	var respStruct struct {
		Id string `json:"id"`
	}

	if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(&respStruct); err != nil {
		return nil, err
	}

	return &respStruct.Id, nil
}

func constructQueryParams(queryParams *QueryParams) (form string) {
	formVals := make(url.Values)
	formVals.Add("wait", "true") // Should always be true (for message id)

	/*
		if queryParams.Wait {
			formVals.Add("wait", "true")
		}
	*/
	if queryParams.ThreadId != "" {
		formVals.Add("thread_id", queryParams.ThreadId)
	}

	return formVals.Encode()
}

func PostMessage(webhookUrl string, message *Message) (messageId *string, err error) {
	if message.QueryParams == nil {
		message.QueryParams = new(QueryParams)
	}
	webhookUrl += "?" + constructQueryParams(message.QueryParams)

	payload, contentType, err := encodeMessage(message)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(webhookUrl, contentType, payload)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return returnMessageIdFromResp(resp)
}

// REQUIRES AN APPLICATION-OWNED WEBHOOK
func EditMessage(webhookUrl string, messageId string, message *Message) (err error) {
	webhookUrl += "/messages/" + url.PathEscape(messageId)
	if message.QueryParams == nil {
		message.QueryParams = new(QueryParams)
	}
	webhookUrl += "?" + constructQueryParams(message.QueryParams)

	payload, contentType, err := encodeMessage(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, webhookUrl, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", contentType)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return errors.New("Bad Discord response (Expected status code 204, got " + fmt.Sprint(resp.StatusCode) + ")")
	}

	return nil
}

func DeleteMessage(webhookUrl string, messageId string, queryParams *QueryParams) (err error) {
	webhookUrl += "/messages/" + url.PathEscape(messageId)
	if queryParams != nil {
		webhookUrl += "?" + constructQueryParams(queryParams)
	}

	req, err := http.NewRequest(http.MethodDelete, webhookUrl, nil)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()

	if resp.StatusCode != 204 {
		return errors.New("Bad Discord response (Expected status code 204, got " + fmt.Sprint(resp.StatusCode) + ")")
	}

	return nil
}

func GetWebhookInfo(webhookUrl string) (webhookInfo *WebhookInfo, err error) {
	resp, err := http.Get(webhookUrl)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Bad Discord response (Expected status code 200, got " + fmt.Sprint(resp.StatusCode) + "): " + string(respBody))
	}

	webhookInfo = new(WebhookInfo)
	if err := json.NewDecoder(bytes.NewReader(respBody)).Decode(webhookInfo); err != nil {
		return nil, err
	}

	return webhookInfo, nil
}
