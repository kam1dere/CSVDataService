package app

import (
	"context"
	"fmt"
	"github.com/kam1dere/CSVDataService/config"
	"github.com/kam1dere/CSVDataService/grpc/genproto/CsvDataService"
	"github.com/kam1dere/CSVDataService/pkg/dataFormCSV"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataService struct {
	log    *zerolog.Logger
	config *config.Config

	CsvDataService.UnimplementedCsvDataServiceServer
}

func NewDataService(
	logger *zerolog.Logger,
	cfg *config.Config,
) (*DataService, error) {
	return &DataService{
		log:    logger,
		config: cfg,
	}, nil
}

func (d *DataService) GetItems(
	ctx context.Context,
	req *CsvDataService.GetItemsRequest,
) (*CsvDataService.GetItemsResponse, error) {

	dataFromCSV, err := dataFormCSV.LoadDataFromCSV()
	if err != nil {
		d.log.Err(err).Msg("load data form csv")
		return nil, status.Error(codes.Internal, ErrUnknown)
	}

	items := make([]*CsvDataService.Item, 0, len(req.Ids))
	for _, id := range req.Ids {
		if data, ok := dataFromCSV[id]; ok {
			items = append(items, parseDataFromCsvInGrpcItems(data))
			continue
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("item with ID=%d not found", id))
	}

	return &CsvDataService.GetItemsResponse{Items: items}, nil
}

func parseDataFromCsvInGrpcItems(data *dataFormCSV.Item) *CsvDataService.Item {
	return &CsvDataService.Item{
		Id:                        data.Id,
		Uid:                       data.Uid,
		Domain:                    data.Domain,
		Cn:                        data.Cn,
		Department:                data.Department,
		Title:                     data.Title,
		Who:                       data.Who,
		LogonCount:                data.LogonCount,
		NumLogons7:                data.NumLogons7,
		NumShare7:                 data.NumShare7,
		NumFile7:                  data.NumFile7,
		NumAd7:                    data.NumAd7,
		NumN7:                     data.NumN7,
		NumLogons14:               data.NumLogons14,
		NumShare14:                data.NumShare14,
		NumFile14:                 data.NumFile14,
		NumAd14:                   data.NumAd14,
		NumN14:                    data.NumN14,
		NumLogons30:               data.NumLogons30,
		NumShare30:                data.NumShare30,
		NumFile30:                 data.NumFile30,
		NumAd30:                   data.NumAd30,
		NumN30:                    data.NumN30,
		NumLogons150:              data.NumLogons150,
		NumShare150:               data.NumShare150,
		NumFile150:                data.NumFile150,
		NumAd150:                  data.NumAd150,
		NumN150:                   data.NumN150,
		NumLogons365:              data.NumLogons365,
		NumShare365:               data.NumShare365,
		NumFile365:                data.NumFile365,
		NumAd365:                  data.NumAd365,
		NumN365:                   data.NumN365,
		HasUserPrincipalName:      data.HasUserPrincipalName,
		HasMail:                   data.HasMail,
		HasPhone:                  data.HasPhone,
		FlagDisabled:              data.FlagDisabled,
		FlagLockout:               data.FlagLockout,
		FlagPasswordNotRequired:   data.FlagPasswordNotRequired,
		FlagPasswordCantChange:    data.FlagPasswordCantChange,
		FlagDontExpirePassword:    data.FlagDontExpirePassword,
		OwnedFiles:                data.OwnedFiles,
		NumMailboxes:              data.NumMailboxes,
		NumMemberOfGroups:         data.NumMemberOfGroups,
		NumMemberOfIndirectGroups: data.NumMemberOfIndirectGroups,
		MemberOfIndirectGroupsIds: data.MemberOfIndirectGroupsIds,
		MemberOfGroupsIds:         data.MemberOfGroupsIds,
		IsAdmin:                   data.IsAdmin,
		IsService:                 data.IsService,
	}
}
