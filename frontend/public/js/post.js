    const selectImage = document.getElementById("selectImage");
    const fileInput = document.getElementById("fileUpload");
    const submitBtn = document.getElementById("submitPost");
    const message = document.getElementById("message");

    let selectedFile = null;

    // Klik pada area dropzone
    selectImage.addEventListener("click", () => fileInput.click());

    // Simpan file yang dipilih
    fileInput.addEventListener("change", () => {
      selectedFile = fileInput.files[0];
      if (selectedFile) {
        selectImage.classList.add("border-blue-500");
        selectImage.innerHTML = `
          <p class="text-sm text-gray-700">Selected: <strong>${selectedFile.name}</strong></p>
          <p class="text-xs text-gray-500 mt-1">Click to change</p>
        `;
      }
    });

    // Simulasi pengiriman post
    submitBtn.addEventListener("click", async () => {
      const caption = document.getElementById("caption").value.trim();

      if (!selectedFile || !caption) {
        message.innerHTML = `<p class="text-red-500">File dan caption wajib diisi.</p>`;
        return;
      }

      const formData = new FormData();
      formData.append("file", selectedFile);
      formData.append("caption", caption);

      try {
        const response = await fetch("http://127.0.0.1:3000/api/posts", {
          method: "POST",
          Authorization: `Bearer ${localStorage.getItem("token")}`,
          body: formData,
        });

        if (!response.ok) throw new Error("Gagal upload");

        message.innerHTML = `<p class="text-green-600">Post berhasil ditambahkan!</p>`;
        fileInput.value = "";
        selectedFile = null;
      } catch (err) {
        message.innerHTML = `<p class="text-red-500">Error: ${err.message}</p>`;
      }
});