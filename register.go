// Package kafka registers the xk6-output-kafka extension
package kafka

import (
	"github.com/pavel-kryvasheyeu/xk6-filtered-output-kafka/pkg/kafka"
	"go.k6.io/k6/output"
)

func init() {
	output.RegisterExtension("xk6-kafka", kafka.New)
}
