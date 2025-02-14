// Auto generated code - DO NOT EDIT.

package monitor

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
)

func createDiagnosticSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorDiagnosticSettingsClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			DiagnosticSettings: mockClient,
		},
	}

	data := insights.DiagnosticSettingsResource{}
	require.Nil(t, faker.FakeObject(&data))

	result := insights.DiagnosticSettingsResourceCollection{Value: &[]insights.DiagnosticSettingsResource{data}}

	mockClient.EXPECT().List(gomock.Any(), "/subscriptions/testSubscription").Return(result, nil)
	mockClient.EXPECT().List(gomock.Any(), "/subscriptions/test/resourceGroups/test/providers/test/test/test").Return(result, nil)
	return s
}
