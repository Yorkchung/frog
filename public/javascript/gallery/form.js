$( "#records-frog" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/gallery/data",
        data: { searchtype: "tag", keyword:"frog" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        console.log(result);
        //console.log(result);

        var key = [];
        jQuery.each(result.GroupByTag, function(name) {
            key.push(name);
        });
        //console.log(key);
        var imgDivHTML = "";
        var count=0;
        for(var i=0;i<key.length;i++){
            jQuery.each(result.GroupByTag[key[i]], function(index, photoSrc){
                imgDivHTML += "<img class='showAlbum' width='100' height='100' src='/storage/photo/" + photoSrc + "'>";
                count++;
                if(count==5){
                imgDivHTML += "<br/>"
                count=0;
                }
            });
        }
        $('#records-data').prepend(imgDivHTML);
        $(".showAlbum").click(function(e){
            swal({
                type: 'info',
                html: '<h1>dffd</h1>'+result.GroupByTag['莫氏樹蛙']+'<img src="">'
            });
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

$( "#close-album-div-button" ).click(function(e) {
    $( "#album" ).hide();
});


//下面這段先留著 別刪！
/*
$( "#records-frog" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/gallery/data",
        data: { searchtype: "tag", keyword:"frog" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        console.log(result)
        var HTML = "<div>";
        HTML += "<table border='1'><tbody>";
        jQuery.each(result.GroupByTag, function(name) {
            HTML += "<tr><td class='dataOrganismName'>" + "<button class='showAlbum' value='" + name + "' >" + name + "</button></td></tr>";
        });
        HTML += "</tbody></table>";
        HTML += "</div></br>";
        HTML += '';
        $('#records-data').prepend(HTML);
        $(".showAlbum").click(function(e) {
            name = $(this).val();
            imgDivHTML = "";
            jQuery.each(result.GroupByTag[name], function(index, photoSrc) {
                imgDivHTML += "<li><img src='/storage/photo/" + photoSrc + "'></li>";
                $('#album-ul').prepend(imgDivHTML);
            });
            $('.album-ul').css('display', 'flex');
            $('#album').show().css('display', 'flex');
            $(".single-album-div ul li").first().show();            
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

$( "#close-album-div-button" ).click(function(e) {
    $( "#album-ul" ).empty();
    $( "#album" ).hide();
});

$(".next").click(function(){
    if ($(".album-ul li:visible").next().length != 0)
        $(".album-ul li:visible").next().show().prev().hide();
    else {
        $(".album-ul li:visible").hide();
        $(".album-ul li:first").show();
        return false;
    }
});

$(".prev").click(function(){
    if ($(".album-ul li:visible").prev().length != 0)
        $(".album-ul li:visible").prev().show().next().hide();
    else {
        $(".album-ul li:visible").hide();
        $(".album-ul li:last").show();
        return false;
    }
});
*/