package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	mapper "github.com/victorvelo/notes/client/unmarshal"
	"github.com/victorvelo/notes/client/yandex"
)

type Note struct {
	S []string `json:"s"`
}

func CheckNote(checkText string) ([]Note, error) {
	httpClient := &http.Client{}

	checkText = strings.ReplaceAll(checkText, " ", "+")

	client := yandex.NewClient(httpClient)
	req, err := getRequest(context.Background(), checkText)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := mapper.UnmarshalData[[]Note](resp)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getRequest(ctx context.Context, text string) (*http.Request, error) {
	path := fmt.Sprintf(yandex.PathCheckText, text)

	str := fmt.Sprintf("%s/%s", yandex.Address, path)
	fmt.Println(str)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, str, nil)
	if err != nil {
		return req, fmt.Errorf("creating object: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
