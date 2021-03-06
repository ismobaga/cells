// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/pydio-sdk-go/models"
)

// AdminAddWorkspaceFeatureReader is a Reader for the AdminAddWorkspaceFeature structure.
type AdminAddWorkspaceFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminAddWorkspaceFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAdminAddWorkspaceFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAdminAddWorkspaceFeatureOK creates a AdminAddWorkspaceFeatureOK with default headers values
func NewAdminAddWorkspaceFeatureOK() *AdminAddWorkspaceFeatureOK {
	return &AdminAddWorkspaceFeatureOK{}
}

/*AdminAddWorkspaceFeatureOK handles this case with default header values.

Workspace object
*/
type AdminAddWorkspaceFeatureOK struct {
	Payload *models.PydioResponse
}

func (o *AdminAddWorkspaceFeatureOK) Error() string {
	return fmt.Sprintf("[POST /admin/workspaces/{workspaceId}/features/{metaId}][%d] adminAddWorkspaceFeatureOK  %+v", 200, o.Payload)
}

func (o *AdminAddWorkspaceFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PydioResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
