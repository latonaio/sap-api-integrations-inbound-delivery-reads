package sap_api_input_reader

type EC_MC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	InboundDelivery struct {
		DeliveryDocument         string `json:"document_no"`
		ShipToParty              string `json:"deliver_to"`
		OriginalDeliveryQuantity string `json:"quantity"`
		ActualDeliveryQuantity   string `json:"picked_quantity"`
		Price                    string `json:"price"`
		Batch                    string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string `json:"api_schema"`
	MaterialCode            string `json:"material_code"`
	ShippingPoint           string `json:"plant/supplier"`
	Stock                   string `json:"stock"`
	DeliveryDocumentType    string `json:"document_type"`
	DeliveryDocument        string `json:"document_no"`
	PlannedGoodsIssueDate   string `json:"planned_date"`
	ActualGoodsMovementDate string `json:"validated_date"`
	Deleted                 bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	InboundDelivery struct {
		DeliveryDocument              string `json:"DeliveryDocument"`
		ReceivingLocationTimeZone     string `json:"ReceivingLocationTimeZone"`
		ActualDeliveryRoute           string `json:"ActualDeliveryRoute"`
		ActualGoodsMovementDate       string `json:"ActualGoodsMovementDate"`
		ActualGoodsMovementTime       string `json:"ActualGoodsMovementTime"`
		BillingDocumentDate           string `json:"BillingDocumentDate"`
		CompleteDeliveryIsDefined     bool   `json:"CompleteDeliveryIsDefined"`
		ConfirmationTime              string `json:"ConfirmationTime"`
		CreationDate                  string `json:"CreationDate"`
		CreationTime                  string `json:"CreationTime"`
		CustomerGroup                 string `json:"CustomerGroup"`
		DeliveryBlockReason           string `json:"DeliveryBlockReason"`
		DeliveryDate                  string `json:"DeliveryDate"`
		DeliveryDocumentBySupplier    string `json:"DeliveryDocumentBySupplier"`
		DeliveryDocumentType          string `json:"DeliveryDocumentType"`
		DeliveryIsInPlant             bool   `json:"DeliveryIsInPlant"`
		DeliveryPriority              string `json:"DeliveryPriority"`
		DeliveryTime                  string `json:"DeliveryTime"`
		DocumentDate                  string `json:"DocumentDate"`
		GoodsIssueOrReceiptSlipNumber string `json:"GoodsIssueOrReceiptSlipNumber"`
		GoodsIssueTime                string `json:"GoodsIssueTime"`
		HeaderBillgIncompletionStatus string `json:"HeaderBillgIncompletionStatus"`
		HeaderBillingBlockReason      string `json:"HeaderBillingBlockReason"`
		HeaderDelivIncompletionStatus string `json:"HeaderDelivIncompletionStatus"`
		HeaderGrossWeight             string `json:"HeaderGrossWeight"`
		HeaderNetWeight               string `json:"HeaderNetWeight"`
		HeaderPackingIncompletionSts  string `json:"HeaderPackingIncompletionSts"`
		HeaderPickgIncompletionStatus string `json:"HeaderPickgIncompletionStatus"`
		HeaderVolume                  string `json:"HeaderVolume"`
		HeaderVolumeUnit              string `json:"HeaderVolumeUnit"`
		HeaderWeightUnit              string `json:"HeaderWeightUnit"`
		IncotermsClassification       string `json:"IncotermsClassification"`
		IsExportDelivery              string `json:"IsExportDelivery"`
		LastChangeDate                string `json:"LastChangeDate"`
		LoadingDate                   string `json:"LoadingDate"`
		LoadingPoint                  string `json:"LoadingPoint"`
		LoadingTime                   string `json:"LoadingTime"`
		MeansOfTransport              string `json:"MeansOfTransport"`
		OrderCombinationIsAllowed     bool   `json:"OrderCombinationIsAllowed"`
		OrderID                       string `json:"OrderID"`
		PickedItemsLocation           string `json:"PickedItemsLocation"`
		PickingDate                   string `json:"PickingDate"`
		PickingTime                   string `json:"PickingTime"`
		PlannedGoodsIssueDate         string `json:"PlannedGoodsIssueDate"`
		ProposedDeliveryRoute         string `json:"ProposedDeliveryRoute"`
		ReceivingPlant                string `json:"ReceivingPlant"`
		RouteSchedule                 string `json:"RouteSchedule"`
		SalesDistrict                 string `json:"SalesDistrict"`
		SalesOffice                   string `json:"SalesOffice"`
		SalesOrganization             string `json:"SalesOrganization"`
		SDDocumentCategory            string `json:"SDDocumentCategory"`
		ShipmentBlockReason           string `json:"ShipmentBlockReason"`
		ShippingCondition             string `json:"ShippingCondition"`
		ShippingPoint                 string `json:"ShippingPoint"`
		ShippingType                  string `json:"ShippingType"`
		ShipToParty                   string `json:"ShipToParty"`
		SoldToParty                   string `json:"SoldToParty"`
		Supplier                      string `json:"Supplier"`
		TotalBlockStatus              string `json:"TotalBlockStatus"`
		TotalCreditCheckStatus        string `json:"TotalCreditCheckStatus"`
		TotalNumberOfPackage          string `json:"TotalNumberOfPackage"`
		TransactionCurrency           string `json:"TransactionCurrency"`
		TransportationGroup           string `json:"TransportationGroup"`
		TransportationPlanningDate    string `json:"TransportationPlanningDate"`
		TransportationPlanningStatus  string `json:"TransportationPlanningStatus"`
		TransportationPlanningTime    string `json:"TransportationPlanningTime"`
		UnloadingPointName            string `json:"UnloadingPointName"`
		Partner                       struct {
			PartnerFunction string `json:"PartnerFunction"`
			Supplier        string `json:"Supplier"`
			Customer        string `json:"Customer"`
			AddressID       string `json:"AddressID"`
			ContactPerson   string `json:"ContactPerson"`
			Personnel       string `json:"Personnel"`
			SDDocument      string `json:"SDDocument"`
			SDDocumentItem  string `json:"SDDocumentItem"`
			Address         struct {
				AddressID              string `json:"AddressID"`
				Building               string `json:"Building"`
				BusinessPartnerName1   string `json:"BusinessPartnerName1"`
				CityName               string `json:"CityName"`
				CorrespondenceLanguage string `json:"CorrespondenceLanguage"`
				Country                string `json:"Country"`
				FaxNumber              string `json:"FaxNumber"`
				Nation                 string `json:"Nation"`
				PhoneNumber            string `json:"PhoneNumber"`
				PostalCode             string `json:"PostalCode"`
				Region                 string `json:"Region"`
				StreetName             string `json:"StreetName"`
			} `json:"Address"`
		} `json:"Partner"`
		DeliveryDocumentItem struct {
			DeliveryDocumentItem           string `json:"DeliveryDocumentItem"`
			DeliveryDocumentItemCategory   string `json:"DeliveryDocumentItemCategory"`
			DeliveryDocumentItemText       string `json:"DeliveryDocumentItemText"`
			ActualDeliveredQtyInBaseUnit   string `json:"ActualDeliveredQtyInBaseUnit"`
			ActualDeliveryQuantity         string `json:"ActualDeliveryQuantity"`
			AdditionalCustomerGroup1       string `json:"AdditionalCustomerGroup1"`
			AdditionalCustomerGroup2       string `json:"AdditionalCustomerGroup2"`
			AdditionalCustomerGroup3       string `json:"AdditionalCustomerGroup3"`
			AdditionalCustomerGroup4       string `json:"AdditionalCustomerGroup4"`
			AdditionalCustomerGroup5       string `json:"AdditionalCustomerGroup5"`
			BaseUnit                       string `json:"BaseUnit"`
			Batch                          string `json:"Batch"`
			BatchBySupplier                string `json:"BatchBySupplier"`
			BOMExplosion                   string `json:"BOMExplosion"`
			BusinessArea                   string `json:"BusinessArea"`
			ControllingArea                string `json:"ControllingArea"`
			CostCenter                     string `json:"CostCenter"`
			CreationDate                   string `json:"CreationDate"`
			CreationTime                   string `json:"CreationTime"`
			DeliveryGroup                  string `json:"DeliveryGroup"`
			DeliveryQuantityUnit           string `json:"DeliveryQuantityUnit"`
			DeliveryRelatedBillingStatus   string `json:"DeliveryRelatedBillingStatus"`
			DistributionChannel            string `json:"DistributionChannel"`
			Division                       string `json:"Division"`
			GLAccount                      string `json:"GLAccount"`
			GoodsMovementReasonCode        string `json:"GoodsMovementReasonCode"`
			GoodsMovementStatus            string `json:"GoodsMovementStatus"`
			GoodsMovementType              string `json:"GoodsMovementType"`
			InternationalArticleNumber     string `json:"InternationalArticleNumber"`
			InventorySpecialStockType      string `json:"InventorySpecialStockType"`
			IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
			IsNotGoodsMovementsRelevant    string `json:"IsNotGoodsMovementsRelevant"`
			IssuingOrReceivingPlant        string `json:"IssuingOrReceivingPlant"`
			IssuingOrReceivingStorageLoc   string `json:"IssuingOrReceivingStorageLoc"`
			ItemBillingBlockReason         string `json:"ItemBillingBlockReason"`
			ItemBillingIncompletionStatus  string `json:"ItemBillingIncompletionStatus"`
			ItemDeliveryIncompletionStatus string `json:"ItemDeliveryIncompletionStatus"`
			ItemGdsMvtIncompletionSts      string `json:"ItemGdsMvtIncompletionSts"`
			ItemGeneralIncompletionStatus  string `json:"ItemGeneralIncompletionStatus"`
			ItemGrossWeight                string `json:"ItemGrossWeight"`
			ItemIsBillingRelevant          string `json:"ItemIsBillingRelevant"`
			ItemNetWeight                  string `json:"ItemNetWeight"`
			ItemPackingIncompletionStatus  string `json:"ItemPackingIncompletionStatus"`
			ItemPickingIncompletionStatus  string `json:"ItemPickingIncompletionStatus"`
			ItemVolume                     string `json:"ItemVolume"`
			ItemVolumeUnit                 string `json:"ItemVolumeUnit"`
			ItemWeightUnit                 string `json:"ItemWeightUnit"`
			LastChangeDate                 string `json:"LastChangeDate"`
			LoadingGroup                   string `json:"LoadingGroup"`
			Material                       string `json:"Material"`
			MaterialByCustomer             string `json:"MaterialByCustomer"`
			MaterialFreightGroup           string `json:"MaterialFreightGroup"`
			MaterialGroup                  string `json:"MaterialGroup"`
			MaterialIsBatchManaged         bool   `json:"MaterialIsBatchManaged"`
			OrderID                        string `json:"OrderID"`
			OrderItem                      string `json:"OrderItem"`
			OriginalDeliveryQuantity       string `json:"OriginalDeliveryQuantity"`
			PackingStatus                  string `json:"PackingStatus"`
			PartialDeliveryIsAllowed       string `json:"PartialDeliveryIsAllowed"`
			PickingConfirmationStatus      string `json:"PickingConfirmationStatus"`
			PickingStatus                  string `json:"PickingStatus"`
			Plant                          string `json:"Plant"`
			ProductAvailabilityDate        string `json:"ProductAvailabilityDate"`
			ProductAvailabilityTime        string `json:"ProductAvailabilityTime"`
			ProfitabilitySegment           string `json:"ProfitabilitySegment"`
			ProfitCenter                   string `json:"ProfitCenter"`
			QuantityIsFixed                bool   `json:"QuantityIsFixed"`
			ReceivingPoint                 string `json:"ReceivingPoint"`
			ReferenceSDDocument            string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem        string `json:"ReferenceSDDocumentItem"`
			SalesDocumentItemType          string `json:"SalesDocumentItemType"`
			SalesGroup                     string `json:"SalesGroup"`
			SalesOffice                    string `json:"SalesOffice"`
			SDDocumentCategory             string `json:"SDDocumentCategory"`
			SDProcessStatus                string `json:"SDProcessStatus"`
			ShelfLifeExpirationDate        string `json:"ShelfLifeExpirationDate"`
			StockType                      string `json:"StockType"`
			StorageLocation                string `json:"StorageLocation"`
			TransportationGroup            string `json:"TransportationGroup"`
			UnlimitedOverdeliveryIsAllowed bool   `json:"UnlimitedOverdeliveryIsAllowed"`
		} `json:"DeliveryDocumentItem"`
	} `json:"DeliveryDocument"`
	APISchema  string   `json:"api_schema"`
	Accepter   []string `json:"accepter"`
	SDDocument string   `json:"delivery_document"`
	Deleted    string   `json:"deleted"`
}
