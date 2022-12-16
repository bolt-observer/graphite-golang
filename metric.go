package graphite

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Metric is a struct that defines the relevant properties of a graphite metric
type Metric struct {
	Name      string
	Value     string
	Timestamp int64
	Tags      map[string]string
}

func NewMetric(name, value string, timestamp int64) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
		Tags:      make(map[string]string),
	}
}

func NewMetricWithTags(name, value string, timestamp int64, tags map[string]string) Metric {
	return Metric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
		Tags:      tags,
	}
}

func (metric Metric) IsUninitialized() bool {
	return metric.Name == "" && metric.Value == "" && metric.Timestamp == 0
}

func (metric Metric) convertTags() string {
	whitespace := regexp.MustCompile(`\s`)
	builder := strings.Builder{}

	if metric.Tags == nil || len(metric.Tags) == 0 {
		return ""
	}
	for k, v := range metric.Tags {
		if whitespace.MatchString(k) || whitespace.MatchString(v) {
			continue
		}

		builder.WriteString(";")
		builder.WriteString(k)
		builder.WriteString("=")
		builder.WriteString(v)
	}

	return builder.String()
}

func (metric Metric) ValueWithTags() string {
	return fmt.Sprintf("%s%s", metric.Value, metric.convertTags())
}

func (metric Metric) String() string {
	return fmt.Sprintf(
		"%s%s %s %s",
		metric.Name,
		metric.convertTags(),
		metric.Value,
		time.Unix(metric.Timestamp, 0).Format("2006-01-02 15:04:05"),
	)
}
