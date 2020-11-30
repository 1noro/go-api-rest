package resources

import (
    "../model"
)

func createInternalServerErrorResponse() model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 500,
        Description: "Internal Server Error",
    }
    return model.JSONHTTPResponse{
        HTTPResponse: response,
    }
}