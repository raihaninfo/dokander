let data = [];
let data2 = [];
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
      $("#fist-select").prop("selected", true);
      $("#email").prop("disabled", true);
      $("#mobile").prop("disabled", true);
      $("#address").prop("disabled", true);

    } else {
      $("#select").css("display", "none");
      $("#product-name").css("display", "block");
      $("#email").prop("disabled", false);
      $("#mobile").prop("disabled", false);
      $("#address").prop("disabled", false);
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


  $("body").on("change", ".product-option", function () {
    let trId = $(this).closest("tr").attr("id");
    let id = $(this).val();
    id = parseInt(id);
    let price = data2.find((item) => item.Id === id).SellAmount;
    $("#price" + trId).val(price);
    let quantity = 1;
    let total = price * quantity;
    $("#total" + trId).val(total);
    totalSum();
  });

  // keyup and change event for quantity
  $("body").on("keyup change", ".q", function () {
    let trId = $(this).closest("tr").attr("id");
    let price = $("#price" + trId).val();
    let quantity = $(this).val();
    let total = price * quantity;
    $("#total" + trId).val(total);
    totalSum();
  });

  // total sum of all the rows
  function totalSum() {
    let sum = 0;
    $(".total").each(function () {
      sum += +$(this).val();
    });
    $("#subTotal").text(sum);
  }

  // add button on click event
  $("#add").on("click", function (event) {
    event.preventDefault();
    // last row id
    let lastRowId = $("#table tr:last").attr("id");
    if (lastRowId == undefined) {
      lastRowId = "0";
    }
    let lastRowIdInt = parseInt(lastRowId);
    rowId = lastRowIdInt + 1;

    html =
      `<tr id="` + rowId + `">
        <td>
            <select class="form-control product-option" name="" id="` + rowId + 100 + `">
            <option id='fist-select' value='none' selected>Select an Option</option>`
    data2.forEach((item) => {

      html += `<option class="product-option" value="${item.Id}">${item.ProductName}</option>`
    });
    // priceId
    let priceId = "price" + rowId;
    html += `</select> </td>
        <td><input class="form-control q" type="text" value="1" name="" id="q` + rowId + `"></td>
        <td><input class="form-control price" disabled type="text" name="" id="` + priceId + `"></td>
        <td><input class="form-control" type="text" name="" id=""></td>
        <td><input class="form-control total" disabled type="text" name="" id="total` + rowId + `"></td>
        <td><button class="btn btn-danger"><i class="bi bi-x-circle"></i></button></td>
    </tr>`;

    $("#table").append(html);
  });
  //    // delete button on click event
  $("#table").on("click", ".btn-danger", function () {
    $(this).closest("tr").remove();
    totalSum();
  });
});

function generateOptions(data) {
  let html = "";
  html +=
    "<option id='fist-select' value='none' selected>Select an Option</option>";
  data.forEach((item3) => {
    html += `<option value="${item3.Id}">${item3.Name}</option>`;
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