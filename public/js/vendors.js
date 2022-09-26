let data = [];
$("document").ready(function () {
    // ajax call to get the data
    $.ajax({
        url: "/api/get-product",
        type: "GET",
        dataType: "json",
        success: function (data) {

            data = data;
            generateOptions(data);
        },
        error: function (err) {
            console.log(err);
        }
    });

    $("#check").on("change", function () {
        if (this.checked) {
            $("#product-name").css("display", "none");
            $("#select").css("display", "block");
        } else {
            $("#select").css("display", "none");
            $("#product-name").css("display", "block");
        }
    });

    $("#select").on("change", function () {
        if (this.value == "xyz") {
            $("#email").val("email ");
            $("#mobile").val("063546565");
            $("#address").val("lksdjflskdjf");
        }
    });

});

function generateOptions(data) {
    let html = "";
    data.forEach((item) => {
        html += `<option value="${item.Name}">${item.Name}</option>`;
    });
    $("#select").append(html);
}