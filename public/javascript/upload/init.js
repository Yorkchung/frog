/*
var getCookie = function (cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
        for(var i = 0; i < ca.length; i++) {
            var c = ca[i];
            while (c.charAt(0) == ' ') {
                c = c.substring(1);
            }
            if (c.indexOf(name) == 0) {
               return c.substring(name.length, c.length);
            }
        }
    return "";
}
*/

/*
// barMenu Trigger
var showDashboard = function () {
    var dashboardStyle = document.getElementById("dashboard").style;
    dashboardStyle.display = "block";
}

window.onclick = function (event) {
    if (!event.target.matches('.setting')) {
        var dashboardStyle = document.getElementById("dashboard").style;
        dashboardStyle.display = "none";
    }
}
*/

/*
function textareaAutoSize(id, maxHeight) {
   var text = id && id.style ? id : document.getElementById(id);

   if (text.clientHeight == text.scrollHeight) {
      text.style.height = "30px";
   }

   var adjustedHeight = text.clientHeight;
   if (!maxHeight || maxHeight > adjustedHeight) {
      adjustedHeight = Math.max(text.scrollHeight, adjustedHeight);
      if (maxHeight)
         adjustedHeight = Math.min(maxHeight, adjustedHeight);
      if (adjustedHeight > text.clientHeight)
         text.style.height = adjustedHeight + "px";
   }
}
*/

/*
//somefile[i].type , somefile[i].size
var previewImage = function () {
    var previewPhotoDiv = document.getElementById("preview-photo-div");
    previewPhotoDiv.style.display = 'flex';
    var somefile = document.getElementById("photos").files;
    var size = getUploadPhototSize(somefile);
    for (let i = 0 ; i < size; i++) {
        var file = somefile[i];
        var reader = new FileReader();
        reader.addEventListener("load", loadPreviewImage);
        reader.readAsDataURL(file);
    }
}

var getUploadPhototSize = function(obj) {
    var size = 0, key;
    for (key in obj) {
        if (obj.hasOwnProperty(key)) size++;
    }
    return size;
}

function loadPreviewImage(e) {
    var previewPhotoDiv = document.getElementById("preview-photo-div");
    var previewtext = '<img src="' + e.target.result + '">';
    previewPhotoDiv.insertAdjacentHTML('beforeend', previewtext);
}
*/