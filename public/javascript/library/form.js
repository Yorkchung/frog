$( "#library-frog" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/library/data",
        data: { label: "frog" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result);
        $( "#library-data" ).empty();
        jQuery.each(result.LibraryDatas, function(i, val) {
            var searchResultHTML = "<div>";
            searchResultHTML += "<table id='" + val.ID +  "' border='1'><tbody>";
            searchResultHTML += "<tr><td>科目</td><td class='dataFamily'>" + val.Family + "</td></tr>";
            searchResultHTML += "<tr><td>種類</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>性狀</td><td class='dataStatus'>" + val.Status + "</td></tr>";
            searchResultHTML += "<tr><td>棲地</td><td class='dataHabitat'>" + val.Habitat + "</td></tr>";
            searchResultHTML += "<tr><td>操作</td><td><button class='deleteDataButton' value='" + val.ID + "'>刪除</button><button class='modifyDataButton' value='" + val.ID + "'>編輯</button><button class='getRecordsByOrganismName' value='" + val.OrganismName + "'>相關記錄</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#library-data').prepend(searchResultHTML);
        });
    });
});

$( "#library-lepidoptera" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/library/data",
        data: { label: "lepidoptera" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result);
        $( "#library-data" ).empty();
        jQuery.each(result.LibraryDatas, function(i, val) {
            var searchResultHTML = "<div>";
            searchResultHTML += "<table id='" + val.ID +  "' border='1'><tbody>";
            searchResultHTML += "<tr><td>科目</td><td class='dataFamily'>" + val.Family + "</td></tr>";
            searchResultHTML += "<tr><td>種類</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>性狀</td><td class='dataStatus'>" + val.Status + "</td></tr>";
            searchResultHTML += "<tr><td>棲地</td><td class='dataHabitat'>" + val.Habitat + "</td></tr>";
            searchResultHTML += "<tr><td>操作</td><td><button class='deleteDataButton' value='" + val.ID + "'>刪除</button><button class='modifyDataButton' value='" + val.ID + "'>編輯</button><button class='getRecordsByOrganismName' value='" + val.OrganismName + "'>相關記錄</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#library-data').prepend(searchResultHTML);
        });
    });
});

$( "#library-plant" ).click(function(e) {
    $.ajax({
        type: "GET",
        url: "/library/data",
        data: { label: "plant" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result);
        $( "#library-data" ).empty();
        jQuery.each(result.LibraryDatas, function(i, val) {
            var searchResultHTML = "<div>";
            searchResultHTML += "<table id='" + val.ID +  "' border='1'><tbody>";
            searchResultHTML += "<tr><td>科目</td><td class='dataFamily'>" + val.Family + "</td></tr>";
            searchResultHTML += "<tr><td>種類</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td>性狀</td><td class='dataStatus'>" + val.Status + "</td></tr>";
            searchResultHTML += "<tr><td>棲地</td><td class='dataHabitat'>" + val.Habitat + "</td></tr>";
            searchResultHTML += "<tr><td>操作</td><td><button class='deleteDataButton' value='" + val.ID + "'>刪除</button><button class='modifyDataButton' value='" + val.ID + "'>編輯</button><button class='getRecordsByOrganismName' value='" + val.OrganismName + "'>相關記錄</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#library-data').prepend(searchResultHTML);
        });
    });
});

$( ".deleteDataButton" ).click(function(e) {
    var dataID = $(this).val();
    $("#deleteDataDiv").show();
    $('input#needDeleteDataID').val(dataID);
});

$( ".modifyDataButton" ).click(function(e) {
    var dataID = $(this).val();
    $("#modifyDataDiv").show();
    $('input#needModifyBookID').val(dataID);
    $("#submit-modify-button").attr("disabled", "true");

    var dataID = [{name:"dataID", value: dataID}];
    $.ajax({
        url : "modify-library-data",
        type: "POST",
        data : dataID,
        success:function(data, textStatus, jqXHR) {
            var result = $.parseJSON(data);
            if (result.BookTitle != "") {
                $('#submit-modify-button').removeAttr('disabled');
            } else {

            }
        },
        error: function(jqXHR, textStatus, errorThrown) {

        },
        always: function() {

        }
    });
});

$("#cancel-delete-button").click(function() {
    $("#deleteDataDiv").hide();
});

$("#cancel-modify-button").click(function() {
    $("#modifyDataDiv").hide();
});
