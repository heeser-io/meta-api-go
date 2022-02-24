package v1

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	UpdateKeyDescription = "updatedKey"
	projectObj           *Project
	testNames            = []string{
		"create",
		"read",
		"update",
		"delete",
		"list",
	}
	tests = []func(t *testing.T){
		func(t *testing.T) {
			keyObj, err := createKey(projectObj.ID)
			if err != nil {
				t.Fatal(err)
			}
			if err := client.ApiKey.Delete(&DeleteApiKeyParams{
				ProjectID: projectObj.ID,
				ApiKeyID:  keyObj.ID,
			}); err != nil {
				t.Fatal(err)
			}
		},
		func(t *testing.T) {
			keyObj, err := createKey(projectObj.ID)
			if err != nil {
				t.Fatal(err)
			}

			readKeyObj, err := client.ApiKey.Read(&ReadApiKeyParams{
				ProjectID: projectObj.ID,
				ApiKeyID:  keyObj.ID,
			})
			if err != nil {
				t.Fatal(err)
			}

			if readKeyObj.ID != keyObj.ID {
				t.Fatal("no key found")
			}
			if err := client.ApiKey.Delete(&DeleteApiKeyParams{
				ProjectID: projectObj.ID,
				ApiKeyID:  keyObj.ID,
			}); err != nil {
				t.Fatal(err)
			}
		},
		func(t *testing.T) {
		},
		func(t *testing.T) {
			keyObj, err := createKey(projectObj.ID)
			if err != nil {
				t.Fatal(err)
			}
			if err := client.ApiKey.Delete(&DeleteApiKeyParams{
				ProjectID: projectObj.ID,
				ApiKeyID:  keyObj.ID,
			}); err != nil {
				t.Fatal(err)
			}

			_, err = client.ApiKey.Read(&ReadApiKeyParams{
				ProjectID: projectObj.ID,
				ApiKeyID:  keyObj.ID,
			})
			if err == nil {
				t.Fatal("expected error")
			}
		},
		func(t *testing.T) {
		},
	}
)

func createKey(projectID string) (*ApiKey, error) {
	keyObj, err := client.ApiKey.Create(generateCreateParams(projectID))
	if err != nil {
		return nil, err
	}
	return keyObj, nil
}

func generateCreateParams(projectID string) *CreateApiKeyParams {
	return &CreateApiKeyParams{
		ProjectID:   projectID,
		Description: "TestKey",
		Scope: []string{
			"auth.createToken",
		},
	}
}
func init() {
	godotenv.Load("../.env")
	client = WithAPIKey(os.Getenv("API_KEY"))
}

func prepareTest(t *testing.T) {
	var err error
	projectObj, err = createProject()
	if err != nil {
		t.Fatal(t)
	}
}

func finishTest(t *testing.T) {
	if err := client.Project.Delete(&DeleteProjectParams{
		ProjectID: projectObj.ID,
	}); err != nil {
		t.Fatal(err)
	}
}

func TestApiKey(t *testing.T) {
	prepareTest(t)
	for i, fn := range tests {
		t.Run(testNames[i], fn)
	}
	finishTest(t)
}
