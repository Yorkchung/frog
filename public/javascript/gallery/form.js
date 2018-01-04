$( "#records-frog" ).click(function(e) {
    $('#records-data').empty();
    $.ajax({
        type: "GET",
        url: "/gallery/data",
        data: { searchtype: "tag", keyword:"frog" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        console.log(result);
        //console.log(result);
//var index;
//        var photoSrc;
        var key = [];
        var photos=[];
        
        jQuery.each(result.GroupByTag, function(name) {
            key.push(name);
        });
        //console.log(key);
        var imgDivHTML = "";
        
        //var image = Object.keys(result.GroupByTag);
        //console.log(result.GroupByTag);
        //console.log(key);
        for(var i =0; i<key.length; i++){
            console.log(result.GroupByTag[key[i]][0]);
            photos.push(result.GroupByTag[key[i]][0]);
        }
        //console.log(photos);
        
        for(var i =0; i<photos.length; i++){
            imgDivHTML += "<button class='showAlbum' value='" + key[i] + "'><img  width='200' height='200' src='/storage/photo/" + photos[i]+"' style='padding:10px;''>"+key[i]+"</button>";
        }
//       for(var i=0;i<key.length;i++){
//           console.log(i);
//            jQuery.each(result.GroupByTag[key[i]], function(index, photoSrc){
//                photos.push(photoSrc);
//                //console.log(photos[0]);
//              });
//                imgDivHTML += "<botton value='" + key[i] + "'><img class='showAlbum' width='200' height='200' src='/storage/photo/" + photos[i]+"' style='padding:10px;''>"+key[i]+"</botton>";    
//            
//             
//       }
        $('#records-data').prepend(imgDivHTML);
        $(".showAlbum").click(function(e) {
            $('#album-ul').empty();
            var HTML ="";
            var number=$(this).val();
            var image=[];
            image.push(result.GroupByTag[number]);
                console.log(image);        
             for(var i =0; i<image[0].length; i++){
             HTML += "<li><img width='400' height='400' src='/storage/photo/" + image[0][i]+"'></li>";
             
             } 
            console.log(HTML);
             $('#album-ul').prepend(HTML);    
            $('.album-ul').css('display', 'flex');
            $('#album').show().css('display', 'flex');
            $(".single-album-div ul li").first().show();            
            });  
        
    });
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
$( "#close-album-div-button" ).click(function(e) {
      $( "#album" ).hide();
  });






/*
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
        var photos=[];
        jQuery.each(result.GroupByTag, function(name) {
            key.push(name);
        });
        //console.log(key);
        var imgDivHTML = "";
        var count=0;
        for(var i=0;i<key.length;i++){
            jQuery.each(result.GroupByTag[key[i]], function(index, photoSrc){
                imgDivHTML += "<img class='showAlbum' width='200' height='200' src='/storage/photo/" + photoSrc +"' style='padding:10px;''>";
                count++;
                if(count==5){
                imgDivHTML += "<br/>"
                count=0;
                }
            });
        }
        
        $('#records-data').prepend(imgDivHTML);
        $(".showAlbum").click(function(e){
            image=$(this).attr("src");
            console.log(image);
           // key.push(result.GroupByTag);
            //console.log(key);
           // photos.push(result.GroupByTag[key[0]]);
            swal({
                
                html: '<img width="400" height="400" src="'+image+'">'
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
*/

//下面這段先留著 別刪！

/*$( "#records-frog" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/gallery/data",
        data: { searchtype: "tag", keyword:"frog" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        console.log(result);
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
});*/
