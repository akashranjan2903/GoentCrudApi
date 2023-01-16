package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewApiResSuccess() ApiResSuccess {
	return ApiResSuccess{}
}

func apiResFuncSuccess(c *gin.Context, message string, data any, status int) {
	c.JSON(status, ApiResSuccess{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// The request succeeded. The result meaning of "success" depends on the HTTP method:
// GET: The resource has been fetched and transmitted in the message body.
// HEAD: The representation headers are included in the response without any message body.
// PUT or POST: The resource describing the result of the action is transmitted in the message body.
// TRACE: The message body contains the request message as received by the server.
func (s *ApiResSuccess) Ok200(c *gin.Context, message string, data any) {
	apiResFuncSuccess(c, message, data, http.StatusOK)
}

// The request succeeded, and a new resource was created as a result.
// This is typically the response sent after POST requests, or some PUT requests.
func (s *ApiResSuccess) Created201(
	c *gin.Context,
	message string,
	data interface{},
) {
	apiResFuncSuccess(c, message, data, http.StatusCreated)
}

// The request has been received but not yet acted upon. It is noncommittal, since there is no way in HTTP to later send an asynchronous response indicating the outcome of the request. It is intended for cases where another process or server handles the request, or for batch processing.
func (s *ApiResSuccess) Accepted202(
	c *gin.Context,
	message string,
	data interface{},
) {
	apiResFuncSuccess(c, message, data, http.StatusAccepted)
}

// There is no content to send for this request, but the headers may be useful. The user agent may update its cached headers for this resource with the new ones.
func (s *ApiResSuccess) NoContent204(c *gin.Context, message string) {
	apiResFuncSuccess(c, message, nil, http.StatusNoContent)
}

func (e *ApiResSuccess) Custom(c *gin.Context, message string, data any, status int) {
	apiResFuncSuccess(c, message, data, status)
}

// Error Responses
type ApiResError struct {
	Status  int    `json:"status"`
	Err     error  `json:"err"`
	Message string `json:"message"`
}

func NewApiResError() ApiResError {
	return ApiResError{}
}

func apiResFunError(c *gin.Context, err error, status int, message string) {

	c.JSON(status, ApiResError{
		Status:  status,
		Err:     err,
		Message: message,
	})
}

// The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).
func (e *ApiResError) BadRequest400(c *gin.Context, err error, message string) {
	apiResFunError(c, err, http.StatusBadRequest, message)
}

// Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated". That is, the client must authenticate itself to get the requested response.
func (e *ApiResError) Unauthorized401(c *gin.Context, err error, message string) {
	apiResFunError(c, err, http.StatusUnauthorized, message)
}

// This response code is reserved for future use. The initial aim for creating this code was using it for digital payment systems, however this status code is used very rarely and no standard convention exists.
func (e *ApiResError) PaymentRequired402(c *gin.Context, err error, message string) {
	apiResFunError(c, err, http.StatusPaymentRequired, message)
}

// The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource. Unlike 401 Unauthorized, the client's identity is known to the server.
func (e *ApiResError) Forbidden403(c *gin.Context, err error, message string) {
	apiResFunError(c, err, http.StatusForbidden, message)
}

// The server can not find the requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist. Servers may also send this response instead of 403 Forbidden to hide the existence of a resource from an unauthorized client. This response code is probably the most well known due to its frequent occurrence on the web.
func (e *ApiResError) NotFound404(c *gin.Context, err error, message string) {
	apiResFunError(c, err, http.StatusNotFound, message)
}

// The request method is known by the server but is not supported by the target resource. For example, an API may not allow calling DELETE to remove a resource
func (e *ApiResError) MethodNotAllowed405(c *gin.Context, err error, message string) {
	apiResFunError(c, err, http.StatusMethodNotAllowed, message)
}

// This response is sent on an idle connection by some servers, even without any previous request by the client. It means that the server would like to shut down this unused connection. This response is used much more since some browsers, like Chrome, Firefox 27+, or IE9, use HTTP pre-connection mechanisms to speed up surfing. Also note that some servers merely shut down the connection without sending this message.
func (e *ApiResError) RequestTimeout408(c *gin.Context, err error, message string) {
	apiResFunError(c, err, http.StatusRequestTimeout, message)
}

// This response is sent when a request conflicts with the current state of the server.
func (e *ApiResError) Conflict409(c *gin.Context, err error, message string) {
	apiResFunError(c, err, http.StatusConflict, message)
}

// InternalServerError500
func (e *ApiResError) InternalServerError500(c *gin.Context, err error) {
	apiResFunError(c, err, http.StatusInternalServerError, "something went wrong")
}

// Custom
func (e *ApiResError) Custom(c *gin.Context, err error, status int, message string) {
	apiResFunError(c, err, status, message)
}
