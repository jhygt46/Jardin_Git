var lista_alu = [];
var prestar_id_alu = 0;
var id_lib = 0;
var fecha_agenda = {};
var code = "";
var open_menu = false;
var list_nav_visita = [0];

function icon(i, w, l){
    GC(i, 0).style.width = w+"px";
    GC(i, 0).style.left = l+"px";
}
function sizeWeb(){

    start_visita();
    if (GC("a3", 0) !== undefined){

        var width = window.innerWidth;
        var cont_site = width * 0.98 > 800 ? 800 : width * 0.98;
        var btns = GC("a3", 0).children;
        var margin = 20;
        var btn_max_width = 142;
        
        var btn_width = (cont_site - margin * 2) / btns.length > btn_max_width ? btn_max_width : (cont_site - margin * 2) / btns.length ;
        var left = (cont_site - btn_width * btns.length) / 2;

        var i = 0;
        for (x of btns){
            x.style.width = btn_width+"px";
            x.style.left = left+"px";
            left += btn_width;
            i++;
        }

        GC("a3", 0).style.height = btn_width+"px";
    }

}

document.addEventListener('DOMContentLoaded', function() {

    sizeWeb();
    window.addEventListener("resize", (event) => {
        sizeWeb();
    });

    var puntos = GCS("punto");
    for(var i=0; i<puntos.length; i++){
        puntos[i].addEventListener("click", ver_puntos);
    }

    var pan = GCS("pan");
    for(var i=0; i<pan.length; i++){
        pan[i].addEventListener("mousemove", mouse_move);
        pan[i].addEventListener("touchmove", touchmove);
        pan[i].addEventListener("touchstart", touchstart);
    }

    var vvimg = GCS("vvimage_chica");
    for(var i=0; i<vvimg.length; i++){
        vvimg[i].addEventListener("click", vvimgclick);
    }
    
    GC("user", 0).addEventListener("click", show_login);
    GC("close", 0).addEventListener("click", hide_login);
    GC("button", 0).addEventListener("click", toogle_menu);
    if (GC("delete_prestamo", 0) !== undefined){
        GC("delete_prestamo", 0).addEventListener("click", delete_prestamo);
    }
    if (GC("back", 0) !== undefined){
        GC("back", 0).addEventListener("click", back);
    }
    if (GI("devolver_libro") !== null){
        GI("devolver_libro").addEventListener("click", devolver_libro);
    }
    if (GI("prestar_libro") !== null){
        GI("prestar_libro").addEventListener("click", prestar_libro);
    }
    if (GI("guardar_libro") !== null){
        GI("guardar_libro").addEventListener("click", guardar_libro);
    }
    if (GI("enviar") !== null){
        GI("enviar").addEventListener("click", enviar);
    }
    if(GC("moddatauser", 0) !== undefined){
        GC("moddatauser", 0).addEventListener("click", showmoddatauser);
    }
    if(GC("loginolvido", 0) !== undefined){
        GC("loginolvido", 0).addEventListener("click", loginolvido);
    }
    if(GC("loginback", 0) !== undefined){
        GC("loginback", 0).addEventListener("click", loginback);
    }
    if(GC("salir", 0) !== undefined){
        GC("salir", 0).addEventListener("click", salir);
    }
    if (GI("mod_datos") !== null){
        GI("mod_datos").addEventListener("click", mod_datos);
    }
    if (GI("reestablecer") !== null){
        GI("reestablecer").addEventListener("click", reestablecer);
    }
    if (GI("login") !== null){
        GI("login").addEventListener("click", login);
    }
    if (GI("loginrec") !== null){
        GI("loginrec").addEventListener("click", loginrec);
    }
    if (GI("alu_nom") !== null){
        GI("alu_nom").addEventListener("keyup", buscaralumno);
    }
    for (x of GCS("bp")){
        x.addEventListener("click", a3);
    }
    for (x of GCS("changeuser")){
        x.addEventListener("blur", change_user);
    }
}, false);

