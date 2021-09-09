module github.com/crossplane-contrib/provider-tf-azure

go 1.16

require (
	github.com/crossplane-contrib/terrajet v0.0.0-20210825052229-d4afdfba2baf
	github.com/crossplane/crossplane-runtime v0.14.1-0.20210805220729-047d9387efbb
	github.com/crossplane/crossplane-tools v0.0.0-20210320162312-1baca298c527
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/hashicorp/terraform-provider-azurerm v1.44.1-0.20210827013307-f1e39e12e4ca
	github.com/iancoleman/strcase v0.2.0
	github.com/pkg/errors v0.9.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
	sigs.k8s.io/controller-runtime v0.9.2
	sigs.k8s.io/controller-tools v0.4.0
)

replace github.com/hashicorp/terraform-provider-azurerm => /Users/alper/data/workspaces/github.com/hashicorp/terraform-provider-azurerm

replace github.com/crossplane-contrib/terrajet => /Users/alper/data/workspaces/github.com/ulucinar/crossplane-contrib/terrajet
