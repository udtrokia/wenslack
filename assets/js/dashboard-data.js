// <-- Get Data -->

;(function(){
  let is_requested = false
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
      _list.push(JSON.parse(list[i]["orbit"]));
    }
    list != null && layout(_list);
  }

  // node expression
  function layout(list){
    const table = document.getElementById("table");
    const total = list.length;
    table.innerHTML = "";
    for(var a in list){
      let p_amount = 0;
      let u_amount = 0;
      for(var b in list){
        list[a].Position === list[b].Position?
        p_amount += 1:""
        list[a].Uint === list[b].Uint?
        u_amount += 1:""
      }
      let advice = document.createElement("div");
      advice.className = "right-row center";
      advice.onmouseover = mouse_over;
      advice.onmouseout  = mouse_out;
      advice.attributes.total = list.length;
      advice.attributes.position = list[a].Position;
      advice.attributes.unit = list[a].Uint;
      advice.attributes.p_amount = p_amount;
      advice.attributes.u_amount = u_amount;
      advice.innerHTML = list[a].Advice;
      table.prepend(advice);
    }
  }

  function mouse_over(node){
    let position = document.getElementById("position");
    let unit = document.getElementById("unit");
    let detail = document.getElementById("detail");
    let attrs = node.target.attributes;
    let attr_position = attrs.position;
    let attr_unit = attrs.unit;
    let total = attrs.total;
    position.innerHTML = "";
    unit.innerHTML = "";
    draw_pie("position", attrs.p_amount, total);
    draw_pie("unit", attrs.u_amount, total);
    detail.innerHTML =  "<p class='line'>"
                      + attr_unit
                      + `</p><p class='line'>`
                      + attr_position
                      + "</p>"
  }

  function mouse_out(node){
    d3.selectAll("svg").remove();
    let detail = document.getElementById("detail");
    let position = document.getElementById("position");
    let unit = document.getElementById("unit");
    position.innerHTML = "领域";
    unit.innerHTML = "单位";
    detail.innerHTML = "详情";
  }
  
  function draw_pie(node, amount, total) {
    var dataset = [{
      label: "rgba(78, 186, 169, .9)",
      count: amount
    },{
      label: "transparent",
      count: total - amount
    }];
    var width = 120;
    var height = 120;
    var radius = Math.min(width, height) / 2;
    var color = d3.scaleOrdinal(d3.schemeCategory20);
    var svg = d3.select("#" + node)
                .append('svg')
                .attr('width',width)
                .attr('height',height)
                .append('g')
                .attr('transform', 'translate(' + (width/2) +  ',' + (height / 2) + ')');

    var arc = d3.arc()
                .innerRadius(0)
                .outerRadius(radius);

    var pie = d3.pie()
                .value(function(d) { return d.count; })
                .sort(null);

    var path = svg.selectAll('path')
                  .data(pie(dataset))
                  .enter()
                  .append('path')
                  .attr('d',arc)
                  .attr('fill',function(d,i){
                    return d.data.label
                  })    
  }
  
})();
