var showLoginDiv = function () {
    document.getElementById('loginDiv').style.display='block';
}

var closeLoginDiv = function () {
    document.getElementById('loginDiv').style.display='none';
}

window.onclick = function(event) {
    if (event.target == loginDiv) {
        loginDiv.style.display = "none";
    }
}

/*
$("#register").click(function() {

});
*/