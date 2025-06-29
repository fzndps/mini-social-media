const selectImage = document.getElementById("selectImage");
const fileInput = document.getElementById("fileUpload");
const submitBtn = document.getElementById("submitPost");
const message = document.getElementById("message");
const token = localStorage.getItem("token")

let selectedFile = null;

selectImage.addEventListener("click", () => fileInput.click());

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

submitBtn.addEventListener("click", async (e) => {
  e.preventDefault();

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
      headers: {
        Authorization: `Bearer ${token}`
      },
      body: formData,
    });

    const newPost = await response.json();
    console.log(newPost.data)
    console.log(newPost)
    if (response.ok) {
      alert('Berhasil Upload')
    } else {
      alert('Gagal Upload', newPost.message)
    }

    // Simpan data postingan di localStorage
    localStorage.setItem("latestPost", JSON.stringify(newPost));

    // Redirect ke dashboard
    window.location.href = "dashboard.html";
  } catch (error) {
    console.error("Upload error:", error);
    message.innerHTML = `<p class="text-red-500">Error: ${error.message}</p>`;
  }
});
