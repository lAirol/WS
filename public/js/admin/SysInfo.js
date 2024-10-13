class SysInfo {
    info = []
    cpu_count
    net_count
    time = 30//кол-во данных которое мы храним
    append_data = false
    time_labels_arr = []
    CpuCharts
    CpuBars
    NetBars
    NetCharts


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
        this.postData("/system/GetNetInterfacesParams").then((response) =>{
            this.net_count = response.data.count;
            this.NetCharts = new NetCharts(this, response.data);
            this.NetBars = new NetBars(this,response.data);
        })
    }

    prepareInfo(data){
        switch (data.type) {
            case "CPU":{
                this.addInfo(data.sysstat.hosts[0]);
            }break;
            case "INTERNET":{ // TODO тут кончно было бы лучше дописать отдельный прикол в addinfo
                this.addInternetInfo(data.info);
            }
        }
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

    addInternetInfo(info){
        try{
            this.NetBars.updateBars(info);
            this.NetCharts.updateCharts(info);
        }catch (e){}
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