function vvimgclick(){

    var id = this.getAttribute("id");
    var imglist = this.parentElement.parentElement.children[0].children;
    for(var i=0; i<imglist.length; i++){
        if(i == id){
            imglist[i].className = "vvimage_grande visible";
        }else{
            imglist[i].className = "vvimage_grande";
        }
    }
}
function loginolvido(event){
    GC("form1", 0).style.display = "none";
    GC("form2", 0).style.display = "block";
    event.preventDefault();
}
function showmoddatauser(event){
    var form = GC("form", 0).children;
    form[0].style.display = "none";
    form[1].style.display = "block";
    event.preventDefault();
}
function loginback(event){
    GC("form1", 0).style.display = "block";
    GC("form2", 0).style.display = "none";
    event.preventDefault();
}
function salir(event){
    if (confirm("¿Realmente desea salir?") == true) {
        window.location.href = "/salir";
    }
    event.preventDefault();
}
function ver_puntos(){
    var id = this.getAttribute("id");
    list_nav_visita.push(id);
    GC("back", 0).style.display = "block";
    ShowVisita(id);
}
function back(){
    list_nav_visita.pop();
    var id = list_nav_visita[list_nav_visita.length-1];
    if(id == 0){
        GC("back", 0).style.display = "none";
    }
    ShowVisita(id);
}
function ShowVisita(id){
    var list = GC("visita-virtual", 0).children;
    for(var i=0; i<list.length - 1; i++){
        if(i == id){
            list[i].className = AddVisible(list[i].className, true);
        }else{
            list[i].className = AddVisible(list[i].className, false);
        }
    }
}
function AddVisible(clase, visible){
    var list = clase.split(" ");
    if(visible){
        return list[0]+" visible";
    }else{
        return list[0];
    }
}
function enviar(){
    const formData = new FormData();
    formData.append('accion', 'enviar_contacto');
    formData.append('nombre', GI("contacto_nombre").value);
    formData.append('correo', GI("contacto_correo").value);
    formData.append('telefono', GI("contacto_telefono").value);
    formData.append('mensaje', GI("contacto_mensaje").value);

    postData2("/save", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Mensaje Enviado", function(){
                window.location.reload();
            });
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });
}
function guardar_libro(){

    const formData = new FormData();
    formData.append('accion', 'guardar_libro');
    formData.append('code', code);
    formData.append('nombre', GI("libro_nom").value);

    postData2("/save", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Cambio Exitoso", function(){
                window.location.reload();
            });
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });
}
function prestar_libro(){
    
    const formData = new FormData();
    formData.append('accion', 'prestar_libro');
    formData.append('id_alu', prestar_id_alu);
    formData.append('id_lib', id_lib);

    postData2("/save", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Cambio Exitoso", function(){
                window.location.reload();
            });
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });
}
function devolver_libro(){
    
    const formData = new FormData();
    formData.append('accion', 'devolver_libro');
    formData.append('id_lib', id_lib);

    postData2("/save", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Cambio Exitoso", function(){
                window.location.reload();
            });
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });
}
function handleChange1(that, tipo, id){

    if (tipo == 6){
        var ids = that.parentElement.id;
        if(ids == "0"){
            sendAgenda(id, 1, tipo);
            that.parentElement.id = "1";
            that.parentElement.className = "alumno_nombre alumno_nombre_color";
            that.parentElement.parentElement.children[1].style.display = "none";
        }
        if(ids == "1"){
            sendAgenda(id, 0, tipo);
            that.parentElement.id = "0";
            that.parentElement.className = "alumno_nombre";
        }
    }else{
        var x = that.parentElement.parentElement.children;
        if(that.className == "selected"){
            that.checked = false;
            that.className = "";
            sendAgenda(id, 0, tipo);
        }else{
            that.className = "selected";
            sendAgenda(id, that.value, tipo);
            for(var i=1; i<x.length; i++){
                if (i != that.value){
                    x[i].children[0].className = "";
                }
            }
        }
    }

}
function sendcommentagenda(that, id){
    sendAgenda(id, that.value, 7);
}
function sendAgenda(id, val, tipo){

    const formData = new FormData();
    formData.append('accion', 'set_agenda');
    formData.append('id_alu', id);
    formData.append('tipo', tipo);
    formData.append('value', val);

    postData2("/accion", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Cambio Exitoso", function(){});
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });
}
function search_agenda(that, n){

    if (that.getAttribute("id") != "no"){

        var fecha = that.parentElement.children;
        var cont_users = that.parentElement.parentElement.children[1];

        console.log(cont_users);

        const formData = new FormData();
        formData.append('accion', 'get_agenda');

        if (n == 1){
            formData.append('fecha', fecha_agenda.fecha_next);
        } else {
            formData.append('fecha', fecha_agenda.fecha_prev);
        }

        postData2("/accion", formData).then((resp) => {
            if (resp.Op == 1){
                mensaje(1, "Datos Actualizados", function(){});

                console.log(resp.Agenda);

                fecha[1].innerHTML = resp.Agenda.FechaStr;
                fecha_agenda.fecha = resp.Agenda.Fecha;
                fecha_agenda.fecha_prev = resp.Agenda.FechaPrev;
                fecha_agenda.fecha_next = resp.Agenda.FechaNext;

                if (resp.Agenda.FechaIsNext == 1){
                    fecha[2].setAttribute("id", "");
                    fecha[2].className = "fecha_next icon";
                }else{
                    fecha[2].setAttribute("id", "no");
                    fecha[2].className = "fecha_next icon opacity";
                }
                
                if(resp.Agenda.Users !== null){
                    cont_users.innerHTML = "";
                    cont_users.appendChild(GenerarAgendaUsers(resp.Agenda.Users));
                }

            } else if (resp.Op == 2){
                mensaje(2, resp.Msg, function(){});
            } else {
                mensaje(3, resp.Msg, function(){});
            }
        });
    }
}
function GenerarAgendaUsers(users){
    var childElement = document.createElement("div");
    if (users.length == 1){
        childElement.className = "cont_agenda ca1";
    }else{
        childElement.className = "cont_agenda ca2";
    }
    for (var i=0; i<users.length; i++){
        var apo_alu = document.createElement("div");
        apo_alu.className = "apo_alu";

        var apo_alu_nom = document.createElement("div");
        apo_alu_nom.className = "apo_alu_nom";
        apo_alu_nom.innerHTML = users[i].Nombre;

        var apo_alu_det = document.createElement("div");
        apo_alu_det.className = "apo_alu_det";

        if(users[i].Data){

            if(users[i].Ausente == 0){

                var cont_alu_det = document.createElement("div");
                cont_alu_det.className = "cont_alu_det clearfix";
                var cont_alu_det1 = document.createElement("div");
                cont_alu_det1.className = "alu_info col1 b1";
                cont_alu_det1.innerHTML = "Colación";
                cont_alu_det.appendChild(cont_alu_det1);
                var cont_alu_det2 = document.createElement("div");
                cont_alu_det2.className = "alu_info col1 b1";
                cont_alu_det2.innerHTML = "Almuerzo";
                cont_alu_det.appendChild(cont_alu_det2);
                var cont_alu_det3 = document.createElement("div");
                cont_alu_det3.className = "alu_info col1 b1";
                cont_alu_det3.innerHTML = "Once";
                cont_alu_det.appendChild(cont_alu_det3);
                apo_alu_det.appendChild(cont_alu_det);

                var cont_alu_det = document.createElement("div");
                cont_alu_det.className = "cont_alu_det clearfix";
                var cont_alu_det1 = document.createElement("div");
                cont_alu_det1.className = "alu_info col1 b2";
                cont_alu_det1.innerHTML = GetName(users[i].Ali1)
                cont_alu_det.appendChild(cont_alu_det1);
                var cont_alu_det2 = document.createElement("div");
                cont_alu_det2.className = "alu_info col1 b2";
                cont_alu_det2.innerHTML = GetName(users[i].Ali2)
                cont_alu_det.appendChild(cont_alu_det2);
                var cont_alu_det3 = document.createElement("div");
                cont_alu_det3.className = "alu_info col1 b2";
                cont_alu_det3.innerHTML = GetName(users[i].Ali3)
                cont_alu_det.appendChild(cont_alu_det3);
                apo_alu_det.appendChild(cont_alu_det);

                var cont_alu_det = document.createElement("div");
                cont_alu_det.className = "cont_alu_det clearfix mtop";
                var cont_alu_det1 = document.createElement("div");
                cont_alu_det1.className = "alu_info col2 b1";
                cont_alu_det1.innerHTML = "Orina";
                cont_alu_det.appendChild(cont_alu_det1);
                var cont_alu_det2 = document.createElement("div");
                cont_alu_det2.className = "alu_info col2 b1";
                cont_alu_det2.innerHTML = "Deposición";
                cont_alu_det.appendChild(cont_alu_det2);
                apo_alu_det.appendChild(cont_alu_det);

                var cont_alu_det = document.createElement("div");
                cont_alu_det.className = "cont_alu_det clearfix";
                var cont_alu_det1 = document.createElement("div");
                cont_alu_det1.className = "alu_info col2 b2";
                cont_alu_det1.innerHTML = users[i].Dep1;
                cont_alu_det.appendChild(cont_alu_det1);
                var cont_alu_det2 = document.createElement("div");
                cont_alu_det2.className = "alu_info col2 b2";
                cont_alu_det2.innerHTML = users[i].Dep2;
                cont_alu_det.appendChild(cont_alu_det2);
                apo_alu_det.appendChild(cont_alu_det);


                if(users[i].Comentario != ""){

                    var comentario1 = document.createElement("div");
                    comentario1.className = "cont_alu_det clearfix mtop";
                    var cont_comentario1 = document.createElement("div");
                    cont_comentario1.className = "alu_info col3a b1";
                    cont_comentario1.innerHTML = "Comentario";
                    comentario1.appendChild(cont_comentario1);

                    var comentario2 = document.createElement("div");
                    comentario2.className = "cont_alu_det clearfix";
                    var cont_comentario2 = document.createElement("div");
                    cont_comentario2.className = "alu_info col3b b2";
                    cont_comentario2.innerHTML = users[i].Comentario;
                    comentario2.appendChild(cont_comentario2);

                    apo_alu_det.appendChild(comentario1);
                    apo_alu_det.appendChild(comentario2);

                }

            }else{
                
                var nodata = document.createElement("div");
                nodata.className = "nodata vhalign";
                nodata.innerHTML = "Hoy se registro como ausente";
                apo_alu_det.appendChild(nodata);

            }
            
        }else{

            var nodata = document.createElement("div");
            nodata.className = "nodata vhalign";
            nodata.innerHTML = "Hoy no se registraron datos";
            apo_alu_det.appendChild(nodata);

        }

        
        apo_alu.appendChild(apo_alu_nom);
        apo_alu.appendChild(apo_alu_det);

        childElement.appendChild(apo_alu);
    }
    return childElement;
}
function GetName(n){
    if (n == 0){
        return "No";
    }
    if (n == 1){
        return "Nada";
    }
    if (n == 2){
        return "Mitad";
    }
    if (n == 3){
        return "Todo";
    }
}
function agenda_alumno(that){

    if (that.parentElement.parentElement.children[1].style.display == "block"){
        that.parentElement.parentElement.children[1].style.display = "none";
    }else{
        if(that.parentElement.id == 0){
            for (x of that.parentElement.parentElement.parentElement.children){
                x.children[1].style.display = "none";
            }
            that.parentElement.parentElement.children[1].style.display = "block";
        }
    }
}
function agenda_curso(that){
    if (that.parentElement.children[1].style.display == "block"){
        that.parentElement.children[1].style.display = "none";
    }else{
        for (x of that.parentElement.parentElement.children){
            x.children[1].style.display = "none";
            for (y of x.children[1].children){
                y.children[1].style.display = "none";
            }
        }
        that.parentElement.children[1].style.display = "block";
    }    
}
function buscaralumno(){
    var lista = GC("lista_busqueda", 0);
    lista.innerHTML = "";
    for (x of lista_alu){
        if (x.Nombre.toUpperCase().indexOf(this.value.toUpperCase()) !== -1){
            lista.appendChild(opc_alu(x));
        }
    }
}
function select_user_libro(){
    prestar_id_alu = this.getAttribute("id");
    GC("selectalum", 0).style.display = "none";
    GC("selectedalum", 0).children[1].innerHTML = this.innerHTML;
    GC("selectedalum", 0).children[1].style.display = "block";
    GC("selectedalum", 0).style.display = "block";
}
function delete_prestamo(){
    prestar_id_alu = 0;
    GC("selectalum", 0).style.display = "block";
    GC("selectedalum", 0).style.display = "none";
}
function opc_alu(x){
    var childElement = document.createElement("div");
    childElement.innerHTML = x.Nombre+" "+x.Apellido1+" "+x.Apellido2;
    childElement.className = "alu";
    childElement.id = x.Id;
    childElement.addEventListener("click", select_user_libro);
    return childElement;
}
function mod_datos(){

    var nom = GI("mod_nombre").value;
    var mail = GI("mod_mail").value;
    var fono = GI("mod_telefono").value;
    
    const formData = new FormData();
    formData.append('accion', 'mod_datos');
    formData.append('nom', nom);
    formData.append('mail', mail);
    formData.append('fono', fono);

    postData2("/accion", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Cambio Exitoso", function(){});
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });

}
function a3(event){
    event.preventDefault();
    SelectPage(this.getAttribute("href"), true);
    
}
function SelectPage(h, p){

    var btns = GC("a3", 0).children;
    var n = 999;
    for(var i=0; i<btns.length; i++){
        if(btns[i].getAttribute("href") == h){
            n = i;
        }
    }

    if(btns.length > n){
        if(!btns[n].classList.contains('selected')){
            var href = btns[n].getAttribute("href");
            for(x of btns){
                if(x.classList.contains('selected')){
                    GC(x.getAttribute("href").substring(1), 0).classList.remove("AnimateTopclass");
                    GC(x.getAttribute("href").substring(1), 0).classList.add("AnimateBottomclass");
                    x.classList.remove("selected");
                }
            }
            GC(href.substring(1), 0).classList.remove("AnimateBottomclass");
            GC(href.substring(1), 0).classList.add("AnimateTopclass");
            btns[n].classList.add("selected");
            if(p){
                history.pushState(null, GetTitle(href), href);
            }
        }
    }
    
}
function GetTitle(s){
    return s.substring(1);
}

