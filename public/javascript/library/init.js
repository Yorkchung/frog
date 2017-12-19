$( document ).ready(function() {
    
    var pathname = document.location.pathname;
    lebelData = pathname.split("/", 3);
    
    $.post( "/search-library-by-label", { label: lebelData[2] }).done(function( data ) {
        var result = $.parseJSON(data);
        //console.log(data);

        jQuery.each(result.LibraryDatas, function(i, val) {
            //console.log(i, val);
            //console.log(result.LibraryDatas[i].ID);
            //console.log(result.LibraryDatas[i].ID)
            //console.log(result.LibraryDatas[i].Family)
            //console.log(result.LibraryDatas[i].OrganismName)
            //console.log(result.LibraryDatas[i].Status)
            //console.log(result.LibraryDatas[i].Habitat)

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

            $( ".getRecordsByOrganismName" ).click(function(e) {
                var OrganismName = $(this).val();
                $.ajax({
                    type: "POST",
                    url: "/search-records-specify-organismname",
                    data: { organismname: OrganismName }
                }).done(function( searchResult ) {
                    var recordData = $.parseJSON(searchResult);
                    console.log(recordData);
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

        });
    });

    $("#cancel-delete-button").click(function() {
        $("#deleteDataDiv").hide();
    });

    $("#cancel-modify-button").click(function() {
        $("#modifyDataDiv").hide();
    });

});