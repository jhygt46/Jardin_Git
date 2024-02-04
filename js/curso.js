var player;
var ytvideo = false;
var turn = false;
function getImages(x){
    var res = "";
    var lista = JSON.parse(x.ListaImagen);
    for (var i=0; i<lista.length; i++){
        res += "<div class='page'><img src='/images_cuentos/"+lista[i].Nom+"' draggable='false' alt='' /></div>";
    }
    return res;
}
function getDims(w1, h1, w2, h2){

    var res = { pt: 0, pl: 0, w: 0, h: 0 }
    x = w2 * h1 / w1
    if (x > h2){
        res.w = w1 * h2 / h1;
        res.pl = (w2 - res.w) / 2;
        res.h = h2;
    }else{
        res.h = x;
        res.pt = (h2 - res.h) / 2;
        res.w = w2;
    }
    return res;
}
function showContent(){

    hide_content();
    var item = items[this.id];
    var cl = GC('cl', 0);
    cl.style.paddingTop = "0px";
    cl.style.paddingLeft = "0px";
    if(ytvideo){ player.stopVideo(); }

    if (item.Tipo == 1){
        var flipbookEL = GI('flipbook');
        if(turn){ $("#flipbook").turn('destroy'); }
        res = getDims(item.Dx*2, item.Dy, cl.clientWidth, cl.clientHeight);
        cl.style.paddingTop = res.pt+"px";
        cl.style.paddingLeft = res.pl+"px";
        flipbookEL.innerHTML = getImages(item);
        $("#flipbook").turn({
            width: res.w,
            height: res.h,
            elevation: 50,
            gradients: true,
            autoCenter: true
        });
        flipbookEL.style.display = "block";
        turn = true;
    }
    if (item.Tipo == 2){
        var video = GI('video');
        video.setAttribute("src", '/videos/'+item.Urlext);
        video.autoplay = true; 
        video.load();
        video.play();
        video.style.display = "block";
    }
    if (item.Tipo == 3){
        var youtube = GI('player');
        player.loadVideoById(item.Urlext);
        player.playVideo();
        youtube.style.display = "block";
        ytvideo = true;
    }
    if (item.Tipo == 4){
        var iframe = GI('iframe');
        iframe.src = item.Urlext;
        iframe.style.display = "block";
    }
    toogle_lista_curso();
}
function showMenuCursos(id, url){

    var lista = GC("lista", 0);
    lista.innerHTML = "";

    for(var i=0; i<cursos.length; i++){
        if(cursos[i].Id == id){
            for(var j=0; j<cursos[i].CursosItems.length; j++){
                lista.append(itemelement(cursos[i].CursosItems[j].Id))
            }
        }
    }
    toogle_lista_curso();
}
function itemelement(x){
    //console.log(x, items[x]);
    const newDiv = document.createElement("div");
    newDiv.className = "items";
    newDiv.onclick = showContent;
    newDiv.setAttribute("id", x);

    const DivImage = document.createElement("div");
    DivImage.className = "itemsimage";
    const Image = document.createElement("img");
    Image.src = "/images_preview/"+items[x].Image;
    DivImage.appendChild(Image)

    const DivText = document.createElement("div");
    DivText.className = "itemstext";
    DivText.innerHTML = items[x].Nombre;

    newDiv.appendChild(DivImage)
    newDiv.appendChild(DivText)

    return newDiv;
}
function onYouTubeIframeAPIReady(){
    player = new YT.Player('player', {
        height: '360',
        width: '640',
        videoId: '',
        playerVars: { 
            'autoplay': 1,
            'controls': 0, 
            'rel' : 0,
            'fs' : 1,
            'loop': 1,
            'showinfo': 0,
            'modestbranding': 0
        },
        events: {
            onReady: onPlayerReady,
            onStateChange: onPlayerStateChange
        }
    });
    console.log(player);
}
function onPlayerReady(event){
    video_duration = event.target.getDuration();
}
function onPlayerStateChange(event){
    if(event.data == YT.PlayerState.PLAYING){
    }
    if(event.data == YT.PlayerState.ENDED){
    }
    if(event.data == YT.PlayerState.PAUSED){
    }
    if(event.data == YT.PlayerState.UNSTARTED){
    }
    if(event.data == YT.PlayerState.BUFFERING){
    }
    if(event.data == YT.PlayerState.CUED){
    }
}
var open_lista = false;
function toogle_lista_curso(){

    var cr = GC("cr", 0).style.right;
    obj = {
        iter: 50,
        els: [
            {el: GC("cr", 0), arr: [{pro: "right", tipo: 0}]},
        ]
    }
    if (cr == "0px"){
        obj.els[0].arr[0].hasta = -145;
    }else{
        obj.els[0].arr[0].hasta = 0;
    }
    animateJs(obj, function(x){
        open_lista = !open_lista
    });
}
function close_lista_curso(){
    obj = { iter: 50, els: [{el: GC("cr", 0), arr: [{pro: "right", tipo: 0, hasta: 0}]}] }
    animateJs(obj, function(){ open_lista = false });
}
function open_lista_curso(){
    obj = { iter: 50, els: [{el: GC("cr", 0), arr: [{pro: "right", tipo: 0, hasta: -145}]}] }
    animateJs(obj, function(){ open_lista = true });
}
function hide_content(){
    GI("flipbook").style.display = "none";
    GI("player").style.display = "none";
    GI("video").style.display = "none";
    GI("video").pause();
    GI("iframe").style.display = "none";
}