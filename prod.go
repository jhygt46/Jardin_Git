package main

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/smtp"
	"strconv"
	"strings"

	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/nfnt/resize"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/fasthttp/router"
	_ "github.com/go-sql-driver/mysql"

	"github.com/valyala/fasthttp"
)

type TemplateConf struct {
	Titulo          string          `json:"Titulo"`
	SubTitulo       string          `json:"SubTitulo"`
	SubTitulo2      string          `json:"SubTitulo2"`
	TituloLista     string          `json:"SubTitulo"`
	FormId          int             `json:"FormId"`
	FormId2         int             `json:"FormId2"`
	Num             int             `json:"Num"`
	Tipo            int             `json:"Tipo"`
	FormAccion      string          `json:"FormAccion"`
	PageMod         string          `json:"PageMod"`
	DelAccion       string          `json:"DelAccion"`
	DelObj          string          `json:"DelObj"`
	Resize          Resize          `json:"Resize"`
	UsuarioTipo     int             `json:"UsuarioTipo"`
	Usuario         Usuario         `json:"Usuario"`
	Curso           Curso           `json:"Curso"`
	Cursos          []Lista         `json:"Cursos"`
	Agenda          Agenda          `json:"Agenda"`
	Prestamos       []Prestamo      `json:"Prestamos"`
	Lista           []Lista         `json:"Lista"`
	Padres          []Padres        `json:"Padres"`
	Padre           Padres          `json:"Padre"`
	UsuarioComplete UsuarioComplete `json:"UsuarioComplete"`
	CursosCat       CursosCat       `json:"CursosCat"`
	CursosCat2      []CursosCat     `json:"CursosCat2"`
	CursosItems     CursosItems     `json:"CursosItems"`
	ListaUsers      []ListaUsers    `json:"ListaUsers"`
}
type Resize struct {
	Ok    bool   `json:"Ok"`
	Image string `json:"Image"`
	Path  string `json:"Path"`
	Block bool   `json:"Block"`
	Dx    int    `json:"Dx"`
	Dy    int    `json:"Dy"`
}
type Lista struct {
	Id        int    `json:"Id"`
	Num       int    `json:"Num"`
	Nombre    string `json:"Nombre"`
	Apellido1 string `json:"Apellido1"`
	Apellido2 string `json:"Apellido2"`
	Tipo      int    `json:"Tipo"`
}
type Config struct {
	Tiempo time.Duration `json:"Tiempo"`
}
type MyHandler struct {
	Conf         Config              `json:"Conf"`
	Passwords    Passwords           `json:"Passwords"`
	CursosOnline *CursosOnlineStruct `json:"CursosOnline"`
	ListaDel     []DelImages         `json:"Lista"`
	VideoPath    string              `json:"VideoPath"`
	Os           string              `json:"Os"`
}
type CursosOnlineStruct struct {
	CursosCat []CursosCat `json:"CursosCat"`
}
type CursosCat struct {
	Id          int           `json:"Id"`
	Nombre      string        `json:"Nombre"`
	Url         string        `json:"Url"`
	Nino        int           `json:"Nino"`
	Visible     int           `json:"Visible"`
	CursosItems []CursosItems `json:"CursosItems"`
}
type CursosItems struct {
	Id          int    `json:"Id"`
	Nombre      string `json:"Nombre"`
	Url         string `json:"Url"`
	Image       string `json:"Image"`
	Tipo        int    `json:"Tipo"`
	Dx          int    `json:"Dx"`
	Dy          int    `json:"Dy"`
	ListaImagen string `json:"ListaImagen"`
	Urlext      string `json:"Urlext"`
	Id_cat      int    `json:"Id_cat"`
}
type ListaImagen struct {
	Nom string `json:"Nom"`
}
type Passwords struct {
	PassDb    string `json:"PassDb"`
	PassEmail string `json:"PassEmail"`
}
type Reestablecer struct {
	Op   int    `json:"Op"`
	Id   int    `json:"Id"`
	Code string `json:"Code"`
}
type Indexs struct {
	Login        int                `json:"Login"`
	Reestablecer Reestablecer       `json:"Reestablecer"`
	Function     string             `json:"Nombre"`
	Page         string             `json:"Page"`
	User         IndexUser          `json:"User"`
	Modulos      []Modulo           `json:"Modulos"`
	Register     bool               `json:"Register"`
	Permisos     IndexPermisos      `json:"Permisos"`
	Libro        Libro              `json:"Libro"`
	Agenda       Agenda             `json:"Agenda"`
	Alumnos      string             `json:"Alumnos"`
	Prestamos    []Prestamo         `json:"Prestamos"`
	CursosOnline CursosOnlineStruct `json:"CursosOnline"`
	UrlOnline    UrlOnline          `json:"UrlOnline"`
}
type UrlOnline struct {
	Cat         CursosCat   `json:"Cat"`
	Cats        []CursosCat `json:"Cats"`
	Item        CursosItems `json:"Item"`
	ListaCursos string      `json:"ListaCursos"`
	ListaItems  string      `json:"CursosItems"`
}
type IndexPermisos struct {
	Admin     bool `json:"Admin"`
	Educadora bool `json:"Educadora"`
	Apoderado bool `json:"Apoderado"`
}
type IndexUser struct {
	Id_usr     int    `json:"Id_usr"`
	Tipo       int    `json:"Tipo"`
	CantAgenda int    `json:"CantAgenda"`
	Nombre     string `json:"Nombre"`
	Correo     string `json:"Correo"`
	Telefono   string `json:"Telefono"`
	Telefono2  string `json:"Telefono2"`
}
type Modulo struct {
	Id     int    `json:"Id"`
	Titulo string `json:"Titulo"`
	Text   string `json:"Text"`
	Icon   string `json:"Icon"`
	Url    string `json:"Url"`
}
type Usuario struct {
	Id_usr          int    `json:"Id_usr"`
	Nombre          string `json:"Nombre"`
	Correo          string `json:"Correo"`
	Telefono        string `json:"Telefono"`
	Nmatricula      string `json:"Nmatricula"`
	Rut             string `json:"Rut"`
	Apellido1       string `json:"Apellido1"`
	Apellido2       string `json:"Apellido2"`
	Genero          int    `json:"Genero"`
	Reglamento      int    `json:"Reglamento"`
	FechaNacimiento string `json:"FechaNacimiento"`
	FechaMatricula  string `json:"FechaMatricula"`
	FechaIngreso    string `json:"FechaIngreso"`
	Direccion       string `json:"Direccion"`
	FechaRetiro     string `json:"FechaRetiro"`
	MotivoRetiro    int    `json:"MotivoRetiro"`
	Observaciones   string `json:"Observaciones"`
	Id_cur          int    `json:"Id_cur"`
}
type UsuarioComplete struct {
	Id_usr     int        `json:"Id_usr"`
	Nombre     string     `json:"Nombre"`
	Apellido1  string     `json:"Apellido1"`
	Apellido2  string     `json:"Apellido2"`
	Tipo       int        `json:"Tipo"`
	Cursos     []Curso    `json:"Cursos"`
	Prestamos  []Prestamo `json:"Prestamos"`
	Parentesco []Padres   `json:"Parentesco"`
}
type Padres struct {
	Id_usr    int    `json:"Id_usr"`
	Nombre    string `json:"Nombre"`
	Telefono  string `json:"Telefono"`
	Email     string `json:"Email"`
	Tipo      int    `json:"Tipo"`      // 1 MADRE - 2 PADRE
	Apoderado int    `json:"Apoderado"` // 1 MADRE - 2 PADRE
}
type Agenda struct {
	Fecha       string       `json:"Fecha"`
	FechaStr    string       `json:"FechaStr"`
	FechaPrev   string       `json:"FechaPrev"`
	FechaNext   string       `json:"FechaNext"`
	FechaIsNext int          `json:"FechaIsNext"`
	FechaIsPrev int          `json:"FechaIsPrev"`
	Users       []AgendaUser `json:"Users"`
	AgendaCurso AgendaCurso  `json:"AgendaCurso"`
}
type AgendaUser struct {
	Data           bool   `json:"Data"`
	Nombre         string `json:"Nombre"`
	Ali1           int    `json:"Ali1"`
	Ali2           int    `json:"Ali2"`
	Ali3           int    `json:"Ali3"`
	Dep1           int    `json:"Dep1"`
	Dep2           int    `json:"Dep2"`
	Comentario     string `json:"Comentario"`
	Ausente        int    `json:"Ausente"`
	UltimaAct      string `json:"UltimaAct"`
	UltimaActAlert string `json:"UltimaActAlert"`
	ShowComentario bool   `json:"ShowComentario"`
}
type AgendaCurso struct {
	Fecha  string        `json:"Fecha"`
	Cursos map[int]Curso `json:"Cursos"`
}
type Curso struct {
	Id_cur int                `json:"Id_cur"`
	Nombre string             `json:"Nombre"`
	Orden  int                `json:"Orden"`
	Users  map[int]AgendaUser `json:"Users"`
}
type Libro struct {
	Found         bool       `json:"Found"`
	Id_Lib        int        `json:"Id_Lib"`
	Nombre        string     `json:"Nombre"`
	Code          string     `json:"Code"`
	Id_Pre        int        `json:"Id_Pre"`
	Prestado      bool       `json:"Prestado"`
	NombreAlu     string     `json:"NombreAlu"`
	FechaEntrega  string     `json:"FechaEntrega"`
	FechaPrestamo string     `json:"FechaPrestamo"`
	Prestamos     []Prestamo `json:"Prestamos"`
}
type Prestamo struct {
	Id               int    `json:"Id"`
	Nombre           string `json:"Nombre"`
	Fecha_Prestamos  string `json:"Fecha_Prestamos"`
	Fecha_Devolucion string `json:"Fecha_Devolucion"`
	Nombre_Alu       string `json:"Nombre_Alu"`
	Duration         string `json:"Duration"`
}
type Response struct {
	Op     uint8  `json:"Op"`
	Msg    string `json:"Msg"`
	Reload int    `json:"Reload"`
	Page   string `json:"Page"`
	Tipo   string `json:"Tipo"`
	Titulo string `json:"Titulo"`
	Texto  string `json:"Texto"`
	Agenda Agenda `json:"Agenda"`
}
type Html_Qr struct {
	List []int `json:"List"`
}
type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}
type ListaImagenes struct {
	I     int     `json:"I"`
	Ok    bool    `json:"Ok"`
	Nom   string  `json:"Nom"`
	Dx    int     `json:"Dx"`
	Dy    int     `json:"Dy"`
	Ratio float64 `json:"Ratio"`
}
type DelImages struct {
	Time  time.Time `json:"Time"`
	Image string    `json:"Image"`
}
type ListaUsers struct {
	Id_usr       int    `json:"Id_usr"`
	Num          int    `json:"Num"`
	Nombre       string `json:"Nombre"`
	Apellido1    string `json:"Apellido1"`
	Apellido2    string `json:"Apellido2"`
	Rut          string `json:"Rut"`
	MamaNom      string `json:"MamaNom"`
	MamaTelefono string `json:"MamaTelefono"`
	PapaNom      string `json:"PapaNom"`
	PapaTelefono string `json:"PapaTelefono"`
	NMatricula   string `json:"NMatricula"`
	FechaNac     string `json:"FechaNac"`
	Direccion    string `json:"Direccion"`
	Genero       int    `json:"Genero"`
}

var (
	imgHandler        fasthttp.RequestHandler
	imgPreviewHandler fasthttp.RequestHandler
	imgCuentosHandler fasthttp.RequestHandler
	cssHandler        fasthttp.RequestHandler
	jsHandler         fasthttp.RequestHandler
	port              string
	favicon           *[]byte
)
var pass = &MyHandler{Conf: Config{}, CursosOnline: &CursosOnlineStruct{}, ListaDel: make([]DelImages, 0)}

