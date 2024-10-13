class NetBars{
    bars = {
        get:[],
        post:[],
        speed:{
            labels:{
                get:[],
                post:[]
            },
            value:{
                get:[],
                post:[]
            }
        }
    }
    info

    constructor(sysinfo, info) {
        this.info = info;
        this.createBars(sysinfo.net_count);
        this.bars.speed.value.get = Array(sysinfo.net_count).fill(null);
        this.bars.speed.value.post = Array(sysinfo.net_count).fill(null);
    }

    createBars(count){
        let net_bars_div = document.getElementById("net_bars");
        for(let i = 0; i<count; i++){
            net_bars_div.append(this.createBarDiv(i));
        }
    }

    createBarDiv(number){
        let bars_div = document.createElement("div");
        bars_div.className = "net_div";
        let span = document.createElement("span");
        span.innerText = "NET INTERFACE "+this.info.name[number];
        let send = document.createElement("div");
        send.className = "bar_div";
        let span_send = document.createElement("span");
        span_send.innerText = "ОТПРАВЛЕНО";
        let span_send_max = document.createElement("span");
        span_send_max.innerText = "MAX BYTE SENT: "
        let get_value = document.createElement("span");
        span_send_max.append(get_value);
        send.append(span_send);

        let post = document.createElement("div");
        post.className = "bar_div";
        let span_post = document.createElement("span");
        span_post.innerText = "ПОЛУЧЕНО";
        let span_post_max = document.createElement("span");
        span_post_max.innerText = "MAX BYTE POST: "
        let post_value = document.createElement("span");
        span_post_max.append(post_value);
        post.append(span_post);

        bars_div.append(span);
        bars_div.append(send);
        bars_div.append(span_send_max);

        bars_div.append(post);
        bars_div.append(span_post_max)

        this.bars.get.push(this.createSendBar(send));
        this.bars.post.push(this.createPostBar(post));
        this.bars.speed.labels.get.push(get_value);
        this.bars.speed.labels.post.push(post_value);
        return bars_div;
    }

    createSendBar(bar_div){
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
            from: {color: '#f69292'},
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

    createPostBar(bar_div){
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
            from: {color: '#b4c0ff'},
            to: {color: '#0035ff'},
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
        Object.keys(this.bars.get).forEach((index,key)=>{
            this.calculateMax(info[key],key);
            let counterElement = this.bars.speed.labels.get[key];
            anime({
                targets: counterElement,
                textContent: [counterElement.textContent,  this.bars.speed.value.get[key]],
                easing: 'easeInOutQuint',
                duration: 1000,
                round: true
            });
            counterElement = this.bars.speed.labels.post[key];
            anime({
                targets: counterElement,
                textContent: [counterElement.textContent, this.bars.speed.value.post[key]],
                easing: 'easeInOutQuint',
                duration: 1000,
                round: true
            });
            // this.bars.speed.labels.get[key].innerText = this.bars.speed.value.get[key] /1024 / 1024;
            // this.bars.speed.labels.post[key].innerText = this.bars.speed.value.post[key] /1024 / 1024;

            if(info[key].recv > 0 && this.bars.speed.value.get[key] > 0){
                this.bars.get[key].animate(info[key].recv / this.bars.speed.value.get[key]);
            }else{
                this.bars.get[key].animate(0.00);
            }
            if(info[key].sent > 0 && this.bars.speed.value.post[key] > 0){
                this.bars.post[key].animate(info[key].sent / this.bars.speed.value.post[key]);
            }else{
                this.bars.post[key].animate(0.00);
            }

            // this.bars.post[key].animate(this.bars.speed.value.post[key] / info[key].recv);
        });
    }

    calculateMax(value, key) {
        this.bars.speed.value.get[key] = Math.max(this.bars.speed.value.get[key], value.recv);
        this.bars.speed.value.post[key] = Math.max(this.bars.speed.value.post[key], value.sent);
    }
}