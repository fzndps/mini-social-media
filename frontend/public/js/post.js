const token = localStorage.getItem("token");

if (!token) {
  alert("Silakan login terlebih dahulu.");
  window.location.href = "login.html";
}

document.getElementById("createPostForm").addEventListener("submit", async function (e) {
  e.preventDefault();

  const form = e.target;
  const formData = new FormData(form);

  try {
    const response = await fetch("http://localhost:3000/api/posts", {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`
      },
      body: formData
    });

    const result = await response.json();
    if (response.ok) {
      alert("Post berhasil dibuat!")
    } else {
      alert("Gagal buat post: " + result.message)
    }

    window.location.href = "../public/dashboard.html"
  } catch (error) {
    console.error("❌ Error:", error);
    alert("Terjadi kesalahan saat membuat post.");
  }
});