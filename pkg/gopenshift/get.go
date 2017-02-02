package gopenshift

import (
	"fmt"

	"k8s.io/kubernetes/pkg/kubectl/resource"
)

// TODO: Since we'll be watching this stuff we should implement the watcher logic
//       from the client
// TODO: We are missing a lot of stuff for getting "all"
func (o *OpenShift) Get(reqResource string) ([]map[string]string, error) {
	var result []map[string]string

	// TODO: Find a way to decide when to print all namespaces
	allNamespaces := false

	cmdNamespace, enforceNamespace, err := o.Factory.DefaultNamespace()
	if err != nil {
		return []map[string]string{}, err
	}

	// TODO: Comment on what this does
	stuff := []string{}
	r := resource.NewBuilder(o.Mapper, o.Typer, resource.ClientMapperFunc(o.Factory.ClientForMapping), o.Factory.Decoder(true)).
		NamespaceParam(cmdNamespace).DefaultNamespace().AllNamespaces(allNamespaces).
		FilenameParam(enforceNamespace, false, stuff...).
		SelectorParam("").
		ExportParam(false).
		ResourceTypeOrNameArgs(true, reqResource).
		ContinueOnError().
		Latest().
		Flatten().
		Do()

	err = r.Err()
	if err != nil {
		return []map[string]string{}, err
	}

	infos, err := r.Infos()
	if err != nil {
		return []map[string]string{}, err
	}

	for ix := range infos {
		currentObject := make(map[string]string)

		currentObject, _ = infos[ix].ResourceMapping().Labels(infos[ix].Object)

		fmt.Println("LABELS")
		for k, v := range currentObject {
			fmt.Printf("%s: %s\n", k, v)
		}

		infos[ix].ResourceMapping()

		result = append(result, currentObject)
	}

	return result, nil
}

func (o *OpenShift) GetPods() ([]map[string]string, error) {
	return o.Get("pods")
}

func (o *OpenShift) GetServices() ([]map[string]string, error) {
	return o.Get("svc")
}

var podColumns = []string{"NAME", "READY", "STATUS", "RESTARTS", "AGE"}
var podTemplateColumns = []string{"TEMPLATE", "CONTAINER(S)", "IMAGE(S)", "PODLABELS"}
var replicationControllerColumns = []string{"NAME", "DESIRED", "CURRENT", "AGE"}
var replicaSetColumns = []string{"NAME", "DESIRED", "CURRENT", "AGE"}
var jobColumns = []string{"NAME", "DESIRED", "SUCCESSFUL", "AGE"}
var serviceColumns = []string{"NAME", "CLUSTER-IP", "EXTERNAL-IP", "PORT(S)", "AGE"}
var ingressColumns = []string{"NAME", "HOSTS", "ADDRESS", "PORTS", "AGE"}
var petSetColumns = []string{"NAME", "DESIRED", "CURRENT", "AGE"}
var endpointColumns = []string{"NAME", "ENDPOINTS", "AGE"}
var nodeColumns = []string{"NAME", "STATUS", "AGE"}
var daemonSetColumns = []string{"NAME", "DESIRED", "CURRENT", "NODE-SELECTOR", "AGE"}
var eventColumns = []string{"LASTSEEN", "FIRSTSEEN", "COUNT", "NAME", "KIND", "SUBOBJECT", "TYPE", "REASON", "SOURCE", "MESSAGE"}
var limitRangeColumns = []string{"NAME", "AGE"}
var resourceQuotaColumns = []string{"NAME", "AGE"}
var namespaceColumns = []string{"NAME", "STATUS", "AGE"}
var secretColumns = []string{"NAME", "TYPE", "DATA", "AGE"}
var serviceAccountColumns = []string{"NAME", "SECRETS", "AGE"}
var persistentVolumeColumns = []string{"NAME", "CAPACITY", "ACCESSMODES", "STATUS", "CLAIM", "REASON", "AGE"}
var persistentVolumeClaimColumns = []string{"NAME", "STATUS", "VOLUME", "CAPACITY", "ACCESSMODES", "AGE"}
var componentStatusColumns = []string{"NAME", "STATUS", "MESSAGE", "ERROR"}
var thirdPartyResourceColumns = []string{"NAME", "DESCRIPTION", "VERSION(S)"}
var roleColumns = []string{"NAME", "AGE"}
var roleBindingColumns = []string{"NAME", "AGE"}
var clusterRoleColumns = []string{"NAME", "AGE"}
var clusterRoleBindingColumns = []string{"NAME", "AGE"}
