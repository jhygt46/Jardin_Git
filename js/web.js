var lista_alu = [];
var prestar_id_alu = 0;
var id_lib = 0;
var fecha_agenda = {};
var code = "";
document.addEventListener('DOMContentLoaded', function() {
    GC("user", 0).addEventListener("click", show_login);
    GC("close", 0).addEventListener("click", hide_login);
    GC("button", 0).addEventListener("click", toogle_menu);

    if (GC("delete_prestamo", 0) !== undefined){
        GC("delete_prestamo", 0).addEventListener("click", delete_prestamo);
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
    
    if (GI("login") !== null){
        GI("login").addEventListener("click", login);
    }
    if (GI("alu_nom") !== null){
        GI("alu_nom").addEventListener("keyup", buscaralumno);
    }
    for (x of GCS("d4")){
        x.addEventListener("click", d4);
    }
    for (x of GCS("changeuser")){
        x.addEventListener("blur", change_user);
    }
}, false);

function guardar_libro(){

    const formData = new FormData();
    formData.append('accion', 'guardar_libro');
    formData.append('code', code);
    formData.append('nombre', GI("libro_nom").value);

    postData2("http://localhost:81/save", formData).then((resp) => {
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

    postData2("http://localhost:81/save", formData).then((resp) => {
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

    postData2("http://localhost:81/save", formData).then((resp) => {
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
function sendcommentagenda(that, id){
    sendAgenda(id, that.value, 6);
}
function sendAgenda(id, val, tipo){

    const formData = new FormData();
    formData.append('accion', 'set_agenda');
    formData.append('id_alu', id);
    formData.append('tipo', tipo);
    formData.append('value', val);

    postData2("http://localhost:81/accion", formData).then((resp) => {
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

        const formData = new FormData();
        formData.append('accion', 'get_agenda');

        if (n == 1){
            formData.append('fecha', fecha_agenda.fecha_next);
        } else {
            formData.append('fecha', fecha_agenda.fecha_prev);
        }

        postData2("http://localhost:81/accion", formData).then((resp) => {
            if (resp.Op == 1){
                mensaje(1, "Datos Actualizados", function(){});

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
                
                cont_users.innerHTML = "";
                cont_users.appendChild(GenerarAgendaUsers(resp.Agenda.Users));

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

    if (that.parentElement.children[1].style.display == "block"){
        that.parentElement.children[1].style.display = "none";
    }else{
        for (x of that.parentElement.parentElement.children){
            x.children[1].style.display = "none";
        }
        that.parentElement.children[1].style.display = "block";
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
    childElement.innerHTML = x.Nombre;
    childElement.className = "alu";
    childElement.id = x.Id;
    childElement.addEventListener("click", select_user_libro);
    return childElement;
}
function closeoptions(data){
    data[1].style.display = "block";
    data[2].style.display = "none";
    data[3].style.display = "block";
}
function change_user(){

    var el = this.parentElement.parentElement.children;
    var id = this.getAttribute("id");
    var value = this.value;
    
    const formData = new FormData();
    formData.append('accion', id);
    formData.append('value', value);

    postData2("http://localhost:81/accion", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Cambio Exitoso", function(){});
            el[1].innerHTML = value;
            closeoptions(el);
        } else if (resp.Op == 2){
            mensaje(2, resp.Msg, function(){});
        } else {
            mensaje(3, resp.Msg, function(){});
        }
    });

}
function d4(){
    var el = this.parentElement.children;
    el[1].style.display = "none";
    el[2].style.display = "block";
    el[3].style.display = "none";
}

var open_menu = false;
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
    for (x of GCS("data")){
        closeoptions(x.children);
    }
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
function login(){

    const formData = new FormData();
    formData.append('user', GI("correo").value);
    formData.append('pass', GI("password").value);

    postData2("http://localhost:81/login", formData).then((resp) => {
        if (resp.Op == 1){
            mensaje(1, "Ingreso Exitoso", function(){ location.reload(true); });
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