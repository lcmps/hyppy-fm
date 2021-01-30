var btn = document.querySelector('#sbmt')
btn.addEventListener('click', function(e){
    e.preventDefault()
    GETCollage()
})

function GETCollage() {

    var clg = document.querySelector(".img-content");
    var dl = document.querySelector('.dl-btn');
    var payload = {
        username: document.querySelector("#username").value,
        period: document.querySelector("#period").value,
        size: parseInt(document.querySelector("#size").value)
    }
    
    var xhr = new XMLHttpRequest();
    xhr.open('POST', cfg.imgAPI);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.addEventListener('readystatechange', function () {
        if (this.readyState == 4) {
            if (this.status == 200) {
                var res = JSON.parse(this.response)
                var fname = res.Path.split("/")
                var holder = document.querySelector('.img-content');
                holder.innerHTML = ''
                dl.innerHTML = ''
    
                var btn = document.createElement('a');
                btn.href = res.Path;
                btn.id = "dl-img";
                btn.target = "_blank";
                btn.innerText = "Download";
                btn.download = fname[2]
                btn.classList.add('btn','btn-primary');
                dl.appendChild(btn)

                var opts = document.createElement('img');
                opts.src = res.Path
                clg.appendChild(opts);
    
                document.querySelector('.img-holder').style.display = "block";
    
            } else {
                var err = JSON.parse(this.response)
                window.alert(err.Reason)
            }
        }
    })
    xhr.send(JSON.stringify(payload))
}