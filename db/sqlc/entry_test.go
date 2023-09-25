package db

import (
	"context"
	"testing"
	"time"

	"github.com/abhisheksatish1999/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry

}
func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	expectedEntry := createRandomEntry(t, account)

	actualEntry, err := testQueries.GetEntry(context.Background(), expectedEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, actualEntry)
	require.Equal(t, expectedEntry.AccountID, actualEntry.AccountID)
	require.Equal(t, expectedEntry.Amount, actualEntry.Amount)
	require.Equal(t, expectedEntry.ID, actualEntry.ID)
	require.WithinDuration(t, expectedEntry.CreatedAt, actualEntry.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    0,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, entry.AccountID, account.ID)
	}
}
