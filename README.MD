<h1>Mars Rover API</h1>
<h5>Uses a caching system to produce the last ten days of rover images with a limit of three images per day.</h5>
<hr>

<h2>Languages And Tools</h2>

* Golang
* Postgres

<h2>Running The Application</h2>

```sh 
git clone https://github.com/NynjaOrca/mars-rover-api.git
```

```sh
go run mars-rover/cmd/svr/main.go
```

<h2>Output</h2>

```json
[
 {
  "EarthDate": "2021-08-04",
  "Images": []
 },
 {
  "EarthDate": "2021-08-03",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03197/opgs/edr/ncam/NLB_681314683EDR_F0900772NCAM00299M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03197/opgs/edr/ncam/NLB_681314652EDR_F0900772NCAM00299M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03197/opgs/edr/ncam/NLB_681314621EDR_F0900772NCAM00299M_.JPG"
  ]
 },
 {
  "EarthDate": "2021-08-02",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03196/opgs/edr/ncam/NLB_681214112EDR_F0900460CCAM15905M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03196/opgs/edr/ncam/NRB_681214112EDR_F0900460CCAM15905M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03196/opgs/edr/ncam/NRB_681233265EDR_S0900460NCAM00536M_.JPG"
  ]
 },
 {
  "EarthDate": "2021-08-01",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03195/opgs/edr/ncam/NLB_681137172EDR_F0900460NCAM00312M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03195/opgs/edr/ncam/NLB_681137148EDR_F0900460NCAM00312M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03195/opgs/edr/ncam/NLB_681137115EDR_F0900460NCAM00312M_.JPG"
  ]
 },
 {
  "EarthDate": "2021-07-31",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03194/opgs/edr/ncam/NRB_681053000EDR_F0900232NCAM00200M_.JPG"
  ]
 },
 {
  "EarthDate": "2021-07-30",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03193/opgs/edr/ncam/NLB_680948088EDR_F0900232CCAM15905M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03193/opgs/edr/ncam/NRB_680954916EDR_F0900232NCAM00599M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03193/opgs/edr/ncam/NRB_680954903EDR_F0900232NCAM00599M_.JPG"
  ]
 },
 {
  "EarthDate": "2021-07-29",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03192/opgs/edr/ncam/NLB_680874801EDR_F0900232NCAM00354M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03192/opgs/edr/ncam/NLB_680874761EDR_F0900232NCAM00354M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03192/opgs/edr/ncam/NLB_680874728EDR_F0900232NCAM00354M_.JPG"
  ]
 },
 {
  "EarthDate": "2021-07-28",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03191/opgs/edr/ncam/NLB_680769824EDR_F0900000CCAM15905M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03191/opgs/edr/ncam/NRB_680769824EDR_F0900000CCAM15905M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03191/opgs/edr/ncam/NRB_680769165EDR_S0900000NCAM00594M_.JPG"
  ]
 },
 {
  "EarthDate": "2021-07-27",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03190/opgs/edr/ncam/NLB_680694657EDR_F0900000NCAM00312M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03190/opgs/edr/ncam/NLB_680694625EDR_F0900000NCAM00312M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03190/opgs/edr/ncam/NLB_680694593EDR_F0900000NCAM00312M_.JPG"
  ]
 },
 {
  "EarthDate": "2021-07-26",
  "Images": [
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03189/opgs/edr/ncam/NLB_680591584EDR_F0892638CCAM15905M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03189/opgs/edr/ncam/NRB_680591584EDR_F0892638CCAM15905M_.JPG",
   "https://mars.nasa.gov/msl-raw-images/proj/msl/redops/ods/surface/sol/03189/opgs/edr/ncam/NRB_680602531EDR_D0892638NCAM00582M_.JPG"
  ]
 }
]
```