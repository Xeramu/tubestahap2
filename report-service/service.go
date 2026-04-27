package main

import "context"

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) GetDailyReport(ctx context.Context, date string) (*DailyReport, error) {
	return nil, nil
}
