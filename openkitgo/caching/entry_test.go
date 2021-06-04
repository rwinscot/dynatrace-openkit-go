package caching

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEntryDataManipulation(t *testing.T) {
	e := BeaconCacheEntry{}

	e.addEventData(NewBeaconCacheRecord(time.Now(), "contents_1"))
	assert.Equal(t, int64(20), e.totalNumBytes)
	assert.Equal(t, 1, len(e.eventData))

	e.addActionData(NewBeaconCacheRecord(time.Now().Add(-10*time.Minute), "contents_2"))
	assert.Equal(t, int64(40), e.totalNumBytes)
	assert.Equal(t, 1, len(e.actionData))

	e.copyDataForSending()
	assert.Equal(t, 0, len(e.actionData))
	assert.Equal(t, 0, len(e.eventData))
	assert.Equal(t, 1, len(e.actionDataBeingSent))
	assert.Equal(t, 1, len(e.eventDataBeingSent))

	e.removeRecordsOlderThan(time.Now().Add(-9 * time.Minute))
	assert.Equal(t, 0, len(e.actionDataBeingSent))
	assert.Equal(t, 1, len(e.eventDataBeingSent))

	e.addEventData(NewBeaconCacheRecord(time.Now(), "contents_3"))
	e.addActionData(NewBeaconCacheRecord(time.Now().Add(-10*time.Minute), "contents_4"))

	e.removeOldestRecords(1)
	assert.Equal(t, 0, len(e.actionData))
	assert.Equal(t, 1, len(e.eventData))

	c := e.getChunk("prefix", 1024*30, '&')
	assert.Equal(t, "prefix&contents_1", c)

	e.removeDataMarkedForSending()
	assert.Equal(t, 0, len(e.actionDataBeingSent))
	assert.Equal(t, 0, len(e.eventDataBeingSent))

}
