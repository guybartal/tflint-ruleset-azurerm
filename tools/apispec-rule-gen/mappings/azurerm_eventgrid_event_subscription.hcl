mapping "azurerm_eventgrid_event_subscription" {
  import_path = "azure-rest-api-specs/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/EventGrid.json"

  event_delivery_schema = EventSubscriptionProperties.eventDeliverySchema
  topic_name            = EventSubscriptionProperties.topic
  included_event_types  = EventSubscriptionFilter.includedEventTypes
  labels                = EventSubscriptionProperties.labels
}
