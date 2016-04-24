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

});

function addTask(taskTable, task) {
  var id = task._id || task.id;
  taskTable.append('<tr id="' + id + '">' +
    '<td>' + task.command + '</td>' +
    '<td>' + task.time + '</td>' +
    '<td>' + task.orderdetails.username + '</td>' +
    '<td>' + task.orderdetails.address + '</td>' +
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
