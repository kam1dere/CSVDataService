package dataFormCSV

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func LoadDataFromCSV() (map[int64]*Item, error) {
	file, err := os.Open("pkg/dataFromCSV/ueba.csv")
	if err != nil {
		return nil, fmt.Errorf("error opening CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data := make(map[int64]*Item)

	for lineNum := 1; ; lineNum++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read err line %d: %s", lineNum, err.Error())
		}

		id, err := strconv.ParseInt(record[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse int err %d: %s", lineNum, err.Error())
		}

		item := &Item{
			Id:                        id,
			Uid:                       record[2],
			Domain:                    record[3],
			Cn:                        record[4],
			Department:                record[5],
			Title:                     record[6],
			Who:                       record[7],
			LogonCount:                parseInt64(record[8]),
			NumLogons7:                parseInt64(record[9]),
			NumShare7:                 parseInt64(record[10]),
			NumFile7:                  parseInt64(record[11]),
			NumAd7:                    parseInt64(record[12]),
			NumN7:                     parseInt64(record[13]),
			NumLogons14:               parseInt64(record[14]),
			NumShare14:                parseInt64(record[15]),
			NumFile14:                 parseInt64(record[16]),
			NumAd14:                   parseInt64(record[17]),
			NumN14:                    parseInt64(record[18]),
			NumLogons30:               parseInt64(record[19]),
			NumShare30:                parseInt64(record[20]),
			NumFile30:                 parseInt64(record[21]),
			NumAd30:                   parseInt64(record[22]),
			NumN30:                    parseInt64(record[23]),
			NumLogons150:              parseInt64(record[24]),
			NumShare150:               parseInt64(record[25]),
			NumFile150:                parseInt64(record[26]),
			NumAd150:                  parseInt64(record[27]),
			NumN150:                   parseInt64(record[28]),
			NumLogons365:              parseInt64(record[29]),
			NumShare365:               parseInt64(record[30]),
			NumFile365:                parseInt64(record[31]),
			NumAd365:                  parseInt64(record[32]),
			NumN365:                   parseInt64(record[33]),
			HasUserPrincipalName:      parseBool(record[34]),
			HasMail:                   parseBool(record[35]),
			HasPhone:                  parseBool(record[36]),
			FlagDisabled:              parseBool(record[37]),
			FlagLockout:               parseBool(record[38]),
			FlagPasswordNotRequired:   parseBool(record[39]),
			FlagPasswordCantChange:    parseBool(record[40]),
			FlagDontExpirePassword:    parseBool(record[41]),
			OwnedFiles:                parseInt64(record[42]),
			NumMailboxes:              parseInt64(record[43]),
			NumMemberOfGroups:         parseInt64(record[44]),
			NumMemberOfIndirectGroups: parseInt64(record[45]),
			MemberOfIndirectGroupsIds: parseSlice(record[46]),
			MemberOfGroupsIds:         parseSlice(record[47]),
			IsAdmin:                   parseBool(record[48]),
			IsService:                 parseBool(record[49]),
		}

		data[id] = item
	}

	return data, nil
}

func parseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func parseSlice(s string) []string {
	return strings.Split(s, ";")
}

func parseBool(s string) bool {
	if s == "1" {
		return true
	}
	return false
}
