$( document ).ready(function() {
    var id = window.location.search.substring(1);

    $.ajax({
        type: "GET",
        url : "/record/data",
        data: { recordid: id }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result.ID);
        //console.log(result.OrganismName);
        //console.log(result.Food);
        //console.log(result.Stage);
        //console.log(result.Habitat);
        //console.log(result.Note);
		//PhotoLatitude
		//PhotoLongitude
        console.log(result);
        
        HTML = "";
        HTML += "<p class='OrganismName'>" + result.OrganismName + "</p><br/>";
        HTML += "<p class='Food'>" + result.Food + "</p><br/>";
        HTML += "<p class='Season'>" + result.Season + "</p><br/>";
        HTML += "<p class='Stage'>" + result.Stage + "</p><br/>";
        HTML += "<p class='Status'>" + result.Status + "</p><br/>";
        HTML += "<p class='Habitat'>" + result.Habitat + "</p><br/>";
        HTML += "<p class='Note'>" + result.Note + "</p><br/>";

        jQuery.each(result.PhotoSrc, function(i, val) {
            HTML += "<img class='PhotoSrc' src='/storage/photo/" + result.PhotoSrc[i] + "' ><br/>";
        });
        $('#record-data').prepend(HTML);                          
    });
});