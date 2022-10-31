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
      $("#cs-name").css("display", "none");
      $("#select").css("display", "block");
      $("#fist-select").prop("selected", true);
      $("#email").prop("disabled", true);
      $("#mobile").prop("disabled", true);
      $("#address").prop("disabled", true);

    } else {
      $("#select").css("display", "none");
      $("#cs-name").css("display", "block");
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
    let email = mapData.get(id).email;
    let mobile = mapData.get(id).mobile;
    let address = mapData.get(id).address;
    $("#email").val(email);
    $("#mobile").val(mobile);
    $("#address").val(address);
  });


  $("body").on("change", ".product-option", function () {
    let trId = $(this).closest("tr").attr("id");
    let id = $(this).val();
    id = parseInt(id);
    let price = data2.find((item) => item.id === id).sell_amount;
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

  // calculate due amount regarding paid amount
  $("#paidAmount").on("keyup change", function () {
    let paidAmount = $("#paidAmount").val();
    let subtotal = $("#subTotal").text();
    let dueAmount = subtotal - paidAmount;
    $("#paid").text(paidAmount);
    $("#due").text(dueAmount);
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

      html += `<option class="product-option" value="${item.id}">${item.product_name}</option>`
    });
    // priceId
    let priceId = "price" + rowId;
    html += `</select> </td>
        <td><input class="form-control q" type="number" value="1" name="" id="q` + rowId + `"></td>
        <td><input class="form-control price" disabled type="text" name="" id="` + priceId + `"></td>
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


  // prepare the data to send to the server
  $("#submit").on("click", function (event) {
    event.preventDefault();

    // customer name check box is checked or not
    let check = $("#check").is(":checked");

    // if checked then get the customer id from the select option
    let customerName;

    if (check) {
      customerName = $("#select option:selected").text();
    } else {
      customerName = $("#cs-name").val();
    }

    let customerEmail = $("#email").val();
    let customerMobile = $("#mobile").val();
    let customerAddress = $("#address").val();
    let customerType = $("#check").is(":checked");
    let customer = {
      customer_name: customerName,
      email: customerEmail,
      mobile: customerMobile,
      address: customerAddress,
      customer_type: customerType,
    };
    let products = [];
    let product = {};
    let rows = $("#table tr");
    rows.each(function (index, item) {
      if (index != 0) {
        let productId = $(item).find(".product-option").val();
        let quantity = $(item).find(".q").val();
        let price = $(item).find(".price").val();
        let total = $(item).find(".total").val();
        product = {
          product_id: productId,
          quantity: quantity,
          price: price,
          total: total,
        };
        products.push(product);
      }
    });
    let paidAmount = $("#paidAmount").val();
    let data = {
      customer: customer,
      products: products,
      paid_amount: paidAmount,
    };
    // data convert to json
    let jsonData = JSON.stringify(data);
    // send data to the server
    console.log(jsonData);
    console.log(data);
    // $.ajax({
    //   url: "/api/save-order",
    //   type: "POST",
    //   dataType: "json",
    //   data: data,
    //   success: function (data) {
    //     console.log(data);
    //   },
    //   error: function (err) {
    //     console.log(err);
    //   },
    // });
  });




});

function generateOptions(data) {
  let html = "";
  html +=
    "<option id='fist-select' value='none' selected>Select an Option</option>";
  data.forEach((item3) => {
    html += `<option value="${item3.id}">${item3.name}</option>`;
  });
  $("#select").append(html);
}

function generateProductOptions(data) {
  let html = "";
  html +=
    "<option id='fist-select' value='none' selected>Select an Option</option>";
  data.forEach((item) => {
    html += `<option value="${item.id}">${item.product_name}</option>`;
  });
  $("#product-option").append(html);
}

function generateMap(data) {
  let tem = new Map();
  for (let i = 0; i < data.length; i++) {
    tem.set(data[i].id, data[i]);
  }
  return tem;
}