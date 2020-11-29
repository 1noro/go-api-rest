package model

// HTTPResponse el contenido de la JSONHTTPResponse
type HTTPResponse struct {
    Code int `json:"code"`
    Description string `json:"description"`
    ExtraText string `json:"extra_text"`
}

// JSONHTTPResponse respuesta JSON que indica el código HTTP y algún comentario extra
type JSONHTTPResponse struct {
    HTTPResponse HTTPResponse `json:"http_response"`
}
