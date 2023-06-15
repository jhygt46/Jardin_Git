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
	"math/rand"
	"strconv"
	"strings"

	//"image/draw"

	"io"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

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
	FormAccion      string          `json:"FormAccion"`
	PageMod         string          `json:"PageMod"`
	DelAccion       string          `json:"DelAccion"`
	DelObj          string          `json:"DelObj"`
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
}
type Lista struct {
	Id     int    `json:"Id"`
	Nombre string `json:"Nombre"`
	Tipo   int    `json:"Tipo"`
}
type Config struct {
	Tiempo time.Duration `json:"Tiempo"`
}
type MyHandler struct {
	Conf Config `json:"Conf"`
}
type Indexs struct {
	Login     int           `json:"Login"`
	Function  string        `json:"Nombre"`
	Page      string        `json:"Page"`
	User      IndexUser     `json:"User"`
	Modulos   []Modulo      `json:"Modulos"`
	Register  bool          `json:"Register"`
	Permisos  IndexPermisos `json:"Permisos"`
	Libro     Libro         `json:"Libro"`
	Agenda    Agenda        `json:"Agenda"`
	Alumnos   string        `json:"Alumnos"`
	Prestamos []Prestamo    `json:"Prestamos"`
}
type IndexPermisos struct {
	Admin     bool `json:"Admin"`
	Educadora bool `json:"Educadora"`
	Apoderado bool `json:"Apoderado"`
}
type IndexUser struct {
	Id_usr    int    `json:"Id_usr"`
	Tipo      int    `json:"Tipo"`
	Nombre    string `json:"Nombre"`
	Correo    string `json:"Correo"`
	Telefono  string `json:"Telefono"`
	Telefono2 string `json:"Telefono2"`
}
type Modulo struct {
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

var (
	imgHandler fasthttp.RequestHandler
	cssHandler fasthttp.RequestHandler
	jsHandler  fasthttp.RequestHandler
	port       string
	favicon    *[]byte
)
var pass = &MyHandler{Conf: Config{}}

func main() {

	if runtime.GOOS == "windows" {
		imgHandler = fasthttp.FSHandler("C:/Go/Jardin/img", 1)
		cssHandler = fasthttp.FSHandler("C:/Go/Jardin/css", 1)
		jsHandler = fasthttp.FSHandler("C:/Go/Jardin/js", 1)
		port = ":81"
	} else {
		imgHandler = fasthttp.FSHandler("/var/Jardin/img", 1)
		cssHandler = fasthttp.FSHandler("/var/Jardin/css", 1)
		jsHandler = fasthttp.FSHandler("/var/Jardin/js", 1)
		port = ":80"
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
		r.GET("/pages/{name}", Pages)
		r.POST("/login", Login)
		r.POST("/nueva", Nueva)
		r.POST("/save", Save)
		r.POST("/delete", Delete)
		r.POST("/accion", Accion)
		r.GET("/salir", Salir)
		r.GET("/admin", Admin)
		r.GET("/libro", LibroInicio)
		r.GET("/libro/{name}", LibroPage)
		r.GET("/agenda", AgendaPage)
		r.GET("/qr/{name}", Qr)
		r.GET("/qrhtml/{name}", HtmlQr)

		// ANTES
		fasthttp.ListenAndServe(port, r.Handler)

		//secureMiddleware := secure.New(secure.Options{SSLRedirect: true})
		//secureHandler := secureMiddleware.Handler(r.Handler)
		//go func() { log.Fatal(fasthttp.ListenAndServe(":80", secureHandler)) }()
		//log.Fatal(fasthttp.ListenAndServeTLS(":443", "/etc/letsencrypt/live/www.valleencantado.cl/fullchain.pem", "/etc/letsencrypt/live/www.valleencantado.cl/privkey.pem", secureHandler))

	}()
	if err := run(con, pass, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
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
	case "user_nom":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Apoderado {
			nombre := string(ctx.FormValue("value"))
			resp.Op, resp.Msg = ChangeUserNom(db, perms.User.Id_usr, nombre)
		} else {
			resp.Msg = "No tiene permisos"
		}
	case "user_correo":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Apoderado {
			correo := string(ctx.FormValue("value"))
			resp.Op, resp.Msg = ChangeUserCorreo(db, perms.User.Id_usr, correo)
		} else {
			resp.Msg = "No tiene permisos"
		}
	case "user_telefono":
		ctx.Response.Header.Set("Content-Type", "application/json")
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Apoderado {
			telefono := string(ctx.FormValue("value"))
			resp.Op, resp.Msg = ChangeUserTelefono(db, perms.User.Id_usr, telefono)
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
			obj := GetTemplateConf("Crear Usuario", "Formulario", "Complete los campos", "Titulo Lista", "guardar_usuarios", fmt.Sprintf("/pages/%s", name), "borrar_usuario", "Usuario")

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
			obj := GetTemplateConf("Crear Cursos", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_cursos", fmt.Sprintf("/pages/%s", name), "borrar_curso", "Curso")

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
			obj := GetTemplateConf("Crear Empresa", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_empresa", fmt.Sprintf("/pages/%s", name), "borrar_empresa", "Empresa")

			t, err := TemplatePage(fmt.Sprintf("html/admin/%s.html", name))
			ErrorCheck(err)

			agenda, found := GetAgendaCurso(db, perms.User.Id_usr, fecha)
			if found {
				obj.Agenda = agenda
			}

			err = t.Execute(ctx, obj)
			ErrorCheck(err)
		}
	case "Prestamos":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {
			obj := GetTemplateConf("Crear Empresa", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_empresa", fmt.Sprintf("/pages/%s", name), "borrar_empresa", "Empresa")

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
			obj := GetTemplateConf("Crear Empresa", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_parentesco", fmt.Sprintf("/pages/%s", name), "borrar_parentesco1", "Empresa")

			if id > 0 {

				obj.FormId = id
				if id_apo > 0 {
					padre, found := GetPadreUser(db, id, id_apo)
					if found {
						obj.Padre = padre
						fmt.Println(padre.Id_usr)
					}
				}

				padres, found := GetPadresUser(db, id)
				if found {
					obj.Padres = padres
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
			obj := GetTemplateConf("Crear Empresa", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_parentesco2", fmt.Sprintf("/pages/%s", name), "borrar_parentesco2", "Empresa")

			if id > 0 {

				obj.FormId = id

				padres, found := GetAlumnosUser(db, id)
				if found {
					obj.Padres = padres
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
			obj := GetTemplateConf("Crear Empresa", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_edu_curso", fmt.Sprintf("/pages/%s", name), "borrar_cur_edu1", "Empresa")
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
			obj := GetTemplateConf("Crear Empresa", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_edu_curso2", fmt.Sprintf("/pages/%s", name), "borrar_cur_edu2", "Empresa")
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
	case "borrarUsuario":
		perms := GetPermisoUser(db, token, false)
		if perms.Permisos.Admin {

			id := Read_uint32bytes(ctx.QueryArgs().Peek("id"))
			obj := GetTemplateConf("Crear Empresa", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_edu_curso", fmt.Sprintf("/pages/%s", name), "borrar_user23", "Usuario")

			obj.FormId = id

			Usuario, found := GetUsuarioComplete(db, id)
			if found {
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
			obj := GetTemplateConf("Eliminar Curso", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_edu_curso", fmt.Sprintf("/pages/%s", name), "borrar_curso", "Curso")
			obj.FormId = id

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

			obj := GetTemplateConf("Crear Empresa", "Subtitulo", "Subtitulo2", "Titulo Lista", "guardar_cursos", fmt.Sprintf("/pages/%s", name), "borrar_empresa", "Empresa")
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

			genero := string(ctx.FormValue("genero"))
			reglamento := string(ctx.FormValue("reglamento"))
			fecha_nacimiento := string(ctx.FormValue("fecha_nacimiento"))
			fecha_matricula := string(ctx.FormValue("fecha_matricula"))
			fecha_ingreso := string(ctx.FormValue("fecha_ingreso"))

			direccion := string(ctx.FormValue("direccion"))
			fecha_retiro := string(ctx.FormValue("fecha_retiro"))
			motivo_retiro := string(ctx.FormValue("motivo_retiro"))
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
			if id > 0 {
				resp.Op, resp.Msg = UpdateCurso(db, id, nombre)
			}
			if id == 0 {
				resp.Op, resp.Msg = InsertCurso(db, nombre)
			}
			if resp.Op == 1 {
				resp.Page = "crearCursos"
				resp.Reload = 1
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
				resp.Op, resp.Msg = UpdateParentesco(db, id_apo, id, apoderado)
			} else {
				if id_usr == 0 {
					if !SelectParentesco(db, id_apo, id) {
						resp.Op, resp.Msg = InsertParentesco(db, id_usr, id, apoderado)
					} else {
						resp.Msg = "Padre ya esta asociado"
					}
				} else {
					resp.Msg = "Debe Seleccionar Padre"
				}
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
			fmt.Println("PRESTAR LIBRO", id_alu, id_lib)
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
					resp.Page = "verPadres?tipo=1"
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
					resp.Page = "verAlumnos?tipo=1"
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
					resp.Page = "verEducadoras?tipo=1"
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
					resp.Page = "verCursos?tipo=1"
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
		fmt.Println("USER NO ENCONTRADO")
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
func Index(ctx *fasthttp.RequestCtx) {

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	ctx.SetContentType("text/html; charset=utf-8")
	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), true)
	index.Login = Read_uint32bytes(ctx.FormValue("login"))
	index.Page = "Inicio"

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func Salir(ctx *fasthttp.RequestCtx) {

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

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	ctx.SetContentType("text/html; charset=utf-8")
	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), false)

	fmt.Println("Cookie:", string(ctx.Request.Header.Cookie("cu")))
	PrintJson(index)

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
func LibroInicio(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("text/html; charset=utf-8")
	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), false)
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

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func LibroPage(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("text/html; charset=utf-8")

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), false)
	index.Page = "Libros"

	code := fmt.Sprintf("%v", ctx.UserValue("name"))
	if len(code) == 32 {
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

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func AgendaPage(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("text/html; charset=utf-8")

	db, err := GetMySQLDB()
	defer db.Close()
	ErrorCheck(err)

	index := GetPermisoUser(db, string(ctx.Request.Header.Cookie("cu")), false)
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

	t, err := TemplatePages("html/web/index.html", "html/web/inicio.html", "html/web/libros.html", "html/web/agenda.html", "html/web/librobase.html")
	ErrorCheck(err)
	err = t.Execute(ctx, index)
	ErrorCheck(err)
}
func Qr(ctx *fasthttp.RequestCtx) {

	id, err := strconv.Atoi(fmt.Sprintf("%v", ctx.UserValue("name")))
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	urlqr := fmt.Sprintf("www.jardinvalleencantado.cl/libro/%v", randSeq(32, id))

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

	ctx.SetContentType("text/html; charset=utf-8")

	cant, err := strconv.Atoi(fmt.Sprintf("%v", ctx.UserValue("name")))
	if err != nil {
		fmt.Println("Error during conversion")
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
	//CREATE DATABASE redigo CHARACTER SET utf8 COLLATE utf8_spanish2_ci;
	db, err = sql.Open("mysql", "root:xFpsM6E1bda@tcp(127.0.0.1:3306)/jardin")
	//db, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/jardin")
	return
}

// USUARIOS DB //
func UpdateUsuario(db *sql.DB, tipo string, id int, nombre string, telefono string, correo string, nmatricula string, rut string, apellido1 string, apellido2 string, genero string, reglamento string, fecha_nacimiento string, fecha_matricula string, fecha_ingreso string, direccion string, fecha_retiro string, motivo_retiro string, observaciones string, id_cur int) (uint8, string) {
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
func InsertUsuario(db *sql.DB, tipo string, nombre string, telefono string, correo string, nmatricula string, rut string, apellido1 string, apellido2 string, genero string, reglamento string, fecha_nacimiento string, fecha_matricula string, fecha_ingreso string, direccion string, fecha_retiro string, motivo_retiro string, observaciones string, id_cur int) (uint8, string) {
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

	list := []int{}
	res, err := db.Query("SELECT id_cur FROM curso_usuarios WHERE id_usr = ?", id_usr)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
	}

	for res.Next() {
		var id_cur int
		err := res.Scan(&id_cur)
		if err != nil {
			ErrorCheck(err)
		}
		list = append(list, id_cur)
	}

	insert := true
	for _, x := range list {
		if x == id_cur {
			insert = false
		} else {
			return DeleteUserClass(db, id_usr, id_cur)
		}
	}
	if insert {
		return InsertUserClass(db, id_usr, id_cur)
	}
	return true
}
func DeleteUserClass(db *sql.DB, id_usr int, id_cur int) bool {
	delForm, err := db.Prepare("DELETE FROM curso_usuarios WHERE id_usr = ? AND id_cur = ?")
	if err != nil {
		ErrorCheck(err)
		return false
	}
	_, err = delForm.Exec(id_usr, id_cur)
	defer db.Close()
	if err != nil {
		ErrorCheck(err)
		return false
	}
	return true
}
func InsertUserClass(db *sql.DB, id_usr int, id_cur int) bool {

	stmt, err := db.Prepare("INSERT INTO curso_usuarios (id_usr, id_cur) VALUES (?,?)")
	if err != nil {
		ErrorCheck(err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(id_usr, id_cur)
	if err == nil {
		return true
	} else {
		ErrorCheck(err)
		return false
	}
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

		}

	} else {
		return Usuario, false
	}
	return Usuario, true
}
func GetUsuarios(db *sql.DB, tipo int) ([]Lista, bool) {

	Listas := []Lista{}

	cn := 0
	res, err := db.Query("SELECT id_usr, nombre, tipo FROM usuarios WHERE tipo = ? AND eliminado = ?", tipo, cn)
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
		return 1, "Nombre actualizada correctamente"
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
			return 1, "Correo actualizada correctamente"
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
		return 1, "Telefono actualizada correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Telefono no pudo ser actualizada"
	}
}
func UpdateCurso(db *sql.DB, id int, nombre string) (uint8, string) {
	stmt, err := db.Prepare("UPDATE cursos SET nombre = ? WHERE id_cur = ?")
	ErrorCheck(err)
	_, err = stmt.Exec(nombre, id)
	if err == nil {
		return 1, "Curso actualizada correctamente"
	} else {
		ErrorCheck(err)
		return 2, "El Curso no pudo ser actualizada"
	}
}
func InsertCurso(db *sql.DB, nombre string) (uint8, string) {
	stmt, err := db.Prepare("INSERT INTO cursos (nombre) VALUES (?)")
	ErrorCheck(err)
	defer stmt.Close()
	_, err = stmt.Exec(nombre)
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
	res, err := db.Query("SELECT id_cur, nombre FROM cursos WHERE id_cur = ? AND eliminado = ?", id, cn)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Curso, false
	}

	if res.Next() {
		err := res.Scan(&Curso.Id_cur, &Curso.Nombre)
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
	res, err := db.Query("SELECT id_cur, nombre FROM cursos WHERE eliminado = ?", cn)
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
		Prestamos = append(Prestamos, Prestamo)
	}
	return Prestamos, true
}
func GetLibroPrestadosUser(db *sql.DB, id int) ([]Prestamo, bool) {

	Prestamos := []Prestamo{}
	Prestamo := Prestamo{}

	//cn := 0
	res, err := db.Query("SELECT t2.id_pre, t2.fecha_prestamo, t3.nombre FROM parentensco t1, prestamos t2, libros t3 WHERE t1.id_apo = ? AND t1.id_usr=t2.id_usr AND t2.fecha_devolucion = '0000-00-00 00:00:00' AND t2.id_lib=t3.id_lib")
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return Prestamos, false
	}

	for res.Next() {

		err := res.Scan(&Prestamo.Id, &Prestamo.Fecha_Prestamos, &Prestamo.Nombre)
		if err != nil {
			ErrorCheck(err)
			return Prestamos, false
		}
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

		fmt.Println("Fecha:", Agenda.Fecha)

		//cn := 0
		res, err := db.Query("SELECT t3.id_usr, t1.id_cur, t3.nombre, t4.nombre FROM educadora_curso t1, curso_usuarios t2, usuarios t3, cursos t4 WHERE t1.id_usr = ? AND t1.id_cur=t2.id_cur AND t2.id_usr=t3.id_usr AND t1.id_cur=t4.id_cur", id)
		defer res.Close()
		if err != nil {
			ErrorCheck(err)
			return Agenda, false
		}

		var id_usr int
		var id_cur int
		var nombre_usr string
		var nombre_cur string

		for res.Next() {
			err := res.Scan(&id_usr, &id_cur, &nombre_usr, &nombre_cur)
			if err != nil {
				ErrorCheck(err)
				return Agenda, false
			}

			if _, Found := Agenda.AgendaCurso.Cursos[id_cur]; Found {
				Agenda.AgendaCurso.Cursos[id_cur].Users[id_usr] = GetAgendaUser(db, id_usr, aux.Fecha, nombre_usr)
			} else {
				Agenda.AgendaCurso.Cursos[id_cur] = Curso{Nombre: nombre_cur, Users: make(map[int]AgendaUser, 0)}
				Agenda.AgendaCurso.Cursos[id_cur].Users[id_usr] = GetAgendaUser(db, id_usr, aux.Fecha, nombre_usr)
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
	res, err := db.Query("SELECT ali1, ali2, ali3, dep1, dep2, comentario, ultima_actualizacion FROM agenda WHERE id_usr = ? AND fecha = ?", id, fecha)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return AgendaUser
	}

	if res.Next() {

		var ultimaact string

		err := res.Scan(&AgendaUser.Ali1, &AgendaUser.Ali2, &AgendaUser.Ali3, &AgendaUser.Dep1, &AgendaUser.Dep2, &AgendaUser.Comentario, &ultimaact)
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
	res, err := db.Query("SELECT t2.id_usr, t2.nombre FROM curso_usuarios t1, usuarios t2 WHERE t1.id_cur = ? AND t1.id_usr=t2.id_usr", id)
	defer res.Close()
	if err != nil {
		ErrorCheck(err)
		return lista, false
	}

	for res.Next() {
		list := Lista{}
		err := res.Scan(&list.Id, &list.Nombre)
		if err != nil {
			ErrorCheck(err)
			return lista, false
		}
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
func BorrarCursoEdu(db *sql.DB, id1 string, id2 string) (string, string, string) {

	delForm, err := db.Prepare("DELETE FROM curso_usuarios WHERE id_cur = ? AND id_usr = ?")
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
func TemplatePages(v1 string, v2 string, v3 string, v4 string, v5 string) (*template.Template, error) {

	t, err := template.ParseFiles(v1, v2, v3, v4, v5)
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
		res, err := db.Query("SELECT t1.id_usr, t1.tipo, t1.nombre, t1.correo, t1.telefono, t1.telefono2 FROM usuarios t1, sesiones t2 WHERE t2.id_ses = ? AND t2.cookie = ? AND t2.id_usr=t1.id_usr AND t1.eliminado = ?", Read_uint32bytes([]byte(tkn[33:])), tkn[0:32], cn)
		defer res.Close()
		ErrorCheck(err)

		if res.Next() {
			Index.Register = true
			err := res.Scan(&Index.User.Id_usr, &Index.User.Tipo, &Index.User.Nombre, &Index.User.Correo, &Index.User.Telefono, &Index.User.Telefono2)
			ErrorCheck(err)

			if Index.User.Tipo == 0 {
				// ADMINISTRADOR
				Index.Permisos.Admin = true
				if complete {
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Administrador", Icon: "admin", Url: "admin"})
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Resumen Libro", Icon: "libro", Url: "libros"})
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Agenda Virtual", Icon: "agenda", Url: "agenda"})
				}
			} else if Index.User.Tipo == 1 {
				// PERSONAL
				Index.Permisos.Educadora = true
				if complete {
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Resumen Libro", Icon: "libro", Url: "libros"})
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Agenda Virtual", Icon: "agenda", Url: "agenda"})
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
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Resumen Libro", Icon: "libro", Url: "libros", Text: textLibro})
					Index.Modulos = append(Index.Modulos, Modulo{Titulo: "Agenda Virtual", Icon: "agenda", Url: "agenda", Text: textAgenda})
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
