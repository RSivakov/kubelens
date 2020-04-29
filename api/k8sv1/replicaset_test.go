package k8sv1

import (
	"testing"

	rbacfakes "github.com/kubelens/kubelens/api/auth/fakes"
	logfakes "github.com/kubelens/kubelens/api/log/fakes"
	"github.com/stretchr/testify/assert"
)

func TestReplicaSetOverviewsDefaultSuccess(t *testing.T) {
	c := setupClient("testns", "test", false, false)

	ls := map[string]string{}
	ls[AppNameLabel] = FriendlyAppName

	d, err := c.ReplicaSetOverviews(ReplicaSetOptions{
		UserRole:      &rbacfakes.RoleAssignment{},
		Logger:        &logfakes.Logger{},
		Namespace:     "testns",
		LabelSelector: ls,
	})

	assert.Nil(t, err)
	assert.Len(t, d, 1)
	assert.Equal(t, d[0].Namespace, "testns")
}

func TestGetReplicaSetOverviewsDefaultFail(t *testing.T) {
	c := setupClient("testns", "test", true, true)

	_, err := c.ReplicaSetOverviews(ReplicaSetOptions{
		UserRole:      &rbacfakes.RoleAssignment{},
		Logger:        &logfakes.Logger{},
		Namespace:     "testns",
		LabelSelector: map[string]string{"random": "labelvalue"},
	})

	assert.NotNil(t, err)
}

func TestReplicaSetAppInfosDefaultSuccess(t *testing.T) {
	c := setupClient("testns", "test", false, false)

	ls := map[string]string{}
	ls[AppNameLabel] = FriendlyAppName

	d, err := c.ReplicaSetAppInfos(ReplicaSetOptions{
		UserRole:  &rbacfakes.RoleAssignment{},
		Logger:    &logfakes.Logger{},
		Namespace: "testns",
	})

	assert.Nil(t, err)
	assert.Len(t, d, 1)
	assert.Equal(t, d[0].Namespace, "testns")
}

func TestGetReplicaSetAppInfosDefaultFail(t *testing.T) {
	c := setupClient("testns", "test", true, true)

	_, err := c.ReplicaSetAppInfos(ReplicaSetOptions{
		UserRole:  &rbacfakes.RoleAssignment{},
		Logger:    &logfakes.Logger{},
		Namespace: "testns",
	})

	assert.NotNil(t, err)
}
