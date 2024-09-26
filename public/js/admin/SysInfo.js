class SysInfo {
    info = []
    cpu_charts = []
    cpu_count
    time = 30//кол-во данных которое мы храним
    append_data = false
    time_labels_arr = []


    constructor() {
        let i = this.time;
        for (i; i>0; i--){
            this.time_labels_arr.push(i)
        }
        this.postData("/system/GetCpuCount").then((response) => {
            this.cpu_count = response.data;
            this.createCharts();
            this.append_data = true;
        });
    }

    prepareInfo(data){
        switch (data.type) {
            case "CPU":{
                sysInfo.addInfo(data.sysstat.hosts[0]);
            }break;
        }
    }

    createCharts(){
        let cpu_charts = document.getElementById("cpu_charts");
        for (let i = 0; i<this.cpu_count; i++){
            let div = document.createElement("div");
            div.className = "chart_div";
            let canvas = document.createElement("canvas");
            canvas.id = "cpu_chart_"+i;
            div.append(canvas);
            cpu_charts.append(div);
            this.createNewCPUChart(canvas, i+1);
        }
    }

    createNewCPUChart(ctx,number){
        let chart = new Chart(ctx, {
            type: 'bar',
            data: {
                labels: this.time_labels_arr,
                datasets: [{
                    label: 'CPU №' + number,
                    data: new Array(this.time).fill(0),
                    borderWidth: 1,
                    backgroundColor: this.getRandomColor(),
                }]
            },
            options: {
                scales: {
                    y: {
                        beginAtZero: true,
                        max: 100
                    }
                },
                plugins: {
                    legend: {
                        position: 'top' // Move the legend to the bottom of the chart
                    },
                    tooltip: {
                        backgroundColor: 'rgba(0, 0, 0, 0.8)',
                        titleColor: '#fff',
                        bodyColor: '#fff'
                    }
                }
            }
        });
        this.cpu_charts.push(chart);
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
                this.update_chart(info);

            }catch (e) {
                console.log("addInfo failed")
            }
        }
    }

    update_chart(info){
        let statistic = info.statistics[0]["cpu-load"];
        statistic.forEach((value, key)=>{
            if(key != 0){
                let cpuLoad = value.usr;
                let latestDataset = this.cpu_charts[key-1].data.datasets[0];
                (latestDataset.data.length>=this.time)? latestDataset.data.shift() : null;
                latestDataset.data.push(cpuLoad);
                this.cpu_charts[key-1].update();
            }
        })
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

