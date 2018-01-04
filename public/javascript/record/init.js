$( document ).ready(function() {
    var id = window.location.search.substring(1);

    $.ajax({
        type: "GET",
        url : "/record/data",
        data: { recordid: id }
    }).done(function( searchResult ) {
        var result = $.parseJSON(searchResult);
        //console.log(result);
        
        HTML = "";
        HTML += "<a href='/records' class='back_to_records'>回前頁</a>";
        HTML += "<button class='edit-record'>編輯</button>";
        
        HTML += "<h2>物種名&nbsp:&nbsp;<class='OrganismName'>" + result.OrganismName + "</h2>";
        HTML += "<table><tr><td class='title'>季節</td><td class='Season'>" + result.Season + "</td><tr/>";
        HTML += "<tr><td class='title'>年齡</td><td class='Stage'>" + result.Stage + "</td><tr/>";
        HTML += "<tr><td class='title'>性狀</td><td class='Status'>" + result.Status + "</td><tr/>";
        HTML += "<tr><td class='title'>棲地</td><td class='Habitat'>" + result.Habitat + "</td><tr/>";
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
        
        $( ".edit-record" ).click(function(e) {
            $(".edit-div").css('display', 'block');
            
            
            $("input[name$='organismname']").attr("required", "true");
            $("input[name$='organismname']").val( result.OrganismName );
            $('select[name="season"] option[value="' + result.Season +'"').attr('selected', 'selected');
            $('select[name="stage"] option[value="' + result.Stage +'"').attr('selected', 'selected');
            $('select[name="tag"] option[value="' + result.Tag +'"').attr('selected', 'selected');
            $('textarea#textarea-status').val(result.Status);
            $('textarea#textarea-habitat').val(result.Habitat);
            $('textarea#textarea-note').val(result.Note);
            EDITHTML = '<input type="hidden" id="recordid" name="recordid"' + 'value="' + result.ID + '">';
            $('#edit-box').prepend(EDITHTML);

            $(".edit-div").css('display', 'block');
        });
        
        $( ".close-edit-div-button" ).click(function(e) {
            $(".edit-div").css("display","none");
        });
        
        $( "#edit-form" ).submit(function(e) {
            e.preventDefault();            
            
            var formData = new FormData(this);
            $.ajax({
                url : "/record/data",
                method: "PATCH",
                data: formData,
                contentType: false,
                processData: false,
                success:function(data, textStatus, jqXHR) {
                    var result = $.parseJSON(data);
                    $("#edit-form")[0].reset();
                    $(".edit-div").css("display","none");
                    location.reload();
                }
            });
        });
    });
});

window.onclick = function(event) {
    if (event.target == editDiv) {
        editDiv.style.display = "none";
    }
}