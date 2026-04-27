package main

import "testing"

// MockResiValidator adalah mock manual dari interface ResiValidator
type MockResiValidator struct{}

func (m MockResiValidator) CheckResi(resi string) bool {
	// Selalu mereturn true untuk mensimulasikan bahwa resi valid di Order Service
	return true
}

func TestInsertTrackingEvent(t *testing.T) {
	mock := MockResiValidator{}

	req := TrackingEvent{
		Resi:      "RESI123",
		Lokasi:    "Gudang Jakarta",
		Event:     "Paket diterima di gudang",
		Timestamp: "2026-04-27T10:00:00Z",
	}

	event, err := InsertTrackingEvent(req, mock)

	if err != nil {
		t.Fatal(err)
	}

	// Test akan FAILED di sini karena InsertTrackingEvent di service.go 
	// saat ini me-return TrackingEvent{} (kosong)
	if event.Resi != "RESI123" {
		t.Fail()
	}
}

func TestGetTrackingStatus(t *testing.T) {
	resi := "RESI123"

	resp := GetTrackingStatus(resi)

	// Test akan FAILED di sini karena GetTrackingStatus di service.go
	// saat ini me-return nil
	if resp == nil {
		t.Fatal("Expected response but got nil")
	}

	if resp.Resi != resi {
		t.Fail()
	}
}

func TestCalculateDistance(t *testing.T) {
	req := DistanceRequest{
		OriginAddress:      "Bandung",
		DestinationAddress: "Jakarta",
	}

	resp := CalculateDistance(req)

	// Test akan FAILED di sini karena CalculateDistance di service.go
	// saat ini me-return nil
	if resp == nil {
		t.Fatal("Expected response but got nil")
	}

	if resp.DistanceKm <= 0 {
		t.Fail()
	}
}