func main() {

	log.SetFlags(0)
	log.SetOutput(ioutil.Discard)

	//SendEmail("diego.gomez.bezmalinovic@gmail.com", "prueba", "mensaje de prueba")

	if runtime.GOOS == "windows" {

		imgHandler = fasthttp.FSHandler("C:/Go/Jardin_Git/img", 1)
		imgPreviewHandler = fasthttp.FSHandler("C:/Go/Jardin_Git/img/preview", 1)
		imgCuentosHandler = fasthttp.FSHandler("C:/Go/Jardin_Git/img/cuentos", 1)
		cssHandler = fasthttp.FSHandler("C:/Go/Jardin_Git/css", 1)
		jsHandler = fasthttp.FSHandler("C:/Go/Jardin_Git/js", 1)
		pass.VideoPath = "C:/Go/Jardin_Git/videos"
		pass.Os = "windows"
		port = ":81"

		passwords, err := os.ReadFile("../password_valleencantado_local.json")
		if err == nil {
			if err := json.Unmarshal(passwords, &pass.Passwords); err == nil {
				//fmt.Println(pass.Passwords)
			}
		}

	} else {

		imgHandler = fasthttp.FSHandler("/var/Jardin_Git/img", 1)
		imgPreviewHandler = fasthttp.FSHandler("/var/Jardin_Git/img/preview", 1)
		imgCuentosHandler = fasthttp.FSHandler("/var/Jardin_Git/img/cuentos", 1)
		cssHandler = fasthttp.FSHandler("/var/Jardin_Git/css", 1)
		jsHandler = fasthttp.FSHandler("/var/Jardin_Git/js", 1)
		pass.VideoPath = "/var/Jardin_Git/videos"
		pass.Os = "linux"
		port = ":80"

		passwords, err := os.ReadFile("../password_valleencantado.json")
		if err == nil {
			if err := json.Unmarshal(passwords, &pass.Passwords); err == nil {
				//fmt.Println(pass.Passwords)
			}
		}

	}

	if err := SetCurso(pass); err == nil {
		//fmt.Println("Cursos Online Ok")
	}

	fav_bytes, err := os.ReadFile("./img/favicon.png")
	if err == nil {
		favicon = &fav_bytes
	}

	con := context.Background()
	con, cancel := context.WithCancel(con)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		for {
			select {
			case s := <-signalChan:
				switch s {
				case syscall.SIGHUP:
					pass.Conf.init()
				case os.Interrupt:
					cancel()
					os.Exit(1)
				}
			case <-con.Done():
				log.Printf("Done.")
				os.Exit(1)
			}
		}
	}()
	go func() {
		r := router.New()
		r.GET("/", Index)
		r.GET("/favicon.ico", Favicon)
		r.GET("/css/{name}", Css)
		r.GET("/js/{name}", Js)
		r.GET("/img/{name}", Img)
		r.GET("/images_preview/{name}", ImgPreview)
		r.GET("/images_cuentos/{name}", ImgCuentos)
		r.GET("/pages/{name}", Pages)
		r.POST("/login", Login)
		r.POST("/recuperar", Recuperar)
		r.POST("/nueva", Nueva)
		r.POST("/save", Save)
		r.POST("/delete", Delete)
		r.POST("/accion", Accion)
		r.GET("/salir", Salir)
		r.GET("/admin", Admin)
		r.GET("/reestablecer", Index)
		r.GET("/libros", LibroInicio)
		r.GET("/libro/{name}", LibroPage)
		r.GET("/videos/{name}", VideoPage)

		r.GET("/cursos_online", CursosOnline)
		r.GET("/cursos_online/{cat}", CursosOnline)
		r.GET("/cursos_online/{cat}/{item}", CursosOnline)

		r.GET("/agenda", AgendaPage)
		r.GET("/qr/{name}", Qr)
		r.GET("/qrhtml/{name}", HtmlQr)

		// ANTES
		if runtime.GOOS == "windows" {
			fasthttp.ListenAndServe(port, r.Handler)
		} else {
			go func() {
				fasthttp.ListenAndServe(":80", redirectHTTP)
			}()
			server := &fasthttp.Server{Handler: r.Handler}
			server.ListenAndServeTLS(":443", "/etc/letsencrypt/live/valleencantado.cl/fullchain.pem", "/etc/letsencrypt/live/valleencantado.cl/privkey.pem")
		}

	}()
	if err := run(con, pass, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
func redirectHTTP(ctx *fasthttp.RequestCtx) {
	ctx.Redirect("https://www.valleencantado.cl", fasthttp.StatusMovedPermanently)
}
func Accion(ctx *fasthttp.RequestCtx) {

	resp := Response{}
	resp.Op = 2
	resp.Msg = "Error Inesperado"

	token := string(ctx.Request.Header.Cookie("cu"))

	db, err := GetMySQLDB()
	defer db.Close()
	if err != nil {
		resp.Msg = "Error Base de Datos"
	}

	switch string(ctx.FormValue("accion")) {
	case "get_agenda":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Apoderado {
			fecha := string(ctx.FormValue("fecha"))
			agenda, found := GetHijosAgenda(db, perms.User.Id_usr, fecha)
			if found {
				resp.Op = 1
				resp.Msg = ""
				resp.Agenda = agenda
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
	case "mod_datos":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Apoderado {
			nombre := string(ctx.FormValue("nom"))
			mail := string(ctx.FormValue("mail"))
			fono := string(ctx.FormValue("fono"))
			resp.Op, resp.Msg = ChangeUserNom(db, perms.User.Id_usr, nombre)
			if resp.Op == 1 {
				resp.Op, resp.Msg = ChangeUserCorreo(db, perms.User.Id_usr, mail)
				if resp.Op == 1 {
					resp.Op, resp.Msg = ChangeUserTelefono(db, perms.User.Id_usr, fono)
				}
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
	case "set_agenda":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin || perms.Permisos.Educadora {
			id_alu := Read_uint32bytes(ctx.FormValue("id_alu"))
			tipo := Read_uint32bytes(ctx.FormValue("tipo"))
			value := string(ctx.FormValue("value"))
			resp.Op, resp.Msg = IngresarAgenda(db, id_alu, tipo, value)
		} else {
			resp.Msg = "No tiene permisos"
		}
	case "reestablecer":
		ctx.Response.Header.Set("Content-Type", "application/json")
		id_usr := Read_uint32bytes(ctx.FormValue("id_usr"))
		code := string(ctx.FormValue("code"))
		pass1 := ctx.FormValue("pass1")
		pass2 := ctx.FormValue("pass2")
		resp.Op, resp.Msg = ReestablecerPassword(db, id_usr, code, pass1, pass2)
	case "enviar_contacto":
		ctx.Response.Header.Set("Content-Type", "application/json")
		nombre := string(ctx.FormValue("nombre"))
		correo := string(ctx.FormValue("correo"))
		telefono := string(ctx.FormValue("telefono"))
		mensaje := string(ctx.FormValue("mensaje"))
		if SendEmail("valle-encantado@hotmail.com", "Contacto ValleEncantado", fmt.Sprintf("Nombre: %v <br/> Correo: %v <br/> Telefono: %v <br/> Mensaje: %v", nombre, correo, telefono, mensaje)) {
			resp.Op = 1
		} else {
			resp.Msg = "Error al enviar el correo"
		}
	default:
		ctx.Response.Header.Set("Content-Type", "application/json")
		json.NewEncoder(ctx).Encode(resp)
	}
	json.NewEncoder(ctx).Encode(resp)
}
func Pages(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("text/html; charset=utf-8")
	name := ctx.UserValue("name")
	token := string(ctx.Request.Header.Cookie("cu"))

	db, err := GetMySQLDB()
	defer db.Close()
	if err != nil {
		ErrorCheck(err)
	}

	switch name {
	case "crearUsuario":

		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			tipo := Read_uint32bytes(ctx.QueryArgs().Peek("tipo"))
			obj := GetTemplateConf("Crear Usuario", "Formulario", "Complete los campos", "Lista Usuarios", "guardar_usuarios", fmt.Sprintf("/pages/%s", name), "borrar_usuario", "Usuario")

			if tipo == 1 {
				obj.Titulo = "Crear Educadora"
				obj.TituloLista = "Lista Educadoras"
			}
			if tipo == 2 {
				obj.Titulo = "Crear Padre"
				obj.TituloLista = "Lista Padres"
			}
			if tipo == 3 {
				obj.Titulo = "Crear Alumno"
				obj.TituloLista = "Lista Alumnos"
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			if id > 0 {
				aux, found1 := GetUsuario(db, id, tipo)
				if found1 {
					obj.FormId = id
					obj.Usuario = aux
				}
			} else {
				obj.FormId = 0
			}

			aux2, found2 := GetUsuarios(db, tipo)
			if found2 {
				obj.Lista = aux2
			}

			if tipo == 3 {
				aux3, found3 := GetCursos(db)
				if found3 {
					obj.Cursos = aux3
				}
			}

			obj.UsuarioTipo = tipo

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "crearCursos":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			obj := GetTemplateConf("Crear Cursos", "Formulario", "Complete los campos", "Lista Cursos", "guardar_cursos", fmt.Sprintf("/pages/%s", name), "borrar_curso", "Curso")

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			if id > 0 {
				aux, found1 := GetCurso(db, id)
				if found1 {
					obj.FormId = id
					obj.Curso = aux
				}
			} else {
				obj.FormId = 0
			}

			aux2, found2 := GetCursos(db)
			if found2 {
				obj.Lista = aux2
			}

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "verAgenda":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			fecha := string(ctx.QueryArgs().Peek("fecha"))
			obj := GetTemplateConf("Agenda", "Lista de Cursos", "Ultima actualización", "Titulo Lista", "", fmt.Sprintf("/pages/%s", name), "", "")

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			agenda, found := GetAgendaCurso(db, perms.User.Id_usr, fecha)
			if found {
				PrintJson(agenda)
				obj.Agenda = agenda
			}

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "Prestamos":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			obj := GetTemplateConf("Libros", "Prestamos", "Libros prestados", "", "", fmt.Sprintf("/pages/%s", name), "borrar_prestamo", "Prestamo")

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			prestamos, found := GetTodosLibroPrestados(db)
			if found {
				obj.Prestamos = prestamos
			}

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "verPadres":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			id_apo := Read_uint32bytes(ctx.QueryArgs().Peek("id_apo"))
			obj := GetTemplateConf("Relacionar Padres", "Formulario", "Complete los campos", "Listado de Padres", "guardar_parentesco", fmt.Sprintf("/pages/%s", name), "borrar_parentesco1", "Padre")

			if id > 0 {

				obj.FormId = id
				if id_apo > 0 {
					padre, found := GetPadreUser(db, id, id_apo)
					if found {
						obj.Padre = padre
					}
				}

				padres, found := GetPadresUser(db, id)
				if found {
					obj.Padres = padres
				}
				aux, found1 := GetUsuario(db, id, 3)
				if found1 {
					obj.Usuario = aux
				}
				aux2, found2 := GetUsuarios(db, 2)
				if found2 {
					obj.Lista = aux2
				}
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)
			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "verAlumnos":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			obj := GetTemplateConf("Relacionar Alumnos", "Formulario", "Complete los campos", "Listado de Alumnos", "guardar_parentesco2", fmt.Sprintf("/pages/%s", name), "borrar_parentesco2", "Alumno")

			if id > 0 {

				obj.FormId = id
				padres, found := GetAlumnosUser(db, id)
				if found {
					obj.Padres = padres
				}
				aux, found1 := GetUsuario(db, id, 3)
				if found1 {
					obj.Usuario = aux
				}
				aux2, found2 := GetUsuarios(db, 3)
				if found2 {
					obj.Lista = aux2
				}
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)
			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "verEducadoras":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			obj := GetTemplateConf("Relacionar Educadora", "Formulario", "Complete los campos", "Listado de Educadoras", "guardar_edu_curso", fmt.Sprintf("/pages/%s", name), "borrar_cur_edu1", "Educadora")
			obj.FormId = id

			educadoras, found := GetUserCursos(db, id)
			if found {
				obj.Padres = educadoras
			}

			aux, found := GetUsuariosEduAdmin(db)
			if found {
				obj.Lista = aux
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "verCursos":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			obj := GetTemplateConf("Relacionar Curso", "Formulario", "Complete los campos", "Listado de Cursos", "guardar_edu_curso2", fmt.Sprintf("/pages/%s", name), "borrar_cur_edu2", "Curso")
			obj.FormId = id

			educadoras, found := GetCursosUser(db, id)
			if found {
				obj.Padres = educadoras
			}

			aux, found := GetCursos(db)
			if found {
				obj.Lista = aux
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "resizeImage":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			num := Read_uint32bytes(ctx.QueryArgs().Peek("num"))
			tipo := Read_uint32bytes(ctx.QueryArgs().Peek("tipo"))
			obj := GetTemplateConf("Recortar Imagen", "Formulario", "Complete los campos", "Listado de Cursos", "guardar_resize", fmt.Sprintf("/pages/%s", name), "borrar_resize", "Imagen")

			obj.Resize.Ok = false
			item, found := GetCursoOnlineItem(db, id)

			if found {
				obj.FormId = id
				obj.Num = num
				if tipo == 0 {
					obj.Resize.Ok = true
					obj.Resize.Path = "images_preview"
					obj.Resize.Image = item.Image
					obj.Resize.Block = true
					obj.Resize.Dx = 130
					obj.Resize.Dy = 90
				} else {
					lista_imagen := []ListaImagen{}
					if err := json.Unmarshal([]byte(item.ListaImagen), &lista_imagen); err == nil {
						if len(lista_imagen) >= num {
							obj.Resize.Ok = true
							obj.Resize.Path = "images_cuentos"
							obj.Resize.Image = lista_imagen[num-1].Nom
							if item.Dx > 0 && item.Dy > 0 {
								obj.Resize.Block = true
								obj.Resize.Dx = item.Dx
								obj.Resize.Dy = item.Dy
							} else {
								obj.Resize.Block = false
							}
						}
					}
				}
			}

			fmt.Println(obj.Resize)

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "uploadCuentosImage":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			id2 := Read_uint32bytes(ctx.QueryArgs().Peek("id2"))
			obj := GetTemplateConf("Recortar Imagen", "Formulario", "Complete los campos", "Listado de Cursos", "guardar_image_cuentos", fmt.Sprintf("/pages/%s", name), "borrar_cuentos_image", "Imagen")

			item, found := GetCursoOnlineItem(db, id)
			if found {
				obj.FormId = id
				obj.FormId2 = id2

				lista_imagen := []ListaImagen{}
				if err := json.Unmarshal([]byte(item.ListaImagen), &lista_imagen); err == nil {
					for i, x := range lista_imagen {
						obj.Lista = append(obj.Lista, Lista{Id: i, Nombre: x.Nom})
					}
				}
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "verCursosOnline":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			obj := GetTemplateConf("Cursos Online", "Formulario", "Complete los campos", "Lista de Categorias", "guardar_cursos_online", fmt.Sprintf("/pages/%s", name), "borrar_cursos_online", "Curso Online")

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			if id > 0 {
				aux, found1 := GetCursoOnline(db, id)
				if found1 {
					obj.FormId = id
					obj.CursosCat = aux
				}
			} else {
				obj.FormId = 0
			}

			aux2, found2 := GetCursosOnline(db)
			if found2 {
				obj.Lista = aux2
			}

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "verItemCursoOnline":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			id_i := Read_uint32bytes(ctx.QueryArgs().Peek("id_item"))
			obj := GetTemplateConf("Items Cursos Online", "Formulario", "Complete los campos", "Lista de Categorias", "guardar_cursos_online_items", fmt.Sprintf("/pages/%s", name), "borrar_cursos_online_item", "Curso Online Items")

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			if id > 0 {
				aux2, found2 := GetCursoOnline1(db)
				if found2 {
					obj.CursosCat2 = aux2
				}
				aux, found1 := GetCursoOnline2(db, id)
				if found1 {
					obj.FormId2 = id
					obj.CursosCat.CursosItems = aux
				}
				if id_i > 0 {
					aux2, found2 := GetCursoOnlineItem(db, id_i)
					if found2 {
						obj.FormId = id
						obj.CursosItems = aux2
					}
				}
			}
			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "borrarUsuario":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			obj := GetTemplateConf("Borrar", "", "Información a eliminar", "", "", fmt.Sprintf("/pages/%s", name), "borrar_user23", "Usuario")

			obj.FormId = id

			Usuario, found := GetUsuarioComplete(db, id)
			if found {
				if Usuario.Tipo == 3 {
					obj.SubTitulo = "Borrar Alumno"
				}
				if Usuario.Tipo == 2 {
					obj.SubTitulo = "Borrar Apoderado"
				}
				obj.UsuarioComplete = Usuario
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "borrarCurso":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			nombre := string(ctx.QueryArgs().Peek("nombre"))
			obj := GetTemplateConf("", "Borrar Curso", "Información a eliminar", "", "", fmt.Sprintf("/pages/%s", name), "borrar_curso", "Curso")
			obj.FormId = id

			obj.Titulo = fmt.Sprintf("Borrar %v", nombre)

			Usuarios, found := GetUserCurso(db, id)
			if found {
				obj.Lista = Usuarios
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "CodigosQr":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			obj := GetTemplateConf("Codigos Qr", "Formulario", "Complete los campos", "", "", fmt.Sprintf("/pages/%s", name), "", "")
			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)
			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "listaCursos":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			obj := GetTemplateConf("Codigos Qr", "Formulario", "Complete los campos", "", "", fmt.Sprintf("/pages/%s", name), "", "")

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			lcursos, found := GetUserCurso(db, id)
			if found {
				obj.Lista = lcursos
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)
			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "resumenAlumnos1":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			obj := GetTemplateConf("Codigos Qr", "Formulario", "Complete los campos", "", "", fmt.Sprintf("/pages/%s", name), "", "")

			lcursos, found := GetAllCurso(db)
			if found {
				obj.ListaUsers = lcursos
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)
			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "resumenAlumnos2":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			obj := GetTemplateConf("Codigos Qr", "Formulario", "Complete los campos", "", "", fmt.Sprintf("/pages/%s", name), "", "")

			lcursos, found := GetAllCurso(db)
			if found {
				obj.ListaUsers = lcursos
			}

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)
			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	default:
		ctx.NotFound()
	}
}
func Save(ctx *fasthttp.RequestCtx) {

	resp := Response{}
	resp.Op = 2
	resp.Msg = "Error Inesperado"

	id := Read_uint32bytes(ctx.FormValue("id"))
	token := string(ctx.Request.Header.Cookie("cu"))

	db, err := GetMySQLDB()
	defer db.Close()
	if err != nil {
		resp.Msg = "Error Base de Datos"
	}

	switch string(ctx.FormValue("accion")) {
	case "guardar_usuarios":

		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			tipo := string(ctx.FormValue("tipo"))

			nombre := string(ctx.FormValue("nombre"))
			telefono := string(ctx.FormValue("telefono"))
			correo := string(ctx.FormValue("correo"))

			nmatricula := string(ctx.FormValue("nmatricula"))
			rut := string(ctx.FormValue("rut"))
			apellido1 := string(ctx.FormValue("apellido1"))
			apellido2 := string(ctx.FormValue("apellido2"))

			genero := int(Read_uint32bytes(ctx.FormValue("genero")))
			reglamento := int(Read_uint32bytes(ctx.FormValue("reglamento")))
			fecha_nacimiento := string(ctx.FormValue("fecha_nacimiento"))
			fecha_matricula := string(ctx.FormValue("fecha_matricula"))
			fecha_ingreso := string(ctx.FormValue("fecha_ingreso"))

			direccion := string(ctx.FormValue("direccion"))
			fecha_retiro := string(ctx.FormValue("fecha_retiro"))
			motivo_retiro := int(Read_uint32bytes(ctx.FormValue("motivo_retiro")))
			observaciones := string(ctx.FormValue("observaciones"))
			id_cur := Read_uint32bytes(ctx.FormValue("curso"))

			if id > 0 {
				if IsCorreo(db, id, correo) {
					resp.Op, resp.Msg = UpdateUsuario(db, tipo, id, nombre, telefono, correo, nmatricula, rut, apellido1, apellido2, genero, reglamento, fecha_nacimiento, fecha_matricula, fecha_ingreso, direccion, fecha_retiro, motivo_retiro, observaciones, id_cur)
				} else {
					resp.Msg = "Correo Existente"
				}
			}
			if id == 0 {
				if IsCorreo(db, 0, correo) {
					resp.Op, resp.Msg = InsertUsuario(db, tipo, nombre, telefono, correo, nmatricula, rut, apellido1, apellido2, genero, reglamento, fecha_nacimiento, fecha_matricula, fecha_ingreso, direccion, fecha_retiro, motivo_retiro, observaciones, id_cur)
				} else {
					resp.Msg = "Correo Existente"
				}
			}
			if resp.Op == 1 {
				resp.Page = fmt.Sprintf("crearUsuario?tipo=%v", tipo)
				resp.Reload = 1
			}
		} else {
			resp.Msg = "No tiene permisos"
		}

		json.NewEncoder(ctx).Encode(resp)
	case "guardar_cursos":

		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			nombre := string(ctx.FormValue("nombre"))
			orden := Read_uint32bytes(ctx.FormValue("orden"))

			if id > 0 {
				resp.Op, resp.Msg = UpdateCurso(db, id, nombre, orden)
			}
			if id == 0 {
				resp.Op, resp.Msg = InsertCurso(db, nombre, orden)
			}
			if resp.Op == 1 {
				resp.Page = "crearCursos"
				resp.Reload = 1
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_resize":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			picture := string(ctx.FormValue("picture"))
			path := string(ctx.FormValue("path"))
			x := BytesDim(ctx.FormValue("x"))
			y := BytesDim(ctx.FormValue("y"))
			w := BytesDim(ctx.FormValue("w"))
			h := BytesDim(ctx.FormValue("h"))

			dx := Read_uint32bytes(ctx.FormValue("dx"))
			dy := Read_uint32bytes(ctx.FormValue("dy"))

			var path2 string
			if path == "images_preview" {
				path2 = "preview"
			}
			if path == "images_cuentos" {
				path2 = "cuentos"
			}

			found, w1, h1 := GetDims(path, w, h)
			if dx > 0 && dy > 0 {
				w1 = uint(dx)
				h1 = uint(dy)
			}

			if found {
				resp.Op, resp.Msg = CropResizeImage(fmt.Sprintf("img/%v", path2), picture, x, y, w, h, w1, h1)
				if resp.Op == 1 {
					op, msg := ActualizarDimensiones(db, id, w1, h1)
					if op == 1 {
						pass.ListaDel = append(pass.ListaDel, DelImages{Time: time.Now(), Image: fmt.Sprintf("img/%v/x%v", path2, picture)})
						resp.Reload = 1
						if path == "images_preview" {
							resp.Page = fmt.Sprintf("verItemCursoOnline?id=%v", id)
						}
						if path == "images_cuentos" {
							resp.Page = fmt.Sprintf("uploadCuentosImage?id=%v", id)
						}
					} else {
						resp.Op = 2
						resp.Msg = msg
					}
				} else {
					resp.Msg = "Error al cortar la imagen"
				}
			} else {
				resp.Msg = "Error Dimensiones"
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_cursos_online":

		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			nombre := string(ctx.FormValue("nombre"))
			url := string(ctx.FormValue("url"))
			visible := string(ctx.FormValue("visible"))
			nino := string(ctx.FormValue("id_nin"))

			if id > 0 {
				resp.Op, resp.Msg = UpdateCursoOnline(db, id, nombre, url, visible, nino)
			}
			if id == 0 {
				resp.Op, resp.Msg = InsertCursoOnline(db, nombre, url, visible, nino)
			}
			if resp.Op == 1 {
				resp.Page = "verCursosOnline"
				resp.Reload = 1
			}

		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_cursos_online_items":

		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			origen := string(ctx.FormValue("origen"))
			id_cuo := string(ctx.FormValue("id2"))
			if origen == "1" {

				nombre := string(ctx.FormValue("nombre"))
				url := string(ctx.FormValue("url"))
				tipo := string(ctx.FormValue("tipo"))
				if id > 0 {
					if IsUrl1(db, id_cuo, url, id) {
						resp.Op, resp.Msg = UpdateCursoOnlineItem(db, id, nombre, url, tipo)
						files, err := ctx.MultipartForm()
						if err == nil {
							errup1, imagen := UploadFile("./img/preview", files.File["preview"], []string{"JPG", "PNG", "JPEG"}, fmt.Sprintf("prev%v.jpg", id_cuo))
							if errup1 {
								UpdateCursoOnlineItemImage(db, id, imagen.Nom)
							}
						}
					} else {
						resp.Op, resp.Msg = 2, "Url existente #1"
					}
				}
				if id == 0 {
					if IsUrl1(db, id_cuo, url, 0) {
						var ids int64
						resp.Op, resp.Msg, ids = InsertCursoOnlineItem(db, id_cuo, nombre, url, tipo)
						files, err := ctx.MultipartForm()
						if err == nil {
							errup1, imagen := UploadFile("./img/preview", files.File["preview"], []string{"JPG", "PNG", "JPEG"}, fmt.Sprintf("prev%v.jpg", id_cuo))
							if errup1 {
								UpdateCursoOnlineItemImage(db, int(ids), imagen.Nom)
							}
						}
					} else {
						resp.Op, resp.Msg = 2, "Url existente #2"
					}
				}
			} else {
				id_coi := string(ctx.FormValue(fmt.Sprintf("catitem-%v", string(ctx.FormValue("cat")))))
				if IsUrl2(db, id_cuo, id_coi) {
					resp.Op, resp.Msg = AsociarCursoOnlineItem(db, id_cuo, id_coi)
				} else {
					resp.Op, resp.Msg = 2, "Url existente #3"
				}
			}
			if resp.Op == 1 && resp.Reload == 0 {
				resp.Page = fmt.Sprintf("verItemCursoOnline?id=%v", id_cuo)
				resp.Reload = 1
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_image_cuentos":

		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.FormValue("id"))
			id2 := Read_uint32bytes(ctx.FormValue("id2"))
			item, found := GetCursoOnlineItem(db, id)
			if found {
				files, err := ctx.MultipartForm()
				if err == nil {

					var name string
					if id2 == 0 {
						lista_imagen := []ListaImagen{}
						if item.ListaImagen != "" {
							if err := json.Unmarshal([]byte(item.ListaImagen), &lista_imagen); err == nil {
								name = fmt.Sprintf("%v_%v.jpg", id, len(lista_imagen))
							}
						} else {
							name = fmt.Sprintf("%v_0.jpg", id)
						}
					} else {
						name = fmt.Sprintf("%v_%v.jpg", id, id2)
					}

					if name != "" {
						errup1, imagen := UploadFile("./img/cuentos", files.File["cuentos"], []string{"JPG", "PNG", "JPEG"}, name)

						fmt.Println(errup1, imagen)

						if errup1 {
							resp.Op, resp.Msg, resp.Reload, resp.Page = UpdateCursoImage(db, id, item, imagen, id2)
						} else {
							resp.Msg = "Error al subir el archivo"
						}
					} else {
						resp.Msg = "Error nombre de archivo"
					}
				}
			} else {
				resp.Msg = "Curso no encontrado"
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_parentesco":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id_usr := Read_uint32bytes(ctx.FormValue("id_usr"))
			id_apo := Read_uint32bytes(ctx.FormValue("id_apo"))
			apoderado := Read_uint32bytes(ctx.FormValue("apoderado"))
			if id_apo == 0 {
				if id_usr == 0 {
					resp.Msg = "Debe Seleccionar Padre"
				} else {
					if !SelectParentesco(db, id_apo, id) {
						resp.Op, resp.Msg = InsertParentesco(db, id_usr, id, apoderado)
					} else {
						resp.Msg = "Padre ya esta asociado"
					}
				}
			} else {
				resp.Op, resp.Msg = UpdateParentesco(db, id_apo, id, apoderado)
			}
			if resp.Op == 1 {
				resp.Page = fmt.Sprintf("verPadres?id=%v", id)
				resp.Reload = 1
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_parentesco2":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id_usr := Read_uint32bytes(ctx.FormValue("id_usr"))
			if id_usr > 0 {
				if !SelectParentesco(db, id, id_usr) {
					resp.Op, resp.Msg = InsertParentesco(db, id, id_usr, 0)
				} else {
					resp.Msg = "Alumno ya esta asociado"
				}
			} else {
				resp.Msg = "Debe Seleccionar Alumno"
			}

			if resp.Op == 1 {
				resp.Page = fmt.Sprintf("verAlumnos?id=%v", id)
				resp.Reload = 1
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_edu_curso":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id_usr := Read_uint32bytes(ctx.FormValue("id_usr"))
			if id_usr > 0 {
				if !SelectEducadoraCurso(db, id, id_usr) {
					resp.Op, resp.Msg = InsertEducadoraCurso(db, id, id_usr)
				} else {
					resp.Msg = "Educadora ya esta asociada"
				}
			} else {
				resp.Msg = "Debe seleccionar Educadora"
			}
			if resp.Op == 1 {
				resp.Page = fmt.Sprintf("verEducadoras?id=%v", id)
				resp.Reload = 1
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_edu_curso2":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id_cur := Read_uint32bytes(ctx.FormValue("id_cur"))
			if id_cur > 0 {
				if !SelectEducadoraCurso(db, id_cur, id) {
					resp.Op, resp.Msg = InsertEducadoraCurso(db, id_cur, id)
				} else {
					resp.Msg = "El curso ya esta asociado"
				}
			} else {
				resp.Msg = "Debe seleccionar Curso"
			}
			if resp.Op == 1 {
				resp.Page = fmt.Sprintf("verCursos?id=%v", id)
				resp.Reload = 1
			}
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "guardar_libro":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin || perms.Permisos.Educadora {
			nombre := string(ctx.FormValue("nombre"))
			code := string(ctx.FormValue("code"))
			resp.Op, resp.Msg = GuardarLibro(db, nombre, code)
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "prestar_libro":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin || perms.Permisos.Educadora {
			id_alu := Read_uint32bytes(ctx.FormValue("id_alu"))
			id_lib := Read_uint32bytes(ctx.FormValue("id_lib"))
			resp.Op, resp.Msg = PrestarLibro(db, id_lib, id_alu, perms.User.Id_usr)
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	case "devolver_libro":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin || perms.Permisos.Educadora {
			id_lib := Read_uint32bytes(ctx.FormValue("id_lib"))
			resp.Op, resp.Msg = DevolverLibro(db, id_lib)
		} else {
			resp.Msg = "No tiene permisos"
		}
		json.NewEncoder(ctx).Encode(resp)
	default:
		ctx.Response.Header.Set("Content-Type", "application/json")
		json.NewEncoder(ctx).Encode(resp)
	}
}
func Delete(ctx *fasthttp.RequestCtx) {

	resp := Response{}
	ctx.Response.Header.Set("Content-Type", "application/json")
	token := string(ctx.Request.Header.Cookie("cu"))

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	switch string(ctx.FormValue("accion")) {
	case "borrar_user23":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			res := strings.Split(id, "/")
			if len(res) == 2 {
				resp.Tipo, resp.Titulo, resp.Texto = BorrarUsuario23(db, res[0])
				if resp.Tipo == "success" {
					resp.Reload = 1
					resp.Page = fmt.Sprintf("crearUsuario?tipo=%v", res[1])
				}
			} else {
				resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "Error inesperado"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "No tiene permiso para esta acción"
		}
	case "borrar_educadora":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			resp.Tipo, resp.Titulo, resp.Texto = BorrarUsuarioEducadora(db, id)
			if resp.Tipo == "success" {
				resp.Reload = 1
				resp.Page = "crearUsuario?tipo=1"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "No tiene permiso para esta acción"
		}
	case "borrar_parentesco1":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			res := strings.Split(id, "/")
			if len(res) == 2 {
				resp.Tipo, resp.Titulo, resp.Texto = BorrarParentesco(db, res[0], res[1])
				if resp.Tipo == "success" {
					resp.Reload = 1
					resp.Page = fmt.Sprintf("verPadres?id=%v", res[0])
				}
			} else {
				resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "Error inesperado"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "No tiene permiso para esta acción"
		}
	case "borrar_parentesco2":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			res := strings.Split(id, "/")
			if len(res) == 2 {
				resp.Tipo, resp.Titulo, resp.Texto = BorrarParentesco(db, res[0], res[1])
				if resp.Tipo == "success" {
					resp.Reload = 1
					resp.Page = fmt.Sprintf("verAlumnos?id=%v", res[1])
				}
			} else {
				resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "Error inesperado"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "No tiene permiso para esta acción"
		}
	case "borrar_cur_edu1":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			res := strings.Split(id, "/")
			if len(res) == 2 {
				resp.Tipo, resp.Titulo, resp.Texto = BorrarCursoEdu(db, res[0], res[1])
				if resp.Tipo == "success" {
					resp.Reload = 1
					resp.Page = fmt.Sprintf("verEducadoras?id=%v", res[0])
				}
			} else {
				resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "Error inesperado"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "No tiene permiso para esta acción"
		}
	case "borrar_cur_edu2":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			res := strings.Split(id, "/")
			if len(res) == 2 {
				resp.Tipo, resp.Titulo, resp.Texto = BorrarCursoEdu(db, res[0], res[1])
				if resp.Tipo == "success" {
					resp.Reload = 1
					resp.Page = fmt.Sprintf("verCursos?id=%v", res[1])
				}
			} else {
				resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "Error inesperado"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "No tiene permiso para esta acción"
		}
	case "borrar_curso":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			resp.Tipo, resp.Titulo, resp.Texto = BorrarCurso(db, id)
			if resp.Tipo == "success" {
				resp.Reload = 1
				resp.Page = "crearCursos"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Curso", "No tiene permiso para esta acción"
		}
	case "borrar_prestamo":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			resp.Tipo, resp.Titulo, resp.Texto = BorrarPrestamo(db, id)
			if resp.Tipo == "success" {
				resp.Reload = 1
				resp.Page = "Prestamos"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Prestamos", "No tiene permiso para esta acción"
		}
	case "borrar_cuentos_image":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			id := string(ctx.FormValue("id"))
			res := strings.Split(id, "/")
			if len(res) == 2 {
				resp.Tipo, resp.Titulo, resp.Texto = BorrarImageCuentos(db, res[0], res[1])
				if resp.Tipo == "success" {
					resp.Reload = 1
					resp.Page = fmt.Sprintf("uploadCuentosImage?id=%v", res[0])
				}
			} else {
				resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Usuario", "Error inesperado"
			}
		} else {
			resp.Tipo, resp.Titulo, resp.Texto = "error", "Error al eliminar Prestamos", "No tiene permiso para esta acción"
		}
	default:
	}

	json.NewEncoder(ctx).Encode(resp)
}
func Login(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.Set("Content-Type", "application/json")
	resp := Response{Op: 2}

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	user := string(ctx.FormValue("user"))

	cn := 0
	res, err := db.Query("SELECT id_usr, pass FROM usuarios WHERE correo = ? AND eliminado = ?", user, cn)
	defer res.Close()
	ErrorCheck(err)

	if res.Next() {

		var id_usr int
		var pass string
		err := res.Scan(&id_usr, &pass)
		ErrorCheck(err)

		if pass == GetMD5Hash(ctx.FormValue("pass")) {

			resp.Op = 1
			resp.Msg = ""
			cookie := randSeq(32, 0)

			stmt, err := db.Prepare("INSERT INTO sesiones(cookie, id_usr, fecha) VALUES(?,?,NOW())")
			ErrorCheck(err)
			defer stmt.Close()
			r, err := stmt.Exec(string(cookie), id_usr)
			ErrorCheck(err)
			id_usr, err := r.LastInsertId()
			cookieset := fmt.Sprintf("%v-%v", string(cookie), id_usr)
			authcookie := CreateCookie("cu", cookieset, 94608000)
			ctx.Response.Header.SetCookie(authcookie)

		} else {
			resp.Msg = "Usuario Contraseña no existen"
		}

	} else {
		resp.Msg = "Usuario Contraseña no existen"
	}

	json.NewEncoder(ctx).Encode(resp)
}
func Recuperar(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.Set("Content-Type", "application/json")
	resp := Response{Op: 2}

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	user := string(ctx.FormValue("user"))

	cn := 0
	res, err := db.Query("SELECT id_usr FROM usuarios WHERE correo = ? AND eliminado = ?", user, cn)
	defer res.Close()
	ErrorCheck(err)

	if res.Next() {

		var id_usr int
		err := res.Scan(&id_usr)
		ErrorCheck(err)

		code := string(randSeq(30, 1))
		stmt, err := db.Prepare("UPDATE usuarios SET mailcode = ? WHERE id_usr = ?")
		ErrorCheck(err)
		_, err = stmt.Exec(code, id_usr)
		if err == nil {

			resp.Op = 1
			resp.Msg = "Correo Enviado"

			SendEmail(user, "Restablecer Contraseña Jardin ValleEncantado", fmt.Sprintf("<a href='https://www.valleencantado.cl/reestablecer?id=%v&code=%v'>REESTABLECER CONTRASEÑA</a>", id_usr, code))

		} else {
			ErrorCheck(err)
			resp.Msg = "Usuario Contraseña no existen"
		}

	} else {
		resp.Msg = "Usuario Contraseña no existen"
	}

	json.NewEncoder(ctx).Encode(resp)
}
func Nueva(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.Set("Content-Type", "application/json")
	resp := Response{Op: 2}

	pass1 := string(ctx.PostArgs().Peek("pass_01"))
	pass2 := string(ctx.PostArgs().Peek("pass_02"))

	if pass1 == pass2 {

		db, err := GetMySQLDB()
		defer db.Close()
		ErrorCheck(err)

		code := string(ctx.PostArgs().Peek("code"))
		cn := 0
		res, err := db.Query("SELECT id_usr FROM usuarios WHERE code = ? AND eliminado = ?", code, cn)
		defer res.Close()
		ErrorCheck(err)

		if res.Next() {

			pass := GetMD5Hash(ctx.PostArgs().Peek("pass_01"))

			var id_usr int
			err := res.Scan(&id_usr)
			ErrorCheck(err)
			st := ""
			stmt, err := db.Prepare("UPDATE usuarios SET pass = ?, code = ? WHERE id_usr = ?")
			ErrorCheck(err)
			_, e := stmt.Exec(pass, st, id_usr)
			ErrorCheck(e)
			if e == nil {
				resp.Op = 1
				resp.Msg = ""
			}

		} else {
			resp.Msg = "Se produjo un error"
		}
	} else {
		resp.Msg = "Se produjo un error"
	}

	json.NewEncoder(ctx).Encode(resp)
}
func Favicon(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "image/x-icon")
	ctx.SetBody(*favicon)
}
func Redirect(ctx *fasthttp.RequestCtx) (bool, string) {

	host := string(ctx.Host())
	if host != "www.valleencantado.cl" && host != "localhost:81" {
		return true, fmt.Sprintf("https://www.valleencantado.cl%v", string(ctx.URI().RequestURI()))
	} else {
		return false, ""
	}
}
func Index(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	ctx.SetContentType("text/html; charset=utf-8")
	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), true)
	index.Login = Read_uint32bytes(ctx.FormValue("login"))
	index.Page = "Inicio"

	if len(ctx.FormValue("code")) == 30 && Read_uint32bytes(ctx.FormValue("id")) > 0 {
		index.Reestablecer = Reestablecer{Op: 1, Code: string(ctx.FormValue("code")), Id: Read_uint32bytes(ctx.FormValue("id"))}
		index.Login = 1
	}

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html", "html/web/cursosonline.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func AgendaPage(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	ctx.SetContentType("text/html; charset=utf-8")

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), true)
	index.Page = "Agenda"

	if index.Permisos.Educadora || index.Permisos.Admin {
		agenda, found := GetAgendaCurso(db, index.User.Id_usr, "")
		if found {
			index.Agenda = agenda
		}
	}
	if index.Permisos.Apoderado {
		agenda, found := GetHijosAgenda(db, index.User.Id_usr, "")
		if found {

			for i := 0; i < len(agenda.Users); i++ {
				if agenda.Users[i].Comentario != "" || len(agenda.Users) == 1 {
					agenda.Users[i].ShowComentario = true
				}
			}

			index.Agenda = agenda
		}
	}

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html", "html/web/cursosonline.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func SetCurso(ref *MyHandler) error {

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	cursos := CursosOnlineStruct{}

	rescats, found := GetCursoOnline1(db)
	if found {
		cursos.CursosCat = rescats
		ref.CursosOnline = &cursos
		return nil
	} else {
		return fmt.Errorf("Buena")
	}
}
func GetCursoOnline1(db *sql.DB) ([]CursosCat, bool) {

	Resp := []CursosCat{}

	res, err := db.Query("SELECT id_cuo, nombre, url, nino FROM curso_online")
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Resp, false
	}

	for res.Next() {

		re := CursosCat{}
		err := res.Scan(&re.Id, &re.Nombre, &re.Url, &re.Nino)
		if err != nil {
			ErrorCheck(err)
			return Resp, false
		}

		items, found := GetCursoOnline2(db, re.Id)
		if found {
			re.CursosItems = items
		}

		Resp = append(Resp, re)
	}
	return Resp, true
}
func GetCursoOnline2(db *sql.DB, id int) ([]CursosItems, bool) {

	Resp := []CursosItems{}

	res, err := db.Query("SELECT t2.id_coi, t2.nombre, t2.url, t2.image, t2.tipo, t2.Dx, t2.Dy, t2.lista_imagen, t2.url_externo FROM curso_online t1, curso_online_items t2, curso_online_rel t3 WHERE t1.id_cuo = ? AND t1.id_cuo = t3.id_cuo AND t3.id_coi = t2.id_coi", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Resp, false
	}

	for res.Next() {
		re := CursosItems{}
		err := res.Scan(&re.Id, &re.Nombre, &re.Url, &re.Image, &re.Tipo, &re.Dx, &re.Dy, &re.ListaImagen, &re.Urlext)
		if err != nil {
			ErrorCheck(err)
			return Resp, false
		}
		Resp = append(Resp, re)
	}
	return Resp, true
}
func GetCursoOnlineItem(db *sql.DB, id int) (CursosItems, bool) {

	Resp := CursosItems{}

	res, err := db.Query("SELECT id_coi, nombre, url, image, tipo, lista_imagen, Dx, Dy FROM curso_online_items WHERE id_coi = ?", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Resp, false
	}

	if res.Next() {
		err := res.Scan(&Resp.Id, &Resp.Nombre, &Resp.Url, &Resp.Image, &Resp.Tipo, &Resp.ListaImagen, &Resp.Dx, &Resp.Dy)
		if err != nil {
			ErrorCheck(err)
			return Resp, false
		}
		return Resp, true
	}
	return Resp, false
}
func ReestablecerPassword(db *sql.DB, id int, code string, pass1 []byte, pass2 []byte) (uint8, string) {

	if len(code) == 30 {
		if string(pass1) == string(pass2) {

			cn := 0
			res, err := db.Query("SELECT id_usr FROM usuarios WHERE mailcode = ? AND id_usr = ? AND eliminado = ?", code, id, cn)
			defer res.Close()
			ErrorCheck(err)

			if res.Next() {

				var id_usr int
				err := res.Scan(&id_usr)
				ErrorCheck(err)

				str := ""
				stmt, err := db.Prepare("UPDATE usuarios SET pass = ?, mailcode = ? WHERE id_usr = ?")
				ErrorCheck(err)
				_, err = stmt.Exec(GetMD5Hash(pass1), str, id_usr)
				if err == nil {
					return 1, ""
				} else {
					ErrorCheck(err)
					return 2, "Se produjo un error"
				}

			} else {
				return 2, "Se produjo un error"
			}

		} else {
			return 2, "Password Diferentes"
		}
	} else {
		return 2, "Se produjo un error"
	}
}
func CursosOnline(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	ctx.SetContentType("text/html; charset=utf-8")
	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), true)
	index.Login = Read_uint32bytes(ctx.FormValue("login"))
	index.Page = "CursoOnline"
	index.CursosOnline = *pass.CursosOnline

	cursos := []CursosCat{}
	listitems := make(map[int]CursosItems, 0)

	for _, x := range index.CursosOnline.CursosCat {
		if x.Visible == 0 {
			items := []CursosItems{}
			for _, y := range x.CursosItems {
				items = append(items, CursosItems{Id: y.Id})
				listitems[y.Id] = CursosItems{Tipo: y.Tipo, Nombre: y.Nombre, Image: y.Image, ListaImagen: y.ListaImagen, Dx: y.Dx, Dy: y.Dy, Urlext: y.Urlext}
			}
			cursos = append(cursos, CursosCat{Id: x.Id, Url: x.Url, Nombre: x.Nombre, CursosItems: items})
		}
	}

	index.UrlOnline.Cats = cursos

	str1, err := json.Marshal(cursos)
	ErrorCheck(err)
	index.UrlOnline.ListaCursos = string(str1)

	str2, err := json.Marshal(listitems)
	ErrorCheck(err)
	index.UrlOnline.ListaItems = string(str2)

	cat := ctx.UserValue("cat")
	if cat != nil {
		index.UrlOnline.Cat.Url = fmt.Sprintf("%v", ctx.UserValue("cat"))
		for _, x := range index.CursosOnline.CursosCat {
			if x.Url == index.UrlOnline.Cat.Url {
				index.UrlOnline.Cat.Id = x.Id
			}
		}
	}

	item := ctx.UserValue("item")
	if item != nil {
		index.UrlOnline.Item.Url = fmt.Sprintf("%v", ctx.UserValue("item"))
		for _, x := range index.CursosOnline.CursosCat {
			if x.Url == index.UrlOnline.Cat.Url {
				for _, y := range x.CursosItems {
					if y.Url == index.UrlOnline.Item.Url {
						index.UrlOnline.Item.Id = y.Id
						index.UrlOnline.Item.Tipo = y.Tipo
					}
				}
			}
		}
	}

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html", "html/web/cursosonline.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func Salir(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	tkn := string(ctx.Request.Header.Cookie("cu"))
	if len(tkn) > 32 {

		db, err := GetMySQLDB()
		defer db.Close()
		ErrorCheck(err)

		delForm, err := db.Prepare("DELETE FROM sesiones WHERE id_ses=? AND cookie=?")
		ErrorCheck(err)
		delForm.Exec(Read_uint32bytes([]byte(tkn[33:])), tkn[0:32])
		defer db.Close()
	}
	ctx.Redirect("/", 200)
}
func Admin(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	ctx.SetContentType("text/html; charset=utf-8")
	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), false)

	if index.Permisos.Admin {
		t, err := TemplatePage("html/admin/inicio.html")
		ErrorCheck(err)
		err = t.Execute(ctx, index)
		ErrorCheck(err)
	} else {
		ctx.Redirect("/?login=1", 200)
	}
}
func Js(ctx *fasthttp.RequestCtx) {
	jsHandler(ctx)
}
func Css(ctx *fasthttp.RequestCtx) {
	cssHandler(ctx)
}
func Img(ctx *fasthttp.RequestCtx) {
	imgHandler(ctx)
}
func ImgPreview(ctx *fasthttp.RequestCtx) {
	imgPreviewHandler(ctx)
}
func ImgCuentos(ctx *fasthttp.RequestCtx) {
	imgCuentosHandler(ctx)
}
func LibroInicio(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	ctx.SetContentType("text/html; charset=utf-8")
	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), true)
	index.Page = "LibroBase"

	if index.Permisos.Admin || index.Permisos.Educadora {
		prestamos, found := GetTodosLibroPrestados(db)
		if found {
			index.Prestamos = prestamos
		}
	}
	if index.Permisos.Apoderado {
		prestamos, found := GetLibroPrestadosUser(db, index.User.Id_usr)
		if found {
			index.Prestamos = prestamos
		}
	}

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html", "html/web/cursosonline.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func VideoPage(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	// Lee el archivo de video
	video, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", pass.VideoPath, ctx.UserValue("name")))
	if err != nil {
		ctx.Error("Error al leer el archivo de video", fasthttp.StatusInternalServerError)
		return
	}

	// Establece las cabeceras para el video
	ctx.Response.Header.Set("Content-Type", "video/mp4")
	ctx.Response.Header.Set("Content-Length", fmt.Sprint(len(video)))

	// Escribe el video en el cuerpo de la respuesta
	ctx.Write(video)
}
func LibroPage(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	ctx.SetContentType("text/html; charset=utf-8")

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), false)
	index.Page = "Libros"

	code := fmt.Sprintf("%v", ctx.UserValue("name"))
	if len(code) == 30 {
		Libro, found := GetLibro(db, code)
		if found {
			index.Libro = Libro
		}
		Alumnos, found := GetUsuarios(db, 3)
		if found {
			lista, err := json.Marshal(Alumnos)
			if err == nil {
				index.Alumnos = string(lista)
			}
		}
	}
	index.Libro.Code = code

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html", "html/web/cursosonline.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func Qr(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	id, err := strconv.Atoi(fmt.Sprintf("%v", ctx.UserValue("name")))
	if err != nil {
		return
	}

	urlqr := fmt.Sprintf("https://www.valleencantado.cl/libro/%v", string(randSeq(30, id)))

	var q *qrcode.QRCode
	q, err = qrcode.New(urlqr, qrcode.Low)
	q.DisableBorder = true

	var png []byte
	png, err = q.PNG(175)
	if err != nil {
		ErrorCheck(err)
	}

	_, err = ctx.Write(png)
	if err != nil {
		ErrorCheck(err)
	}
}
func HtmlQr(ctx *fasthttp.RequestCtx) {

	if redirect, redirectURL := Redirect(ctx); redirect {
		ctx.Redirect(redirectURL, fasthttp.StatusMovedPermanently)
		return
	}

	ctx.SetContentType("text/html; charset=utf-8")

	cant, err := strconv.Atoi(fmt.Sprintf("%v", ctx.UserValue("name")))
	if err != nil {
		return
	}

	index := Html_Qr{}
	index.List = make([]int, cant)

	t, err := TemplatePage("html/admin/qr.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}

// FUNCTION DB //
func GetMySQLDB() (db *sql.DB, err error) {
	//CREATE DATABASE jardin CHARACTER SET utf8 COLLATE utf8_spanish2_ci;
	db, err = sql.Open("mysql", fmt.Sprintf("root:%v@tcp(127.0.0.1:3306)/jardin", pass.Passwords.PassDb))
	return
}

// USUARIOS DB //
func UpdateUsuario(db *sql.DB, tipo string, id int, nombre string, telefono string, correo string, nmatricula string, rut string, apellido1 string, apellido2 string, genero int, reglamento int, fecha_nacimiento string, fecha_matricula string, fecha_ingreso string, direccion string, fecha_retiro string, motivo_retiro int, observaciones string, id_cur int) (uint8, string) {

	if fecha_nacimiento == "" {
		fecha_nacimiento = "0000-00-00"
	}
	if fecha_matricula == "" {
		fecha_matricula = "0000-00-00"
	}
	if fecha_ingreso == "" {
		fecha_ingreso = "0000-00-00"
	}
	if fecha_retiro == "" {
		fecha_retiro = "0000-00-00"
	}

	stmt, err := db.Prepare("UPDATE usuarios SET nombre = ?, telefono = ?, correo = ?, nmatricula = ?, rut = ?, apellido1 = ?, apellido2 = ?, genero = ?, reglamento = ?, fecha_nacimiento = ?, fecha_matricula = ?, fecha_ingreso = ?, direccion = ?, fecha_retiro = ?, motivo_retiro = ?, observaciones = ? WHERE id_usr = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(nombre, telefono, correo, nmatricula, rut, apellido1, apellido2, genero, reglamento, fecha_nacimiento, fecha_matricula, fecha_ingreso, direccion, fecha_retiro, motivo_retiro, observaciones, id)
	if err == nil {
		if tipo == "3" {
			if AddUserClass(db, id, id_cur) {
				return 1, "Usuario actualizada correctamente"
			} else {
				return 2, "Usuario actualizado pero no se asigno el curso"
			}
		} else {
			return 1, "Usuario ingresado correctamente"
		}
	} else {
		ErrorCheck(err)
		return 2, "El Usuario no pudo ser actualizada"
	}
}
func ActualizarDimensiones(db *sql.DB, id int, w1 uint, h1 uint) (uint8, string) {

	stmt, err := db.Prepare("UPDATE curso_online_items SET Dx = ?, Dy = ? WHERE id_coi = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(w1, h1, id)
	if err == nil {
		return 1, "Item modificado correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Item no pudo ser actualizado"
	}
}
func IsUrl1(db *sql.DB, id_cuo string, url string, id_coi int) bool {

	if id_coi == 0 {
		res, err := db.Query("SELECT * FROM curso_online_rel t1, curso_online_items t2 WHERE t2.url = ? AND t2.id_coi = t1.id_coi AND t1.id_cuo = ?", url, id_cuo)
		defer res.Close()
		if err != nil {
			ErrorCheck(err)
			return false
		}
		if res.Next() {
			return false
		} else {
			return true
		}
	} else {
		res, err := db.Query("SELECT * FROM curso_online_rel t1, curso_online_items t2 WHERE t2.url = ? AND t2.id_coi = t1.id_coi AND t1.id_cuo = ? AND t2.id_coi <> ?", url, id_cuo, id_coi)
		defer res.Close()
		if err != nil {
			ErrorCheck(err)
			return false
		}
		if res.Next() {
			return false
		} else {
			return true
		}
	}
}
func IsUrl2(db *sql.DB, id_cuo string, id_coi string) bool {
	res, err := db.Query("SELECT url FROM curso_online_items WHERE id_coi = ?", id_coi)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
	}
	if res.Next() {
		var url string
		err := res.Scan(&url)
		if err != nil {
			ErrorCheck(err)
			return false
		}
		res2, err := db.Query("SELECT * FROM curso_online_rel t1, curso_online_items t2 WHERE t2.url = ? AND t2.id_coi = t1.id_coi AND t1.id_cuo = ?", url, id_cuo)
		defer res2.Close()
		if err != nil {
			ErrorCheck(err)
		}
		if res2.Next() {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
func InsertCursoOnlineItem(db *sql.DB, id_cuo string, nombre string, url string, tipo string) (uint8, string, int64) {

	stmt, err := db.Prepare("INSERT INTO curso_online_items (nombre, url, tipo) VALUES (?,?,?)")
	ErrorCheck(err)
	defer stmt.Close()
	r, err := stmt.Exec(nombre, url, tipo)
	if err == nil {
		id_coi, err := r.LastInsertId()
		if err == nil {
			stmt2, err := db.Prepare("INSERT INTO curso_online_rel (id_cuo, id_coi) VALUES (?,?)")
			ErrorCheck(err)
			defer stmt2.Close()
			r, err = stmt2.Exec(id_cuo, id_coi)
			if err == nil {
				return 1, "Item ingresado correctamente", id_coi
			} else {
				return 2, "Se produjo un Error", 0
			}
		} else {
			return 2, "Se produjo un Error", 0
		}
	} else {
		ErrorCheck(err)
		return 2, "Se produjo un Error", 0
	}
}
func UpdateCursoOnlineItem(db *sql.DB, id_coi int, nombre string, url string, tipo string) (uint8, string) {

	stmt, err := db.Prepare("UPDATE curso_online_items SET nombre = ?, url = ?, tipo = ? WHERE id_coi = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(nombre, url, tipo, id_coi)
	if err == nil {
		return 1, "Item modificado correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Item no pudo ser actualizado"
	}
}
func UpdateCursoOnlineItemImage(db *sql.DB, id_coi int, image string) (uint8, string) {

	stmt, err := db.Prepare("UPDATE curso_online_items SET image = ? WHERE id_coi = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(image, id_coi)
	if err == nil {
		return 1, "Item modificado correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Item no pudo ser actualizado"
	}
}
func AsociarCursoOnlineItem(db *sql.DB, id_cuo string, id_coi string) (uint8, string) {

	res, err := db.Query("SELECT * FROM curso_online_rel WHERE id_cuo = ? AND id_coi = ?", id_cuo, id_coi)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
	}

	if res.Next() {
		return 2, "Ya existe este Item"
	} else {
		stmt, err := db.Prepare("INSERT INTO curso_online_rel (id_cuo, id_coi) VALUES (?,?)")
		ErrorCheck(err)
		defer stmt.Close()
		_, err = stmt.Exec(id_cuo, id_coi)
		if err == nil {
			return 1, "Usuario ingresado correctamente"
		} else {
			ErrorCheck(err)
			return 2, "Se produjo un Error"
		}
	}
}
func InsertUsuario(db *sql.DB, tipo string, nombre string, telefono string, correo string, nmatricula string, rut string, apellido1 string, apellido2 string, genero int, reglamento int, fecha_nacimiento string, fecha_matricula string, fecha_ingreso string, direccion string, fecha_retiro string, motivo_retiro int, observaciones string, id_cur int) (uint8, string) {

	if fecha_nacimiento == "" {
		fecha_nacimiento = "0000-00-00"
	}
	if fecha_matricula == "" {
		fecha_matricula = "0000-00-00"
	}
	if fecha_ingreso == "" {
		fecha_ingreso = "0000-00-00"
	}
	if fecha_retiro == "" {
		fecha_retiro = "0000-00-00"
	}

	stmt, err := db.Prepare("INSERT INTO usuarios (nombre, telefono, tipo, correo, nmatricula, rut, apellido1, apellido2, genero, reglamento, fecha_nacimiento, fecha_matricula, fecha_ingreso, direccion, fecha_retiro, motivo_retiro, observaciones) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	ErrorCheck(err)
	defer stmt.Close()
	r, err := stmt.Exec(nombre, telefono, tipo, correo, nmatricula, rut, apellido1, apellido2, genero, reglamento, fecha_nacimiento, fecha_matricula, fecha_ingreso, direccion, fecha_retiro, motivo_retiro, observaciones)
	if err == nil {
		if tipo == "3" {
			id_usr, err := r.LastInsertId()
			if err == nil {
				if AddUserClass(db, int(id_usr), id_cur) {
					return 1, "Usuario ingresado correctamente"
				} else {
					return 2, "Usuario ingresado pero no se asigno el curso"
				}
			} else {
				ErrorCheck(err)
				return 2, "El Usuario no pudo ser ingresada"
			}
		} else {
			return 1, "Usuario ingresado correctamente"
		}
	} else {
		ErrorCheck(err)
		return 2, "El Usuario no pudo ser ingresada"
	}
}
func AddUserClass(db *sql.DB, id_usr int, id_cur int) bool {

	res, err := db.Query("SELECT id_cur FROM curso_usuarios WHERE id_usr = ? AND id_cur = ?", id_usr, id_cur)
	defer res.Close()
	ErrorCheck(err)

	if !res.Next() {

		delForm, err := db.Prepare("DELETE FROM curso_usuarios WHERE id_usr = ?")
		if err != nil {
			ErrorCheck(err)
			return false
		}
		_, err = delForm.Exec(id_usr)
		if err != nil {
			ErrorCheck(err)
			return false
		}

		if id_cur > 0 {
			stmt, err := db.Prepare("INSERT INTO curso_usuarios (id_usr, id_cur) VALUES (?,?)")
			if err != nil {
				ErrorCheck(err)
				return false
			}
			defer stmt.Close()
			_, err = stmt.Exec(id_usr, id_cur)
			if err != nil {
				ErrorCheck(err)
				return false
			}
		}
	}
	return true
}
func GetUsuario(db *sql.DB, id int, tipo int) (Usuario, bool) {

	Usuario := Usuario{}

	cn := 0
	res, err := db.Query("SELECT id_usr, nombre, correo, telefono, rut, nmatricula, apellido1, apellido2, genero, reglamento, fecha_nacimiento, fecha_matricula, fecha_ingreso, direccion, fecha_retiro, motivo_retiro, observaciones FROM usuarios WHERE id_usr = ? AND eliminado = ?", id, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Usuario, false
	}

	if res.Next() {
		err := res.Scan(&Usuario.Id_usr, &Usuario.Nombre, &Usuario.Correo, &Usuario.Telefono, &Usuario.Rut, &Usuario.Nmatricula, &Usuario.Apellido1, &Usuario.Apellido2, &Usuario.Genero, &Usuario.Reglamento, &Usuario.FechaNacimiento, &Usuario.FechaMatricula, &Usuario.FechaIngreso, &Usuario.Direccion, &Usuario.FechaRetiro, &Usuario.MotivoRetiro, &Usuario.Observaciones)
		if err != nil {
			ErrorCheck(err)
			return Usuario, false
		}
		if tipo == 3 {
			aux4, found4 := GetCursoUser(db, id)
			if found4 {
				Usuario.Id_cur = aux4
			}
		}
	} else {
		return Usuario, false
	}
	return Usuario, true
}
func GetUsuarioComplete(db *sql.DB, id int) (UsuarioComplete, bool) {

	Usuario := UsuarioComplete{}

	cn := 0
	res, err := db.Query("SELECT id_usr, nombre, apellido1, apellido2, tipo FROM usuarios WHERE id_usr = ? AND eliminado = ?", id, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Usuario, false
	}

	if res.Next() {
		err := res.Scan(&Usuario.Id_usr, &Usuario.Nombre, &Usuario.Apellido1, &Usuario.Apellido2, &Usuario.Tipo)
		if err != nil {
			ErrorCheck(err)
			return Usuario, false
		}

		if Usuario.Tipo == 3 {
			cursos, found := GetCursoUserComplete(db, Usuario.Id_usr)
			if found {
				Usuario.Cursos = cursos
			}
			prestamos, found := GetPrestamosUserComplete(db, Usuario.Id_usr)
			if found {
				Usuario.Prestamos = prestamos
			}
			parentesco, found := GetPadresUser(db, Usuario.Id_usr)
			if found {
				Usuario.Parentesco = parentesco
			}
		}
		if Usuario.Tipo == 2 {
			parentesco, found := GetAlumnosUser(db, Usuario.Id_usr)
			if found {
				Usuario.Parentesco = parentesco
			}
		}

	} else {
		return Usuario, false
	}
	return Usuario, true
}
func GetUsuarios(db *sql.DB, tipo int) ([]Lista, bool) {

	Listas := []Lista{}

	cn := 0
	res, err := db.Query("SELECT id_usr, nombre, apellido1, apellido2, tipo FROM usuarios WHERE tipo = ? AND eliminado = ? ORDER BY apellido1, apellido2, nombre", tipo, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Listas, false
	}

	i := 1
	for res.Next() {
		Lista := Lista{}
		err := res.Scan(&Lista.Id, &Lista.Nombre, &Lista.Apellido1, &Lista.Apellido2, &Lista.Tipo)
		if err != nil {
			ErrorCheck(err)
			return Listas, false
		}
		Lista.Num = i
		i++
		Listas = append(Listas, Lista)
	}
	return Listas, true
}
func GetUsuariosEduAdmin(db *sql.DB) ([]Lista, bool) {

	Listas := []Lista{}

	tipo := 2
	cn := 0
	res, err := db.Query("SELECT id_usr, nombre, tipo FROM usuarios WHERE tipo < ? AND eliminado = ?", tipo, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Listas, false
	}

	for res.Next() {
		Lista := Lista{}
		err := res.Scan(&Lista.Id, &Lista.Nombre, &Lista.Tipo)
		if err != nil {
			ErrorCheck(err)
			return Listas, false
		}
		Listas = append(Listas, Lista)
	}
	return Listas, true
}
func IsCorreo(db *sql.DB, id int, correo string) bool {

	if correo == "" {
		return true
	}

	cn := 0
	var err error
	var res *sql.Rows

	if id == 0 {
		res, err = db.Query("SELECT id_usr FROM usuarios WHERE correo = ? AND eliminado = ?", correo, cn)
	} else {
		res, err = db.Query("SELECT id_usr FROM usuarios WHERE correo = ? AND eliminado = ? AND id_usr <> ?", correo, cn, id)
	}
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return false
	}

	if res.Next() {
		var id_usr int
		err := res.Scan(&id_usr)
		if err != nil {
			ErrorCheck(err)
			return false
		}
		return false
	}
	return true
}
func SelectParentesco(db *sql.DB, id_apo int, id_alu int) bool {

	res, err := db.Query("SELECT apoderado FROM parentensco WHERE id_apo = ? AND id_alu = ?", id_apo, id_alu)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return false
	}

	if res.Next() {
		var apoderado int
		err := res.Scan(&apoderado)
		if err != nil {
			ErrorCheck(err)
			return false
		}
		return true
	}
	return false
}
func InsertParentesco(db *sql.DB, id_apo int, id_alu int, apoderado int) (uint8, string) {

	stmt, err := db.Prepare("INSERT INTO parentensco (apoderado, id_apo, id_alu) VALUES (?,?,?)")
	if err != nil {
		ErrorCheck(err)
		return 2, "El Usuario no pudo ser actualizada"
	}
	defer stmt.Close()
	_, err = stmt.Exec(apoderado, id_apo, id_alu)
	if err == nil {
		return 1, "Usuario ingresado correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Usuario no pudo ser actualizada"
	}
}
func UpdateParentesco(db *sql.DB, id_apo int, id_alu int, apoderado int) (uint8, string) {

	stmt, err := db.Prepare("UPDATE parentensco SET apoderado = ? WHERE id_apo = ? AND id_alu = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(apoderado, id_apo, id_alu)
	if err == nil {
		return 1, "Usuario ingresado correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Usuario no pudo ser actualizada"
	}
}
func UpdateCursoImage(db *sql.DB, id int, item CursosItems, imagen ListaImagenes, id2 int) (uint8, string, int, string) {

	var page string

	lista_imagen := []ListaImagen{}
	if item.ListaImagen != "" {
		if err := json.Unmarshal([]byte(item.ListaImagen), &lista_imagen); err != nil {
			return 2, "Error unmarshal", 0, page
		}
	}

	if id2 == 0 {
		lista_imagen = append(lista_imagen, ListaImagen{Nom: imagen.Nom})
		id2 = len(lista_imagen)
	} else {
		lista_imagen[id2-1].Nom = imagen.Nom
	}

	if item.Dx == 0 && item.Dy == 0 {
		page = fmt.Sprintf("resizeImage?id=%v&num=%v&tipo=1", id, id2)
	} else {
		if item.Dx == imagen.Dx && item.Dy == imagen.Dy {
			page = fmt.Sprintf("verItemCursoOnline?id=%v", id)
		} else {
			page = fmt.Sprintf("resizeImage?id=%v&num=%v&tipo=1", id, id2)
		}
	}

	data, errm := json.Marshal(lista_imagen)
	if errm == nil {
		stmt, err := db.Prepare("UPDATE curso_online_items SET lista_imagen = ? WHERE id_coi = ?")
		ErrorCheck(err)
		_, err = stmt.Exec(string(data), id)
		if err == nil {
			return 1, "Usuario ingresado correctamente", 1, page
		} else {
			ErrorCheck(err)
			return 2, "El Usuario no pudo ser actualizada", 0, page
		}
	} else {
		return 2, "El Usuario no pudo ser actualizada", 0, page
	}
}
func SelectEducadoraCurso(db *sql.DB, id_cur int, id_edu int) bool {

	res, err := db.Query("SELECT id_usr FROM educadora_curso WHERE id_cur = ? AND id_usr = ?", id_cur, id_edu)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return false
	}

	if res.Next() {
		var id_usr int
		err := res.Scan(&id_usr)
		if err != nil {
			ErrorCheck(err)
			return false
		}
		return true
	}
	return false
}
func InsertEducadoraCurso(db *sql.DB, id_cur int, id_edu int) (uint8, string) {

	stmt, err := db.Prepare("INSERT INTO educadora_curso (id_cur, id_usr) VALUES (?,?)")
	if err != nil {
		ErrorCheck(err)
		return 2, "La Educadora no pudo ser asociada"
	}
	defer stmt.Close()
	_, err = stmt.Exec(id_cur, id_edu)
	if err == nil {
		return 1, "La educadora ha sido asociada correctamente"
	} else {
		ErrorCheck(err)
		return 2, "La Educadora no pudo ser asociada"
	}
}

// LIBROS DB //
func GuardarLibro(db *sql.DB, nombre string, code string) (uint8, string) {
	stmt, err := db.Prepare("INSERT INTO libros (nombre, code) VALUES (?,?)")
	ErrorCheck(err)
	defer stmt.Close()
	_, err = stmt.Exec(nombre, code)
	if err == nil {
		return 1, "Libro ingresado correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Libro no pudo ser ingresado"
	}
}
func PrestarLibro(db *sql.DB, id_lib int, id_alu int, id_edu int) (uint8, string) {
	stmt, err := db.Prepare("INSERT INTO prestamos (id_lib, id_alu, fecha_prestamo, id_edu) VALUES (?,?,NOW(),?)")
	ErrorCheck(err)
	defer stmt.Close()
	_, err = stmt.Exec(id_lib, id_alu, id_edu)
	if err == nil {
		return 1, "Libro prestado correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Libro no pudo ser prestado"
	}
}
func DevolverLibro(db *sql.DB, id_lib int) (uint8, string) {
	stmt, err := db.Prepare("UPDATE prestamos SET fecha_devolucion = NOW() WHERE id_lib = ? AND fecha_devolucion = '0000-00-00'")
	ErrorCheck(err)
	_, err = stmt.Exec(id_lib)
	if err == nil {
		return 1, "Libro devuelto correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Libro no pudo ser devuelto"
	}
}

// AGENDA DB //
func IngresarAgenda(db *sql.DB, id_alu int, tipo int, value string) (uint8, string) {

	date := time.Now()
	fecha := date.Format("2006-01-02")

	res, err := db.Query("SELECT id_age FROM agenda WHERE id_usr = ? AND fecha = ?", id_alu, fecha)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return 2, "Error DB"
	}

	var id_age int

	if res.Next() {
		err := res.Scan(&id_age)
		if err != nil {
			ErrorCheck(err)
			return 2, "Error DB"
		}
	} else {
		stmt, err := db.Prepare("INSERT INTO agenda (id_usr, fecha) VALUES (?,NOW())")
		ErrorCheck(err)
		defer stmt.Close()
		r, err := stmt.Exec(id_alu)
		if err != nil {
			ErrorCheck(err)
			return 2, "La agenda no pudo ser actualiza"
		}
		aux_id, err := r.LastInsertId()
		if err != nil {
			ErrorCheck(err)
			return 2, "La agenda no pudo ser actualiza"
		}
		id_age = int(aux_id)
	}

	var x string
	if tipo == 1 {
		x = "ali1"
	} else if tipo == 2 {
		x = "ali2"
	} else if tipo == 3 {
		x = "ali3"
	} else if tipo == 4 {
		x = "dep1"
	} else if tipo == 5 {
		x = "dep2"
	} else if tipo == 6 {
		x = "ausente"
	} else {
		x = "comentario"
	}

	sql := fmt.Sprintf("UPDATE agenda SET ultima_actualizacion = NOW(), %v = ? WHERE id_age = ?", x)
	stmt, err := db.Prepare(sql)
	ErrorCheck(err)
	_, err = stmt.Exec(value, id_age)
	if err == nil {
		return 1, "La agenda fue actualizada exitosamente"
	} else {
		ErrorCheck(err)
		return 2, "La agenda no pudo ser actualiza"
	}
}

// CURSOS DB //
func ChangeUserNom(db *sql.DB, id int, nombre string) (uint8, string) {
	stmt, err := db.Prepare("UPDATE usuarios SET nombre = ? WHERE id_usr = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(nombre, id)
	if err == nil {
		return 1, "Datos actualizados correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Nombre no pudo ser actualizada"
	}
}
func ChangeUserCorreo(db *sql.DB, id int, correo string) (uint8, string) {

	if IsCorreo(db, id, correo) {
		stmt, err := db.Prepare("UPDATE usuarios SET correo = ? WHERE id_usr = ?")
		ErrorCheck(err)
		_, err = stmt.Exec(correo, id)
		if err == nil {
			return 1, "Datos actualizados correctamente"
		} else {
			ErrorCheck(err)
			return 2, "El Correo no pudo ser actualizada"
		}
	} else {
		return 2, "El Correo ya existe"
	}
}
func ChangeUserTelefono(db *sql.DB, id int, nombre string) (uint8, string) {
	stmt, err := db.Prepare("UPDATE usuarios SET telefono2 = ? WHERE id_usr = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(nombre, id)
	if err == nil {
		return 1, "Datos actualizados correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Telefono no pudo ser actualizada"
	}
}
func ExisteUrlCurso(db *sql.DB, url string, id int) (bool, string) {

	cn := 0
	res, err := db.Query("SELECT id_cuo FROM curso_online WHERE url = ? AND eliminado = ?", url, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return false, "Error Base de Datos"
	}

	if res.Next() {
		var id_cuo int
		err := res.Scan(&id_cuo)
		if err != nil {
			ErrorCheck(err)
			return false, "Error Base de Datos"
		}
		if id == id_cuo {
			return true, ""
		} else {
			return false, "Url Existente"
		}
	} else {
		return true, ""
	}
}
func UpdateCursoOnline(db *sql.DB, id int, nombre string, url string, visible string, nino string) (uint8, string) {
	i, errs := ExisteUrlCurso(db, url, id)
	if i {
		stmt, err := db.Prepare("UPDATE curso_online SET nombre = ?, url = ?, visible = ?, nino = ? WHERE id_cuo = ?")
		ErrorCheck(err)
		_, err = stmt.Exec(nombre, url, visible, nino, id)
		if err == nil {
			return 1, "Curso Online actualizada correctamente"
		} else {
			ErrorCheck(err)
			return 2, "El Curso Online no pudo ser actualizada"
		}
	} else {
		return 2, errs
	}
}
func InsertCursoOnline(db *sql.DB, nombre string, url string, visible string, nino string) (uint8, string) {
	i, errs := ExisteUrlCurso(db, url, 0)
	if i {
		stmt, err := db.Prepare("INSERT INTO curso_online (nombre, url, visible, nino) VALUES (?,?,?,?)")
		ErrorCheck(err)
		defer stmt.Close()
		_, err = stmt.Exec(nombre, url, visible, nino)
		if err == nil {
			return 1, "Curso Online ingresado correctamente"
		} else {
			ErrorCheck(err)
			return 2, "El Curso Online no pudo ser ingresado"
		}
	} else {
		return 2, errs
	}
}
func GetAllCurso(db *sql.DB) ([]ListaUsers, bool) {

	LUsers := make([]ListaUsers, 0)

	cn := 0
	res, err := db.Query("SELECT id_usr, nombre, apellido1, apellido2, nmatricula, rut, fecha_nacimiento, direccion, genero FROM usuarios WHERE tipo = 3 AND eliminado = ?", cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return LUsers, false
	}

	i := 1
	for res.Next() {
		User := ListaUsers{}
		err := res.Scan(&User.Id_usr, &User.Nombre, &User.Apellido1, &User.Apellido2, &User.NMatricula, &User.Rut, &User.FechaNac, &User.Direccion, &User.Genero)
		if err != nil {
			ErrorCheck(err)
			return LUsers, false
		}

		User.Num = i
		i++
		padres, found := GetInfoPadres(db, User.Id_usr)
		PrintJson(padres)
		if found {
			for _, x := range padres {
				if x.Tipo == 1 {
					User.MamaNom = x.Nombre
					User.MamaTelefono = x.Telefono
				} else {
					User.PapaNom = x.Nombre
					User.PapaTelefono = x.Telefono
				}
			}
		}

		LUsers = append(LUsers, User)
	}
	return LUsers, true
}
func GetInfoPadres(db *sql.DB, id int) ([]Padres, bool) {

	LPadres := make([]Padres, 0)

	cn := 0
	res, err := db.Query("SELECT t1.nombre, t1.telefono, t1.correo, t1.genero, t2.apoderado FROM usuarios t1, parentensco t2 WHERE t2.id_alu = ? AND t2.id_apo=t1.id_usr AND t1.eliminado = ?", id, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return LPadres, false
	}
	for res.Next() {
		User := Padres{}
		err := res.Scan(&User.Nombre, &User.Telefono, &User.Email, &User.Tipo, &User.Apoderado)
		if err != nil {
			ErrorCheck(err)
			return LPadres, false
		}
		LPadres = append(LPadres, User)
	}
	return LPadres, true
}
func UpdateCurso(db *sql.DB, id int, nombre string, orden int) (uint8, string) {
	stmt, err := db.Prepare("UPDATE cursos SET nombre = ?, orden = ? WHERE id_cur = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(nombre, orden, id)
	if err == nil {
		return 1, "Curso actualizada correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Curso no pudo ser actualizada"
	}
}
func InsertCurso(db *sql.DB, nombre string, orden int) (uint8, string) {
	stmt, err := db.Prepare("INSERT INTO cursos (nombre, orden) VALUES (?,?)")
	ErrorCheck(err)
	defer stmt.Close()
	_, err = stmt.Exec(nombre, orden)
	if err == nil {
		return 1, "Curso ingresado correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Curso no pudo ser ingresado"
	}
}
func GetCurso(db *sql.DB, id int) (Curso, bool) {

	Curso := Curso{}

	cn := 0
	res, err := db.Query("SELECT id_cur, nombre, orden FROM cursos WHERE id_cur = ? AND eliminado = ?", id, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Curso, false
	}

	if res.Next() {
		err := res.Scan(&Curso.Id_cur, &Curso.Nombre, &Curso.Orden)
		if err != nil {
			ErrorCheck(err)
			return Curso, false
		}
	} else {
		return Curso, false
	}
	return Curso, true
}
func GetCursoOnline(db *sql.DB, id int) (CursosCat, bool) {

	Curso := CursosCat{}

	cn := 0
	res, err := db.Query("SELECT id_cuo, nombre, nino, url, visible FROM curso_online WHERE id_cuo = ? AND eliminado = ?", id, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Curso, false
	}

	if res.Next() {
		err := res.Scan(&Curso.Id, &Curso.Nombre, &Curso.Nino, &Curso.Url, &Curso.Visible)
		if err != nil {
			ErrorCheck(err)
			return Curso, false
		}
	} else {
		return Curso, false
	}
	return Curso, true
}
func GetCursos(db *sql.DB) ([]Lista, bool) {

	Listas := []Lista{}

	cn := 0
	res, err := db.Query("SELECT id_cur, nombre FROM cursos WHERE eliminado = ? ORDER BY orden", cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Listas, false
	}

	i := 1
	for res.Next() {
		Lista := Lista{}
		err := res.Scan(&Lista.Id, &Lista.Nombre)
		if err != nil {
			ErrorCheck(err)
			return Listas, false
		}
		Lista.Num = i
		i++
		Listas = append(Listas, Lista)
	}
	return Listas, true
}
func GetCursosOnline(db *sql.DB) ([]Lista, bool) {

	Listas := []Lista{}

	cn := 0
	res, err := db.Query("SELECT id_cuo, nombre FROM curso_online WHERE eliminado = ?", cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Listas, false
	}

	for res.Next() {
		Lista := Lista{}
		err := res.Scan(&Lista.Id, &Lista.Nombre)
		if err != nil {
			ErrorCheck(err)
			return Listas, false
		}
		Listas = append(Listas, Lista)
	}
	return Listas, true
}
func GetHijos(db *sql.DB, id int) ([]Usuario, bool) {

	var Usuarios = []Usuario{}
	var Usuario = Usuario{}

	cn := 0
	res, err := db.Query("SELECT t1.id_usr, t1.nombre FROM usuarios t1, parentensco t2 WHERE t2.id_apo = ? AND t2.id_alu=t1.id_usr AND t1.eliminado = ?", id, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Usuarios, false
	}

	for res.Next() {
		err := res.Scan(&Usuario.Id_usr, &Usuario.Nombre)
		if err != nil {
			ErrorCheck(err)
			return Usuarios, false
		}
		Usuarios = append(Usuarios, Usuario)
	}
	return Usuarios, true
}
func GetHijosAgenda(db *sql.DB, id int, fecha string) (Agenda, bool) {

	var Agenda Agenda
	aux, found := GetPrevNextDate(fecha)
	if found {

		Agenda = aux

		cn := 0
		res, err := db.Query("SELECT t1.id_usr, t1.nombre FROM usuarios t1, parentensco t2 WHERE t2.id_apo = ? AND t2.id_alu=t1.id_usr AND t1.eliminado = ?", id, cn)
		defer res.Close()
		if err != nil {
			ErrorCheck(err)
			return Agenda, false
		}

		var id int
		var nombre string

		for res.Next() {
			err := res.Scan(&id, &nombre)
			if err != nil {
				ErrorCheck(err)
				return Agenda, false
			}
			Agenda.Users = append(Agenda.Users, GetAgendaUser(db, id, aux.Fecha, nombre))
		}
		return Agenda, true
	} else {
		return Agenda, false
	}
}
func GetLibro(db *sql.DB, code string) (Libro, bool) {

	Libros := Libro{}

	//cn := 0
	res, err := db.Query("SELECT nombre, id_lib FROM libros WHERE code = ?", code)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Libros, false
	}

	for res.Next() {

		err := res.Scan(&Libros.Nombre, &Libros.Id_Lib)
		if err != nil {
			ErrorCheck(err)
			return Libros, false
		}
		Libros.Found = true

		prestado, found := GetLibroIsPrestado(db, Libros.Id_Lib)
		if found {
			if prestado.Prestado {
				Libros.Prestado = true
				Libros.NombreAlu = prestado.NombreAlu
				Libros.FechaPrestamo = prestado.FechaPrestamo
			}
		}
	}
	return Libros, true
}
func GetLibroIsPrestado(db *sql.DB, id int) (Libro, bool) {

	libro := Libro{}

	res, err := db.Query("SELECT t2.nombre, t1.fecha_prestamo FROM prestamos t1, usuarios t2 WHERE t1.id_lib = ? AND t1.fecha_devolucion = '0000-00-00 00:00:00' AND t1.id_alu=t2.id_usr", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return libro, false
	}

	for res.Next() {

		libro.Prestado = true
		err := res.Scan(&libro.NombreAlu, &libro.FechaPrestamo)
		if err != nil {
			ErrorCheck(err)
			return libro, false
		}

	}
	return libro, true
}
func GetLibroPrestados(db *sql.DB, id int) ([]Prestamo, bool) {

	Prestamos := []Prestamo{}

	//cn := 0
	res, err := db.Query("SELECT t1.fecha_prestamo, t2.nombre, t3.nombre FROM prestamos t1, usuarios t2, libros t3 WHERE t1.id_edu = ? AND t1.fecha_devolucion = '0000-00-00 00:00:00' AND t1.id_lib=t3.id_lib AND t1.id_alu=t2.id_usr ORDER BY t1.fecha_prestamo", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Prestamos, false
	}

	for res.Next() {

		Prestamo := Prestamo{}

		err := res.Scan(&Prestamo.Fecha_Prestamos, &Prestamo.Nombre_Alu, &Prestamo.Nombre)
		if err != nil {
			ErrorCheck(err)
			return Prestamos, false
		}

		Prestamos = append(Prestamos, Prestamo)
	}
	return Prestamos, true
}
func GetLibroPrestadosHijos(db *sql.DB, id int) (int, bool) {

	var ids int

	res, err := db.Query("SELECT COUNT(*) FROM parentensco t1, usuarios t2, prestamos t3 WHERE t1.id_apo = ? AND t1.id_alu = t2.id_usr AND t2.id_usr = t3.id_alu AND t3.fecha_devolucion = '0000-00-00 00:00:00'", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return ids, false
	}

	if res.Next() {
		err := res.Scan(&ids)
		if err != nil {
			ErrorCheck(err)
			return ids, false
		}
		return ids, true
	}
	return ids, false
}
func GetTodosLibroPrestados(db *sql.DB) ([]Prestamo, bool) {

	Prestamos := []Prestamo{}
	Prestamo := Prestamo{}

	//cn := 0
	res, err := db.Query("SELECT t1.id_pre, t1.fecha_prestamo, t2.nombre, t3.nombre FROM prestamos t1, usuarios t2, libros t3 WHERE t1.fecha_devolucion = '0000-00-00 00:00:00' AND t1.id_lib=t3.id_lib AND t1.id_alu=t2.id_usr ORDER BY t1.fecha_prestamo")
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Prestamos, false
	}

	for res.Next() {

		err := res.Scan(&Prestamo.Id, &Prestamo.Fecha_Prestamos, &Prestamo.Nombre_Alu, &Prestamo.Nombre)
		if err != nil {
			ErrorCheck(err)
			return Prestamos, false
		}

		date, err := time.Parse("2006-01-02", Prestamo.Fecha_Prestamos)
		if err != nil {
			ErrorCheck(err)
		}
		Prestamo.Duration = GetDuration2(int(time.Since(date).Minutes()))
		Prestamos = append(Prestamos, Prestamo)
	}
	return Prestamos, true
}
func GetLibroPrestadosUser(db *sql.DB, id int) ([]Prestamo, bool) {

	Prestamos := []Prestamo{}
	Prestamo := Prestamo{}

	res, err := db.Query("SELECT t2.id_pre, t2.fecha_prestamo, t3.nombre, t4.nombre FROM parentensco t1, prestamos t2, libros t3, usuarios t4 WHERE t1.id_apo = ? AND t1.id_alu=t2.id_alu AND t2.id_alu=t4.id_usr AND t2.fecha_devolucion = '0000-00-00 00:00:00' AND t2.id_lib=t3.id_lib", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Prestamos, false
	}

	for res.Next() {

		err := res.Scan(&Prestamo.Id, &Prestamo.Fecha_Prestamos, &Prestamo.Nombre, &Prestamo.Nombre_Alu)
		if err != nil {
			ErrorCheck(err)
			return Prestamos, false
		}
		date, err := time.Parse("2006-01-02", Prestamo.Fecha_Prestamos)
		if err != nil {
			ErrorCheck(err)
		}
		Prestamo.Duration = GetDuration2(int(time.Since(date).Minutes()))
		Prestamos = append(Prestamos, Prestamo)
	}
	return Prestamos, true
}
func GetAgendaCurso(db *sql.DB, id int, fecha string) (Agenda, bool) {

	var Agenda Agenda
	aux, found := GetPrevNextDate(fecha)
	if found {

		Agenda = aux
		Agenda.AgendaCurso = AgendaCurso{Cursos: make(map[int]Curso, 0)}

		//cn := 0
		res, err := db.Query("SELECT t3.id_usr, t1.id_cur, t3.nombre, t3.apellido1, t4.nombre FROM educadora_curso t1, curso_usuarios t2, usuarios t3, cursos t4 WHERE t1.id_usr = ? AND t1.id_cur=t2.id_cur AND t2.id_usr=t3.id_usr AND t1.id_cur=t4.id_cur", id)
		defer res.Close()
		if err != nil {
			ErrorCheck(err)
			return Agenda, false
		}

		var id_usr int
		var id_cur int
		var nombre_usr string
		var apellido_usr string
		var nombre_cur string

		for res.Next() {
			fmt.Println("HOLA")
			err := res.Scan(&id_usr, &id_cur, &nombre_usr, &apellido_usr, &nombre_cur)
			if err != nil {
				ErrorCheck(err)
				return Agenda, false
			}

			if _, Found := Agenda.AgendaCurso.Cursos[id_cur]; Found {
				Agenda.AgendaCurso.Cursos[id_cur].Users[id_usr] = GetAgendaUser(db, id_usr, aux.Fecha, fmt.Sprintf("%s %s", nombre_usr, apellido_usr))
			} else {
				Agenda.AgendaCurso.Cursos[id_cur] = Curso{Nombre: nombre_cur, Users: make(map[int]AgendaUser, 0)}
				Agenda.AgendaCurso.Cursos[id_cur].Users[id_usr] = GetAgendaUser(db, id_usr, aux.Fecha, fmt.Sprintf("%s %s", nombre_usr, apellido_usr))
			}
		}
		return Agenda, true
	} else {
		return Agenda, false
	}
}
func GetAgendaUser(db *sql.DB, id int, fecha string, nombre string) AgendaUser {

	AgendaUser := AgendaUser{}
	AgendaUser.Nombre = nombre

	//cn := 0
	res, err := db.Query("SELECT ali1, ali2, ali3, dep1, dep2, comentario, ausente, ultima_actualizacion FROM agenda WHERE id_usr = ? AND fecha = ?", id, fecha)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return AgendaUser
	}

	if res.Next() {

		var ultimaact string

		err := res.Scan(&AgendaUser.Ali1, &AgendaUser.Ali2, &AgendaUser.Ali3, &AgendaUser.Dep1, &AgendaUser.Dep2, &AgendaUser.Comentario, &AgendaUser.Ausente, &ultimaact)
		if err != nil {
			ErrorCheck(err)
			return AgendaUser
		}

		date, err := time.Parse("2006-01-02 15:04:05", ultimaact)
		if err != nil {
			ErrorCheck(err)
		}
		AgendaUser.UltimaAct = GetDuration(int(time.Since(date).Minutes()))
		if int(time.Since(date).Minutes()) > 90 {
			AgendaUser.UltimaActAlert = "alert"
		}

		AgendaUser.Data = true
		return AgendaUser
	} else {
		return AgendaUser
	}
}
func GetDuration(minutes int) string {

	months := minutes / (30 * 24 * 60)
	minutes %= 30 * 24 * 60
	days := minutes / (24 * 60)
	minutes %= 24 * 60
	hours := minutes / 60
	minutes %= 60

	result := ""
	if months > 0 {
		result += fmt.Sprintf("%v Meses ", months)
	}
	if days > 0 {
		result += fmt.Sprintf("%v Dias ", days)
	}
	if hours > 0 {
		result += fmt.Sprintf("%v Hrs ", hours)
	}
	if minutes > 0 {
		result += fmt.Sprintf("%v Mins ", minutes)
	}

	return result
}
func GetDuration2(minutes int) string {

	months := minutes / (30 * 24 * 60)
	minutes %= 30 * 24 * 60
	days := minutes / (24 * 60)

	result := ""
	if months > 0 {
		result += fmt.Sprintf("%v Meses ", months)
	}
	if days > 0 {
		result += fmt.Sprintf("%v Dias ", days)
	} else {
		result += fmt.Sprintf("%v Dias ", days)
	}

	return result
}
func GetCursoUser(db *sql.DB, id int) (int, bool) {

	//cn := 0
	res, err := db.Query("SELECT id_cur FROM curso_usuarios WHERE id_usr = ?", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return 0, false
	}

	var ids int
	if res.Next() {

		err := res.Scan(&ids)
		if err != nil {
			ErrorCheck(err)
			return 0, false
		}
		return ids, true
	} else {
		return 0, false
	}
}
func GetUserCurso(db *sql.DB, id int) ([]Lista, bool) {

	lista := []Lista{}
	res, err := db.Query("SELECT t2.id_usr, t2.nombre, t2.apellido1, t2.apellido2 FROM curso_usuarios t1, usuarios t2 WHERE t1.id_cur = ? AND t1.id_usr=t2.id_usr", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return lista, false
	}

	i := 1
	for res.Next() {
		list := Lista{}
		err := res.Scan(&list.Id, &list.Nombre, &list.Apellido1, &list.Apellido2)
		if err != nil {
			ErrorCheck(err)
			return lista, false
		}
		list.Num = i
		i++
		lista = append(lista, list)
	}
	return lista, true
}
func GetCursoUserComplete(db *sql.DB, id int) ([]Curso, bool) {

	cursos := []Curso{}

	res, err := db.Query("SELECT t1.id_cur, t2.nombre FROM curso_usuarios t1, cursos t2 WHERE t1.id_usr = ? AND t1.id_cur=t2.id_cur", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return cursos, false
	}

	for res.Next() {
		curso := Curso{}
		err := res.Scan(&curso.Id_cur, &curso.Nombre)
		if err != nil {
			ErrorCheck(err)
			return cursos, false
		}
		cursos = append(cursos, curso)
	}
	return cursos, true
}
func GetPrestamosUserComplete(db *sql.DB, id int) ([]Prestamo, bool) {

	Prestamos := []Prestamo{}
	Prestamo := Prestamo{}

	//cn := 0
	res, err := db.Query("SELECT t1.id_pre, t1.fecha_prestamo, t2.nombre, t3.nombre FROM prestamos t1, usuarios t2, libros t3 WHERE t1.fecha_devolucion = '0000-00-00 00:00:00' AND t1.id_alu=? AND t1.id_lib=t3.id_lib AND t1.id_alu=t2.id_usr", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Prestamos, false
	}

	for res.Next() {

		err := res.Scan(&Prestamo.Id, &Prestamo.Fecha_Prestamos, &Prestamo.Nombre_Alu, &Prestamo.Nombre)
		if err != nil {
			ErrorCheck(err)
			return Prestamos, false
		}
		Prestamos = append(Prestamos, Prestamo)
	}
	return Prestamos, true
}
func BorrarUsuario23(db *sql.DB, id string) (string, string, string) {

	delForm, err := db.Prepare("DELETE FROM usuarios WHERE id_usr = ?")
	ErrorCheck(err)
	_, e := delForm.Exec(id)
	defer db.Close()

	ErrorCheck(e)
	if e == nil {
		return "success", "Usuario eliminado", "Usuario eliminado correctamente"
	} else {
		return "error", "Error al eliminar el usuario", "El usuario no pudo ser eliminada"
	}
}
func BorrarUsuarioEducadora(db *sql.DB, id string) (string, string, string) {
	del := 1
	stmt, err := db.Prepare("UPDATE usuarios SET eliminado = ? WHERE id_usr = ?")
	_, err = stmt.Exec(del, id)
	if err == nil {
		return "success", "Educadora eliminada", "Educadora eliminada correctamente"
	} else {
		return "error", "Error al eliminar Educadora", "La Educadora no pudo ser eliminada"
	}
}
func BorrarImageCuentos(db *sql.DB, id_cois string, nums string) (string, string, string) {

	id_coi, err := strconv.Atoi(fmt.Sprintf("%v", id_cois))
	if err != nil {
		return "error", "", ""
	}

	num, err := strconv.Atoi(fmt.Sprintf("%v", nums))
	if err != nil {
		return "error", "", ""
	}

	item, found := GetCursoOnlineItem(db, id_coi)

	if found {
		var listaImage []ListaImagen
		var listaImage2 []ListaImagen
		if err := json.Unmarshal([]byte(item.ListaImagen), &listaImage); err == nil {
			for i, x := range listaImage {
				if i == num {
					pass.ListaDel = append(pass.ListaDel, DelImages{Time: time.Now(), Image: fmt.Sprintf("img/cuentos/%v", x.Nom)})
				} else {
					listaImage2 = append(listaImage2, ListaImagen{Nom: x.Nom})
				}
			}
			u, err := json.Marshal(listaImage2)
			if err == nil {
				stmt, err := db.Prepare("UPDATE curso_online_items SET lista_imagen = ? WHERE id_coi = ?")
				_, err = stmt.Exec(string(u), id_coi)
				if err == nil {
					return "success", "Imagen eliminada", "Imagen eliminada correctamente"
				} else {
					return "error", "Error al eliminar la Imagen", "La Imagen no pudo ser eliminada"
				}
			} else {
				return "error", "", ""
			}
		} else {
			return "error", "", ""
		}
	} else {
		return "error", "", ""
	}
}
func GetPadresUser(db *sql.DB, id int) ([]Padres, bool) {

	padres := []Padres{}
	obj := Padres{}
	//cn := 0
	res, err := db.Query("SELECT t2.id_usr, t2.nombre, t1.tipo, t1.apoderado, t2.correo, t2.telefono FROM parentensco t1, usuarios t2 WHERE t1.id_alu = ? AND t1.id_apo=t2.id_usr", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return padres, false
	}

	for res.Next() {

		err := res.Scan(&obj.Id_usr, &obj.Nombre, &obj.Tipo, &obj.Apoderado, &obj.Email, &obj.Telefono)
		if err != nil {
			ErrorCheck(err)
			return padres, false
		}
		padres = append(padres, obj)
	}
	return padres, true
}
func GetAlumnosUser(db *sql.DB, id int) ([]Padres, bool) {

	padres := []Padres{}
	obj := Padres{}
	//cn := 0
	res, err := db.Query("SELECT t2.id_usr, t2.nombre, t1.tipo, t1.apoderado, t2.correo, t2.telefono FROM parentensco t1, usuarios t2 WHERE t1.id_apo = ? AND t1.id_alu=t2.id_usr", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return padres, false
	}

	for res.Next() {

		err := res.Scan(&obj.Id_usr, &obj.Nombre, &obj.Tipo, &obj.Apoderado, &obj.Email, &obj.Telefono)
		if err != nil {
			ErrorCheck(err)
			return padres, false
		}
		padres = append(padres, obj)
	}
	return padres, true
}
func GetPadreUser(db *sql.DB, id int, id2 int) (Padres, bool) {

	padre := Padres{}
	res, err := db.Query("SELECT t2.id_usr, t2.nombre, t1.tipo, t1.apoderado, t2.correo, t2.telefono FROM parentensco t1, usuarios t2 WHERE t1.id_alu = ? AND t1.id_apo = ? AND t1.id_apo=t2.id_usr", id, id2)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return padre, false
	}

	if res.Next() {
		err := res.Scan(&padre.Id_usr, &padre.Nombre, &padre.Tipo, &padre.Apoderado, &padre.Email, &padre.Telefono)
		if err != nil {
			ErrorCheck(err)
			return padre, false
		}
		return padre, true
	} else {
		return padre, false
	}
}
func GetUserCursos(db *sql.DB, id int) ([]Padres, bool) {

	padres := []Padres{}

	res, err := db.Query("SELECT t2.id_usr, t2.nombre FROM educadora_curso t1, usuarios t2 WHERE t1.id_cur = ? AND t1.id_usr=t2.id_usr", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return padres, false
	}

	for res.Next() {
		obj := Padres{}
		err := res.Scan(&obj.Id_usr, &obj.Nombre)
		if err != nil {
			ErrorCheck(err)
			return padres, false
		}
		padres = append(padres, obj)
	}
	return padres, true
}
func GetCursosUser(db *sql.DB, id int) ([]Padres, bool) {

	padres := []Padres{}

	res, err := db.Query("SELECT t2.id_cur, t2.nombre FROM educadora_curso t1, cursos t2 WHERE t1.id_usr = ? AND t1.id_cur=t2.id_cur", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return padres, false
	}

	for res.Next() {
		obj := Padres{}
		err := res.Scan(&obj.Id_usr, &obj.Nombre)
		if err != nil {
			ErrorCheck(err)
			return padres, false
		}
		padres = append(padres, obj)
	}
	return padres, true
}
func BorrarParentesco(db *sql.DB, id1 string, id2 string) (string, string, string) {

	delForm, err := db.Prepare("DELETE FROM parentensco WHERE id_alu = ? AND id_apo = ?")
	ErrorCheck(err)
	_, e := delForm.Exec(id1, id2)
	defer db.Close()

	ErrorCheck(e)
	if e == nil {
		return "success", "Parentesco eliminado", "Parentesco eliminado correctamente"
	} else {
		return "error", "Error al eliminar el parentesco", "El parentesco no pudo ser eliminada"
	}
}
func BorrarCurso(db *sql.DB, id string) (string, string, string) {

	delForm, err := db.Prepare("DELETE FROM cursos WHERE id_cur = ?")
	ErrorCheck(err)
	_, e := delForm.Exec(id)
	defer db.Close()

	ErrorCheck(e)
	if e == nil {
		return "success", "Curso eliminado", "Curso eliminado correctamente"
	} else {
		return "error", "Error al eliminar el curso", "El curso no pudo ser eliminada"
	}
}
func BorrarPrestamo(db *sql.DB, id string) (string, string, string) {

	delForm, err := db.Prepare("DELETE FROM prestamos WHERE id_pre = ?")
	ErrorCheck(err)
	_, e := delForm.Exec(id)
	defer db.Close()

	ErrorCheck(e)
	if e == nil {
		return "success", "Prestamo eliminado", "Prestamo eliminado correctamente"
	} else {
		return "error", "Error al eliminar el Prestamo", "El Prestamo no pudo ser eliminada"
	}
}
func BorrarCursoEdu(db *sql.DB, id1 string, id2 string) (string, string, string) {

	fmt.Printf("ID_CUR %v - ID_USR %v\n", id1, id2)

	delForm, err := db.Prepare("DELETE FROM educadora_curso WHERE id_cur = ? AND id_usr = ?")
	ErrorCheck(err)
	_, e := delForm.Exec(id1, id2)
	defer db.Close()

	ErrorCheck(e)
	if e == nil {
		return "success", "Relacion eliminada", "Relacion eliminada correctamente"
	} else {
		return "error", "Error al eliminar el relacion", "La relacion no pudo ser eliminada"
	}
}
func GetEducadoraCurso(db *sql.DB, id int, id2 int) (Padres, bool) {

	padre := Padres{}
	res, err := db.Query("SELECT t2.id_usr, t2.nombre, t1.tipo, t1.apoderado, t2.correo, t2.telefono FROM curso_usuarios t1, usuarios t2 WHERE t1.id_usr = ? AND t1.id_cur = ? AND t1.id_usr=t2.id_usr", id, id2)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return padre, false
	}

	if res.Next() {
		err := res.Scan(&padre.Id_usr, &padre.Nombre, &padre.Tipo, &padre.Apoderado, &padre.Email, &padre.Telefono)
		if err != nil {
			ErrorCheck(err)
			return padre, false
		}
		return padre, true
	} else {
		return padre, false
	}
}

// OTHER FUNCTIONS //
func GetPrevNextDate(fecha string) (Agenda, bool) {

	var Agenda = Agenda{FechaIsNext: 1, FechaIsPrev: 1}

	var date time.Time
	var dateprev time.Time
	var datenext time.Time
	var err error

	if fecha == "" {
		date = time.Now()
	} else {
		date, err = time.Parse("2006-01-02", fecha)
		if err != nil {
			ErrorCheck(err)
			return Agenda, false
		}
	}

	dateprev = date
	datenext = date

	weekday := int(date.Weekday())
	if weekday == 0 {
		date = date.Add(-48 * time.Hour)
		dateprev = dateprev.Add(-72 * time.Hour)
		datenext = datenext.Add(24 * time.Hour)
	} else if weekday == 6 {
		date = date.Add(-24 * time.Hour)
		dateprev = dateprev.Add(-48 * time.Hour)
		datenext = datenext.Add(48 * time.Hour)
	} else if weekday == 5 {
		dateprev = dateprev.Add(-24 * time.Hour)
		datenext = datenext.Add(72 * time.Hour)
	} else if weekday == 1 {
		dateprev = dateprev.Add(-72 * time.Hour)
		datenext = datenext.Add(24 * time.Hour)
	} else {
		dateprev = dateprev.Add(-24 * time.Hour)
		datenext = datenext.Add(24 * time.Hour)
	}

	Agenda.FechaStr = GetFechaStr(date)
	Agenda.Fecha = date.Format("2006-01-02")
	Agenda.FechaPrev = dateprev.Format("2006-01-02")
	Agenda.FechaNext = datenext.Format("2006-01-02")

	if datenext.Sub(time.Now()).Hours() > 0 {
		Agenda.FechaIsNext = 0
	}

	return Agenda, true
}
func GetFechaStr(fecha time.Time) string {

	var mes string
	var wd string

	_, month, day := fecha.Date()
	switch int(month) {
	case 1:
		mes = "Enero"
	case 2:
		mes = "Febrero"
	case 3:
		mes = "Marzo"
	case 4:
		mes = "Abril"
	case 5:
		mes = "Mayo"
	case 6:
		mes = "Junio"
	case 7:
		mes = "Julio"
	case 8:
		mes = "Agosto"
	case 9:
		mes = "Septiembre"
	case 10:
		mes = "Octubre"
	case 11:
		mes = "Noviembre"
	case 12:
		mes = "Diciembre"
	}

	weekday := int(fecha.Weekday())
	switch weekday {
	case 0:
		wd = "Domingo"
	case 1:
		wd = "Lunes"
	case 2:
		wd = "Martes"
	case 3:
		wd = "Miercoles"
	case 4:
		wd = "Jueves"
	case 5:
		wd = "Viernes"
	case 6:
		wd = "Sabado"
	}

	return fmt.Sprintf("%v %v de %v", wd, day, mes)
}
func TemplatePage(v string) (*template.Template, error) {

	t, err := template.ParseFiles(v)
	if err != nil {
		log.Print(err)
		return t, err
	}
	return t, nil
}
func TemplatePages(v1 string, v2 string, v3 string, v4 string, v5 string, v6 string) (*template.Template, error) {

	t, err := template.ParseFiles(v1, v2, v3, v4, v5, v6)
	if err != nil {
		log.Print(err)
		return t, err
	}
	return t, nil
}
func GetMD5Hash(text []byte) string {
	hasher := md5.New()
	hasher.Write(text)
	return hex.EncodeToString(hasher.Sum(nil))
}
func CreateCookie(key string, value string, expire int) *fasthttp.Cookie {
	if strings.Compare(key, "") == 0 {
		key = "GoLog-Token"
	}
	authCookie := fasthttp.Cookie{}
	authCookie.SetKey(key)
	authCookie.SetValue(value)
	authCookie.SetMaxAge(expire)
	authCookie.SetHTTPOnly(true)
	authCookie.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	return &authCookie
}
func GetRandom(min int, max int, id int) int {
	rand.Seed(time.Now().UnixNano() * int64(id))
	return rand.Intn(max-min) + min
}
func randSeq(n int, x int) []byte {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = letters[GetRandom(0, len(letters), x+i)]
	}
	return b
}
func GetTemplateConf(titulo string, subtitulo string, subtitulo2 string, titulolista string, formaccion string, pagemod string, delaccion string, delobj string) TemplateConf {
	return TemplateConf{Titulo: titulo, SubTitulo: subtitulo, SubTitulo2: subtitulo2, TituloLista: titulolista, FormAccion: formaccion, PageMod: pagemod, DelAccion: delaccion, DelObj: delobj}
}
func PrintJson(b interface{}) {
	u, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u))
}
func ErrorCheck(e error) {
	if e != nil {
		fmt.Println("ERROR:", e)
	}
}
func Read_uint32bytes(data []byte) int {
	var x int
	for _, c := range data {
		x = x*10 + int(c-'0')
	}
	return x
}
func GetPermisoUser(db *sql.DB, tkn string, complete bool) Indexs {

	Index := Indexs{}

	if len(tkn) > 32 {

		cn := 0
		res, err := db.Query("SELECT t1.id_usr, t1.cant_agenda, t1.tipo, t1.nombre, t1.correo, t1.telefono, t1.telefono2 FROM usuarios t1, sesiones t2 WHERE t2.id_ses = ? AND t2.cookie = ? AND t2.id_usr=t1.id_usr AND t1.eliminado = ?", Read_uint32bytes([]byte(tkn[33:])), tkn[0:32], cn)
		defer res.Close()
		ErrorCheck(err)

		if res.Next() {
			Index.Register = true
			err := res.Scan(&Index.User.Id_usr, &Index.User.CantAgenda, &Index.User.Tipo, &Index.User.Nombre, &Index.User.Correo, &Index.User.Telefono, &Index.User.Telefono2)
			ErrorCheck(err)

			if Index.User.Tipo == 0 {
				// ADMINISTRADOR
				Index.Permisos.Admin = true
				if complete {
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Administrador", Icon: "admin", Url: "admin"})
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Resumen Libro", Icon: "iclibro", Url: "libros"})
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Agenda Virtual", Icon: "icagenda", Url: "agenda"})
				}
			} else if Index.User.Tipo == 1 {
				// PERSONAL
				Index.Permisos.Educadora = true
				if complete {
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Resumen Libro", Icon: "iclibro", Url: "libros"})
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Agenda Virtual", Icon: "icagenda", Url: "agenda"})
				}
			} else if Index.User.Tipo == 2 {
				// PADRES

				Index.Permisos.Apoderado = true
				if complete {

					textAgenda := "Sin información"
					Agenda, found := GetHijosAgenda(db, Index.User.Id_usr, "")
					if found {
						textAgenda = fmt.Sprintf("hay %v registro", len(Agenda.Users))
					}
					textLibro := "Sin libros prestados"
					Libro, found := GetLibroPrestadosHijos(db, Index.User.Id_usr)
					if found {
						textLibro = fmt.Sprintf("%v libro prestado", Libro)
					}
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Biblioteca", Icon: "iclibro", Url: "libros", Text: textLibro})
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Agenda Virtual", Icon: "icagenda", Url: "agenda", Text: textAgenda})
				}
			} else {
				// OTROS
			}
		} else {
			Index.Register = false
		}
	}
	return Index
}

// OTHER FUNCTIONS //

// DAEMON //
func (h *MyHandler) StartDaemon() {
	h.Conf.Tiempo = 10 * time.Second
	for i, x := range h.ListaDel {
		e1 := os.Remove(x.Image)
		if e1 == nil {
			h.ListaDel = removeSlice(h.ListaDel, i)
			break
		}
	}
}
func removeSlice(s []DelImages, i int) []DelImages {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (c *Config) init() {
	var tick = flag.Duration("tick", 1*time.Second, "Ticking interval")
	c.Tiempo = *tick
}
func run(con context.Context, c *MyHandler, stdout io.Writer) error {
	c.Conf.init()
	log.SetOutput(os.Stdout)
	for {
		select {
		case <-con.Done():
			return nil
		case <-time.Tick(c.Conf.Tiempo):
			c.StartDaemon()
		}
	}
}

// DAEMON //
func SendEmail(to string, subject string, body string) bool {

	from := "valleencantado.cl@gmail.com"
	sub := fmt.Sprintf("From:%v\nTo:%v\nSubject:%v\n", from, to, subject)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass.Passwords.PassEmail, "smtp.gmail.com"), from, []string{to}, []byte(sub+mime+body))
	if err != nil {
		return false
	}
	return true
}
func CropResizeImage(path string, files string, x int, y int, w int, h int, w1 uint, h1 uint) (uint8, string) {

	fmt.Println(png.UnsupportedError("aa"))
	fmt.Println(jpeg.UnsupportedError("aa"))

	file, err := os.Open(fmt.Sprintf("%v/x%v", path, files))
	if err != nil {
		fmt.Println("ERROR OPEN IMAGE", err)
		return 2, ""
	}
	defer file.Close()

	originalImage, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("ERROR DECODE IMAGE", err)
		return 2, ""
	}

	img := originalImage.(SubImager).SubImage(image.Rect(x, y, w+x, h+y))

	rgbaImg := image.NewRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			rgbaImg.Set(x, y, img.At(x, y))
		}
	}

	resizedImg := resize.Resize(w1, h1, rgbaImg, resize.Lanczos3)

	croppedImageFile, err := os.Create(fmt.Sprintf("%v/%v", path, files))
	if err != nil {
		fmt.Println("ERROR CREATE IMAGE", err)
		return 2, ""
	}

	defer croppedImageFile.Close()
	if err := jpeg.Encode(croppedImageFile, resizedImg, &jpeg.Options{Quality: 75}); err != nil {
		fmt.Println("ERROR ENCODE IMAGE", err)
		return 2, ""
	}

	return 1, "Imagen ha sido redimensionada con exito"
}
func Extension(filename string, extends []string) bool {

	upper := strings.ToUpper(filename)
	bupper := []byte(upper)
	point := -1
	for i, x := range bupper {
		if x == 46 {
			point = i
		}
	}
	if point > 0 {
		for _, x := range extends {
			if x == string(bupper[point+1:]) {
				return true
			}
		}
	}
	return false
}
func NombreExtension(filename string) (bool, string, string) {

	upper := strings.ToUpper(filename)
	bupper := []byte(upper)
	point := -1
	for i, x := range bupper {
		if x == 46 {
			point = i
		}
	}
	if point > 0 {
		return true, strings.ToLower(string(bupper[:point])), strings.ToLower(string(bupper[point+1:]))
	}
	return false, "", ""
}
func ImageDimenciones(path string) (bool, int, int, float64) {

	archivoImagen, err := os.Open(path)
	if err != nil {
		fmt.Println("Error al abrir la imagen:", err)
		return false, 0, 0, 0
	}
	defer archivoImagen.Close()

	img, _, err := image.Decode(archivoImagen)
	if err != nil {
		fmt.Println("Error al decodificar la imagen:", err)
		return false, 0, 0, 0
	}

	return true, img.Bounds().Dx(), img.Bounds().Dy(), float64(img.Bounds().Dx()) / float64(img.Bounds().Dy())
}
func BytesDim(data []byte) int {
	var x int
	for _, c := range data {
		if c == 46 {
			return x
		} else {
			x = x*10 + int(c-'0')
		}
	}
	return x
}
func FileExist(path string, file string) bool {
	if _, serr := os.Stat(path); serr != nil {
		err := os.MkdirAll(path, os.ModeDir)
		if err != nil {
			return true
		}
	}
	return false
}
func UploadFile(path string, files []*multipart.FileHeader, extends []string, filename string) (bool, ListaImagenes) {

	if len(files) == 0 {
		return false, ListaImagenes{}
	}

	fexist := false

	file, err := files[0].Open()
	if err != nil {
		return false, ListaImagenes{}
	}
	defer file.Close()
	if filename == "" {
		filename = files[0].Filename
	}

	if FileExist(path, filename) {
		fexist = true
		e := os.Rename(fmt.Sprintf("%v/%v", path, filename), fmt.Sprintf("%v/temp_%v", path, filename))
		if e != nil {
			return false, ListaImagenes{}
		}
	}

	if Extension(filename, extends) {
		out, err := os.Create(fmt.Sprintf("%v/x%v", path, filename))
		defer out.Close()
		if err == nil {
			_, err = io.Copy(out, file)
			if err == nil {
				if fexist {
					e1 := os.Remove(fmt.Sprintf("%v/temp_%v", path, filename))
					if e1 != nil {
						return false, ListaImagenes{}
					}
				}
				found, dx, dy, ratio := ImageDimenciones(fmt.Sprintf("%v/x%v", path, filename))
				if found {
					return true, ListaImagenes{Nom: filename, Dx: dx, Dy: dy, Ratio: ratio}
				} else {
					return false, ListaImagenes{}
				}
			} else {
				return false, ListaImagenes{}
			}
		} else {
			return false, ListaImagenes{}
		}
	} else {
		return false, ListaImagenes{}
	}
}
func GetDims(path string, w int, h int) (bool, uint, uint) {
	if path == "images_preview" {
		return true, 135, 90
	}
	if path == "images_cuentos" {

		var maxwidth float32 = 500
		var maxheight float32 = 400

		newheight := float32(maxwidth) * float32(h) / float32(w)

		if newheight <= maxheight {
			return true, uint(maxwidth), uint(newheight)
		} else {
			newwidth := uint(float32(w) * float32(maxheight) / float32(h))
			return true, uint(newwidth), uint(maxheight)
		}
	}
	return false, 0, 0
}
