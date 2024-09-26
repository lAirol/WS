class Window {
    modal
    innerModal
    closeSpan

    createWindow(){
        this.modal = document.createElement("div");
        this.modal.className = "modal";

        this.innerModal = document.createElement("div");
        this.innerModal.className = "modal-content";

        this.topBtnDiv = document.createElement("div");
        this.topBtnDiv.className = "topBtnDiv";

        this.closeSpan = document.createElement("span");
        this.closeSpan.className = "close";
        this.closeSpan.onclick = () => {
            this.modal.remove();
        };
        let closeIcon = document.createElement("i");
        closeIcon.className = "zmdi zmdi-close";
        this.closeSpan.appendChild(closeIcon);

        this.fullscreenBtn = document.createElement("span");
        this.fullscreenBtn.className = "fullscreen";
        this.fullscreenBtn.onclick = () => {
            this.modal.classList.toggle("fullscreen");
        };
        let fullscreenIcon = document.createElement("i");
        fullscreenIcon.className = "zmdi zmdi-window-maximize";
        this.fullscreenBtn.appendChild(fullscreenIcon);

        this.minimizeBtn = document.createElement("span");
        this.minimizeBtn.className = "minimize";
        this.minimizeBtn.onclick = () => {
            this.modal.classList.toggle("minimized");
        };
        let minimizeIcon = document.createElement("i");
        minimizeIcon.className = "zmdi zmdi-window-minimize";
        this.minimizeBtn.appendChild(minimizeIcon);

        this.topBtnDiv.append(this.minimizeBtn, this.fullscreenBtn, this.closeSpan)
        this.innerModal.append(this.topBtnDiv);
        this.modal.append(this.innerModal);
        return this.modal;
    }

    showModal(){
        this.modal.style.display = "block";
    }
}

function openDialog(){
    let window = new Window();
    document.body.append(window.createWindow());
    window.showModal();
}