class SysInfo {
    info = []
    cpu_count
    time = 30//кол-во данных которое мы храним
    append_data = false
    time_labels_arr = []
    CpuCharts
    CpuBars


    constructor() {
        let i = this.time;
        for (i; i>0; i--){
            this.time_labels_arr.push(i)
        }
        this.postData("/system/GetCpuCount").then((response) => {
            this.cpu_count = response.data;
            this.CpuCharts = new CpuCharts(this);
            this.CpuBars = new CpuBars(this);
            this.append_data = true;
        });
    }

    prepareInfo(data){
        switch (data.type) {
            case "CPU":{
                this.addInfo(data.sysstat.hosts[0]);
            }break;
        }
    }

    createNewCPUBar(){
        let bar_div = document.createElement("div");
        let bar = new ProgressBar.SemiCircle(bar_div, {
            strokeWidth: 6,
            color: '#FFEA82',
            trailColor: '#eee',
            trailWidth: 1,
            easing: 'easeInOut',
            duration: 1400,
            svgStyle: null,
            text: {
                value: '',
                alignToBottom: false
            },
            from: {color: '#FFEA82'},
            to: {color: '#ED6A5A'},
            // Set default step function for all animate calls
            step: (state, bar) => {
                bar.path.setAttribute('stroke', state.color);
                var value = Math.round(bar.value() * 100);
                if (value === 0) {
                    bar.setText('');
                } else {
                    bar.setText(value);
                }

                bar.text.style.color = state.color;
            }
        });

    }

    getRandomColor() {
        const letters = '0123456789ABCDEF';
        let color = '#';
        for (let i = 0; i < 6; i++) {
            color += letters[Math.floor(Math.random() * 16)];

        }
        return color;
    }

    addInfo(info){
        if(this.append_data){
            try{
                (this.info.length>this.time) ? this.info.shift() : null;
                this.info.push(info)
                this.CpuCharts.updateChart(info);
                this.CpuBars.updateBars(info);
            }catch (e) {
                console.log("addInfo failed")
            }
        }
    }



    async postData(url = "", data = {}){
        const response = await fetch(url, {
            method: "POST", // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // no-cors, *cors, same-origin
            cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
            credentials: "same-origin", // include, *same-origin, omit
            headers: {
                "Content-Type": "application/json",
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            redirect: "follow", // manual, *follow, error
            referrerPolicy: "no-referrer", // no-referrer, *client
            body: JSON.stringify(data), // body data type must match "Content-Type" header
        });
        return await response.json(); // parses JSON response into native JavaScript objects
    }

}

