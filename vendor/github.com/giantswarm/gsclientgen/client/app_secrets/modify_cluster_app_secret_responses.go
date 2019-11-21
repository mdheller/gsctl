// Code generated by go-swagger; DO NOT EDIT.

package app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// ModifyClusterAppSecretReader is a Reader for the ModifyClusterAppSecret structure.
type ModifyClusterAppSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyClusterAppSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyClusterAppSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyClusterAppSecretBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewModifyClusterAppSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewModifyClusterAppSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyClusterAppSecretOK creates a ModifyClusterAppSecretOK with default headers values
func NewModifyClusterAppSecretOK() *ModifyClusterAppSecretOK {
	return &ModifyClusterAppSecretOK{}
}

/*ModifyClusterAppSecretOK handles this case with default header values.

Success
*/
type ModifyClusterAppSecretOK struct {
	Payload *models.V4GenericResponse
}

func (o *ModifyClusterAppSecretOK) Error() string {
	return fmt.Sprintf("[PATCH /v4/clusters/{cluster_id}/apps/{app_name}/secret/][%d] modifyClusterAppSecretOK  %+v", 200, o.Payload)
}

func (o *ModifyClusterAppSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyClusterAppSecretBadRequest creates a ModifyClusterAppSecretBadRequest with default headers values
func NewModifyClusterAppSecretBadRequest() *ModifyClusterAppSecretBadRequest {
	return &ModifyClusterAppSecretBadRequest{}
}

/*ModifyClusterAppSecretBadRequest handles this case with default header values.

Invalid input
*/
type ModifyClusterAppSecretBadRequest struct {
	Payload *models.V4GenericResponse
}

func (o *ModifyClusterAppSecretBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v4/clusters/{cluster_id}/apps/{app_name}/secret/][%d] modifyClusterAppSecretBadRequest  %+v", 400, o.Payload)
}

func (o *ModifyClusterAppSecretBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyClusterAppSecretUnauthorized creates a ModifyClusterAppSecretUnauthorized with default headers values
func NewModifyClusterAppSecretUnauthorized() *ModifyClusterAppSecretUnauthorized {
	return &ModifyClusterAppSecretUnauthorized{}
}

/*ModifyClusterAppSecretUnauthorized handles this case with default header values.

Permission denied
*/
type ModifyClusterAppSecretUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *ModifyClusterAppSecretUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v4/clusters/{cluster_id}/apps/{app_name}/secret/][%d] modifyClusterAppSecretUnauthorized  %+v", 401, o.Payload)
}

func (o *ModifyClusterAppSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyClusterAppSecretDefault creates a ModifyClusterAppSecretDefault with default headers values
func NewModifyClusterAppSecretDefault(code int) *ModifyClusterAppSecretDefault {
	return &ModifyClusterAppSecretDefault{
		_statusCode: code,
	}
}

/*ModifyClusterAppSecretDefault handles this case with default header values.

Error
*/
type ModifyClusterAppSecretDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the modify cluster app secret default response
func (o *ModifyClusterAppSecretDefault) Code() int {
	return o._statusCode
}

func (o *ModifyClusterAppSecretDefault) Error() string {
	return fmt.Sprintf("[PATCH /v4/clusters/{cluster_id}/apps/{app_name}/secret/][%d] modifyClusterAppSecret default  %+v", o._statusCode, o.Payload)
}

func (o *ModifyClusterAppSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}