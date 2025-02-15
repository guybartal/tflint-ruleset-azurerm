mapping "azurerm_kubernetes_cluster" {
  import_path = "azure-rest-api-specs/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-03-02-preview/managedClusters.json"

  name                       = ResourceNameParameter
  resource_group_name        = any //ResourceGroupNameParameter
  dns_prefix                 = ContainerServiceMasterProfile.dnsPrefix
  enable_pod_security_policy = ManagedClusterProperties.enablePodSecurityPolicy
  kubernetes_version         = ManagedClusterProperties.kubernetesVersion
  node_resource_group        = ManagedClusterProperties.nodeResourceGroup

  default_node_pool = {
    vm_size = any //ContainerServiceVMSize
  }
}
