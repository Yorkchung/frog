//要加上傳中離開頁面會跳遺失檔案警告的浮動視窗

$("#upload-form").submit(function(e) {
    e.preventDefault();
    //$("#gogogo").prop('disabled', true);

    var formURL = $(this).attr("action");
    var formData = new FormData(this);
    if (formData.get("recorddate") == "") {
        formData.delete("recorddate");
    }
    console.log(formData);
    $.ajax({
        url : formURL,
        type: "POST",
        data : formData,
        processData: false,
        contentType: false,
        success:function(data, textStatus, jqXHR) {
            var result = $.parseJSON(data);
            //$("#gogogo").prop('disabled', false);
            $("#upload-form")[0].reset();
        },
        error: function(jqXHR, textStatus, errorThrown) {
            console.log("ajax post error", textStatus);
        },
        always: function() {

        }
    });
});
