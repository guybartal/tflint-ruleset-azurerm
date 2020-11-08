mapping "azurerm_bot_channel_ms_teams" {
  import_path = "azure-rest-api-specs/specification/botservice/resource-manager/Microsoft.BotService/preview/2018-07-12/botservice.json"

  resource_group_name = resourceGroupNameParameter
  bot_name            = resourceNameParameter
  calling_web_hook    = MsTeamsChannelProperties.callingWebHook
  enable_calling      = MsTeamsChannelProperties.enableCalling
}