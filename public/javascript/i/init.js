
/*
var previewImage = function () {
    var previewPhotoDiv = document.getElementById("previewPhotoDiv");
    previewPhotoDiv.style.display = 'flex';
    var somefile = document.getElementById("photos").files;
    var size = getUploadPhototSize(somefile);
    for (let i = 0 ; i < size; i++) {
        var file = somefile[i]
        //if (!file.type.match('image')) continue;
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
    var previewPhotoDiv = document.getElementById("previewPhotoDiv");
    var previewtext = '<img src="' + e.target.result + '">';
    previewPhotoDiv.insertAdjacentHTML('beforeend', previewtext);
}
*/