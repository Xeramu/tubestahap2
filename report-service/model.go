package main

type DailyReport struct {
	TotalPaket  int     `json:"total_paket"`
	Delivered   int     `json:"delivered"`
	Pending     int     `json:"pending"`
	Terlambat   int     `json:"terlambat"`
	RataRataETA float64 `json:"rata_rata_eta"`
}