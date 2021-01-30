var btn = document.querySelector('#btn-user');
btn.addEventListener('click', function(e) {
    e.preventDefault()
    GETUserInfo()
})

function GETUserInfo(){    
    var usrName = document.querySelector('#username-info').value;

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