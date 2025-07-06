const loginForm = document.getElementById("loginForm");

loginForm.addEventListener("submit", async function (e) {
  e.preventDefault();

  const username = document.getElementById("loginUsername").value.trim();
  const password = document.getElementById("loginPassword").value.trim();

  if (username.length < 5) {
    alert("Username minimal 5 karakter");
    return;
  }
  if (password.length < 8) {
    alert("Password minimal 8 karakter");
    return;
  }


    }
  } catch (error) {
    alert("Gagal terhubung ke server.");
    console.error(error);
  }
});
