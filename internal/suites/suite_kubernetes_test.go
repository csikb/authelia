package suites

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type KubernetesSuite struct {
	*RodSuite
}

func NewKubernetesSuite() *KubernetesSuite {
	return &KubernetesSuite{RodSuite: new(RodSuite)}
}

func (s *KubernetesSuite) Test1FAScenario() {
	suite.Run(s.T(), New1FAScenario())
}

func (s *KubernetesSuite) Test2FAScenario() {
	suite.Run(s.T(), New2FAScenario())
}

func (s *KubernetesSuite) TestRedirectionURLScenario() {
	suite.Run(s.T(), NewRedirectionURLScenario())
}

func TestKubernetesSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping suite test in short mode")
	}

	suite.Run(t, NewKubernetesSuite())
}
