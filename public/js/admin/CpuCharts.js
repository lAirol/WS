class CpuCharts{
    sysInfo
    cpu_charts = []

    constructor(sysInfo) {
        this.sysInfo = sysInfo;
        this.createCharts();
    }

    createCharts(){
        let cpu_charts = document.getElementById("cpu_charts");
        for (let i = 0; i<this.sysInfo.cpu_count; i++){
            let div = document.createElement("div");
            div.className = "chart_div";
            let canvas = document.createElement("canvas");
            canvas.id = "cpu_chart_"+i;
            div.append(canvas);
            cpu_charts.append(div);
            this.cpu_charts.push(this.createNewCPUChart(canvas, i+1));
        }
    }

    createNewCPUChart(ctx,number){
        return new Chart(ctx, {
            type: 'bar',
            data: {
                labels: this.sysInfo.time_labels_arr,
                datasets: [{
                    label: 'CPU №' + number,
                    data: new Array(this.sysInfo.time).fill(0),
                    borderWidth: 1,
                    backgroundColor: this.sysInfo.getRandomColor(),
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
    }

    updateChart(info){
        let statistic = info.statistics[0]["cpu-load"];
        statistic.forEach((value, key)=>{
            if(key != 0){
                let cpuLoad = value.usr;
                let latestDataset = this.cpu_charts[key-1].data.datasets[0];
                (latestDataset.data.length>=this.sysInfo.time)? latestDataset.data.shift() : null;
                latestDataset.data.push(cpuLoad);
                this.cpu_charts[key-1].update();
            }
        });
    }
}