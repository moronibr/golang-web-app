$(document).ready(function () {
    $(".login-btn").click(function () {
        $("#loginModal").css("display", "flex");
    });

    $(".close").click(function () {
        $("#loginModal").css("display", "none");
    });

    // Fecha o modal se clicar fora dele
    $(window).click(function (event) {
        if ($(event.target).is("#loginModal")) {
            $("#loginModal").css("display", "none");
        }
    });
});
