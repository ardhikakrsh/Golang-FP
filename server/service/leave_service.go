package service

import (
	"fmt"
	"time"

	"github.com/biskitsx/Leave-Management-System/logs"
	"github.com/biskitsx/Leave-Management-System/repository"
	"gorm.io/gorm"
)

type leaveService struct {
	db *gorm.DB
}

func NewLeaveService(db *gorm.DB) LeaveService {
	db.AutoMigrate(&repository.Leave{})
	return &leaveService{
		db: db,
	}
}

func (s leaveService) AddLeave(req AddLeaveRequest, userId uint) (*LeaveResponse, error) {
	timeStartDate, err := time.Parse(time.RFC3339, req.TimeStart)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	timeEndDate, err := time.Parse(time.RFC3339, req.TimeEnd)
	if err != nil {
		return nil, err
	}
	fmt.Println(timeEndDate)

	leave := repository.Leave{
		TimeStart: timeStartDate,
		TimeEnd:   timeEndDate,
		Type:      req.Type,
		Detail:    req.Detail,
		UserID:    userId,
		LeaveDay:  uint(timeEndDate.Sub(timeStartDate).Hours()/24) + 1,
		Status:	"requested",
	}
	s.db.Create(&leave)

	return &LeaveResponse{
		ID:        leave.ID,
		Type:      leave.Type,
		Detail:    leave.Detail,
		TimeStart: leave.TimeStart.Format(time.RFC3339),
		TimeEnd:   leave.TimeEnd.Format(time.RFC3339),
		CreatedAt: leave.CreatedAt.Format(time.RFC3339),
		UpdatedAt: leave.UpdatedAt.Format(time.RFC3339),
		LeaveDay:  leave.LeaveDay,
		Status:   leave.Status,
	}, nil
}

func (s leaveService) GetLeavesByUser(userId uint) (*LeaveResponseWithCount, error) {
	var leaves []repository.Leave
	s.db.Where("user_id = ?", userId).Find(&leaves)
	countSick := 0
	countBusiness := 0
	countVacation := 0

	for _, leave := range leaves {
		if leave.Type == "sakit" {
			countSick += 1
		} else if leave.Type == "absen" {
			countBusiness += 1
		} else if leave.Type == "liburan" {
			countVacation += 1
		}
	}
	var res []LeaveResponse
	for _, leave := range leaves {
		res = append(res, LeaveResponse{
			ID:        leave.ID,
			Type:      leave.Type,
			Detail:    leave.Detail,
			TimeStart: leave.TimeStart.Format("2006-01-02"),
			TimeEnd:   leave.TimeEnd.Format("2006-01-02"),
			CreatedAt: leave.CreatedAt.Format(time.RFC3339),
			UpdatedAt: leave.UpdatedAt.Format(time.RFC3339),
			LeaveDay:  leave.LeaveDay,
			Status:   leave.Status,
		})
	}
	return &LeaveResponseWithCount{
		LeaveResponses: res,
		CountSick:      uint(countSick),
		CountBusiness:  uint(countBusiness),
		CountVacation:  uint(countVacation),
	}, nil
}

func (s leaveService) GetLeaves() ([]LeaveResponse, error) {
	var leaves []repository.Leave
	s.db.Find(&leaves)
	var res []LeaveResponse
	for _, leave := range leaves {
		res = append(res, LeaveResponse{
			ID:        leave.ID,
			Type:      leave.Type,
			Detail:    leave.Detail,
			TimeStart: leave.TimeStart.Format("2006-01-02"),
			TimeEnd:   leave.TimeEnd.Format("2006-01-02"),
			CreatedAt: leave.CreatedAt.Format(time.RFC3339),
			UpdatedAt: leave.UpdatedAt.Format(time.RFC3339),
			LeaveDay:  leave.LeaveDay,
			Status:   leave.Status,
		})
	}
	return res, nil
}
