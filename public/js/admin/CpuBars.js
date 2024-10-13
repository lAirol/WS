class CpuBars {
    cpu_bars = []
    sysInfo

    constructor(sysInfo) {
        this.sysInfo = sysInfo;
        this.createCpuBars();
    }

    createCpuBars(){
        let cpu_bars = document.getElementById("cpu_bars");
        for (let i = 0; i<this.sysInfo.cpu_count; i++){
            let div = document.createElement("div");
            let span = document.createElement("span");
            span.innerText = "CPU №"+(i+1);
            div.className = "bar_div";
            div.append(span)
            cpu_bars.append(div);

            this.cpu_bars.push(this.createNewCPUBar(div));

        }
    }

    createNewCPUBar(bar_div){
        return new ProgressBar.SemiCircle(bar_div, {
            strokeWidth: 6,
            color: '#ffd700',
            trailColor: '#eee',
            trailWidth: 1,
            easing: 'easeInOut',
            duration: 1400,
            svgStyle: null,
            text: {
                value: '',
                alignToBottom: false
            },
            // from: {color: '#00ff8d'},
            // to: {color: '#d000ff'},
            from: {color: '#0021ff'},
            to: {color: '#ff1a00'},
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

    updateBars(info){
        let statistic = info.statistics[0]["cpu-load"];
        statistic.forEach((value, key)=>{
            if(key != 0){
                let cpuLoad = value.usr;
                this.cpu_bars[key-1].animate(cpuLoad/100);
            }
        });
    }
}