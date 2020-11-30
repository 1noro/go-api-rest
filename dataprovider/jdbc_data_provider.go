package dataprovider

import (
    "log"
    "time"
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
func (jdbcDataProvider JDBCDataProvider) GetProducts() ([]model.Product, int) {
    var products []model.Product
    httpState := 200
    // Open up our database connection.
    // db, err := sql.Open("mysql", "developer:pass@tcp(gitlab.afundacionfp.com:3306)/mysql")
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
        httpState = 500
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT referencia, nombre, urlImagen, CONCAT(SUBSTRING(descripcion, 1, 117), '...') FROM TablaCamiones")
    if err != nil {
        log.Print(err.Error())
        httpState = 500
    } else {
        for results.Next() {
            var product model.Product
            // for each row, scan the result into our tag composite object
            err = results.Scan(&product.Reference, &product.Name, &product.ImagePath, &product.ShortDescription)
            if err != nil {
                log.Print(err.Error())
                httpState = 500
            } else {
                products = append(products, product)
            }
        }
    }
    return products, httpState
}

// GetFullProduct devuelve el detalle de un producto
func (jdbcDataProvider JDBCDataProvider) GetFullProduct(reference string) (model.Product, int) {
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    if err != nil {log.Print(err.Error())}
    defer db.Close()

    var product model.Product

    sql := "SELECT referencia, nombre, urlImagen, CONCAT(SUBSTRING(descripcion, 1, 117), '...'), descripcion, precio, unidades " + 
                "FROM TablaCamiones WHERE referencia = ?"
    err = db.QueryRow(sql, reference).Scan(
    // err = db.QueryRow("SELECT * FROM TablaCamiones WHERE referencia = '"+reference+"'").Scan(
        &product.Reference,
        &product.Name,
        &product.ImagePath,
        &product.ShortDescription,
        &product.ProductInfo.Description,
        &product.ProductInfo.Price,
        &product.ProductInfo.AvailableAmount,
    )

    if err != nil {panic(err.Error())}

    return product, 200
}

func makeTimestamp(enterTime time.Time) int64 {
    return enterTime.UnixNano() / int64(time.Millisecond)
}

// GetReserves devuelve la lista de reservas de un usuario
func (jdbcDataProvider JDBCDataProvider) GetReserves(username string, passwordSha string) ([]model.Reserve, int) {
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    if err != nil {log.Print(err.Error())}
    defer db.Close()
    var user model.User
    sql := "SELECT usuario, contrasenaSha1, salt FROM TablaClientes WHERE usuario = ?"
    err = db.QueryRow(sql, username).Scan(&user.Username, &user.ConcatenatedPasswordSha, &user.Salt)
    if err != nil {panic(err.Error())}
    if user.CheckPassword(passwordSha) {
        var reserves []model.Reserve
        sql = "SELECT c.referencia, c.nombre, c.urlImagen, c.descripcion, c.precio, c.unidades, UNIX_TIMESTAMP(r.fecha) " +
                    "FROM TablaReservas AS r, TablaCamiones AS c " +
                    "WHERE r.idCliente = ? AND r.refCamion = c.referencia"
        results, err := db.Query(sql, username)
        if err != nil {panic(err.Error())}
        for results.Next() {
            var reserve model.Reserve
            err = results.Scan(
                &reserve.Product.Reference,
                &reserve.Product.Name,
                &reserve.Product.ImagePath,
                &reserve.Product.ProductInfo.Description,
                &reserve.Product.ProductInfo.Price,
                &reserve.Product.ProductInfo.AvailableAmount,
                &reserve.ReserveDate,
            )
            if err != nil {panic(err.Error())}
            reserves = append(reserves, reserve)
        }
        return reserves, 200
    }
    return []model.Reserve{model.Reserve{}}, 200
}

// PostReserve crea una reserva nueva para un usuario
func (jdbcDataProvider JDBCDataProvider) PostReserve(reference string, username string, passwordSha string) int {
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    if err != nil {log.Print(err.Error())}
    defer db.Close()
    var user model.User
    sql := "SELECT usuario, contrasenaSha1, salt FROM TablaClientes WHERE usuario = ?"
    err = db.QueryRow(sql, username).Scan(&user.Username, &user.ConcatenatedPasswordSha, &user.Salt)
    if err != nil {panic(err.Error())}
    if user.CheckPassword(passwordSha) {
        sql = "INSERT INTO TablaReservas(idCliente, refCamion) VALUES(?, ?)"
        insert, err := db.Query(sql, username, reference)
        if err != nil {panic(err.Error())}
        defer insert.Close()
    }
    return 200
}

// DeleteReserve borra una reserva nueva para un usuario
func (jdbcDataProvider JDBCDataProvider) DeleteReserve(reference string, username string, passwordSha string) int {
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    if err != nil {log.Print(err.Error())}
    defer db.Close()
    var user model.User
    sql := "SELECT usuario, contrasenaSha1, salt FROM TablaClientes WHERE usuario = ?"
    err = db.QueryRow(sql, username).Scan(&user.Username, &user.ConcatenatedPasswordSha, &user.Salt)
    if err != nil {panic(err.Error())}
    if user.CheckPassword(passwordSha) {
        sql = "DELETE FROM TablaReservas WHERE (idCliente = ?) AND (refCamion = ?)"
        delete, err := db.Query(sql, username, reference)
        if err != nil {panic(err.Error())}
        defer delete.Close()
    }
    return 200
}

// CheckLogin comprueba si el usuario y la contraseña son correctos
func (jdbcDataProvider JDBCDataProvider) CheckLogin(username string, passwordSha string) (model.JSONHTTPResponse, int) {
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    if err != nil {log.Print(err.Error())}
    defer db.Close()
    var user model.User
    sql := "SELECT usuario, contrasenaSha1, salt FROM TablaClientes WHERE usuario = ?"
    err = db.QueryRow(sql, username).Scan(&user.Username, &user.ConcatenatedPasswordSha, &user.Salt)
    if err != nil {panic(err.Error())}
    if user.CheckPassword(passwordSha) {
        return model.JSONHTTPResponse{HTTPResponse:model.HTTPResponse{Code:200, Description: "OK", ExtraText: "Login check OK"}}, 200
    }
    return model.JSONHTTPResponse{HTTPResponse:model.HTTPResponse{Code:401, Description: "Unauthorized", ExtraText: "Login check FAILED"}}, 200
}