class Options {
    back_color
    opt_div = document.getElementById("options");
    colors_elems = {}
    start_color = {}
    new_color = {}
    displayDiv
    constructor(options) {
        this.start_color =  Elements.randomRgbColor();
        this.new_color = this.start_color;
        this.createBackColor(this.start_color.red,this.start_color.green,this.start_color.blue);
        this.createOptionsDiv();
        this.opt_div.onclick = (e) => {
            if(e.target.id === this.opt_div.id || e.target.id === "options_img") {
                let opt_main = document.getElementById("options_main");
                (opt_main.classList.contains("active")) ? opt_main.classList.remove("active"): opt_main.classList.add("active");
            }
        }
    }

    createOptionsDiv(){
        let mainDiv = document.createElement("div");
        mainDiv.id = "options_main";
        mainDiv.className = "options_main";

        let innerDiv = document.createElement("div");
        innerDiv.className = "options_inner";
        innerDiv.id = "options_inner";

        let img = document.createElement("img");
        img.id = "options_img";
        img.src = "https://cdn-icons-png.flaticon.com/512/6812/6812607.png";

        let colorDiv = document.createElement("div");
        colorDiv.id = "options_color";
        colorDiv.className = "options_color";

        this.createColorSelect(colorDiv);

        this.opt_div.append(innerDiv);
        innerDiv.append(img);
        innerDiv.appendChild(mainDiv);
        mainDiv.appendChild(colorDiv);
    }

    createColorSelect(colorDiv) {
        let redValue = this.start_color.red;
        let greenValue = this.start_color.green;
        let blueValue = this.start_color.blue;

        let redPicker = this.createColorPicker("Красный", redValue, "red");
        let greenPicker = this.createColorPicker("Зелёный", greenValue, "green");
        let bluePicker = this.createColorPicker("Синий", blueValue, "blue");

        let displayDiv = document.createElement("div");
        displayDiv.className = "display";
        displayDiv.id = "display_selected_color_div";
        displayDiv.style.backgroundColor = this.back_color;

        this.displayDiv = displayDiv;

        colorDiv.append(redPicker, greenPicker, bluePicker, displayDiv);
    }

    createColorPicker(colorLabel, colorValue, colorId) {
        let label = document.createElement("label");
        label.innerText = colorLabel;
        label.className = "color_label";
        label.setAttribute("for", `${colorId}_input`);

        let input = document.createElement("input");
        input.id = `${colorId}_input`;
        input.className = "color_picker";
        input.type = "range";
        input.min = 0;
        input.max = 255;
        input.value = colorValue;
        input.oninput = (evt) => {
            this.new_color[colorId] = input.value;
            this.createBackColor(this.start_color.red,this.start_color.green,this.start_color.blue);
            this.displayDiv.style.backgroundColor = this.back_color;
        }


        let wrapperDiv = document.createElement("div");
        wrapperDiv.className = "color_picker_wrapper";
        wrapperDiv.appendChild(label);
        wrapperDiv.appendChild(input);

        return wrapperDiv;
    }


    createBackColor(r,g,b){
        this.back_color = `rgb(${r}, ${g}, ${b})`;
    }

}