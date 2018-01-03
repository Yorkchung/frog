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
            searchResultHTML += "<tr><td>上傳時間</td><td class='dataStatus'>" + val.CrateTime + "</td></tr>";
            searchResultHTML += "<tr><td>物種名</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>操作</td><td><button class='getRecordByRecordID' value='" + val.ID + "'>細節</button><button class='deleteRecordByRecordID' value='" + val.ID + "'>刪除</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#records-data').prepend(searchResultHTML);
        });

        $( ".getRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            location.href = "/record?" + id;
        });
        
        $( ".deleteRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            $.ajax({
                url : "/record/data" + '?' + $.param({"recordid" : id}),
                type: "DELETE"
            }).done(function( deleteResult ) {;
                 var result = $.parseJSON(deleteResult);
                 if (result.DeleteStatus == true) {
                     //location.reload();
                 }
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
            searchResultHTML += "<tr><td>上傳時間</td><td class='dataStatus'>" + val.CrateTime + "</td></tr>";
            searchResultHTML += "<tr><td>物種名</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>操作</td><td><button class='getRecordByRecordID' value='" + val.ID + "'>細節</button><button class='deleteRecordByRecordID' value='" + val.ID + "'>刪除</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#records-data').prepend(searchResultHTML);
        });

        $( ".getRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            location.href = "/record?" + id;
        });
        
        $( ".deleteRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            $.ajax({
                url : "/record/data" + '?' + $.param({"recordid" : id}),
                type: "DELETE"
            }).done(function( deleteResult ) {;
                 var result = $.parseJSON(deleteResult);
                 if (result.DeleteStatus == true) {
                     //location.reload();
                 }
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
            searchResultHTML += "<tr><td>上傳時間</td><td class='dataStatus'>" + val.CrateTime + "</td></tr>";
            searchResultHTML += "<tr><td>物種名</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>操作</td><td><button class='getRecordByRecordID' value='" + val.ID + "'>細節</button><button class='deleteRecordByRecordID' value='" + val.ID + "'>刪除</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#records-data').prepend(searchResultHTML);
        });
        
        $( ".getRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            location.href = "/record?" + id;
        });
        
        $( ".deleteRecordByRecordID" ).click(function(e) {
            id = $(this).val();
            $.ajax({
                url : "/record/data" + '?' + $.param({"recordid" : id}),
                type: "DELETE"
            }).done(function( deleteResult ) {;
                 var result = $.parseJSON(deleteResult);
                 if (result.DeleteStatus == true) {
                     //location.reload();
                 }
            });
        });

    });
});

/*
$( "#close-single-record-data-div-button" ).click(function(e) {
    $( "#single-record-data" ).hide();
});
*/