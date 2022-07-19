function getUsers() {
    let url = 'http://localhost:8080/wish';
    var content=document.getElementById("greet")
    var contenttime=document.getElementById("time")
    var contentquote=document.getElementById("quote")
    fetch(url)
    .then((response) => {
      return response.json();
    })
    .then((data) => {
      content.innerHTML=data.name;
      contenttime.innerHTML=data.time;
      contentquote.innerHTML=data.sendquote;
      });

}

