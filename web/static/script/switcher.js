load()
function save() {
    var checkbox = document.getElementById('button');
    localStorage.setItem('button-save', checkbox.checked);

    if (checkbox.checked) {
        window.location.href = "/weapons";
    } else {
        window.location.href = "/descendants";
    }
}
function load() {
    var ischecked = localStorage.getItem('button-save');
    if (ischecked === 'true') {
        document.getElementById("button").checked = true;
    } else {
        document.getElementById("button").checked = false;
    }
}