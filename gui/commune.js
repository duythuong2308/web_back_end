// this func assumes all elements in input array are maps with same keys,
// :arg tableElem: example: document.getElementById("table0")
function loadArrayToTable(tableElem, array) {
    if (array.length < 1) {
        console.log("empty data:", array);
        return
    }
    let keys = Object.keys(array[0]);

    let tbody = document.createElement("tbody");
    let tr = document.createElement("tr");

    for (let col = 0; col < keys.length; col ++) {
        let bold = document.createElement("strong");
        bold.appendChild(document.createTextNode(keys[col]));
        let tdKey = document.createElement("td");
        tdKey.appendChild(bold);
        tr.appendChild(tdKey);

    }
    tbody.appendChild(tr);

    for (let row = 0; row < array.length; row++) {
        let tr = document.createElement("tr");

        for (let col = 0; col < keys.length; col++) {
            let td = document.createElement("td");
            let value = array[row][keys[col]];
            if (!value) {
                tr.appendChild(td);
                continue
            }
            if (keys[col] === "Id") {
                let a = document.createElement('a');
                a.appendChild(document.createTextNode(value));
                a.href = `/gui/village.html?communeId=${value}`;
                td.appendChild(a);
            } else {
                td.appendChild(document.createTextNode(value));
            }
            tr.appendChild(td)
        }
        let td = document.createElement("td");
        var x = document.createElement("BUTTON");
        x.id = `deleteCommune${array[row]["Id"]}`;
        var t = document.createTextNode("Xóa");
        x.appendChild(t);
        td.appendChild(x);
        tr.appendChild(td);
        tbody.appendChild(tr);

        x.addEventListener("click",function () {
            var xhr = new XMLHttpRequest();
            xhr.open("DELETE", '/api/commune', true);
            xhr.send(JSON.stringify({"Id": array[row]["Id"]}));
            window.location.reload();
        })
    }

    tableElem.innerHTML = "";
    tableElem.appendChild(tbody);
}

window.onload = async () => {
    console.log("load communes");
    let resp = await FetchBackendRows(GetBackend(),
        "/api/commune"+window.location.search);
    loadArrayToTable(document.getElementById("communes"), resp["Data"]);
};
