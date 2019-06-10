/*
 * Dashboards API
 *
 * API for creating, retrieving, updating, and deleting dashboards and dashboard groups. <br> Dashboards are groups of charts. In a dashboard, all the charts that belong to the dashboard appear at the same time and follow the same filtering options. ## Dashboard layout The system lays out dashboards and the charts they contain with these dimensions: <br>   * The web UI reserves a 12x100 grid for each dashboard and assigns one or     more charts to specific locations within the grid.   * A chart associated with the dashboard can be any size from 1x1 to 12x3.   * If you assign overlapping dashboard locations for charts, the system     attempts to resize or reorganize the layout. This ensures that all     of the charts fit within the space alloted to the dashboard.  ## Dashboard access By default, all users in an organization can edit and delete dashboards and dashboard groups. If SignalFx has enabled the \"write permissions\" feature for your organization, you can limit editing or deleting of specific dashboards to specific individuals or teams, or both. Use this feature to prevent unauthorized or accidental modifications to dashboards and the charts they contain. ## Cloning dashboards Users who don't have permission to edit a dashboard can still clone it and modify the clone. ## View dashboards You can view dashboards you create using the API in the web UI by specifying their \"id\" property in a web UI URL, by following this syntax: <br> <code>https://app.signalfx.com/#/dashboard/&lt;DASHBOARD_ID&gt;</code> <br> Dashboards you create using the API also appear by name in the web UI catalog and in their dashboard group.
 *
 * API version: 3.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dashboard

// A filter that selects charts to overlay with events, based on event names and types.
type DashboardEventSignal struct {
	// The event name or partial name that the system uses to selecct events to suggest as overlays on the charts in the dashboard.
	EventSearchText string `json:"eventSearchText"`
	// Controls the source of the event. You can specify the following: <br>   * detectorEvents: Select events that come from a detector   * eventTimeSeries: Select events that come from a time series <br> The API doesn't accept other event types.
	EventType string `json:"eventType,omitempty"`
}
