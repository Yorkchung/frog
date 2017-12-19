$( document ).ready(function() {
    
    var pathname = document.location.pathname;
    recordData = pathname.split("/", 3);
    
    $.post( "/search-record-by-record-id", { recordid: recordData[2] }).done(function( data ) {
        var result = $.parseJSON(data);
        console.log(result);

        $('#library-data').prepend(searchResultHTML);

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

    $("#cancel-delete-button").click(function() {
        $("#deleteDataDiv").hide();
    });

    $("#cancel-modify-button").click(function() {
        $("#modifyDataDiv").hide();
    });

});