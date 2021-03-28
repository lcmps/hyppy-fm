var btn_sbmt = document.querySelector('#sbmt')
btn_sbmt.addEventListener('click', function (e) {
    e.preventDefault()
    btn_sbmt.setAttribute('disabled', true)
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
                btn.classList.add('btn', 'btn-primary');
                dl.appendChild(btn)

                var opts = document.createElement('img');
                opts.src = res.Path
                clg.appendChild(opts);

                document.querySelector('.img-holder').style.display = "block";
                GETUserInfo()

                btn_sbmt.removeAttribute("disabled");

            } else {
                var err = JSON.parse(this.response)
                window.alert(err.Reason)
                btn_sbmt.removeAttribute("disabled");
            }
        }
    })
    xhr.send(JSON.stringify(payload))
}

function GETUserInfo() {
    var usrName = document.querySelector("#username").value;

    var xhr = new XMLHttpRequest();
    xhr.open('GET', cfg.usrAPI + "?username=" + usrName);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.addEventListener('readystatechange', function () {
        if (this.readyState == 4) {
            if (this.status == 200) {
                var res = JSON.parse(this.response)

                var usrImg = document.querySelector('#userImg');
                var uname = document.querySelector('#uName'); uname.value = res.name;
                var rname = document.querySelector('#realName'); rname.value = res.realName;
                var url = document.querySelector('#urlUser'); url.value = res.url;
                var country = document.querySelector('#country'); country.value = res.country;
                var count = document.querySelector('#playcount'); count.value = res.playCount;
                var regDate = document.querySelector('#registered'); regDate.value = res.registered;

                usrImg.src = res.image

                document.querySelector('#usrInfoHolder').style.display = 'block'
            } else {
                var err = JSON.parse(this.response)
                window.alert(err.Reason)
            }
        }
    })
    xhr.send()
}