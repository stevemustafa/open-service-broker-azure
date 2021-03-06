package storage

import (
	"fmt"

	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (s *serviceManager) Bind(
	service.Instance,
	service.BindingParameters,
) (service.BindingDetails, error) {
	return nil, nil
}

// nolint: lll
func (s *serviceManager) GetCredentials(
	instance service.Instance,
	_ service.Binding,
) (service.Credentials, error) {
	dt := instance.Details.(*instanceDetails)
	credential := credentials{
		StorageAccountName:          dt.StorageAccountName,
		AccessKey:                   dt.AccessKey,
		ContainerName:               dt.ContainerName,
		PrimaryBlobServiceEndPoint:  fmt.Sprintf("https://%s.blob.core.windows.net/", dt.StorageAccountName),
		PrimaryTableServiceEndPoint: fmt.Sprintf("https://%s.table.core.windows.net/", dt.StorageAccountName),
	}
	storeKind, _ := instance.Plan.GetProperties().Extended[kindKey].(storageKind)
	if storeKind == storageKindGeneralPurposeStorageAcccount ||
		storeKind == storageKindGeneralPurposeV2StorageAccount {
		credential.PrimaryFileServiceEndPoint = fmt.Sprintf("https://%s.file.core.windows.net/", dt.StorageAccountName)
		credential.PrimaryQueueServiceEndPoint = fmt.Sprintf("https://%s.queue.core.windows.net/", dt.StorageAccountName)
	}
	return credential, nil
}
