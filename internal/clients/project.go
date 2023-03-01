package clients

import (
	"context"
	"fmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/client/project_service"
	resourcemodels "github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/models"
)

// GetProjectByID gets a project by its ID
func GetProjectByID(ctx context.Context, client *Client, projectID string) (*resourcemodels.HashicorpCloudResourcemanagerProject, error) {
	getParams := project_service.NewProjectServiceGetParams()
	getParams.Context = ctx
	getParams.ID = projectID
	getResponse, err := client.Project.ProjectServiceGet(getParams, nil)
	if err != nil {
		return nil, err
	}

	return getResponse.Payload.Project, nil
}

// GetParentOrganizationIDByProjectID gets the parent organization ID of a project
func GetParentOrganizationIDByProjectID(ctx context.Context, client *Client, projectID string) (string, error) {
	project, err := GetProjectByID(ctx, client, projectID)
	if err != nil {
		return "", err
	}

	return project.Parent.ID, nil
}

func CreateProject(ctx context.Context, client *Client, name, organizationID string) (*resourcemodels.HashicorpCloudResourcemanagerProject, error) {
	projectOrg := &resourcemodels.HashicorpCloudResourcemanagerResourceID{
		ID:   organizationID,
		Type: resourcemodels.NewHashicorpCloudResourcemanagerResourceIDResourceType("ORGANIZATION"),
	}
	projectParams := project_service.NewProjectServiceCreateParamsWithContext(ctx)
	projectParams.Body = &resourcemodels.HashicorpCloudResourcemanagerProjectCreateRequest{
		Name:   name,
		Parent: projectOrg,
	}

	createProjectResp, err := client.Project.ProjectServiceCreate(projectParams, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create project '%s' with organization ID %s: %v", name, organizationID, err)
	}

	return createProjectResp.Payload.Project, nil
}
