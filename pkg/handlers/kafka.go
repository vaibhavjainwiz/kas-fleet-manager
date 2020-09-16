package handlers

import (
	"net/http"

	"gitlab.cee.redhat.com/service/managed-services-api/pkg/api/openapi"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/api/presenters"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/errors"
	"gitlab.cee.redhat.com/service/managed-services-api/pkg/services"
)

type kafkaHandler struct {
	service services.KafkaService
}

func NewKafkaHandler(service services.KafkaService) *kafkaHandler {
	return &kafkaHandler{
		service: service,
	}
}

func (h kafkaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var kafkaRequest openapi.KafkaRequest
	cfg := &handlerConfig{
		MarshalInto: &kafkaRequest,
		Validate: []validate{
			validateEmpty(&kafkaRequest.Id, "id"),
			validateNotEmpty(&kafkaRequest.Region, "region"),
			validateNotEmpty(&kafkaRequest.CloudProvider, "cloud_provider"),
			validateNotEmpty(&kafkaRequest.Name, "cluster_name"),
		},
		Action: func() (interface{}, *errors.ServiceError) {
			convKafka := presenters.ConvertKafkaRequest(kafkaRequest)
			err := h.service.RegisterKafkaJob(convKafka)
			if err != nil {
				return nil, err
			}
			return presenters.PresentKafkaRequest(convKafka), nil
		},
		ErrorHandler: handleError,
	}

	// return 202 status accepted
	handle(w, r, cfg, http.StatusAccepted)
}