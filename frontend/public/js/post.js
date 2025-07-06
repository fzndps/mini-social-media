document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  const form = document.getElementById("createPostForm");
  const message = document.getElementById("message");

  if (!token) {
    alert("Silakan login terlebih dahulu.");
    window.location.href = "login.html";
    return;
  }

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const formData = new FormData(form);

    try {
      const response = await fetch("http://127.0.0.1:3000/api/posts", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: formData,
      });

      if (!response.ok) {
        const errorText = await response.text();
        message.textContent = `Gagal membuat postingan: ${errorText}`;
        message.className = "text-red-500";
        return;
      }
const selectImage = document.getElementById("dropzone");
const fileInput = document.getElementById("fileUpload");
const submitBtn = document.getElementById("submitPost");
const message = document.getElementById("message");
const postsContainer = document.getElementById("postsContainer");


      // Redirect setelah berhasil
      window.location.href = "dashboard.html";
    } catch (error) {
      console.error("Error saat membuat postingan:", error);
      message.textContent = "Terjadi kesalahan saat membuat postingan.";
      message.className = "text-red-500";
    }
  });

