package cluster_health_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClusterHealth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ClusterHealth Suite")
}
