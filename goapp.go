package main

import (
	"fmt"
	"io/ioutil"
)

//Page sirve para estructurar de las páginas
type Page struct {
	Title string
	//Para procesar mejor es mejor un slice de bytes mas optimo
	Body []byte
}

//sirve para guardar una pagina
func (p *Page) save() error {
	filename := "./data/" + p.Title + ".txt"
	//con esta funcion se guarda, permisos de tipo linux 0600 solo puede escribir el owner de la pagina
	//nombre del archivo con su ubicacion el nombre y luego el propietario de la pagina solo un permiso, solo ler y escribir return retorna el error
	err := ioutil.WriteFile(filename, p.Body, 0600)
	return err
}

//Lee el archivo
func loadPage(title string) *Page {
	filename := "./data/" + title + ".txt"
	body, _ := ioutil.ReadFile(filename)
	page := &Page{Title: title, Body: body}
	return page
}

func main() {
	//sirve para ver probar si sirve para guardar el archivo
	//page := &Page{Title: "primer", Body: []byte("Nuestra primera pgina")}
	//page.save()
	//ahora se probara si sirve la carga
	page := loadPage("primer")
	//string( sirve para convertir en en string )
	fmt.Println(page.Title, string(page.Body))
}
