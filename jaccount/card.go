/*
Copyright 2021 The Go jAccount Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package jaccount

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// CardService handles communications with the card data related methods of the jAccount API.
//
// See https://developer.sjtu.edu.cn/api/card.html for more information.
type CardService service

type CardInfo struct {
	User         *Profile `json:"user,omitempty"`
	CardNO       string   `json:"cardNo,omitempty"`
	CardID       string   `json:"cardId,omitempty"`
	BankNO       string   `json:"bankNo,omitempty"`
	CardBalance  float64  `json:"cardBalance,omitempty"`
	Transbalance float64  `json:"transbalance,omitempty"`
	Lost         bool     `json:"lost,omitempty"`
	Frozen       bool     `json:"frozen,omitempty"`
}

// GetCardInfo returns the card information for the user.
//
// See https://developer.sjtu.edu.cn/api/card.html#%E8%8E%B7%E5%8F%96%E6%A0%A1%E5%9B%AD%E5%8D%A1%E4%BF%A1%E6%81%AF for more information.
func (s *CardService) GetCardInfo(ctx context.Context) (*CardInfo, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/v1/me/card", nil)
	if err != nil {
		return nil, err
	}

	cardInfo := make([]CardInfo, 1)
	_, err = s.client.Do(ctx, req, &cardInfo)
	if err != nil {
		return nil, err
	}

	return &cardInfo[0], nil
}

type CardTransaction struct {
	DateTime    int64   `json:"dateTime,omitempty"`
	System      string  `json:"system,omitempty"`
	Merchant    string  `json:"merchant,omitempty"`
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	CardBalance float64 `json:"cardBalance,omitempty"`
}

type CardListTransactionsOptions struct {
	CardNo    string `url:"cardNo,omitempty"`
	BeginDate int64  `url:"beginDate,omitempty"`
	EndDate   int64  `url:"endDate,omitempty"`
}

// ListTransactions returns a list of transactions for the given card.
//
// See https://developer.sjtu.edu.cn/api/card.html#%E8%8E%B7%E5%8F%96%E4%BA%A4%E6%98%93%E8%AE%B0%E5%BD%95%E4%BF%A1%E6%81%AF for more information.
func (s *CardService) ListTransactions(ctx context.Context, opts *CardListTransactionsOptions) ([]*CardTransaction, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, "/v1/me/card/transactions", values)
	if err != nil {
		return nil, err
	}

	var transactions []*CardTransaction
	_, err = s.client.Do(ctx, req, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
