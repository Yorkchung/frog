// disable enter's submit
$('input').keypress(function(e) {
    if(e.keyCode == 13) {
        e.preventDefault();
    }
});

$("#login-form").submit(function(e) {
    
    e.preventDefault();
        
    //$("#registerDiv").hide();

    var formData = $(this).serializeArray();
    
    $.ajax({
        url : "requestlogin",
        type: "POST",
        data : formData,
        success:function(data, textStatus, jqXHR) {
            console.log("post success");
            var result = $.parseJSON(data);
            if (result.LoginStatus == true) {
                $(location).attr('href', '/');
            }
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log("error");
        }
    });
});