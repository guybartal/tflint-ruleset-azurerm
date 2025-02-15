mapping "azurerm_network_watcher_flow_log" {
  import_path = "azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/networkWatcher.json"

  network_security_group_id = FlowLogPropertiesFormat.targetResourceId
  storage_account_id        = FlowLogPropertiesFormat.storageId
  enabled                   = FlowLogPropertiesFormat.enabled
  version                   = FlowLogFormatParameters.version
}
