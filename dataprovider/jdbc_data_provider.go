package dataprovider

import (
    "strconv"
    "strings"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // go get -u github.com/go-sql-driver/mysql
    "../model"
)

// JDBCDataProvider conexión entre el API REST y la base de datos
type JDBCDataProvider struct {}

var dbServer = "mysql"
var dbUsername = "developer"
var dbPass = "pass"
var dbProtocol = "tcp"
var dbURL = "gitlab.afundacionfp.com"
var dbPort = "3306"
var dbName = "mysql"

// GetProducts devuelve una lista de productos para el catálogo
func (jdbcDataProvider JDBCDataProvider) GetProducts() []model.Product {
    // Open up our database connection.
    // db, err := sql.Open("mysql", "developer:pass@tcp(gitlab.afundacionfp.com:3306)/mysql")
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT referencia, nombre, urlImagen, CONCAT(SUBSTRING(descripcion, 1, 117), '...') FROM TablaCamiones")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    var products []model.Product
    for results.Next() {
        var product model.Product
        // for each row, scan the result into our tag composite object
        err = results.Scan(&product.Reference, &product.Name, &product.ImagePath, &product.ShortDescription)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        products = append(products, product)
    }
    return products
}

// GetFullProduct devuelve el detalle de un producto
func (jdbcDataProvider JDBCDataProvider) GetFullProduct(reference string) model.Product {
    // Open up our database connection.
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

    var product model.Product
    var price []uint8

    // Execute the query
    err = db.QueryRow("SELECT referencia, nombre, urlImagen, CONCAT(SUBSTRING(descripcion, 1, 117), '...'), descripcion, precio, unidades FROM TablaCamiones WHERE referencia = ?", reference).Scan(
    // err = db.QueryRow("SELECT * FROM TablaCamiones WHERE referencia = '"+reference+"'").Scan(
        &product.Reference,
        &product.Name,
        &product.ImagePath,
        &product.ShortDescription,
        &product.ProductInfo.Description,
        &price,
        &product.ProductInfo.AvailableAmount,
    )
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // ESTO AÚN NO FUNSIONAAAAA
    valuesText := []string{}

    // Create a string slice using strconv.Itoa.
    // ... Append strings to it.
    for i := range price {
        number := price[i]
        text := strconv.FormatUint(uint64(number), 10)
        valuesText = append(valuesText, text)
    }

    // Join our string slice.
    result := strings.Join(valuesText, "+")

    product.ProductInfo.Price, err = strconv.Atoi(result)

    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    return product
}

// GetReserves devuelve la lista de reservas de un usuario
func (jdbcDataProvider JDBCDataProvider) GetReserves(username string, passwordSha string) []model.Reserve {
    return []model.Reserve{model.Reserve{}}
}

// PostReserve crea una reserva nueva para un usuario
func (jdbcDataProvider JDBCDataProvider) PostReserve(reference string, username string, passwordSha string) {}

// DeleteReserve borra una reserva nueva para un usuario
func (jdbcDataProvider JDBCDataProvider) DeleteReserve(reference string, username string, passwordSha string) {}

// CheckLogin comprueba si el usuario y la contraseña son correctos
func (jdbcDataProvider JDBCDataProvider) CheckLogin(username string, passwordSha string) model.JSONHTTPResponse {
    return model.JSONHTTPResponse{HTTPResponse:model.HTTPResponse{Code:200, Description: "OK", ExtraText: "Login check OK"}}
}