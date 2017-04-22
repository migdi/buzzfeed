package feed_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFeed(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Feed Suite")
}
