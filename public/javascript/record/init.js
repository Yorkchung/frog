$( document ).ready(function() {
    var id = window.location.search.substring(1);

    $.ajax({
        type: "GET",
        url : "/record/data",
        data: { recordid: id }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        console.log(result);
        
        HTML = "";
        // HTML += "<a herf='records.html'></a>";
        HTML += "<h2>物種名&nbsp:&nbsp;<class='OrganismName'>" + result.OrganismName + "</h2><a href='http://127.0.0.1/records' class='back_to_records'>返回</a>";
        // HTML += "<tr><td></td><td class='Food'>" + result.Food + "</td><tr/>";
        HTML += "<table><tr><td class='title'>季節</td><td class='Season'>" + result.Season + "</td><tr/>";
        HTML += "<tr><td class='title'>年齡</td><td class='Stage'>" + result.Stage + "</td><tr/>";
        HTML += "<tr><td class='title'>性狀</td><td class='Status'>" + result.Status + "</td><tr/>";
        HTML += "<tr><td class='title'>棲息地</td><td class='Habitat'>" + result.Habitat + "</td><tr/>";
        HTML += "<tr><td class='title'>備註</td><td class='Note'>" + result.Note + "</td><tr/></table>";
        HTML += "<div class='Photo'>";

        jQuery.each(result.PhotoSrc, function(i, val) {
            //HTML += "<a target='_blank' href='/storage/photo/"+result.PhotoSrc[i] +"' >";
            HTML += "<img class='PhotoSrc' src='/storage/photo/" + result.PhotoSrc[i] + "' ><br/>";
            
        });
        HTML += "</div>";
        $('#record-data').prepend(HTML);
        
        (function(){
            Galleria.loadTheme('/resource/galleria/themes/classic/galleria.classic.js');
            //Galleria.loadTheme('/resource/galleria/themes/fullscreen/galleria.fullscreen.min.js');
            Galleria.run('.Photo');
        }());                          
    });
});