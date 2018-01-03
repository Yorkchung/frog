$( document ).ready(function() {
    $.ajax({
        type: "GET",
        url: "/library/data",
        data: { label: "" }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result);
        $( "#library-data" ).empty();
        jQuery.each(result.LibraryDatas, function(i, val) {
            var searchResultHTML = "<div>";
            searchResultHTML += "<table id='" + val.ID +  "' border='1'><tbody>";
            searchResultHTML += "<tr><td class='title'>科目</td><td class='dataFamily'>" + val.Family + "</td></tr>";
            searchResultHTML += "<tr><td class='title'>種類</td><td class='dataOrganismName'>"+ val.OrganismName +"</td></tr>";
            searchResultHTML += "<tr><td class='title'>性狀</td><td class='dataStatus'>" + val.Status + "</td></tr>";
            searchResultHTML += "<tr><td class='title'>棲地</td><td class='dataHabitat'>" + val.Habitat + "</td></tr>";
            searchResultHTML += "<tr><td class='title'>操作</td><td><button class='deleteDataButton' value='" + val.ID + "'>刪除</button><button class='modifyDataButton' value='" + val.ID + "'>編輯</button><button class='getRecordsByOrganismName' value='" + val.OrganismName + "'>相關記錄</button></td></tr>";
            searchResultHTML += "</tbody></table>";
            searchResultHTML += "</div></br>";
            searchResultHTML += '';
            $('#library-data').prepend(searchResultHTML);
        });
    });
});