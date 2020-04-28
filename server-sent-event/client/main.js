function onloaded() {
    var source = new EventSource("/sse/dashboard");
    source.onmessage = function (event) {
        var dashboard = JSON.parse(event.data);
        var items = dashboard["inventory"]["items"];
        document.getElementById("bprice").innerHTML = items["bicycle"].price;
        document.getElementById("bquantity").innerHTML = items["bicycle"].quantity;
    }
}