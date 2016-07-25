package main

import (
	"encoding/json"
	"fmt"

	"github.com/egonelbre/exp/spec/soap"
)

var data = `
<Envelope xmlns:s='http://schemas.xmlsoap.org/soap/envelope/'>
  <Body>
    <GetAllColumnsResponse xmlns='http://replicon.com/'>
      <GetAllColumnsResult xmlns:i='http://www.w3.org/2001/XMLSchema-instance'>
        <ListColumnGroup1>
          <columns>
            <ListColumn1>
              <dataType>urn:replicon:list-type:string</dataType>
              <displayText>Project Name</displayText>
              <uri>urn:replicon:project-list-column:name</uri>
            </ListColumn1>
            <ListColumn1>
              <dataType>urn:replicon:list-type:string</dataType>
              <displayText>Code</displayText>
              <uri>urn:replicon:project-list-column:code</uri>
            </ListColumn1>
            <ListColumn1>
              <dataType>urn:replicon:list-type:object</dataType>
              <displayText>Status</displayText>
              <uri>urn:replicon:project-list-column:status</uri>
            </ListColumn1>
            <ListColumn1>
              <dataType>urn:replicon:list-type:object</dataType>
              <displayText>Client</displayText>
              <uri>urn:replicon:project-list-column:client</uri>
            </ListColumn1>
            <ListColumn1>
              <dataType>urn:replicon:list-type:object</dataType>
              <displayText>Program</displayText>
              <uri>urn:replicon:project-list-column:program</uri>
            </ListColumn1>
          </columns>
        </ListColumnGroup1>
      </GetAllColumnsResult>
    </GetAllColumnsResponse>
  </Body>
</Envelope>
`

type AllColumnsResponse struct {
	Columns []Column `json:"columns"`
}

type Column struct {
	DisplayText string `json:"displayText"`
	DataType    string `json:"dataType"`
	URI         string `json:"uri"`
}

func (response *AllColumnsResponse) Spec() soap.Spec {
	return soap.Tag("Envelope",
		soap.Tag("Body",
			soap.Tag("GetAllColumnsResponse",
				soap.Tag("GetAllColumnsResult",
					soap.Tag("ListColumnGroup1",
						soap.TagList("columns", &response.Columns),
					),
				),
			),
		),
	)
}

func (column *Column) Spec() soap.Spec {
	return soap.Tag("ListColumn1",
		soap.String("dataType", &column.DataType),
		soap.String("displayText", &column.DisplayText),
		soap.String("uri", &column.URI),
	)
}

func main() {
	node, err := soap.Parse([]byte(data))
	if err != nil {
		panic(err)
	}

	var response AllColumnsResponse
	response.Spec().Decode(node)

	data, err := json.MarshalIndent(&response, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("-----------------")
	fmt.Println(string(data))

	data, err = response.Spec().Encode().Encode()
	if err != nil {
		panic(err)
	}
	fmt.Println("-----------------")
	fmt.Println(string(data))
}
