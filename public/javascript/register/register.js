// disable enter's submit
$('input').keypress(function(e) {
    if(e.keyCode == 13) {
        e.preventDefault();
    }
});

$("#register-form").submit(function(e) {
    
    e.preventDefault();
        
    //$("#registerDiv").hide();

    var formData = $(this).serializeArray();
    
    $.ajax({
        url : "requestregister",
        type: "POST",
        data : formData,
        success:function(data, textStatus, jqXHR) {
            //var result = $.parseJSON(data);
            console.log("post success");
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log("error");
        }
    });
});