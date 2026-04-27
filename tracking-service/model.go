package main

// TrackingEvent merepresentasikan data tabel tracking_events
type TrackingEvent struct {
	ID        int    `json:"id"`
	Resi      string `json:"resi"`
	Lokasi    string `json:"lokasi"`
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
}

// TrackingResponse adalah output untuk endpoint GET /tracking/{resi}
type TrackingResponse struct {
	Resi     string          `json:"resi"`
	Status   string          `json:"status"`
	Timeline []TrackingEvent `json:"timeline"`
}

// DistanceRequest adalah input untuk POST /distance
type DistanceRequest struct {
	OriginAddress      string `json:"origin_address"`
	DestinationAddress string `json:"destination_address"`
}

// DistanceResponse adalah output untuk POST /distance
type DistanceResponse struct {
	DistanceKm    float64 `json:"distance_km"`
	DurationMin   int     `json:"duration_min"`
	PolylineRoute string  `json:"polyline_route"`
}

// RouteRequest adalah input untuk POST /route
type RouteRequest struct {
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	Waypoints   []string `json:"waypoints"`
}

// CourierLocation adalah output untuk GET /location/{courier_id}
type CourierLocation struct {
	CourierID string  `json:"courier_id"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}