package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
    "fmt"
)

var addr = flag.String("addr", ":8000", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {

    fmt.Println("******************************************************************************")
    fmt.Println("---------------------------------Gerador De Qr Code-----------------------------")
    fmt.Println("******************************************************************************")
    fmt.Println("")
    fmt.Println(" Abra seu navegador em http://localhost:8000 , para usar seu gerador de qr code.")
    fmt.Println("")


    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("input_formulario_qr_code"))
}

const templateStr = `
<html>
<head>

<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-Zenh87qX5JnK2Jl0vWa8Ck2rdkQ2Bzep5IDxbcnCeuOxjzrPF/et3URy9Bv1WTRi" crossorigin="anonymous">
<title>Gerador de qr code</title>
</head>



<body style="background-image: linear-gradient(white, blue);">
<div class="d-flex  justify-content-center rounded-3">
<div class="card  d-flex " style="width: 22rem; margin:5em;">
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
  <div class="card-body">
    {{end}}
    <h5 class="fs-6">{{.}}</h5>
    <form class="d-flex flex-column  justify-content-center" action="/" name=f method="POST">  

    <input class="mb-4"  name=input_formulario_qr_code placeholder="Digite a informação do qr code" value="" title="Text to QR Encode">
   
    <input class="btn btn-primary "type=submit value="Gerar QR" name=qr>
</form>
</div>    
  </div>
</div>



</body>
</html>
`