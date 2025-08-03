package stats

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestPrometheus(t *testing.T) {
	p := NewPrometheus("test_cache",
		WithNamespace("test_namespace"),
		WithSubsystem("test_subsystem"),
		WithName("test_name"),
		WithHelp("test_help"))
	assert.Equal(t, "test_cache", p.cacheName)
	assert.Equal(t, "test_namespace", p.namespace)
	assert.Equal(t, "test_subsystem", p.subsystem)
	assert.Equal(t, "test_name", p.name)
	assert.Equal(t, "test_help", p.help)

	assert.Equal(t, 0, testutil.CollectAndCount(prometheusCounterVec))

	p.IncrHit()
	assert.Equal(t, 1, testutil.CollectAndCount(prometheusCounterVec))
	assert.Equal(t, float64(1), testutil.ToFloat64(prometheusCounterVec.WithLabelValues("test_cache", "total", "hit", "")))

	p.IncrMiss()
	assert.Equal(t, 2, testutil.CollectAndCount(prometheusCounterVec))
	assert.Equal(t, float64(1), testutil.ToFloat64(prometheusCounterVec.WithLabelValues("test_cache", "total", "miss", "")))

	p.IncrLocalHit()
	assert.Equal(t, 3, testutil.CollectAndCount(prometheusCounterVec))
	assert.Equal(t, float64(1), testutil.ToFloat64(prometheusCounterVec.WithLabelValues("test_cache", "local", "hit", "")))

	p.IncrLocalMiss()
	assert.Equal(t, 4, testutil.CollectAndCount(prometheusCounterVec))
	assert.Equal(t, float64(1), testutil.ToFloat64(prometheusCounterVec.WithLabelValues("test_cache", "local", "miss", "")))

	p.IncrRemoteHit()
	assert.Equal(t, 5, testutil.CollectAndCount(prometheusCounterVec))
	assert.Equal(t, float64(1), testutil.ToFloat64(prometheusCounterVec.WithLabelValues("test_cache", "remote", "hit", "")))

	p.IncrRemoteMiss()
	assert.Equal(t, 6, testutil.CollectAndCount(prometheusCounterVec))
	assert.Equal(t, float64(1), testutil.ToFloat64(prometheusCounterVec.WithLabelValues("test_cache", "remote", "miss", "")))

	p.IncrQuery()
	assert.Equal(t, 7, testutil.CollectAndCount(prometheusCounterVec))
	assert.Equal(t, float64(1), testutil.ToFloat64(prometheusCounterVec.WithLabelValues("test_cache", "query", "query", "")))

	p.IncrQueryFail(assert.AnError)
	assert.Equal(t, 8, testutil.CollectAndCount(prometheusCounterVec))
	assert.Equal(t, float64(1),
		testutil.ToFloat64(prometheusCounterVec.WithLabelValues("test_cache", "query", "queryFail", assert.AnError.Error())))
}

func TestNewPrometheus_MultipleInstances(t *testing.T) {
	assert.NotPanics(t, func() {
		NewPrometheus("test_cache_1")
		NewPrometheus("test_cache_2")
	})
}
