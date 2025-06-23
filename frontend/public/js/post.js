    const selectImage = document.getElementById("selectImage");
    const fileInput = document.getElementById("fileUpload");
    const submitBtn = document.getElementById("submitPost");
    const message = document.getElementById("message");

    let selectedFile = null;

    // Klik pada area dropzone
    dropzone.addEventListener("click", () => fileInput.click());

    // Simpan file yang dipilih
    fileInput.addEventListener("change", () => {
      selectedFile = fileInput.files[0];
      if (selectedFile) {
        dropzone.classList.add("border-blue-500");
        dropzone.innerHTML = `
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
        const res = await fetch("https://your-api-endpoint/posts", {
          method: "POST",
          body: formData,
        });

        if (!res.ok) throw new Error("Gagal upload");

        message.innerHTML = `<p class="text-green-600">Post berhasil ditambahkan!</p>`;
        fileInput.value = "";
        selectedFile = null;
      } catch (err) {
        message.innerHTML = `<p class="text-red-500">Error: ${err.message}</p>`;
      }
});