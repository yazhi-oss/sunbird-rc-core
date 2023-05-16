// Code generated by go-swagger; DO NOT EDIT.

package download_file_report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"bulk_issuance/swagger_gen/models"
)

// GetV1IDReportOKCode is the HTTP code returned for type GetV1IDReportOK
const GetV1IDReportOKCode int = 200

/*
GetV1IDReportOK OK

swagger:response getV1IdReportOK
*/
type GetV1IDReportOK struct {
	/*

	 */
	ContentDisposition string `json:"Content-Disposition"`

	/*
	  In: Body
	*/
	Payload models.FileDownload `json:"body,omitempty"`
}

// NewGetV1IDReportOK creates GetV1IDReportOK with default headers values
func NewGetV1IDReportOK() *GetV1IDReportOK {

	return &GetV1IDReportOK{}
}

// WithContentDisposition adds the contentDisposition to the get v1 Id report o k response
func (o *GetV1IDReportOK) WithContentDisposition(contentDisposition string) *GetV1IDReportOK {
	o.ContentDisposition = contentDisposition
	return o
}

// SetContentDisposition sets the contentDisposition to the get v1 Id report o k response
func (o *GetV1IDReportOK) SetContentDisposition(contentDisposition string) {
	o.ContentDisposition = contentDisposition
}

// WithPayload adds the payload to the get v1 Id report o k response
func (o *GetV1IDReportOK) WithPayload(payload models.FileDownload) *GetV1IDReportOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 Id report o k response
func (o *GetV1IDReportOK) SetPayload(payload models.FileDownload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1IDReportOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Disposition

	contentDisposition := o.ContentDisposition
	if contentDisposition != "" {
		rw.Header().Set("Content-Disposition", contentDisposition)
	}

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetV1IDReportForbiddenCode is the HTTP code returned for type GetV1IDReportForbidden
const GetV1IDReportForbiddenCode int = 403

/*
GetV1IDReportForbidden Forbidden

swagger:response getV1IdReportForbidden
*/
type GetV1IDReportForbidden struct {
	/*

	  Default: "application/json"
	*/
	ContentType string `json:"Content-Type"`

	/*
	  In: Body
	*/
	Payload *models.ErrorPayload `json:"body,omitempty"`
}

// NewGetV1IDReportForbidden creates GetV1IDReportForbidden with default headers values
func NewGetV1IDReportForbidden() *GetV1IDReportForbidden {

	var (
		// initialize headers with default values

		contentTypeDefault = string("application/json")
	)

	return &GetV1IDReportForbidden{

		ContentType: contentTypeDefault,
	}
}

// WithContentType adds the contentType to the get v1 Id report forbidden response
func (o *GetV1IDReportForbidden) WithContentType(contentType string) *GetV1IDReportForbidden {
	o.ContentType = contentType
	return o
}

// SetContentType sets the contentType to the get v1 Id report forbidden response
func (o *GetV1IDReportForbidden) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithPayload adds the payload to the get v1 Id report forbidden response
func (o *GetV1IDReportForbidden) WithPayload(payload *models.ErrorPayload) *GetV1IDReportForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 Id report forbidden response
func (o *GetV1IDReportForbidden) SetPayload(payload *models.ErrorPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1IDReportForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Type

	contentType := o.ContentType
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetV1IDReportNotFoundCode is the HTTP code returned for type GetV1IDReportNotFound
const GetV1IDReportNotFoundCode int = 404

/*
GetV1IDReportNotFound Not found

swagger:response getV1IdReportNotFound
*/
type GetV1IDReportNotFound struct {
	/*

	  Default: "application/json"
	*/
	ContentType string `json:"Content-Type"`

	/*
	  In: Body
	*/
	Payload *models.ErrorPayload `json:"body,omitempty"`
}

// NewGetV1IDReportNotFound creates GetV1IDReportNotFound with default headers values
func NewGetV1IDReportNotFound() *GetV1IDReportNotFound {

	var (
		// initialize headers with default values

		contentTypeDefault = string("application/json")
	)

	return &GetV1IDReportNotFound{

		ContentType: contentTypeDefault,
	}
}

// WithContentType adds the contentType to the get v1 Id report not found response
func (o *GetV1IDReportNotFound) WithContentType(contentType string) *GetV1IDReportNotFound {
	o.ContentType = contentType
	return o
}

// SetContentType sets the contentType to the get v1 Id report not found response
func (o *GetV1IDReportNotFound) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithPayload adds the payload to the get v1 Id report not found response
func (o *GetV1IDReportNotFound) WithPayload(payload *models.ErrorPayload) *GetV1IDReportNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get v1 Id report not found response
func (o *GetV1IDReportNotFound) SetPayload(payload *models.ErrorPayload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetV1IDReportNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Content-Type

	contentType := o.ContentType
	if contentType != "" {
		rw.Header().Set("Content-Type", contentType)
	}

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
