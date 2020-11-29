package dataprovider

// GetDataProvider devuelve el DataProvider seleccionado
func GetDataProvider() DataProvider {
    return JDBCDataProvider{}
}