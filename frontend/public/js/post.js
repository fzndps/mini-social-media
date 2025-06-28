const selectImage = document.getElementById("dropzone");
const fileInput = document.getElementById("fileUpload");
const submitBtn = document.getElementById("submitPost");
const message = document.getElementById("message");
const postsContainer = document.getElementById("postsContainer");

let selectedFile = null;

// Klik pada dropzone membuka file picker
selectImage.addEventListener("click", () => fileInput.click());

// Saat file dipilih
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

// Saat submit ditekan
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

    // Bisa diganti sesuai respons API jika sudah ada backend
    const postData = await fetch("https:/localhost/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        name,
        password
      })
    });

    renderNewPost(postData);

    message.innerHTML = `<p class="text-green-600">Post berhasil ditambahkan!</p>`;

    // Reset form
    fileInput.value = "";
    selectedFile = null;
    document.getElementById("caption").value = "";
    selectImage.classList.remove("border-blue-500");
    selectImage.innerHTML = `Click to select image`;
  } catch (err) {
    message.innerHTML = `<p class="text-red-500">Error: ${err.message}</p>`;
  }
});

// Fungsi untuk render HTML post baru ke atas container
function renderNewPost(data) {
  const postHTML = `
    <div class="max-w-md mx-auto bg-white border border-gray-200 rounded-2xl shadow-sm overflow-hidden mb-6">
      <div class="flex items-center justify-between p-4">
        <div class="flex items-center">
          <img src="../src/assets/profile.jpeg" alt="User" class="w-10 h-10 rounded-full mr-3" />
          <div>
            <p class="text-sm font-semibold text-gray-900">${data.username}</p>
            <p class="text-xs text-gray-500">${data.location}</p>
          </div>
        </div>
        <button class="text-gray-500 hover:text-black">
          <svg class="w-5 h-5" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 12h.01M12 12h.01M18 12h.01" />
          </svg>
        </button>
      </div>
      <div class="w-full">
        <img src="${data.imageUrl}" alt="Post" class="w-full object-cover" />
      </div>
      <div class="flex items-center justify-between px-4 py-3">
        <div class="flex space-x-4">
          <button>
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
              stroke-width="1.5" stroke="currentColor" class="size-6">
              <path stroke-linecap="round" stroke-linejoin="round"
                d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.5 `
}