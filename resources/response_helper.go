package resources

import (
    "../model"
)

func createNotFoundResponse() model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 404,
        Description: "Not Found",
    }
    return model.JSONHTTPResponse{
        HTTPResponse: response,
    }
}

func createInternalServerErrorResponse() model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 500,
        Description: "Internal Server Error",
    }
    return model.JSONHTTPResponse{
        HTTPResponse: response,
    }
}