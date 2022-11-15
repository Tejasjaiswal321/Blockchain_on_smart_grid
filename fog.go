




// FogNode describes basic details of what makes up a simple fognode

//Insert struct field in alphabetic order => to achieve determinism across languages

// golang keeps the order when marshal to json but doesn't order automatically

type FogNode struct {

	ID             string `json:"ID"`
	
	H_pk          string `json:"H_pk"`
	
	Sec_param          string `json:"Sec_param"`

}



// InitLedger adds a base set of fognodes to the ledger

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	fognodes := []FogNode{

		{ID: "fognode1", Sec_param: "654365734652345", H_pk: "fwerq23rffawrtfaewfaewfwaef"},

		{ID: "fognode2", Sec_param: "532452345t235", H_pk: "htikuilkuhijmkyfhjngf"},

		{ID: "fognode3", Sec_param: "421345321", H_pk: "gryjertjyukfjft"},

		{ID: "fognode4", Sec_param: "53245134521", H_pk: "jtdhgjndrthbrfdbdf"},

		// {ID: "fognode5", Sec_param: "black"5, H_pk: "Adriana"},

		// {ID: "fognode6", Sec_param: "white"5, H_pk: "Michel"},

	}



	for _, fognode := range fognodes {

		fognodeJSON, err := json.Marshal(fognode)

		if err != nil {

			return err

		}



		err = ctx.GetStub().PutState(fognode.ID, fognodeJSON)

		if err != nil {

			return fmt.Errorf("failed to put to world state. %v", err)

		}

	}



	return nil

}



// AddFogNode issues a new fognode to the world state with given details.

func (s *SmartContract) AddFogNode(ctx contractapi.TransactionContextInterface, id string, sec_param string, h_pk string) error {

	exists, err := s.FogNodeExists(ctx, id)

	if err != nil {

		return err

	}

	if exists {

		return fmt.Errorf("the fognode %s already exists", id)

	}



	fognode := FogNode{

		ID:             id,

		Sec_param:          sec_param,

		H_pk:          h_pk,


	}

	fognodeJSON, err := json.Marshal(fognode)

	if err != nil {

		return err

	}



	return ctx.GetStub().PutState(id, fognodeJSON)

}



// ReadFogNode returns the fognode stored in the world state with given id.

func (s *SmartContract) ReadFogNode(ctx contractapi.TransactionContextInterface, id string) (*FogNode, error) {

	fognodeJSON, err := ctx.GetStub().GetState(id)

	if err != nil {

		return nil, fmt.Errorf("failed to read from world state: %v", err)

	}

	if fognodeJSON == nil {

		return nil, fmt.Errorf("the fognode %s does not exist", id)

	}



	var fognode FogNode

	err = json.Unmarshal(fognodeJSON, &fognode)

	if err != nil {

		return nil, err

	}



	return &fognode, nil

}



// UpdateFogNode updates an existing fognode in the world state with provided parafognodes.

func (s *SmartContract) UpdateFogNode(ctx contractapi.TransactionContextInterface, id string, sec_param string, h_pk string) error {

	exists, err := s.FogNodeExists(ctx, id)

	if err != nil {

		return err

	}

	if !exists {

		return fmt.Errorf("the fognode %s does not exist", id)

	}



	// overwriting original fognode with new fognode

	fognode := FogNode{

		ID:             id,

		Sec_param:          sec_param,


		H_pk:          h_pk,


	}

	fognodeJSON, err := json.Marshal(fognode)

	if err != nil {

		return err

	}



	return ctx.GetStub().PutState(id, fognodeJSON)

}



// RemoveFogNode deletes an given fognode from the world state.

func (s *SmartContract) RemoveFogNode(ctx contractapi.TransactionContextInterface, id string) error {

	exists, err := s.FogNodeExists(ctx, id)

	if err != nil {

		return err

	}

	if !exists {

		return fmt.Errorf("the fognode %s does not exist", id)

	}



	return ctx.GetStub().DelState(id)

}



// FogNodeExists returns true when fognode with given ID exists in world state

func (s *SmartContract) FogNodeExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	fognodeJSON, err := ctx.GetStub().GetState(id)

	if err != nil {

		return false, fmt.Errorf("failed to read from world state: %v", err)

	}



	return fognodeJSON != nil, nil

}



// TransferFogNode updates the h_pk field of fognode with given id in world state, and returns the old h_pk.





// GetAllFogNodes returns all fognodes found in world state

func (s *SmartContract) GetAllFogNodes(ctx contractapi.TransactionContextInterface) ([]*FogNode, error) {

	// range query with empty string for startSec_param and endSec_param does an

	// open-ended query of all fognodes in the chaincode namespace.

	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")

	if err != nil {

		return nil, err

	}

	defer resultsIterator.Close()



	var fognodes []*FogNode

	for resultsIterator.HasNext() {

		queryResponse, err := resultsIterator.Next()

		if err != nil {

			return nil, err

		}



		var fognode FogNode

		err = json.Unmarshal(queryResponse.Value, &fognode)

		if err != nil {

			return nil, err

		}

		fognodes = append(fognodes, &fognode)

	}



	return fognodes, nil

}