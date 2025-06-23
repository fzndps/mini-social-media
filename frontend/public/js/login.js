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

    try {
        const response = await fetch("http://127.0.0.1:3000/auth/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username, password })
        });

        const data = await response.json();
        console.log(data.data.access_token)

        if (response.ok) {
            alert("Login berhasil!");
            // Simpan token (kalau dikasih)
            localStorage.setItem("token", data.data.access_token);
            // Redirect
            window.location.href = "../public/dashboard.html";
        } else {
            alert(data.message || "Login gagal");
        }
    } catch (error) {
        alert("Gagal terhubung ke server.");
        console.error(error);
    }
});