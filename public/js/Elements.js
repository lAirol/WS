class Elements {
    static createRoundElement(data, color) {
        const round = document.createElement("div");
        round.setAttribute("class", "round");
        round.style.left = data.offsetX-7.5 + "px";
        round.style.top = data.offsetY-7.5 + "px";
        round.classList.add("fade-out");
        // round.style.backgroundColor = randomRgbColor();
        round.style.backgroundColor = color;
        document.getElementById("output").append(round);
        setTimeout(() => {
        round.remove();
        }, 1000);
    }

    static createSquareElement(data) {
        const square = document.createElement("div");
        square.setAttribute("class", "square");
        square.style.left = data.offsetX-15 + "px";
        square.style.top = data.offsetY-15 + "px";
        square.classList.add("fade-out1");
        document.getElementById("output").append(square);
        setTimeout(() => {
            square.remove();
        }, 5000);
    }

    static randomRgbColor(){
        const red = Math.floor(Math.random() * 256);
        const green = Math.floor(Math.random() * 256);
        const blue = Math.floor(Math.random() * 256);
        return {red:red,green:green,blue:blue};
    };
}