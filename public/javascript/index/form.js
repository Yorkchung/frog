$("#loginform").submit(function(e) {
    e.preventDefault();
    var postData = $(this).serializeArray();
    var formURL = $(this).attr("action");
    $.ajax({
        url : formURL,
        type: "POST",
        data : postData,
        success:function(data, textStatus, jqXHR) {
            var result = $.parseJSON(data);
            console.log(result.LoginStatus);            

            if (result.LoginStatus == false) {

            } else {
                document.location.reload();                
            }
        },
        error: function(jqXHR, textStatus, errorThrown) {
            alert("ajax pust login data error");
        }
    });
});