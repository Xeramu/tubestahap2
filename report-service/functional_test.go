//go:build functional

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDailyReportEndpoint_ReturnsOK(t *testing.T) {
	svc := NewReportService()
	h := NewReportHandler(svc)

	req := httptest.NewRequest("GET", "/report/daily?date=2026-04-25", nil)
	w := httptest.NewRecorder()

	h.DailyReport(w, req)

	// Pastikan status 200 dulu (tapi nanti akan gagal karena data kosong)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// 🔥 CEK DATA TIDAK KOSONG (ini yang akan membuat test FAIL)
	var report DailyReport
	err := json.NewDecoder(w.Body).Decode(&report)
	if err != nil {
		t.Errorf("Failed to decode: %v", err)
	}
	if report.TotalPaket == 0 {
		t.Errorf("Expected TotalPaket > 0, got 0") // ← akan gagal
	}
}
