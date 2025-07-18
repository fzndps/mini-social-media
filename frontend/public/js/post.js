document.addEventListener("DOMContentLoaded", () => {
  const token = localStorage.getItem("token");
  const form = document.getElementById("createPostForm");
  const message = document.getElementById("message");
  const imageInput = document.getElementById("imageInput");
  const dropzone = document.getElementById("dropzone");
  const imagePreview = document.getElementById("imagePreview");
  const previewImage = imagePreview.querySelector("img");

  // Cek token
  if (!token) {
    alert("Silakan login terlebih dahulu.");
    window.location.href = "index.html";
    return;
  }

  // Dragzone click
  dropzone.addEventListener("click", () => {
    imageInput.click();
  });

  // Preview image
  imageInput.addEventListener("change", () => {
    const file = imageInput.files[0];
    if (file && file.type.startsWith("image/")) {
      previewImage.src = URL.createObjectURL(file);
      imagePreview.classList.remove("hidden");
    } else {
      imagePreview.classList.add("hidden");
    }
  });

  // Submit
  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    message.textContent = "";
    message.className = "text-center text-sm";

    const formData = new FormData(form);
    const content = formData.get("content").trim();
    const imageFile = formData.get("image");

    // Validasi konten
    if (!content) {
      showError("Isi postingan tidak boleh kosong.");
      return;
    }

    // Validasi file gambar jika diupload
    if (imageFile && imageFile.size > 0 && !imageFile.type.startsWith("image/")) {
      showError("File yang diunggah harus berupa gambar.");
      return;
    }

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
        showError(`Gagal membuat postingan: ${errorText}`);
        return;
      }

      // Sukses
      window.location.href = "../public/dashboard.html";
    } catch (error) {
      console.error("Error:", error);
      showError("Terjadi kesalahan saat membuat postingan.");
    }
  });

  function showError(msg) {
    message.textContent = msg;
    message.className = "text-red-500 text-center text-sm mt-2";
  }
});
