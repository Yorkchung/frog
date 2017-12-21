$( "#records-frog" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/records/data",
        data: { searchtype: "tag", keyword:"frog" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        console.log(result);
        $( "#records-data" ).empty();
        jQuery.each(result.Records, function(i, val) {
            var searchResultHTML = "<div>";
            searchResultHTML += "<table id='" + val.ID +  "' border='1'><tbody>";
            searchResultHTML += "<tr><td>種類</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>性狀</td><td class='dataStatus'>" + val.Status + "</td></tr>";
            searchResultHTML += "<tr><td>棲地</td><td class='dataHabitat'>" + val.Habitat + "</td></tr>";
            searchResultHTML += "<tr><td>操作</td><td><button class='getRecordByRecordID' value='" + val.ID + "'>細節</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#records-data').prepend(searchResultHTML);
        });

        $( ".getRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            $.ajax({
                type: "GET",
                url : "/record/data",
                data: { recordid: id }
            }).done(function( searchResult ) {;
                var result = $.parseJSON(searchResult);
                console.log(result);
            });
        });
        
    });
});

$( "#records-lepidoptera" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/records/data",
        data: { searchtype: "tag", keyword:"lepidoptera" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result);
        $( "#records-data" ).empty();
        jQuery.each(result.Records, function(i, val) {
            var searchResultHTML = "<div>";
            searchResultHTML += "<table id='" + val.ID +  "' border='1'><tbody>";
            searchResultHTML += "<tr><td>種類</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>日期</td><td class='dataStatus'>" + val.CrateTime + "</td></tr>";
            /*
                searchResultHTML += "<tr><td>性狀</td><td class='dataStatus'>" + val.Status + "</td></tr>";
                searchResultHTML += "<tr><td>棲地</td><td class='dataHabitat'>" + val.Habitat + "</td></tr>";
            */
            searchResultHTML += "<tr><td>操作</td><td><button class='getRecordByRecordID' value='" + val.ID + "'>細節</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#records-data').prepend(searchResultHTML);
        });

        $( ".getRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            $.ajax({
                type: "GET",
                url : "/record/data",
                data: { recordid: id }
            }).done(function( searchResult ) {;
                var result = $.parseJSON(searchResult);
                console.log(result); //單筆記錄
            });
        });

    });
});

$( "#records-plant" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/records/data",
        data: { searchtype: "tag", keyword:"plant" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result);
        $( "#records-data" ).empty();
        jQuery.each(result.Records, function(i, val) {
            var searchResultHTML = "<div>";
            searchResultHTML += "<table id='" + val.ID +  "' border='1'><tbody>";
            searchResultHTML += "<tr><td>種類</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>性狀</td><td class='dataStatus'>" + val.Status + "</td></tr>";
            searchResultHTML += "<tr><td>棲地</td><td class='dataHabitat'>" + val.Habitat + "</td></tr>";
            searchResultHTML += "<tr><td>操作</td><td><button class='getRecordByRecordID' value='" + val.ID + "'>細節</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#records-data').prepend(searchResultHTML);
        });
        
        $( ".getRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            $.ajax({
                type: "GET",
                url : "/record/data",
                data: { recordid: id }
            }).done(function( searchResult ) {;
                var result = $.parseJSON(searchResult);
                console.log(result);
            });
        });

    });
});