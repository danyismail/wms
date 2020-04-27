package models

type User struct {
    Id          int     `json:"id"`
    Username    string  `json:"username"`
    Email       string  `json:"email"`
    Password    string  `json:"password"`
    Role        int     `json:"role"`
}

type Unit struct {
  Id          int     `json:"id"`
  Code        string  `json:"code"`
  Name        string  `json:"name"`
}

type Incoming struct {
  Id            int     `json:"id"`
  TransaksiId   string  `json:"transaksi_id"`
  Tanggal       string  `json:"tanggal"`
  Lokasi        string  `json:"lokasi"`
  KodeBarang    string  `json:"kodebarang"`
  NamaBarang    string  `json:"namabarang"`
  Satuan        string  `json:"satuan"`
  Jumlah        string  `json:"jumlah"`
}

type Outcoming struct {
  Id            int     `json:"id"`
  TransaksiId           string   `json:"transaksi_id"`
  TanggalMasuk  string  `json:"tanggalmasuk"`
  TanggalKeluar string  `json:"tanggalkeluar"`
  Lokasi        string  `json:"lokasi"`
  KodeBarang    string  `json:"kodebarang"`
  NamaBarang    string  `json:"namabarang"`
  Satuan        string  `json:"satuan"`
  Jumlah        string  `json:"jumlah"`
}

