package v1

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	client       *Client
	createParams *CreateProjectParams
	UpdateName   = "updatedProject"
)

func init() {
	godotenv.Load("../.env")
	client = WithAPIKey(os.Getenv("API_KEY"))

	createParams = &CreateProjectParams{
		Name: "TestProjekt",
	}
}

func createProject() (*Project, error) {
	createdProject, err := client.Project.Create(createParams)
	if err != nil {
		return nil, err
	}
	return createdProject, nil
}

func TestCreate(t *testing.T) {
	createdProject, err := createProject()
	if err != nil {
		t.Fatal(err)
	}

	if createdProject.Name != createParams.Name {
		t.Fatalf("expected name %s, got %s", createParams.Name, createdProject.Name)
	}

	if err := client.Project.Delete(&DeleteProjectParams{
		ProjectID: createdProject.ID,
	}); err != nil {
		t.Fatal(err)
	}
}

func TestRead(t *testing.T) {
	createdProject, err := createProject()
	if err != nil {
		t.Fatal(err)
	}

	if createdProject.Name != createParams.Name {
		t.Fatalf("expected name %s, got %s", createParams.Name, createdProject.Name)
	}

	readProject, err := client.Project.Read(&ReadProjectParams{
		ProjectID: createdProject.ID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if readProject.Name != createParams.Name {
		t.Fatalf("expected name %s, got %s", createParams.Name, createdProject.Name)
	}

	if err := client.Project.Delete(&DeleteProjectParams{
		ProjectID: createdProject.ID,
	}); err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	createdProject, err := createProject()
	if err != nil {
		t.Fatal(err)
	}

	if createdProject.Name != createParams.Name {
		t.Fatalf("expected name %s, got %s", createParams.Name, createdProject.Name)
	}

	updatedProject, err := client.Project.Update(&UpdateProjectParams{
		ProjectID: createdProject.ID,
		Name:      UpdateName,
	})
	if err != nil {
		t.Fatal(err)
	}
	if updatedProject.Name != UpdateName {
		t.Fatalf("expected name %s, got %s", UpdateName, updatedProject.Name)
	}

	readProject, err := client.Project.Read(&ReadProjectParams{
		ProjectID: createdProject.ID,
	})
	if err != nil {
		t.Fatal(err)
	}

	if readProject.Name != UpdateName {
		t.Fatalf("expected name %s, got %s", UpdateName, readProject.Name)
	}

	if err := client.Project.Delete(&DeleteProjectParams{
		ProjectID: createdProject.ID,
	}); err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	createdProject, err := createProject()
	if err != nil {
		t.Fatal(err)
	}

	if createdProject.Name != createParams.Name {
		t.Fatalf("expeced name %s, got %s", createParams.Name, createdProject.Name)
	}

	if err := client.Project.Delete(&DeleteProjectParams{
		ProjectID: createdProject.ID,
	}); err != nil {
		t.Fatal(err)
	}

	readProject, err := client.Project.Read(&ReadProjectParams{
		ProjectID: createdProject.ID,
	})
	log.Println(readProject)
	if err == nil {
		t.Fatal("expected an error")
	}
}

func TestList(t *testing.T) {
}