window.addEventListener('popstate', historyfunc);
function historyfunc(){
    console.log("BACK: ", document.location.pathname);
    SelectPage(document.location.pathname, false);
}

var btn_active = 1;
function aparece(pag){
    if(btn_active == 1){
        btn_active = 0;
        //desaparece();

        GC(pag, 0).style.top = "100px";

        obj = { iter: 50, els: [{el: GC(pag, 0), arr: [{pro: "right", tipo: 0, hasta: 0}]}] }
        animateJs(obj, function(){ btn_active = 1; });
        /*
        $("."+pag).animate({
            top: "0px",
            opacity: 1
        }, 1000, function() {
            pagina = pag;
            btn_active = 1;
        });
        */
    }
}
function desaparece(){
    var pag = pagina;
    $("."+pag).animate({
        top: "100px",
        opacity: 0
    }, 1000, function() {
        $("."+pag ).css({top: "500px"});
    });
}
function animateJs(obj, callback){

    if (!obj.hasOwnProperty("times")){
        obj.times = 0
    }else{
        obj.times++
    }
    if (obj.times == 0) {
        for (let value of obj.els) {
            if (value.el.getAttribute("id") == "move"){
                return;
            }else{
                value.el.setAttribute("id", "move");
                for (let prop of value.arr) {
                    if (value.el.style[prop.pro] == "") {
                        var style = window.getComputedStyle(value.el);
                        prop.desde = parseInt(style[prop.pro]);
                    }else{
                        prop.desde = parseInt(value.el.style[prop.pro].replace("px", ""));
                    }
                }
            }
        }
    }
    if (obj.times == obj.iter) {
        for (let value of obj.els) {
            value.el.removeAttribute('id');
        }
        callback(obj);
        return;
    }
    for (let value of obj.els) {
        for (let prop of value.arr) {
            value.el.style[prop.pro] = calcvalue(prop.desde, prop.hasta, obj.times+1, obj.iter, prop.tipo || value.tipo)+"px";
        }
    }
    setTimeout(function(){
        animateJs(obj, callback);
    }, 10)
}
function calcvalue(desde, hasta, i, total, tipo){
    if (tipo == 0) {
        return desde + i/total * (hasta - desde);
    }else if (tipo == 1){
        return desde + easeInOutSine(i/total) * (hasta - desde);
    }else{
        return desde + easeInOutSine(i/total) * (hasta - desde);
    }
}
function show_login(){
    GC("login", 0).style.display = "block";
}
function hide_login(){
    GC("login", 0).style.display = "none";
}
function toogle_menu(){
    obj = {
        iter: 50,
        els: [
            {el: GC("button", 0), arr: [{pro: "left", tipo: 0}]},
            {el: GC("menu", 0), arr: [{pro: "left", tipo: 1}]}
        ]
    }
    if (!open_menu) {
        obj.els[0].arr[0].hasta = 170;
        obj.els[1].arr[0].hasta = 0;
    }else{
        obj.els[0].arr[0].hasta = 250;
        obj.els[1].arr[0].hasta = -230;
    }
    animateJs(obj, function(x){
        open_menu = !open_menu
    });
}
function GC(name, pos){
    return document.getElementsByClassName(name)[pos];
}
function GCS(name){
    return document.getElementsByClassName(name);
}
function GI(name){
    return document.getElementById(name);
}
function easeInOutSine(x) {
    return x*x*x;
}
async function postData2(url, formData) {
    const response = await fetch(url, {
        method: 'POST',
        body: formData
    });
    return response.json();
}
function loginrec(){

    const formData = new FormData();
    formData.append('user', GI("correo_rec").value);

    postData2("/recuperar", formData).then((resp) => {

        console.log("RESP", resp);

        if (resp.Op == 1){
            mensaje(1, "Correo Enviado", function(){ location.reload(true); });
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });

}
function login(){

    const formData = new FormData();
    formData.append('user', GI("correo").value);
    formData.append('pass', GI("password").value);

    postData2("/login", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Ingreso Exitoso", function(){ location.reload(true); });
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });
    
}
function reestablecer(){
    const formData = new FormData();
    formData.append('accion', 'reestablecer');
    formData.append('pass1', GI("pass1").value);
    formData.append('pass2', GI("pass2").value);
    formData.append('id_usr', GI("id_usr").value);
    formData.append('code', GI("code").value);

    postData2("/accion", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Contraseña Reestablecida", function(){ window.location.href = "/?login=1"; });
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });
}
function mensaje(op, message, func){
    var el = GC("mensaje", 0);

    if (op == 1) {
        el.className = "mensaje msgok";
    } else if (op == 2) {
        el.className = "mensaje msgerr";
    } else {
        el.className = "mensaje msgwar";
    }

    msg = el.children[0]
    msg.innerHTML = message;
    obj = { iter: 10, els: [{el: el, arr: [{pro: "bottom", hasta: 0, tipo: 0}]}]}
    animateJs(obj, function(){
        setTimeout(function() {
            obj = { iter: 10, els: [{el: el, arr: [{pro: "bottom", hasta: -30, tipo: 0}]}]}
            animateJs(obj, func)
        }, 3000);
    });

}
function initMap(){

    var myCenter = new google.maps.LatLng(-33.480455,-70.5534333);
    var mapProp = {
        center: myCenter,
        zoom: 15,
        mapTypeControl: false,
        fullscreenControl: false,
        streetViewControl: false,
        zoomControl: false,
        keyboardShortcuts: false,
        mapTypeId:google.maps.MapTypeId.ROADMAP
    };
    var map = new google.maps.Map(document.getElementById("map"), mapProp);
    var marker = new google.maps.Marker({
        position: myCenter,
    });
    marker.setMap(map);

    const infowindow = new google.maps.InfoWindow({
        content: "<strong>Alberto Valenzuela Llanos 2705</strong>",
        ariaLabel: "Uluru",
    });
    infowindow.open({
        anchor: marker,
        map,
    });

    
}
function start_visita(){

    if(GC("visita-virtual", 0) !== undefined){
        var width = window.innerWidth;
        var cont_site = width * 0.98 > 800 ? 800 : width * 0.98;
        var images = GC("visita-virtual", 0).children;

        for(var i=0; i<images.length; i++){
            if(images[i].className == "pan" || images[i].className == "pan visible"){
                var left = (cont_site - images[i].getAttribute("id")) / 2
                images[i].style.left = left+"px";
            }
        }
    }
}
function mouse_move(e){
    var div = GC("visible", 0);
    if(div.className == "pan visible"){
        var width = window.innerWidth;
        var cont_site = width * 0.98 > 800 ? 800 : width * 0.98;
        var j = (e.clientX - (width - cont_site)/2)/cont_site;
        var left = (cont_site - div.offsetWidth)*j;
        div.style.left = left+"px";
    }
}
var touchX = 0;
function touchmove(e){
    
    var move = touchX - e.touches[0].pageX;
    touchX = e.touches[0].pageX;

    var width = window.innerWidth;
    var cont_site = width * 0.98 > 800 ? 800 : width * 0.98;
    var left_max = cont_site - this.id;

    var left = parseInt(this.style.left) - move;

    if(left <= 0 && left > left_max){
        this.style.left = left+"px";
    }
    e.preventDefault();

}
function touchstart(e){
    touchX = e.touches[0].pageX;
}