package secrets

import "github.com/khulnasoft-lab/cloud-metadata/pkg/metadata"

var SensitiveInLocal = metadata.Metadata{
	ID:          "AVD-GEN-0003",
	Title:       "Potentially sensitive data stored in local value.",
	Description: "Sensitive attributes such as passwords and API tokens should not be available in your templates, especially in a plaintext form. You can declare variables to hold the secrets, assuming you can provide values for those variables in a secure fashion. Alternatively, you can store these secrets in a secure secret store, such as AWS KMS.

*NOTE: It is also recommended to store your Terraform state in an encrypted form.*",
	Impact:      "Local value could be leaking secrets",
	Severity:    "CRITICAL",
	Links:       []string {
		
	},
}

