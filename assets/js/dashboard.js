;(function(){
  var list = [];
  var xhr = new XMLHttpRequest();
  xhr.open("POST", "/api/get_list" , true);
  xhr.send();
  xhr.onreadystatechange = function(res){
    var _list = [];    
    var list = res.target.response
              ?JSON.parse(res.target.response)
              :null;

    for(var i in list){
      _list.push(JSON.parse(list[i]["orbit"]))
    }
    
    list != null && layout(_list);
  }

  function layout(list){
    const table = document.getElementById("dashboard-table");
    table.innerHTML = "<caption><h2>List</h2></caption>";
    for(var a in list){
      let tr = document.createElement("tr");
      tr.className = "dashboard-row";
      tr.innerHTML =
        "<td>" + list[a].Uint + "</td>" +
        "<td>" + list[a].Position + "</td>" +
        "<td>" + list[a].Advice + "</td>"      
      table.appendChild(tr);
    }
  }
})();
