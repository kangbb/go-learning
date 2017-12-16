$(document).ready(function() {
    $.ajax({
        url: "/home/date"
    }).then(function(data) {
       $('.time').append(data);
    });
});
