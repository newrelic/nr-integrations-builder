package main

import (
	"encoding/json"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/data/event"
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	// Insert here the logic for your tests

	i, err := integration.New("test", "1.0.0")
	assert.NoError(t, err)

	e, err := i.Entity("testEntity", "testNamespace")
	assert.NoError(t, err)

	e.AddEvent(event.New("testSummary", "testCategory"))

	assert.Equal(t, e.Events[0].Summary, "testSummary")
	assert.Equal(t, e.Events[0].Category, "testCategory")
	assert.Equal(t, 1, len(e.Events))

}

func TestMetrics(t *testing.T) {
	// Insert here the logic for your tests

	i, err := integration.New("test", "1.0.0")
	assert.NoError(t, err)

	e, err := i.Entity("testEntity", "testNamespace")
	assert.NoError(t, err)

	m, err := e.NewMetricSet("testMetrics")
	assert.NoError(t, err)

	m.SetMetric("testMetric1", 1, metric.GAUGE)
	m.SetMetric("testMetric2", "foo", metric.ATTRIBUTE)

	metrics := m.Metrics
	assert.Equal(t, 1, metrics["testMetric1"])
	assert.Equal(t, "foo", metrics["testMetric2"])

}

func TestCreateRemoteEntity(t *testing.T) {
	i, err := integration.New("test", "1.0.0")
	assert.NoError(t, err)

	e, err := i.Entity("testEntity", "testNamespace")
	assert.Equal(t, "testEntity", e.Metadata.Name)
	assert.Equal(t, "testNamespace", e.Metadata.Namespace)

	j, err := json.Marshal(i)

	assert.Equal(t, `{"name":"test","protocol_version":"2","integration_version":"1.0.0","data":[{"entity":{"name":"testEntity","type":"testNamespace"},"metrics":[],"inventory":{},"events":[]}]}`, string(j))

}
