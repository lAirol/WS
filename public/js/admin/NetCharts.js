class NetCharts{
    charts = {
        get:[],
        post:[]
    }
    net_charts = []
    info
    sysInfo

    constructor(sysinfo, info) {
        this.info = info;
        this.sysInfo = sysinfo;
        this.createCharts(sysinfo.net_count);
    }

    createCharts(count){
        let net_bars_div = document.getElementById("net_charts");
        for(let i = 0; i<count; i++){
            net_bars_div.append(this.createBarDiv(i));
        }
    }

    createBarDiv(number){
        let charts_div = document.createElement("div");
        charts_div.className = "chart_div";
        let canvas = document.createElement("canvas");
        let span = document.createElement("span");
        span.innerText = "NET INTERFACE "+this.info.name[number];
        charts_div.append(span);
        charts_div.append(canvas);

        this.net_charts.push(this.createNewNetworkChart(canvas, number+1));

        return charts_div;
    }

    createNewNetworkChart(ctx, number) {
        return new Chart(ctx, {
            type: 'line',
            data: {
                labels: this.sysInfo.time_labels_arr,
                datasets: [
                    {
                        label: 'Отправлено',
                        data: new Array(this.sysInfo.time).fill(null),
                        borderWidth: 2,
                        backgroundColor: "rgba(255, 0, 0, 0.2)",
                        borderColor: "red",
                        fill: true,
                        pointRadius: 0, // Убираем точки
                        tension: 0.4, // Добавляем сглаживание
                    },
                    {
                        label: 'Получено',
                        data: new Array(this.sysInfo.time).fill(null),
                        borderWidth: 2,
                        backgroundColor: "rgba(0, 0, 255, 0.2)",
                        borderColor: "blue",
                        fill: true,
                        pointRadius: 0, // Убираем точки
                        tension: 0.4, // Добавляем сглаживание
                    }
                ]
            },
            options: {
                scales: {
                    x:{
                        beginAtZero: false
                    },
                    y: {
                        beginAtZero: true,
                        //max: 100
                    }
                },
                plugins: {
                    legend: {
                        position: 'top' // Переместить легенду в верхнюю часть графика
                    },
                    tooltip: {
                        backgroundColor: 'rgba(0, 0, 0, 0.8)',
                        titleColor: '#fff',
                        bodyColor: '#fff'
                    }
                },
                animation: {
                    //duration: 0,
                    //easing: 'linear' // Устанавливаем линейное изменение
                }
            }
        });
    }



    updateCharts(info) {
        Object.entries(info).forEach(([key, value]) => {
            const chart = this.net_charts[key];
            chart.data.datasets[0].data.shift();
            chart.data.datasets[0].data.push(value.sent/125000);
            chart.data.datasets[1].data.shift();
            chart.data.datasets[1].data.push(value.recv/125000);

            chart.update();
        });
    }

}