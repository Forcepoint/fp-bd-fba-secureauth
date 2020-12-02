package structs

type MonitoredEntities struct {
	Entities []Entity `json:"entities"`
}

type Entity struct {
	ID string `json:"id"`
	ActorID string `json:"actor_id"`
}