$( "#records-frog" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/gallery/data",
        data: { searchtype: "tag", keyword:"frog" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result);
        //console.log(result.GroupByTag);
        var HTML = "<div>";
        HTML += "<table border='1'><tbody>";
        jQuery.each(result.GroupByTag, function(name) {
            //console.log(name); //莫氏樹蛙
            HTML += "<tr><td class='dataOrganismName'>" + "<button class='showAlbum' value='" + name + "' >" + name + "</button></td></tr>";
            /*
            jQuery.each(result.GroupByTag[name], function(galleryKey, records) {
                //console.log(galleryKey);
                //console.log(records); //莫氏樹蛙的每筆記錄
                jQuery.each(records, function(recordID, record) {
                    //console.log(recordID); //莫氏樹蛙的每筆記錄的ID
                    jQuery.each(record["PhotoSrc"], function(photoID, PhotoSrc) {
                        console.log(photoID); //該記錄當中圖片的ID
                        console.log(PhotoSrc); //該記錄當中圖片的位址
                    });
                });
            });
            */
        });
        HTML += "</tbody></table>";
        HTML += "</div></br>";
        HTML += '';
        $('#records-data').prepend(HTML);
        $(".showAlbum").click(function(e) {
            name = $(this).val();
            //console.log(name);
            imgDivHTML = "";
            jQuery.each(result.GroupByTag[name], function(galleryKey, records) {
                //console.log(galleryKey);
                //console.log(records); //莫氏樹蛙的每筆記錄
                imgDivHTML += "<div class='single-album-div'>";
                imgDivHTML += "<ul>";
                jQuery.each(records, function(recordID, record) {
                    //console.log(recordID); //莫氏樹蛙的每筆記錄的ID
                    jQuery.each(record["PhotoSrc"], function(photoID, PhotoSrc) {
                        imgDivHTML += "<li>";
                        imgDivHTML += "<img src='/storage/photo/" + PhotoSrc + "'>";
                        imgDivHTML += "</li>";
                        //console.log(photoID); //該記錄當中圖片的ID
                        //console.log(PhotoSrc); //該記錄當中圖片的位址
                    });
                    //imgDivHTML += "<button class='getRecord' value='" + recordID + "' >來自</button>";
                });
                imgDivHTML += "</ul>";
                imgDivHTML += "</div>";
                $('#album').prepend(imgDivHTML);
                $("#album div ul li").first().addClass('selected');
            });
            $('#album').show();

                //$('#album').hide();

        });

    });
});

$( "#records-lepidoptera" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/gallery/data",
        data: { searchtype: "tag", keyword:"lepidoptera" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        console.log(result);
    });
});

$( "#records-plant" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/gallery/data",
        data: { searchtype: "tag", keyword:"plant" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        console.log(result);
    });
});