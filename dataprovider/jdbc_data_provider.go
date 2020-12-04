package dataprovider

import (
    "strconv"
    "log"
    "time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // go get -u github.com/go-sql-driver/mysql
    "../model"
)

// makeTimestamp: funcion de utilidad genérica (debería moverse en un futuro)
func makeTimestamp(enterTime time.Time) int64 {
    return enterTime.UnixNano() / int64(time.Millisecond)
}

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
    defer db.Close()
    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error()) // Error al conectar
        httpState = 500
    } else {
        // Execute the query
        results, err := db.Query("SELECT referencia, nombre, urlImagen, CONCAT(SUBSTRING(descripcion, 1, 117), '...') FROM TablaCamiones")
        if err != nil {
            log.Print(err.Error()) // Error de la query
            httpState = 500
        } else {
            for results.Next() {
                var product model.Product
                // for each row, scan the result into our tag composite object
                err = results.Scan(&product.Reference, &product.Name, &product.ImagePath, &product.ShortDescription)
                if err != nil {
                    log.Print(err.Error()) // Error de la obtención de resultados
                    httpState = 500
                } else {
                    products = append(products, product)
                }
            }
        }
    }
    return products, httpState
}

// GetFullProduct devuelve el detalle de un producto
func (jdbcDataProvider JDBCDataProvider) GetFullProduct(reference string) (model.Product, int) {
    var product model.Product
    httpState := 200
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    defer db.Close()
    if err != nil {
        log.Print(err.Error())
        httpState = 500
    } else {
        sql := "SELECT referencia, nombre, urlImagen, CONCAT(SUBSTRING(descripcion, 1, 117), '...'), descripcion, precio, unidades " + 
                    "FROM TablaCamiones WHERE referencia = ?"
        err = db.QueryRow(sql, reference).Scan(
            &product.Reference,
            &product.Name,
            &product.ImagePath,
            &product.ShortDescription,
            &product.ProductInfo.Description,
            &product.ProductInfo.Price,
            &product.ProductInfo.AvailableAmount,
        )
        if err != nil {
            log.Print(err.Error())
            if err.Error() == "sql: no rows in result set" {
                httpState = 404
            } else {
                httpState = 500
            }
        }
    }
    return product, httpState
}

// GetReserves devuelve la lista de reservas de un usuario
func (jdbcDataProvider JDBCDataProvider) GetReserves(username string, passwordSha string) ([]model.Reserve, int) {
    var reserves []model.Reserve
    httpState := 200
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    defer db.Close()
    if err != nil {
        log.Print(err.Error()) // falla la conexion
        httpState = 500
    } else {
        var user model.User
        sql := "SELECT usuario, contrasenaSha1, salt FROM TablaClientes WHERE usuario = ?"
        err = db.QueryRow(sql, username).Scan(&user.Username, &user.ConcatenatedPasswordSha, &user.Salt)
        if err != nil {
            log.Print(err.Error()) // falla la query
            if err.Error() == "sql: no rows in result set" {
                httpState = 404
            } else {
                httpState = 500
            }
        } else {
            if user.CheckPassword(passwordSha) {
                sql = "SELECT c.referencia, c.nombre, c.urlImagen, c.descripcion, c.precio, c.unidades, UNIX_TIMESTAMP(r.fecha) " +
                            "FROM TablaReservas AS r, TablaCamiones AS c " +
                            "WHERE r.idCliente = ? AND r.refCamion = c.referencia"
                results, err := db.Query(sql, username)
                if err != nil {
                    log.Print(err.Error()) // falla la query
                    httpState = 500
                } else {
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
                        if err != nil {
                            log.Print(err.Error()) // falla la obtención de resultados
                            httpState = 500
                        } else {
                            reserves = append(reserves, reserve)
                        }
                    }
                }
            } else {
                httpState = 401 // contraseña incorrecta
            }
        }
    }
    return reserves, httpState
}

// PostReserve crea una reserva nueva para un usuario
func (jdbcDataProvider JDBCDataProvider) PostReserve(reference string, username string, passwordSha string) int {
    httpState := 201
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    defer db.Close()
    if err != nil {
        log.Print(err.Error()) // falla la conexion
        httpState = 500
    } else {
        var user model.User
        sql := "SELECT usuario, contrasenaSha1, salt FROM TablaClientes WHERE usuario = ?"
        err = db.QueryRow(sql, username).Scan(&user.Username, &user.ConcatenatedPasswordSha, &user.Salt)
        if err != nil {
            log.Print(err.Error()) // falla la query
            if err.Error() == "sql: no rows in result set" {
                httpState = 404
            } else {
                httpState = 500
            }
        } else {
            if user.CheckPassword(passwordSha) {
                sql = "INSERT INTO TablaReservas(idCliente, refCamion) VALUES(?, ?)"
                result, err := db.Exec(sql, username, reference)
                lastInsertID, _ := result.LastInsertId()
                rowsAffected, _ := result.RowsAffected()
                log.Print("LastInsertId: " + strconv.Itoa(int(lastInsertID)))
                log.Print("RowsAffected: " + strconv.Itoa(int(rowsAffected)))
                if err != nil {
                    log.Print(err.Error()) // falla la query
                    httpState = 500
                }
            } else {
                httpState = 401 // contraseña incorrecta
            }
        }
    }
    return httpState
}

// DeleteReserve borra una reserva nueva para un usuario
func (jdbcDataProvider JDBCDataProvider) DeleteReserve(reference string, username string, passwordSha string) int {
    httpState := 200
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    defer db.Close()
    if err != nil {
        log.Print(err.Error()) // falla la conexion
        httpState = 500
    } else {
        var user model.User
        sql := "SELECT usuario, contrasenaSha1, salt FROM TablaClientes WHERE usuario = ?"
        err = db.QueryRow(sql, username).Scan(&user.Username, &user.ConcatenatedPasswordSha, &user.Salt)
        if err != nil {
            log.Print(err.Error()) // falla la query
            if err.Error() == "sql: no rows in result set" {
                httpState = 404
            } else {
                httpState = 500
            }
        } else {
            if user.CheckPassword(passwordSha) {
                sql = "DELETE FROM TablaReservas WHERE (idCliente = ?) AND (refCamion = ?)"
                result, err := db.Exec(sql, username, reference)
                if err != nil {
                    log.Print(err.Error()) // falla la query
                    httpState = 500
                } else {
                    rowsAffected, _ := result.RowsAffected()
                    if rowsAffected == 0 {
                        httpState = 404 // No se ha borrado nada
                    }
                }
            } else {
                httpState = 401 // contraseña incorrecta
            }
        }
    }
    return httpState
}

// CheckLogin comprueba si el usuario y la contraseña son correctos
func (jdbcDataProvider JDBCDataProvider) CheckLogin(username string, passwordSha string) int {
    httpState := 200
    db, err := sql.Open(dbServer, dbUsername+":"+dbPass+"@"+dbProtocol+"("+dbURL+":"+dbPort+")/"+dbName)
    defer db.Close()
    if err != nil {
        log.Print(err.Error())
        httpState =  500
    } else {
        var user model.User
        sql := "SELECT usuario, contrasenaSha1, salt FROM TablaClientes WHERE usuario = ?"
        err = db.QueryRow(sql, username).Scan(&user.Username, &user.ConcatenatedPasswordSha, &user.Salt)
        if err != nil {
            log.Print(err.Error())
            if err.Error() == "sql: no rows in result set" {
                httpState = 404
            } else {
                httpState = 500
            }
        } else {
            if !user.CheckPassword(passwordSha) {
                httpState =  401
            }
        }
    }
    return httpState
}