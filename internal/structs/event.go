package structs

type Event struct {
	Type string `json:"type"`
	Timestamp string `json:"timestamp"`
	Entities []EntityRoles `json:"entities"`
	SourceEventID string `json:"source_event_id"`
	Subject string `json:"subject"`
	Labels []string `json:"labels"`
	Attributes []Attribute `json:"attributes"`
}

type EntityRoles struct {
	Role string `json:"role"`
	Entities []string `json:"entities"`
}

type Attribute interface {}

type StringAttribute struct {
	Attribute `json:"-"`
	Name string `json:"name"`
	Value string `json:"value"`
	Type string `json:"type"`
}

type BooleanAttribute struct {
	Attribute `json:"-"`
	Name string `json:"name"`
	Value bool `json:"value"`
	Type string `json:"type"`
}

type Attachments struct {
	Name string `json:"name"`
	Data string `json:"data"`
	ContentType string `json:"content_type"`
	ExtractedText string `json:"extracted_text"`
}

type EventReferences struct {
	ReferencesName string `json:"references_name"`
	EventIDField string `json:"event_id_field"`
	EventIDs []string `json:"event_ids"`
}

func CreateEvent(details [][]string, application string, realm string) Event {

	// Common details
	eventTypeID := details[2][2]
	eventType := EventTypes[eventTypeID]
	eventID := details[2][2] + details[9][2]
	timestamp := details[3][2]

	// Populate entities fields with relevant entities
	var eventEntities []EntityRoles

	entity1 := EntityRoles{
		Role:     "Vendor",
		Entities: []string{"SecureAuth"},
	}

	entity2 := EntityRoles{
		Role:     "App",
		Entities: []string{application},
	}

	entity3 := EntityRoles{
		Role:     "Realm",
		Entities: []string{realm},
	}

	eventEntities = append(eventEntities, entity1, entity2, entity3)

	if details[10][2] != "" {
		entity := EntityRoles{
			Role:     "Source IP",
			Entities: []string{details[10][2]},
		}
		eventEntities = append(eventEntities, entity)
	}

	if details[7][2] != "User" {
		entity := EntityRoles{
			Role:     "User",
			Entities: []string{details[7][2]},
		}
		eventEntities = append(eventEntities, entity)
	}

	// Populate attributes with details of event
	attribute1 := StringAttribute{
		Name:  "Timestamp",
		Value: timestamp,
		Type:  "String",
	}
	
	attribute2 := StringAttribute{
		Name:  "Event ID",
		Value: eventID,
		Type:  "String",
	}

	attribute3 := BooleanAttribute{
		Name:  "Success",
		Value: EventSuccess[eventTypeID],
		Type:  "Boolean",
	}

	attribute4 := StringAttribute{
		Name:       "Reason",
		Value:      eventType,
		Type:       "String",
	}

	var attributes []Attribute
	attributes = append(attributes, attribute1, attribute2, attribute3, attribute4)

	// Populate Event
	event := Event{
		Type:            "Login Attempt",
		Timestamp:       timestamp,
		Entities:        eventEntities,
		SourceEventID:   eventID,
		Subject:         eventType,
		Labels:          []string{"policy"},
		Attributes:      attributes,
	}

	return event
}