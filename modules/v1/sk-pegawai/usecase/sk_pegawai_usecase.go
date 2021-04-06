package usecase

//  func HandleGetSKPegawai(a app.App) echo.HandlerFunc {
// h := func(c echo.Context) error {
// IDPegawai := c.QueryParam("id_pegawai")
// skPegawai, err := repo.GetSKPegawai(a, IDPegawai)

// if err != nil {
// fmt.Printf("[ERROR] %s\n", err.Error())
// return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
// }
// for i, sk := range skPegawai {
// switch sk.IDJenisSK {
// case "1":
// skPengangkatan, err := repo.GetSKPengangkatan(a, sk.UUID)
// if err != nil {
// fmt.Printf("[ERROR] %s\n", err.Error())
// return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
// }
// if skPengangkatan != skPengangkatanModel.EmptySKPengangkatanPegawai() {
//  fmt.Printf("\nsk pengangkatan ada kak \n")
// skPegawai[i].SetSKPengangkatan(skPengangkatan)
// sk.SetSKPengangkatan(skPengangkatan)
// fmt.Printf("\nDEBUG SK : %+v \n", sk)
// }
// }
// }
// fmt.Printf("\nDEBUG list SK setelah for: %+v \n", skPegawai)

// return c.JSON(http.StatusOK, model.SKPegawaiResponse{
// Data: skPegawai,
// })
// }
// return echo.HandlerFunc(h)
//  }
