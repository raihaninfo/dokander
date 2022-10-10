let data = [];
let mapData = new Map();
$("document").ready(function () {
  // ajax call to get the data
  $.ajax({
    url: "/api/get-customers",
    type: "GET",
    dataType: "json",
    success: function (data) {
      data = data;
      generateOptions(data);
      mapData = generateMap(data);
    },
    error: function (err) {
      console.log(err);
    },
  });
  // ajax call to get the product data
  $.ajax({
    url: "/api/get-products",
    type: "GET",
    dataType: "json",
    success: function (data) {
      data2 = data;
      generateProductOptions(data);
    },
    error: function (err) {
      console.log(err);
    },
  });

  $("#check").on("change", function () {
    if (this.checked) {
      $("#product-name").css("display", "none");
      $("#select").css("display", "block");
      // html attribute to disable the button
      $("#fist-select").prop("selected", true);

      console.log($("#select").val());
    } else {
      $("#select").css("display", "none");
      $("#product-name").css("display", "block");
      $("#email").val("");
      $("#mobile").val("");
      $("#address").val("");
    }
  });

  $("#select").on("change", function () {
    let id = $(this).val();
    id = parseInt(id);
    let email = mapData.get(id).Email;
    let mobile = mapData.get(id).Mobile;
    let address = mapData.get(id).Address;
    $("#email").val(email);
    $("#mobile").val(mobile);
    $("#address").val(address);
  });

  // add button on click event

  $("#add").on("click", function (event) {
    event.preventDefault();
    let rowId = $("#table tr").length + 1;
    // console.log(data2);

    rowId = rowId + 1;
    html =
      `<tr id="` +
      rowId +
      `">
        <td>
            <select class="form-control" name="" id="">
            "<option id='fist-select' value='none' selected>Select an Option</option>";
            `
    data2.forEach((item) => {

      html += `<option value="${item.Id}">${item.ProductName}</option>
            `
    });
    html += `</select> </td>
        <td><input class="form-control" type="text" name="" id=""></td>
        <td><input class="form-control" type="text" name="" id=""></td>
        <td><input class="form-control" type="text" name="" id=""></td>
        <td><input class="form-control" type="text" name="" id=""></td>
        <td><button class="btn btn-danger"><i class="bi bi-x-circle"></i></button></td>
    </tr>`;

    $("#table").append(html);
  });
  //    // delete button on click event
  $("#table").on("click", ".btn-danger", function () {
    $(this).closest("tr").remove();
  });
});

function generateOptions(data) {
  let html = "";
  html +=
    "<option id='fist-select' value='none' selected>Select an Option</option>";
  data.forEach((item) => {
    html += `<option value="${item.Id}">${item.Name}</option>`;
  });
  $("#select").append(html);
}

function generateProductOptions(data) {
  let html = "";
  html +=
    "<option id='fist-select' value='none' selected>Select an Option</option>";
  data.forEach((item) => {
    html += `<option value="${item.Id}">${item.ProductName}</option>`;
  });
  $("#product-option").append(html);
}

function generateMap(data) {
  let tem = new Map();
  for (let i = 0; i < data.length; i++) {
    tem.set(data[i].Id, data[i]);
  }
  return tem;
}