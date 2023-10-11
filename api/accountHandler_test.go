package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_db "github.com/abhisheksatish1999/simplebank/db/mock"
	db "github.com/abhisheksatish1999/simplebank/db/sqlc"
	"github.com/abhisheksatish1999/simplebank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := createRandomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_db.NewMockStore(ctrl)
	store.EXPECT().GetAccount(gomock.Any(), account.ID).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, req)
	
	//check response code and body
	require.Equal(t, http.StatusOK, recorder.Code)
	body, err := io.ReadAll(recorder.Body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(body, &gotAccount)
	require.NoError(t, err)
	require.Equal(t,account,gotAccount)

}

func createRandomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
