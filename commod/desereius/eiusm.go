import (
	"context"
	"fmt"
	"io"

	healthcare "google.golang.org/api/healthcare/v1"
)

// approveDPX approves a Data Protection Policy.
func approveDPX(w io.Writer, projectID, location, datasetID, dpxID string) error {
	ctx := context.Background()

	healthcareService, err := healthcare.NewService(ctx)
	if err != nil {
		return fmt.Errorf("healthcare.NewService: %v", err)
	}

	datasetsService := healthcareService.Projects.Locations.Datasets

	name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/dataProtectionPolicies/%s", projectID, location, datasetID, dpxID)

	req := &healthcare.ApproveDataProtectionPolicyRequest{}

	if _, err := datasetsService.DataProtectionPolicies.Approve(name, req).Do(); err != nil {
		return fmt.Errorf("Approve: %v", err)
	}

	fmt.Fprintf(w, "Approved Data Protection Policy: %q\n", name)
	return nil
}
  
