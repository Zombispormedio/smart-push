package response

import (
	"net/http"

	"github.com/labstack/echo"
)

type MessageT struct{
    Status int  `json:"status"`
    Message string `json:"message"`
}

type ErrorMessageT struct{
    Status int  `json:"status"`
    Error string `json:"error"`
}

type DataT struct{
    Status int  `json:"status"`
    Data interface{} `json:"data"`
}

func Success(e echo.Context, message string) error {
    msg:=MessageT{
        Status: 0,
        Message: message,
    }
   
	return e.JSON(http.StatusOK, msg)
}

func ExpectFail(e echo.Context, message string) error {
    
     msg:=ErrorMessageT{
        Status: 1,
        Error:message,
    }
   
	return e.JSON(http.StatusExpectationFailed, msg)
}

func Forbidden(e echo.Context, message string) error {
    
     msg:=ErrorMessageT{
        Status: 1,
        Error:message,
    }
   
	return e.JSON(http.StatusForbidden, msg)
}
	
    
    