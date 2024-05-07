package dpfm_api_output_formatter

import (
	"data-platform-api-act-purpose-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToActPurpose(rows *sql.Rows) (*[]ActPurpose, error) {
	defer rows.Close()
	actPurpose := make([]ActPurpose, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ActPurpose{}

		err := rows.Scan(
			&pm.ActPurpose,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &actPurpose, nil
		}

		data := pm
		actPurpose = append(actPurpose, ActPurpose{
			ActPurpose:				data.ActPurpose,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &actPurpose, nil
}

func ConvertToText(rows *sql.Rows) (*[]Text, error) {
	defer rows.Close()
	text := make([]Text, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Text{}

		err := rows.Scan(
			&pm.ActPurpose,
			&pm.Language,
			&pm.ActPurposeName,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &text, err
		}

		data := pm
		text = append(text, Text{
			ActPurpose:     		data.ActPurpose,
			Language:          		data.Language,
			ActPurposeName:			data.ActPurposeName,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
		})
	}

	return &text, nil
}
