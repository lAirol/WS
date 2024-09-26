let sysInfo;
window.onload = ()=>{
    menuExpand();
    sysInfo = new SysInfo();
}

function toggleSidebar(){
    let sidebar = document.getElementById("sidebar");

    if (document.activeElement === sidebar && document.activeElement === sidebar) {
        sidebar.blur();
    } else {
        sidebar.focus();
    }
}

function menuExpand(){
    let menu = document.getElementById("sidebar-elements");
    menu.querySelectorAll("li.parent").forEach((li) => {
         li.addEventListener("click",(event)=>{
             event.stopPropagation();
             expandTree(li);
         })
    });
}

function expandTree(li){
    const children = li.querySelectorAll('ul.sub-menu');
    children[0].style.display = children[0].style.display === 'block' ? 'none' : 'block';
}

function showDropdown(target){
    let drop = document.getElementById(target);
    drop.style.display = drop.style.display === 'block' ? 'none' : 'block';

    drop.focus();
    drop.addEventListener('blur', () => {
        drop.style.display = 'none';
    });
}

const socket = new WebSocket('ws://localhost:8080/wsadmin');
socket.onopen = () => {
    console.log('Connected to the WebSocket server');
};

socket.onmessage = (event) => {
    let data = event.data;
    try {
        data = JSON.parse(data);
    } catch (ex) {}
    switch (data.target){
        case wsConst.chat:{
            try{
                addMessage(data.message,false)
            }catch (e) {
                console.log(e);
            }
        }break;
        case wsConst.sys_info:{
            sysInfo.prepareInfo(data);
        }break;
    }
};

window.addEventListener("load", () => {

})

function sendMessage(message,target) {
    message = {
        message:message,
        target: wsConst.chat
    }
    socket.send(JSON.stringify(message));
}

const wsConst =  {
        "points": 1,
        "chat": 2,
        "sys_info": 3
}