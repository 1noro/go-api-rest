package resources

import (
    "../model"
)

func createOkResponse() model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 200,
        Description: "OK",
    }
    return model.JSONHTTPResponse{
        HTTPResponse: response,
    }
}

func createOkResponseExtra(extraText string) model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 200,
        Description: "OK",
        ExtraText: extraText,
    }
    return model.JSONHTTPResponse{
        HTTPResponse: response,
    }
}

func createCreatedResponse() model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 201,
        Description: "Created",
    }
    return model.JSONHTTPResponse{
        HTTPResponse: response,
    }
}

func createUnauthorizedResponse() model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 401,
        Description: "Unauthorized",
    }
    return model.JSONHTTPResponse{
        HTTPResponse: response,
    }
}

func createNotFoundResponse() model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 404,
        Description: "Not Found",
    }
    return model.JSONHTTPResponse{
        HTTPResponse: response,
    }
}

func createConflictResponse() model.JSONHTTPResponse {
    response := model.HTTPResponse {
        Code: 409,
        Description: "Conflict",
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