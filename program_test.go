package main

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_SelectTimeFrame(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(SelectTimeFrame(), 0, "Select timeFrame failed")
}

func Test_GetApiStockData(t *testing.T) {
	// 1D
	res, err := http.Get("https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1D.json")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	// 1W
	res, err = http.Get("https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1W.json")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	// 1M
	res, err = http.Get("https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1M.json")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	// 1Y
	res, err = http.Get("https://storage.googleapis.com/finno-ex-re-v2-static-staging/recruitment-test/fund-ranking-1Y.json")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

func Test_GetApiStockData2(t *testing.T) {
	GetApiStockData(1)
	GetApiStockData(5)
}

func Test_GetTimeFrameString(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(GetTimeFrameString(1), "1D", "Get timeFrame failed in '1D'")
	assert.Equal(GetTimeFrameString(2), "1W", "Get timeFrame failed in '1W'")
	assert.Equal(GetTimeFrameString(3), "1M", "Get timeFrame failed in '1M'")
	assert.Equal(GetTimeFrameString(4), "1Y", "Get timeFrame failed in '1Y'")
}


