$(document).ready(function() {
  var connection = new WebSocket(connectionType() + getCurrentUrl() + "/socket");

  connection.onopen = function(data) {
    console.log('connection established');
  };

  connection.onmessage = function(data) {
    console.log(data.data);
  };

  $.ajax({
    method: 'GET',
    url: '/tasks',
  }).done(function(data) {
    console.log(data);
    var taskTable = $('#tasks table tbody');

    data.forEach(function(task) {
      taskTable.append('<tr id="' + task._id + '">' +
        '<td>' + task.command + '</td>' +
        '<td>' + task.time + '</td>' +
        '<td>' + task.orderdetails.username + '</td>' +
        '<td>' + task.orderdetails.address + '</td>' +
        '<td>' + task.orderdetails.paymenttype + '</td>' +
        '<td>' + task.orderlist + '</td>' +
        '<td>' + task.pizzerianame + '</td>' +
        '<td class="status">' + task.status + '</td>' +
      '</tr>');
    });
  });

});

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
