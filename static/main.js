$(document).ready(function() {
  var connection = new WebSocket(connectionType() + getCurrentUrl() + "/socket");

  connection.onopen = function(data) {
    console.log('connection established');
  };

  connection.onmessage = function(data) {
    var message = JSON.parse(data.data.toLowerCase());

    switch (message.action) {
      case "taskadded":
        addTask($('#tasks table tbody'), message.message);
        break;

      case "taskcompleted":
        var id = message.message._id || message.message.id;
        $('#' + id + ' td.status').text(message.message.status);
        break;
      default:
      console.log('Unrecoginised task');
    }
  };

  $.ajax({
    method: 'GET',
    url: '/tasks',
  }).done(function(data) {
    var taskTable = $('#tasks table tbody');

    data.forEach(function(task) {
      addTask(taskTable, task);
    });
  });

  $.ajax({
    method: 'GET',
    url: '/user',
    contentType: "application/json; charset=utf-8"
  }).done(function(data) {
    if (data.username != null) {
      toggleAuthButtons();
    }
  });

  $('#logout').click(function(e) {
    e.preventDefault();

    $.ajax({
      method: 'DELETE',
      url: '/logout',
      contentType: "application/json; charset=utf-8"
    }).done(function(data) {
      if (data.message.length > 0) {
        toggleAuthButtons();
      }
    }).fail(function(data, textStatus, errorThrown) {
    });
  })

  $("#loginForm").unbind("submit");
  $('#loginForm').submit(function(e) {
    var me = $(this);
    e.preventDefault();

    var username = $('#usernameLogin').val();
    var password = $('#passwordLogin').val();

    $.ajax({
      method: 'POST',
      url: '/login',
      contentType: "application/json; charset=utf-8",
      data: JSON.stringify({ "username": username, "password": password, "phone": 'phone' }),
      complete: function() {
            me.data('requestRunning', false);
        }
    }).done(function(data) {
      toggleAuthButtons();
      $('#modalLogin').modal('hide');
    }).fail(function(data, textStatus, errorThrown) {
      $('#modalLogin').modal('hide');
    });
  });

  $("#registerForm").unbind("submit");
  $('#registerForm').submit(function(e) {
    var me = $(this);
    e.preventDefault();

    var username = $('#usernameRegister').val();
    var password = $('#passwordRegister').val();
    var phone = $('#phoneRegister').val();

    $.ajax({
      method: 'POST',
      url: '/register',
      contentType: "application/json; charset=utf-8",
      data: JSON.stringify({ "username": username, "password": password, "phone": phone }),
      complete: function() {
            me.data('requestRunning', false);
        }
    }).done(function(data) {
      toggleAuthButtons();
      $('#modalRegister').modal('hide');
    }).fail(function(data, textStatus, errorThrown) {
      $('#modalRegister').modal('hide');
    });
  });

});

function addTask(taskTable, task) {
  var id = task._id || task.id;
  taskTable.append('<tr id="' + id + '">' +
    '<td>' + task.command + '</td>' +
    '<td>' + task.time + '</td>' +
    '<td>' + task.orderdetails.username + '</td>' +
    '<td>' + task.orderdetails.address + '</td>' +
    '<td>' + task.orderdetails.phone + '</td>' +
    '<td>' + task.orderdetails.paymenttype + '</td>' +
    '<td>' + task.orderlist + '</td>' +
    '<td>' + task.pizzerianame + '</td>' +
    '<td class="status">' + task.status + '</td>' +
  '</tr>');
}

function connectionType() {
  var docUrl = document.URL;
  var url;

  if (docUrl.indexOf('http://') > -1) {
    connectionType = 'ws://';
  } else if (docUrl.indexOf('https://') > -1) {
    connectionType = 'wss://';
  } else {
    connectionType = 'ws://';
  }

  return connectionType;
}

function getCurrentUrl() {
  var docUrl = document.URL;
  var url;

  if (docUrl.indexOf('http://') > -1) {
    url = docUrl.substring(7, docUrl.length - 1);
  } else if (docUrl.indexOf('https://') > -1) {
    url = docUrl.substring(8, docUrl.length - 1);
  } else {
    url = docUrl;
  }

  return url;
}

function toggleAuthButtons() {
  $('#register').toggleClass('hidden');
  $('#login').toggleClass('hidden');
  $('#logout').toggleClass('hidden');
}
