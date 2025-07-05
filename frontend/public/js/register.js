const registerForm = document.getElementById("registerForm");

registerForm.addEventListener("submit", async function (e) {
  e.preventDefault();

  // Ambil nilai input
  const username = document.getElementById("registerUsername").value.trim();
  const email = document.getElementById("registerEmail").value.trim();
  const password = document.getElementById("registerPassword").value.trim();
  const confirmPassword = document.getElementById("registerConfirmPassword").value.trim();

  // Reset error
  let valid = true;

  // Validasi nama
  if (username === "") {
    alert("Username tidak boleh kosong");
    valid = false;
  }

  // Validasi email
  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!emailPattern.test(email)) {
    alert("Format email tidak valid");
    valid = false;
  }

  // Validasi password
  if (password.length < 8) {
    alert("Password minimal 8 karakter");
    valid = false;
  }

  // Validasi konfirmasi password
  if (password !== confirmPassword) {
    alert("Password dan konfirmasi password tidak cocok");
    valid = false;
  }

  if (!valid) return;

  // Kirim data ke backend
  try {
    const response = await fetch("http://127.0.0.1:3000/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        username,
        email,
        password
      })
    });

    const data = await response.json();

    if (response.ok) {
      alert("Registrasi berhasil!");
      // Redirect
      window.location.href = "../public/login.html";
    } else {
      alert(data.message || "Registrasi gagal");
    }
  } catch (err) {
    console.error(err);
    alert("Gagal terhubung ke server.");
  }
});
