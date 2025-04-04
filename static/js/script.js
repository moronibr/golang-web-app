$(document).ready(function () {
    // Buscar usuários existentes
    function carregarUsuarios() {
      $.get("http://localhost:8000/users", function (data) {
        let html = "<ul>";
        data.forEach(function (user) {
          html += `<li>${user.nome} (${user.email})</li>`;
        });
        html += "</ul>";
        $("#user-list").html(html);
      });
    }
  
    carregarUsuarios();
  
    // Enviar novo usuário
    $("#user-form").submit(function (e) {
      e.preventDefault();
  
      const nome = $("#nome").val();
      const email = $("#email").val();
  
      $.ajax({
        url: "http://localhost:8000/users",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({ nome, email }),
        success: function () {
          alert("Usuário adicionado!");
          $("#user-form")[0].reset();
          carregarUsuarios();
        },
      });
    });
  });
  